# BYML CMS Recovery Storage

This directory is the repo-level recovery copy for CMS data and uploaded assets. MySQL is still the live database used by the site, but this folder is what makes the project portable if the database is lost, the server expires, or someone deletes content by mistake.

- `content/content.json` is the latest readable export of CMS tables. It includes draft, published, and soft-deleted rows, so deleted content can still be recovered from git history or from the current snapshot.
- `content/checkpoints/weekly/` and `content/checkpoints/monthly/` store long-lived readable JSON checkpoints. They are small and should be tracked by git. These are the main safety net when a problem is discovered weeks or months later.
- `uploads/` stores files uploaded from the admin media library. Keep this directory tracked because the project data volume is small.
- `backups/` stores timestamped SQL dumps created by `make backup`. These files are ignored to avoid repeated large git diffs, and old folders are pruned automatically.
- `public/` is outside this folder, but it is also tracked by git and contains the original site images, PDFs, PPT files, and videos that ship with the project.

Recommended routine:

1. Keep the Docker Compose `backup` service running on the server. It refreshes `storage/content/content.json` every 6 hours by default.
2. After important admin edits, use the admin Guide page's "刷新备份快照" button if you want the JSON snapshot updated immediately.
3. If git write access is configured on the server, use the admin Guide page's "一键同步到 Git" button, or start `docker compose --profile git-sync up -d git-sync`, to automatically commit and push `storage/content/content.json` plus new files under `storage/uploads/`.
4. If `git-sync` is not enabled, periodically commit `storage/content/content.json` and new files under `storage/uploads/` by hand.
5. Keep timestamped files under `storage/backups/` local unless you explicitly want an offline archive.
6. If you do not use the Compose backup service, run `make backup-cron-command` and add the printed line to crontab as an alternative.

Safety behavior:

- `backup` writes `storage/content/content.json` atomically. If MySQL is unavailable or the generated snapshot is empty/too small, it keeps the previous good snapshot.
- `backup` refuses to replace `content/content.json` if the new export is dramatically smaller than the previous one. The default threshold is `BACKUP_MIN_RELATIVE_CONTENT_PERCENT=30`, which catches many bad exports or unexpected data loss cases.
- `backup` creates weekly and monthly JSON checkpoints under `storage/content/checkpoints/`. An existing healthy checkpoint for the same week/month is not overwritten, so a bad backup later in the same period does not erase the earlier checkpoint.
- `backup` keeps only the newest `BACKUP_KEEP_COUNT` timestamp folders under `storage/backups/` by default. The Compose default is `28`, roughly 7 days at a 6-hour interval. Set `BACKUP_KEEP_COUNT=0` only if you intentionally want to keep everything.
- `backup` does not repeatedly archive `public/` and `storage/uploads/` by default, because those files are tracked directly by git. Set `BACKUP_INCLUDE_ASSET_ARCHIVES=true` if you want each timestamp backup to include asset tarballs too.
- `git-sync` refuses to push an empty or unexpectedly small `storage/content/content.json`.
- `git-sync` adds `storage/content/content.json`, new content checkpoints, and new/modified files under `storage/uploads/`, but it does not auto-stage upload deletions by default. Set `GIT_SYNC_ALLOW_UPLOAD_DELETES=true` only when you intentionally want remote git to record removed uploaded files.
- Even if a bad snapshot is pushed by mistake, git history can still recover an earlier version.

Local development to server:

1. Local MySQL data lives in your local Docker volume, not automatically in source control.
2. Run `make backup` locally, or use the admin Guide page's backup button, before moving content to the server.
3. Commit `storage/content/content.json` and any new files under `storage/uploads/`.
4. On the server, pull the project, copy `.env.prod.example` to `.env.prod`, and change the passwords/secrets.
5. Start production MySQL and API with `docker compose --env-file .env.prod -f compose.prod.yaml up -d --build mysql api`.
6. Run `make restore-content-prod-dry-run`, then `make restore-content-prod`, then `make media-import-prod`.
7. Start the full production stack with `docker compose --env-file .env.prod -f compose.prod.yaml up -d --build`.

Recovery routine:

1. Restore repo files from git.
2. For local development, start MySQL with `docker compose up -d mysql`, then run `make restore-content-dry-run` and `make restore-content`.
3. For production, start MySQL and API with `docker compose --env-file .env.prod -f compose.prod.yaml up -d --build mysql api`, then run `make restore-content-prod-dry-run` and `make restore-content-prod`.
4. Run `make media-import` locally or `make media-import-prod` in production to scan repo assets under `public/` back into the media library.
5. Start the remaining services.

Long-delayed recovery:

If a problem is discovered much later, inspect `storage/content/checkpoints/monthly/` or git history and restore a known-good checkpoint:

```bash
make restore-content-file-dry-run FILE=../storage/content/checkpoints/monthly/2026-05.json
make restore-content-file FILE=../storage/content/checkpoints/monthly/2026-05.json
make restore-content-file-prod-dry-run FILE=storage/content/checkpoints/monthly/2026-05.json
make restore-content-file-prod FILE=storage/content/checkpoints/monthly/2026-05.json
```
