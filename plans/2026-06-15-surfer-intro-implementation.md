# Surfer Intro Page Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

`Goal:` Create a full-screen tech HUD-themed surfer introduction page inside `sample/intro.html` to showcase surfer stats and live telemetry.

`Architecture:` Implements a responsive three-column grid layout inside `sample/intro.html`. The left column hosts bio and heartbeat SVG biometrics, the center column shows the surfer visual with HUD rotating grids, and the right column draws SVG metrics charts. A JavaScript loop runs heart rate simulation and keypress triggers.

`Tech Stack:` HTML5, Vanilla CSS (CSS Grid, Flexbox, Keyframe Animations, SVG), Vanilla JavaScript (ES6).

---

### Task 1: 建立衝浪者主視覺圖片 (Generate Hero Surfer Image)

`Files:`
- Create: `assets/images/personal/surfer_action.jpg`

- [ ] `Step 1: 產生衝浪者的高品質動作照片`
  呼叫 `generate_image` 產生一張大浪衝浪的動態照片，為科技風格與雙色調過濾器做準備。
  Prompt: `A high-contrast professional action photo of a male surfer carving on a massive dramatic deep blue ocean wave, sun spray, premium sports photography style, centered`
  Image Name: `surfer_action`
  儲存路徑為 `/Users/shuk/projects/bizshuk.github.io/assets/images/personal/surfer_action.jpg` (以生成工具產生的檔案為準，若是 PNG 則為 `.png`)。

---

### Task 2: 更新 `sample/intro.html` 基礎結構與設計標記 (Update HTML & Base Styles)

`Files:`
- Modify: `sample/intro.html`

- [ ] `Step 1: 修改 sample/intro.html 導入全新的 HUD 變數與 Reset`
  修改 `<style>` 中的變數，定義科技感調色盤（深海暗藍、亮天藍、霓虹紅、螢光黃）。
  替換 `sample/intro.html` 中的整個內容為新的選手介紹結構。

```html
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="referrer" content="no-referrer">
<title>COMPETITOR PROFILE // KOA MERCER</title>
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=Syne:wght@400;600;700;800&family=DM+Mono:wght@300;400;500&family=Plus+Jakarta+Sans:wght@300;400;500;600&display=swap" rel="stylesheet">
<style>
:root {
  --bg: #040810;
  --bg-gradient: radial-gradient(circle at center, #0c1424 0%, #040810 100%);
  --surface: rgba(12, 20, 36, 0.6);
  --border: rgba(0, 240, 255, 0.15);
  --border-hi: rgba(0, 240, 255, 0.3);
  --text: #e0f7fc;
  --text-dim: #8ba7b0;
  --text-dark: #3a535c;
  --accent: #00f0ff;          /* tech cyan */
  --accent-glow: rgba(0, 240, 255, 0.35);
  --warning: #ff0055;         /* heart beat / bio warning */
  --warning-glow: rgba(255, 0, 85, 0.35);
  --marker: #ffd24a;          /* high-voltage yellow */
  --f-head: 'Syne', sans-serif;
  --f-body: 'Plus Jakarta Sans', sans-serif;
  --f-mono: 'DM Mono', monospace;
  --t: 1;
}

*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
html, body { height: 100%; overflow: hidden; background: var(--bg); background-image: var(--bg-gradient); color: var(--text); font-family: var(--f-body); -webkit-font-smoothing: antialiased; }

/* HUD border frames */
.stage { position: relative; width: 100vw; height: 100vh; display: grid; place-items: center; padding: 1.5vh 2vw; }
.frame { position: relative; width: 100%; height: 100%; display: grid; grid-template-rows: auto 1fr auto; padding: 20px; border: 1px solid var(--border); background: rgba(4, 8, 16, 0.4); backdrop-filter: blur(10px); }
.frame::before { content: ''; position: absolute; inset: 0; border: 1px solid var(--border); pointer-events: none; margin: -5px; }

/* Grid corners */
.corner { position: absolute; width: 15px; height: 15px; border: 2px solid var(--accent); }
.corner-tl { top: -3px; left: -3px; border-right: none; border-bottom: none; }
.corner-tr { top: -3px; right: -3px; border-left: none; border-bottom: none; }
.corner-bl { bottom: -3px; left: -3px; border-right: none; border-top: none; }
.corner-br { bottom: -3px; right: -3px; border-left: none; border-top: none; }

/* Header & Telemetry */
header { display: flex; justify-content: space-between; font-family: var(--f-mono); font-size: 11px; letter-spacing: 0.15em; color: var(--text-dim); border-bottom: 1px solid var(--border); padding-bottom: 10px; text-transform: uppercase; }
header span b { color: var(--accent); }

/* Main Grid Layout */
.main-grid { display: grid; grid-template-columns: 280px 1fr 300px; gap: 20px; height: 100%; min-height: 0; margin-top: 15px; margin-bottom: 15px; }

/* Panels styling */
.panel { background: var(--surface); border: 1px solid var(--border); border-radius: 4px; padding: 15px; display: flex; flex-direction: column; min-height: 0; overflow: hidden; position: relative; }
.panel-title { font-family: var(--f-mono); font-size: 12px; letter-spacing: 0.2em; text-transform: uppercase; color: var(--accent); margin-bottom: 15px; border-bottom: 1px solid var(--border); padding-bottom: 5px; display: flex; align-items: center; justify-content: space-between; }
.panel-title::after { content: '///'; color: var(--text-dark); }
</style>
</head>
<body>
</body>
</html>
```

