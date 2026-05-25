#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
STAMP="$(date +%Y%m%d-%H%M%S)"
BACKUP_DIR="${ROOT_DIR}/storage/backups/${STAMP}"
UPLOADS_DIR="${ROOT_DIR}/storage/uploads"
CONTENT_DIR="${ROOT_DIR}/storage/content"
CONTENT_SNAPSHOT="${CONTENT_DIR}/content.json"
CONTENT_TMP="${CONTENT_SNAPSHOT}.tmp"
MIN_CONTENT_BYTES="${BACKUP_MIN_CONTENT_BYTES:-2000}"
MIN_RELATIVE_CONTENT_PERCENT="${BACKUP_MIN_RELATIVE_CONTENT_PERCENT:-30}"
BACKUP_KEEP_COUNT="${BACKUP_KEEP_COUNT:-28}"
BACKUP_INCLUDE_ASSET_ARCHIVES="${BACKUP_INCLUDE_ASSET_ARCHIVES:-false}"
CONTENT_CHECKPOINT_DIR="${CONTENT_DIR}/checkpoints"
MYSQL_HOST_VALUE="${MYSQL_HOST:-}"
MYSQL_PORT_VALUE="${MYSQL_PORT:-3306}"
MYSQL_DATABASE_VALUE="${MYSQL_DATABASE:-byml}"
MYSQL_USER_VALUE="${MYSQL_USER:-byml}"
MYSQL_PASSWORD_VALUE="${MYSQL_PASSWORD:-byml_password}"
MYSQL_ROOT_USER_VALUE="${MYSQL_ROOT_USER:-root}"
MYSQL_ROOT_PASSWORD_VALUE="${MYSQL_ROOT_PASSWORD:-root_password}"
MYSQL_CLI="${MYSQL_CLI:-mysql}"
MYSQLDUMP_CLI="${MYSQLDUMP_CLI:-mysqldump}"
DB_BACKUP_CREATED="false"
CONTENT_SNAPSHOT_CREATED="false"
CONTENT_SNAPSHOT_REJECTED="false"

mkdir -p "${BACKUP_DIR}"
mkdir -p "${CONTENT_DIR}"
mkdir -p "${CONTENT_CHECKPOINT_DIR}/weekly" "${CONTENT_CHECKPOINT_DIR}/monthly"
trap 'rm -f "${CONTENT_TMP}"' EXIT

if ! command -v "${MYSQL_CLI}" >/dev/null 2>&1 && command -v mariadb >/dev/null 2>&1; then
  MYSQL_CLI="mariadb"
fi

if ! command -v "${MYSQLDUMP_CLI}" >/dev/null 2>&1 && command -v mariadb-dump >/dev/null 2>&1; then
  MYSQLDUMP_CLI="mariadb-dump"
fi

echo "Creating BYML backup at ${BACKUP_DIR}"
echo "Refreshing git-visible CMS snapshot at ${CONTENT_SNAPSHOT}"

has_direct_mysql() {
  [ -n "${MYSQL_HOST_VALUE}" ] && command -v "${MYSQL_CLI}" >/dev/null 2>&1 && command -v "${MYSQLDUMP_CLI}" >/dev/null 2>&1
}

has_compose_mysql() {
  [ -n "$(docker compose -f "${ROOT_DIR}/compose.yaml" ps --status running -q mysql 2>/dev/null)" ]
}

install_content_snapshot() {
  if [ ! -s "${CONTENT_TMP}" ]; then
    echo "Generated content snapshot is empty; keeping previous ${CONTENT_SNAPSHOT}." >&2
    return 1
  fi

  content_bytes="$(wc -c < "${CONTENT_TMP}" | tr -d ' ')"
  if [ "${content_bytes}" -lt "${MIN_CONTENT_BYTES}" ]; then
    echo "Generated content snapshot is unexpectedly small (${content_bytes} bytes); keeping previous ${CONTENT_SNAPSHOT}." >&2
    return 1
  fi

  if [ -s "${CONTENT_SNAPSHOT}" ] && [[ "${MIN_RELATIVE_CONTENT_PERCENT}" =~ ^[0-9]+$ ]] && [ "${MIN_RELATIVE_CONTENT_PERCENT}" -gt 0 ]; then
    previous_bytes="$(wc -c < "${CONTENT_SNAPSHOT}" | tr -d ' ')"
    minimum_relative_bytes=$((previous_bytes * MIN_RELATIVE_CONTENT_PERCENT / 100))
    if [ "${minimum_relative_bytes}" -gt "${MIN_CONTENT_BYTES}" ] && [ "${content_bytes}" -lt "${minimum_relative_bytes}" ]; then
      echo "Generated content snapshot shrank from ${previous_bytes} to ${content_bytes} bytes; keeping previous ${CONTENT_SNAPSHOT}." >&2
      return 1
    fi
  fi

  create_content_checkpoints
  mv "${CONTENT_TMP}" "${CONTENT_SNAPSHOT}"
  cp "${CONTENT_SNAPSHOT}" "${BACKUP_DIR}/content.json"
  CONTENT_SNAPSHOT_CREATED="true"
}

