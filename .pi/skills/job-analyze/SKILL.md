---
name: job-analyze
description: Use when a job URL or job description is pasted, linked, or fetched and should be recorded, scored against the resume, and analyzed for role value, opportunities, and roadmap - triggers on "note this JD", "evaluate this job", "analyze this job", "job-analyze", "這個職缺適合嗎", "這個職缺值得嗎", "匹配度", "幫我評估這份 JD", "分析這個角色", "職涯路線圖", "roadmap for this role", on re-scoring existing JD files after Resume.md changes, and on any write into pkg/resume/jd/.
---

# Job Analyze

## Overview

把一個職缺連結或 JD 全文收進 [pkg/resume/jd/](../../pkg/resume/jd/) 職缺庫, 並回答兩個問題:

| 問題 | 產出 | 區塊 |
| --- | --- | --- |
| `我夠不夠格` | 格式化 JD, 匹配度分數, 缺口分析 | `評估摘要`, `缺口分析` |
| `這個位子值不值得坐, 坐上去之後去哪` | 角色評估, 機會, 路線圖 | `角色評估`, `機會`, `路線圖` |

`核心原則`:

- 分數是`履歷對 JD 硬性要求的覆蓋率`, 不是`錄取機率`. 兩者混淆會讓整個庫失去可比性, 因為覆蓋率可查證, 錄取機率不可.
- 匹配度與角色價值`分開判斷`. 一個 `85` 分的職缺可能是職涯死路, 一個 `55` 分的職缺可能是最好的跳板.

## Modes

| 模式 | 觸發 | 執行步驟 |
| --- | --- | --- |
| `完整分析` (預設) | 貼上 URL 或 JD, 問值不值得, 要機會或路線圖 | 全部 |
| `只建檔評分` | 明說只記錄, 或 `Resume.md` 更新後重算既有檔案分數 | 跳過步驟 5-7 |
| `只分析不建檔` | 明說不要寫入職缺庫 | 跳過步驟 2-4, 8, 9; 三個分析區塊直接輸出在對話 |

`不適用`: 撰寫求職信 (屬 `coverletter`), 面試題庫, 履歷改寫本身 (那些屬 `docs/interview/` 與 `Resume.md` 的職責).

## Non-Negotiables

- `輸出根目錄固定`: `~/projects/product/bizshuk.github.io/pkg/resume/jd/<company>/`. 動手前先檢查 `~/projects/product/bizshuk.github.io` 存在; 不存在就回報 `ERROR: repo not found: <path>` 並`立即停止`, 不建檔, 不改寫到其他位置, 不自行建立目錄.
- `先讀 [jd/README.md](../../pkg/resume/jd/README.md)` 再動筆. 它是 schema, 評分基準與特例規則的單一事實來源, 且會演進; 不要憑本檔記憶的欄位值硬寫.
- `建立時間`: 每份 JD 的 front matter 必有 `created: YYYY-MM-DD`, 建檔當天寫入, 之後重抓或重評分`不得改動`. 舊檔缺 `created` 時以 `fetched` 值補上.
- `不得杜撰要求`. 只寫 JD 原文有的內容. 抓取被阻擋而使用者只給了片段時, 標明來源為人工貼入, 不要用搜尋引擎摘要補齊要求條目.
- `必寫不符合的部分`. 只列符合項的評估等於沒有評估; 缺口是這個庫的核心產出.
- `分數要與鄰居對齊`. 給分後在排序總表找上下各一檔, 確認相對位置說得通, 再定案.
- `證據分級`: 每個角色判斷標注來源 - `原文` (JD 明寫), `推論` (由 JD 用詞, 職級, 團隊位置推得), `外部` (公司公告, 財報, 新聞, 需附連結). 不得把推論寫成事實.
- `必寫風險`: 只講機會的分析是廣告. `角色評估` 一定要有 `風險與紅旗` 一項, 即使結論是`未發現`.
- `路線圖要可執行`: 每一步有動作與時間窗, 不寫`提升領導力`這類無法驗收的句子.
- `重複職缺合併 (Merge on Duplicate)`: 若職缺已在庫內 (相同 URL, job_post_id, 或同公司同職務), 不得建立重複檔案; 一律就地合併更新 (merge in-place), 保留原始 `created` 日期, 更新 `status`, `status_updated`, `fetched`, 補齊分析區塊並同步 `README.md` 索引.
- `格式`: 不使用粗體, 一律 backtick 強調; 半形標點; Mermaid 邊線文字加雙引號.

## Workflow

