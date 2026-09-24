---
company: "Meta"
title: "Software Engineer, Infrastructure (Singapore)"
department: "Engineering"
role_type: ic
role_facing: internal
domain: cloud_platform
job_post_id: null
url: null
source: direct
posted: null
expiry: null
fetched: 2026-09-10
created: 2026-09-10
min_years: 8
education: bachelor_cs
score: 80
status: applied
status_updated: 2026-09-10
---
# Meta - Software Engineer, Infrastructure (Singapore)

> ## 評估摘要 (Evaluation Summary)
>
> `評估日期`: 2026-09-10, 對照 [Resume.md](../../Resume.md)
>
> `匹配度 (Profile Matching Score)`: `80 / 100`
>
> `結論`: 這是`本庫所有 Meta 職缺中最貼合履歷的一檔`. 它與 [Software Engineer, Product](software_engineer_product.singapore.md) 是同一份通用職缺的兩個方向, 八條硬性要求幾乎完全相同, 但把`產品面向`換成`基礎設施`後, Product 版最大的扣分項 (客製化 UI 與消費者產品迭代) 整條消失, 換上的正是履歷近九年的重心. 唯一實質缺口是`日常語言`: 職責明寫 `primarily C/C++ and Python`, 而履歷的系統深度建立在 Java 與 Go 上.
>
> `符合的部分`:
>
> - `8 年以上程式開發經驗`: 12+ 年.
> - `6 年以上建置大規模基礎設施應用`: 這一條是履歷最厚的部分 - TikTok 的跨資料中心資料通道與全球備援, TSMC 的 Kubernetes AI 平台與 GPU 推論排程器, Droi 的 BaaS 平台 CI/CD, Change Healthcare 的 AWS 事件驅動架構. 四段皆為基礎設施本業, 合計遠超 6 年.
> - `成功領導重大倡議, 並具專案與團隊領導經驗`: ByteDance 的 Tech Lead 與錢包資料架構 roadmap 擁有權, Change Healthcare 帶 5 人並跨團隊協調, TSMC 兼任技術專案經理.
> - `跨多位工程師領導複雜技術工作`: change-region 合規框架被錢包以外的業務線採用, 這是跨團隊領導的直接證據.
> - `分析並改善系統資源的效率, 可擴充性與穩定性`: 全球資料中心備援`使實作與維護成本下降兩個 job grade`, 跨資料中心通道由多條一次性管線整併為單一 transport, GPU 排程器在固定加速器預算下維持資源池飽和.
> - `設計核心軟體元件與平台`: EventBus 共用框架, Spring Secret Manager plugin, Lambda 與事件註解框架, data SDLC 本身.
> - `透過程式碼審查, 測試, rollout, 監控與主動變更提升品質`: 覆蓋率 90% 以上與函式 30 行以內的專案標準, GitOps (Argo CD) 的可重現 rollout, EFK 與 Prometheus / Grafana, CloudWatch 與 SNS 告警.
> - `協助新成員上手並提供指導`: 30 場以上面試與 5 人團隊的帶人紀錄.
> - `具 Python, JavaScript 或 Hack 等腳本語言經驗`: JavaScript 為長期主力之一, Bash 列在技能表; `Python 僅有 LeetCode 使用紀錄`, 這一條算命中但偏弱.
> - `以資料與分析解釋技術問題`: Gamesofa 的資料工程, A/B testing 與使用者分群分析工具.
> - `資工學士學位`: 資工碩士.
> - `Java (加分)`: 4.5 年 Spring 全家桶.
> - `持續的 AI 技能發展: prompt / context engineering, agent orchestration (加分)`: 2026 年起的跨供應商 Agent SDK 與多模型編排, 直接命中.
> - `整合 AI 工具重新設計工作流 (加分)`: 有實作 (個人基礎設施的帳號, 支付與雲端編排), 但缺可量測數字, 見下.
>
> `不符合的部分`:
>
> - `以 C/C++ 與 Python 為主要開發語言`: 這是職責條目而非硬性要求, 但它定義了到職後的日常. 履歷的 C/C++ 只有 2016 年的 LeetCode 與 Google C++ education, 無生產環境經驗; Python 同樣只在 LeetCode. 面試的 coding 輪可以用 Java 或 Go 通過, 但團隊配對階段會被追問.
> - `AI 工具帶來可量測影響 (加分)`: AI 工作為個人專案, 沒有`效率提升 X%`或`缺陷率下降 Y%`這類可審計數字.
> - `負責任, 合乎倫理的 AI 實務 (加分)`: 風險評估與偏誤緩解皆無舉證.
> - `職涯連續性`: 2026 年 8 月離開 ByteDance 後為 self-directed, 需準備說法.
>
> `投遞前必問`:
>
> - 這是`通用職缺 (generic req)`, 團隊在流程後段才配對. 應確認可否指定`儲存, 資料基礎設施或平台方向`的團隊, 因為那是履歷最強的落點; 配到系統核心 (kernel, 網路堆疊, 編譯器) 方向會直接放大 C/C++ 缺口.
> - 實際職等對應 E5 還是 E6, 以及 8 年以上門檻對應哪一級.
> - 新加坡辦公室的基礎設施團隊組成與 headcount, 新加坡 PR 直接滿足工作權.
> - 面試 coding 輪是否接受 Java 或 Go 作答, 以及系統設計輪能否以跨資料中心與資料治理題材切入.
> - `+ 1 more` 的第二地點是哪裡, 是否影響團隊配對範圍.