---

### Task 3: 實作左欄 - 選手基本檔案與即時生理數據 (Implement Left Column - Profile & Biometrics)

`Files:`
- Modify: `sample/intro.html`

- [ ] `Step 1: 新增左欄 HTML 元件與 CSS`
  在 `.main-grid` 下加入左邊欄，呈現選手資料與心電圖 SVG 動畫。

```html
<!-- Left Column: Profile & Biometrics -->
<div class="panel">
  <div class="panel-title">ATHLETE.PROFILE</div>
  <div class="bio-card">
    <div class="bio-row"><span class="label">NAME:</span> <span class="val name-highlight">KOA "STORM" MERCER</span></div>
    <div class="bio-row"><span class="label">ORIGIN:</span> <span class="val">MAUI, HAWAII</span></div>
    <div class="bio-row"><span class="label">RANK:</span> <span class="val rankings">WCT #03</span></div>
    <div class="bio-row"><span class="label">STATUS:</span> <span class="val green-glow">SYSTEM ACTIVE</span></div>
  </div>

  <div class="panel-title" style="margin-top: 20px;">BIOMETRIC.TELEM</div>
  <div class="biometrics">
    <div class="bio-heart">
      <span class="heart-bpm" id="bpm-val">142</span>
      <span class="heart-unit">BPM</span>
      <svg class="heart-pulse-svg" viewBox="0 0 100 30">
        <path class="ecg-line" d="M0,15 L30,15 L35,10 L40,25 L45,5 L50,18 L55,15 L100,15" />
      </svg>
    </div>
    <div class="bio-gforce">
      <div class="g-header">
        <span>G-FORCE SENSOR</span>
        <span id="g-val">3.2 G</span>
      </div>
      <div class="g-bar-container">
        <div class="g-bar" id="g-bar" style="width: 64%;"></div>
      </div>
    </div>
  </div>
</div>
```

- [ ] `Step 2: 新增左欄 CSS 動畫與佈局`
  加入 ECG 動畫 (`stroke-dasharray` 循環) 以及心跳發光效果。

