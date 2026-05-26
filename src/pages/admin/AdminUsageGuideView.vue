<template>
  <div class="usage-guide-page">
    <div class="admin-page-header">
      <div>
        <h1>使用说明</h1>
        <p>日常录入、编辑和发布官网内容的操作流程</p>
      </div>
    </div>

    <section class="guide-section intro-section">
      <div>
        <h2>日常维护只记住这条线</h2>
        <p>先上传资源，再新增内容，预览无误后发布；重要修改完成后去运维说明刷新备份快照。</p>
      </div>
      <router-link class="primary-link" to="/admin/guide">去运维说明</router-link>
    </section>

    <section class="guide-section">
      <h2>新增一篇论文</h2>
      <ol>
        <li>进入左侧 <strong>论文管理</strong>，点击右上角 <strong>新增论文</strong>。</li>
        <li>填写论文图片、标题、作者、期刊/会议、年份和类型。类型选择 <strong>期刊</strong> 或 <strong>会议</strong> 后，前台会自动归到对应年份和分类。</li>
        <li>如果想让论文出现在首页 Publications 区域，打开 <strong>首页精选</strong>。</li>
        <li>在 <strong>链接</strong> 里添加 Paper、Code、PPT、Poster、Video 等按钮；可以粘贴外部地址，也可以上传文件后自动填入。</li>
        <li>确认无误后点 <strong>保存并发布</strong>。如果还没准备好，点 <strong>存为草稿</strong>，草稿不会显示在前台。</li>
      </ol>
    </section>

    <section class="guide-section">
      <h2>论文带视频怎么做</h2>
      <ol>
        <li>进入 <strong>页面内容</strong>，点击 <strong>新增论文视频页</strong>。</li>
        <li>填写视频页标题、访问路径和视频文件。访问路径建议类似 <code>video-paper-title</code>，系统会生成前台路径 <code>/video/video-paper-title</code>。</li>
        <li>保存视频页后，回到 <strong>论文管理</strong> 编辑对应论文。</li>
        <li>在链接区域点击 <strong>添加Video</strong>，类型选择 Video，然后从 <strong>选择站内视频页</strong> 下拉框里选刚创建的视频页。</li>
        <li>保存论文后，前台论文卡片里的 Video 按钮会跳到对应视频播放页。</li>
      </ol>
    </section>

    <section class="guide-section">
      <h2>上传和选择媒体文件</h2>
      <ul>
        <li>图片、PDF、PPT/PPTX、ZIP、MP4 都可以在对应编辑弹窗里直接上传并使用。</li>
        <li>也可以先进入 <strong>媒体文件</strong> 页面统一上传，再回到内容编辑页点击 <strong>选择媒体</strong>。</li>
        <li>媒体文件名可以在媒体库里改成更容易辨认的显示名，前台访问地址不会因此变化。</li>
        <li>不要直接在服务器里手动删除 <code>storage/uploads/</code> 文件；需要隐藏内容时先在后台操作。</li>
      </ul>
    </section>

    <section class="guide-section">
      <h2>成员头像怎么调</h2>
      <ol>
        <li>进入 <strong>成员管理</strong>，新增或编辑成员。</li>
        <li>选择或上传头像后，在头像预览区域拖动图片，让人脸处于合适位置。</li>
        <li>用水平位置、垂直位置和缩放控制细调；前台成员页和校友页会使用同一套裁剪参数。</li>
      </ol>
    </section>

    <section class="guide-section">
      <h2>发布、隐藏、删除和恢复</h2>
      <ul>
        <li><strong>发布</strong>：内容会显示到前台公开页面。</li>
        <li><strong>隐藏</strong>：内容仍保留在后台，但前台不显示，适合暂时下线。</li>
        <li><strong>移入回收站</strong>：软删除，误操作后可以去 <strong>回收站</strong> 恢复。</li>
        <li>所有删除、清空、状态切换都有二次确认；真正恢复历史快照请看 <strong>运维说明</strong>。</li>
      </ul>
    </section>

    <section class="guide-section warning-section">
      <h2>改完重要内容后</h2>
      <p>
        后台内容会先写入 MySQL。重要修改完成后，进入 <strong>运维说明</strong> 点
        <strong>刷新备份快照</strong>，服务器配置好 Git 写权限时再点 <strong>一键同步到 Git</strong>。
        这样以后 MySQL 出问题，也能从项目里的 <code>storage/content/content.json</code> 和
        <code>storage/uploads/</code> 恢复。
      </p>
    </section>
  </div>
</template>

<style scoped>
.usage-guide-page {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.admin-page-header h1 {
  margin: 0;
  color: #25313b;
  font-size: 24px;
  line-height: 1.25;
}

.admin-page-header p {
  margin: 6px 0 0;
  color: #6b7280;
}

.guide-section {
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  background: #ffffff;
  padding: 18px 20px;
}

.intro-section {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  border-color: rgba(125, 18, 49, 0.22);
  background: #fffafb;
}

.guide-section h2 {
  margin: 0 0 10px;
  color: #7d1231;
  font-size: 18px;
}

.guide-section p,
.guide-section li {
  color: #4b5563;
  line-height: 1.75;
}

.guide-section p {
  margin: 0;
}

.guide-section ol,
.guide-section ul {
  margin: 0;
  padding-left: 22px;
}

.guide-section code {
  border-radius: 5px;
  background: #f3f4f6;
  padding: 2px 6px;
  color: #7d1231;
}

.primary-link {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-height: 36px;
  border-radius: 6px;
  background: #7d1231;
  padding: 0 14px;
  color: #ffffff;
  text-decoration: none;
  font-weight: 700;
  white-space: nowrap;
}

.warning-section {
  border-color: #f5c2cf;
  background: #fff7fa;
}

@media (max-width: 760px) {
  .intro-section {
    align-items: flex-start;
    flex-direction: column;
  }
}
</style>