## 基本資訊 (Basics)

| 項目 | 內容 |
| --- | --- |
| 公司 (Company) | Meta (Facebook, Instagram, WhatsApp, Messenger 母公司) |
| 職稱 (Title) | Software Engineer, Infrastructure (Singapore) |
| 部門 (Department) | Engineering |
| 角色類型 (Role Type) | `ic` · `internal` · `cloud_platform` |
| 地點 (Location) | 新加坡 (Singapore), 職缺頁另標 `+ 1 more` |
| 職缺編號 (Job Post ID) | 未提供 |
| 原始連結 (Original Link) | 未提供 (JD 全文由人工貼入) |
| 薪酬 (Package) | 未公開 |
| 最低年資 (Min Experience) | `8 年以上` 程式開發, 其中 `6 年以上` 為大規模基礎設施 |
| 學歷要求 (Min Qualification) | 資工, 資訊工程或相關技術領域學士, 或同等實務經驗 |
| 投遞狀態 (Status) | `applied` 於 2026-09-10, 目前階段 `Application` (更新於 2026-09-10) |
| 抓取日期 (Fetched) | 2026-09-10 (metacareers 對非瀏覽器請求回 HTTP 400, 全文為人工貼入) |

## 團隊背景 (Team Context)

這是 Meta 新加坡的`通用基礎設施工程職缺`, 未綁定特定團隊, 團隊配對在招募流程後段進行. JD 本文描述的是 Meta 的整體命題 (連結全球數十億使用者, 橫跨行動與網頁的多個平台), 而非某一條基礎設施產品線, 因此`要求的通用性本身就是這一檔的特徵`.

與 [Software Engineer, Product](software_engineer_product.singapore.md) 對照可見, 兩者是同一份 generic req 的兩個方向: 硬性要求八條中有七條逐字相同, 差異全在職責段 - Product 版要求`實作客製化使用者介面`, 本檔則要求`分析並改善系統資源的效率, 可擴充性與穩定性`並`以 C/C++ 與 Python 開發`.

## 崗位職責 (Responsibilities)

- `跨多位工程師`領導複雜的技術工作.
- 為團隊設定方向與目標, 涵蓋專案影響力, 產品品質與工程效率.
- 領導重大倡議, 專案, 團隊, 上線與`分階段釋出 (phased-releases)`.
- 與其他團隊介接, 吸收對方的創新成果, 反之亦然.
- 主動辨識並推動所負責的程式庫, 產品領域或系統所需的變更.
- 分析並改善各類系統資源的`效率, 可擴充性與穩定性`.
- 設計`核心軟體元件與平台`.
- 執行設計評審與程式碼審查.
- 協助新成員上手, 提供指導, 讓對方順利熟悉團隊程式庫.
- `主要以 C/C++ 與 Python 開發`.

