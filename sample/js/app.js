/* app.js — single-page app router
 *
 *   - Menu 資料結構 (MENU) 定義在 page state 中,由 JS 注入 sidebar
 *   - localStorage 記住最後一次檢視的 view,跨 session 復原
 *   - hash 路由支援深連結 (#resume) 與瀏覽器上一頁/下一頁
 *   - data-jump 按鈕讓 hero 等區塊直接切換 view
 */

const STORAGE_KEY = 'bizshuk:active-view';

// ─── Menu 資料結構 (存放於 page state) ────────────────────────────
const MENU = [
  { id: 'home',    label: 'Home',    icon: '✦' },
  { id: 'resume',  label: 'Resume',  icon: '◧' },
  { id: 'socials', label: 'Socials', icon: '◉' },
  { id: 'links',   label: 'Links',   icon: '⇗' }
];

// ─── 工具 ─────────────────────────────────────────────────────────
const isValidView = id => MENU.some(m => m.id === id);

function readInitialView() {
  const fromHash = (location.hash || '').replace('#', '');
  if (isValidView(fromHash)) return fromHash;

  try {
    const saved = localStorage.getItem(STORAGE_KEY);
    if (isValidView(saved)) return saved;
  } catch (_) { /* localStorage 不可用時忽略 */ }

  return MENU[0].id;
}

// ─── 渲染 sidebar menu (menu 從資料結構生成) ──────────────────────
function renderMenu() {
  const ul = document.getElementById('menu');
  if (!ul) return;

  ul.innerHTML = MENU.map(m => `
    <li>
      <a href="#${m.id}" class="menu-item" data-view="${m.id}" role="button">
        <span class="menu-item-icon" aria-hidden="true">${m.icon}</span>
        <span>${m.label}</span>
      </a>
    </li>
  `).join('');
}

// ─── View 切換 ─────────────────────────────────────────────────────
function showView(id, { persist = true, scroll = true } = {}) {
  if (!isValidView(id)) id = MENU[0].id;

  document.querySelectorAll('.view').forEach(v => {
    v.classList.toggle('active', v.dataset.view === id);
  });
  document.querySelectorAll('.menu-item').forEach(a => {
    a.classList.toggle('active', a.dataset.view === id);
  });

  if (persist) {
    try { localStorage.setItem(STORAGE_KEY, id); } catch (_) {}
  }
  if (location.hash !== '#' + id) {
    history.replaceState(null, '', '#' + id);
  }
  if (scroll) {
    window.scrollTo({ top: 0, behavior: 'smooth' });
  }

  // 重新觸發當前 view 內 fade-in 元素的進場動畫
  rebindFadeIn(id);
}

// IntersectionObserver 在 module-level 維持單一實例
const fadeObserver = new IntersectionObserver(entries => {
  entries.forEach(e => {
    if (e.isIntersecting) {
      e.target.classList.add('visible');
      fadeObserver.unobserve(e.target);
    }
  });
}, { threshold: 0.1 });

function rebindFadeIn(viewId) {
  const view = document.querySelector(`.view[data-view="${viewId}"]`);
  if (!view) return;

  view.querySelectorAll('.fade-in').forEach(el => {
    el.classList.remove('visible');
    fadeObserver.observe(el);
  });
}

// ─── 初始化 ───────────────────────────────────────────────────────
document.addEventListener('DOMContentLoaded', () => {
  renderMenu();

  // menu 點擊: 攔截交給 showView (避免頁面跳動)
  document.querySelectorAll('.menu-item').forEach(a => {
    a.addEventListener('click', e => {
      e.preventDefault();
      showView(a.dataset.view);
    });
  });

  // hero 等區塊的 data-jump 按鈕
  document.querySelectorAll('[data-jump]').forEach(a => {
    a.addEventListener('click', e => {
      e.preventDefault();
      showView(a.dataset.jump);
    });
  });

  // 瀏覽器上一頁/下一頁 → 跟著 hash 走
  window.addEventListener('hashchange', () => {
    const id = (location.hash || '').replace('#', '');
    if (isValidView(id)) showView(id, { persist: false });
  });

  // 首次載入
  showView(readInitialView(), { persist: false, scroll: false });
});
