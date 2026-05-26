<template>
  <div class="guide-page">
    <div class="admin-page-header">
      <div>
        <h1>运维说明</h1>
        <p>后台内容、资源文件、备份与恢复流程</p>
      </div>
    </div>

    <section class="guide-section takeover-section">
      <div class="section-title-row">
        <div>
          <h2>第一次接手先看这里</h2>
          <p>日常只需要记住三件事：后台改内容、重要修改后备份、服务器出问题先 dry run 再恢复。</p>
        </div>
      </div>
      <div class="takeover-grid">
        <div v-for="item in takeoverCards" :key="item.title" class="takeover-card">
          <strong>{{ item.title }}</strong>
          <p>{{ item.text }}</p>
        </div>
      </div>
    </section>

    <section class="guide-section">
      <h2>这套后台管理了什么</h2>
      <div class="guide-grid">
        <div class="guide-item">
          <strong>MySQL</strong>
          <span>网站运行时读取的正式数据，包括新闻、成员、论文、专利、项目和媒体索引。</span>
        </div>
        <div class="guide-item">
          <strong>storage/content/content.json</strong>
          <span>项目内的可恢复内容快照，包含已发布、草稿和回收站里的软删除数据。</span>
        </div>
        <div class="guide-item">
          <strong>storage/uploads/</strong>
          <span>后台上传的图片、PDF、PPT、压缩包、视频等文件，应随项目一起提交到 git。</span>
        </div>
        <div class="guide-item">
          <strong>public/</strong>
          <span>官网原有静态资源目录，部署后可通过 Media 页面扫描进媒体库。</span>
        </div>
      </div>
    </section>

    <section class="guide-section">
      <h2>日常维护流程</h2>
      <div class="step-flow">
        <div v-for="(step, index) in dailySteps" :key="step.title" class="step-card">
          <span>{{ index + 1 }}</span>
          <strong>{{ step.title }}</strong>
          <p>{{ step.text }}</p>
        </div>
      </div>
    </section>

    <section class="guide-section">
      <div class="section-title-row">
        <div>
          <h2>一键备份与同步</h2>
          <p>备份按钮会直接从 MySQL 导出内容快照；同步按钮会先刷新快照，再调用 Git 同步脚本。</p>
        </div>
        <el-button :loading="statusLoading" @click="loadStatus">刷新状态</el-button>
      </div>
      <div class="guide-grid status-grid">
        <div class="guide-item">
          <strong>内容快照</strong>
          <span>{{ snapshotText }}</span>
        </div>
        <div class="guide-item">
          <strong>上传资源</strong>
          <span>{{ uploadsText }}</span>
        </div>
        <div class="guide-item">
          <strong>本地时间点备份</strong>
          <span>{{ backupsText }}</span>
        </div>
        <div class="guide-item">
          <strong>长期内容检查点</strong>
          <span>{{ checkpointsText }}</span>
        </div>
        <div class="guide-item">
          <strong>Git 状态</strong>
          <span>{{ gitText }}</span>
        </div>
      </div>
      <div class="guide-actions">
        <el-popconfirm
          title="确认现在从 MySQL 刷新项目内恢复快照？这不会删除数据。"
          @confirm="handleBackup"
        >
          <template #reference>
            <el-button type="primary" :loading="backupRunning">刷新备份快照</el-button>
          </template>
        </el-popconfirm>
        <el-popconfirm
          title="确认执行一键同步到 Git？会先刷新快照，再提交恢复文件和新增上传资源。"
          @confirm="handleGitSync"
        >
          <template #reference>
            <el-button :loading="syncRunning">一键同步到 Git</el-button>
          </template>
        </el-popconfirm>
      </div>
      <pre v-if="commandOutput" class="command-output">{{ commandOutput }}</pre>
    </section>

    <section class="guide-section">
      <h2>自动备份</h2>
      <p>
        Docker Compose 里有 <code>backup</code> 服务，默认每 6 小时执行一次 <code>scripts/backup.sh</code>。
        它会刷新 <code>storage/content/content.json</code>，同时生成本地短期备份到 <code>storage/backups/</code>。
        默认保留最近 28 份，也就是 6 小时一次时大约保留 7 天。
        它还会生成 git 可跟踪的周/月内容检查点，放在 <code>storage/content/checkpoints/</code>。
        如果数据库暂时不可用，脚本会保留上一份正常快照，不会用空文件覆盖它，也不会留下新的无效时间点备份。
      </p>
      <div class="command-list">
        <code>docker compose up -d backup</code>
        <code>docker compose logs -f backup</code>
      </div>
    </section>

    <section class="guide-section">
      <h2>自动同步到 Git</h2>
      <p>
        <code>git-sync</code> 是可选服务。服务器配置好 git 写权限后，它只会提交
        <code>storage/content/content.json</code> 和 <code>storage/uploads/</code>，
        不会提交代码文件或 <code>storage/backups/</code> 压缩包。
        后台“一键同步到 Git”会先刷新备份快照再同步；定时 <code>git-sync</code> 服务默认只提交 backup 服务已经生成好的快照，所以需要和 <code>backup</code> 服务一起运行。
        为了防止误删扩散，默认不会自动提交 <code>storage/uploads/</code> 里的删除操作。
      </p>
      <div class="command-list">
        <code>docker compose up -d backup</code>
        <code>docker compose --profile git-sync up -d git-sync</code>
        <code>docker compose logs -f git-sync</code>
      </div>
    </section>

    <section class="guide-section">
      <h2>误删内容怎么恢复</h2>
      <ol>
        <li>优先进入左侧 Trash 页面，从回收站恢复内容。</li>
        <li>如果回收站里没有，查看 <code>storage/content/checkpoints/</code> 或 git 历史里的 <code>storage/content/content.json</code>。</li>
        <li>找到误删前的版本后，可以对照 JSON 手动补回，也可以由维护者执行恢复命令。</li>
      </ol>
    </section>

    <section class="guide-section">
      <h2>MySQL 崩了怎么恢复</h2>
      <ol>
        <li>从 git 拉回完整项目，确认 <code>public/</code>、<code>storage/uploads/</code> 和 <code>storage/content/content.json</code> 都在。</li>
        <li>本地开发恢复：<code>docker compose up -d mysql</code>，再执行 <code>make restore-content-dry-run</code> 和 <code>make restore-content</code>。</li>
        <li>生产服务器恢复：先确认 <code>.env.prod</code> 已配置，再执行 <code>docker compose --env-file .env.prod -f compose.prod.yaml up -d --build mysql api</code>。</li>
        <li>生产环境验证快照：<code>make restore-content-prod-dry-run</code></li>
        <li>生产环境恢复内容表：<code>make restore-content-prod</code></li>
        <li>如果要恢复某个月的长期检查点，先 dry run：<code>make restore-content-file-prod-dry-run FILE=storage/content/checkpoints/monthly/2026-05.json</code></li>
        <li>确认后恢复该文件：<code>make restore-content-file-prod FILE=storage/content/checkpoints/monthly/2026-05.json</code></li>
        <li>扫描官网静态资源：生产环境执行 <code>make media-import-prod</code>，本地执行 <code>make media-import</code>。</li>
        <li>启动全部生产服务：<code>docker compose --env-file .env.prod -f compose.prod.yaml up -d --build</code></li>
      </ol>
    </section>

    <section class="guide-section">
      <h2>本地开发到服务器部署</h2>
      <ol>
        <li>本地开发阶段，MySQL 数据在本机 Docker volume 里，不会天然跟着源码走。</li>
        <li>本地确认内容后，点“刷新备份快照”或执行 <code>make backup</code>，生成最新 <code>storage/content/content.json</code>。</li>
        <li>把 <code>storage/content/content.json</code> 和 <code>storage/uploads/</code> 里新增资源提交到 git。</li>
        <li>服务器拉取项目后，复制 <code>.env.prod.example</code> 为 <code>.env.prod</code>，修改管理员密码、JWT 密钥和 MySQL 密码。</li>
        <li>服务器先启动生产 MySQL 和 API：<code>docker compose --env-file .env.prod -f compose.prod.yaml up -d --build mysql api</code></li>
        <li>执行 <code>make restore-content-prod-dry-run</code>，确认快照无误后执行 <code>make restore-content-prod</code>。</li>
        <li>执行 <code>make media-import-prod</code>，把 <code>public/</code> 里的旧资源登记进媒体库。</li>
        <li>之后服务器上的后台就是正式数据入口；启用 <code>backup</code> 和可选 <code>git-sync</code> 后，新增内容也会继续沉淀回项目目录和 git。</li>
      </ol>
    </section>

    <section class="guide-section">
      <h2>服务器首次部署空 MySQL</h2>
      <ol>
        <li>部署前在本地后台点“刷新备份快照”，确认 <code>storage/content/content.json</code> 是最新内容。</li>
        <li>把最新代码、<code>storage/content/content.json</code>、<code>storage/content/checkpoints/</code> 和 <code>storage/uploads/</code> 推送到 git。</li>
        <li>服务器拉取项目后，复制 <code>.env.prod.example</code> 为 <code>.env.prod</code>，修改所有默认密码和密钥。</li>
        <li>宝塔已有旧站点时，建议先保留旧站点目录；新版先用 <code>WEB_PORT=127.0.0.1:18080</code> 在本机端口跑起来，验证后再用宝塔反向代理切过去。</li>
        <li>先启动生产 MySQL 和 API：<code>docker compose --env-file .env.prod -f compose.prod.yaml up -d --build mysql api</code></li>
        <li>验证项目内快照能恢复：<code>make restore-content-prod-dry-run</code></li>
        <li>确认无误后恢复内容到服务器 MySQL：<code>make restore-content-prod</code></li>
        <li>把 <code>public/</code> 里的旧图片、PDF、PPT、视频登记进媒体库：<code>make media-import-prod</code></li>
        <li>启动全部服务：<code>docker compose --env-file .env.prod -f compose.prod.yaml up -d --build</code></li>
        <li>访问前台和后台确认内容正常；之后服务器后台就是正式内容入口。</li>
      </ol>
    </section>

    <section class="guide-section">
      <h2>生产环境密码说明</h2>
      <ul>
        <li><code>WEB_PORT</code> 是生产 Web 容器暴露端口；宝塔反代部署推荐写成 <code>127.0.0.1:18080</code>，避免容器直接占用公网 80 端口。</li>
        <li><code>MYSQL_PASSWORD</code> 是后端连接业务数据库用户 <code>byml</code> 的密码。</li>
        <li><code>MYSQL_ROOT_PASSWORD</code> 是 MySQL <code>root</code> 管理员密码，备份、恢复和健康检查会用到。</li>
        <li><code>ADMIN_PASSWORD</code> 是后台管理系统登录密码，不是 MySQL 密码。</li>
        <li><code>JWT_SECRET</code> 用来签发后台登录 token，生产环境必须使用很长的随机字符串。</li>
        <li>MySQL 第一次初始化后，密码会写入 Docker volume；之后直接改 <code>.env.prod</code> 里的 MySQL 密码不会自动改库内密码，反而可能导致后端连不上。</li>
      </ul>
    </section>

    <section class="guide-section">
      <h2>宝塔站点怎么切换新版</h2>
      <ol>
        <li>旧的宝塔静态站点先不要删除，保留一段时间作为回滚方案。</li>
        <li>新版 Docker Compose 正常启动后，先在服务器验证 <code>curl http://127.0.0.1:18080</code> 能返回页面。</li>
        <li>宝塔站点里添加反向代理，目标地址填 <code>http://127.0.0.1:18080</code>；域名和 SSL 仍由宝塔原站点处理。</li>
        <li>确认 Nginx 或宝塔站点允许大文件上传，建议设置 <code>client_max_body_size 100m</code>。</li>
        <li>切换后检查首页、新闻、成员、论文、后台登录和媒体上传。</li>
        <li>如果新版异常，关闭宝塔反向代理即可临时切回旧静态站点。</li>
      </ol>
    </section>

    <section class="guide-section">
      <h2>之后更新代码怎么部署</h2>
      <ol>
        <li>如果只是后台里新增、编辑、发布、隐藏内容，不需要重新部署；内容已经在服务器 MySQL 里。</li>
        <li>如果改了硬编码页面、导航栏、样式、前端组件或 Go 后端代码，需要在服务器拉取最新代码并重新构建相关容器。</li>
        <li>只改前端源码时，一般执行：<code>git pull</code>，再执行 <code>docker compose --env-file .env.prod -f compose.prod.yaml up -d --build web</code></li>
        <li>改了后端、数据库模型、备份脚本或 Compose 配置时，执行：<code>git pull</code>，再执行 <code>docker compose --env-file .env.prod -f compose.prod.yaml up -d --build</code></li>
        <li>正常代码更新不会清空 MySQL；不要因为更新前端就执行 <code>make restore-content-prod</code>。</li>
        <li>只有在明确要用项目内 <code>storage/content/content.json</code> 覆盖服务器内容，或者 MySQL 崩了需要恢复时，才执行 restore 命令，并且必须先 dry run。</li>
      </ol>
    </section>

    <section class="guide-section warning-section">
      <h2>注意事项</h2>
      <ul>
        <li><code>make restore-content</code> 和 <code>make restore-content-prod</code> 会重建内容表和媒体索引，执行前一定先跑 dry run。</li>
        <li>管理员账号不在 <code>content.json</code> 里，由环境变量 <code>ADMIN_USERNAME</code> 和 <code>ADMIN_PASSWORD</code> 初始化。</li>
        <li>生产环境必须修改默认密码和 <code>JWT_SECRET</code>。</li>
        <li>真实生产配置在 <code>.env.prod</code>，这个文件已被 git 忽略，不要提交到仓库。</li>
        <li>只有启用 <code>git-sync</code> 后，恢复快照和新增上传文件才会自动推送到 git。</li>
        <li><code>git-sync</code> 默认拒绝异常小的内容快照，也不会自动推送上传资源删除。</li>
        <li><code>backup</code> 会拒绝比上一份突然小很多的内容快照，默认阈值是上一份的 30%。</li>
        <li><code>BACKUP_KEEP_COUNT</code> 控制本地时间点备份保留数量；设置为 <code>0</code> 才会关闭自动清理。</li>
      </ul>
    </section>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { getMaintenanceStatus, runBackup, runGitSync } from '../../api/admin'
