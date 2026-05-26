.PHONY: backup content-snapshot restore-content restore-content-dry-run restore-content-file restore-content-file-dry-run restore-content-prod restore-content-prod-dry-run restore-content-file-prod restore-content-file-prod-dry-run backup-cron-command git-sync-backup git-sync-backup-dry-run media-import media-import-prod

PROD_COMPOSE ?= docker compose --env-file .env.prod -f compose.prod.yaml
ADMIN_BASE_URL ?= http://127.0.0.1:8080
ADMIN_PROD_BASE_URL ?= http://127.0.0.1

backup:
	./scripts/backup.sh

content-snapshot:
	./scripts/backup.sh

restore-content:
	cd server && GOCACHE=$(CURDIR)/.cache/go-build go run ./cmd/restore -file ../storage/content/content.json

restore-content-dry-run:
	cd server && GOCACHE=$(CURDIR)/.cache/go-build go run ./cmd/restore -file ../storage/content/content.json -dry-run

restore-content-file:
	@test -n "$(FILE)" || (echo "Usage: make restore-content-file FILE=../storage/content/checkpoints/monthly/2026-05.json"; exit 1)
	cd server && GOCACHE=$(CURDIR)/.cache/go-build go run ./cmd/restore -file "$(FILE)"

restore-content-file-dry-run:
	@test -n "$(FILE)" || (echo "Usage: make restore-content-file-dry-run FILE=../storage/content/checkpoints/monthly/2026-05.json"; exit 1)
	cd server && GOCACHE=$(CURDIR)/.cache/go-build go run ./cmd/restore -file "$(FILE)" -dry-run

restore-content-prod:
	$(PROD_COMPOSE) exec -T api /app/byml-restore -file /workspace/storage/content/content.json

restore-content-prod-dry-run:
	$(PROD_COMPOSE) exec -T api /app/byml-restore -file /workspace/storage/content/content.json -dry-run

restore-content-file-prod:
	@test -n "$(FILE)" || (echo "Usage: make restore-content-file-prod FILE=storage/content/checkpoints/monthly/2026-05.json"; exit 1)
	$(PROD_COMPOSE) exec -T api /app/byml-restore -file "/workspace/$(FILE)"

restore-content-file-prod-dry-run:
	@test -n "$(FILE)" || (echo "Usage: make restore-content-file-prod-dry-run FILE=storage/content/checkpoints/monthly/2026-05.json"; exit 1)
	$(PROD_COMPOSE) exec -T api /app/byml-restore -file "/workspace/$(FILE)" -dry-run

backup-cron-command:
	@echo "0 */6 * * * cd $(CURDIR) && BACKUP_KEEP_COUNT=28 make backup >> storage/backups/backup.log 2>&1"

git-sync-backup:
	./scripts/git-sync-backup.sh

git-sync-backup-dry-run:
	GIT_SYNC_DRY_RUN=true ./scripts/git-sync-backup.sh

media-import:
	curl -sS -H 'Content-Type: application/json' \
		-d "{\"username\":\"$${ADMIN_USERNAME:-admin}\",\"password\":\"$${ADMIN_PASSWORD:-admin123456}\"}" \
		$(ADMIN_BASE_URL)/api/admin/auth/login \
		| sed -n 's/.*"token":"\([^"]*\)".*/\1/p' \
		| xargs -I {} curl -sS -X POST -H "Authorization: Bearer {}" $(ADMIN_BASE_URL)/api/admin/media/import-public

media-import-prod:
	$(PROD_COMPOSE) exec -T api sh -lc 'set -e; token=$$(curl -sS -H "Content-Type: application/json" -d "{\"username\":\"$${ADMIN_USERNAME}\",\"password\":\"$${ADMIN_PASSWORD}\"}" http://127.0.0.1:8080/api/admin/auth/login | sed -n "s/.*\"token\":\"\([^\"]*\)\".*/\1/p"); test -n "$$token"; curl -sS -X POST -H "Authorization: Bearer $$token" http://127.0.0.1:8080/api/admin/media/import-public'
