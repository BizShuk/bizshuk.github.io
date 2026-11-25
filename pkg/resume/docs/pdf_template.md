# Resume PDF Template Rule

本文件定義 `pkg/resume/` 兩份對外履歷 PDF 的版型規範與產出契約, 確保每次重建 PDF 都可重現且不漂移.

## 模板與產出對應 (Templates & Outputs)

| 模板 (Template) | 產出 (Output PDF) | 用途 | 頁數上限 |
| --- | --- | --- | --- |
| `templates/resume-tech.html` | `assets/Resume-ShukLiu.pdf` | 科技業投遞專用精簡版 | 嚴格 `1` 頁 |
| `templates/resume-full.html` | `assets/Resume-ShukLiu-full.pdf` | 完整經歷版 (含 `Other Experience`) | `2` 頁 |

## 硬性規範 (Non-Negotiables)

1. **事實單一來源**: 所有經歷, 日期, 聯絡資訊與技術堆疊必須來自 [Resume.md](../Resume.md), 不得在模板內編造未見於主稿的內容. 聯絡列固定使用 `biz.shuk@gmail.com`, `https://www.linkedin.com/in/bizshuk/` 與 `WhatsApp (+65 8012 9716)`, 並標註與 `Resume.md` 一致的 `Last updated`.
2. **禁止高中學歷 (No High School in PDF)**: 所有 PDF 版本只列高等教育學歷 (`M.S.` 與 `B.S.`), **不得**包含高中學歷 (`Taipei Municipal Heping High School` 或 `High School`), 即使 `Resume.md` 與 `index.html` 保留該項.
3. **版面與頁數紀律**:
   - 統一使用 `A4` 紙張 (`@page { size: A4; }`), 系統無襯線字體, 雙色層級 (`#111827` 主色 + `#1d6fa5` 連結), 不得使用點陣截圖取代可選取文字.
   - 每個經歷區塊 (`article.entry`) 必須設定 `break-inside: avoid`, 標題 (`h2`) 設定 `break-after: avoid`, 避免標題或公司列跨頁斷裂.
   - `resume-tech.html` 超出 1 頁或 `resume-full.html` 超出 2 頁時, 必須精簡條列文字, 不得以極端縮小邊界 (< 9mm) 或極小字級 (< 8pt) 硬塞.
4. **雙版本內容邊界**:
   - `Tech PDF` 省略 `Other Experience`, 段落順序為 `Summary -> Skills -> Work Experience -> Education`.
   - `Full PDF` 保留完整條列與 `Other Experience`, 段落順序與 [Resume.md](../Resume.md) 一致 (`Summary -> Education -> Skills -> Work Experience -> Other Experience`).
5. **歸檔與生成方式**: 原 `.pages` 封裝檔封存於 `archive/`, 不再從 Pages 匯出; 一律透過 [scripts/build-pdf.sh](../scripts/build-pdf.sh) 產出與驗證.

## 重建與驗證 (Build & Verify)

於 repo 根目錄執行:

```bash
npm run build:pdf
```

(或直接執行 `./pkg/resume/scripts/build-pdf.sh`) 腳本會以 headless Chrome 列印兩份 PDF, 並自動驗證:
- `Resume-ShukLiu.pdf` 為 `1` 頁, `Resume-ShukLiu-full.pdf` 為 `2` 頁.
- 兩份 PDF 抽出文字後皆不含 `High School` 與 `Heping`.
