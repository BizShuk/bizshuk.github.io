---
name: coverletter
description: Use when a cover letter / motivation letter PDF is needed for a specific job, written from pkg/resume/Resume.md and tailored to a job description - triggers on "cover letter", "coverletter", "write a cover letter for this job", "motivation letter", "求職信", "自薦信", "幫我寫 cover letter", and on any write of cover letters in pkg/resume/jd/. Requires a JD (pasted text, URL, or a pkg/resume/jd/ file); stops when none is given.
---

# Cover Letter

## Overview

以 [Resume.md](../../pkg/resume/Resume.md) 為唯一事實來源, 針對`一份指定的 JD` 產出`單頁 (1 page)` 的英文求職信 PDF.
PDF 產製遵循 `pdf` skill (`/pdf`) 的 `reportlab` 建檔與 `pypdf` 驗證做法, 由 [scripts/render_pdf.py](scripts/render_pdf.py) 封裝.

`核心原則`: 求職信是`JD 要求 -> 履歷證據`的對位, 不是履歷的散文版. 每一段都要能指回 JD 的某條要求與 Resume.md 的某行事實.

`不適用`: JD 建檔與匹配度評分 (屬 `job-analyze`), 履歷改寫本身 (屬 `Resume.md`).

## Output Location (與 JD 路徑對齊)

求職信是 JD 檔的`伴隨檔 (companion)`, 與 JD `同目錄, 同檔名主幹 (stem)`, 只換副檔名:

| 檔案 | 路徑 |
| --- | --- |
| JD | `pkg/resume/jd/<company>/<title_slug>.<team_or_domain>.md` |
| 原稿 | `pkg/resume/jd/<company>/<title_slug>.<team_or_domain>.coverletter.md` |
| PDF | `pkg/resume/jd/<company>/<title_slug>.<team_or_domain>.coverletter.pdf` |

- `先有 JD 檔才有路徑`: JD 以全文或 URL 提供且尚未入庫時, 先以 `job-analyze` 的 `只建檔評分` 模式建檔, 再以該檔路徑推導輸出路徑. `不得`自行編造 `<company>/<stem>`, 也`不得`寫到 `pkg/resume/coverletter/` 等其他目錄.
- `已封存的 JD`: JD 位於 `jd/archive/<company>/` 時, 求職信也寫在同一個 archive 目錄.
- `前置檢查`: 原稿的 front matter `jd:` 必須等於 JD 檔的 repo 相對路徑, 且 `test -f <jd>` 成立.
- `封存連動`: `job-analyze` 的封存腳本會把 `.coverletter.md` / `.coverletter.pdf` 隨 JD 一起移到 `archive/`, 不需另外處理.

## JD Gate (先做, 不可跳過)

在讀履歷或寫任何字`之前`, 確認使用者這次`明確提供`了 JD, 形式限於:

| 形式 | 判定 |
| --- | --- |
| 貼上的 JD 全文或主要要求條目 | 通過 |
| 職缺 URL, 且能抓到全文 | 通過; 抓取被擋則視同未提供 |
| [pkg/resume/jd/](../../pkg/resume/jd/) 下的檔案路徑, 或能唯一對應到單一檔案的 `公司 + 職稱` | 通過 |

全文或 URL 通過 Gate 後, 仍須先入庫取得 JD 檔路徑, 見 `Output Location`.

不通過時`立即停止`, 不建目錄, 不寫檔, 不產 PDF, 只回覆:

```text
JD 未提供, 已停止產出 cover letter.
請提供下列任一: JD 全文, 職缺 URL, 或 pkg/resume/jd/ 下的檔案路徑.
```

`禁止`的替代做法: 寫一封通用 (generic) 求職信, 從職缺庫自行挑一份 JD, 用搜尋摘要或公司官網介紹拼湊要求. `公司 + 職稱`對應到多個檔案時, 列出候選並停止, 讓使用者選.

## Non-Negotiables

