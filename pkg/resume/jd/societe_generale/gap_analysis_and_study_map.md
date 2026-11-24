# 能力缺口分析與角色地圖 (Gap Analysis and Role Map)

對照 [front_office_software_engineer.fic_automated_trading.md](front_office_software_engineer.fic_automated_trading.md)
的「不符合的部分」, 逐項說明`那是什麼`, `缺了會怎樣`, 以及`怎麼補`.

`評估日期`: 2026-09-03
`對照履歷`: [Resume.md](../../Resume.md)
`職缺截止`: 2026-09-16 (`本檔的學習路徑趕不上這一檔投遞, 是為下一次 sell-side 機會準備的`)

---

## 先確認手上已有的槓桿 (Transferable Assets)

補缺口之前先確認哪些既有經驗可以直接折算, 否則會把已經有的東西重學一遍.

| 既有經驗 | 對應 JD 的哪一條 | 折算後的定位 |
| --- | --- | --- |
| TSMC 的 30 行函式上限, 90% 覆蓋率, 架構明確化 | 硬性要求 2 `品質導向, 重視軟體設計與架構` | `直接可用`, 是唯一能正面對打的硬性要求 |
| Change Healthcare 的共用框架 (Lambda, EventBus, Secret Manager) | 優先項 `物件導向設計與常見設計模式` | `直接可用` |
| ByteDance 跨資料中心通道與全球冗餘設計 | 職責 `確保交易平台的效能與韌性` | `部分可用`, 韌性思維可用, 延遲量級不同 |
| Change Healthcare 的 CloudWatch, SNS 告警與 Fortify 修補 | 職責 `L1/L2 支援與 production 事故追蹤` | `直接可用` |
| Change Healthcare 的 KMS, JWT, 安全設計 | 職責 `交易面與 IT 面的風險意識` (IT 面) | `部分可用`, 只涵蓋 IT 風險, 不含交易風險 |
| 中央資工碩士 + 東華資工學士 | 硬性要求 3 `技術科系學士學位` | `直接可用`, 超出要求 |
| 12 年經驗, 跨半導體, 醫療, 內容平台 | MCF 欄位 `6 年以上` | `直接可用` |

結論: `工程品質層與事故處理層已經到位, 缺的全部集中在系統程式語言層與交易領域層`.
這兩層彼此獨立, 可以並行補, 但`語言層是門檻, 領域層是加分`, 投入順序不能顛倒.

---

## 缺口一: C / C++ 系統程式語言 (門檻級)

### 這是什麼

sell-side front office 的交易系統以 `C++` 為主 (少數遺留元件為 `C`), 原因不是偏好而是約束:
需要`手動控制記憶體配置時機`, 避免 GC 造成的不可預期停頓, 並在 microsecond 量級上做確定性的延遲.
與履歷熟悉的 Golang, Java 的本質差異在:

- `記憶體模型`: 手動生命週期管理, RAII, smart pointer, 沒有 runtime 幫你回收.
- `零成本抽象`: template 與編譯期多型, 抽象不能有 runtime 代價.
- `可預期性`: 熱路徑上禁止動態配置, 禁止例外, 禁止系統呼叫.

### 缺了會怎樣

`這是唯一的硬性技術要求, 缺了等於整份履歷不進入審閱`. 具體後果:

- 履歷篩選階段直接淘汰, 連對位敘事的機會都沒有.
- 即使進入面試, 現場 coding 一律以 C++ 出題, Go 的寫法在此無法折算.
- 無法閱讀既有 codebase, 而 JD 明寫`在既有程式碼內實作`.

### 怎麼補

最小可驗證成果: `用 modern C++ (17/20) 重寫一個履歷上已完成的 Go 專案,
並在 GitHub 上留下含 benchmark 與單元測試的完整倉庫`.

