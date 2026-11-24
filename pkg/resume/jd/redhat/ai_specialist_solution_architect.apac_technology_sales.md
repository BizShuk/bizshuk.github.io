---
company: "Red Hat"
title: "AI Specialist Solution Architect, Southeast Asia"
department: "APAC Technology Sales"
role_type: architect
role_facing: customer_facing
domain: ml_infra
job_post_id: null
url: null
source: direct
posted: 2026-09-02
expiry: null
fetched: 2026-09-07
created: 2026-09-07
min_years: 15
education: bachelor_cs
score: 40
status: not_applied
status_updated: 2026-09-07
---
# Red Hat - AI Specialist Solution Architect, Southeast Asia (新加坡)

> ## 評估摘要 (Evaluation Summary)
>
> `評估日期`: 2026-09-07, 對照 [Resume.md](../../Resume.md)
>
> `匹配度 (Profile Matching Score)`: `40 / 100`
>
> `結論`: `不建議直接投遞, 但值得作為學習標的`. 三項門檻級缺口同時存在: `15+ 年企業架構年資` (履歷 12+ 年), `7+ 年 ML 使用案例開發與 MLOps 實作` (履歷的 AI 資歷起於 2026/06, 且是 API 編排而非模型生命週期), 以及 `pre-sales / SSA 客戶面向軌道`. 與 [Dell 的硬體軌道職缺](../dell/senior_principal_solutions_architect.hyperscale_ai_infrastructure.md) 不同的是, 本職缺的技術缺口`落在 Kubernetes 之上的平台層`, 正好是履歷已站穩的地基, 屬`可自學補齊`而非需要換一段任職經歷. 中文流利與新加坡在地兩項則是明確加分.
>
> `符合的部分`:
>
> - `Kubernetes 設計建置, Operator, RBAC, GitOps`: TSMC 的 AI 平台與 Argo CD / Terraform 資歷, 是 OpenShift AI 的直接地基.
> - `Ansible 與 CI/CD 自動化 MLOps / LLMOps 流程`: Ansible 與 Azure DevOps / GitLab 皆在履歷上.
> - `端到端 AI 系統與分散式架構設計`: ByteDance 的跨資料中心資料通道, 全球備援與任務編排框架, 是可直接搬用的 distributed system 敘事.
> - `現代 AI 模式 (RAG, Agentic workflow, vector database)`: 2026 年起的 Agent SDK 與語音 agent 實作.
> - `LLM 提示工程與多供應商評估`: 跨八家模型供應商的比較實作.
> - `Mandarin fluency strongly preferred`: 母語, 這是 JD 明列的強加分項.
> - `CS 相關學位`: 資工碩士.
> - `新加坡在地 hybrid`: 已是新加坡 PR, 無搬遷成本.
>
> `不符合的部分`:
>
> - `15+ 年企業架構經驗`: 12+ 年, 差 3 年, 屬硬性年資缺口.
> - `7+ 年 ML 使用案例開發與 MLOps 實作`: 接近零. TSMC 的 GPU job scheduler 是排程層, 不是模型訓練與上線的生命週期.
> - `Deep Learning 架構 (CNN, RNN, LSTM) 的取捨論述能力`: 零. 這是 JD 用來篩 SME 的題目, 面試現場問得出來.
> - `Python 為主的統計程式語言與資料科學生命週期`: 履歷主力是 Java 與 Golang, Python 僅止於 LeetCode.
> - `Red Hat 產品面 (RHOAI, RHEL AI, InstructLab, Granite)`: 零實作.
> - `推論服務層 (vLLM, llm-d, KServe, ModelMesh)`: 零實作, 但為本檔缺口分析的補齊重點.
> - `PyTorch, TensorFlow, Jupyter 生態`: 零.
> - `pre-sales, PoC 主導, bake-off, C-level 簡報, 研討會演講`: 職涯軌道缺口, 與 [Anthropic](../anthropic/applied_ai_architect.technical_success.md) 與 [AWS](../amazon/senior_solution_architect.ai_engineering_asean_tech.md) 兩檔是同一道門檻.
>
> `投遞前必問`:
>
> - `15+ 年` 與 `7+ 年 MLOps` 是否為硬性篩選, 還是可用 distributed system 架構深度替代.
> - 團隊實際比重: `賣 OpenShift AI 平台` 與 `做 GenAI 解決方案架構` 各佔多少, 前者履歷接得上, 後者接不上.
> - 公共部門 (public sector) 客戶佔比與所需的安全審查, 影響入職前置時間.

