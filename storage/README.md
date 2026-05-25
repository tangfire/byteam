# BYML CMS Recovery Storage

This directory is the repo-level recovery copy for CMS data and uploaded assets. MySQL is still the live database used by the site, but this folder is what makes the project portable if the database is lost, the server expires, or someone deletes content by mistake.

- `content/content.json` is the latest readable export of CMS tables. It includes draft, published, and soft-deleted rows, so deleted content can still be recovered from git history or from the current snapshot.
- `uploads/` stores files uploaded from the admin media library. Keep this directory tracked because the project data volume is small.
- `backups/` stores timestamped SQL dumps and tar archives created by `make backup`. These files are ignored to avoid repeated large git diffs.
- `public/` is outside this folder, but it is also tracked by git and contains the original site images, PDFs, PPT files, and videos that ship with the project.

Recommended routine:

1. After important admin edits, run `make backup`.
2. Commit `storage/content/content.json` and any new files under `storage/uploads/`.
3. Keep timestamped files under `storage/backups/` local unless you explicitly want an offline archive.
4. Keep the Docker Compose `backup` service running on the server. It refreshes the snapshot every 6 hours by default.
5. If you do not use the Compose backup service, run `make backup-cron-command` and add the printed line to crontab as an alternative.

Recovery routine:

1. Restore repo files from git.
2. Start MySQL with `docker compose up -d mysql`.
3. Run `make restore-content-dry-run` to validate the JSON snapshot.
4. Run `make restore-content` to rebuild CMS tables from `storage/content/content.json`.
5. Run `make media-import` to scan repo assets under `public/` back into the media library.