- `事實只來自 Resume.md`: 年資, 職稱, 公司, 數字, 技術棧都必須能在 Resume.md 找到原句. 沒有的經驗`不得宣稱`, 也不得把 `熟悉` 升級成 `精通`. 這是最常見的失敗: 為了貼合 JD 而把 JD 用詞寫進自己的經歷.
- `JD 缺口不硬湊`: JD 的硬性要求履歷沒有覆蓋時, 不提或以可轉移能力 (transferable) 誠實帶過, 不虛構.
- `單頁`: 本文約 `250-350` 英文字, 3-4 段. 渲染腳本會在 `10.5 / 10 / 9.5 pt` 間自動縮字, 仍超過一頁則退出碼 `2`, 必須刪減內容重跑, 不得調小邊界或字級繞過.
- `語言`: 預設英文; JD 為其他語言時才跟隨 JD 語言 (中文 JD 需另行處理字型, 先告知使用者).
- `聯絡資訊不手寫`: 姓名, email, 電話, LinkedIn 由腳本從 Resume.md 開頭讀取, 信件原稿不重複.

## Letter Structure

| 段落 | 內容 | 長度 |
| --- | --- | --- |
| `開場 (Hook)` | 應徵的職位, 為什麼是這家公司的這個團隊 (取自 JD 原文的使命或產品), 一句定位 | 2-3 句 |
| `證據一` | JD 最重的一條要求 -> Resume.md 最直接的一段經歷, 帶一個具體結果 | 3-4 句 |
| `證據二` | JD 第二重的要求 (常為領導, 跨團隊, 規模) -> 另一段經歷 | 3-4 句 |
| `收尾 (Close)` | 能帶來的價值, 地點與身分 (Singapore PR, 若 JD 相關), 邀約面談 | 2-3 句 |

## Workflow

1. `JD Gate`: 見上. 未通過即停止.
2. `定位 JD 檔`: 取得 JD 檔的 repo 相對路徑 `<jd>`; 尚未入庫則先跑 `job-analyze` 的 `只建檔評分`. 輸出路徑一律由 `<jd>` 推導, 見 `Output Location`.
3. `讀 JD`: 萃取 3-5 條硬性要求與團隊使命, 一併讀 JD 檔的 `缺口分析` 與 `對位敘事`, 避免寫到缺口上.
4. `讀 Resume.md`: 為每條要求找對應原句, 找不到的標為缺口.
5. `寫原稿`: 寫到 `<jd 去掉 .md>.coverletter.md`, 已存在則就地覆蓋. 格式:

   ```markdown
   ---
   company: Stripe
   role: Staff Software Engineer, Payments
   recipient: Hiring Team
   date: 2026-09-24
   jd: pkg/resume/jd/stripe/staff_software_engineer.payments.md
   ---

   Dear Hiring Team,

   <paragraph 1>

   <paragraph 2>

   Sincerely,
   Shuk Liu
   ```

   `company`, `role`, `date`, `jd` 必填, `jd` 為 JD 檔的 repo 相對路徑; `recipient` 僅在 JD 寫明聯絡人時填姓名, 否則 `Hiring Team`. 段落之間空一行, 段內不換行.
6. `渲染 PDF` (在 repo 根目錄執行):

   ```bash
   uv run --with reportlab --with pypdf python skills/coverletter/scripts/render_pdf.py \
     pkg/resume/jd/<company>/<title_slug>.<team_or_domain>.coverletter.md
   ```

   輸出同目錄的 `.coverletter.pdf`, 不要用 `--out` 改到別處. 退出碼 `2` 表示超過一頁, 回步驟 5 刪減. 這台機器的 `python3` 沒有 `reportlab`, 一律經 `uv run --with` 帶入.
7. `驗證`: 以 `pypdf` 抽出 PDF 文字, 逐句確認沒有 Resume.md 以外的事實, 公司名與職稱與 JD 一致.
8. `回覆使用者`: 一行對位摘要 (哪兩條 JD 要求對到哪兩段經歷), 未覆蓋的 JD 硬性要求, 以及原稿與 PDF 的`絕對路徑`.

## Common Mistakes

| 錯誤 | 修正 |
| --- | --- |
| 沒有 JD 仍寫一封通用信 | JD Gate 停止並回覆固定訊息 |
| 把 JD 的技術關鍵字塞進自己的經歷 | 只用 Resume.md 原句, 缺口不提 |
| 重述整份履歷時間線 | 只挑兩段最對位的經歷 |
| 超過一頁就縮邊界 | 刪減內容, 邊界與字級下限固定 |
| 寫到 `pkg/resume/coverletter/` 或自訂檔名 (`.cover.md`) | 一律 `<jd stem>.coverletter.md`, 與 JD 同目錄 |
| JD 未入庫就自行編造 `<company>/<stem>` | 先以 `job-analyze` 建檔, 再由 JD 檔路徑推導 |
