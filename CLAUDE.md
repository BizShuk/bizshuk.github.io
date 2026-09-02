# bizshuk.github.io — 技術脈絡 (Technical Context)

## 專案結構 (Project Structure)

```tree
bizshuk.github.io/
├── index.html              # 入口頁 (gallery landing,樣式與腳本 inline)
├── README.md
├── CLAUDE.md
│
├── assets/                 # 圖片等靜態資源
│   └── images/
│       ├── personal/       # 個人相關圖示 / 校園照
│       │   ├── icon.gif
│       │   ├── edu/
│       │   │   ├── ncu/ncu.jpg
│       │   │   └── ndhu/ndhu.jpg
│       │   └── surf/uluwatu-dropin.jpg   # surf 頁 hero (本人 Uluwatu 照)
│       ├── photos/         # 歷史生活照片 (未在現有頁面中使用)
│       └── profile/shuk-profile.jpg
│
├── data/
│   ├── gallery.json        # gallery tile 資料來源
│   └── params.json         # legacy GitHub Pages metadata (未使用)
│
└── pkg/                    # 各自封裝的子頁面 (subpage packages)
    ├── resume/
    │   ├── index.html      # 履歷頁 (sunny 風格,樣式與腳本 inline)
    │   ├── Resume.md       # 履歷內容主稿
    │   └── assets/         # 可下載的履歷 PDF 與作品截圖
    └── surf/
        └── index.html      # 衝浪頁 (ocean editorial, Uluwatu hero,樣式與腳本 inline)
```

`.gitignore` 內容：

- `.DS_Store`
- `*.*.sw[op]` — vim swap files
- `life/` — 個人生活筆記子模組 (內容不在 deploy 範圍)

## 技術棧 (Tech Stack)

- Language: HTML5 + vanilla JavaScript (ES2017, `fetch` API)
- Framework: 無 (no React, no jQuery, no Vue — 自 2026-06 重構移除)
- Build tool: 無 (靜態檔案直接 deploy)
- Font: Google Fonts — `Montserrat` 300/400/600(/700)
- Image hosting: `i.imgur.com` (外部 hot-link)、`via.placeholder.com` (教育卡片 fallback)
- 部署: GitHub Pages (假設 `master` 分支 / `gh-pages` 流程;未偵測到 workflow 檔)

關鍵相依：

- 瀏覽器原生 `fetch` (gallery.js)
- Google Fonts CDN
- 第三方圖床 (imgur.com)

## 關鍵決策 (Key Decisions)

- Decision 1: 不使用任何前端框架。`index.html` 透過原生 `fetch` 從
  `data/gallery.json` 拉資料,以保持頁面輕量、避免 bundler 依賴。
- Decision 2: 將頁面邏輯切到外部 `js/` 與樣式切到 `css/` 進行分離,讓
  `index.html` 只保留 semantic markup;`pkg/` 下的子頁面則自成一包,樣式與腳本隨頁面放在同一資料夾。
- Decision 3: 圖片以外部 CDN (imgur) 為主,減少 repo 體積,但代價是依賴
  第三方 hot-link 政策 (見 README 改善建議)。
- Decision 4: 履歷頁採用純手寫 HTML 而非模板引擎,內容以 `pkg/resume/Resume.md`
  為主稿手動同步,換得零工具鏈成本。
- Decision 5: 在兩個 HTML 頁面都使用 `<meta name="referrer" content="no-referrer">`,
  避免將 referrer 洩漏給第三方 (imgur、外部社交連結)。
- Decision 6: 教育卡片使用 inline `onerror="this.src='https://via.placeholder.com/...'"` 處理圖片失效情境,確保卡片 layout 不會破。
- Decision 7: gallery 渲染時若 `link` 為空字串則降級為純 `<div>`,讓
  decorative tile 與有連結的 tile 共用同一份資料 schema。

## 模組對應 (Module Mapping)

| 業務領域 (Domain)               | 套件/模組 (Package/Module)                                          | 進入點 (Entry Point)                           |
| ------------------------------- | ------------------------------------------------------------------- | ---------------------------------------------- |
| 個人作品集 (Personal Portfolio) | `index.html` (樣式與腳本 inline), `data/gallery.json`            | `index.html` inline `<script>` 的 `DOMContentLoaded` |
| 履歷展示 (Resume Showcase)      | `pkg/resume/index.html` (樣式與腳本 inline)                         | 頁尾 inline `<script>` (年份 + 捲動顯影)       |
| 衝浪頁 (Surf Page)              | `pkg/surf/index.html` (樣式與腳本 inline)                           | `pkg/surf/index.html` inline `<script>` 的 `DOMContentLoaded` |

## 開發指南 (Development Guide)

### 前置需求 (Prerequisites)

- 一個現代瀏覽器 (支援 `fetch` 與 CSS `scroll-snap`)
- (選擇性) 本機 HTTP server,例如 `python3 -m http.server`,因為 `fetch`
  對 `file://` protocol 在多數瀏覽器會被 CORS 擋下

### 安裝 (Installation)

無 `package.json`,直接 `git clone` 後即可使用。

### 建置 (Build)

無 — 所有檔案即 production artifact。

### 測試 (Test)

未偵測到測試設定 (No test config detected)。建議手動驗證:

1. 開啟 `index.html` 並確認 gallery 16 個 tile 都正確載入
2. 將 `data/gallery.json` 暫時改成非法 JSON,確認 console 有錯誤訊息
3. 開啟 `pkg/resume/index.html` 確認 PDF 下載連結,以及 `pkg/surf/index.html` 的捲動顯影

### 部署 (Deploy)

未偵測到部署設定 (No deployment config detected)。預設推送到 `master`
分支後,GitHub Pages 會自動透過預設路徑提供服務。
若使用 `gh-pages` 分支,流程同 `params.json` 內描述的 boilerplate。

## 慣例 (Conventions)

- Naming: kebab-case 用於檔名 (`gallery.json`, `uluwatu-dropin.jpg`);camelCase
  用於 CSS class 中的複合名 (`gallery-item`, `gallery-item-overlay`)。
- CSS variables: 兩個 stylesheet 都以 `:root` 宣告 design tokens
  (`--primary-color`, `--bg-color`, ...)。各頁面分別維護自己的 token 集合,
  沒有抽出 shared token 檔。
- Error handling: JS 端僅在 `gallery.js` 的 `fetch` chain 結尾有
  `console.error`,沒有 UI 層級的 fallback;HTML 端用 inline `onerror`
  處理圖片失效。
- Logging: 無,僅瀏覽器 console。
- Testing: 無 (見上方 測試 章節)。
- Image policy: 優先使用外部 `i.imgur.com` URL;本地 `assets/images/`
  僅用於教育卡片這類需要穩定控制的場景。
- Referrer policy: 兩個 HTML 頁面都明確設定
  `<meta name="referrer" content="no-referrer">`。
- 註解風格: JS 檔頭以 `/* path - purpose */` 形式說明歸屬;
  HTML 內部以 `<!-- 註解 -->` 標示 section 邊界。
