#!/usr/bin/env sh
set -eu

ROOT_DIR="$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)"
cd "$ROOT_DIR"

GIT_SYNC_REMOTE_VALUE="${GIT_SYNC_REMOTE:-origin}"
GIT_SYNC_BRANCH_VALUE="${GIT_SYNC_BRANCH:-}"
GIT_SYNC_NAME_VALUE="${GIT_SYNC_NAME:-BYML Backup Bot}"
GIT_SYNC_EMAIL_VALUE="${GIT_SYNC_EMAIL:-byml-backup@example.local}"
GIT_SYNC_RUN_BACKUP_VALUE="${GIT_SYNC_RUN_BACKUP:-true}"
GIT_SYNC_DRY_RUN_VALUE="${GIT_SYNC_DRY_RUN:-false}"
GIT_SYNC_ALLOW_UPLOAD_DELETES_VALUE="${GIT_SYNC_ALLOW_UPLOAD_DELETES:-false}"
MIN_CONTENT_BYTES="${GIT_SYNC_MIN_CONTENT_BYTES:-2000}"

if ! command -v git >/dev/null 2>&1; then
  echo "git is not available; cannot sync CMS backup." >&2
  exit 1
fi

if [ "$GIT_SYNC_RUN_BACKUP_VALUE" = "true" ]; then
  ./scripts/backup.sh
fi

git_cmd() {
  git -c safe.directory="$ROOT_DIR" "$@"
}

if [ "$GIT_SYNC_DRY_RUN_VALUE" = "true" ]; then
  echo "CMS backup working tree changes:"
  git_cmd status --short -- storage/content/content.json storage/uploads
  exit 0
fi

if [ ! -s storage/content/content.json ]; then
  echo "storage/content/content.json is missing or empty; refusing to sync." >&2
  exit 1
fi

CONTENT_BYTES="$(wc -c < storage/content/content.json | tr -d ' ')"
if [ "$CONTENT_BYTES" -lt "$MIN_CONTENT_BYTES" ]; then
  echo "storage/content/content.json is unexpectedly small (${CONTENT_BYTES} bytes); refusing to sync." >&2
  exit 1
fi

git_cmd add -- storage/content/content.json
git_cmd add --ignore-removal -- storage/uploads

if [ "$GIT_SYNC_ALLOW_UPLOAD_DELETES_VALUE" = "true" ]; then
  git_cmd add -- storage/uploads
fi

if git_cmd diff --cached --name-status -- storage/uploads | grep -q '^D'; then
  if [ "$GIT_SYNC_ALLOW_UPLOAD_DELETES_VALUE" != "true" ]; then
    git_cmd restore --staged -- storage/uploads
    echo "Detected upload deletions; upload deletes are not auto-synced. Set GIT_SYNC_ALLOW_UPLOAD_DELETES=true to allow them." >&2
  fi
fi

if git_cmd diff --cached --quiet -- storage/content/content.json storage/uploads; then
  echo "No CMS backup changes to sync."
  exit 0
fi

if [ -z "$GIT_SYNC_BRANCH_VALUE" ]; then
  GIT_SYNC_BRANCH_VALUE="$(git_cmd branch --show-current 2>/dev/null || true)"
fi

if [ -z "$GIT_SYNC_BRANCH_VALUE" ]; then
  echo "Cannot determine git branch. Set GIT_SYNC_BRANCH before running git sync." >&2
  exit 1
fi

STAMP="$(date '+%Y-%m-%d %H:%M:%S %z')"
git_cmd -c user.name="$GIT_SYNC_NAME_VALUE" -c user.email="$GIT_SYNC_EMAIL_VALUE" commit -m "chore(cms): backup content ${STAMP}"
git_cmd push "$GIT_SYNC_REMOTE_VALUE" "HEAD:${GIT_SYNC_BRANCH_VALUE}"

echo "CMS backup synced to ${GIT_SYNC_REMOTE_VALUE}/${GIT_SYNC_BRANCH_VALUE}"
