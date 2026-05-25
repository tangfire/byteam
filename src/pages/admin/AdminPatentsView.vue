<template>
  <AdminCrudView title="专利与标准" description="管理授权专利、实审专利和团体标准" resource="patents" :defaults="defaults" :columns="columns" :fields="fields" sortable sort-group-key="category" />
</template>

<script setup lang="ts">
import AdminCrudView, { type FieldConfig } from '../../components/admin/AdminCrudView.vue'

const defaults = { authors: '', title: '', date: '', country: '中国', number: '', category: 'granted', status: 'published' }
const columns: FieldConfig[] = [
  { prop: 'category', label: '分类', width: 120 },
  { prop: 'title', label: '标题', width: 360 },
  { prop: 'authors', label: '作者', width: 240 },
  { prop: 'number', label: '编号', width: 180 },
  { prop: 'status', label: '状态', width: 110 },
]
const fields: FieldConfig[] = [
  { prop: 'authors', label: '作者', type: 'textarea', rows: 2 },
  { prop: 'title', label: '标题', type: 'textarea', rows: 3 },
  { prop: 'date', label: '日期' },
  { prop: 'country', label: '国家/组织' },
  { prop: 'number', label: '编号' },
  { prop: 'category', label: '分类', type: 'select', options: [{ label: '授权', value: 'granted' }, { label: '实审', value: 'review' }, { label: '团体标准', value: 'standard' }] },
  { prop: 'status', label: '发布状态', type: 'select', options: [{ label: '已发布', value: 'published' }, { label: '草稿', value: 'draft' }] },
]
</script>
