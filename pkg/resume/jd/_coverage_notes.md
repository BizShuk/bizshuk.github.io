# 蒐集覆蓋範圍與缺口 (Coverage and Gaps)

`蒐集日期`: 2026-09-02 (2026-09-03 追加 Societe Generale 與 Airwallex Billing 兩檔, 2026-09-07 追加 Red Hat 一檔, 2026-09-08 追加 Grab 一檔與 `TSMC 台灣 14 檔`, 2026-09-10 追加 `Meta 6 檔`, 2026-09-11 追加 `CarHarbor 1 檔`, 2026-09-14 追加 `Visa 1 檔`, 2026-09-22 追加 Built In 來源之 `Airwallex 1 檔`、`JPMorganChase 2 檔`、`Mastercard 2 檔`、`Ambiq 1 檔` 與 Google 官方直取之 `Google 1 檔`)

## 已完成的公司 (Covered)

| 公司 | 職缺數 | 資料來源 | 全文取得 |
| --- | --- | --- | --- |
| Airwallex | 8 | Ashby job board API + Built In + MCF API | 是 |
| JPMorganChase | 2 | Built In / 官網直取驗證 | 是 |
| Mastercard | 2 | Built In / 官網直取驗證 | 是 |
| Ambiq | 1 | Built In / 官網直取驗證 | 是 |
| OKX | 5 | Greenhouse board API (`boards-api.greenhouse.io/v1/boards/okx`) | 是 |
| Google | 7 | careers 搜尋頁與職缺頁直取; Geo Expansion Lead 經 `api.mycareersfuture.gov.sg/v2/jobs/<uuid>` | 是 |
| Stripe | 4 | stripe.com 職缺索引 + Greenhouse board API | 是 |
| OpenAI | 4 | Ashby job board API | 是 |
| Anthropic | 1 | Greenhouse board API | 是 |
| Squarepoint Capital | 1 | Greenhouse board API | 是 |
| Tower Research Capital | 1 | Greenhouse board API | 是 |
| Jump Trading | 2 | Greenhouse board API | 是 |
| DRW | 1 | Greenhouse board API | 是 |
| Thoughtworks | 1 | 使用者提供連結, 官網職缺頁直取驗證 | 是 |
| Apple | 4 | jobs.apple.com 職缺詳情頁直取 (搜尋 API 需授權, 詳情頁可 SSR 抓取) | 是 |
| Societe Generale | 1 | 使用者提供 MyCareersFuture 連結, 以 `api.mycareersfuture.gov.sg/v2/jobs/<uuid>` 取全文 | 是 |
| Red Hat | 1 | 使用者提供 LinkedIn 職缺全文, 未附連結 | 是 |
| Grab | 1 | 使用者提供 LinkedIn 職缺全文, 未附連結 | 是 |
| CarHarbor | 1 | 使用者人工貼入 JD 全文, `未附連結` | 是 (全文, 但雇主本身未經查證) |
| TSMC (台灣) | 14 | 使用者提供搜尋頁連結, 經瀏覽器自動化取得 `JobDetail` 全文 | 是 |
| Visa | 1 | 使用者貼入 LinkedIn 全文, 以 Workday CXS API (`visa.wd5.myworkdayjobs.com/wday/cxs/visa/Visa`) 反查職缺編號並比對全文 | 是 |

## 未完成的公司 (Not Covered)

### PayPal - `已查證, 新加坡目前沒有工程職缺`

PayPal 的對外職缺唯一來源是 Workday (`wd1.myworkdaysite.com/en-US/recruiting/paypal/jobs`).
以空白查詢取得的全站職缺總數為 `26 筆`, 其中新加坡僅 `1 筆`, 且為 `Sr Manager, Employee Relations` (人資職).

結論: PayPal 新加坡目前 `沒有` 任何 tech lead, architect 或 engineering manager 職缺可投.
PayPal 的國際總部雖設於新加坡, 但工程招募目前集中在美國 (San Jose, Austin) 與日本 (Tokyo).
建議做法是訂閱該 Workday 頁面的新加坡地點通知, 而非持續人工檢查.

### Microsoft - `技術性阻擋, 需人工開啟瀏覽器`

Microsoft 的職缺站已從 `jobs.careers.microsoft.com` 遷移到 `apply.careers.microsoft.com`
(改用 Eightfold 平台). 該站為純前端 SPA, 且其後端 API 對未授權來源回傳
`Not authorized for PCSX`; 舊的 `gcsservices.careers.microsoft.com` API 亦已失效
(TLS 憑證不符 + 404). 因此無法以無頭方式取得新加坡職缺清單與職缺全文.

`已知可用的入口`: <https://apply.careers.microsoft.com/careers?lc=Singapore>

需要人工在瀏覽器開啟並篩選 `Singapore` + `Software Engineering` + `Experienced professionals`,
或啟用瀏覽器自動化後重跑. 本次工作階段中 Claude 的 Chrome 擴充未連線.

### Meta - `技術性阻擋, 已人工補入 9 檔`