import type { MaintenanceStatus } from '../../api/client'

const status = ref<MaintenanceStatus | null>(null)
const statusLoading = ref(false)
const backupRunning = ref(false)
const syncRunning = ref(false)
const commandOutput = ref('')

const dailySteps = [
  { title: '后台维护内容', text: '新增、编辑、发布或隐藏新闻、成员、论文、页面和媒体文件。' },
  { title: '自动刷新快照', text: 'backup 服务默认每 6 小时把 MySQL 内容写入 storage/content/content.json。' },
  { title: '重要修改手动备份', text: '本页“刷新备份快照”会立刻导出当前数据库内容，适合发布前后使用。' },
  { title: '同步到 Git', text: '服务器配置好写权限后，“一键同步到 Git”会提交恢复快照和新增上传文件。' },
  { title: '短期备份自动清理', text: 'storage/backups/ 默认保留最近 28 份，不会被提交到 git。' },
]

const takeoverCards = [
  { title: '日常更新', text: '进入对应管理页新增或编辑内容，确认无误后发布；不需要改前端源码，也不需要重新构建前端。' },
  { title: '重要修改', text: '发布重要内容后，来这里点“刷新备份快照”；服务器配置了 Git 写权限时，再点“一键同步到 Git”。' },
  { title: '误删恢复', text: '先去回收站恢复；如果发现太晚，再从 storage/content/checkpoints 或 git 历史恢复。' },
  { title: '服务器空库', text: '拉取项目后先配置 .env.prod，启动生产 MySQL 和 API，再执行 make restore-content-prod-dry-run，确认没问题后执行 make restore-content-prod。' },
]

