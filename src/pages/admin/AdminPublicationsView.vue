<template>
  <AdminCrudView
    ref="crudRef"
    title="Publications"
    description="管理论文、链接和首页精选论文"
    resource="publications"
    :defaults="defaults"
    :columns="columns"
    :fields="fields"
    :row-actions="rowActions"
    @row-action="handleRowAction"
  />
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import AdminCrudView, { type FieldConfig, type RowActionConfig } from '../../components/admin/AdminCrudView.vue'
import { movePublication } from '../../api/admin'

const crudRef = ref<InstanceType<typeof AdminCrudView> | null>(null)
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
  { prop: 'image', label: '图片 URL' },
  { prop: 'title', label: '标题', type: 'textarea', rows: 3 },
  { prop: 'authors', label: '作者', type: 'textarea', rows: 3 },
  { prop: 'venue', label: '期刊/会议', type: 'textarea', rows: 2 },
  { prop: 'year', label: '年份', type: 'number' },
  { prop: 'kind', label: '类型', type: 'select', options: [{ label: 'Journal', value: 'journal' }, { label: 'Conference', value: 'conference' }] },
  { prop: 'featured', label: '首页精选', type: 'boolean' },
  { prop: 'links', label: '链接', type: 'links' },
  { prop: 'status', label: '发布状态', type: 'select', options: [{ label: 'Published', value: 'published' }, { label: 'Draft', value: 'draft' }] },
]

const rowActions: RowActionConfig[] = [
  { command: 'top', label: '置顶到同年同类型' },
  { command: 'up', label: '上移一位' },
  { command: 'down', label: '下移一位' },
]

const handleRowAction = async (command: string, row: Record<string, any>) => {
  if (!row.id) return
  if (command !== 'top' && command !== 'up' && command !== 'down') return
  await movePublication(row.id, command)
  ElMessage.success('位置已更新')
  await crudRef.value?.load()
}
</script>
