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
MYSQL_HOST_VALUE="${MYSQL_HOST:-}"
MYSQL_PORT_VALUE="${MYSQL_PORT:-3306}"
MYSQL_DATABASE_VALUE="${MYSQL_DATABASE:-byml}"
MYSQL_USER_VALUE="${MYSQL_USER:-byml}"
MYSQL_PASSWORD_VALUE="${MYSQL_PASSWORD:-byml_password}"
MYSQL_ROOT_USER_VALUE="${MYSQL_ROOT_USER:-root}"
MYSQL_ROOT_PASSWORD_VALUE="${MYSQL_ROOT_PASSWORD:-root_password}"

mkdir -p "${BACKUP_DIR}"
mkdir -p "${CONTENT_DIR}"
trap 'rm -f "${CONTENT_TMP}"' EXIT

echo "Creating BYML backup at ${BACKUP_DIR}"
echo "Refreshing git-visible CMS snapshot at ${CONTENT_SNAPSHOT}"

has_direct_mysql() {
  [ -n "${MYSQL_HOST_VALUE}" ] && command -v mysql >/dev/null 2>&1 && command -v mysqldump >/dev/null 2>&1
}

has_compose_mysql() {
  docker compose -f "${ROOT_DIR}/compose.yaml" ps mysql >/dev/null 2>&1
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

  mv "${CONTENT_TMP}" "${CONTENT_SNAPSHOT}"
  cp "${CONTENT_SNAPSHOT}" "${BACKUP_DIR}/content.json"
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
  mysqldump \
    -h"${MYSQL_HOST_VALUE}" \
    -P"${MYSQL_PORT_VALUE}" \
    -u"${MYSQL_ROOT_USER_VALUE}" \
    -p"${MYSQL_ROOT_PASSWORD_VALUE}" \
    --default-character-set=utf8mb4 \
    --databases "${MYSQL_DATABASE_VALUE}" \
    | gzip > "${BACKUP_DIR}/byml.sql.gz"
elif has_compose_mysql; then
  docker compose -f "${ROOT_DIR}/compose.yaml" exec -T mysql \
    mysqldump -uroot -proot_password --default-character-set=utf8mb4 --databases byml \
    | gzip > "${BACKUP_DIR}/byml.sql.gz"
else
  echo "MySQL container is not running; skipped SQL dump." >&2
fi

if has_direct_mysql; then
  mysql \
    -h"${MYSQL_HOST_VALUE}" \
    -P"${MYSQL_PORT_VALUE}" \
    -u"${MYSQL_USER_VALUE}" \
    -p"${MYSQL_PASSWORD_VALUE}" \
    --default-character-set=utf8mb4 \
    "${MYSQL_DATABASE_VALUE}" \
    --batch --raw --skip-column-names \
    -e "${CONTENT_SQL}" \
    > "${CONTENT_TMP}"
  install_content_snapshot
elif has_compose_mysql; then
  docker compose -f "${ROOT_DIR}/compose.yaml" exec -T mysql \
    mysql -ubyml -pbyml_password --default-character-set=utf8mb4 byml --batch --raw --skip-column-names \
    -e "${CONTENT_SQL}" \
    > "${CONTENT_TMP}"
  install_content_snapshot
else
  echo "MySQL container is not running; kept existing ${CONTENT_SNAPSHOT}." >&2
  if [ -s "${CONTENT_SNAPSHOT}" ]; then
    cp "${CONTENT_SNAPSHOT}" "${BACKUP_DIR}/content.json"
  fi
fi

if [ -d "${UPLOADS_DIR}" ]; then
  tar -czf "${BACKUP_DIR}/uploads.tar.gz" -C "${ROOT_DIR}/storage" uploads
else
  mkdir -p "${UPLOADS_DIR}"
  tar -czf "${BACKUP_DIR}/uploads.tar.gz" -C "${ROOT_DIR}/storage" uploads
fi

tar -czf "${BACKUP_DIR}/public.tar.gz" -C "${ROOT_DIR}" public

cat > "${BACKUP_DIR}/README.txt" <<EOF
BYML backup created at ${STAMP}

Files:
- byml.sql.gz: full MySQL dump for the byml database
- content.json: JSON export of CMS content tables
- uploads.tar.gz: uploaded files from storage/uploads
- public.tar.gz: versioned public static assets currently deployed

The latest git-visible recovery snapshot is also written to:
${CONTENT_SNAPSHOT}
EOF

echo "Backup complete: ${BACKUP_DIR}"
