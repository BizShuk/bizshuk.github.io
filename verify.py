from playwright.sync_api import sync_playwright

def run_cuj(page):
    print("Navigating to index.html...")
    page.goto("http://localhost:8000/index.html")
    page.wait_for_timeout(500)
    page.screenshot(path="/app/gallery.png", full_page=True)

    print("Navigating to resume.html...")
    page.goto("http://localhost:8000/resume.html")
    page.wait_for_timeout(500)
    page.screenshot(path="/app/resume_hero.png")

    page.evaluate("window.scrollBy(0, 800)")
    page.wait_for_timeout(500)
    page.screenshot(path="/app/resume_exp.png")

if __name__ == "__main__":
    with sync_playwright() as p:
        browser = p.chromium.launch(headless=True)
        context = browser.new_context(
            record_video_dir="/app/"
        )
        page = context.new_page()
        page.set_viewport_size({"width": 1280, "height": 800})
        try:
            run_cuj(page)
        finally:
            context.close()
            browser.close()