```css
.bio-card { display: flex; flex-direction: column; gap: 8px; font-family: var(--f-mono); font-size: 13px; }
.bio-row { display: flex; justify-content: space-between; border-bottom: 1px dashed rgba(0, 240, 255, 0.08); padding-bottom: 6px; }
.bio-row .label { color: var(--text-dim); }
.bio-row .val { color: var(--text); text-align: right; }
.name-highlight { color: var(--marker) !important; font-weight: 600; }
.rankings { color: var(--accent) !important; text-shadow: 0 0 8px var(--accent-glow); }
.green-glow { color: #00ff66 !important; text-shadow: 0 0 8px rgba(0, 255, 102, 0.4); }

.biometrics { flex: 1; display: flex; flex-direction: column; justify-content: space-around; }
.bio-heart { background: rgba(255, 0, 85, 0.05); border: 1px solid rgba(255, 0, 85, 0.2); border-radius: 4px; padding: 12px; position: relative; overflow: hidden; display: flex; align-items: baseline; justify-content: space-between; }
.heart-bpm { font-family: var(--f-mono); font-size: 36px; font-weight: bold; color: var(--warning); text-shadow: 0 0 10px var(--warning-glow); animation: heartbeat 0.82s infinite alternate; }
.heart-unit { font-family: var(--f-mono); font-size: 12px; color: var(--warning); }
.heart-pulse-svg { width: 120px; height: 35px; }
.ecg-line { fill: none; stroke: var(--warning); stroke-width: 1.5; stroke-dasharray: 100; stroke-dashoffset: 100; animation: ecgRun 1.5s linear infinite; }

@keyframes heartbeat {
  0% { transform: scale(1); }
  30% { transform: scale(1.05); }
  40% { transform: scale(1); }
  60% { transform: scale(1.1); }
  100% { transform: scale(1); }
}
@keyframes ecgRun {
  to { stroke-dashoffset: 0; }
}

.bio-gforce { display: flex; flex-direction: column; gap: 5px; font-family: var(--f-mono); font-size: 11px; }
.g-header { display: flex; justify-content: space-between; }
.g-bar-container { width: 100%; height: 6px; background: var(--text-dark); border-radius: 3px; overflow: hidden; }
.g-bar { height: 100%; background: var(--accent); box-shadow: 0 0 8px var(--accent-glow); transition: width 0.3s ease; }
```

---

### Task 4: 實作中欄 - 選手全螢幕大圖與雷達掃描疊加層 (Implement Center Column - Surfer & Radar HUD)

`Files:`
- Modify: `sample/intro.html`

- [ ] `Step 1: 新增中欄 HTML 結構`
  置入生成的衝浪選手照片，並以雙色調 filter 濾鏡處理。疊加旋轉的 SVG 雷達網格與掃描線。

```html
<!-- Center Column: Surfer Visual & Radar Overlay -->
<div class="surfer-container">
  <div class="hud-scanner"></div>
  <div class="radar-hud">
    <svg class="radar-svg" viewBox="0 0 200 200">
      <circle cx="100" cy="100" r="90" class="radar-circle-outer" />
      <circle cx="100" cy="100" r="60" class="radar-circle-mid" />
      <circle cx="100" cy="100" r="30" class="radar-circle-inner" />
      <line x1="10" y1="100" x2="190" y2="100" class="radar-axis" />
      <line x1="100" y1="10" x2="100" y2="190" class="radar-axis" />
      <path d="M100,100 L170,50" class="radar-sweep-line" />
    </svg>
  </div>
  <img src="../assets/images/personal/surfer_action.jpg" alt="Koa Mercer" class="surfer-img" id="surfer-image" onerror="handleImageError(this)">
  <div class="surfer-silhouette" id="silhouette" style="display:none;">
    <!-- SVG silhouette fallback -->
    <svg viewBox="0 0 100 100">
      <path fill="#00f0ff" opacity="0.2" d="M50 10 Q65 40 80 90 L20 90 Z" />
    </svg>
  </div>
  <div class="surfer-name-hud">
    <h1 class="surfer-title">KOA MERCER</h1>
    <p class="surfer-subtitle">THE STRIKE ZONE // MAUI SWELL STAGE</p>
  </div>
</div>
```

- [ ] `Step 2: 新增中欄 CSS 樣式與動畫`
  設定 `surfer-container` 的滿版效果、雙色調濾鏡與雷達旋轉。

