# 2026-09-05 客製化 404 頁面與「小朋友下樓梯」遊戲規格

為 `bizshuk.github.io` 的 GitHub Pages 客製化 404 頁面 (`404.html`) 與內嵌「小朋友下樓梯」（NS-SHAFT）遊戲規格說明。

## 需求與目標 (Requirements & Goals)

1. **GitHub Pages 404 規格**：在 repo 根目錄提供 `404.html`，當訪客進入不存在網址時自動回傳此頁面。
2. **單一檔案 (Single HTML File)**：HTML、CSS 樣式與 Canvas JavaScript 邏輯一體化封裝在 `404.html`，零外部打包或框架依賴。
3. **「小朋友下樓梯」遊戲 (NS-SHAFT Mechanics)**：
   - 玩家角色物理運動（重力、左右移動加速度、與階梯碰撞、站在階梯上移動）。
   - 多種階梯（普通、針刺、彈簧、左右輸送帶、易碎翻轉板）。
   - 天花板尖刺（頂撞扣生命值）與底部深淵（掉落出鏡直接死亡）。
   - 生命值（HP）、地下樓層（BxxF）計分與 LocalStorage 最高分記錄。
   - 原生 Web Audio API 復古 8-bit 音效與靜音開關。
4. **介面控制**：
   - 「開始遊戲」按鈕（Start Button）與說明覆蓋層。
   - 「再玩一次」按鈕（Retry Button）與 Game Over 結算覆蓋層。
   - 「返回首頁」導覽連結。
   - 支援鍵盤（方向鍵、A/D 鍵）與手機/觸控虛擬按鈕。
