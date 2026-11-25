# Resume & Job Hiring Workbench

個人履歷產出與 E2E job hiring 工作台, 涵蓋公開履歷展示, 職缺分析與投遞追蹤,
跨市場就業結果資料集, 以及招募隱私法遵.

- [線上履歷頁 (index.html)](index.html)
- [Tech Resume — 科技業專用 1 頁版 (PDF)](assets/Resume-ShukLiu.pdf)
- [Full Resume — 完整經歷版 (PDF)](assets/Resume-ShukLiu-full.pdf)
- [LinkedIn](https://www.linkedin.com/in/bizshuk/)

## 業務領域 (Business Domains)

### 履歷產出 (Resume Publishing)

維護完整經歷與科技業精簡版兩條對外履歷產出線. [Resume.md](Resume.md),
[index.html](index.html) 與 [Resume-ShukLiu-full.pdf](assets/Resume-ShukLiu-full.pdf)
為完整履歷的三端同步基準; [Resume-ShukLiu.pdf](assets/Resume-ShukLiu.pdf)
為科技業專用 1 頁精簡版. 原先的 `.pages` 原稿已移至 `archive/`, 後續由
`templates/` 與 [scripts/build-pdf.sh](scripts/build-pdf.sh) 直接產出 PDF (版型與省略高中學歷等規則見 [docs/pdf_template.md](docs/pdf_template.md)).

`領域流程 (Domain Flow):`

1. 更新純文字主稿 [Resume.md](Resume.md).
2. 同步至線上履歷頁 [index.html](index.html) 與 `templates/` 內的 HTML 模板.
3. 執行 [scripts/build-pdf.sh](scripts/build-pdf.sh) 生成並驗證 `assets/` 雙版本 PDF, 由 [index.html](index.html) 提供下載.

`核心實體 (Key Entities):` `Resume.md` (內容主稿), `index.html` (線上展示頁),
`templates/` (PDF 版型模板), `Resume-ShukLiu.pdf` (科技業 1 頁版), `Resume-ShukLiu-full.pdf` (完整經歷版)

---

### 職缺分析與投遞追蹤 (Job Description Library)

蒐集目標職缺並依履歷內容評估適配度, 支撐缺口分析, 對位改寫與投遞狀態管理.

`領域流程 (Domain Flow):`

1. 在 `jd/` 按公司建檔, 記錄原文任職要求與公開薪酬.
2. 以 [Resume.md](Resume.md) 計算硬性要求覆蓋率 (`0-100`), 產出缺口分析, 角色地圖與對位敘事.
3. 在 [jd/README.md](jd/README.md) 維護排序總表與投遞追蹤狀態.

`核心實體 (Key Entities):` JD 檔案 (`jd/<company>/`), 排序與追蹤總表 (`jd/README.md`), 覆蓋註記 (`jd/_coverage_notes.md`)

---

### 跨市場就業結果資料集 (Employment Outcomes Dataset)

週期性追蹤 7 個市場 (`US`, `SG`, `TW`, `JP`, `AU`, `GB`, `EU`) 的畢業生與轉職者就業指標,
提供求職市場脈絡; 各市場口徑獨立記錄, 不做跨市場換算.

`領域流程 (Domain Flow):`

1. 依 [schedule.json](schedule.json) 的發布月份展開到期任務至 [README.todo](README.todo) 受管帳本.
2. 在 `data/<period>/` 產出同幹名 `.json` (數值) 與 `.md` (口徑定義) 配對檔並執行驗證.
3. 將稽核報告寫入 `docs/audits/` 並回報更新帳本狀態.

`核心實體 (Key Entities):` 事件排程 (`schedule.json`), 任務帳本 (`README.todo`), 資料配對 (`data/<period>/<period>__<CC>.{json,md}`)

---

### 招募隱私法遵 (Hiring Privacy Compliance)

依各國法源拆解跨國招募隱私與揭露義務 (`privacy/`), 作為收集與處理求職及就業資料的合規依據.

## 領域關聯 (Domain Relationships)

- [Resume.md](Resume.md) 同時驅動對外履歷產出 (`index.html`, 兩份 PDF) 與 `jd/` 的匹配度評分及對位敘事.
- `data/` 的就業結果資料集提供市場基準背景, 與 `jd/` 共同服務投遞決策.
- 術語單一事實來源與適用範圍見 [docs/terminology.md](docs/terminology.md) 與
  [docs/scope.md](docs/scope.md); 角色地圖與領域導覽見 `docs/tutorials/`.