## 基本資訊 (Basics)

| 項目 | 內容 |
| --- | --- |
| 公司 (Company) | Red Hat |
| 職稱 (Title) | AI Specialist Solution Architect, Southeast Asia |
| 部門 (Department) | APAC Technology Sales |
| 角色類型 (Role Type) | `architect` · `customer_facing` · `ml_infra` |
| 地點 (Location) | 新加坡 (Singapore), `hybrid`, 全職 |
| 職缺編號 (Job Post ID) | 未載明 |
| 原始連結 (Original Link) | 未附連結, 來源為使用者提供的 LinkedIn 職缺全文 |
| 薪酬 (Package) | 未公開 |
| 最低年資 (Min Experience) | `15 年` 企業架構 + `7 年` ML / MLOps |
| 招聘團隊 (Hiring Team) | Director, AI at Red Hat (Asia Pacific) |
| 投遞狀態 (Status) | `not_applied` (更新於 2026-09-07) |
| 抓取日期 (Fetched) | 2026-09-07 |

## 團隊背景 (Team Context)

Red Hat APAC Technology Sales 團隊, 職缺定位為 Red Hat AI 相關 `Go-To-Market (GTM)` 策略執行的
`技術重心 (technical center of gravity)`, 主要服務新加坡的公共部門, 企業與商用客戶.
任務是移除關鍵客戶內部在組織, 技術與競品上的採用障礙, 把 Red Hat 建立為企業 AI 工作負載的
首選平台. 角色需同時對話 C-level, 資料科學團隊與 DevOps 團隊, 屬典型 `專家型 SSA (Specialist SA)`,
與一般 account SA 的差別在於`只帶一條產品線但要帶到最深`.

## 崗位職責 (Responsibilities)

- 架構與設計: 規劃, 設計並執行端到端 AI 系統, 含資料管線, ML 管線與分散式訓練 / 服務架構.
- 技術銷售支援: 為 GTM 策略提供深度技術支援, 包含售前與售後的技術深潛, PoC, bake-off 與 hackathon.
- 利害關係人溝通: 對 C-level 到資料科學與 DevOps 團隊闡述 Red Hat AI 的價值主張.
- AI 營運化: 導入 AI 系統設計與 MLOps 最佳實務, 把客戶專案從實驗推進到可上線環境.
- 產品倡議: 作為現場與 Red Hat AI 事業部 / 工程團隊之間的雙向管道, 回饋市場意見並測試前沿功能.
- 社群推廣: 研討會演講, 技術部落格與技術社群志願參與.

## 任職要求 (Requirements)

`硬性要求 (Must-have)` — 原文 `What you will bring`:

- `SME 資格`: 15+ 年企業架構經驗, 其中 7+ 年為企業環境中 ML 使用案例開發與 MLOps 的實作經驗.
- `深度學習專業`: CNN, RNN, LSTM 等架構的基礎與應用知識, 並能說明其在真實場景中的取捨.
- `生成式 AI 架構`: 具備設計與部署 GenAI 解決方案架構的實績, 聚焦 LLM, 提示工程與微調策略.
- `技術深度`: 精通統計程式語言 (以 Python 為主) 與端到端資料科學生命週期.
- `架構能力`: 能對技術與非技術受眾架構並解釋複雜的資料 / ML 管線, 含分散式訓練與高擴充推論.
- `溝通`: 能為 C-level 與工程團隊主持商業價值場次與技術工作坊.
- `語言`: 英文流利為必要, 中文流利為強烈偏好.

`加分要求 (Preferred)` — 原文 `Education`:

- 電腦科學, 數學或相關技術領域學位.

`其他條件 (Other)` — 原文 `Technical Scope: The Red Hat AI Portfolio`:

- `平台掌握度`: Red Hat OpenShift AI (RHOAI) 與 Red Hat Enterprise Linux AI (RHEL AI) 實作經驗.
- `模型生態`: Granite 開源模型家族與 InstructLab 對齊調校方法論.
- `推論服務與執行期`: vLLM 推論最佳化, 以及 KServe, ModelMesh 等模型服務框架.
- `現代 AI 模式`: RAG, Agentic workflow 與 vector database 的實作經驗.
- `基礎架構即程式碼`: 以 Ansible 與 CI/CD 方案自動化 MLOps / LLMOps 流程.
- `開源框架`: PyTorch, TensorFlow 與 Jupyter 生態.
- 銷售標的為 2026 年 Red Hat AI 堆疊: Red Hat AI Enterprise, Red Hat AI Inference Server (vLLM 與
  llm-d, 跨 NVIDIA / AMD / Intel / Google TPU 加速器), Open Hybrid Cloud AI (bare metal 到 edge 的
  一致營運模型), 以及 Agentic AI 與編排.