const loadStatus = async () => {
  statusLoading.value = true
  try {
    status.value = await getMaintenanceStatus()
  } catch (error) {
    const message = error instanceof Error ? error.message : '维护状态读取失败'
    commandOutput.value = message
    ElMessage.error(message)
  } finally {
    statusLoading.value = false
  }
}

const handleBackup = async () => {
  backupRunning.value = true
  commandOutput.value = ''
  try {
    const result = await runBackup()
    commandOutput.value = result.output || '备份完成'
    ElMessage.success('备份快照已刷新')
    await loadStatus()
  } catch (error) {
    const message = error instanceof Error ? error.message : '备份失败'
    commandOutput.value = message
    ElMessage.error(message)
  } finally {
    backupRunning.value = false
  }
}

const handleGitSync = async () => {
  syncRunning.value = true
  commandOutput.value = ''
  try {
    const result = await runGitSync()
    commandOutput.value = result.output || '同步完成'
    ElMessage.success('已同步到 Git')
    await loadStatus()
  } catch (error) {
    const message = error instanceof Error ? error.message : '同步失败'
    commandOutput.value = message
    ElMessage.error(message)
  } finally {
    syncRunning.value = false
  }
}

const snapshotText = computed(() => {
  const snapshot = status.value?.contentSnapshot
  if (!snapshot?.exists) return '还没有生成 content.json'
  return `${formatSize(snapshot.size)}，更新时间 ${formatDate(snapshot.updatedAt)}`
})

