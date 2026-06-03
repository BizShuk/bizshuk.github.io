const { chromium } = require('playwright');

(async () => {
  const browser = await chromium.launch();
  const page = await browser.newPage();
  await page.setViewportSize({ width: 1280, height: 800 });

  console.log("Navigating to index.html...");
  await page.goto('http://localhost:8000/index.html', { waitUntil: 'networkidle' });
  await page.screenshot({ path: 'gallery.png', fullPage: true });

  console.log("Navigating to resume.html...");
  await page.goto('http://localhost:8000/resume.html', { waitUntil: 'networkidle' });
  await page.screenshot({ path: 'resume_hero.png' });

  // Scroll down a bit to see the timeline
  await page.evaluate(() => window.scrollBy(0, 800));
  await page.screenshot({ path: 'resume_exp.png' });

  await browser.close();
  console.log("Screenshots captured successfully.");
})();
