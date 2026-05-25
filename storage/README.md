# BYML CMS Recovery Storage

This directory is the repo-level recovery copy for CMS data and uploaded assets. MySQL is still the live database used by the site, but this folder is what makes the project portable if the database is lost, the server expires, or someone deletes content by mistake.

- `content/content.json` is the latest readable export of CMS tables. It includes draft, published, and soft-deleted rows, so deleted content can still be recovered from git history or from the current snapshot.
- `uploads/` stores files uploaded from the admin media library. Keep this directory tracked because the project data volume is small.
- `backups/` stores timestamped SQL dumps and tar archives created by `make backup`. These files are ignored to avoid repeated large git diffs.
- `public/` is outside this folder, but it is also tracked by git and contains the original site images, PDFs, PPT files, and videos that ship with the project.

Recommended routine:

1. Keep the Docker Compose `backup` service running on the server. It refreshes `storage/content/content.json` every 6 hours by default.
2. If git write access is configured on the server, start `docker compose --profile git-sync up -d git-sync` to automatically commit and push `storage/content/content.json` plus new files under `storage/uploads/`. Keep the `backup` service running too, because `git-sync` only pushes the snapshot that `backup` has generated.
3. If `git-sync` is not enabled, periodically commit `storage/content/content.json` and new files under `storage/uploads/` by hand.
4. Keep timestamped files under `storage/backups/` local unless you explicitly want an offline archive.
5. If you do not use the Compose backup service, run `make backup-cron-command` and add the printed line to crontab as an alternative.

Safety behavior:

- `backup` writes `storage/content/content.json` atomically. If MySQL is unavailable or the generated snapshot is empty/too small, it keeps the previous good snapshot.
- `git-sync` refuses to push an empty or unexpectedly small `storage/content/content.json`.
- `git-sync` adds new and modified files under `storage/uploads/`, but it does not auto-stage upload deletions by default. Set `GIT_SYNC_ALLOW_UPLOAD_DELETES=true` only when you intentionally want remote git to record removed uploaded files.
- Even if a bad snapshot is pushed by mistake, git history can still recover an earlier version.

Recovery routine:

1. Restore repo files from git.
2. Start MySQL with `docker compose up -d mysql`.
3. Run `make restore-content-dry-run` to validate the JSON snapshot.
4. Run `make restore-content` to rebuild CMS tables from `storage/content/content.json`.
5. Run `make media-import` to scan repo assets under `public/` back into the media library.
