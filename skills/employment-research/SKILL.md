---
name: employment-research
description: 跨市場就業結果資料集 (Employment Outcomes Dataset) 的週期性收集, 稽核與排程. 每次執行依 schedule.json 展開到期任務, 產出本次 TODO; 未完成的任務自動延續到下一次執行並依失敗原因決定重試間隔. 觸發詞 - "跑就業資料收集", "employment outcomes", "collect graduate employment", "畢業生就業率", "本週排程", "jobs 排程", "collection plan", "GES", "NACE", "Graduate Outcomes", "就職状況調査".
---

# Employment Outcomes — 週期性收集排程

跨市場 (US, SG, TW, JP, AU, GB, EU) 的新鮮人與轉職者就業結果追蹤. 這個 skill
管的不是"做一次收集", 而是`讓收集這件事持續發生`: 官方統計會延後發布, 網址會改版,
定義會查不到 — 這些是常態不是例外, 所以工作狀態存在帳本 (ledger) 裡, 沒完成的自己
會回來.

## 核心模型 (Core Model)

```
schedule.json  ──展開──▶  README.todo 受管區塊  ──▶  資料配對
  事件排程                    待辦即帳本                data/<period>/
                                   ▲                          │
                                   └── record / 檔案對帳 ◀────┘
```

- `事件 (Event)` — schedule.json 裡一筆會週期性到期的工作, 錨定在該市場的`實際發布月份`.
- `任務 (Task)` — `<event_id>@<period>`, 例如 `sg_collect@2026-H1`. 這是帳本主鍵, 具冪等性:
  重複執行 plan 不會產生重複任務.
- `待辦即帳本 (The TODO is the Ledger)` — 狀態`只存在` [README.todo](../../README.todo)
  的受管區塊, 沒有第二份檔案. 勾選框與散文給人看, 行尾 HTML 註解給腳本讀, 兩者同時寫入.
  區塊外的手寫待辦一律原樣保留.
- `配對 (The Pair)` — 一筆紀錄是兩個同幹名檔案: `.json` 放數值, `.md` 放口徑散文.
  規範由資料集的 [README.md](../../README.md) 單一擁有.

## 執行方式 (Usage)

假設`每週執行一次`. 全部指令的 `--root` 指向資料集根目錄 (`jobs/`).

```bash
# 1. 展開本次排程, 對帳, 重寫 README.todo 的待辦區塊
python3 skills/employment-research/scripts/plan.py --root . plan

# 2. 回報單一任務結果
python3 skills/employment-research/scripts/plan.py --root . record \
    --task sg_collect@2026-H1 --status failed --reason not_yet_published \
    --note "GES 2026 尚未發布"

# 3. 查看帳本全貌
python3 skills/employment-research/scripts/plan.py --root . status

# 4. 驗證資料配對 (需 jsonschema)
.venv/bin/python3 skills/employment-research/scripts/validate.py data/2026-H1/

# 5. 展平為 tidy CSV
python3 skills/employment-research/scripts/extract.py data/ --out outcomes
```

補充旗標: `--as-of YYYY-MM-DD` 指定基準日 (回填或重演用), `--dry-run` 只印區塊不寫檔,
`--no-reconcile` 跳過檔案系統對帳, `--reopen` 重開已結案任務並歸零嘗試次數.

## 工作流程 (Workflow)

### 步驟 1 — 產生本次 TODO

跑 `plan`. 它會做四件事: 展開新到期的任務, 對帳 (檔案已存在者自動結案),
依 backoff 與相依性分類, 重寫 `README.todo` 的受管區塊.

`不要手動編輯受管區塊`. 手改會在下次執行時被覆寫, 所有狀態變更一律經 `record`.
區塊外的手寫待辦可自由編輯, 腳本不會碰.

### 步驟 2 — 逐項執行

只做報告中`本次待辦 (TODO this run)`區塊的項目. 等待重試與阻塞的不要碰.

| task 類型 | 做什麼 | 依據 |
| --- | --- | --- |
| `collect` | 收集單一 (國家, 期別) 的配對檔 | [prompts/collect.md](prompts/collect.md) |
| `verify` | 稽核他人產出的配對, 只出報告不改檔 | [prompts/verify.md](prompts/verify.md) |
| `extract` | 展平全資料集為 CSV | `scripts/extract.py` |
| `review` | 人工檢視, 無腳本 | — |

