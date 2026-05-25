# BYML CMS Storage

This directory keeps the small CMS data and upload assets visible to git.

- `content/content.json` is the latest readable export of MySQL CMS tables, including draft, published, and soft-deleted rows.
- `uploads/` stores files uploaded from the admin media library. Keep this directory tracked because the project data volume is small.
- `backups/` stores timestamped SQL dumps and tar archives created by `make backup`. These files are ignored to avoid repeated large git diffs.

Run `make media-import` after deploying repo assets to scan `public/` into the media library, and run `make backup` after content changes when you want to refresh the git-visible snapshot.
