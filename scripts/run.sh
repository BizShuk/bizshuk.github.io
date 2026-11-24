#!/usr/bin/env bash
# run.sh —— 設定前置作業, 不安裝, 不啟動服務, 可重複執行 (idempotent).
# 由 package.json 的 `run:setup` 叫用.
set -euo pipefail

CONFIG_ROOT="${HOME}/.config/bizshuk.github.io"
REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

mkdir -p "${CONFIG_ROOT}/data" "${CONFIG_ROOT}/logs" "${REPO_ROOT}/tmp"

if [[ ! -e "${CONFIG_ROOT}/settings.json" ]]; then
    printf '{}\n' > "${CONFIG_ROOT}/settings.json"
    echo "created ${CONFIG_ROOT}/settings.json"
else
    echo "kept    ${CONFIG_ROOT}/settings.json"
fi

chmod 0700 "${CONFIG_ROOT}" "${CONFIG_ROOT}/data" "${CONFIG_ROOT}/logs"
chmod 0600 "${CONFIG_ROOT}/settings.json"

if [[ -L "${REPO_ROOT}/tmp/config" ]]; then
    echo "kept    symlink tmp/config -> $(readlink "${REPO_ROOT}/tmp/config")"
elif [[ ! -e "${REPO_ROOT}/tmp/config" ]]; then
    ln -s "${CONFIG_ROOT}" "${REPO_ROOT}/tmp/config"
    echo "created symlink tmp/config -> ${CONFIG_ROOT}"
else
    echo "warning: tmp/config exists and is not a symlink"
fi
