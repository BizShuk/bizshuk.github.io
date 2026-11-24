---
company: "Societe Generale"
title: "Front Office Software Engineer - FIC Automated Trading"
department: "Front Office Development / FIC (Fixed Income & Currencies) Automated Trading"
role_type: ic
role_facing: internal
domain: low_latency_trading
job_post_id: "`MCF-2026-1537690`"
url: "https://www.mycareersfuture.gov.sg/job/information-technology/front-office-software-engineer-%E2%80%93-fic-automated-trading-societe-generale-5b56fc795c3abb18474b522325625ae3"
source: mycareersfuture
posted: 2026-09-02
expiry: 2026-09-16
fetched: 2026-09-03
created: 2026-09-03
min_years: 6
education: bachelor_stem
score: 35
status: not_applied
status_updated: 2026-09-03
---
# Societe Generale - Front Office Software Engineer, FIC Automated Trading

> ## 評估摘要 (Evaluation Summary)
>
> `評估日期`: 2026-09-03, 對照 [Resume.md](../../Resume.md)
>
> `匹配度 (Profile Matching Score)`: `35 / 100`
>
> `結論`: `薪幅吸引, 但硬性門檻直接卡死`. 這是坐在交易大廳上的 front office 開發職, 唯一的技術硬性要求是 `C 語言的紮實經驗`, 履歷完全沒有 C 或 C++ 的生產級經驗. 職級為 `Executive (IC)`, 不含帶人或架構治理範圍, 與履歷往 lead / architect 發展的方向相反. 缺口分析與補法見 [gap_analysis_and_study_map.md](gap_analysis_and_study_map.md).
>
> `符合的部分`:
>
> - `品質導向, 重視軟體設計與架構`: TSMC 的函式 30 行上限與 90% 覆蓋率, Change Healthcare 的共用框架設計, 是履歷上最強的一塊.
> - `技術科系學士學位`: 中央資工碩士, 東華資工學士, 超出要求.
> - `6 年以上經驗 (MCF 欄位要求)`: 12 年以上.
> - `物件導向設計與常見設計模式`: Java (Spring Boot), SOLID, design pattern.
> - `能在整合型團隊中工作`: ByteDance 跨部門推動, Change Healthcare 帶 5 人.
> - `Level 1 & 2 支援與production 事故追蹤`: ByteDance 的資料中心切換與 Change Healthcare 的 CloudWatch 告警體系.
>
> `不符合的部分`:
>
> - `C 語言的紮實經驗` — `唯一的硬性技術要求, 直接不符`. 履歷的 C++ 僅止於 2016 年的 LeetCode 與 Google C++ education.
> - `多執行緒與並行問題`: 只有 Go 的 goroutine 與事件驅動層級, 沒有 lock-free, memory ordering, cache line 這類低延遲語境的並行經驗.
> - `網路協定與 socket 程式設計`: 停留在 Nginx, Istio, 跨資料中心通道的應用層, 沒有直接寫過 socket.
> - `股票與衍生性商品產品知識`, `交易演算法與策略開發`: 完全空白.
> - `FIC / rates algo trading`, `低延遲交易系統`, `亞洲市場連線 (market connectivity)`: 完全空白.
> - `Windows 平台開發`: 完全空白, 履歷全在 Linux.
> - 職級是 IC 而非 lead, 且職責含 L1/L2 支援, title 對後續職涯的訊號偏向支援與維護.
>
> `投遞前必問`:
>
> - `C` 是否為筆誤, 實際指的是 `C++` 或 `C#`. 這一條決定整個職缺能不能碰.
> - JD 內文寫 `坐在香港的 front office 使用者旁邊`, 但職缺地點是新加坡 Marina Bay. 需確認實際辦公地與是否需要外派.
> - 這個職級是否有往 tech lead 或 desk architect 發展的路徑, 或本質上就是 desk 的長期支援位.
> - L1/L2 支援佔工時比例, 以及是否有跟隨亞洲市場開盤的早班或 on-call.
> - SGD 12,500 - 22,000 的區間中, 沒有交易領域背景會被定位在哪一段.

## 基本資訊 (Basics)

| 項目 | 內容 |
| --- | --- |
| 公司 (Company) | Societe Generale (法國興業銀行, 歐洲主要金融服務集團, 66 國 117,000 名員工) |
| 職稱 (Title) | Front Office Software Engineer - FIC Automated Trading |
| 部門 (Department) | Front Office Development / FIC (Fixed Income & Currencies) Automated Trading |
| 角色類型 (Role Type) | `ic` · `internal` · `low_latency_trading` |
| 地點 (Location) | 新加坡 Marina Bay Financial Centre, 8 Marina Boulevard (018981) |
| 職缺編號 (Job Post ID) | `MCF-2026-1537690` |
| 原始連結 (Original Link) | <https://www.mycareersfuture.gov.sg/job/information-technology/front-office-software-engineer-%E2%80%93-fic-automated-trading-societe-generale-5b56fc795c3abb18474b522325625ae3> |
| 刊登與截止 (Posted / Expiry) | 2026-09-02 刊登, `2026-09-16 截止` |
| 僱用形式 (Employment Type) | Full Time, 1 個名額 |
| 薪酬 (Package) | `SGD 12,500 - 22,000 / 月` (公開揭露) |
| 最低年資 (Min Experience) | `6 年以上` |
| 學歷要求 (Min Qualification) | 學士 (SSEC EQA 70), 主修資訊, 數學或工程 |
| 投遞狀態 (Status) | `not_applied` (更新於 2026-09-03) |
| 抓取日期 (Fetched) | 2026-09-03, 來源 `api.mycareersfuture.gov.sg/v2/jobs/<uuid>` |

