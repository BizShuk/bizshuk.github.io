# Surf Page Design Spec — 海洋編輯風 (Ocean Editorial)

以使用者本人於 Uluwatu 衝浪的真實照片 (`0U4A1604.jpg`) 為主視覺，
新增獨立頁面 `surf.html`，走雜誌長卷軸 (magazine long-scroll) 編輯風格。

## 定位

- 獨立 surf 頁，與 `resume.html` 同層；不動現有 `index.html` / `resume.html`
- 從 gallery 連入（`data/gallery.json` 新增一個 tile）
- 頁面本身不設回主站導航，footer 僅極小 © 標記

## 檔案結構

```tree
bizshuk.github.io/
├── surf.html                                  # 新增
├── css/surf.css                               # 新增
├── js/surf.js                                 # 新增（IntersectionObserver reveal）
├── assets/images/personal/surf/
│   └── uluwatu-dropin.jpg                     # 1983x1281 原檔壓縮（~1800px, q80）
└── data/gallery.json                          # 加 tile 連入 surf.html
```

## 視覺系統

全部取自照片本身：

- Palette：深海藍 `#0b3a63`、浪面藍 `#1f6fae`、浪花白 `#eef6fb`、
  彩虹板 accent（coral→gold 漸層，僅用於強調線與 hover）
- 字體：Montserrat（repo 慣例）；大標 300 weight、uppercase、寬字距
- 雜誌元素：章節編號 `01 / 02 / 03`、細線分隔、攝影圖說（`Uluwatu, Bali`）

## 內容區塊（由上而下）

1. Hero — 全幅照片，底部漸層壓字 `SHUK · SURF — Uluwatu, Bali`，scroll cue
2. `01 The Story` — 衝浪者視角短文：這張 drop-in 的故事（英文編輯文，配合全站語言）
3. `02 Footprints` — 浪點足跡，依地區分組：
    - `BALI`：Uluwatu（本照）等
    - `LOMBOK`：Gerupuk（Inside/Outside/Don Don）、Tanjung Aan、Ekas（Inside/Baby）、
      Mawi、Seger、Awang — 一句話注記，自 life 筆記提煉、去私人化
    - `TAIWAN`：先以常見浪點隨機產生（中角灣、外澳、佳樂水、都蘭），待使用者修正
4. `03 Quiver` — `Firewire Cymatic`、`Pyzel Shadow`
5. Footer — 極小 © 標記

## 技術決策

- 內容純手寫 inline HTML（同 `resume.html` Decision 4），不走 fetch/JSON
- 照片放本地 `assets/`（穩定控制場景，符合 image policy），以 `sips` 壓縮
- reveal 動畫用 `IntersectionObserver`，`prefers-reduced-motion` 時停用
- `<meta name="referrer" content="no-referrer">`（慣例）
- 無框架、無 build step（repo Decision 1）

## 驗證

1. 本機 `python3 -m http.server` 開 `surf.html`，桌機與行動視窗各截圖確認
2. gallery tile 連結可達 `surf.html`
3. `prefers-reduced-motion` 下無動畫、內容仍完整可見