- 語言核心: RAII, move semantics, smart pointer, template 基礎, `std::atomic` 與 memory order.
- 工具鏈: CMake, `perf`, `valgrind`, Google Benchmark, sanitizer (ASan, TSan).
- 讀物: `Effective Modern C++`, CppCon 上的低延遲交易演講 (Carl Cook 的
  `When a Microsecond Is an Eternity` 是這個領域的共同語彙起點).
- 節奏: 這是`以季為單位`的缺口, 不是面試前速讀能補的.

---

## 缺口二: 低延遲並行與 socket 程式設計

### 這是什麼

履歷的並行經驗停在`應用層的吞吐量`, 交易系統要的是`尾端延遲的確定性`. 差別在:

| 層級 | 履歷已有 | 交易系統要求 |
| --- | --- | --- |
| 並行原語 | goroutine, channel, mutex | `lock-free` 佇列, `std::atomic`, memory ordering |
| 排程 | Kubernetes 排程器決定 | `CPU pinning`, isolcpus, NUMA-aware 配置 |
| 網路 | HTTP, gRPC, Istio service mesh | 原生 `socket`, `epoll`, `TCP_NODELAY`, kernel bypass |
| 衡量 | QPS, p99 延遲 (毫秒) | p99.9 延遲 (微秒), jitter, `cache line` 爭用 |

### 缺了會怎樣

這是 JD 兩條優先項 (`多執行緒與並行`, `網路協定與 socket 程式設計`) 的共同底座:

- 面試時被問到「兩個執行緒之間傳一筆訂單, 你怎麼做到不用鎖」會完全答不出來.
- 無法判斷既有平台的效能瓶頸, 而職責明寫`確保交易平台的效能與韌性`.
- L1/L2 支援時看到延遲尖峰, 無法區分是 GC, 是 context switch, 還是網路重傳.

### 怎麼補

最小可驗證成果: `實作一個 single-producer single-consumer 的 lock-free ring buffer,
用 Google Benchmark 量出與 mutex 版本的延遲分布差異, 並解釋差異來源`.

- socket 層: 用 C 直接寫一個 `epoll` 的 TCP echo server, 對比 `TCP_NODELAY` 開關的延遲差異.
- 觀測: `perf stat` 看 cache miss 與 context switch, 學會用 `HdrHistogram` 看尾端延遲.
- 概念: Mechanical Sympathy 系列, LMAX Disruptor 的設計文件 (Java 寫的, 但概念與語言無關).

---

## 缺口三: FIC 交易領域知識

### 這是什麼

`FIC (Fixed Income & Currencies)` 指固定收益與外匯, 是 sell-side 的一大業務條線.
`Automated trading` 在這裡通常指`做市 (market making)` 與`演算法執行`, 而非自營策略:

| 名詞 | 意思 | 為何 front office 開發要懂 |
| --- | --- | --- |
| `Market making` | 同時掛買賣報價賺價差 | 系統的核心迴圈就是報價更新與風險再平衡 |
| `Rates algo trading` | 利率商品 (公債, IRS) 的演算法交易 | JD 的 nice-to-have 直接點名 |
| `Market connectivity` | 接交易所與 ECN 的連線層 | 亞洲市場各自有規格, JD 點名`亞洲市場特性` |
| `FIX protocol` | 金融資訊交換的產業標準協定 | 訂單, 成交回報的共同語言 |
| `Order lifecycle` | new, ack, partial fill, cancel, reject | L1/L2 支援時客戶問的每一句話都在這條線上 |
| `Trading risk` | 部位, 限額, kill switch | JD 明寫`交易面的風險意識` |

### 缺了會怎樣

領域缺口不會讓履歷被過濾, 但會讓面試的每一個情境題失去回答基礎:

- 「報價延遲 200 微秒, desk 說他們在虧錢, 你先看哪裡」— 答不出交易面的因果.
- 無法理解需求的優先順序, 而 JD 職責第一條就是`蒐集使用者需求並與 desk 主管達成一致`.
- `risk aware on both trading and IT perspective` 這條只能答一半.

### 怎麼補

