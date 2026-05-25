<template>
  <AdminCrudView title="Research Projects" description="管理科研项目列表" resource="research-projects" :defaults="defaults" :columns="columns" :fields="fields" sortable />
</template>

<script setup lang="ts">
import AdminCrudView, { type FieldConfig } from '../../components/admin/AdminCrudView.vue'

const defaults = { title: '', fund: '', number: '', period: '', amount: '', projectStatus: '在研', role: '', status: 'published' }
const columns: FieldConfig[] = [
  { prop: 'title', label: '项目', width: 360 },
  { prop: 'fund', label: '基金来源', width: 260 },
  { prop: 'projectStatus', label: '项目状态', width: 120 },
  { prop: 'role', label: '角色', width: 140 },
  { prop: 'status', label: '发布状态', width: 110 },
]
const fields: FieldConfig[] = [
  { prop: 'title', label: '项目名称', type: 'textarea', rows: 2 },
  { prop: 'fund', label: '基金来源', type: 'textarea', rows: 2 },
  { prop: 'number', label: '项目编号' },
  { prop: 'period', label: '执行期限' },
  { prop: 'amount', label: '项目经费' },
  { prop: 'projectStatus', label: '项目状态', type: 'select', options: [{ label: '在研', value: '在研' }, { label: '已结题', value: '已结题' }] },
  { prop: 'role', label: '角色' },
  { prop: 'status', label: '发布状态', type: 'select', options: [{ label: 'Published', value: 'published' }, { label: 'Draft', value: 'draft' }] },
]
</script>
