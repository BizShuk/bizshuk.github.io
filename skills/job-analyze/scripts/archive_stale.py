#!/usr/bin/env python3
"""archive_stale.py - move JD files whose `created` date is older than N months into jd/archive/<company>/."""
import argparse
import calendar
import datetime
import os
import re
import subprocess
import sys

REPO = os.path.expanduser("~/projects/product/bizshuk.github.io")
JD_DIR = os.path.join(REPO, "pkg", "resume", "jd")
DATE_KEYS = ("created", "fetched")
ACTIVE_STATUSES = {"applied", "screening", "interviewing", "offer"}


def months_ago(today, months):
    month = today.month - months
    year = today.year
    while month <= 0:
        month += 12
        year -= 1
    day = min(today.day, calendar.monthrange(year, month)[1])
    return datetime.date(year, month, day)


def front_matter_date(path):
    with open(path, encoding="utf-8") as f:
        text = f.read()
    match = re.match(r"^---\n(.*?)\n---\n", text, re.S)
    if not match:
        return None, None, None
    fields = dict(re.findall(r"^(\w+):\s*(.*)$", match.group(1), re.M))
    status = fields.get("status", "").strip().strip('"')
    for key in DATE_KEYS:
        value = fields.get(key, "").strip().strip('"')
        if re.fullmatch(r"\d{4}-\d{2}-\d{2}", value):
            return key, datetime.date.fromisoformat(value), status
    return None, None, status


def is_tracked(path):
    result = subprocess.run(
        ["git", "-C", REPO, "ls-files", "--error-unmatch", path],
        capture_output=True,
    )
    return result.returncode == 0


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--months", type=int, default=3)
    parser.add_argument("--dry-run", action="store_true")
    args = parser.parse_args()

    if not os.path.isdir(REPO):
        print(f"ERROR: repo not found: {REPO}", file=sys.stderr)
        return 1

    cutoff = months_ago(datetime.date.today(), args.months)
    print(f"cutoff: {cutoff} (created before this date is archived)")

    for company in sorted(os.listdir(JD_DIR)):
        company_dir = os.path.join(JD_DIR, company)
        if company == "archive" or not os.path.isdir(company_dir):
            continue
        for name in sorted(os.listdir(company_dir)):
            if not name.endswith(".md"):
                continue
            src = os.path.join(company_dir, name)
            key, date, status = front_matter_date(src)
            if date is None:
                print(f"SKIP no-date: {src}")
                continue
            if date >= cutoff:
                continue
            if status in ACTIVE_STATUSES:
                print(f"SKIP active status={status}: {src}")
                continue
            dst = os.path.join(JD_DIR, "archive", company, name)
            if os.path.exists(dst):
                print(f"SKIP exists: {dst}")
                continue
            print(f"ARCHIVE {key}={date}: {src} -> {dst}")
            if args.dry_run:
                continue
            os.makedirs(os.path.dirname(dst), exist_ok=True)
            if is_tracked(src):
                subprocess.run(["git", "-C", REPO, "mv", src, dst], check=True)
            else:
                os.rename(src, dst)
        if not args.dry_run and os.path.isdir(company_dir) and not os.listdir(company_dir):
            os.rmdir(company_dir)
    return 0


if __name__ == "__main__":
    sys.exit(main())
