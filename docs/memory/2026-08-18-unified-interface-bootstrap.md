# 2026-08-18 — 統一介面補齊 (Unified Interface Bootstrap)

## 背景 (Context)

`~/projects/` 全域稽核發現本 repo 缺三件統一介面必備檔案。本次只補缺件,
既有 `README.md` / `CLAUDE.md` / `AGENTS.md` symlink / `plans/` 一律未動。

## 補齊項目 (What Was Added)

- `README.todo` — 待辦事項。來源是 `README.md` 既有的`改善建議`七項、
  `plans/` 內未勾選的 step,以及本次閱讀時發現的文件分岔。
- `docs/terminology.md` — 術語表。分五個領域,每筆都指到真實檔案或識別符。
- `docs/memory/` — 本檔即目錄的第一份留底。

## 值得留底的現況 (Findings Worth Keeping)

- `README.md` 與 `CLAUDE.md` 都宣稱「無 `package.json`」「無 build step」
  「未偵測到部署設定」。這三句`現在都是假的`: repo 已有 `package.json`
  (`dev` / `build` / `deploy` / `clean` / `destroy` / `release:store`)、多階段
  `Dockerfile` 與 `docker/nginx.conf`。文件停在純靜態站的年代,實作已經走到容器化。
  已列入 `README.todo`,本次刻意不改既有文件以免和其他修復撞在一起。
- `pkg/` 是`第三個業務領域`,兩份正典文件都沒提。它裝的是六個 iOS app 的
  App Store 政策頁,而且`不是在本 repo 編輯的` — 真正的來源在
  `iphone/<project>/appstore/`,由 `scripts/release-store.sh` 以
  `rsync -a --delete` 單向覆蓋過來。這代表`直接改 pkg/ 底下的檔案會在下次發佈時被無聲蓋掉`,
  是這個 repo 最容易踩的坑。
- 同一個腳本會先跑外部的 `POLICY_VALIDATOR`(住在 `iphone/` 的 skill 裡),
  驗證器不存在就整個中止。也就是`本 repo 的發佈流程對外部 repo 有硬相依`,
  但兩邊都沒有文件寫下這條契約。
- `surf.html` 是三頁裡唯一做 progressive enhancement 的: `js/surf.js` 先在
  `documentElement` 掛 `js` class,並在 `prefers-reduced-motion` 或缺
  `IntersectionObserver` 時直接把所有 `.reveal` 設為可見。另外兩頁沒有等價的降級路徑。
- `data/params.json` 從 GitHub Pages 產生器留到現在,沒有任何 handler 讀它。
  README 已把它列為改善建議,至今未處理。