最小可驗證成果: `能完整口述一筆客戶訂單從進入系統到成交回報的生命週期,
並指出每一段可能的失效模式`.

- 讀 `FIX protocol` 的 spec 摘要, 理解常見 message type (`NewOrderSingle`, `ExecutionReport`).
- 這一塊`可以在面試前兩週密集補`, 屬於名詞與流程層級, 不需要實作.
- 加分做法: 用已熟悉的 Golang 寫一個迷你 order book 與 matching engine, 兩三天可完成,
  是最有效率的`領域理解證明`.

---

## 缺口四: Windows 平台開發

### 這是什麼

sell-side 的 front office 桌面工具與部分交易元件仍跑在 Windows Server 上,
JD 將此列為 nice-to-have. 涉及 Visual Studio 工具鏈, Windows API, 與 `IOCP` 這類與 epoll 對應的非同步 I/O 模型.

### 缺了會怎樣

`影響最小的一項`. 這是 nice-to-have 的最後一條, 不會單獨造成淘汰,
只在既有 codebase 剛好是 Windows 元件時造成上手期拉長.

### 怎麼補

`不建議主動投入`. 補完前三項之後, 這一項的邊際效益最低.
若確定要投, 只需在面試時誠實說明`跨平台移植經驗可折算, Windows 特有 API 需要上手期`.

---

## 角色地圖 (Role Map)

`目前定位` 與 `這個職缺` 之間的差距不只在技能, 更在`角色類型`. 這張圖是用來判斷值不值得轉向:

| 維度 | 履歷目前的位置 | 這個職缺要的位置 | 差距性質 |
| --- | --- | --- | --- |
| 職級 | Tech Lead / Architect 方向 | `Executive (IC)`, 不帶人 | `方向相反`, 是一次職級與軌道的後退 |
| 語言 | Golang, Java, JavaScript | `C / C++` | `結構性`, 以季為單位 |
| 系統層級 | 分散式系統, 資料中心, 平台 | 單機熱路徑, 微秒級延遲 | `結構性`, 思維模式不同 |
| 領域 | 半導體, 醫療, 內容平台 | `FIC 固定收益與外匯` | `可補`, 以週為單位 |
| 工作型態 | 建設為主 | 建設 + `L1/L2 支援`, 貼著 desk | `生活型態`, 需自行判斷 |
| 薪酬 | - | `SGD 12,500 - 22,000 / 月` | `唯一的正向誘因` |

### 三條可能的路徑 (Three Paths)

`路徑 A - 不轉向 (建議)`: 這一檔的分數 `35` 落在職缺庫的 `50 以下` 區間, 意義是
`缺口為職涯類型或核心語言等級, 不建議投遞, 保留作為門檻參照`.
既有的 `80 分以上` 主投名單 (Apple, OKX) 與履歷方向一致, 應優先消化.

`路徑 B - 只補領域, 走金融的平台側`: 不碰 C++, 只補 FIC 與訂單生命週期的名詞層,
改投銀行與交易所的`基礎設施, 平台, SRE` 職缺 — 職缺庫中 Squarepoint 的
`Trading Infrastructure Support Lead` (65 分) 正是這個模式. `投入產出比最高的一條`.

`路徑 C - 補到底, 轉向低延遲工程`: 需要 `6 至 12 個月`的 C++ 與低延遲專職投入,
且轉入後要接受從 IC 重新累積年資. 只有在`確定要以交易系統為長期職涯`時才成立.

### 若選路徑 B 的執行順序

1. 補`缺口三 (FIC 領域知識)` 的名詞與訂單生命週期 — 兩週, 邊際效益最高.
2. 用 Golang 寫迷你 order book 與 matching engine — 三天, 產出可談的證據.
3. 補`缺口二`中不需要 C++ 的部分 (socket, epoll, 尾端延遲觀測) — 兩週, 對平台側職缺同樣加分.
4. `缺口一與缺口四暫緩`, 直到確定要走路徑 C.
