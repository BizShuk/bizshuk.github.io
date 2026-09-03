# MyCareersFuture API — 實測記錄 (Verified Surface)

`api.mycareersfuture.gov.sg` 是 `www.mycareersfuture.gov.sg` 前端直接呼叫的 backend。
`不需要 API key`, 不需要 cookie, 未帶任何憑證即回 `200`。代價是它`不是官方對外契約`:
沒有公開的 OpenAPI 文件 (`/openapi.json`, `/swagger.json`, `/api-docs` 皆 404),
端點與欄位可能無預警改版, 也可能隨時加上流量限制。

本檔記錄的每一條都以實際請求驗證過, 驗證日期 `2026-09-03`。

## 端點 (Endpoints)

| 端點                                    | 方法   | 授權      | 用途                                                |
| --------------------------------------- | ------ | --------- | --------------------------------------------------- |
| `/v2/search`                            | `POST` | 免         | 關鍵字與條件搜尋。`不含` description                |
| `/v2/jobs`                              | `GET`  | 免         | 全站職缺列表, 依張貼日倒序。`含` description        |
| `/v2/jobs/{uuid}`                       | `GET`  | 免         | 單一職缺全文。`唯一`拿得到 JD 全文的定點路徑        |
| `/v2/jobs/{uuid}/screening-questions`   | `GET`  | 免         | 申請篩選題。多數職缺回 `[]`                         |
| `/v2/companies`                         | `GET`  | 免         | 雇主列表, 可依 `name` 模糊查詢。全站 151,211 家     |
| `/v2/companies/{uen}`                   | `GET`  | 免         | 單一雇主。UEN 為新加坡統一企業號                    |
| `/v2/skills`                            | `GET`  | 免         | 回 `[]`, 未找到可用的查詢參數                        |
| `/v2/companies/{uen}/jobs`              | `GET`  | `需授權`   | 回 `401`                                            |
| `/v2/companies/{uen}/addresses`         | `GET`  | `需授權`   | 回 `401`                                            |
| `/v2/companies/{uen}/schemes`           | `GET`  | `需授權`   | 回 `401`                                            |

`探索方法`: 每個回應都帶 HATEOAS 的 `_links` 區塊, 那是唯一可靠的 API 表面線索;
猜路徑幾乎都是 404。另有一層 OpenAPI request validator, 送錯型別會逐欄列出期望型別,
可用來反推 schema。

## 兩段式抓取 (Two-stage Fetch)

`/v2/search` 的 `results[].description` `恆為空字串`。這不是欄位遺漏, 是端點設計:
要 JD 全文一定得拿 `uuid` 再打一次 `/v2/jobs/{uuid}`。`mcf fetch` 就是這條流程的封裝。

## 搜尋 body 欄位 (POST /v2/search)

驗證器只認下列欄位, 送出未知欄位不報錯也不生效。

| 欄位                       | 型別       | 生效 | 說明                                                         |
| -------------------------- | ---------- | ---- | ------------------------------------------------------------ |
| `search`                   | `string`   | 是   | 關鍵字                                                       |
| `sessionId`                | `string`   | 是   | 追蹤用, 對結果無影響                                          |
| `salary`                   | `integer`  | 是   | 月薪下限                                                     |
| `categories`               | `[]string` | 是   | enum, 例 `Information Technology`                            |
| `positionLevels`           | `[]string` | 是   | enum, 例 `Manager`, `Senior Management`                      |
| `employmentTypes`          | `[]string` | 是   | enum, 例 `Full Time`, `Permanent`                            |
| `postingCompany`           | `[]string` | 是   | enum, `只有` `Direct` 與 `Third Party` 兩值                  |
| `flexibleWorkArrangements` | `[]int`    | 是   | `以 id 而非名稱`篩選, 見下表                                 |
| `sortBy`                   | `[]string` | 是   | enum, `只有` `new_posting_date` 與 `min_monthly_salary`      |
| `schemes`                  | `boolean`  | `否` | 送 true 或 false 結果數完全相同, 在公開路徑上無作用          |
| `uen`                      | `string`   | `否` | 通過驗證但不篩選, 結果等同未帶                               |
| `jobStatuses`              | `[]int`    | `陷阱` | 一旦帶上, 關鍵字與其餘所有篩選`全部失效`, 回應也不再帶 `countWithoutFilters` |

`flexibleWorkArrangements` id 對照: `1` Flexi-Hours, `2` Telecommuting,
`3` Employees Choice of Days Off, `4` Staggered Time, `5` Compressed Work Schedule,
`6` Creative Scheduling。

`sortBy` 留白代表以相關性排序, 那是預設行為且沒有對應的 enum 字串;
送 `relevancy` 會被擋成 `400`。

## 查詢參數 (Query Parameters)

- 全部分頁端點: `limit` (integer, `上限 100`), `page` (integer, 由 0 起算)
- `/v2/jobs`: 另有 `salary` (integer, 月薪下限)
- `/v2/companies`: 另有 `name` (模糊比對), `orderBy` (enum, `只有` `uen` 與 `name`),
  `orderDirection` (enum), `responsiveEmployer` (boolean)

## 已知的欄位形狀陷阱 (Shape Traps)

- `status.id` 在 `/v2/search` 是`字串` `"102"`, 在 `/v2/jobs` 是`數字` `102`。
  本套件以 `FlexInt` 同時吃下兩種。
- `postedCompany` 在獵頭代發的職缺上是`仲介`, 實際雇主在 `hiringCompany`。
- `metadata.isHideSalary` 為真的職缺`仍會帶` salary, 那是平台推估值, 不是雇主公開數字。
- `description` 是 employer 貼上的 rich text (HTML), 不是純文字。

## 公開路徑上做不到的事 (Not Available)

- `列出某雇主的所有職缺`: `/v2/companies/{uen}/jobs` 需授權, 而搜尋的 `uen` 欄位不生效。
  只能用 `mcf search` 以雇主名稱當關鍵字近似取得。
- `技能字典查詢`: `/v2/skills` 恆回 `[]`。