```css
.surfer-container { position: relative; border: 1px solid var(--border); border-radius: 4px; overflow: hidden; display: flex; justify-content: center; align-items: center; background: #03060d; }
.surfer-img { width: 100%; height: 100%; object-fit: cover; filter: grayscale(1) contrast(1.1) brightness(0.7) sepia(0.2) hue-rotate(150deg) saturate(2.5); mix-blend-mode: screen; opacity: 0.85; transition: transform 0.5s; }
.surfer-container:hover .surfer-img { transform: scale(1.03); }

/* Radar graphic rotating overlay */
.radar-hud { position: absolute; inset: 0; display: grid; place-items: center; pointer-events: none; z-index: 5; }
.radar-svg { width: 70%; height: 70%; opacity: 0.25; animation: spinRadar 25s linear infinite; }
.radar-circle-outer, .radar-circle-mid, .radar-circle-inner { fill: none; stroke: var(--accent); stroke-width: 0.7; stroke-dasharray: 4 2; }
.radar-axis { stroke: var(--accent); stroke-width: 0.5; opacity: 0.4; }
.radar-sweep-line { stroke: var(--accent); stroke-width: 2; opacity: 0.8; stroke-linecap: round; filter: drop-shadow(0 0 5px var(--accent-glow)); }

@keyframes spinRadar {
  to { transform: rotate(360deg); }
}

/* Vertical sweep line */
.hud-scanner { position: absolute; left: 0; right: 0; height: 4px; background: linear-gradient(180deg, transparent, var(--accent), transparent); opacity: 0.4; z-index: 4; pointer-events: none; animation: scanLines 4.5s linear infinite; }
@keyframes scanLines {
  0% { top: 0%; }
  50% { top: 100%; }
  100% { top: 0%; }
}

.surfer-name-hud { position: absolute; bottom: 20px; left: 20px; z-index: 6; font-family: var(--f-head); }
.surfer-title { font-size: 32px; font-weight: 800; color: #fff; letter-spacing: -1px; text-shadow: 0 2px 10px rgba(0, 0, 0, 0.8); }
.surfer-subtitle { font-family: var(--f-mono); font-size: 11px; color: var(--accent); letter-spacing: 2px; }
```

---

### Task 5: 實作右欄 - 賽事表現統計與 SVG 技能雷達圖 (Implement Right Column - Stats & Skills Radar Chart)

`Files:`
- Modify: `sample/intro.html`

- [ ] `Step 1: 新增右欄 HTML 元件與統計圖`
  在 `.main-grid` 下加入右邊欄，以 SVG 繪製五角技能雷達圖與柱狀圖。

```html
<!-- Right Column: Stats & Performance -->
<div class="panel">
  <div class="panel-title">PERFORMANCE.STATS</div>
  <div class="stats-grid">
    <div class="stat-box"><span class="stat-lbl">MAX SPEED</span><span class="stat-val">48.5 <small>KM/H</small></span></div>
    <div class="stat-box"><span class="stat-lbl">AIR TIME</span><span class="stat-val">3.2 <small>SEC</small></span></div>
    <div class="stat-box"><span class="stat-lbl">MAX WAVE</span><span class="stat-val">12.8 <small>MTR</small></span></div>
  </div>

  <div class="panel-title" style="margin-top: 20px;">SKILLS.RATING</div>
  <div class="radar-chart-container">
    <svg class="skills-radar-svg" viewBox="0 0 100 100">
      <!-- Polygon backdrops -->
      <polygon points="50,10 88,38 73,82 27,82 12,38" class="radar-poly-bg" />
      <polygon points="50,23 79,44 67,74 33,74 21,44" class="radar-poly-bg" />
      <polygon points="50,37 69,50 62,65 38,65 31,50" class="radar-poly-bg" />
      <!-- Axis lines -->
      <line x1="50" y1="50" x2="50" y2="10" class="radar-axis-line" />
      <line x1="50" y1="50" x2="88" y2="38" class="radar-axis-line" />
      <line x1="50" y1="50" x2="73" y2="82" class="radar-axis-line" />
      <line x1="50" y1="50" x2="27" y2="82" class="radar-axis-line" />
      <line x1="50" y1="50" x2="12" y2="38" class="radar-axis-line" />
      <!-- Data Area -->
      <polygon points="50,18 84,40 68,76 38,70 19,32" class="radar-data-poly" />
      <!-- Text Labels -->
      <text x="50" y="7" class="radar-label">SPEED</text>
      <text x="91" y="38" class="radar-label">POWER</text>
      <text x="76" y="88" class="radar-label">FLOW</text>
      <text x="18" y="88" class="radar-label">TUBE</text>
      <text x="4" y="38" class="radar-label">AIR</text>
    </svg>
  </div>

  <div class="panel-title" style="margin-top: 15px;">STAGE.LOGS</div>
  <div class="bar-chart">
    <div class="bar-col"><div class="bar-fill" style="height: 85%;"></div><span class="bar-lbl">S1</span></div>
    <div class="bar-col"><div class="bar-fill" style="height: 92%;"></div><span class="bar-lbl">S2</span></div>
    <div class="bar-col"><div class="bar-fill" style="height: 70%;"></div><span class="bar-lbl">S3</span></div>
    <div class="bar-col"><div class="bar-fill" style="height: 95%;"></div><span class="bar-lbl">S4</span></div>
    <div class="bar-col"><div class="bar-fill" style="height: 88%;"></div><span class="bar-lbl">S5</span></div>
  </div>
</div>
```