## 團隊背景 (About the Team)

`Front Office development` 團隊負責為證券與融資交易 desk 設計, 實作, 維護與支援電子交易軟體方案.
JD 強調的環境特徵: 多文化, 鼓勵探索新技術, `扁平管理結構`, 高曝光度與`對交易活動的直接影響`.
工作位置在`交易大廳上 (on the trading floor)`, 與 front office 使用者近距離協作.

## 崗位職責 (Responsibilities)

`開發 (Development)`:

- 蒐集使用者需求, 設計符合 front office 業務需要的軟體方案, 並在成本與時程限制下與 IT 管理層及 desk 主管達成一致.
- 主動提出對業務或 IT 有益的改進建議.
- 確保交易平台的`效能與韌性 (performance and resiliency)`.
- 同時具備`交易面與 IT 面的風險意識`.
- 在既有程式碼與約定的限制內完成實作, 測試與部署.

`支援 (Support)`:

- 對已上線的方案負責 `Level 1 與 Level 2 支援`.
- 事故發生時提供良好的反應速度.
- 對 production 事故進行必要的後續追蹤, 以`預先偵測潛在的技術問題`.

`協作對象 (Collaboration)`:

- 坐在香港的 front office 使用者附近, 提供交易平台支援 (`註: 與職缺地點新加坡不一致, 需查證`).
- 與全球團隊成員及使用者聯繫互動.
- 與負責平台其他技術元件的內部 IT 團隊 (開發與支援) 協作.

## 任職要求 (Requirements)

`硬性要求 (Must-have)` — 原文 `Minimum qualifications (硬性)`:

- `Solid experience with C` (JD 原文如此, 未標示 `C++` 或 `C#`).
- 品質導向, 重視`軟體設計與架構`.
- 技術科系學士學位 (資訊, 數學, 工程).
- 能在整合型團隊中工作.

`加分要求 (Preferred)`:

原文 `Preferred qualifications (優先)`:

- 物件導向設計與常見設計模式的經驗.
- `多執行緒與並行 (multi-threading and concurrency)` 問題的經驗.
- `網路協定與 socket 程式設計`的知識.
- `股票與衍生性商品 (Stocks and Derivatives)` 產品經驗.
- 交易演算法與策略的軟體開發經驗.

原文 `Nice to have (加分)`:

- `Fixed Income & Currencies` - rates algo trading 或 FIC algo trading.
- `低延遲交易系統 (low latency trading systems)` 的開發經驗.
- 支援`面向客戶的交易系統`的經驗.
- 亞洲電子市場的特性知識, 或`市場連線 (market connectivity)` 經驗.
- `Windows 平台`的開發經驗.

`其他條件 (Other)`:

原文 `語言`: 英文.

原文 `行為能力 (Behavioral Skills)`:

- `Client focus`: 與客戶及內部夥伴培養長期關係.
- `Team Spirit`: 開放心態與尊重, 傾聽並分享專業見解.
- `Innovation - Technology`: 將技術突破納入所實施的策略.
- `Responsibility - Performance`: 產出卓越的個人與集體績效.

原文 `MCF 標記技能`: Level 2 Support, Teamwork, Multithreading, Development, Optimization of Costs,
Software Quality Engineering, Computer Science, Business Improvement, Derivative Products,
Trading Platforms, IT Risk Management, Technical Engineering, Incident Handling,
Socket Programming, Object-Oriented Design, Solution Deployment.

## 我的對位敘事 (Positioning Angle)

`前提`: 除非 `C` 確認為 `C++` 或 `C#` 的筆誤, 否則不建議投遞 — 硬性要求只有四條, 而技術那條完全不符,
履歷會在第一關被過濾.

若確認可投, 唯一站得住的敘事是 `我不是低延遲交易工程師, 我是把 desk 的軟體品質工程化的人`:
JD 的四條硬性要求中有三條是`品質, 架構與團隊`, 這正是履歷最強的部分 —
TSMC 的 30 行函式與 90% 覆蓋率標準, Change Healthcare 的共用框架與安全設計,
ByteDance 的跨資料中心韌性設計. 交易平台的 `performance and resiliency` 與
`風險意識` 這兩條, 可以用 ByteDance 的全球冗餘設計與合規遷移框架對位, 論證`我理解高風險系統的失效模式`.

領域缺口 (FIC, rates algo, market connectivity) 無法在面試前補齊, 應誠實承認並以
`我在半導體與醫療兩個高度受管制的領域都從零上手` 換取領域學習速度的信任.