1. `前置檢查`: `test -d ~/projects/product/bizshuk.github.io`, 失敗即報錯停止. 職缺已在庫內則進入合併模式 (merge in-place), 不另立新檔.
2. `取得全文`: 先試 `pkg/resume/` 的 `resume mcf` CLI 或直接抓取. 被阻擋且改由人工貼入時, 於 [_coverage_notes.md](../../pkg/resume/jd/_coverage_notes.md) 記下該公司的阻擋狀況與已補檔數.
3. `建檔 / 合併`: 路徑 `~/projects/product/bizshuk.github.io/pkg/resume/jd/<company>/<title_slug>.<team_or_domain>.md`, 全小寫加底線. 既有檔案保留原始 `created` 日期並就地覆蓋合併; 新檔寫入今日 `created`. front matter 欄位與列舉值以 README 的 `資料欄位 (Schema)` 表為準, 未載明一律 `null`, 不要以 `0` 或空字串冒充. 本文區塊順序固定, 見 [template.md](template.md).
4. `評分`: 逐條比對硬性要求, 再以加分要求與領域相鄰度調整, 對照 README 的 `分數的意義` 區間表. 觸發條件成立時寫 `缺口分析`.
5. `角色評估`: 依 `Role Assessment` 六個面向逐項判斷.
6. `機會`: 依 `Opportunities` 四類找出這個角色能打開的東西.
7. `路線圖`: 依 `Roadmap` 四個階段寫行動. 對位敘事若因分析結果改變主軸, 一併修訂.
8. `更新索引`: 排序總表插入到正確分數位置 (降冪), 更新檔數與`最近一次蒐集`日期; 有寫缺口分析則加入缺口分析清單並在總表列標記 `+缺口分析`.
9. `封存過期職缺`: 見 `Archive Sweep`.
10. `回覆使用者`: 給三行結論 - `值不值得投` (含分數), `最大機會`, `最優先的一個動作`; 接著`一定要列出檔案路徑`: 本次建立或更新的 JD 檔絕對路徑, 以及本次封存的檔案 (`原路徑 -> 新路徑`, 沒有則寫 `封存: 無`). `只分析不建檔` 模式路徑一行寫 `未建檔`. 細節留在檔案.

## Scoring

分數來自硬性要求的逐條覆蓋, 常見的判定陷阱:

| 情境 | 判法 |
| --- | --- |
| 年資達標但領域不同 | 算部分覆蓋. 年資是門檻, 不是覆蓋率 |
| 能力有但只在個人專案 | 算`舉證缺口`, 扣分低於`能力缺口`, 但必須在`不符合的部分`明講規模落差 |
| 職級跨越兩級以上 (如 Tech Lead 對 Principal) | 這是`結構性缺口`, 壓進 `50-64` 區間, 不因其他條目全中而拉高 |
| 台灣職缺 | `不`因地點折抵分數, 但搬遷與再任前提寫進`我的對位敘事` |
| 同雇主多職級掛同一份 JD | 投薪幅高的那一檔, 低薪版不提高錄取率只壓低錨點 |
| 履歷有相鄰工具經驗 (EFK vs DataDog) | 算命中, 於面試準備註明等價概念即可 |

## Gap Analysis

缺口分析是`選備`, 但出現以下任一情況時應內嵌:

- 分數落在 `50-64` (有結構性缺口, 需要說明缺什麼與值不值得轉向)
- 該檔代表一條`新軌道` (客戶面向, 商業策略, 低延遲交易等), 值得留下判斷依據

固定回答四個問題: `缺什麼` (表格: 缺口 / 類型 / 補法 / 可補期), `缺了會怎樣`, `怎麼補` (指出最高槓桿的單一動作), `值不值得轉向`.

缺口類型只有四種: `舉證`, `技能`, `工具`, `規模 / 職級`. 分類本身就是資訊 - 舉證缺口是短期可補的, 職級缺口不是.

## Role Assessment

| 面向 | 要回答的問題 | 常見訊號 |
| --- | --- | --- |
| `真實職級` | 標題與職責描述的 scope 是否一致? | 標題 Lead 但職責全是實作 = 資深 IC; 寫 `set direction for org` = 高於標題 |
| `影響範圍 (Scope)` | 影響一個服務, 一個團隊, 一條產品線, 還是全公司? | 使用者數, 跨團隊字眼, 匯報對象層級 |
| `組織位置` | 成本中心還是營收中心? 新建團隊還是維運既有系統? | `greenfield`, `build from zero`, `modernize`, `migrate` |
| `業務壓力` | 這個職缺為什麼現在開? | 新市場, 新法規, 產品上線, 補離職, 組織重整 |
| `成長天花板` | 這個位子往上一級是什麼, 內部有沒有路? | 團隊規模, 是否有 Staff/Principal 或 Director 層 |
| `風險與紅旗` | 什麼會讓這份工作變差? | 職責過多 (一人抵三職), JD 貼錯段落, 反覆重開, 團隊剛經歷裁員 |

