---
company: "Apple"
title: "Senior Site Reliability Engineer"
department: "Apple B2B Engineering (供應鏈夥伴整合)"
role_type: ic
role_facing: internal
domain: cloud_platform
job_post_id: "200675997"
url: "https://jobs.apple.com/en-sg/details/200675997-3278/senior-site-reliability-engineer"
source: direct
posted: 2026-08-05
expiry: null
fetched: 2026-11-24
created: 2026-09-02
min_years: 5
education: null
score: 80
status: closed
status_updated: 2026-11-24
---
# Apple - Senior Site Reliability Engineer (B2B)

> ## 評估摘要 (Evaluation Summary)
>
> `評估日期`: 2026-09-02, 對照 [Resume.md](../../Resume.md)
>
> `匹配度 (Profile Matching Score)`: `85 / 100`
>
> `結論`: 要求清單長達 19 條, 履歷命中其中約十六條, 且 `從零建置 DevOps 平台`, `GitOps (ArgoCD)` 與 `IaC` 三條在履歷中都有直接對應 (Droi 的 BaaS CI/CD 與 IaC, TSMC 的 Argo CD 與 Terraform). 扣分在於這是 `資深 IC (ICT4)` 而非 lead, 且職責含 on-call.
>
> `符合的部分`:
>
> - `5 年以上 SRE, DevOps 或基礎設施角色`: Droi 的 Cloud Engineer 加 TSMC 的平台工作, 加 ByteDance 的基礎設施建設.
> - `從零建置 DevOps 平台的實績` 與 `端對端設計建置 DevOps 平台`: Droi 從零建立 BaaS 平台的 CI/CD 模型 (Jenkins + Docker + Kubernetes) 並自建端對端測試框架與 CLI, 這一條是本檔的最強命中.
> - `GitOps 部署模型 (ArgoCD / Flux) 的實作經驗`: TSMC 以 Argo CD 從 Git 對帳叢集狀態.
> - `建立 Infrastructure as Code (IaC) 實務`: TSMC 以 Terraform 描述雲端與叢集資源, Droi 以 Terraform 加 Ansible 讓環境可重建.
> - `建置並維運 Kubernetes 平台 (EKS / AKS / GKE / OpenShift)`: TSMC 與 Droi 兩段 K8s 生產經驗, AWS 有 SAA 認證.
> - `大規模分散式系統工作經驗`: ByteDance 跨資料中心.
> - `使用 Splunk, Grafana, Prometheus 等監控與可觀測性工具實作遙測`: EFK, Prometheus, Grafana 為履歷明列項.
> - `從作業系統與應用兩個視角解決 Kubernetes 問題`: Droi 的 Linux 底層加 TSMC 的 K8s 平台.
> - `強 Java 程式經驗, 建置 web, middleware 或後端應用`: Change Healthcare 四年 Spring Boot.
> - `深入理解關聯式資料庫與 NoSQL (MongoDB)`: RDB 與 MongoDB 皆為履歷明列項.
> - `HTTP/S, TCP, DNS, web 應用負載平衡知識`: Droi 的 Nginx 與 Linux 網路工作.
> - `安全概念與協定: 認證, 授權, 簽章, 加密, SSL/TLS, SSH/SFTP, PKI, X509, PGP`: JWT 授權, Jenkins KMS, Spring Secret Manager plugin.
> - `主導事故回應, RCA 與可靠性改善`: Tech Lead 經驗.
> - preferred `service mesh (Istio, Linkerd)`: TSMC 明確使用 Istio.
> - preferred `多雲或混合雲環境`: AWS 四年加 Azure DevOps.
> - preferred `大規模高度分散的生產環境`: ByteDance.
>
> `不符合的部分`:
>
> - `Python 等程式語言的熟練度`: Python 僅見於 LeetCode, 非生產主力 (主力為 Golang 與 Java).
> - `支援對外生產服務, 含部署, on-call 與事故管理`: on-call 輪值的實際承擔經驗不足, 且需評估生活型態.
> - preferred `容器編排內部機制 (scheduler, CNI, CSI)`: 未曾深入插件層.
> - preferred `進階可觀測性: 分散式追蹤, SLO/SLI 設計, error budget`: 有可觀測性工具經驗但缺 SLO 與 error budget 的制度化實作.
> - preferred `middleware 平台如 WebMethods Integration Server`: 零.
> - `職等 ICT4 為資深 IC`, 與往 lead 與 manager 發展的方向不一致.
>
> `投遞前必問`:
>
> - 職等 ICT4 的定位: 是資深 IC 還是有 tech lead 成分, 是否有帶人路徑.
> - on-call 的實際輪值頻率與夜間覆蓋安排.
> - JD 說 `從零設計建置現代 DevOps 與 SRE 生態系`, 這是全新建置還是既有平台改造? 目前的成熟度為何.
> - `管理生產與非生產環境的機器學習模型生命週期` 這條在日常工作中的比重.
> - IaC 目前使用什麼工具, 是 Terraform 還是 Pulumi, state 與 module 的管理方式為何.

## 基本資訊 (Basics)

