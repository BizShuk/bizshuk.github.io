---
company: "OKX"
title: "Staff/Senior Staff Engineer, Kubernetes"
department: "Engineering / Cloud Native Infrastructure"
role_type: staff_ic
role_facing: internal
domain: cloud_platform
job_post_id: "7767880003"
url: "https://job-boards.greenhouse.io/okx/jobs/7767880003"
source: greenhouse
posted: null
expiry: null
fetched: 2026-11-24
created: 2026-09-02
min_years: 4
education: bachelor_cs
score: 68
status: not_applied
status_updated: 2026-09-03
---
# OKX - Staff/Senior Staff Engineer, Kubernetes

> ## 評估摘要 (Evaluation Summary)
>
> `評估日期`: 2026-09-02, 對照 [Resume.md](../../Resume.md)
>
> `匹配度 (Profile Matching Score)`: `68 / 100`
>
> `結論`: Kubernetes 與 Istio 的核心要求命中, 但這是一檔 `深度維運` 而非 `架構設計` 的職缺, 且雙雲要求含阿里雲 ACK. 適合作為備選, 不建議當主力.
>
> `符合的部分`:
>
> - `4+ 年生產級 Kubernetes 叢集操作`: TSMC 與 Droi 皆為 K8s 生產環境.
> - `K8s 核心原理與元件 (Pod, Deployment, StatefulSet, Service, Ingress, CRD, controller, 調度策略, 網路模型, 儲存掛載)`: 具備.
> - `Istio / Envoy service mesh, 東西向流量治理, 灰度發布`: TSMC 明確使用 Istio.
> - `精通 Linux 系統管理, 系統調優, 權限控制, 行程管理, 日誌分析與線上排錯`: Droi 的 Linux/Nginx/Bash 底層工作.
> - `AWS: EKS, EC2, S3, VPC, IAM, 負載平衡, CloudWatch, 安全策略`: 四年 AWS 實作 + SAA 認證.
> - `容器 runtime (containerd/Docker)`: 具備.
>
> `不符合的部分`:
>
> - `阿里雲 ACK, ECS, SLB, VPC, RAM, RDS, OSS, CloudMonitor` 為明列的雙雲之一, 履歷零經驗.
> - `CNI (Calico/Flannel)` 與 `CSI` 的插件層細節未曾深入.
> - 職務性質偏 SRE 與叢集維運, 與履歷近年往架構與領導發展的方向相反, 有 title 與職涯的錯位風險.
>
> `投遞前必問`:
>
> - 這個職缺是否含 on-call 輪值與夜間故障處理.
> - 叢集規模 (節點數, 叢集數), 以及阿里雲與 AWS 的實際比重.
> - Staff 與 Senior Staff 的分野, 是否有帶人的成分.
> - OKX 明訂 `will be prioritising applicants who have a current right to work in Singapore`, 新加坡 PR 直接滿足.

## 基本資訊 (Basics)

| 項目 | 內容 |
| --- | --- |
| 公司 (Company) | OKX (加密貨幣交易所, 隸屬 OKG 集團, 產品含 OKX Wallet, OKLink) |
| 職稱 (Title) | Staff/Senior Staff Engineer, Kubernetes |
| 部門 (Department) | Engineering / Cloud Native Infrastructure |
| 角色類型 (Role Type) | `staff_ic` · `internal` · `cloud_platform` |
| 地點 (Location) | 新加坡 (Singapore) |
| 職缺編號 (Job Post ID) | `7767880003` |
| 原始連結 (Original Link) | <https://job-boards.greenhouse.io/okx/jobs/7767880003> |
| 薪酬 (Package) | 未公開 |
| 最低年資 (Min Experience) | `4 年以上` |
| 學歷要求 (Min Qualification) | 資訊工程相關學士以上 |
| 投遞狀態 (Status) | `not_applied` (更新於 2026-09-03) |
| 抓取日期 (Fetched) | 2026-09-02 |

## 崗位職責 (Responsibilities)

JD 主體為要求清單, 職責為: 維運生產級 Kubernetes 叢集, 獨立排除複雜叢集故障與效能瓶頸, 執行阿里雲與 AWS 雙雲的生產環境維運與災難復原, 管理多叢集, service mesh 與網路儲存插件.

## 任職要求 (Requirements)

`硬性要求 (Must-have)` — 原文 `What We Look For In You`:

- 資訊相關科系學士以上; `4+ 年` 生產級 Kubernetes 叢集操作經驗; 精通 K8s 核心原理與元件, 含 `Pod`, `Deployment`, `StatefulSet`, `Service`, `Ingress`, `CRD`, controller, 調度策略, 網路模型與儲存掛載; 能獨立解決複雜叢集故障與效能瓶頸.
- 精通阿里雲與 AWS 雙雲維運, 具備雙雲生產環境的獨立經驗:
  - `阿里雲`: ACK 容器服務, ECS, SLB, VPC, RAM, RDS, OSS, CloudMonitor, 安全群組與快照備份.
  - `AWS`: EKS, EC2, S3, VPC, IAM, TGW, 負載平衡, CloudWatch 與安全策略; 具海外雲端部署, 維運與災備的實務經驗.
- 精通 Linux 系統管理, 熟悉系統調優, 權限控制, 行程管理, 日誌分析與線上排錯.
- 熟悉主流容器 runtime (`containerd` / `Docker`); 理解 K8s 網路 (`CNI`, 如 Calico / Flannel), 儲存 (`CSI`) 與多叢集管理; 熟悉 `Istio` / `Envoy` service mesh, 東西向流量治理與灰度發布.

## 我的對位敘事 (Positioning Angle)

只有在其他四檔 OKX 職缺都沒有進展時才投這一檔. 若要投, 強調 TSMC 的 Istio 與 Argo 實作以及 Droi 從零建立 K8s CI/CD 的經驗, 並明確表態阿里雲可在到職前自學 (ACK 與 EKS 的心智模型高度重疊). 同時要在面試中確認這個職缺不會把人鎖死在維運線上.
