<template>
  <div class="admin-crud">
    <div class="admin-page-header">
      <div>
        <h1>{{ title }}</h1>
        <p>{{ description }}</p>
      </div>
      <el-button type="primary" @click="openCreate">新增</el-button>
    </div>

    <div class="admin-toolbar">
      <el-input v-model="query.q" placeholder="搜索" clearable @keyup.enter="load" />
      <el-select v-model="query.status" placeholder="状态" clearable>
        <el-option label="Published" value="published" />
        <el-option label="Draft" value="draft" />
      </el-select>
      <el-button @click="load">刷新</el-button>
    </div>

    <el-table v-loading="loading" :data="items" border class="admin-table">
      <el-table-column v-for="column in columns" :key="column.prop" :prop="column.prop" :label="column.label" :min-width="column.width || 120">
        <template #default="{ row }">
          <el-image v-if="column.type === 'image' && row[column.prop]" :src="row[column.prop]" fit="cover" class="table-image" />
          <el-tag v-else-if="column.prop === 'status'" :type="row.status === 'published' ? 'success' : 'info'">{{ row.status }}</el-tag>
          <span v-else>{{ formatCell(row[column.prop]) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" fixed="right" width="150">
        <template #default="{ row }">
          <el-button link type="primary" @click="openEdit(row)">编辑</el-button>
          <el-popconfirm title="确认移入回收站？之后可在回收站恢复。" @confirm="remove(row)">
            <template #reference>
              <el-button link type="danger">移入回收站</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      v-model:current-page="query.page"
      v-model:page-size="query.pageSize"
      class="admin-pagination"
      layout="total, sizes, prev, pager, next"
      :total="total"
      @current-change="load"
      @size-change="load"
    />

    <el-dialog v-model="dialogVisible" :title="editing?.id ? '编辑' : '新增'" width="760px">
      <el-form :model="editing" label-width="120px" class="admin-form" v-if="editing">
        <el-form-item v-for="field in fields" :key="field.prop" :label="field.label">
          <el-input v-if="!field.type || field.type === 'text'" v-model="editing[field.prop]" />
          <el-input v-else-if="field.type === 'textarea'" v-model="editing[field.prop]" type="textarea" :rows="field.rows || 4" />
          <el-input-number v-else-if="field.type === 'number'" v-model="editing[field.prop]" :min="0" />
          <el-switch v-else-if="field.type === 'boolean'" v-model="editing[field.prop]" />
          <el-select v-else-if="field.type === 'select'" v-model="editing[field.prop]">
            <el-option v-for="option in field.options || []" :key="option.value" :label="option.label" :value="option.value" />
          </el-select>
          <el-date-picker v-else-if="field.type === 'date'" v-model="editing[field.prop]" value-format="YYYY-MM-DD" type="date" />
          <el-input v-else-if="field.type === 'list'" :model-value="(editing[field.prop] || []).join('\n')" type="textarea" :rows="5" @update:model-value="editing[field.prop] = splitList($event)" />
          <div v-else-if="field.type === 'links'" class="links-editor">
            <div v-for="(link, index) in editing.links" :key="index" class="link-row">
              <el-select v-model="link.type" placeholder="类型">
                <el-option label="Paper" value="paper" />
                <el-option label="Code" value="code" />
                <el-option label="Video" value="video" />
                <el-option label="PPT" value="ppt" />
                <el-option label="Poster" value="poster" />
              </el-select>
              <el-input v-model="link.label" placeholder="标签" />
              <el-input v-model="link.url" placeholder="URL 或文件路径" />
              <el-input v-model="link.routeName" placeholder="视频路由名" />
              <el-button @click="editing.links.splice(index, 1)">删除</el-button>
            </div>
            <el-button @click="editing.links.push({ type: 'paper', label: 'Paper', url: '', routeName: '', sortOrder: editing.links.length + 1 })">添加链接</el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="saving" @click="save">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { createAdmin, deleteAdmin, listAdmin, updateAdmin } from '../../api/admin'

export interface FieldConfig {
  prop: string
  label: string
  width?: number
  type?: 'text' | 'textarea' | 'number' | 'select' | 'date' | 'boolean' | 'list' | 'links' | 'image'
  rows?: number
  options?: { label: string; value: string | number }[]
}

const props = defineProps<{
  title: string
  description: string
  resource: string
  defaults: Record<string, unknown>
  columns: FieldConfig[]
  fields: FieldConfig[]
}>()

const loading = ref(false)
const saving = ref(false)
const items = ref<Record<string, any>[]>([])
const total = ref(0)
const dialogVisible = ref(false)
const editing = ref<Record<string, any> | null>(null)
const query = reactive({ page: 1, pageSize: 20, q: '', status: '' })

const load = async () => {
  loading.value = true
  try {
    const result = await listAdmin<Record<string, any>>(props.resource, query)
    items.value = result.items
    total.value = result.total
  } finally {
    loading.value = false
  }
}

const openCreate = () => {
  editing.value = structuredClone(props.defaults)
  dialogVisible.value = true
}

const openEdit = (row: Record<string, any>) => {
  editing.value = structuredClone(row)
  if (props.resource === 'publications' && !editing.value.links) {
    editing.value.links = []
  }
  dialogVisible.value = true
}

const save = async () => {
  if (!editing.value) return
  saving.value = true
  try {
    if (editing.value.id) {
      await updateAdmin(props.resource, editing.value)
    } else {
      await createAdmin(props.resource, editing.value)
    }
    ElMessage.success('已保存')
    dialogVisible.value = false
    await load()
  } finally {
    saving.value = false
  }
}

const remove = async (row: Record<string, any>) => {
  await deleteAdmin(props.resource, row.id)
  ElMessage.success('已移入回收站')
  await load()
}

const splitList = (value: string) => value.split('\n').map((item) => item.trim()).filter(Boolean)

const formatCell = (value: unknown) => {
  if (Array.isArray(value)) return value.join('；')
  if (typeof value === 'boolean') return value ? '是' : '否'
  return value ?? ''
}

onMounted(load)
</script>

<style scoped>
.admin-crud {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.admin-page-header,
.admin-toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.admin-page-header h1 {
  margin: 0;
  color: #25313b;
  font-size: 24px;
}

.admin-page-header p {
  margin: 6px 0 0;
  color: #6b7280;
}

.admin-toolbar {
  justify-content: flex-start;
}

.admin-toolbar .el-input {
  width: 280px;
}

.admin-toolbar .el-select {
  width: 160px;
}

.admin-table {
  width: 100%;
}

.table-image {
  width: 72px;
  height: 48px;
  border-radius: 6px;
}

.admin-pagination {
  align-self: flex-end;
}

.admin-form {
  max-height: 65vh;
  overflow: auto;
  padding-right: 12px;
}

.links-editor {
  display: flex;
  flex-direction: column;
  gap: 10px;
  width: 100%;
}

.link-row {
  display: grid;
  grid-template-columns: 110px 110px 1fr 150px 70px;
  gap: 8px;
}
</style>