const uploadsText = computed(() => {
  const uploads = status.value?.uploads
  if (!uploads?.exists) return '上传目录还不存在'
  return `${uploads.fileCount} 个文件，合计 ${formatSize(uploads.totalSize)}`
})

const backupsText = computed(() => {
  const backups = status.value?.backups
  if (!backups) return '正在读取'
  const newest = backups.newest ? `，最新 ${backups.newest}` : ''
  return `${backups.count} 份，本地最多保留 ${backups.keepCount} 份${newest}`
})

const checkpointsText = computed(() => {
  const weekly = status.value?.weeklyCheckpoints
  const monthly = status.value?.monthlyCheckpoints
  if (!weekly || !monthly) return '正在读取'
  const weeklyCount = weekly.count || 0
  const monthlyCount = monthly.count || 0
  const latestMonthly = monthly.newest ? `，最近月检查点 ${monthly.newest}` : ''
  return `周 ${weeklyCount} 份，月 ${monthlyCount} 份${latestMonthly}`
})

const gitText = computed(() => {
  const git = status.value?.git
  if (!git) return '正在读取'
  if (!git.available) return git.error || '当前环境无法读取 git 状态'
  const changes = Array.isArray(git.changes) ? git.changes : []
  if (changes.length === 0) return `分支 ${git.branch || '-'}，恢复文件没有待提交变化`
  return `分支 ${git.branch || '-'}，${changes.length} 个恢复文件变更待提交`
})