`verify` 必須在`乾淨的 context` 下執行, 且不得由寫出該檔的同一次 session 執行 —
同 context 自我複查基本上只會自我確認. 稽核報告寫到
`docs/audits/<period>__<CC>-audit.md`, 對帳器會據此自動結案.

### 步驟 3 — 驗證

`collect` 產出後一定要跑 `validate.py`, 修到通過為止. 它擋四層: 配對完整性,
JSON Schema, 跨欄位不變式, 以及`每個有值的欄位都必須在 .md 的 Metric definitions
說明口徑`. 最後一層是這個資料集的核心契約, 不可繞過.

### 步驟 4 — 回報

每一項都要 `record`, 包含成功的. 沒回報的任務下週會原樣再出現一次.
`collect` 與 `verify` 若產出檔案已落地, 下次 plan 會自動結案, 但明確回報仍較好 —
自動結案不帶失敗原因與備註.

## 失敗原因決定重試策略 (Reason → Retry Policy)

沿用 [schema/vocabularies.md](schema/vocabularies.md) 的 `gap_reason` 詞彙.
`"這次沒拿到"`與`"這個市場根本不發布"`在排程層也必須是兩件不同的事.

| `--reason` | 意義 | 計入嘗試次數 | 結果 |
| --- | --- | --- | --- |
| `not_yet_published` | 對方還沒發布 | 否 | 14 天後重試 |
| `not_found` | 可能存在, 這次沒找到 | 是 | 依 backoff 重試 |
| `definition_unclear` | 找到了但分母無法確定 | 是 | 依 backoff 重試 |
| `source_unreachable` | 網址失效或無法存取 | 是 | 依 backoff 重試 |
| `not_measured` | 該市場根本不量這個 | 否 | 終態 `skipped` |
| `discontinued` | 曾經發布, 已停止 | 否 | 終態 `skipped` |
| `needs_human` | 需要人工判斷 | 是 | 轉 `expired`, 列在需人工區塊 |

預設 backoff 為 `7 → 7 → 14 → 28` 天, 嘗試 8 次後轉 `expired`.
`不計入次數`的原因不會消耗重試額度 — 等待官方發布不該把任務逼成人工案件.

相依任務 (如 `sg_verify` 依賴 `sg_collect`) 在上游完成前停在 `blocked`, 不佔用嘗試次數;
上游若停在 `skipped`/`expired`, 下游會`串聯放棄`, 不會無限阻塞.

## 硬規則 (Hard Rules)

- `誠實的 null 勝過合理的猜測`. 沒有出版者, 發布日期, 可解析網址三者齊全的數字, 不進資料集.
- `不跨市場比較`. 七個市場的畢業生就業率不是同一個量, 本資料集不提供換算. 詳見
  [docs/scope.md](../../docs/scope.md).
- `不外插, 不內插, 不沿用上期數值`. 該期沒有新發布就設 `stale: true`.
- `不改他人期別的檔案`. 一次 collect 只寫一組配對.
- `record 過的歷史不覆寫`. 要重開已結案任務用 `--reopen`, 會留下軌跡.

## 檔案地圖 (File Map)

| 路徑 | 角色 |
| --- | --- |
| `schedule.json` | 事件排程. 新增市場或改發布月份就改這裡 |
| `README.todo` | 手寫待辦 + 受管的任務帳本區塊. 區塊由腳本維護 |
| `data/<period>/<period>__<CC>.{json,md}` | 資料配對 |
| `docs/audits/<period>__<CC>-audit.md` | 稽核報告 |
| `skills/employment-research/schema/` | 資料 schema, 排程 schema, 控制詞彙 |
| `skills/employment-research/prompts/` | collect 與 verify 提示詞 |
| `skills/employment-research/scripts/` | `plan.py`, `validate.py`, `extract.py` |

## 新增市場 (Adding a Market)

1. 在 `schema/country-period.schema.json` 的 `country` enum 加代碼.
2. 在 `schedule.json` 加一組 `<cc>_collect` 與 `<cc>_verify` 事件, `due.months`
   填該市場的`實際發布月份`, 不是希望收集的月份.
3. `verify` 事件加 `depends_on: ["<cc>_collect"]`.
4. 跑 `plan --dry-run` 確認展開的期別與到期日符合預期.
