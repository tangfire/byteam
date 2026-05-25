<template>
  <AdminCrudView
    ref="crudRef"
    title="Publications"
    description="管理论文、链接和首页精选论文"
    resource="publications"
    :defaults="defaults"
    :columns="columns"
    :fields="fields"
    sortable
    :sort-group-key="['year', 'kind']"
  />
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
  { prop: 'kind', label: '类型', type: 'select', options: [{ label: 'Journal', value: 'journal' }, { label: 'Conference', value: 'conference' }] },
  { prop: 'featured', label: '首页精选', type: 'boolean' },
  { prop: 'links', label: '链接', type: 'links' },
  { prop: 'status', label: '发布状态', type: 'select', options: [{ label: 'Published', value: 'published' }, { label: 'Draft', value: 'draft' }] },
]
</script>
