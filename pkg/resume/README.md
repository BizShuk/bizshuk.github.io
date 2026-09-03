# Resume

履歷本體 (resume artifacts): 可編輯的 Pages 原稿、匯出的 PDF、履歷與作品集引用的圖片,
以及支撐履歷條目的成就清單。面試與生涯規劃資料不放這裡,見 `BizShuk/resume` repo。

- [線上履歷頁 (index.html)](index.html)
- [Engineer Resume (PDF)](assets/Resume-ShukLiu.pdf)
- [Engineer Resume — full (PDF)](assets/Resume-ShukLiu-full.pdf)
- [LinkedIn](https://www.linkedin.com/in/initialshuk/)

## 內容 (Contents)

| 路徑                       | 說明                                                       |
| -------------------------- | ---------------------------------------------------------- |
| `index.html`               | 線上履歷頁 (sunny 風格),內容同步自 `Resume.md`             |
| `Resume.md`                | 履歷純文字主稿,線上頁與 Pages 的內容來源                    |
| `Resume-ShukLiu.pages`     | 履歷原稿 (Pages),精簡版,PDF 由此匯出                     |
| `Resume-ShukLiu-full.pages`| 履歷原稿 (Pages),完整版                                    |
| `assets/`                  | 匯出的 PDF,以及履歷與作品集截圖 (專案畫面、公司/團隊圖表)   |
| `achievement.todo`         | 各任職公司的成就與專案清單,履歷條目的原始素材              |
| `jd/`                      | 職缺庫 (job description library),含匹配度評分              |
| `svc/`, `cmd/`, `main.go`  | `resume` CLI:從 MyCareersFuture 蒐集職缺資料到 `jd/` 的原始素材 |

## 職缺蒐集 CLI (Job Collection CLI)

`resume` 是一支 Go CLI,從 `api.mycareersfuture.gov.sg` 蒐集職缺,供 `jd/` 的職缺庫使用。
該 API `不需要 API key`,實測過的端點、欄位與陷阱記在 [svc/README.md](svc/README.md)。

```bash
go build -o bin/resume .

bin/resume mcf search "engineering manager" --salary 12000 --level Manager
bin/resume mcf detail <uuid>                       # 唯一拿得到 JD 全文的路徑
bin/resume mcf fetch "engineering manager" --max 50 # 搜尋 + 逐筆補全文 + 落檔
bin/resume mcf jobs --salary 15000                 # 全站列表,不吃關鍵字
bin/resume mcf company capgemini                   # 雇主查詢,解析出 UEN
```

設定走 gosdk 的扁平大寫 key,可由同名環境變數覆寫:
`MCF_BASE_URL`, `MCF_USER_AGENT`, `MCF_TIMEOUT`, `MCF_MAX_ATTEMPTS`, `COMMAND_TIMEOUT`。
落檔位置預設為 `~/.config/resume-jd/data/mcf/jobs/`,可用 `--out` 覆寫。

## 關聯 (Related)

- `index.html` — 線上履歷頁 (原 `../../resume.html` 已移除),內容應與本資料夾的 PDF 保持一致