create_content_checkpoints() {
  weekly_checkpoint="${CONTENT_CHECKPOINT_DIR}/weekly/$(date +%Y-W%U).json"
  monthly_checkpoint="${CONTENT_CHECKPOINT_DIR}/monthly/$(date +%Y-%m).json"
  install_content_checkpoint "${weekly_checkpoint}" "weekly"
  install_content_checkpoint "${monthly_checkpoint}" "monthly"
}

install_content_checkpoint() {
  checkpoint_path="$1"
  checkpoint_label="$2"
  if [ -s "${checkpoint_path}" ]; then
    checkpoint_bytes="$(wc -c < "${checkpoint_path}" | tr -d ' ')"
    if [ "${checkpoint_bytes}" -ge "${MIN_CONTENT_BYTES}" ]; then
      echo "Keeping existing ${checkpoint_label} checkpoint ${checkpoint_path}"
      return
    fi
    echo "Replacing too-small ${checkpoint_label} checkpoint ${checkpoint_path}"
  fi
  cp "${CONTENT_TMP}" "${checkpoint_path}"
  echo "Created ${checkpoint_label} checkpoint ${checkpoint_path}"
}

cleanup_old_backups() {
  if ! [[ "${BACKUP_KEEP_COUNT}" =~ ^[0-9]+$ ]]; then
    echo "BACKUP_KEEP_COUNT=${BACKUP_KEEP_COUNT} is not a number; skipped pruning old backups." >&2
    return
  fi

  if [ "${BACKUP_KEEP_COUNT}" -eq 0 ]; then
    echo "Backup pruning disabled because BACKUP_KEEP_COUNT=0."
    return
  fi

  backup_total="$(find "${ROOT_DIR}/storage/backups" -mindepth 1 -maxdepth 1 -type d -name '????????-??????' | wc -l | tr -d ' ')"
  if [ "${backup_total}" -le "${BACKUP_KEEP_COUNT}" ]; then
    return
  fi

  remove_count=$((backup_total - BACKUP_KEEP_COUNT))
  find "${ROOT_DIR}/storage/backups" -mindepth 1 -maxdepth 1 -type d -name '????????-??????' \
    | sort \
    | head -n "${remove_count}" \
    | while IFS= read -r old_backup; do
        echo "Pruning old backup ${old_backup}"
        rm -rf -- "${old_backup}"
      done
}

