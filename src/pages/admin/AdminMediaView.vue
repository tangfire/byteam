<template>
  <div class="media-page">
    <div class="admin-page-header">
      <div>
        <h1>媒体文件</h1>
        <p>上传和管理图片、文档、压缩包和视频资源</p>
      </div>
      <div class="media-actions">
        <el-button :loading="importing" @click="handleImport">扫描现有资源</el-button>
        <el-upload :show-file-list="false" :http-request="handleUpload">
          <el-button type="primary" :loading="uploading">上传文件</el-button>
        </el-upload>
      </div>
    </div>

    <div class="media-toolbar">
      <el-input v-model="query" placeholder="搜索显示名称、文件名或 URL" clearable @keyup.enter="refresh" />
      <el-select v-model="kind" placeholder="类型" clearable>
        <el-option label="图片" value="image" />
        <el-option label="文档" value="document" />
        <el-option label="压缩包" value="archive" />
        <el-option label="视频" value="video" />
      </el-select>
      <el-segmented v-model="usage" :options="usageOptions" @change="refresh" />
      <div class="toolbar-spacer" />
      <span class="toolbar-count">共 {{ total }} 个文件</span>
      <el-button :loading="loading" @click="refresh">搜索</el-button>
    </div>

    <el-table v-loading="loading" :data="items" border class="media-table">
      <el-table-column label="预览" width="110">
        <template #default="{ row }">
          <el-image v-if="row.kind === 'image'" :src="row.url" fit="cover" class="preview" :preview-src-list="[row.url]" preview-teleported />
          <button v-else type="button" class="file-kind" @click="openURL(row.url)">{{ formatKind(row.kind) }}</button>
        </template>
      </el-table-column>
      <el-table-column label="显示名称" min-width="260">
        <template #default="{ row }">
          <div class="media-name-cell">
            <strong>{{ displayName(row) }}</strong>
            <span>{{ row.originalName }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="使用状态" width="110">
        <template #default="{ row }">
          <el-tag :type="row.inUse ? 'success' : 'info'">{{ row.inUse ? '已使用' : '未使用' }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="url" label="URL" min-width="300" show-overflow-tooltip />
      <el-table-column label="大小" width="120">
        <template #default="{ row }">{{ formatSize(row.size) }}</template>
      </el-table-column>
      <el-table-column label="操作" width="260">
        <template #default="{ row }">
          <el-button link type="primary" @click="openRename(row)">改显示名</el-button>
          <el-button link type="primary" @click="copyURL(row.url)">复制 URL</el-button>
          <el-button link @click="openURL(row.url)">打开</el-button>
          <el-popconfirm title="确认移入回收站？被内容引用时后端会拒绝删除。" @confirm="remove(row.id)">
            <template #reference>
              <el-button link type="danger">移入回收站</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
      <template #empty>
        <el-empty :description="emptyText" :image-size="92">
          <el-button v-if="hasFilter" @click="resetFilters">清空筛选</el-button>
          <el-upload v-else :show-file-list="false" :http-request="handleUpload">
            <el-button type="primary" :loading="uploading">上传文件</el-button>
          </el-upload>
        </el-empty>
      </template>
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

  <el-dialog v-model="renameVisible" title="修改显示名称" width="480px">
    <el-form label-width="86px">
      <el-form-item label="显示名称">
        <el-input v-model="renameValue" placeholder="例如：2025 ICIP 论文视频、FedCD 首页配图" maxlength="255" show-word-limit />
      </el-form-item>
      <el-form-item label="原文件名" v-if="renaming">
        <span class="muted-text">{{ renaming.originalName }}</span>
      </el-form-item>
      <el-form-item label="URL" v-if="renaming">
        <span class="muted-text">{{ renaming.url }}</span>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="renameVisible = false">取消</el-button>
      <el-button type="primary" :loading="renamingSaving" @click="saveRename">保存</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import type { UploadRequestOptions } from 'element-plus'
import { ElMessage } from 'element-plus'
import { deleteAdmin, importPublicMedia, listAdmin, updateMediaName, uploadMedia } from '../../api/admin'
import type { MediaAsset } from '../../api/client'

const loading = ref(false)
const uploading = ref(false)
const importing = ref(false)
const items = ref<MediaAsset[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(20)
const query = ref('')
const kind = ref('')
const usage = ref('used')
const renameVisible = ref(false)
const renamingSaving = ref(false)
const renaming = ref<MediaAsset | null>(null)
const renameValue = ref('')

const usageOptions = [
  { label: '已使用', value: 'used' },
  { label: '全部', value: '' },
  { label: '未使用', value: 'unused' },
]

const hasFilter = computed(() => Boolean(query.value || kind.value || usage.value !== 'used'))
const emptyText = computed(() => hasFilter.value ? '没有找到符合筛选条件的媒体文件' : '还没有登记已使用的媒体文件')

const load = async () => {
  loading.value = true
  try {
    const result = await listAdmin<MediaAsset>('media', { page: page.value, pageSize: pageSize.value, q: query.value, kind: kind.value, usage: usage.value })
    items.value = result.items
    total.value = result.total
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '媒体列表加载失败')
  } finally {
    loading.value = false
  }
}

const displayName = (row: MediaAsset) => row.displayName || row.originalName || row.fileName

const refresh = async () => {
  page.value = 1
  await load()
}

const resetFilters = async () => {
  query.value = ''
  kind.value = ''
  usage.value = 'used'
  page.value = 1
  await load()
}

const handleUpload = async (options: UploadRequestOptions) => {
  uploading.value = true
  try {
    await uploadMedia(options.file)
    usage.value = ''
    page.value = 1
    ElMessage.success('上传成功')
    await load()
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '上传失败')
  } finally {
    uploading.value = false
  }
}

const handleImport = async () => {
  importing.value = true
  try {
    const result = await importPublicMedia()
    usage.value = ''
    ElMessage.success(`扫描完成：新增 ${result.created}，更新 ${result.updated}，跳过 ${result.skipped}`)
    page.value = 1
    await load()
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '扫描失败')
  } finally {
    importing.value = false
  }
}

