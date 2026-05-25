<template>
  <AdminCrudView title="本科生培养" description="管理本科生培养成果" resource="undergraduates" :defaults="defaults" :columns="columns" :fields="fields" sortable />
</template>

<script setup lang="ts">
import AdminCrudView, { type FieldConfig } from '../../components/admin/AdminCrudView.vue'

const defaults = { name: '', grade: '', major: '', direction: '', achievements: [], status: 'published' }
const columns: FieldConfig[] = [
  { prop: 'name', label: '姓名', width: 160 },
  { prop: 'grade', label: '年级', width: 100 },
  { prop: 'major', label: '专业', width: 220 },
  { prop: 'direction', label: '方向', width: 240 },
  { prop: 'status', label: '状态', width: 110 },
]
const fields: FieldConfig[] = [
  { prop: 'name', label: '姓名' },
  { prop: 'grade', label: '年级' },
  { prop: 'major', label: '专业' },
  { prop: 'direction', label: '方向' },
  { prop: 'achievements', label: '成果', type: 'list' },
  { prop: 'status', label: '发布状态', type: 'select', options: [{ label: '已发布', value: 'published' }, { label: '草稿', value: 'draft' }] },
]
</script>
