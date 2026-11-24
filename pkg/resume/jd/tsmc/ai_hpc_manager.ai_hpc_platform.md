---
company: "TSMC"
title: "AI HPC Manager"
department: "Information Technology (AI HPC)"
role_type: manager
role_facing: internal
domain: ml_infra
job_post_id: "18441"
url: "https://careers.tsmc.com/en_US/careers/JobDetail?jobId=18441"
source: direct
posted: 2026-03-25
expiry: null
fetched: 2026-09-08
created: 2026-09-08
min_years: null
education: null
score: 56
status: not_applied
status_updated: 2026-09-08
---
# TSMC - AI HPC Manager

> ## 評估摘要 (Evaluation Summary)
>
> `評估日期`: 2026-09-08, 對照 [Resume.md](../../Resume.md)
>
> `匹配度 (Profile Matching Score)`: `56 / 100`
>
> `結論`: 職稱掛 Manager, 但 JD 本文從第一句起就寫 `AI HPC Architect`, 九條要求全是技術深度而非帶人. 履歷有一個非常精準的鉤子 — TSMC 時期自建的 GPU 推論任務排程器 (把 n 個排隊推論任務裝進 m 個 GPU 資源, 有任務結束就立刻補位), 這正是第 4 條 performance tuning 與 workload management 的實作. 但缺口是`結構性的`: HPC 傳統技術棧 (MPI, OpenMP, CUDA, OpenCL) 履歷完全沒有, C++ 只到 LeetCode 等級, Python 非生產主力, 也沒有訓練側的模型開發經驗.
>
> `符合的部分`:
>
> - `系統架構與設計, 可擴充且高效率的架構`: ByteDance 跨資料中心備援與高階任務編排框架.
> - `效能調校與 workload management`: TSMC 的 GPU 推論排程器, 在固定加速器預算下維持 GPU pool 飽和且佇列持續消化.
> - `可擴充性與韌性 (冗餘, 負載平衡, failover)`: 履歷 Concept 欄明列 HA 與 scalability, ByteDance 那段即為全球備援設計.
> - `整合與部署, 與軟體工程及 DevOps 團隊協作把 AI 模型推上生產`: 履歷 MLOps 條目即為推論層的生產化.
> - `系統管理 (Linux, 監控工具, 叢集管理框架)`: Droi 的 Linux 底層, K8s 叢集, EFK, Prometheus, Grafana.
> - `程式與腳本能力`: Java 4.5 年, Golang, Bash 皆為生產主力.
> - `溝通與協作, 對技術與非技術對象傳達複雜概念`: Tech Lead 與跨部門推動經驗.
> - `持續學習`: 2026 年起自主橫跨八家模型供應商並自建 Agent SDK.
>
> `不符合的部分`:
>
> - `HPC 工具與框架 (MPI, OpenMP, CUDA, OpenCL)`: 履歷零. 這是本檔最大的硬缺口, 且不是幾週能補的.
> - `AI 與機器學習的深厚背景, 演算法與框架, 開發並部署 AI 模型`: 經驗集中在`推論服務與排程`, 不在模型本身; 無 PyTorch / TensorFlow 訓練經驗.
> - `HPC 架構與平行運算, 分散式運算, GPU 加速的深入知識`: 只有排程層的視角, 缺 interconnect (InfiniBand, NVLink), 平行檔案系統等 HPC 本業知識.
> - `硬體選型 (運算, 網路, 儲存)`: 未曾主導伺服器與網路硬體採購規格.
> - `Python 或 C++ 的熟練度`: Python 僅見於 LeetCode 與腳本, C++ 同; 主力為 Golang 與 Java.
> - `profiling 工具與 benchmarking`: 缺 HPC 等級的量測經驗.
>
> `投遞前必問`:
>
> - 這個職缺實際上是帶人 (Manager) 還是架構師 (JD 本文用 Architect), 團隊規模多少.
> - 服務對象是 R&D 的模型訓練叢集, 還是廠務端的推論工作負載? 兩者需要的技能完全不同.
> - 叢集規模與加速器型號, 是否有 Slurm 或 Kubernetes 為主的排程層.
> - MPI / CUDA 這條是硬門檻還是加分, 團隊裡是否已有專責的 HPC 工程師.

## 基本資訊 (Basics)

| 項目 | 內容 |
| --- | --- |
| 公司 (Company) | 台積電 (Taiwan Semiconductor Manufacturing Company Limited) |
| 職稱 (Title) | AI HPC Manager (JD 本文自稱 `AI HPC Architect`) |
| 部門 (Department) | Information Technology · Job Type `Manager / Executive` |
| 角色類型 (Role Type) | `manager` · `internal` · `ml_infra` |
| 地點 (Location) | 台灣 (Taiwan) |
| 職缺編號 (Job Post ID) | `18441` |
| 原始連結 (Original Link) | <https://careers.tsmc.com/en_US/careers/JobDetail?jobId=18441> |
| 刊登與截止 (Posted / Expiry) | 2026-03-25 刊登 |
| 薪酬 (Package) | 未公開 |
| 最低年資 (Min Experience) | 未載明 |
| 僱用型態 (Employment Type) | Regular |
| 投遞狀態 (Status) | `not_applied` (更新於 2026-09-08) |
| 抓取日期 (Fetched) | 2026-09-08 |