## 任職要求 (Requirements)

`硬性要求 (Must-have)` — 原文 `Minimum Qualifications`:

- `8 年以上`相關程式語言的開發經驗.
- 成功領導重大倡議的經驗, 並具相應的專案與團隊領導經驗.
- `6 年以上`建置`大規模基礎設施應用`或同等經驗.
- 具透過`程式碼審查`, 適切`測試`, 正確 `rollout`, `監控`與主動變更來提升品質的經驗.
- 資工, 資訊工程, 相關技術領域學士學位, 或同等實務經驗.
- 具 `Python`, `JavaScript` 或 `Hack` 等腳本語言經驗.
- 具建置與交付`高品質工作`並達成`高可靠度`的經驗.
- 能運用`資料與分析`解釋技術問題, 並提供詳盡回饋與解法.

`加分要求 (Preferred)` — 原文 `Preferred Qualifications`:

- 具 `C, C++, Java` 等程式語言經驗.
- 能整合 AI 工具以優化或重新設計工作流, 並帶來`可量測的影響` (效率提升, 品質改善).
- 遵循並落實`負責任, 合乎倫理的 AI 實務` (如風險評估, 偏誤緩解, 品質與正確性審查).
- 持續的 AI 技能發展 (如 prompt / context engineering, agent orchestration), 並跟進新興 AI 技術.

`其他條件 (Other)`:

- Meta 為 Equal Employment Opportunity 雇主, 並提供身心障礙者的申請流程協助.
- 職缺頁標示 `+ 1 more` 的其他地點, 實際可選地點需於流程中確認.

## 我的對位敘事 (Positioning Angle)

這一檔應列為 `Meta 的主投`, 且優先於本庫既有的其他 Meta 職缺 - 它是唯一一檔`職責敘述與履歷近九年的工作幾乎逐條對應`的.

敘事順序: `我在 TikTok 擁有錢包資料架構的 roadmap, 把資料落在哪裡從一次性遷移專案變成平台屬性` -> `為此我對齊了新資料中心, 重新設計跨資料中心通道, 並設計交易資料的跨中心備援, 使實作與維護成本下降兩個 job grade` -> `這套 change-region 流程後來被錢包以外的業務線採用` -> `這正好是 JD 說的分析並改善系統資源的效率, 可擴充性與穩定性, 以及設計核心元件與平台`. 三句話覆蓋掉硬性要求的第二, 三, 七條.

`必須主動處理`的是語言: 不要等對方問. 正確做法是把 C/C++ 定位成`可補的工具層`而非能力層 - 明講履歷的系統深度來自 Java 與 Go 的生產環境經驗 (記憶體行為, 併發模型, 效能剖析的概念是共通的), 並在投遞前用 C++ 重寫一個既有元件 (例如 Agent SDK 的傳輸層或 GPU 排程器的核心邏輯) 作為可展示的佐證. 同時把 Python 從`LeetCode 語言`升級為`有作品的語言`, 這一條的成本極低但直接影響 JD 明列的日常語言條目.

`第二個要主動處理`的是 AI 加分項的量化: 現有 Agent SDK 與多模型編排在敘事上很強, 但缺數字. 投遞前為它訂出一組前後對比指標 (例如 provider 切換成本, 單一任務的人工介入次數), 這同時補上`可量測影響`這一條加分要求, 也對得上 Meta 近期所有 JD 都在問的 AI 主軸.

`團隊配對是本檔的實際風險`, 不是分數. 通用職缺的分發結果決定了 C/C++ 缺口是致命還是無關: 配到資料基礎設施或平台方向, 履歷是強候選; 配到系統核心方向, 語言缺口會直接主導評估. 因此`指定方向`應在 recruiter 第一通電話就提出, 而非等到 team matching 階段.