- `績效指標 (KPI)`: PoC 成功率, 客戶從實驗到生產的轉換, AI 商機管道貢獻, 夥伴與內部團隊賦能,
  技術社群影響力, 以及對全球產品方向的現場回饋.

## 缺口分析 (Gap Analysis)

`缺什麼`: 三層缺口, 但層層的補法完全不同.

| 缺口 | 類型 | 補法 | 可補期 |
| --- | --- | --- | --- |
| 15+ 年企業架構年資 | 硬性數字 | 無法補, 只能靠架構深度與 SSA 面試表現換取例外 | - |
| 7+ 年 ML / MLOps 實作 | 年資類型 | 無法在短期補足年份, 但可用`可展示的推論平台實作`置換敘事 | 中期 |
| DL 架構 (CNN / RNN / LSTM) 論述 | 知識 | 純知識題, 可讀通並準備取捨說法 | 短期 |
| 推論服務層 (vLLM, llm-d, KServe) | 技能 | 站在既有 Kubernetes 地基上自建可跑的 demo | 短期 |
| Red Hat 產品面 (RHOAI, RHEL AI) | 產品 | OpenShift 開發者沙箱與 InstructLab 本機安裝 | 短期 |
| pre-sales 客戶面向軌道 | 職涯類型 | 需一段任職經歷或以技術社群演講 / 部落格代償 | 長期 |

`缺了會怎樣`: 履歷會被歸類為 `平台工程師 (Platform Engineer)` 而非 `AI SME`.
本職缺篩的是`能在客戶白板前把 KV cache, batching 與 TTFT / TPOT 取捨畫出來的人`,
Kubernetes 熟練度只是入場券.

`怎麼補`: 技術缺口中影響最大且最快見效的是 `推論服務層`, 因為它同時命中
`Deep Learning 取捨`, `GenAI 架構` 與 `分散式高擴充推論` 三項硬性要求.
細節見 [docs/tutorials/ai_inference_serving.md](../../docs/tutorials/ai_inference_serving.md),
該文涵蓋 vLLM 推論最佳化, llm-d 的跨加速器分散式服務, KServe 與 ModelMesh 的定位差異,
以及生成式 AI 服務的端到端上線架構.

`值不值得轉向`: 值得, 但不是靠這一檔.
這條軌道 (`AI 平台 / 推論基礎架構`) 與履歷的距離, 明顯近於低延遲交易或支付領域,
且與 [AWS Senior Solution Architect, AI Engineering](../amazon/senior_solution_architect.ai_engineering_asean_tech.md) `66` 分那檔屬同一條軌道.
合理路徑是`先投 AWS 那檔`, 把 Red Hat 這檔的內容當成該面試的技術準備素材.

## 我的對位敘事 (Positioning Angle)

若要投, 敘事不應該假裝有 ML 年資, 而是`把 12 年架構經驗接到 AI 工作負載的營運面`:

- `分散式系統可信度`: ByteDance 的跨資料中心資料通道與全球備援, 對位 `分散式訓練 / 高擴充推論`
  所需的容錯, 流量調度與一致性思維. 推論叢集的 P/D 拆分本質上是同一類問題.
- `平台落地可信度`: TSMC 的 AI 平台架構與 GPU job scheduler, 是 `把 GPU 資源交付給資料科學家`
  的真實經驗, 正是 RHOAI 要賣的東西.
- `框架推動可信度`: 兩次成功推動跨部門框架採用, 對位 SSA 的 `移除組織採用障礙`.
- `多模型評估可信度`: 跨八家供應商的實測比較, 是 bake-off 的原型能力.
- `語言與在地`: 中文母語加新加坡 PR, 對 SEA 公共部門與企業客戶是直接優勢.

不應該碰的說法: 不要宣稱 DL 模型訓練經驗, 也不要把 API 編排說成 MLOps.
面對 `你 fine-tune 過什麼模型` 這題, 誠實答`沒有生產級微調, 但能說清楚何時該 RAG, 何時該
fine-tune, 何時該只換 prompt 與量化`, 這個取捨判斷本身就是 SSA 的日常工作.
