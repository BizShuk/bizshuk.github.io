---
company: "Airwallex"
title: "Senior Cloud Infra Engineer (Cloud Networking)"
department: "Engineering / Infrastructure & Productivity (Global Infrastructure)"
role_type: ic
role_facing: internal
domain: cloud_platform
job_post_id: "11249329"
url: "https://builtin.com/job/senior-cloud-infrastructure-engineer-cloud-networking/11249329"
source: builtin
posted: 2026-09-18
expiry: 2026-10-18
fetched: 2026-11-24
created: 2026-09-22
min_years: 5
education: bachelor_cs
score: 84
status: closed
status_updated: 2026-11-24
---
# Airwallex - Senior Cloud Infra Engineer (Cloud Networking)

> ## 評估摘要 (Evaluation Summary)
>
> `評估日期`: 2026-09-22, 對照 [Resume.md](../../Resume.md)
>
> `匹配度 (Profile Matching Score)`: `84 / 100`
>
> `結論`: `高度契合的雲端網路與基礎設施資深職缺`. 履歷在 TSMC (Terraform, K8s, Istio, GitOps) 與 Droi (Terraform, Ansible, K8s) 具備扎實的雲端 IaC 自動化經驗，且在 ByteDance 擁有跨區域跨資料中心 (Cross-DC) 傳輸通道與資料備援的架構實績。Airwallex 作為新加坡高成長金融科技獨角獸，此職位主導 AWS/GCP 跨雲骨幹與銀行端直連線路，技能重疊度極高。
>
> `符合的部分`:
>
> - `Infrastructure as Code (IaC)`: 精通 Terraform，Droi 與 TSMC 兩段均以 Terraform 為單一真理源進行雲端與叢集資源宣告。
> - `跨雲與網路基礎設施`: AWS (VPC, TGW, Direct Connect, Network Firewall) 與 GCP 概念；具備 Linux 系統調優、Shell/Bash/Python 腳本能力。
> - `大規模分散式架構與資料通道`: ByteDance 錢包業務期間主導跨資料中心通道重構，與金融直連網絡、BGP/IPsec 的高可用容錯需求完全相符。
> - `CI/CD 與 GitOps`: Droi (Jenkins + Docker + K8s) 與 TSMC (Argo CD GitOps + Azure DevOps)。
> - `金融級高可用與合規`: 具備新加坡 PR 身分，無工作簽證摩擦；ByteDance 錢包合規與歐盟資料合規落地經驗可直接遷移至 Airwallex 支付平台。
>
> `不符合的部分`:
>
> - `硬體/電信級網路認證`: JD 偏好 CCNA/JNCIA 或 AWS Advanced Networking Specialty 認證，履歷持有的是 AWS Solutions Architect Associate (SAA)。
> - `實體 POP 點與硬體網路設備`: 本職缺涉及部分區域 POP 點與銀行邊界路由直連，履歷背景更偏向純雲端軟體定義網路 (SDN/Cloud Native)。
> - `輪值待命 (On-call)`: 該職位需參與全球 SRE/網路 On-call 輪值。
>
> `投遞前必問`:
>
> - AWS 與 GCP 在 Airwallex 的架構比重為何，跨雲流量的主要治理策略。
> - 網路團隊在新加坡的編制規模，與墨爾本/香港基礎設施團隊的分工模式。
> - 辦公型態：JD 標記 In-Office (Guoco Tower)，需確認是否有每週 1-2 天彈性遠距空間。

## 基本資訊 (Basics)

| 項目 | 內容 |
| --- | --- |
| 公司 (Company) | Airwallex (全球跨境支付獨角獸，估值 US$11B，全球 2,300+ 人) |
| 職稱 (Title) | Senior Cloud Infra Engineer (Cloud Networking) |
| 部門 (Department) | Global Infrastructure Team |
| 角色類型 (Role Type) | `ic` · `internal` · `cloud_platform` |
| 地點 (Location) | 新加坡 (1 Wallich St, #20-01 Guoco Tower, Singapore 078881) |
| 職缺編號 (Job Post ID) | `11249329` |
| 原始連結 (Original Link) | <https://builtin.com/job/senior-cloud-infrastructure-engineer-cloud-networking/11249329> |
| 薪酬 (Package) | 未公開（新加坡市場同級約 SGD 138k - 192k 底薪 + 獎金 + Pre-IPO 股權） |
| 最低年資 (Min Experience) | 5 年以上 |
| 學歷要求 (Min Qualification) | Bachelor's in CS / Engineering 或同等實務經驗 |
| 投遞狀態 (Status) | `not_applied` (更新於 2026-09-22) |
| 抓取日期 (Fetched) | 2026-09-22 |

## 團隊背景 (Team Context)

Airwallex 基礎設施團隊負責構建與擴展全球分散式金融平台，確保跨多雲環境的高可用性、嚴格安全邊界與雲端彈性擴展，同時賦能全球研發團隊的交付生產力。

## 崗位職責 (Responsibilities)

- **雲端網路架構設計**：設計、構建與優化高可用、多帳號與多區域之 AWS 及 GCP 網路架構 (L3 routing, Transit Gateway, Shared VPC, Cloud Router, Direct Connect)。
- **基礎設施即代碼 (IaC)**：透過 Git-driven CI/CD 流水線，使用 Terraform 自動化編排管理全球雲端網路。
- **混合雲與跨雲直連**：管理公有雲 (AWS/GCP)、區域 POP 點與外部清算銀行夥伴間的 BGP 動態路由、IPsec VPN 與線路調度，符合嚴格金融監管要求。
- **維運卓越與可觀測性**：作為複雜雲端網路事故的二線升級核心，優化遙測與可觀測性，並參與全球 On-call 輪值。
- **跨職能協同**：與 Security、DevOps、Platform 團隊緊密合作，建立安全的 Landing Zones 與自動化網路防禦機制。

## 任職要求 (Requirements)

- 熟練掌握大規模 AWS 網路架構設計與營運 (VPC, TGW, Route Tables, Direct Connect, AWS Network Firewall)。
- 精通 Terraform 或同等 IaC 自動化工具。
- 具備 GCP 網路概念 (Shared VPC, Cloud Router) 或多雲架構實作經驗。
- 熟悉 Linux 發行版、Shell/Python 腳本自動化及 CI/CD 工具鏈。
- 紮實的網路基礎：核心路由交換、子網劃分、TCP/IP 與 BGP 動態路由協定。
