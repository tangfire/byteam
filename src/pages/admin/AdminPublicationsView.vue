<template>
  <AdminCrudView
    ref="crudRef"
    title="论文管理"
    description="管理论文、链接和首页精选论文"
    resource="publications"
    :defaults="defaults"
    :columns="columns"
    :fields="fields"
    sortable
    :sort-group-key="['year', 'kind']"
  >
    <template #header-actions>
      <router-link class="usage-link" to="/admin/usage-guide">查看使用说明</router-link>
    </template>

    <template #before-table>
      <el-alert
        class="publication-help"
        type="info"
        show-icon
        :closable="false"
      >
        <template #title>
          新增论文时先填年份和类型，前台会自动生成对应年份下的期刊/会议分组；视频论文先到“页面内容”创建视频页，再在论文链接里添加 Video 并选择该视频页。
        </template>
      </el-alert>
    </template>
  </AdminCrudView>
</template>

<script setup lang="ts">
import AdminCrudView, { type FieldConfig } from '../../components/admin/AdminCrudView.vue'

const defaults = { image: '', title: '', authors: '', venue: '', year: new Date().getFullYear(), kind: 'conference', status: 'published', featured: false, links: [] }
const columns: FieldConfig[] = [
  { prop: 'image', label: '图片', type: 'image', width: 90 },
  { prop: 'year', label: '年份', width: 90 },
  { prop: 'kind', label: '类型', width: 120 },
  { prop: 'title', label: '标题', width: 360 },
  { prop: 'featured', label: '精选', width: 90 },
  { prop: 'status', label: '状态', width: 110 },
]
const fields: FieldConfig[] = [
  { prop: 'image', label: '图片', type: 'image' },
  { prop: 'title', label: '标题', type: 'textarea', rows: 3 },
  { prop: 'authors', label: '作者', type: 'textarea', rows: 3 },
  { prop: 'venue', label: '期刊/会议', type: 'textarea', rows: 2 },
  { prop: 'year', label: '年份', type: 'number' },
  { prop: 'kind', label: '类型', type: 'select', options: [{ label: '期刊', value: 'journal' }, { label: '会议', value: 'conference' }] },
  { prop: 'featured', label: '首页精选', type: 'boolean' },
  { prop: 'links', label: '链接', type: 'links' },
  { prop: 'status', label: '发布状态', type: 'select', options: [{ label: '已发布', value: 'published' }, { label: '草稿', value: 'draft' }] },
]
</script>

<style scoped>
.usage-link {
  color: #7d1231;
  font-size: 14px;
  font-weight: 700;
  text-decoration: none;
}

.usage-link:hover {
  text-decoration: underline;
}

.publication-help {
  border-color: rgba(125, 18, 49, 0.16);
  background: #fffafb;
}

.publication-help :deep(.el-alert__title) {
  color: #4b5563;
  line-height: 1.6;
}
</style>
