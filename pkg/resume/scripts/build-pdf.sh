#!/usr/bin/env bash
set -euo pipefail

RESUME_DIR="$(cd "$(dirname "$0")/.." && pwd)"
CHROME_BIN="${CHROME_BIN:-/Applications/Google Chrome.app/Contents/MacOS/Google Chrome}"

if [[ ! -x "$CHROME_BIN" ]]; then
  echo "ERROR: Google Chrome not found at $CHROME_BIN" >&2
  exit 1
fi

render_pdf() {
  local html_file="$1"
  local pdf_file="$2"
  "$CHROME_BIN" \
    --headless=new \
    --disable-gpu \
    --no-pdf-header-footer \
    --print-to-pdf="$pdf_file" \
    "file://$html_file"
}

render_pdf "$RESUME_DIR/templates/resume-tech.html" "$RESUME_DIR/assets/Resume-ShukLiu.pdf"
render_pdf "$RESUME_DIR/templates/resume-full.html" "$RESUME_DIR/assets/Resume-ShukLiu-full.pdf"

uv run --with pypdf python3 - "$RESUME_DIR/assets/Resume-ShukLiu.pdf" "$RESUME_DIR/assets/Resume-ShukLiu-full.pdf" <<'PY'
import sys
from pypdf import PdfReader

expected_pages = {
    sys.argv[1]: 1,
    sys.argv[2]: 2,
}
forbidden_terms = ["high school", "heping"]

for path, want_pages in expected_pages.items():
    reader = PdfReader(path)
    got_pages = len(reader.pages)
    if got_pages != want_pages:
        raise SystemExit(f"FAIL {path}: expected {want_pages} page(s), got {got_pages}")
    text = "\n".join(page.extract_text() or "" for page in reader.pages).lower()
    for term in forbidden_terms:
        if term in text:
            raise SystemExit(f"FAIL {path}: forbidden term {term!r} found in PDF text")
    print(f"OK {path}: {got_pages} page(s), no high school")
PY
