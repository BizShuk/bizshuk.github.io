### bizshuk.github.io

`https://bizshuk.github.io` is the personal portfolio site for Shuk Liu. The site
runs as a static GitHub Pages deployment with one gallery landing page plus
self-contained subpages under `pkg/`, a small JSON data layer, and external
CSS/JS files. There is no backend, no build step, and no framework runtime —
the production output is the working tree itself.

- `index.html` is the main entry point. It renders a gallery of links driven by
  [data/gallery.json](data/gallery.json), and from there the user can navigate
  to [pkg/resume/index.html](pkg/resume/index.html),
  [pkg/surf/index.html](pkg/surf/index.html) and the rest of the site.
- `pkg/resume/index.html` is the résumé page — a sunny, single-column CV whose
  content is kept in step with [pkg/resume/Resume.md](pkg/resume/Resume.md), with
  the exported PDF offered as a download.
- `pkg/surf/index.html` is an ocean-editorial magazine page built around the
  owner's own Uluwatu drop-in photo — story, surf-spot footprints
  (Bali / Lombok / Taiwan), and quiver.

## 業務領域 (Business Domains)

### 個人作品集 (Personal Portfolio)

The site is a single-page portfolio plus a long-form resume. It is the public
face of the owner on GitHub and is consumed as a static artifact by GitHub
Pages.

`領域流程 (Domain Flow):`

1. Visitor opens `https://bizshuk.github.io/` (served as `index.html`).
2. Browser executes the page's inline script, which `fetch`-es `data/gallery.json` on
   `DOMContentLoaded`.
3. Each gallery entry is materialised as either an `<a>` (when `link` is set)
   or a `<div>` (decorative tile), with the `img_title` rendered as a hover
   overlay and the `style.backgroundImage` applied to the tile itself.
4. Tiles fade in with a staggered `fadeInUp` animation (delay =
   `index * 0.05s`).
5. The résumé tile links to `pkg/resume/index.html`, which renders the full
   CV. The remaining tiles fan out to external profiles (GitHub, LinkedIn,
   Instagram, Imgur, LeeCode).

`核心實體 (Key Entities):`

- `GalleryItem` — `{ link, img_title, style.backgroundImage }`. Lives in
  `data/gallery.json` and is consumed by the inline script in `index.html`.
- `ResumeSection` — `hero`, `education`, `skills`, `work experience`,
  `other experience`. Each section is a static block in
  `pkg/resume/index.html`, styled by that page's own inline CSS.

`相關處理器 (Related Handlers):`

- `DOMContentLoaded` listener in `index.html` — bootstraps the gallery.
- `IntersectionObserver` in the inline scripts of `pkg/resume/index.html` and
  `pkg/surf/index.html` — reveals sections on scroll as progressive enhancement.

---

### 履歷展示 (Resume Showcase)

A single long-form HTML page that presents the owner's career history,
skills, and education. All content is hand-written into
`pkg/resume/index.html`, mirroring `pkg/resume/Resume.md`; there is no
template engine and no data binding.

`領域流程 (Domain Flow):`

1. Visitor clicks a gallery tile or navigates directly to
   `https://bizshuk.github.io/pkg/resume/`.
2. The hero introduces the owner (portrait, one-paragraph summary, contact
   chips) and carries the last-updated stamp; a floating icon in the top
   right downloads `assets/Resume-ShukLiu.pdf`.
3. Sections render from inline HTML and are visible without JavaScript; when
   scripting is available they fade in on scroll and the footer year updates.

`核心實體 (Key Entities):`

- `JobEntry` — one role: title, company and location, period, and bullet
  list. Seven entries, from the current self-directed AI work back to
  National Central University.
- `SkillTag` — pill-style chip inside a skill row (Certificate, Language,
  Test, Infra / Tooling, Concept, Methodology).
- `EducationCard` — school card with degree and date range: NCU (MS),
  NDHU (BS), Heping High School.
- `OtherExperience` — non-employment entries (working holidays, self-study,
  alternative civilian service).

`相關處理器 (Related Handlers):`

- `year` span update in the page's inline script (footer copyright year).
- `IntersectionObserver` reveal, armed only after the script runs so a
  script failure never hides content.

---

## 領域關聯 (Domain Relationships)

- The Gallery domain is the only entry point and the only page that issues a
  network request at runtime (`fetch('data/gallery.json')`).
- The Resume domain is reachable only from the Gallery (via the résumé tile)
  or by direct URL. It issues no network requests of its own.
- Every page shares the Montserrat font and a fade-in idiom but is otherwise
  independent — there is no shared layout component or shared header/footer
  file, and each `pkg/` subpage carries its own styles.
- `data/params.json` is a legacy GitHub Pages metadata file; it is not
  consumed by any current handler and can be considered dead weight.

## 使用方式 (Usage)

- View the portfolio: open `index.html` in a browser, or visit
  `https://bizshuk.github.io/`.
- View the resume: click the résumé tile, or visit
  `https://bizshuk.github.io/pkg/resume/` directly. The PDF is at
  `https://bizshuk.github.io/pkg/resume/assets/Resume-ShukLiu.pdf`.
- View the surf page: `https://bizshuk.github.io/pkg/surf/`.
- Edit the gallery: append or remove entries in `data/gallery.json`. An
  entry with an empty `link` renders as a decorative tile (no anchor, no
  hover navigation); an entry with `link` becomes an anchor that opens in a
  new tab with `rel="noopener noreferrer"`.

## 改善建議 (Improvement Suggestions)

- [ ] The shared head boilerplate (font import, favicon, referrer policy) is
      duplicated across `index.html` and the `pkg/` subpages, and each page now
      inlines its own CSS/JS. Consider a small generator or include so shared
      layout changes are made in one place.
- [ ] Add a build/lint step (e.g. an `html-validate` or `markdownlint` hook)
      to catch dead links, broken image refs, and stale copy. The site has
      no CI and `data/params.json` is still the original GitHub Pages
      boilerplate — a clear sign that drift goes unnoticed.
- [ ] Move the gallery image URLs out of `data/gallery.json` and onto local
      files in `assets/images/gallery/`. The site currently depends on
      `i.imgur.com` for the gallery tiles; if Imgur hot-link policy changes
      (as it has before) the portfolio silently loses its cover art.
- [ ] Add a `link` for the seven decorative tiles in `data/gallery.json` or
      document explicitly that they are placeholders. Right now seven
      tiles look like navigation options but do nothing on click, which is
      confusing on the landing page.
- [ ] The resume page duplicates `pkg/resume/Resume.md` by hand. Consider
      generating the page from that Markdown (or a JSON derivative) so the
      page, the PDF and the Markdown cannot drift apart.
- [ ] Add `aria-label` / `alt` text and keyboard focus styles for the
      gallery tiles, which are currently reachable but visually unmarked
      when focused.
- [ ] `data/params.json` is unused legacy content from the original
      GitHub Pages generator. Either delete it or document its role in
      `README.md` so the next maintainer does not assume it is a source of
      truth.
