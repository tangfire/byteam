<template>
  <AdminCrudView title="News" description="管理官网新闻和首页最新动态" resource="news" :defaults="defaults" :columns="columns" :fields="fields" />
</template>

<script setup lang="ts">
import AdminCrudView, { type FieldConfig } from '../../components/admin/AdminCrudView.vue'

const today = new Date().toISOString().slice(0, 10)
const defaults = { title: '', content: '', excerpt: '', type: 'publication', typeLabel: '', date: today, color: '#7d1231', status: 'published' }
const columns: FieldConfig[] = [
  { prop: 'date', label: '日期', width: 120 },
  { prop: 'title', label: '标题', width: 260 },
  { prop: 'typeLabel', label: '类型', width: 120 },
  { prop: 'status', label: '状态', width: 110 },
]
const fields: FieldConfig[] = [
  { prop: 'title', label: '标题' },
  { prop: 'content', label: '内容', type: 'textarea', rows: 5 },
  { prop: 'excerpt', label: '首页摘要', type: 'textarea', rows: 3 },
  { prop: 'type', label: '类型', type: 'select', options: [{ label: 'Publication', value: 'publication' }, { label: 'Team', value: 'team' }, { label: 'Award', value: 'award' }, { label: 'Event', value: 'event' }, { label: 'General', value: 'general' }] },
  { prop: 'typeLabel', label: '类型显示（可留空）' },
  { prop: 'date', label: '日期', type: 'date' },
  { prop: 'color', label: '颜色' },
  { prop: 'status', label: '发布状态', type: 'select', options: [{ label: 'Published', value: 'published' }, { label: 'Draft', value: 'draft' }] },
]
</script>