- [ ] `Step 2: 新增右欄 CSS 樣式與圖表細節`
  設定技能雷達圖背景與數值多邊形的透明發光天藍色樣式。

```css
.stats-grid { display: grid; grid-template-columns: repeat(3, 1fr); gap: 10px; font-family: var(--f-mono); }
.stat-box { display: flex; flex-direction: column; background: rgba(0, 240, 255, 0.03); border: 1px solid rgba(0, 240, 255, 0.08); padding: 8px 4px; border-radius: 4px; text-align: center; }
.stat-lbl { font-size: 9px; color: var(--text-dim); margin-bottom: 4px; }
.stat-val { font-size: 16px; font-weight: bold; color: var(--accent); }
.stat-val small { font-size: 9px; font-weight: normal; color: var(--text-dim); }

.radar-chart-container { display: flex; justify-content: center; align-items: center; flex: 1; padding: 10px; }
.skills-radar-svg { width: 85%; height: auto; }
.radar-poly-bg { fill: none; stroke: var(--text-dark); stroke-width: 0.5; }
.radar-axis-line { stroke: var(--text-dark); stroke-width: 0.5; stroke-dasharray: 2 2; }
.radar-data-poly { fill: rgba(0, 240, 255, 0.22); stroke: var(--accent); stroke-width: 1.5; filter: drop-shadow(0 0 5px var(--accent-glow)); }
.radar-label { font-family: var(--f-mono); font-size: 7px; fill: var(--text-dim); text-anchor: middle; }

.bar-chart { display: flex; justify-content: space-around; align-items: flex-end; height: 75px; padding-top: 10px; border-bottom: 1px dashed rgba(0, 240, 255, 0.1); }
.bar-col { display: flex; flex-direction: column; align-items: center; width: 14%; height: 100%; justify-content: flex-end; }
.bar-fill { width: 100%; background: linear-gradient(0deg, var(--accent) 0%, rgba(0,240,255,0.4) 100%); border-radius: 2px 2px 0 0; box-shadow: 0 0 8px var(--accent-glow); animation: growBars 1.2s cubic-bezier(0.2, 0.8, 0.2, 1) forwards; transform-origin: bottom; }
.bar-lbl { font-family: var(--f-mono); font-size: 8px; color: var(--text-dim); margin-top: 4px; }

@keyframes growBars {
  0% { transform: scaleY(0); }
  100% { transform: scaleY(1); }
}
```

---

### Task 6: 實作底欄控制按鈕與互動 JS 邏輯 (Implement Footer Controls & JS Logic)

`Files:`
- Modify: `sample/intro.html`