| 項目 | 內容 |
| --- | --- |
| 公司 (Company) | Apple |
| 職稱 (Title) | Senior Site Reliability Engineer |
| 部門 (Department) | Apple B2B Engineering (供應鏈夥伴整合) · 內部職等 `Software Engineering Applications ICT4` |
| 角色類型 (Role Type) | `ic` · `internal` · `cloud_platform` |
| 地點 (Location) | 新加坡 (Singapore) |
| 職缺編號 (Job Post ID) | `200675997` |
| 原始連結 (Original Link) | <https://jobs.apple.com/en-sg/details/200675997-3278/senior-site-reliability-engineer> (Role Number `200675997-3278`, 張貼於 2026-08-05) |
| 刊登與截止 (Posted / Expiry) | 2026-08-05 刊登 |
| 薪酬 (Package) | 未公開 |
| 最低年資 (Min Experience) | `5 年以上` |
| 投遞狀態 (Status) | `not_applied` (更新於 2026-09-03) |
| 抓取日期 (Fetched) | 2026-09-02 |

## 崗位職責 (Responsibilities)

- 加入實作並維護一流 DevOps 實務的團隊, 處理與 Apple B2B 系統的可規模化, 可靠性與效能相關的複雜技術挑戰.
- 管理生產與非生產環境中機器學習模型的生命週期.
- 持續評估並改善系統流程, 偵測異常, 找出優化空間並實作解法以提升系統可靠性與效能.
- 對運行環境 (JVM, 作業系統, 檔案系統, 網路協定) 有良好的概念理解.
- 與 Apple 內部的技術與非技術團隊, 以及外部供應鏈夥伴協作.
- `從零設計, 建置並規模化現代 DevOps 與 SRE 生態系`; 建立以 GitOps 驅動的雲原生 CI/CD 平台.
- 擔任 foundational engineer 與技術領導者, 為整個組織定義標準, 工具, 自動化與可靠性實務.

## 任職要求 (Requirements)

`硬性要求 (Must-have)` — 原文 `Minimum Qualifications`:

- `5 年以上` SRE, `DevOps (必要)` 或基礎設施導向角色的實績經驗.
- 端對端設計並建置 DevOps 平台, 並同時涵蓋 SRE 與平台工程.
- 從零建置 DevOps 平台的實績.
- `GitOps` 部署模型 (`ArgoCD` / `Flux`) 的實作經驗.
- 建立 `Infrastructure as Code (IaC)` 實務.
- 建置並維運 Kubernetes 平台 (`EKS` / `AKS` / `GKE` / `OpenShift`).
- 大規模分散式系統的工作經驗.
- 強問題解決與架構能力.
- 熟練一種以上程式語言 (如 `Python`).
- 透過部署, on-call 與事故管理支援對外生產服務與分散式系統.
- 主導事故回應, RCA 與可靠性改善.
- 熟練使用 `Splunk`, `Grafana`, `Prometheus` 或類似工具實作與協調遙測.
- 能從作業系統與應用兩個視角解決 Kubernetes 問題.
- 建置並維運 Kubernetes 或 EKS 等容器編排系統.
- 強 `Java` 程式經驗, 建置 web, middleware 或後端應用.
- 深入理解 Oracle 或類似關聯式資料庫, 以及 MongoDB 等 NoSQL 資料庫.
- 應用與資料庫效能調校的第一手經驗.
- HTTP/S, TCP, DNS, web 應用負載平衡的知識.
- 深入理解基本安全概念與協定: 認證, 授權, 簽章, 加密, SSL/TLS, SSH/SFTP, PKI, X509 憑證與 PGP.

`加分要求 (Preferred)` — 原文 `Preferred Qualifications`:

- 強 Java 程式經驗 (後端, middleware 或 web 應用).
- NoSQL 資料庫經驗 (MongoDB, Cassandra, DynamoDB 等).
- 深入理解關聯式資料庫 (Oracle, PostgreSQL, MySQL 等).
- 應用與資料庫效能調校的實作經驗.
- 進階可觀測性實務: 分散式追蹤, SLO/SLI 設計, error budget.
- 大規模, 高度分散生產環境的經驗.
- 容器編排內部機制經驗 (scheduler, CNI, CSI 等).
- middleware 平台知識, 如 WebMethods Integration Server 或類似.
- 多雲或混合雲環境經驗.
- 熟悉 service mesh 技術 (Istio, Linkerd 等).

## 我的對位敘事 (Positioning Angle)

這一檔與同公司的 Engineering Manager 一檔可以`同時投遞`, 因為分屬不同組織 (B2B 與 Apple Store Online) 且職等路線不同 (IC 與 Manager). 若 Apple 有意願, 兩條線都開著能提高進場機率.

`敘事主軸`: JD 用了 `foundational engineer`, `from scratch`, `defining standards` 這幾個詞, 代表他們要的是建立者而非維運者. 履歷有兩個從零建立的實例可以直接對應: Droi 的 BaaS CI/CD (整合 Jenkins, Docker, Kubernetes) 與自建的端對端測試框架, 以及 TSMC 的 Azure CI/CD 流程與工程標準. 這兩段應該放在最前面.

`要補的`: `SLO/SLI 與 error budget` 是概念性的, 讀完 Google SRE Book 的相關章節並準備好一個 `如果讓我為某個服務定 SLO 我會怎麼做` 的答案即可.
