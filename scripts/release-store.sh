#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SITE_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
PROJECTS_ROOT="$(cd "${SITE_ROOT}/../.." && pwd)"
IPHONE_PROJECTS_ROOT="${IPHONE_PROJECTS_ROOT:-${PROJECTS_ROOT}/iphone}"
POLICY_VALIDATOR="${POLICY_VALIDATOR:-${IPHONE_PROJECTS_ROOT}/.agents/skills/policy/scripts/validate_policy_site.py}"

APPS=(
  "MinimalBrowser:MinimalBrowser"
  "iphone_sync:iphone_sync"
  "live-casting:live-casting"
  "md-viewer:md-viewer"
  "push:push"
  "tally:tally"
)

SELECTED_APPS=("${APPS[@]}")

if [[ $# -gt 1 ]]; then
  echo "Usage: npm run release:store -- [project]" >&2
  exit 1
fi

if [[ $# -eq 1 ]]; then
  SELECTED_APPS=()
  for entry in "${APPS[@]}"; do
    if [[ "${entry%%:*}" == "$1" ]]; then
      SELECTED_APPS+=("${entry}")
    fi
  done

  if [[ ${#SELECTED_APPS[@]} -eq 0 ]]; then
    echo "Unknown App Store project: $1" >&2
    exit 1
  fi
fi

if [[ ! -f "${POLICY_VALIDATOR}" ]]; then
  echo "Policy validator not found: ${POLICY_VALIDATOR}" >&2
  exit 1
fi

for entry in "${SELECTED_APPS[@]}"; do
  project="${entry%%:*}"
  source_root="${IPHONE_PROJECTS_ROOT}/${project}"
  python3 "${POLICY_VALIDATOR}" "${source_root}"
done

for entry in "${SELECTED_APPS[@]}"; do
  project="${entry%%:*}"
  route="${entry##*:}"
  source_dir="${IPHONE_PROJECTS_ROOT}/${project}/appstore"
  destination_dir="${SITE_ROOT}/pkg/${route}"

  case "${destination_dir}" in
    "${SITE_ROOT}/pkg/"*) ;;
    *)
      echo "Refusing destination outside pkg/: ${destination_dir}" >&2
      exit 1
      ;;
  esac

  mkdir -p "${destination_dir}"
  rsync -a --delete --exclude '.DS_Store' "${source_dir}/" "${destination_dir}/"
  echo "Released ${project} policy to pkg/${route}/"
done

echo "Store policy release completed for ${#SELECTED_APPS[@]} apps."