const remove = async (id: number) => {
  try {
    await deleteAdmin('media', id)
    ElMessage.success('已移入回收站')
    await load()
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '删除失败')
  }
}

const formatSize = (size: number) => {
  if (size > 1024 * 1024) return `${(size / 1024 / 1024).toFixed(1)} MB`
  if (size > 1024) return `${(size / 1024).toFixed(1)} KB`
  return `${size} B`
}

const formatKind = (value: string) => {
  const labels: Record<string, string> = {
    image: '图片',
    document: '文档',
    archive: '压缩包',
    video: '视频',
  }
  return labels[value] || value
}

const copyURL = async (url: string) => {
  try {
    await navigator.clipboard.writeText(url)
    ElMessage.success('URL 已复制')
  } catch {
    ElMessage.error('复制失败，请手动复制 URL')
  }
}

const openURL = (url: string) => {
  window.open(url, '_blank', 'noopener,noreferrer')
}

const openRename = (row: MediaAsset) => {
  renaming.value = row
  renameValue.value = displayName(row)
  renameVisible.value = true
}

const saveRename = async () => {
  if (!renaming.value) return
  renamingSaving.value = true
  try {
    await updateMediaName(renaming.value.id, renameValue.value)
    ElMessage.success('显示名称已更新')
    renameVisible.value = false
    await load()
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '保存失败')
  } finally {
    renamingSaving.value = false
  }
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

.media-toolbar {
  display: flex;
  align-items: center;
  gap: 10px;
  flex-wrap: wrap;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  background: #ffffff;
  padding: 12px;
}

.media-toolbar .el-input {
  width: 320px;
}

.media-toolbar .el-select {
  width: 150px;
}

.toolbar-spacer {
  flex: 1;
  min-width: 16px;
}

.toolbar-count {
  color: #6b7280;
  font-size: 13px;
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

.file-kind {
  width: 70px;
  height: 48px;
  border: 1px solid #e5e7eb;
  border-radius: 6px;
  background: #f9fafb;
  color: #4b5563;
  cursor: pointer;
}

.file-kind:hover {
  border-color: #7d1231;
  color: #7d1231;
}

.media-name-cell {
  display: flex;
  flex-direction: column;
  gap: 4px;
  min-width: 0;
}

.media-name-cell strong {
  color: #25313b;
  font-weight: 600;
}

.media-name-cell span,
.muted-text {
  color: #6b7280;
  font-size: 12px;
  word-break: break-all;
}

.pager {
  align-self: flex-end;
}

.media-table {
  border-radius: 8px;
  overflow: hidden;
}

@media (max-width: 760px) {
  .admin-page-header {
    align-items: flex-start;
    flex-direction: column;
  }

  .media-toolbar .el-input,
  .media-toolbar .el-select {
    width: 100%;
  }

  .toolbar-spacer,
  .toolbar-count {
    display: none;
  }
}
</style>
