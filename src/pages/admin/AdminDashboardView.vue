<template>
  <div class="dashboard-page">
    <div class="admin-page-header">
      <div>
        <h1>后台首页</h1>
        <p>官网内容概览和常用维护入口</p>
      </div>
      <el-button :loading="loading" @click="loadSummary">刷新</el-button>
    </div>

    <div class="quick-actions">
      <router-link v-for="action in actions" :key="action.path" :to="action.path" class="quick-action">
        <strong>{{ action.title }}</strong>
        <span>{{ action.description }}</span>
      </router-link>
    </div>

    <div class="summary-grid">
      <el-card v-for="item in cards" :key="item.key" shadow="never">
        <span>{{ item.label }}</span>
        <strong>{{ summary[item.key] || 0 }}</strong>
      </el-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { getSummary } from '../../api/admin'

const summary = ref<Record<string, number>>({})
const loading = ref(false)
const cards = computed(() => [
  { key: 'sitePages', label: '页面内容' },
  { key: 'news', label: '新闻动态' },
  { key: 'people', label: '成员' },
  { key: 'undergraduates', label: '本科生培养' },
  { key: 'publications', label: '论文' },
  { key: 'patents', label: '专利与标准' },
  { key: 'researchProjects', label: '科研项目' },
  { key: 'media', label: '媒体文件' },
])

const actions = [
  { title: '发布新闻', description: '新增或编辑首页新闻动态', path: '/admin/news' },
  { title: '整理媒体', description: '上传资源、修改显示名称', path: '/admin/media' },
  { title: '恢复误删', description: '从回收站找回内容', path: '/admin/trash' },
  { title: '备份同步', description: '刷新快照并同步到 Git', path: '/admin/guide' },
]

const loadSummary = async () => {
  loading.value = true
  try {
    summary.value = (await getSummary()).counts
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '概览加载失败')
  } finally {
    loading.value = false
  }
}

onMounted(loadSummary)
</script>

<style scoped>
.dashboard-page {
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.admin-page-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.admin-page-header h1 {
  margin: 0;
  color: #25313b;
}

.admin-page-header p {
  margin: 6px 0 0;
  color: #6b7280;
}

.quick-actions {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(190px, 1fr));
  gap: 12px;
}

.quick-action {
  display: flex;
  flex-direction: column;
  gap: 6px;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  background: #ffffff;
  padding: 14px;
  color: inherit;
  text-decoration: none;
  transition: border-color 0.18s ease, box-shadow 0.18s ease;
}

.quick-action:hover {
  border-color: #7d1231;
  box-shadow: 0 10px 24px rgba(17, 24, 39, 0.07);
}

.quick-action strong {
  color: #25313b;
}

.quick-action span {
  color: #6b7280;
  font-size: 13px;
}

.summary-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 16px;
}

.summary-grid :deep(.el-card__body) {
  display: flex;
  flex-direction: column;
  gap: 12px;
  border-radius: 8px;
}

.summary-grid span {
  color: #6b7280;
}

.summary-grid strong {
  color: #7d1231;
  font-size: 30px;
}
</style>