## 崗位職責 (Responsibilities)

- `方案設計`: 與資料科學家, 研究人員與 IT 團隊釐清需求並轉譯為 AI HPC 方案, 為 AI 工作負載設計兼顧效能, 可擴充性與效率的架構.
- `HPC 基礎設施規劃`: 評估並選用合適的硬體, 網路與儲存資源以支撐大規模 AI 運算, 考量平行處理, GPU 加速, interconnect 技術與儲存系統.
- `演算法最佳化`: 與資料科學與機器學習專家找出最佳化機會, 以平行處理, 分散式運算與 GPU 加速提升模型效能與效率.
- `效能調校`: 調整系統設定, 資源配置與工作負載管理, 找出並解除運算, 記憶體, 儲存或網路頻寬的瓶頸.
- `可擴充性與韌性`: 設計能承載大規模 AI 工作負載並容納未來成長的方案, 以冗餘, 負載平衡與 failover 確保高可用與容錯.
- `整合與部署`: 與軟體工程及 DevOps 團隊協作, 把 AI 模型與演算法整合進生產環境, 制定部署策略與工作流程.
- `研究與創新`: 追蹤 AI, ML 與 HPC 的最新進展, 評估新興技術, 框架與工具.
- `文件與溝通`: 記錄系統設計, 組態與效能最佳化, 並對技術與非技術對象有效傳達.

## 任職要求 (Requirements)

`硬性要求 (Must-have)` — 原文 `Job Qualifications`:

- `AI 與機器學習的深厚背景`: 深入理解概念, 演算法與框架, 具備開發並部署 AI 模型的經驗.
- `高效能運算專精`: 深入掌握 HPC 架構, 技術與最佳實務; 熟悉 HPC 框架, 平行處理, 分散式運算與 GPU 加速.
- `系統架構與設計`: 能設計可擴充且高效率的系統架構; 具備硬體元件, 網路技術與儲存系統的選型與組態經驗.
- `程式與腳本能力`: 熟練 `Python`, `C++` 或 `Java`; 熟悉 `Bash` 或 PowerShell 等腳本語言.
- `HPC 工具與框架`: 熟悉 `MPI`, `OpenMP`, `CUDA` 或 `OpenCL`, 並知道如何運用於平行運算與 GPU 加速.
- `效能最佳化`: 大規模運算系統的效能調校經驗, 熟悉 profiling 工具, benchmarking 與工作負載管理.
- `系統管理`: 理解 Linux / UNIX 作業系統, 系統監控工具與叢集管理框架.
- `溝通與協作`: 能與資料科學家, 研究人員, 軟體工程師與 IT 專業人員跨職能協作.
- `持續學習`: 追蹤研究論文, 參與研討會與相關社群.

`加分要求 (Preferred)`: JD 未區分, 無獨立 preferred 區塊.

## 缺口分析 (Gap Analysis)

`缺什麼`: HPC 本業技術棧. 履歷的 GPU 經驗是`排程與服務層` (哪個任務跑在哪張卡, 何時補位), JD 要的是`運算層` (單一任務如何在多張卡與多台機器上平行展開). 這兩者共用 GPU 這個詞, 但知識體系不同.

`缺了會怎樣`: 面試若進到 `如何為一個 8 卡節點的訓練任務選 all-reduce 策略` 或 `NCCL 與 InfiniBand 的拓撲影響` 這類問題, 會直接失守. 現有的排程器故事只能撐住推論側的討論.

`怎麼補`: 若要認真轉這條軌道, 順序是 (1) CUDA 程式模型與記憶體階層的基本觀念, (2) NCCL 與 collective communication 的語意, (3) Slurm 與 Kubernetes device plugin 兩種排程模型的對照. 這是`數月級`的投入, 不是投遞前能補完的.

`值不值得轉`: 若目標是 AI 基礎設施軌道, 這條的市場需求確實在成長, 但更務實的路徑是先從`推論服務層` (自己的既有強項) 站穩, 再往訓練側延伸. 參考同庫的 [redhat/ai_specialist_solution_architect.apac_technology_sales.md](../redhat/ai_specialist_solution_architect.apac_technology_sales.md) 與教材 [docs/tutorials/ai_inference_serving.md](../../docs/tutorials/ai_inference_serving.md).

## 我的對位敘事 (Positioning Angle)

`定位為拉伸型`, 面試本身有情報價值 — 能問出 TSMC 內部 AI 算力的實際規模與組織分工, 這是同批其他 AI 職缺 (AI Manager, AI Product Manager) 的重要背景資料.

`敘事主軸`: 只講一件事 — GPU 推論排程器. 說清楚問題 (固定加速器預算下佇列一直塞車), 做法 (n 對 m 的裝箱加上完成即補位), 結果 (GPU pool 保持飽和). 然後誠實把邊界畫出來: `我做的是排程層, 不是 kernel 層`, 並說明願意補的方向. 誠實畫界比硬撐 HPC 詞彙更能拿到下一輪.

`不要做的`: 不要在履歷或談話中把 MPI, CUDA 寫成會. 這種職缺的面試官就是 HPC 本業出身, 一問就破.
