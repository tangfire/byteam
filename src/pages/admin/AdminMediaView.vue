<template>
  <div class="media-page">
    <div class="admin-page-header">
      <div>
        <h1>Media</h1>
        <p>上传和管理图片、文档、压缩包和视频资源</p>
      </div>
      <div class="media-actions">
        <el-button :loading="importing" @click="handleImport">扫描现有资源</el-button>
        <el-upload :show-file-list="false" :http-request="handleUpload">
          <el-button type="primary" :loading="uploading">上传文件</el-button>
        </el-upload>
      </div>
    </div>

    <el-table v-loading="loading" :data="items" border>
      <el-table-column label="预览" width="110">
        <template #default="{ row }">
          <el-image v-if="row.kind === 'image'" :src="row.url" fit="cover" class="preview" />
          <el-tag v-else>{{ row.kind }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="originalName" label="文件名" min-width="220" />
      <el-table-column prop="url" label="URL" min-width="300">
        <template #default="{ row }">
          <el-input :model-value="row.url" readonly @focus="($event.target as HTMLInputElement).select()" />
        </template>
      </el-table-column>
      <el-table-column label="大小" width="120">
        <template #default="{ row }">{{ formatSize(row.size) }}</template>
      </el-table-column>
      <el-table-column prop="mimeType" label="MIME" width="180" />
      <el-table-column label="操作" width="110">
        <template #default="{ row }">
          <el-popconfirm title="确认移入回收站？被内容引用时后端会拒绝删除。" @confirm="remove(row.id)">
            <template #reference>
              <el-button link type="danger">移入回收站</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      v-model:current-page="page"
      v-model:page-size="pageSize"
      class="pager"
      layout="total, sizes, prev, pager, next"
      :total="total"
      @current-change="load"
      @size-change="load"
    />
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import type { UploadRequestOptions } from 'element-plus'
import { ElMessage } from 'element-plus'
import { deleteAdmin, importPublicMedia, listAdmin, uploadMedia } from '../../api/admin'
import type { MediaAsset } from '../../api/client'

const loading = ref(false)
const uploading = ref(false)
const importing = ref(false)
const items = ref<MediaAsset[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(20)

const load = async () => {
  loading.value = true
  try {
    const result = await listAdmin<MediaAsset>('media', { page: page.value, pageSize: pageSize.value })
    items.value = result.items
    total.value = result.total
  } finally {
    loading.value = false
  }
}

const handleUpload = async (options: UploadRequestOptions) => {
  uploading.value = true
  try {
    await uploadMedia(options.file)
    ElMessage.success('上传成功')
    await load()
  } finally {
    uploading.value = false
  }
}

const handleImport = async () => {
  importing.value = true
  try {
    const result = await importPublicMedia()
    ElMessage.success(`扫描完成：新增 ${result.created}，更新 ${result.updated}，跳过 ${result.skipped}`)
    page.value = 1
    await load()
  } finally {
    importing.value = false
  }
}

const remove = async (id: number) => {
  await deleteAdmin('media', id)
  ElMessage.success('已移入回收站')
  await load()
}

const formatSize = (size: number) => {
  if (size > 1024 * 1024) return `${(size / 1024 / 1024).toFixed(1)} MB`
  if (size > 1024) return `${(size / 1024).toFixed(1)} KB`
  return `${size} B`
}

onMounted(load)
</script>

<style scoped>
.media-page {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.admin-page-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.media-actions {
  display: flex;
  align-items: center;
  gap: 10px;
}

.admin-page-header h1 {
  margin: 0;
  color: #25313b;
}

.admin-page-header p {
  margin: 6px 0 0;
  color: #6b7280;
}

.preview {
  width: 70px;
  height: 48px;
  border-radius: 6px;
}

.pager {
  align-self: flex-end;
}
</style>
