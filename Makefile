.PHONY: backup media-import

backup:
	./scripts/backup.sh

media-import:
	curl -sS -H 'Content-Type: application/json' \
		-d "{\"username\":\"$${ADMIN_USERNAME:-admin}\",\"password\":\"$${ADMIN_PASSWORD:-admin123456}\"}" \
		http://127.0.0.1:8080/api/admin/auth/login \
		| sed -n 's/.*"token":"\([^"]*\)".*/\1/p' \
		| xargs -I {} curl -sS -X POST -H "Authorization: Bearer {}" http://127.0.0.1:8080/api/admin/media/import-public
