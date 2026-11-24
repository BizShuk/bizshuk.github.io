# 取得 JD 全文的操作細節 (Access Reference)

[SKILL.md](../SKILL.md) Workflow 步驟 2 的三層抓取順序, 每一層的實際做法.
公司別的阻擋紀錄與已驗證管道不在這裡, 一律寫在 [_coverage_notes.md](../../../../pkg/resume/jd/_coverage_notes.md).

## 第 1 層: 公開抓取 (WebFetch / curl)

| 來源 | 端點 |
| --- | --- |
| Greenhouse | `https://boards-api.greenhouse.io/v1/boards/<company>/jobs?content=true` |
| Ashby | `https://api.ashbyhq.com/posting-api/job-board/<company>` |
| Workday CXS | `https://<tenant>.wd<N>.myworkdayjobs.com/wday/cxs/<tenant>/<site>/jobs` (POST 搜尋, 詳情另取 `job/<path>`) |
| MyCareersFuture | `https://api.mycareersfuture.gov.sg/v2/jobs/<uuid>` |

`失敗訊號`: HTTP 400/403, Cloudflare challenge 頁, 回傳的 HTML 沒有職缺內文 (前端渲染). 出現任一個就換第 2 層.

## 第 2 層: 桌面操作 (agent-desktop)

Chrome 的網頁內容不在 agent-desktop 讀得到的 AX tree 裡, `find` 與 `snapshot` 只看得到瀏覽器外框, 所以內文一律`以截圖讀取`.

```bash
agent-desktop launch "Google Chrome" --activate
agent-desktop press cmd+t
agent-desktop press cmd+l
agent-desktop find --app "Google Chrome" --role textfield      # 取得網址列 ref
agent-desktop type <ref> "<url>" --headed                       # 必須 --headed, 語意寫入不會導航
agent-desktop press return
agent-desktop wait 6000
agent-desktop list-windows --app "Google Chrome"               # 以視窗標題確認已載入
agent-desktop screenshot --app "Google Chrome" <path>.png      # 讀取畫面內文
agent-desktop scroll ... / press space                          # 翻頁後再截圖, 直到 JD 結尾
```

- 截圖存 scratchpad, 不進 repo. 逐張讀出內文後拼成全文, 截不完整的段落標明缺漏, 不自行補寫.
- 列表頁要先點開單筆職缺, 讓詳情面板出現再截圖.
- 登入, 驗證碼, cookie 同意等需要使用者介入時, 停下來請使用者處理, 不代填帳密.

## 第 3 層: 人工貼入

請使用者貼 JD 全文與原始連結. 建檔時來源欄位標明 `人工貼入`, 不用搜尋引擎摘要補齊要求條目.

## 已知網站

搜尋來源清單見 [job.site](job.site).

| 網站 | 第 1 層 | 備註 |
| --- | --- | --- |
| JobStreet (`sg.jobstreet.com`) | 未驗證 | 2026-11-24 經第 2 層可開啟搜尋頁, 未被 bot check 擋 |
| LinkedIn | 不可用 | 需登入, 直接走第 2 或第 3 層 |

## 公司職缺入口 (Top 40)

前 20 家依職缺庫檔數排序, 取有可用入口者; 第 21-40 家為庫外公司, 見表後說明. 入口取自庫內 JD 的 `url` 欄位與 `_coverage_notes.md`, 為`已實際取得過全文`的管道; 第 1 層欄位表示可否不開瀏覽器直取.

| 公司 | 庫內檔數 | 入口 | 第 1 層 |
| --- | --- | --- | --- |
| TSMC | 14 | `https://careers.tsmc.com/en_US/careers/JobDetail?jobId=<id>` | 不可用, Cloudflare 擋, 走第 2 層 |
| OKX | 9 | `https://job-boards.greenhouse.io/okx`, API `boards-api.greenhouse.io/v1/boards/okx` | 可用 |
| Meta | 9 | `https://www.metacareers.com/jobsearch/` | 不可用, 非瀏覽器請求回 400 |
| Airwallex | 8 | `https://jobs.ashbyhq.com/airwallex` + MyCareersFuture | 可用, 兩來源需並用 |
| Google | 7 | `https://www.google.com/about/careers/applications/jobs/results/` | 可用 |
| Apple | 7 | `https://jobs.apple.com/en-sg/search?location=singapore-SGP` | 詳情頁可用, 搜尋 API 需授權 |
| Stripe | 4 | `https://stripe.com/jobs` | 可用 |
| OpenAI | 4 | `https://jobs.ashbyhq.com/openai` | 可用 |
| Societe Generale | 2 | MyCareersFuture | 可用 (MCF API) |
| Mastercard | 2 | Built In (`builtin.com/job/...`) | 可用 |
| Jump Trading | 2 | `https://www.jumptrading.com/hr/job?gh_jid=<id>` (Greenhouse) | 可用 |
| JPMorganChase | 2 | Built In (`builtin.com/job/...`) | 可用 |
| Visa | 1 | `https://visa.wd5.myworkdayjobs.com/Visa`, CXS `.../wday/cxs/visa/Visa/jobs` | 可用 |
| Tower Research | 1 | `https://www.tower-research.com/open-positions/` (Greenhouse) | 可用 |
| Thoughtworks | 1 | `https://www.thoughtworks.com/en-sg/careers/jobs` | 可用 |
| Squarepoint | 1 | `https://www.squarepoint-capital.com/open-opportunities` (Greenhouse) | 可用 |
| NUS | 1 | `https://careers.nus.edu.sg/` | 可用 |
| Amazon | 1 | `https://www.amazon.jobs/en/` | 可用 |
| DRW | 1 | `https://job-boards.greenhouse.io/drweng` | 可用 |
| Anthropic | 1 | `https://job-boards.greenhouse.io/anthropic` | 可用 |