- [ ] `Step 1: 新增底欄 HTML 與按鈕`
  在 `<footer>` 下新增 `[REPLAY]` 按鈕以及 `[PROCEED]` 按鈕。

```html
<!-- Footer Control Bar -->
<footer>
  <div class="keyboard-tips">PRESS [ENTER] TO JOIN COMPETITION / [R] REPLAY DETAILED MATRIX</div>
  <div class="action-buttons">
    <button class="btn btn-replay" id="btn-replay">↻ REPLAY DATA</button>
    <a href="sys-zero.html" class="btn btn-proceed">ENTER ARENA →</a>
  </div>
</footer>
```

- [ ] `Step 2: 加入 CSS 按鈕與底欄佈局樣式`
  實作 HUD 樣式的玻璃發光按鈕效果。

```css
footer { display: flex; justify-content: space-between; align-items: center; font-family: var(--f-mono); border-top: 1px solid var(--border); padding-top: 15px; margin-top: auto; }
.keyboard-tips { font-size: 9.5px; color: var(--text-dark); letter-spacing: 0.1em; }
.action-buttons { display: flex; gap: 10px; }
.btn { font-family: var(--f-mono); font-size: 11px; letter-spacing: 0.15em; text-transform: uppercase; padding: 8px 16px; border-radius: 3px; cursor: pointer; transition: all 0.2s ease; border: 1px solid var(--border); background: rgba(12, 20, 36, 0.5); color: var(--accent); text-decoration: none; }
.btn:hover { background: rgba(0, 240, 255, 0.1); border-color: var(--accent); box-shadow: 0 0 12px var(--accent-glow); transform: translateY(-1px); }
.btn-proceed { background: var(--accent); color: var(--bg); font-weight: 600; }
.btn-proceed:hover { background: #5df7ff; color: var(--bg); }
```

- [ ] `Step 3: 實作心跳與重播 JS 邏輯`
  在 `sample/intro.html` 尾部置入 JS，以動態改變 BPM 數值，以及實作 `R` 鍵重播開機動畫、`Enter` 鍵跳轉。

```html
<script>
(function() {
  const bpmVal = document.getElementById('bpm-val');
  const gVal = document.getElementById('g-val');
  const gBar = document.getElementById('g-bar');
  const btnReplay = document.getElementById('btn-replay');

  // 心率與 G-Force 即時模擬
  setInterval(() => {
    // BPM 在 138 - 146 之間微調
    const currentBpm = Math.floor(138 + Math.random() * 9);
    bpmVal.textContent = currentBpm;

    // G-Force 在 2.9 - 3.8 之間微調
    const currentG = (2.9 + Math.random() * 0.9).toFixed(1);
    gVal.textContent = currentG + ' G';
    gBar.style.width = (parseFloat(currentG) / 5.0 * 100) + '%';
  }, 1000);

  // 重新撥放動畫邏輯
  function replaySequence() {
    const stage = document.querySelector('.stage');
    const fresh = stage.cloneNode(true);
    stage.replaceWith(fresh);
    bindEvents(fresh);
  }

  function bindEvents(root) {
    const replay = root.querySelector('#btn-replay');
    if (replay) {
      replay.addEventListener('click', replaySequence);
    }
  }

  bindEvents(document);

  // 鍵盤監聽事件
  document.addEventListener('keydown', (e) => {
    if (e.key === 'Enter' || e.key === ' ') {
      e.preventDefault();
      window.location.href = 'sys-zero.html';
    } else if (e.key.toLowerCase() === 'r') {
      replaySequence();
    }
  });
})();

function handleImageError(img) {
  img.style.display = 'none';
  document.getElementById('silhouette').style.display = 'block';
}
</script>
```

---

### Task 7: 選手照片失效降級處理與驗證 (Verify & Fallback Verification)

`Files:`
- Modify: `sample/intro.html`

- [ ] `Step 1: 驗證圖片加載失敗時的降級 SVG 元件`
  手動將 `<img>` src 設為錯誤路徑，確認會隱藏主圖並顯示 `.surfer-silhouette` 的備用 SVG 衝浪板背景。