CONTENT_SQL="$(cat <<'SQL'
SELECT JSON_PRETTY(JSON_OBJECT(
  'news', COALESCE((
    SELECT JSON_ARRAYAGG(item)
    FROM (
      SELECT JSON_OBJECT(
        'id', id,
        'title', title,
        'content', content,
        'excerpt', excerpt,
        'type', type,
        'typeLabel', type_label,
        'date', event_date,
        'color', color,
        'status', status,
        'sortOrder', sort_order,
        'createdAt', DATE_FORMAT(created_at, '%Y-%m-%dT%H:%i:%s'),
        'updatedAt', DATE_FORMAT(updated_at, '%Y-%m-%dT%H:%i:%s'),
        'deletedAt', DATE_FORMAT(deleted_at, '%Y-%m-%dT%H:%i:%s')
      ) AS item
      FROM news_items
      ORDER BY sort_order ASC, event_date DESC, id ASC
    ) ordered_news
  ), JSON_ARRAY()),
  'people', COALESCE((
    SELECT JSON_ARRAYAGG(item)
    FROM (
      SELECT JSON_OBJECT(
        'id', id,
        'name', name,
        'avatarUrl', avatar_url,
        'category', category,
        'research', research,
        'graduationDate', graduation_date,
        'status', status,
        'sortOrder', sort_order,
        'createdAt', DATE_FORMAT(created_at, '%Y-%m-%dT%H:%i:%s'),
        'updatedAt', DATE_FORMAT(updated_at, '%Y-%m-%dT%H:%i:%s'),
        'deletedAt', DATE_FORMAT(deleted_at, '%Y-%m-%dT%H:%i:%s')
      ) AS item
      FROM people
      ORDER BY sort_order ASC, category ASC, id ASC
    ) ordered_people
  ), JSON_ARRAY()),
  'undergraduates', COALESCE((
    SELECT JSON_ARRAYAGG(item)
    FROM (
      SELECT JSON_OBJECT(
        'id', id,
        'name', name,
        'grade', grade,
        'major', major,
        'direction', direction,
        'achievements', achievements,
        'status', status,
        'sortOrder', sort_order,
        'createdAt', DATE_FORMAT(created_at, '%Y-%m-%dT%H:%i:%s'),
        'updatedAt', DATE_FORMAT(updated_at, '%Y-%m-%dT%H:%i:%s'),
        'deletedAt', DATE_FORMAT(deleted_at, '%Y-%m-%dT%H:%i:%s')
      ) AS item
      FROM undergraduate_educations
      ORDER BY sort_order ASC, grade DESC, id ASC
    ) ordered_undergraduates
  ), JSON_ARRAY()),
  'publications', COALESCE((
    SELECT JSON_ARRAYAGG(item)
    FROM (
      SELECT JSON_OBJECT(
        'id', id,
        'image', image_url,
        'title', title,
        'authors', authors,
        'venue', venue,
        'year', year,
        'kind', kind,
        'status', status,
        'featured', featured,
        'sortOrder', sort_order,
        'createdAt', DATE_FORMAT(created_at, '%Y-%m-%dT%H:%i:%s'),
        'updatedAt', DATE_FORMAT(updated_at, '%Y-%m-%dT%H:%i:%s'),
        'deletedAt', DATE_FORMAT(deleted_at, '%Y-%m-%dT%H:%i:%s')
      ) AS item
      FROM publications
      ORDER BY sort_order ASC, year DESC, id ASC
    ) ordered_publications
  ), JSON_ARRAY()),
  'publicationLinks', COALESCE((
    SELECT JSON_ARRAYAGG(item)
    FROM (
      SELECT JSON_OBJECT(
        'id', id,
        'publicationId', publication_id,
        'type', type,
        'label', label,
        'url', url,
        'routeName', route_name,
        'sortOrder', sort_order,
        'createdAt', DATE_FORMAT(created_at, '%Y-%m-%dT%H:%i:%s'),
        'updatedAt', DATE_FORMAT(updated_at, '%Y-%m-%dT%H:%i:%s')
      ) AS item
      FROM publication_links
      ORDER BY publication_id ASC, sort_order ASC, id ASC
    ) ordered_publication_links
  ), JSON_ARRAY()),
  'patents', COALESCE((
    SELECT JSON_ARRAYAGG(item)
    FROM (
      SELECT JSON_OBJECT(
        'id', id,
        'authors', authors,
        'title', title,
        'date', date,
        'country', country,
        'number', number,
        'category', category,
        'status', status,
        'sortOrder', sort_order,
        'createdAt', DATE_FORMAT(created_at, '%Y-%m-%dT%H:%i:%s'),
        'updatedAt', DATE_FORMAT(updated_at, '%Y-%m-%dT%H:%i:%s'),
        'deletedAt', DATE_FORMAT(deleted_at, '%Y-%m-%dT%H:%i:%s')
      ) AS item
      FROM patents
      ORDER BY sort_order ASC, date DESC, id ASC
    ) ordered_patents
  ), JSON_ARRAY()),
  'researchProjects', COALESCE((
    SELECT JSON_ARRAYAGG(item)
    FROM (
      SELECT JSON_OBJECT(
        'id', id,
        'title', title,
        'fund', fund,
        'number', number,
        'period', period,
        'amount', amount,
        'projectStatus', project_status,
        'role', role,
        'status', status,
        'sortOrder', sort_order,
        'createdAt', DATE_FORMAT(created_at, '%Y-%m-%dT%H:%i:%s'),
        'updatedAt', DATE_FORMAT(updated_at, '%Y-%m-%dT%H:%i:%s'),
        'deletedAt', DATE_FORMAT(deleted_at, '%Y-%m-%dT%H:%i:%s')
      ) AS item
      FROM research_projects
      ORDER BY sort_order ASC, id ASC
    ) ordered_research_projects
  ), JSON_ARRAY()),
  'media', COALESCE((
    SELECT JSON_ARRAYAGG(item)
    FROM (
      SELECT JSON_OBJECT(
        'id', id,
        'fileName', file_name,
        'originalName', original_name,
        'url', url,
        'mimeType', mime_type,
        'size', size,
        'kind', kind,
        'createdAt', DATE_FORMAT(created_at, '%Y-%m-%dT%H:%i:%s'),
        'updatedAt', DATE_FORMAT(updated_at, '%Y-%m-%dT%H:%i:%s'),
        'deletedAt', DATE_FORMAT(deleted_at, '%Y-%m-%dT%H:%i:%s')
      ) AS item
      FROM media_assets
      ORDER BY kind ASC, url ASC, id ASC
    ) ordered_media
  ), JSON_ARRAY())
));
SQL
)"