`未列入`: Xingyun, Red Hat, Grab (庫內無 `url`), Dell (僅 LinkedIn), Ambiq (僅 Built In 單檔). 補到入口後再加入.

### 庫外公司 (21-40)

`選取方式`: 2026-11-24 以第 1 層 API 實測, 依新加坡職缺數排序取前 17 家, 另補 3 家新加坡大型雇主 (API 未驗證, 只確認首頁回 200). 新加坡職缺數為當日快照.

| 公司 | 新加坡職缺 (2026-11-24) | 入口 | 第 1 層 |
| --- | --- | --- | --- |
| Binance | 107 | `https://boards-api.greenhouse.io/v1/boards/binance/jobs?content=true` | 可用 (Greenhouse) |
| Databricks | 33 | `https://boards-api.greenhouse.io/v1/boards/databricks/jobs?content=true` | 可用 (Greenhouse) |
| Datadog | 18 | `https://boards-api.greenhouse.io/v1/boards/datadog/jobs?content=true` | 可用 (Greenhouse) |
| Jane Street | 15 | `https://boards-api.greenhouse.io/v1/boards/janestreet/jobs?content=true` | 可用 (Greenhouse) |
| GitLab | 11 | `https://boards-api.greenhouse.io/v1/boards/gitlab/jobs?content=true` | 可用 (Greenhouse) |
| Xendit | 10 | `https://boards-api.greenhouse.io/v1/boards/xendit/jobs?content=true` | 可用 (Greenhouse) |
| Coinbase | 9 | `https://boards-api.greenhouse.io/v1/boards/coinbase/jobs?content=true` | 可用 (Greenhouse) |
| Twilio | 9 | `https://boards-api.greenhouse.io/v1/boards/twilio/jobs?content=true` | 可用 (Greenhouse) |
| Zscaler | 8 | `https://boards-api.greenhouse.io/v1/boards/zscaler/jobs?content=true` | 可用 (Greenhouse) |
| Airbnb | 7 | `https://boards-api.greenhouse.io/v1/boards/airbnb/jobs?content=true` | 可用 (Greenhouse) |
| MongoDB | 7 | `https://boards-api.greenhouse.io/v1/boards/mongodb/jobs?content=true` | 可用 (Greenhouse) |
| Okta | 7 | `https://boards-api.greenhouse.io/v1/boards/okta/jobs?content=true` | 可用 (Greenhouse) |
| Pinterest | 7 | `https://boards-api.greenhouse.io/v1/boards/pinterest/jobs?content=true` | 可用 (Greenhouse) |
| Notion | 6 | `https://api.ashbyhq.com/posting-api/job-board/notion` | 可用 (Ashby) |
| Elastic | 5 | `https://boards-api.greenhouse.io/v1/boards/elastic/jobs?content=true` | 可用 (Greenhouse) |
| Linear | 5 | `https://api.ashbyhq.com/posting-api/job-board/linear` | 可用 (Ashby) |
| Rubrik | 4 | `https://boards-api.greenhouse.io/v1/boards/rubrik/jobs?content=true` | 可用 (Greenhouse) |
| Microsoft | 未知 | `https://apply.careers.microsoft.com/careers?lc=Singapore` | 未驗證, 首頁可開 |
| ByteDance / TikTok | 未知 | `https://jobs.bytedance.com/en/` | 未驗證, 首頁可開 |
| Shopee | 未知 | `https://careers.shopee.sg/` | 未驗證, 首頁可開 |
