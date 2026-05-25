<template>
  <div>
    <div class="admin-page-header">
      <div>
        <h1>后台首页</h1>
        <p>官网内容概览</p>
      </div>
    </div>
    <div class="summary-grid">
      <el-card v-for="item in cards" :key="item.key">
        <span>{{ item.label }}</span>
        <strong>{{ summary[item.key] || 0 }}</strong>
      </el-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { getSummary } from '../../api/admin'

const summary = ref<Record<string, number>>({})
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

onMounted(async () => {
  summary.value = (await getSummary()).counts
})
</script>

<style scoped>
.admin-page-header h1 {
  margin: 0;
  color: #25313b;
}

.admin-page-header p {
  margin: 6px 0 20px;
  color: #6b7280;
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
}

.summary-grid span {
  color: #6b7280;
}

.summary-grid strong {
  color: #7d1231;
  font-size: 30px;
}
</style>