if has_direct_mysql; then
  "${MYSQLDUMP_CLI}" \
    -h"${MYSQL_HOST_VALUE}" \
    -P"${MYSQL_PORT_VALUE}" \
    -u"${MYSQL_ROOT_USER_VALUE}" \
    -p"${MYSQL_ROOT_PASSWORD_VALUE}" \
    --default-character-set=utf8mb4 \
    --databases "${MYSQL_DATABASE_VALUE}" \
    | gzip > "${BACKUP_DIR}/byml.sql.gz"
  DB_BACKUP_CREATED="true"
elif has_compose_mysql; then
  docker compose -f "${ROOT_DIR}/compose.yaml" exec -T mysql \
    mysqldump -uroot -proot_password --default-character-set=utf8mb4 --databases byml \
    | gzip > "${BACKUP_DIR}/byml.sql.gz"
  DB_BACKUP_CREATED="true"
else
  echo "MySQL container is not running; skipped SQL dump." >&2
fi

if has_direct_mysql; then
  "${MYSQL_CLI}" \
    -h"${MYSQL_HOST_VALUE}" \
    -P"${MYSQL_PORT_VALUE}" \
    -u"${MYSQL_USER_VALUE}" \
    -p"${MYSQL_PASSWORD_VALUE}" \
    --default-character-set=utf8mb4 \
    "${MYSQL_DATABASE_VALUE}" \
    --batch --raw --skip-column-names \
    -e "${CONTENT_SQL}" \
    > "${CONTENT_TMP}"
  if ! install_content_snapshot; then
    CONTENT_SNAPSHOT_REJECTED="true"
  fi
elif has_compose_mysql; then
  docker compose -f "${ROOT_DIR}/compose.yaml" exec -T mysql \
    mysql -ubyml -pbyml_password --default-character-set=utf8mb4 byml --batch --raw --skip-column-names \
    -e "${CONTENT_SQL}" \
    > "${CONTENT_TMP}"
  if ! install_content_snapshot; then
    CONTENT_SNAPSHOT_REJECTED="true"
  fi
else
  echo "MySQL container is not running; kept existing ${CONTENT_SNAPSHOT}." >&2
fi

if [ "${BACKUP_INCLUDE_ASSET_ARCHIVES}" = "true" ]; then
  if [ -d "${UPLOADS_DIR}" ]; then
    tar -czf "${BACKUP_DIR}/uploads.tar.gz" -C "${ROOT_DIR}/storage" uploads
  else
    mkdir -p "${UPLOADS_DIR}"
    tar -czf "${BACKUP_DIR}/uploads.tar.gz" -C "${ROOT_DIR}/storage" uploads
  fi

  tar -czf "${BACKUP_DIR}/public.tar.gz" -C "${ROOT_DIR}" public
else
  echo "Skipped repeated asset archives. Repo-tracked storage/uploads and public keep file recovery copies."
fi

cat > "${BACKUP_DIR}/README.txt" <<EOF
BYML backup created at ${STAMP}

Files:
- byml.sql.gz: full MySQL dump for the byml database
- content.json: JSON export of CMS content tables
- uploads.tar.gz: optional, only when BACKUP_INCLUDE_ASSET_ARCHIVES=true
- public.tar.gz: optional, only when BACKUP_INCLUDE_ASSET_ARCHIVES=true

The latest git-visible recovery snapshot is also written to:
${CONTENT_SNAPSHOT}
EOF

if [ "${DB_BACKUP_CREATED}" != "true" ] && [ "${CONTENT_SNAPSHOT_CREATED}" != "true" ]; then
  echo "No fresh database backup was created; removing incomplete timestamp backup." >&2
  rm -rf -- "${BACKUP_DIR}"
  exit 1
fi

cleanup_old_backups

if [ "${CONTENT_SNAPSHOT_REJECTED}" = "true" ]; then
  echo "Backup SQL dump was created, but the content snapshot was rejected. Kept previous ${CONTENT_SNAPSHOT}." >&2
  exit 1
fi

echo "Backup complete: ${BACKUP_DIR}"