const formatSize = (size = 0) => {
  if (size > 1024 * 1024 * 1024) return `${(size / 1024 / 1024 / 1024).toFixed(2)} GB`
  if (size > 1024 * 1024) return `${(size / 1024 / 1024).toFixed(1)} MB`
  if (size > 1024) return `${(size / 1024).toFixed(1)} KB`
  return `${size} B`
}

const formatDate = (value?: string) => {
  if (!value) return '-'
  return new Date(value).toLocaleString()
}

onMounted(loadStatus)
</script>

<style scoped>
.guide-page {
  display: flex;
  flex-direction: column;
  gap: 18px;
  max-width: 1120px;
}

.admin-page-header h1 {
  margin: 0;
  color: #25313b;
}

.admin-page-header p {
  margin: 6px 0 0;
  color: #6b7280;
}

.guide-section {
  background: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 20px;
}

.section-title-row {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 16px;
}

.section-title-row p {
  margin: -6px 0 14px;
}

.status-grid {
  margin-top: 2px;
}

.takeover-section {
  border-color: #ead2d9;
  background: #fffafb;
}

.takeover-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 12px;
}

.takeover-card {
  border: 1px solid #ead2d9;
  border-radius: 8px;
  background: #ffffff;
  padding: 14px;
}

.takeover-card strong {
  color: #7d1231;
}

.takeover-card p {
  margin: 8px 0 0;
  font-size: 13px;
}

.guide-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  margin-top: 14px;
}

.command-output {
  margin: 14px 0 0;
  max-height: 240px;
  overflow: auto;
  white-space: pre-wrap;
  word-break: break-word;
  background: #111827;
  border-radius: 8px;
  color: #e5e7eb;
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
  font-size: 12px;
  line-height: 1.6;
  padding: 12px;
}

.guide-section h2 {
  margin: 0 0 14px;
  color: #25313b;
  font-size: 18px;
}

.guide-section p,
.guide-section li {
  color: #4b5563;
  line-height: 1.7;
}

.guide-section ol,
.guide-section ul {
  margin: 0;
  padding-left: 22px;
}

.step-flow {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 12px;
}

.step-card {
  display: flex;
  flex-direction: column;
  gap: 8px;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  background: #fbfbfc;
  padding: 14px;
}

.step-card span {
  width: 26px;
  height: 26px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: #7d1231;
  color: #ffffff;
  font-size: 13px;
  font-weight: 700;
}

.step-card strong {
  color: #25313b;
}

.step-card p {
  margin: 0;
  color: #4b5563;
  font-size: 13px;
}

.guide-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 12px;
}

.guide-item {
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 14px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.guide-item strong {
  color: #7d1231;
}

.guide-item span {
  color: #4b5563;
  line-height: 1.6;
}

.command-list {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  margin-top: 14px;
}

code {
  background: #f3f4f6;
  border: 1px solid #e5e7eb;
  border-radius: 6px;
  color: #111827;
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
  font-size: 13px;
  padding: 2px 6px;
}

.command-list code {
  padding: 8px 10px;
}

.warning-section {
  border-color: #f3d6dc;
  background: #fff8fa;
}
</style>
