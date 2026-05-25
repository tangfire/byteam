#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
STAMP="$(date +%Y%m%d-%H%M%S)"
BACKUP_DIR="${ROOT_DIR}/storage/backups/${STAMP}"
UPLOADS_DIR="${ROOT_DIR}/storage/uploads"
CONTENT_DIR="${ROOT_DIR}/storage/content"
CONTENT_SNAPSHOT="${CONTENT_DIR}/content.json"

mkdir -p "${BACKUP_DIR}"
mkdir -p "${CONTENT_DIR}"

echo "Creating BYML backup at ${BACKUP_DIR}"

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

if docker compose -f "${ROOT_DIR}/compose.yaml" ps mysql >/dev/null 2>&1; then
  docker compose -f "${ROOT_DIR}/compose.yaml" exec -T mysql \
    mysqldump -uroot -proot_password --default-character-set=utf8mb4 --databases byml \
    | gzip > "${BACKUP_DIR}/byml.sql.gz"
else
  echo "MySQL container is not running; skipped SQL dump." >&2
fi

if docker compose -f "${ROOT_DIR}/compose.yaml" ps mysql >/dev/null 2>&1; then
  docker compose -f "${ROOT_DIR}/compose.yaml" exec -T mysql \
    mysql -ubyml -pbyml_password --default-character-set=utf8mb4 byml --batch --raw --skip-column-names \
    -e "${CONTENT_SQL}" \
    > "${CONTENT_SNAPSHOT}"
  cp "${CONTENT_SNAPSHOT}" "${BACKUP_DIR}/content.json"
else
  echo "{}" > "${CONTENT_SNAPSHOT}"
  cp "${CONTENT_SNAPSHOT}" "${BACKUP_DIR}/content.json"
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
EOF

echo "Backup complete: ${BACKUP_DIR}"
