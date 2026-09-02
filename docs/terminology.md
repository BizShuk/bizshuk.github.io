# bizshuk.github.io — 術語表 (Terminology)

本檔是本站術語的`單一定義來源 (single source of truth)`。`README.md`、`CLAUDE.md`、
程式碼註解與 commit message 一律沿用此處用詞,同一概念不得有第二種說法。

## 個人作品集 (Personal Portfolio)

| 術語 (Term) | 英文 (English) | 定義 (Definition) | 出處 (Source) |
| ----------- | -------------- | ----------------- | ------------- |
| 作品格 | Gallery Item | 首頁格狀清單的一格。schema 為 `{ link, img_title, style.backgroundImage }`,由 JS 於執行期 fetch 後渲染 | `data/gallery.json`, `js/gallery.js` |
| 格標題 | `img_title` | 作品格的顯示文字,渲染成 hover 時浮現的 overlay 標題 | `data/gallery.json`, `js/gallery.js` (`gallery-item-title`) |
| 裝飾格 | Decorative Tile | `link` 為空字串或空白的作品格。渲染成 `<div>` 而非 `<a>`,外觀相同但不可點擊 | `js/gallery.js` (`isLink` 判斷式) |
| 疊層 | Overlay | 覆蓋在作品格底圖上的標題層,class 為 `gallery-item-overlay` | `js/gallery.js`, `css/index.css` |
| 錯位進場 | Staggered Animation | 每格依索引遞增 `animationDelay`(`index * 0.05s`)造成的依序淡入 | `js/gallery.js` |

## 履歷展示 (Resume Showcase)

| 術語 (Term) | 英文 (English) | 定義 (Definition) | 出處 (Source) |
| ----------- | -------------- | ----------------- | ------------- |
| 經歷項 | Timeline Item | 工作經歷區塊的單筆條目(職稱、公司與地點、期間、條列)。內容 inline 手寫,無資料層 | `pkg/resume/index.html` `.job` |
| 技能標籤 | Skill Tag | 技能分類下的藥丸狀 chip | `pkg/resume/index.html` `.tag` |

## 衝浪頁 (Surf Editorial)

| 術語 (Term) | 英文 (English) | 定義 (Definition) | 出處 (Source) |
| ----------- | -------------- | ----------------- | ------------- |
| 章 | Chapter | 長卷軸雜誌版面的一個編號段落,含 `chapter-no` / `chapter-label` / `chapter-rule` | `pkg/surf/index.html` `.chapter` |
| 足跡 | Footprints | 衝浪地點列表區塊(Bali / Lombok / Taiwan) | `pkg/surf/index.html` `#footprints` |
| 板櫃 | Quiver | 衝浪板收藏區塊,每張板有 `board-brand` / `board-model` / `board-role` / `board-note` | `pkg/surf/index.html` `#quiver` |
| 浪點 | Spot | 足跡區塊內的單一衝浪點,含 `spot-name` / `spot-tag` / `spot-note` | `pkg/surf/index.html` `.spot` |
| 捲動顯影 | Reveal | 元素進入視窗時由 `IntersectionObserver` 加上 `is-visible` 的漸入效果。threshold 為 `0.15` | `pkg/surf/surf.js`, `pkg/surf/index.html` `.reveal` |
| JS 旗標 | `js` class | `pkg/surf/surf.js` 在 `documentElement` 掛上的 class。作為 progressive enhancement 開關 — 沒有 JS 時內容維持全可見 | `pkg/surf/surf.js`, `pkg/surf/surf.css` |

## App Store 政策頁 (Store Policy Pages)

| 術語 (Term) | 英文 (English) | 定義 (Definition) | 出處 (Source) |
| ----------- | -------------- | ----------------- | ------------- |
| 政策發佈 | Store Policy Release | 把 iPhone 專案的 `appstore/` 目錄同步進本站 `pkg/<route>/` 的流程,以 `rsync -a --delete` 執行 | `scripts/release-store.sh`, `package.json` (`release:store`) |
| 路由名 | Route | `APPS` 陣列中 `project:route` 的右半,決定政策頁在站上的路徑 `pkg/<route>/` | `scripts/release-store.sh` |
| 政策驗證器 | Policy Validator | 發佈前逐專案執行的外部 Python 檢查腳本,路徑由 `POLICY_VALIDATOR` 覆寫。找不到即中止發佈 | `scripts/release-store.sh` |
| 來源專案根 | `IPHONE_PROJECTS_ROOT` | iPhone 專案所在目錄,預設為 repo 上兩層的 `iphone/`。政策頁的原始內容住在這裡,不在本 repo 編輯 | `scripts/release-store.sh` |
| 政策頁組 | Policy Page Set | 每個 app 固定四件套 + 預覽圖:`index.html`、`privacy.html`、`copyright.html`、`privacy.md`、`preview/` | `pkg/<app>/` |

## 部署 (Deployment)

| 術語 (Term) | 英文 (English) | 定義 (Definition) | 出處 (Source) |
| ----------- | -------------- | ----------------- | ------------- |
| 打包階段 | `image-build` stage | Dockerfile 第一階段。把工作樹複製成 `/home/app/out/html`,並剔除 `.git` / `Dockerfile` / `docker/` 與 `*.map` | `Dockerfile` |
| 產物目錄 | `/home/app/out` | runtime 階段唯一帶入的目錄,內含 `html/`、`nginx.conf`、`entrypoint.sh` | `Dockerfile` |
| 入口符號連結 | `/home/app/app` | 指向 `out/entrypoint.sh` 的 symlink,是容器 `ENTRYPOINT` | `Dockerfile` |
| 設定根 | `/home/app/.config/bizshuk.github.io/` | 容器內的 per-app 設定根,`data/` 放 nginx pid、`logs/` 保留 | `Dockerfile`, `docker/nginx.conf` |

## 狀態值 (Status Values)

未偵測到列舉型狀態值 (Not detected) — 本站無執行期狀態機,唯一的分支是
`js/gallery.js` 以 `link` 是否為非空字串決定渲染成 `<a>` 或 `<div>`。

## 待補 (To be extended)

- `css/index.css`、`pkg/surf/surf.css` 與 `pkg/resume/index.html` 的 inline 樣式各自維護一組 `:root` design token,
  尚未整理出共用命名。統一 token 之後再把 token 名收進本表。
