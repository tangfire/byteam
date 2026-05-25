.PHONY: backup content-snapshot restore-content restore-content-dry-run backup-cron-command media-import

backup:
	./scripts/backup.sh

content-snapshot:
	./scripts/backup.sh

restore-content:
	cd server && GOCACHE=$(CURDIR)/.cache/go-build go run ./cmd/restore -file ../storage/content/content.json

restore-content-dry-run:
	cd server && GOCACHE=$(CURDIR)/.cache/go-build go run ./cmd/restore -file ../storage/content/content.json -dry-run

backup-cron-command:
	@echo "0 */6 * * * cd $(CURDIR) && make backup >> storage/backups/backup.log 2>&1"

media-import:
	curl -sS -H 'Content-Type: application/json' \
		-d "{\"username\":\"$${ADMIN_USERNAME:-admin}\",\"password\":\"$${ADMIN_PASSWORD:-admin123456}\"}" \
		http://127.0.0.1:8080/api/admin/auth/login \
		| sed -n 's/.*"token":"\([^"]*\)".*/\1/p' \
		| xargs -I {} curl -sS -X POST -H "Authorization: Bearer {}" http://127.0.0.1:8080/api/admin/media/import-public
