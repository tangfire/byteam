# BYML Lab Website

BYML 官网项目，包含 Vue 前台、Vue 后台管理、Go API、MySQL、媒体上传和内容备份恢复。

## Local Development

```bash
docker compose up -d
```

- 前台和后台：http://localhost:5173
- 后台入口：http://localhost:5173/#/admin
- API：http://localhost:8080

本地 Compose 是开发模式：前端使用 Vite dev server，后端使用 `go run`。

## Production Deployment

服务器正式部署请使用生产 Compose：

```bash
cp .env.prod.example .env.prod
# 修改 .env.prod 里的密码和密钥
docker compose --env-file .env.prod -f compose.prod.yaml up -d --build
```

生产 Compose 会：

- 用 Nginx 托管构建后的前端静态文件。
- 通过 Nginx 反向代理 `/api/` 和 `/uploads/` 到 Go API。
- 使用编译后的 Go 二进制运行后端。
- 默认只暴露 Web 端口，不把 MySQL 端口暴露到公网。
- 启动 `backup` 服务，每 6 小时刷新项目内恢复快照。
- 提供容器内恢复命令，不需要把 MySQL 端口暴露到公网。

## Empty Server Database

如果服务器 MySQL 是空库，按这个顺序恢复项目内内容快照：

```bash
cp .env.prod.example .env.prod
# 修改 .env.prod 里的 ADMIN_PASSWORD、JWT_SECRET、MYSQL_PASSWORD、MYSQL_ROOT_PASSWORD

docker compose --env-file .env.prod -f compose.prod.yaml up -d --build mysql api
make restore-content-prod-dry-run
make restore-content-prod
make media-import-prod
docker compose --env-file .env.prod -f compose.prod.yaml up -d --build
```

如果 `storage/content/content.json` 不是最新内容，先在本地或旧服务器后台“运维说明”页点击“刷新备份快照”，并提交最新 `storage/content/content.json`、`storage/content/checkpoints/` 和 `storage/uploads/`。当前快照较旧时，`make restore-content-prod-dry-run` 可能显示 `sitePages=0`；后端会先 seed 默认页面，恢复工具不会用旧快照清掉这些页面，但正式迁移前仍建议刷新一次快照。

## Backup And Recovery

- 最新内容快照：`storage/content/content.json`
- 长期检查点：`storage/content/checkpoints/weekly/` 和 `storage/content/checkpoints/monthly/`
- 上传资源：`storage/uploads/`
- 本地短期备份：`storage/backups/`，默认不提交 Git

常用命令：

```bash
make backup
make restore-content-dry-run
make restore-content
make restore-content-prod-dry-run
make restore-content-prod
make git-sync-backup-dry-run
make git-sync-backup
```

后台“运维说明”页也提供备份、Git 同步和恢复流程说明。
