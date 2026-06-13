### bizshuk.github.io

`https://bizshuk.github.io` is the personal portfolio site for Shuk Liu. The site
runs as a static GitHub Pages deployment with two top-level pages, a small JSON
data layer, and external CSS/JS files. There is no backend, no build step, and
no framework runtime — the production output is the working tree itself.

- `index.html` is the main entry point. It renders a gallery of links driven by
  [data/gallery.json](data/gallery.json), and from there the user can navigate
  to [resume.html](resume.html) and the rest of the site.

## 業務領域 (Business Domains)

### 個人作品集 (Personal Portfolio)

The site is a single-page portfolio plus a long-form resume. It is the public
face of the owner on GitHub and is consumed as a static artifact by GitHub
Pages.

`領域流程 (Domain Flow):`

1. Visitor opens `https://bizshuk.github.io/` (served as `index.html`).
2. Browser executes `js/gallery.js`, which `fetch`-es `data/gallery.json` on
   `DOMContentLoaded`.
3. Each gallery entry is materialised as either an `<a>` (when `link` is set)
   or a `<div>` (decorative tile), with the `img_title` rendered as a hover
   overlay and the `style.backgroundImage` applied to the tile itself.
4. Tiles fade in with a staggered `fadeInUp` animation (delay =
   `index * 0.05s`).
5. The `Resume (non-updated)` tile links to `resume.html`, which renders the
   full CV. The remaining tiles fan out to external profiles (GitHub,
   LinkedIn, Facebook, Imgur, LeeCode, certificates, projects).

`核心實體 (Key Entities):`

- `GalleryItem` — `{ link, img_title, style.backgroundImage }`. Lives in
  `data/gallery.json` and is consumed by `js/gallery.js`.
- `ResumeSection` — `hero`, `photo-gallery`, `experience` (timeline),
  `skills`, `education`, `demo`. Each section is a static block in
  `resume.html` styled by `css/resume.css`.
- `Photo` — image entry in the horizontal scroll strip on `resume.html`
  (15 Imgur-hosted photos).

`相關處理器 (Related Handlers):`

- `DOMContentLoaded` listener in `js/gallery.js` — bootstraps the gallery.
- `mousedown` / `mousemove` / `mouseleave` / `mouseup` listeners in
  `js/resume.js` — enable drag-to-scroll on the photo strip.
- `onerror` attribute on the two `<img>` tags in the education section —
  fallback to a `via.placeholder.com` URL when the local JPG is missing.

---

### 履歷展示 (Resume Showcase)

A single long-form HTML page that presents the owner's career history,
skills, education, and project demos. All content is hand-written into
`resume.html`; there is no template engine and no data binding.

`領域流程 (Domain Flow):`

1. Visitor clicks a gallery tile or navigates directly to `resume.html`.
2. `js/resume.js` injects the current year into the footer on first load
   (`document.getElementById('year').textContent`).
3. The photo strip at the top of the page supports click-and-drag scrolling
   on desktop via four mouse listeners in `js/resume.js`.
4. Each section (`#experience`, `#skills`, `#education`, `#demo`) is rendered
   from inline HTML; education cards and demo cards each have an `onerror`
   handler for image fallbacks.

`核心實體 (Key Entities):`

- `TimelineItem` — work experience entry (employer, role, period, bullet
  list). Currently three items: Gamesofa Inc., Department of Civil Servant
  Development, National Central University.
- `SkillTag` — pill-style chip inside a `Skills & Technologies` category
  (Programming Languages, Frameworks & Libraries, Environment & Tools,
  Concepts & APIs).
- `EducationCard` — university card with a campus photo, name, degree, and
  date range. Cards: NCU (MS), NDHU (BS).
- `DemoCard` — anchor-tag card linking to a Flickr-hosted screenshot of a
  past project (User Analysis Dashboard, Survey System, Demo Application).

`相關處理器 (Related Handlers):`

- `year` span update in `js/resume.js` (footer copyright year).
- Drag-to-scroll handlers in `js/resume.js` for `.photo-gallery`.
- `<img onerror="this.src='https://via.placeholder.com/400x200?text=...'">`
  inline fallbacks in the education section.

---

## 領域關聯 (Domain Relationships)

- The Gallery domain is the only entry point and the only page that issues a
  network request at runtime (`fetch('data/gallery.json')`).
- The Resume domain is reachable only from the Gallery (via the
  `Resume (non-updated)` tile) or by direct URL. It issues no network
  requests of its own.
- Both pages share visual language (Montserrat font, hover-overlay
  transitions, fade-in animations) but are otherwise independent — there is
  no shared layout component or shared header/footer file.
- `data/params.json` is a legacy GitHub Pages metadata file; it is not
  consumed by any current handler and can be considered dead weight.

## 使用方式 (Usage)

- View the portfolio: open `index.html` in a browser, or visit
  `https://bizshuk.github.io/`.
- View the resume: click the `Resume (non-updated)` tile, or visit
  `https://bizshuk.github.io/resume.html` directly.
- Edit the gallery: append or remove entries in `data/gallery.json`. An
  entry with an empty `link` renders as a decorative tile (no anchor, no
  hover navigation); an entry with `link` becomes an anchor that opens in a
  new tab with `rel="noopener noreferrer"`.

## 改善建議 (Improvement Suggestions)

- [ ] Extract the shared header/hero/footer markup from `index.html` and
      `resume.html` into a small client-side include (e.g. a `js/partials.js`
      or a templated fragment) so layout changes only need to be made in
      one place. The two pages currently duplicate the font import and the
      favicon link.
- [ ] Add a build/lint step (e.g. an `html-validate` or `markdownlint` hook)
      to catch dead links, broken image refs, and stale copy. The site has
      no CI and `data/params.json` is still the original GitHub Pages
      boilerplate — a clear sign that drift goes unnoticed.
- [ ] Move the gallery image URLs out of `data/gallery.json` and onto local
      files in `assets/images/gallery/`. The site currently depends on
      `i.imgur.com` for both the gallery tiles and the resume photo strip;
      if Imgur hot-link policy changes (as it has before) the portfolio
      silently loses its cover art.
- [ ] Add a `link` for the seven decorative tiles in `data/gallery.json` or
      document explicitly that they are placeholders. Right now seven
      tiles look like navigation options but do nothing on click, which is
      confusing on the landing page.
- [ ] The resume page hard-codes three jobs and two schools. Consider
      splitting the structured content (`#experience`, `#skills`,
      `#education`, `#demo`) out of `resume.html` into a JSON file (mirroring
      how the gallery is data-driven) so the resume can be re-rendered in
      different layouts (PDF, JSON-LD for hResume, etc.) without editing
      HTML by hand.
- [ ] Add `aria-label` / `alt` text and keyboard focus styles for the
      gallery tiles. The current drag-to-scroll handler on the resume photo
      strip is mouse-only and ignores `scroll-snap` keyboard navigation.
- [ ] `data/params.json` is unused legacy content from the original
      GitHub Pages generator. Either delete it or document its role in
      `README.md` so the next maintainer does not assume it is a source of
      truth.
