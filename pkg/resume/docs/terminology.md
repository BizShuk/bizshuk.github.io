# 履歷與職缺庫 — 術語表 (Terminology)

本檔是 `pkg/resume/` 底下履歷與職缺庫術語的`單一定義來源 (single source of truth)`.
職缺檔案, 排序總表, 缺口分析與教學文件一律沿用此處用詞, 同一概念不得有第二種說法.

網站本身 (首頁, 履歷頁, 衝浪頁, 政策頁, 部署) 的術語見
[docs/terminology.md](../../../docs/terminology.md).

推論服務的技術術語 (`TTFT`, `TPOT`, `KV cache`, `prefill / decode`, `goodput` 等)
不重複定義於此, 以 [tutorials/ai_inference_serving.md](tutorials/ai_inference_serving.md)
的術語表為準.

## 職缺庫 (Job Description Library)

住在 `jd/`, 是履歷的對照面: 履歷描述`我有什麼`, 職缺庫描述`市場要什麼`,
兩者之間的差值就是缺口. 角色軌道的完整導覽見 [tutorials/jd_role_map.md](tutorials/jd_role_map.md).

| 術語 (Term) | 英文 (English) | 定義 (Definition) | 出處 (Source) |
| ----------- | -------------- | ----------------- | ------------- |
| 職缺庫 | JD Library | 一份 JD 一個檔案的職缺蒐集. 檔名為 `<role_type_or_title>.<team>.md`, 依公司分目錄 | `jd/` |
| 排序總表 | Ranked Table | `jd/README.md` 中依匹配度排序的索引. 由各檔 front matter 產生, 不手工維護 | `jd/README.md` |
| 角色類型 | `role_type` | 職缺在職涯軌道上的位置. 判準是`產出物與決策單位`, 不是職稱字面 | 各 JD front matter |
| 角色面向 | `role_facing` | 工作對象是內部團隊 (`internal`) 或外部客戶 (`customer_facing`). 決定面試考的是系統設計還是客戶情境 | 各 JD front matter |
| 領域 | `domain` | 業務領域. 決定`領域缺口`要不要補, 以及補的成本 | 各 JD front matter |
| 匹配度 | `score` | 履歷對該 JD 硬性要求的覆蓋率, 加軟性要求與領域相鄰度調整, 範圍 0-100. `不是錄取機率` | 各 JD front matter, 評分基準見 `jd/README.md` |
| 硬性要求 | Must-have | JD 中不滿足即被過濾的條件. 對應原文 `Minimum Qualifications` 等標題, 原標題以 `— 原文` 保留 | 各 JD `任職要求` |
| 加分要求 | Preferred | 不滿足不會淘汰, 但決定面試深度的條件 | 各 JD `任職要求` |
| 其他條件 | Other | 國籍, 工作權, 出差比例, 系統標記技能等非能力條件 | 各 JD `任職要求` |
| 缺口分析 | Gap Analysis | 對單一職缺的三問拆解: `缺什麼`, `缺了會怎樣`, `怎麼補`. 內嵌於 JD 本文 | `google/geo_expansion_lead.google_cloud.md` |
| 門檻級缺口 | Blocking Gap | 落在硬性要求且無 `or equivalent practical experience` 的缺口. 在 recruiter screen 就被過濾, 面試表現補不回來 | 缺口分析區塊 |
| 角色地圖 | Role Map | 以`產出物`, `決策單位`, `對話對象`, `技術深度`, `時間尺度`, `成功衡量` 六個面向量出履歷與目標角色的距離, 並標出中間站的圖表 | `google/geo_expansion_lead.google_cloud.md` |
| 中間站 | Intermediate Role | 角色地圖中連接目前位置與目標角色的過渡職缺. 同時接受現有背景, 又能累積目標角色缺的那項責任 | 角色地圖區塊 |
| 對位敘事 | Positioning Angle | 針對該職缺的履歷改寫策略與投遞判斷, 每份 JD 的最後一個區塊 | 各 JD 檔尾 |
| 語彙轉換 | Vocabulary Translation | 同一段經歷改寫成目標軌道的語彙. 例如`設計全球資料中心冗餘`改寫為`使實作與維護成本下降兩個 job grade` | `對位敘事` 區塊 |
| 判重 | Duplicate Posting | 同一雇主以不同職級掛出內文相同的職缺. 以 `description` 字串相等判定, 判準是投薪幅高的那一檔 | `jd/README.md` |
| 覆蓋註記 | Coverage Notes | 記錄哪些公司已蒐集, 哪些被技術性阻擋, 以及各職缺站可用的抓取管道 | `jd/_coverage_notes.md` |

`role_type` 的九個列舉值:

| 值 (Value) | 中文 | 定義 (Definition) |
| ---------- | ---- | ----------------- |
| `ic` | 個人貢獻者 | 產出是程式與系統, 決策單位是服務或模組 |
| `staff_ic` | 資深個人貢獻者 | 不帶人但跨團隊, 產出是共用元件, 標準與技術決策 |
| `architect` | 架構師 | 產出是架構方案與設計審查, 決策單位是業務線 |
| `lead` | 技術主管 | 帶人且仍動手寫程式, 對交付與團隊同時負責 |
| `manager` | 工程經理 | 產出是團隊產能, 對人員, 排程與跨團隊對齊負責 |
| `director` | 總監 | 決策單位是多個團隊與組織設計 |
| `program_manager` | 專案經理 | 產出是跨團隊交付的確定性, 對里程碑, 相依與風險負責 |
| `pre_sales` | 售前 | 產出是客戶方案與參考架構, 成功衡量掛在商機上 |
| `strategy_lead` | 策略主管 | 產出是損益與資本配置決策, 屬商業軌道而非工程軌道 |

## 投遞狀態 (Application Status)

`status` 是本庫唯一的列舉型狀態值, 記錄於各 JD front matter,
並在本文的`基本資訊`表以 `投遞狀態 (Status)` 呈現:

| 值 (Value) | 意義 (Meaning) |
| ---------- | -------------- |
| `not_applied` | 已建檔, 尚未投遞 |
| `shortlisted` | 已決定投遞, 尚未送出 |
| `applied` | 已送出履歷 |
| `screening` | 進入履歷篩選或 recruiter screen |
| `interviewing` | 進入面試流程 |
| `offer` | 收到錄取通知 |
| `rejected` | 對方拒絕 |
| `withdrawn` | 本人撤回 |
| `closed` | 職缺下架或截止 |

每次變更 `status` 須同步更新 `status_updated` (格式 `YYYY-MM-DD`).
