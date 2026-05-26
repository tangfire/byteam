<template>
  <AdminCrudView title="成员管理" description="管理研究生、研究生校友和本科生校友" resource="people" :defaults="defaults" :columns="columns" :fields="fields" sortable sort-group-key="category" />
</template>

<script setup lang="ts">
import AdminCrudView, { type FieldConfig } from '../../components/admin/AdminCrudView.vue'

const defaults = {
  name: '',
  avatarUrl: '',
  avatarObjectX: 50,
  avatarObjectY: 50,
  avatarScale: 100,
  category: 'graduate',
  research: '',
  graduationDate: '',
  status: 'published',
}
const columns: FieldConfig[] = [
  { prop: 'avatarUrl', label: '头像', type: 'avatar', width: 90 },
  { prop: 'name', label: '姓名', width: 180 },
  { prop: 'category', label: '分类', width: 160 },
  { prop: 'graduationDate', label: '毕业信息', width: 180 },
  { prop: 'status', label: '状态', width: 110 },
]
const fields: FieldConfig[] = [
  { prop: 'name', label: '姓名' },
  { prop: 'avatarUrl', label: '头像', type: 'image' },
  { prop: 'avatarDisplay', label: '头像显示', type: 'avatarCrop' },
  { prop: 'category', label: '分类', type: 'select', options: [{ label: '研究生', value: 'graduate' }, { label: '研究生校友', value: 'graduate_alumni' }, { label: '本科生校友', value: 'undergraduate_alumni' }] },
  { prop: 'research', label: '研究方向' },
  { prop: 'graduationDate', label: '毕业信息' },
  { prop: 'status', label: '发布状态', type: 'select', options: [{ label: '已发布', value: 'published' }, { label: '草稿', value: 'draft' }] },
]
</script>
