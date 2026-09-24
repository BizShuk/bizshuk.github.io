"""skills/coverletter/scripts/render_pdf.py - render a cover letter source (.md) into a one-page A4 PDF.

Usage:
    uv run --with reportlab --with pypdf python render_pdf.py <jd_stem>.coverletter.md [--resume <Resume.md>]

The letter header (name, email, phone, LinkedIn) is read from Resume.md so contact details keep a
single source of truth. The letter body comes from <letter.md>: a front matter block followed by
plain-text paragraphs separated by blank lines.

The letter is a companion of its JD file: it must sit beside the JD named in front matter `jd`, as
<jd_stem>.coverletter.md, and the PDF is written next to it as <jd_stem>.coverletter.pdf.

Exit codes: 0 ok, 1 bad input, 2 letter does not fit on one page even at the smallest font size.
"""

import argparse
import re
import sys
from pathlib import Path
from xml.sax.saxutils import escape

from pypdf import PdfReader
from reportlab.lib.colors import HexColor
from reportlab.lib.pagesizes import A4
from reportlab.lib.styles import ParagraphStyle
from reportlab.lib.units import mm
from reportlab.platypus import HRFlowable, Paragraph, SimpleDocTemplate, Spacer

REPO = Path(__file__).resolve().parents[3]
DEFAULT_RESUME = REPO / "pkg" / "resume" / "Resume.md"
REQUIRED_FIELDS = ("company", "role", "date", "jd")
LETTER_SUFFIX = ".coverletter.md"
# Body font sizes tried in order; the first one that fits on a single page wins.
FONT_SIZES = (10.5, 10.0, 9.5)
ACCENT = HexColor("#1f3a5f")
MUTED = HexColor("#555555")


def fail(code, msg):
    print(f"ERROR: {msg}", file=sys.stderr)
    sys.exit(code)


def parse_letter(path):
    text = path.read_text(encoding="utf-8")
    m = re.match(r"^---\n(.*?)\n---\n(.*)$", text, re.S)
    if not m:
        fail(1, f"{path}: missing front matter block")
    meta = {}
    for line in m.group(1).splitlines():
        if ":" in line:
            key, value = line.split(":", 1)
            meta[key.strip()] = value.strip().strip('"')
    missing = [f for f in REQUIRED_FIELDS if not meta.get(f)]
    if missing:
        fail(1, f"{path}: front matter missing {', '.join(missing)}")
    blocks = [b.strip() for b in re.split(r"\n\s*\n", m.group(2)) if b.strip()]
    if len(blocks) < 3:
        fail(1, f"{path}: body needs salutation, paragraphs and closing")
    return meta, blocks


def check_alignment(letter, meta):
    jd = REPO / meta["jd"]
    if not jd.is_file():
        fail(1, f"{letter}: jd not found: {meta['jd']}")
    expected = jd.with_name(jd.stem + LETTER_SUFFIX)
    if letter.resolve() != expected.resolve():
        fail(1, f"{letter}: letter must sit beside its JD as {expected.relative_to(REPO)}")


def parse_contact(resume_path):
    lines = resume_path.read_text(encoding="utf-8").splitlines()
    name = lines[0].lstrip("# ").strip()
    head = "\n".join(lines[:6])
    email = re.search(r"<([^>@\s]+@[^>\s]+)>", head)
    linkedin = re.search(r"\((https://www\.linkedin\.com/[^)]+)\)", head)
    phone = re.search(r"phone=(\d+)", head)
    parts = []
    if email:
        parts.append(email.group(1))
    if phone:
        digits = phone.group(1)
        parts.append(f"+{digits[:2]} {digits[2:6]} {digits[6:]}")
    if linkedin:
        parts.append(linkedin.group(1).replace("https://www.", ""))
    return name, " · ".join(parts)


def build(meta, blocks, name, contact, out, size):
    lead = size * 1.38
    body = ParagraphStyle("body", fontName="Helvetica", fontSize=size, leading=lead, spaceAfter=lead * 0.6)
    styles = {
        "name": ParagraphStyle("name", fontName="Helvetica-Bold", fontSize=18, leading=22, textColor=ACCENT),
        "contact": ParagraphStyle("contact", fontName="Helvetica", fontSize=9, leading=12, textColor=MUTED),
        "meta": ParagraphStyle("meta", parent=body, spaceAfter=0),
        "subject": ParagraphStyle("subject", parent=body, fontName="Helvetica-Bold"),
        "body": body,
    }
    story = [
        Paragraph(escape(name), styles["name"]),
        Paragraph(escape(contact), styles["contact"]),
        Spacer(1, 3 * mm),
        HRFlowable(width="100%", thickness=0.8, color=ACCENT),
        Spacer(1, 5 * mm),
        Paragraph(escape(meta["date"]), styles["meta"]),
        Spacer(1, lead * 0.6),
        Paragraph(escape(meta.get("recipient") or "Hiring Team"), styles["meta"]),
        Paragraph(escape(meta["company"]), styles["meta"]),
        Spacer(1, lead * 0.6),
        Paragraph(escape(f"Re: {meta['role']}"), styles["subject"]),
    ]
    for block in blocks:
        # Single newlines inside a block are kept, so the closing can span lines.
        story.append(Paragraph(escape(block).replace("\n", "<br/>"), styles["body"]))
    doc = SimpleDocTemplate(
        str(out), pagesize=A4,
        leftMargin=22 * mm, rightMargin=22 * mm, topMargin=18 * mm, bottomMargin=18 * mm,
        title=f"Cover Letter - {meta['role']} - {meta['company']}", author=name,
    )
    doc.build(story)
    return len(PdfReader(str(out)).pages)


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("letter", type=Path)
    ap.add_argument("--resume", type=Path, default=DEFAULT_RESUME)
    args = ap.parse_args()
    if not args.letter.is_file():
        fail(1, f"letter not found: {args.letter}")
    if not args.resume.is_file():
        fail(1, f"resume not found: {args.resume}")

    meta, blocks = parse_letter(args.letter)
    check_alignment(args.letter, meta)
    name, contact = parse_contact(args.resume)
    out = args.letter.with_suffix(".pdf")
    for size in FONT_SIZES:
        pages = build(meta, blocks, name, contact, out, size)
        if pages == 1:
            words = len(" ".join(blocks).split())
            print(f"OK: {out} · pages=1 · font={size}pt · words={words}")
            return
    fail(2, f"{out}: still {pages} pages at {FONT_SIZES[-1]}pt, shorten the letter body")


if __name__ == "__main__":
    main()
