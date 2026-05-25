<template>
  <div class="trash-page">
    <div class="admin-page-header">
      <div>
        <h1>回收站</h1>
        <p>查看已删除内容，并把误删的数据恢复到原管理列表</p>
      </div>
    </div>

    <div class="admin-toolbar">
      <el-select v-model="query.resource" placeholder="类型" @change="handleResourceChange">
        <el-option v-for="resource in resources" :key="resource.value" :label="resource.label" :value="resource.value" />
      </el-select>
      <el-input v-model="query.q" placeholder="搜索已删除内容" clearable @keyup.enter="load" />
      <el-button @click="load">刷新</el-button>
    </div>

    <el-table v-loading="loading" :data="items" border>
      <el-table-column prop="label" label="类型" width="150" />
      <el-table-column prop="title" label="标题/名称" min-width="260" show-overflow-tooltip />
      <el-table-column prop="subtitle" label="说明" min-width="260" show-overflow-tooltip />
      <el-table-column label="原状态" width="120">
        <template #default="{ row }">
          <el-tag v-if="row.status" :type="row.status === 'published' ? 'success' : 'info'">{{ formatStatus(row.status) }}</el-tag>
          <span v-else>-</span>
        </template>
      </el-table-column>
      <el-table-column label="删除时间" width="190">
        <template #default="{ row }">{{ formatDate(row.deletedAt) }}</template>
      </el-table-column>
      <el-table-column label="操作" fixed="right" width="110">
        <template #default="{ row }">
          <el-popconfirm title="确认恢复这条内容？" @confirm="restore(row)">
            <template #reference>
              <el-button link type="primary">恢复</el-button>
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
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { listTrash, restoreTrash } from '../../api/admin'
import type { TrashItem } from '../../api/client'

const resources = [
  { label: '页面内容', value: 'pages' },
  { label: '新闻动态', value: 'news' },
  { label: '成员', value: 'people' },
  { label: '本科生培养', value: 'undergraduates' },
  { label: '论文', value: 'publications' },
  { label: '专利与标准', value: 'patents' },
  { label: '科研项目', value: 'research-projects' },
  { label: '媒体文件', value: 'media' },
]

const loading = ref(false)
const items = ref<TrashItem[]>([])
const total = ref(0)
const query = reactive({
  resource: 'news',
  page: 1,
  pageSize: 20,
  q: '',
})

const load = async () => {
  loading.value = true
  try {
    const result = await listTrash(query)
    items.value = result.items
    total.value = result.total
  } finally {
    loading.value = false
  }
}

const handleResourceChange = async () => {
  query.page = 1
  await load()
}

const restore = async (row: TrashItem) => {
  await restoreTrash(row.resource, row.id)
  ElMessage.success('已恢复')
  await load()
}

const formatDate = (value?: string) => {
  if (!value) return '-'
  return new Date(value).toLocaleString()
}

const formatStatus = (status: string) => status === 'published' ? '已发布' : '草稿'

onMounted(load)
</script>

<style scoped>
.trash-page {
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
}

.admin-page-header p {
  margin: 6px 0 0;
  color: #6b7280;
}

.admin-toolbar {
  justify-content: flex-start;
}

.admin-toolbar .el-select {
  width: 220px;
}

.admin-toolbar .el-input {
  width: 320px;
}

.admin-pagination {
  align-self: flex-end;
}
</style>
