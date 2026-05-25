#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
STAMP="$(date +%Y%m%d-%H%M%S)"
BACKUP_DIR="${ROOT_DIR}/storage/backups/${STAMP}"
UPLOADS_DIR="${ROOT_DIR}/storage/uploads"

mkdir -p "${BACKUP_DIR}"

echo "Creating BYML backup at ${BACKUP_DIR}"

if docker compose -f "${ROOT_DIR}/compose.yaml" ps mysql >/dev/null 2>&1; then
  docker compose -f "${ROOT_DIR}/compose.yaml" exec -T mysql \
    mysqldump -uroot -proot_password --databases byml \
    | gzip > "${BACKUP_DIR}/byml.sql.gz"
else
  echo "MySQL container is not running; skipped SQL dump." >&2
fi

if docker compose -f "${ROOT_DIR}/compose.yaml" ps mysql >/dev/null 2>&1; then
  docker compose -f "${ROOT_DIR}/compose.yaml" exec -T mysql \
    mysql -ubyml -pbyml_password byml --batch --raw --skip-column-names \
    -e "SELECT JSON_OBJECT('news',(SELECT JSON_ARRAYAGG(JSON_OBJECT('id',id,'title',title,'content',content,'excerpt',excerpt,'type',type,'typeLabel',type_label,'date',event_date,'color',color,'status',status,'sortOrder',sort_order)) FROM news_items WHERE deleted_at IS NULL),'people',(SELECT JSON_ARRAYAGG(JSON_OBJECT('id',id,'name',name,'avatarUrl',avatar_url,'category',category,'research',research,'graduationDate',graduation_date,'status',status,'sortOrder',sort_order)) FROM people WHERE deleted_at IS NULL),'undergraduates',(SELECT JSON_ARRAYAGG(JSON_OBJECT('id',id,'name',name,'grade',grade,'major',major,'direction',direction,'achievements',achievements,'status',status,'sortOrder',sort_order)) FROM undergraduate_educations WHERE deleted_at IS NULL),'publications',(SELECT JSON_ARRAYAGG(JSON_OBJECT('id',id,'image',image_url,'title',title,'authors',authors,'venue',venue,'year',year,'kind',kind,'status',status,'featured',featured,'sortOrder',sort_order)) FROM publications WHERE deleted_at IS NULL),'publicationLinks',(SELECT JSON_ARRAYAGG(JSON_OBJECT('id',id,'publicationId',publication_id,'type',type,'label',label,'url',url,'routeName',route_name,'sortOrder',sort_order)) FROM publication_links),'patents',(SELECT JSON_ARRAYAGG(JSON_OBJECT('id',id,'authors',authors,'title',title,'date',date,'country',country,'number',number,'category',category,'status',status,'sortOrder',sort_order)) FROM patents WHERE deleted_at IS NULL),'researchProjects',(SELECT JSON_ARRAYAGG(JSON_OBJECT('id',id,'title',title,'fund',fund,'number',number,'period',period,'amount',amount,'projectStatus',project_status,'role',role,'status',status,'sortOrder',sort_order)) FROM research_projects WHERE deleted_at IS NULL),'media',(SELECT JSON_ARRAYAGG(JSON_OBJECT('id',id,'fileName',file_name,'originalName',original_name,'url',url,'mimeType',mime_type,'size',size,'kind',kind)) FROM media_assets WHERE deleted_at IS NULL))" \
    > "${BACKUP_DIR}/content.json"
else
  echo "{}" > "${BACKUP_DIR}/content.json"
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