`metacareers.com` 對非瀏覽器的請求一律回傳 HTTP 400, 個別職缺頁對外部抓取回傳 404.
無法取得新加坡職缺清單與職缺全文.

`已知可用的入口`: <https://www.metacareers.com/jobs?offices%5B0%5D=Singapore>

搜尋引擎索引顯示 Meta 新加坡確實有 `Software Engineer, Infrastructure`,
`Software Engineering Manager, Infrastructure`, `Production Engineering Manager` 等職缺,
但未經直接查證, 因此`不`在此建立 JD 檔案, 以免記錄未驗證的要求內容.

`2026-09-10 更新`: [Software Engineer, Principal](meta/software_engineer_principal.xf_apac_products.md)
(XF APAC Products, job id `1298625432375018`) 由人工貼入 JD 全文並附連結後建檔,
抓取管道的阻擋狀況不變. 其餘 Meta 新加坡職缺仍待人工開啟瀏覽器補齊.

`2026-09-10 第二次更新`: 再由人工貼入 `5 檔`全文建檔, 皆`未附原始連結`, 故 `url` 與 `job_post_id` 為 `null`:
[FDE, Enterprise AI (Senior)](meta/forward_deployed_engineer_senior.enterprise_ai.md),
[FDE, Enterprise AI (Staff)](meta/forward_deployed_engineer_staff.enterprise_ai.md),
[Software Engineer, Product](meta/software_engineer_product.singapore.md),
[Business Support Engineer - AdTech](meta/business_support_engineer.adtech.md),
[Business Support Engineer - Meta Business Agents](meta/business_support_engineer.business_agents.md).
Meta 累計 `6 檔`. 抓取管道的阻擋狀況`仍然不變`, 全部 6 檔都是人工貼入.
另一個可用入口是 <https://www.metacareers.com/jobsearch/> (需人工在瀏覽器開啟).

`2026-09-10 第三次更新`: [Director, Compliance & Assurance](meta/director_compliance_assurance.regulatory_compliance_programs.md)
(Regulatory Compliance Programs, job id `1728142584963155`) 由人工貼入全文並附連結後建檔, Meta 累計 `7 檔`.
本次另試過三條繞道, `全部失敗`, 記錄以免重複嘗試: 帶完整瀏覽器標頭的 `curl` 仍回 HTTP 400;
鏡像站 `metacareers.dejobs.org` 的搜尋結果為 JS 渲染, 靜態抓取只得到空殼;
`WebSearch` 無法以 job id 反查職缺. 唯一可行的自動化路徑是`已渲染的瀏覽器`
(Claude 的 Chrome 擴充或 headless Chrome `--dump-dom`), 本次工作階段中該擴充未連線.

`2026-09-10 第四次更新`: 再由人工貼入 [Software Engineer, Infrastructure (Singapore)](meta/software_engineer_infrastructure.singapore.md)
全文建檔, `未附原始連結`, 故 `url` 與 `job_post_id` 為 `null`. Meta 累計 `8 檔`.
該檔與既有的 Software Engineer, Product 為`同一份 generic req 的兩個方向`,
硬性要求八條中七條逐字相同, 差異僅在職責段; 未來比對 Meta 通用職缺時可據此直接判親緣.

`2026-09-10 第五次更新`: 再由人工貼入 [AI Data Center Development Manager](meta/ai_data_center_development_manager.data_center_strategy.md)
全文建檔, `未附原始連結`. Meta 累計 `9 檔`.
該檔的 JD 條列美國稅制與環境法規名詞 (`federal, state, and local tax structures`, `wetland delineation`),
但職缺掛在新加坡, 判定為`美國模板的地區複用`; 遇到 Meta 這類全球職缺時,
`職責文本的法域`與`職缺掛牌地點`可能不一致, 應在檔案中標註而非照單全收.

### Apple - `已補齊 (2026-09-02 追加)`

Apple 的職缺搜尋 API (`jobs.apple.com/api/v1/search`) 對外回 `401 User Unauthorized`,
搜尋頁本身也只吐出殼層; 但`職缺詳情頁是 server-rendered`, 可直接抓取全文.

因此做法是先以搜尋引擎索引取得 Role Number, 再直取
`https://jobs.apple.com/en-sg/details/<roleNumber>/<slug>` 解析頁內 JSON.

已建檔 4 檔. 另有 1 檔 `Software Engineering Manager, Apple Services Engineering`
(Role Number `200623624-3278`) 經查證`已下架`, 頁面顯示職缺不再開放, 故未建檔.

`注意`: 此法依賴搜尋引擎已索引該職缺, 無法保證涵蓋 Apple 新加坡的`全部`開放職缺.
若要完整清單, 仍需人工開啟 <https://jobs.apple.com/en-sg/search?location=singapore-SGP>.

### TSMC - `需瀏覽器自動化, 無頭抓取全站被擋`