結論以一句 `角色本質` 收尾, 例如 `掛 Manager 標題的平台救火隊, 真實 scope 為單一系統`.

## Opportunities

四類各找最多兩點, 找不到就寫`無`, 不湊數:

- `能力資產`: 做完這份工作後, 履歷上會多出什麼`可舉證`的東西 (規模, 領域, 職級).
- `軌道轉換`: 是否打開新軌道 (客戶面向, 管理, AI 平台, 特定產業), 對照 [jd/README.md](../../pkg/resume/jd/README.md) 的軌道與庫內相鄰職缺.
- `市場訊號`: 這類職缺在庫內與市場上是變多還是變少, 可附上庫內同類檔數.
- `槓桿與網絡`: 雇主品牌, 內部轉調機會, 接觸的人脈層級.

## Roadmap

四個階段, 每階段 2-4 項 `動作 + 時間窗 + 驗收`:

| 階段 | 時間窗 | 內容重心 |
| --- | --- | --- |
| `投遞前` | 0-2 週 | 補舉證缺口, 履歷調整, 內推, `投遞前必問`的答案 |
| `面試` | 2-6 週 | 預期考題類型, 要準備的故事, 要反問的問題 |
| `到職 90 天` | 入職後 | 30/60/90 天各一個可交付的成果, 對準`業務壓力`那一項 |
| `1-3 年` | 長期 | 這個角色通往的下一個位子, 以及到那裡需要累積的證據 |

`1-3 年` 階段附一張 Mermaid 路徑圖, 標出`現在 → 本職缺 → 下一步 (最多兩條分支)`.

## Archive Sweep

每次建檔或重評分後都對整個職缺庫跑一次, 規則由腳本決定, 不要手算日期:

```bash
python3 ~/projects/product/bizshuk.github.io/skills/job-analyze/scripts/archive_stale.py --dry-run
python3 ~/projects/product/bizshuk.github.io/skills/job-analyze/scripts/archive_stale.py
```

- `判準`: `created` (缺則 `fetched`) 早於今天往前推 `3 個月`, 就從 `jd/<company>/` 移到 `jd/archive/<company>/`. 已在 git 追蹤的檔案用 `git mv`.
- `SKIP no-date`: 沒有 front matter 日期的檔案不動, 在回覆中列出請使用者處理.
- `投遞中不封存`: `status` 為 `applied`, `screening`, `interviewing`, `offer` 的檔案由腳本跳過 (`SKIP active`), 因為進行中的流程仍需要它留在主表.
- `移動後同步`:
  - [jd/README.md](../../pkg/resume/jd/README.md): 從 `排序總表` 與 `投遞追蹤` 移除該列, 加到 `已封存` 清單, 更新總檔數. 已封存的檔案之後`不得`再加回排序總表.
  - `相對連結`: `grep -rn '<檔名>' pkg/resume/jd/ pkg/resume/docs/` 找出其他檔案指向它的連結, 改成 `archive/<company>/` 下的新路徑.

## Common Mistakes

- `把分數當錄取機率`: 導致 Principal 這類高門檻職缺被壓到 30 分以下, 失去`還差什麼`的參照價值.
- `把匹配度當角色價值`: 高分只代表門檻低於能力, 也可能代表這是一個往下走的位子.
- `評估摘要寫成 JD 摘要`: 摘要區塊要寫`對照履歷的結論`, 不是複述職責.
- `漏掉投遞前必問`: 這一段是把評估轉成行動的介面, 沒有它整份檔案只能讀不能用.
- `機會全是雇主文案`: `impactful`, `fast-growing` 不是機會, 必須換成履歷上可舉證的東西.
- `路線圖只有投遞前`: 沒有 90 天與 1-3 年, 就無法回答`坐上去之後去哪`.
- `只更新 JD 檔忘了索引`: 排序總表, 檔數, 蒐集日期, 缺口分析清單四處要同步.
- `我的對位敘事寫成自誇`: 它的作用是`敘事順序`, 要明確指出先講哪句後講哪句, 以及哪個弱點必須主動處理而非迴避.
- `對話輸出整份分析`: 使用者要的是結論與檔案路徑, 不是在終端機滾動整頁.
