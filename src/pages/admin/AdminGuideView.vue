<template>
  <div class="guide-page">
    <div class="admin-page-header">
      <div>
        <h1>运维说明</h1>
        <p>后台内容、资源文件、备份与恢复流程</p>
      </div>
    </div>

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
      <ol>
        <li>在后台新增、编辑、发布或隐藏内容。</li>
        <li>重要修改后在服务器或本机执行 <code>make backup</code>，刷新项目内恢复快照。</li>
        <li>提交 <code>storage/content/content.json</code> 和 <code>storage/uploads/</code> 中新增的资源文件。</li>
        <li>不要提交 <code>storage/backups/</code> 里的时间戳压缩包，除非需要单独做离线归档。</li>
      </ol>
    </section>

    <section class="guide-section">
      <h2>自动备份</h2>
      <p>
        Docker Compose 里有 <code>backup</code> 服务，默认每 6 小时执行一次 <code>scripts/backup.sh</code>。
        它会刷新 <code>storage/content/content.json</code>，同时生成本地压缩备份到 <code>storage/backups/</code>。
      </p>
      <div class="command-list">
        <code>docker compose up -d backup</code>
        <code>docker compose logs -f backup</code>
      </div>
    </section>

    <section class="guide-section">
      <h2>误删内容怎么恢复</h2>
      <ol>
        <li>优先进入左侧 Trash 页面，从回收站恢复内容。</li>
        <li>如果回收站里没有，查看 git 历史里的 <code>storage/content/content.json</code>。</li>
        <li>找到误删前的版本后，可以对照 JSON 手动补回，也可以由维护者执行恢复命令。</li>
      </ol>
    </section>

    <section class="guide-section">
      <h2>MySQL 崩了怎么恢复</h2>
      <ol>
        <li>从 git 拉回完整项目，确认 <code>public/</code>、<code>storage/uploads/</code> 和 <code>storage/content/content.json</code> 都在。</li>
        <li>启动数据库：<code>docker compose up -d mysql</code></li>
        <li>验证快照：<code>make restore-content-dry-run</code></li>
        <li>恢复内容表：<code>make restore-content</code></li>
        <li>扫描官网静态资源：<code>make media-import</code></li>
        <li>启动全部服务：<code>docker compose up -d</code></li>
      </ol>
    </section>

    <section class="guide-section warning-section">
      <h2>注意事项</h2>
      <ul>
        <li><code>make restore-content</code> 会重建内容表和媒体索引，执行前一定先跑 dry run。</li>
        <li>管理员账号不在 <code>content.json</code> 里，由环境变量 <code>ADMIN_USERNAME</code> 和 <code>ADMIN_PASSWORD</code> 初始化。</li>
        <li>生产环境必须修改默认密码和 <code>JWT_SECRET</code>。</li>
        <li>自动备份不等于自动提交 git；重要更新后仍需要提交恢复快照和上传资源。</li>
      </ul>
    </section>
  </div>
</template>

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