`careers.tsmc.com` `整站`置於 Cloudflare bot check 之後: curl 與無頭 Chrome 對首頁, robots.txt,
sitemap.xml 與搜尋頁一律得到 `HTTP 403` 的 challenge 頁, WebFetch 同樣失敗.

`可用做法`: 以已登入的真實 Chrome 通過驗證後, 在該分頁的 same-origin context 內
`fetch('/en_US/careers/JobDetail?jobId=<id>')` 取回 server-rendered HTML 再解析.
職缺清單頁的每個標題連結都帶 `jobId`, 因此先從搜尋頁取 id 清單, 再逐一取全文即可.

`本次篩選條件`: TSMC Group + Location `Taiwan` + Job Category `Information Technology`
+ Job Type `Manager / Executive`, 共 `14 筆` (兩頁, 每頁 10 筆).

`注意`: 該職類過濾`會混入非 IT 職缺` — 例如 `Loss Prevention Manager` 實為實體資產保全,
與軟體工程零交集. 以 `Information Technology` 過濾出的清單需人工複核.

### MyCareersFuture - `已驗證可用的抓取管道`

`mycareersfuture.gov.sg` 的職缺頁是純前端 SPA, 直接抓 HTML 只會拿到殼層.
但其後端 API `https://api.mycareersfuture.gov.sg/v2/jobs/<uuid>` `無須授權即可取得完整 JSON`,
uuid 即職缺網址結尾的 32 碼十六進位字串.

該 JSON 比網頁多出可用欄位: `salary` (公開薪幅), `minimumYearsExperience`, `positionLevels`,
`metadata.expiryDate` (截止日), `metadata.jobPostId`, `skills` (系統標記的技能標籤), `address`.
`這是目前所有來源中資訊密度最高的一個`, 且涵蓋所有在新加坡合法招募的雇主 (含銀行與非科技業).

`已知落差`: Airwallex 的 Ashby job board `未刊登` Billing 團隊的兩檔職缺, 需靠 MCF 才看得到;
反之 MCF 也未涵蓋 Airwallex 全部職缺. `兩個來源必須並用`, 單靠任一邊都會漏.
另外 MCF 上同一雇主可能以 `不同職級掛出內文完全相同的職缺`, 比對時應以 `description` 字串相等直接判重.

`建議`: 後續蒐集應優先走 MCF, 再以公司官方 job board 補齊 MCF 未刊登的職缺.

### Visa - `Workday CXS API 可直取, 無須授權`

使用者貼入的是 LinkedIn 版本 (`Responses managed off LinkedIn`, 無 apply 連結),
但 Visa 的實際投遞管道是 Workday, 且其 `CXS API` 無須授權即可用:

- 搜尋: `POST https://visa.wd5.myworkdayjobs.com/wday/cxs/visa/Visa/jobs`,
  body 為 `{"appliedFacets":{},"limit":20,"offset":0,"searchText":"<關鍵字>"}`,
  回傳含 `externalPath` 與 `bulletFields` (即 `REF` 職缺編號).
- 全文: `GET https://visa.wd5.myworkdayjobs.com/wday/cxs/visa/Visa<externalPath>`,
  回傳 `jobPostingInfo`, 欄位比網頁多出 `startDate` (實際掛出日), `jobReqId`, `remoteType`,
  `jobRequisitionLocation` (含辦公室, 本檔為 `Singapore, SingPost`).

`可複用性`: 這是 Workday 的通用路徑格式 (`/wday/cxs/<tenant>/<site>/jobs`),
先前判定`技術性阻擋`的 PayPal 同樣走 Workday, 未來應以此法重試而非人工檢查.

`本次查得的新加坡 Visa 工程職缺`: 除本檔外另有 `Staff Software Engineer (AI Engineer)`
(`REF084670W`, API 顯示 13 天前掛出), `尚未建檔`.

## 觀察 (Observations)

- 新加坡的 `tech lead / architect / engineering manager` 供給, 目前集中在
  `金融科技 (Airwallex)` 與 `加密貨幣 (OKX)` 兩個產業, 兩者合計提供了本輪十檔中的十檔高分職缺.
- `AI 原生公司 (OpenAI, Anthropic)` 在新加坡的職缺以 `業務與客戶交付` 為主,
  幾乎沒有內部工程職缺, 代表其新加坡站尚未建立工程組織.
- `量化交易公司` 的新加坡職缺高度集中在 quant 與 C++ 低延遲, 對非 C++ 背景者實質封閉,
  唯一入口是基礎設施支援與平台維運.
- `TSMC 台灣 14 檔是本庫第一組非新加坡職缺`, 收錄時的地點前提與其餘檔案不同, 見
  [README.md](README.md) 的 `台灣職缺的地點前提` 一節.
- TSMC 的 `Manager / Executive` 職類中, 真正屬於軟體工程領導軌道的只有 `IT DevOps Manager`
  與 `Solutions Architect` 兩檔; 其餘多為`領域專業職` (資安本業, HR/財會產品, 半導體設備),
  缺口屬領域資格而非工具層, 無法以工程經驗替代.
