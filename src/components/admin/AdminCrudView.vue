<template>
  <div ref="crudRoot" class="admin-crud">
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
        <el-option label="已发布" value="published" />
        <el-option label="草稿" value="draft" />
      </el-select>
      <el-button @click="load">刷新</el-button>
    </div>

    <el-alert v-if="sortable" class="sort-hint" type="info" :closable="false" show-icon>
      <template #title>按住左侧手柄拖动，可调整当前列表的展示顺序；跨年份或跨分类的拖动会被自动拦截。</template>
    </el-alert>

    <el-table
      v-loading="loading"
      :data="items"
      border
      class="admin-table"
      row-key="id"
      :row-class-name="rowClassName"
    >
      <el-table-column v-if="sortable" label="" width="54">
        <template #default="{ row }">
          <button
            class="drag-handle"
            type="button"
            draggable="true"
            aria-label="拖动排序"
            title="拖动排序"
            @dragstart="handleDragStart(row, $event)"
            @dragend="handleDragEnd"
          >
            ⋮⋮
          </button>
        </template>
      </el-table-column>
      <el-table-column v-for="column in columns" :key="column.prop" :prop="column.prop" :label="column.label" :min-width="column.width || 120">
        <template #default="{ row }">
          <el-image v-if="column.type === 'image' && row[column.prop]" :src="row[column.prop]" fit="cover" class="table-image" />
          <span v-else-if="column.type === 'image'" class="empty-image">未设置</span>
          <el-tag v-else-if="column.prop === 'status'" :type="row.status === 'published' ? 'success' : 'info'">{{ formatStatus(row.status) }}</el-tag>
          <span v-else>{{ formatCell(row[column.prop], column) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" fixed="right" width="210">
        <template #default="{ row }">
          <el-button link type="primary" @click="openEdit(row)">编辑</el-button>
          <el-button
            link
            :type="row.status === 'published' ? 'warning' : 'success'"
            :loading="togglingID === row.id"
            @click="toggleStatus(row)"
          >
            {{ row.status === 'published' ? '隐藏' : '发布' }}
          </el-button>
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

    <el-dialog v-model="dialogVisible" :title="editing?.id ? '编辑' : '新增'" width="860px">
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
          <div v-else-if="field.type === 'image'" class="media-picker">
            <el-image v-if="editing[field.prop]" :src="editing[field.prop]" fit="cover" class="field-image-preview" />
            <div v-else class="field-image-empty">未选择</div>
            <div class="media-picker-controls">
              <el-input v-model="editing[field.prop]" placeholder="可粘贴 URL，也可从媒体库选择" />
              <div class="media-picker-buttons">
                <el-button @click="openMediaPicker(field.prop, 'image')">选择媒体</el-button>
                <el-upload accept="image/*" :show-file-list="false" :http-request="(options: UploadRequestOptions) => uploadFieldMedia(options, field.prop)">
                  <el-button :loading="mediaUploading">上传并使用</el-button>
                </el-upload>
                <el-button v-if="editing[field.prop]" @click="editing[field.prop] = ''">清空</el-button>
              </div>
            </div>
          </div>
          <div v-else-if="field.type === 'links'" class="links-editor">
            <div class="links-toolbar">
              <el-button v-for="option in linkTypeOptions" :key="option.value" size="small" @click="addPublicationLink(option.value)">
                添加{{ option.shortLabel }}
              </el-button>
            </div>
            <el-empty v-if="!editing.links?.length" description="暂无附件或链接" :image-size="72" />
            <div
              v-for="(link, index) in editing.links"
              :key="link.id || index"
              class="link-card"
              :class="{ 'link-card-over': linkDragOverIndex === index }"
              @dragover.prevent="handleLinkDragOver(index)"
              @drop.prevent="handleLinkDrop(index)"
            >
              <button
                class="drag-handle link-drag-handle"
                type="button"
                draggable="true"
                aria-label="拖动排序"
                title="拖动排序"
                @dragstart="handleLinkDragStart(index, $event)"
                @dragend="handleLinkDragEnd"
              >
                ⋮⋮
              </button>
              <div class="link-fields">
                <div class="link-line">
                  <el-select v-model="link.type" class="link-type" placeholder="类型" @change="handleLinkTypeChange(link)">
                    <el-option v-for="option in linkTypeOptions" :key="option.value" :label="option.label" :value="option.value" />
                  </el-select>
                  <el-input v-model="link.label" class="link-label" placeholder="官网按钮文字" />
                  <el-button link type="danger" @click="removePublicationLink(index)">删除</el-button>
                </div>
                <div class="link-line">
                  <el-input v-model="link.url" :placeholder="linkURLPlaceholder(link)" />
                  <template v-if="linkCanUseMedia(link)">
                    <el-button @click="openLinkMediaPicker(index)">选择文件</el-button>
                    <el-upload :accept="linkUploadAccept(link)" :show-file-list="false" :http-request="(options: UploadRequestOptions) => uploadLinkMedia(options, index)">
                      <el-button :loading="mediaUploading">上传并使用</el-button>
                    </el-upload>
                  </template>
                </div>
                <div v-if="link.type === 'video'" class="link-line">
                  <el-input v-model="link.routeName" placeholder="站内视频页面路由名；如果上传 MP4，可只填写上面的文件 URL" />
                </div>
              </div>
            </div>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="saving" @click="save">保存</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="mediaPickerVisible" title="选择媒体" width="860px">
      <div class="media-picker-toolbar">
        <el-input v-model="mediaQuery" placeholder="按文件名或 URL 搜索" clearable @keyup.enter="loadMediaOptions" />
        <el-button @click="loadMediaOptions">搜索</el-button>
      </div>
      <div v-loading="mediaLoading" class="media-grid">
        <button v-for="asset in mediaOptions" :key="asset.id" class="media-option" type="button" @click="chooseMedia(asset.url)">
          <el-image v-if="asset.kind === 'image'" :src="asset.url" fit="cover" />
          <span v-else class="media-kind">{{ formatMediaKind(asset.kind) }}</span>
          <span class="media-name">{{ asset.originalName }}</span>
        </button>
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, onUpdated, reactive, ref, toRaw } from 'vue'
import type { UploadRequestOptions } from 'element-plus'
import { ElMessage } from 'element-plus'
import { createAdmin, deleteAdmin, listAdmin, placeAdmin, updateAdmin, uploadMedia } from '../../api/admin'
import type { MediaAsset, PublicationLink } from '../../api/client'

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
  sortable?: boolean
  sortGroupKey?: string | string[]
}>()

const loading = ref(false)
const saving = ref(false)
const togglingID = ref<number | null>(null)
const items = ref<Record<string, any>[]>([])
const total = ref(0)
const dialogVisible = ref(false)
const editing = ref<Record<string, any> | null>(null)
const query = reactive({ page: 1, pageSize: 20, q: '', status: '' })
const draggingRow = ref<Record<string, any> | null>(null)
const dragOverID = ref<number | null>(null)
const dragSaving = ref(false)
const mediaPickerVisible = ref(false)
const mediaLoading = ref(false)
const mediaUploading = ref(false)
const mediaTargetProp = ref('')
const mediaTargetLinkIndex = ref<number | null>(null)
const mediaKind = ref('')
const mediaQuery = ref('')
const mediaOptions = ref<MediaAsset[]>([])
const crudRoot = ref<HTMLElement | null>(null)
const dragEventsBound = ref(false)
const linkDragIndex = ref<number | null>(null)
const linkDragOverIndex = ref<number | null>(null)

const linkTypeOptions = [
  { label: '论文 / PDF', shortLabel: '论文', value: 'paper', defaultLabel: 'Paper' },
  { label: '代码', shortLabel: '代码', value: 'code', defaultLabel: 'Code' },
  { label: '视频', shortLabel: '视频', value: 'video', defaultLabel: 'Video' },
  { label: 'PPT', shortLabel: 'PPT', value: 'ppt', defaultLabel: 'PPT' },
  { label: 'Poster', shortLabel: 'Poster', value: 'poster', defaultLabel: 'Poster' },
]

const sortGroupKeys = computed(() => {
  if (!props.sortGroupKey) return []
  return Array.isArray(props.sortGroupKey) ? props.sortGroupKey : [props.sortGroupKey]
})

const clonePlain = (value: Record<string, any>) => JSON.parse(JSON.stringify(toRaw(value)))

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
  editing.value = clonePlain(props.defaults)
  ensurePublicationLinks()
  dialogVisible.value = true
}

const openEdit = (row: Record<string, any>) => {
  const next = clonePlain(row)
  if (props.resource === 'publications' && !next.links) {
    next.links = []
  }
  editing.value = next
  ensurePublicationLinks()
  dialogVisible.value = true
}

const save = async () => {
  if (!editing.value) return
  normalizeEditingBeforeSave()
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

const toggleStatus = async (row: Record<string, any>) => {
  togglingID.value = row.id
  try {
    const payload = clonePlain(row)
    payload.status = row.status === 'published' ? 'draft' : 'published'
    normalizeLinks(payload)
    await updateAdmin(props.resource, payload)
    ElMessage.success(payload.status === 'published' ? '已发布' : '已隐藏')
    await load()
  } finally {
    togglingID.value = null
  }
}

const remove = async (row: Record<string, any>) => {
  await deleteAdmin(props.resource, row.id)
  ElMessage.success('已移入回收站')
  await load()
}

const rowClassName = ({ row }: { row: Record<string, any> }) => {
  if (!props.sortable) return ''
  return ['admin-sort-row', `admin-sort-row-${row.id}`, dragOverID.value === row.id ? 'drag-over-row' : ''].filter(Boolean).join(' ')
}

const sameSortGroup = (source: Record<string, any>, target: Record<string, any>) => {
  return sortGroupKeys.value.every((key) => source[key] === target[key])
}

const handleDragStart = (row: Record<string, any>, event: DragEvent) => {
  if (!props.sortable || dragSaving.value) return
  event.dataTransfer?.setData('text/plain', String(row.id))
  if (event.dataTransfer) {
    event.dataTransfer.effectAllowed = 'move'
  }
  draggingRow.value = row
}

const handleDragEnd = () => {
  draggingRow.value = null
  dragOverID.value = null
}

const rowFromDragEvent = (event: DragEvent) => {
  const rowElement = (event.target as HTMLElement | null)?.closest('.el-table__body tbody tr')
  if (!rowElement || !crudRoot.value?.contains(rowElement)) return null
  const idClass = Array.from(rowElement.classList).find((className) => className.startsWith('admin-sort-row-'))
  const rowID = idClass ? Number(idClass.replace('admin-sort-row-', '')) : 0
  return items.value.find((item) => item.id === rowID) || null
}

const bindTableDragEvents = () => {
  if (!props.sortable || dragEventsBound.value || !crudRoot.value) return
  dragEventsBound.value = true
  crudRoot.value.addEventListener('dragover', (event) => {
    const row = rowFromDragEvent(event as DragEvent)
    if (!row || !draggingRow.value || draggingRow.value.id === row.id) return
    event.preventDefault()
    dragOverID.value = row.id
  })
  crudRoot.value.addEventListener('dragleave', (event) => {
    const row = rowFromDragEvent(event as DragEvent)
    if (row && dragOverID.value === row.id) {
      dragOverID.value = null
    }
  })
  crudRoot.value.addEventListener('drop', async (event) => {
    const row = rowFromDragEvent(event as DragEvent)
    if (!row) return
    event.preventDefault()
    const rowElement = (event.target as HTMLElement | null)?.closest('.el-table__body tbody tr')
    const rect = rowElement?.getBoundingClientRect()
    const position = rect && event instanceof DragEvent && event.clientY > rect.top + rect.height / 2 ? 'after' : 'before'
    await handleRowDrop(row, position)
  })
}

const handleRowDrop = async (row: Record<string, any>, position: 'before' | 'after') => {
  const source = draggingRow.value
  dragOverID.value = null
  draggingRow.value = null
  if (!source || source.id === row.id || dragSaving.value) return
  if (!sameSortGroup(source, row)) {
    ElMessage.warning('只能在同一展示分组内排序')
    return
  }
  dragSaving.value = true
  try {
    await placeAdmin(props.resource, source.id, row.id, position)
    ElMessage.success('顺序已更新')
    await load()
  } finally {
    dragSaving.value = false
  }
}

defineExpose({ load })

const openMediaPicker = async (prop: string, kind = '') => {
  mediaTargetProp.value = prop
  mediaTargetLinkIndex.value = null
  mediaKind.value = kind
  mediaPickerVisible.value = true
  await loadMediaOptions()
}

const openLinkMediaPicker = async (index: number) => {
  const link = editing.value?.links?.[index]
  if (!link) return
  mediaTargetProp.value = ''
  mediaTargetLinkIndex.value = index
  mediaKind.value = linkMediaKind(link)
  mediaPickerVisible.value = true
  await loadMediaOptions()
}

const loadMediaOptions = async () => {
  mediaLoading.value = true
  try {
    const result = await listAdmin<MediaAsset>('media', { page: 1, pageSize: 48, q: mediaQuery.value, kind: mediaKind.value })
    mediaOptions.value = result.items
  } finally {
    mediaLoading.value = false
  }
}

const chooseMedia = (url: string) => {
  if (!editing.value) return
  if (mediaTargetLinkIndex.value !== null) {
    const link = editing.value.links?.[mediaTargetLinkIndex.value]
    if (link) {
      link.url = url
    }
  } else if (mediaTargetProp.value) {
    editing.value[mediaTargetProp.value] = url
  }
  mediaPickerVisible.value = false
}

const uploadFieldMedia = async (options: UploadRequestOptions, prop: string) => {
  if (!editing.value) return
  mediaUploading.value = true
  try {
    const asset = await uploadMedia(options.file)
    editing.value[prop] = asset.url
    ElMessage.success('已上传并填入')
  } finally {
    mediaUploading.value = false
  }
}

const uploadLinkMedia = async (options: UploadRequestOptions, index: number) => {
  if (!editing.value?.links?.[index]) return
  mediaUploading.value = true
  try {
    const asset = await uploadMedia(options.file)
    editing.value.links[index].url = asset.url
    ElMessage.success('已上传并填入')
  } finally {
    mediaUploading.value = false
  }
}

const defaultLinkLabel = (type: string) => linkTypeOptions.find((option) => option.value === type)?.defaultLabel || 'Link'

const addPublicationLink = (type = 'paper') => {
  if (!editing.value) return
  if (!Array.isArray(editing.value.links)) {
    editing.value.links = []
  }
  editing.value.links.push({ type, label: defaultLinkLabel(type), url: '', routeName: '', sortOrder: editing.value.links.length + 1 })
}

const removePublicationLink = (index: number) => {
  editing.value?.links?.splice(index, 1)
}

const handleLinkTypeChange = (link: PublicationLink) => {
  const defaults = linkTypeOptions.map((option) => option.defaultLabel)
  if (!link.label || defaults.includes(link.label)) {
    link.label = defaultLinkLabel(link.type)
  }
  if (link.type !== 'video') {
    link.routeName = ''
  }
}

const handleLinkDragStart = (index: number, event: DragEvent) => {
  linkDragIndex.value = index
  event.dataTransfer?.setData('text/plain', String(index))
  if (event.dataTransfer) {
    event.dataTransfer.effectAllowed = 'move'
  }
}

const handleLinkDragOver = (index: number) => {
  if (linkDragIndex.value === null || linkDragIndex.value === index) return
  linkDragOverIndex.value = index
}

const handleLinkDrop = (index: number) => {
  if (!editing.value?.links || linkDragIndex.value === null || linkDragIndex.value === index) {
    handleLinkDragEnd()
    return
  }
  const links = editing.value.links
  const [moved] = links.splice(linkDragIndex.value, 1)
  links.splice(index, 0, moved)
  normalizeLinks(editing.value)
  handleLinkDragEnd()
}

const handleLinkDragEnd = () => {
  linkDragIndex.value = null
  linkDragOverIndex.value = null
}

const ensurePublicationLinks = () => {
  if (props.resource !== 'publications' || !editing.value) return
  if (!Array.isArray(editing.value.links)) {
    editing.value.links = []
  }
  normalizeLinks(editing.value)
}

const normalizeEditingBeforeSave = () => {
  if (!editing.value) return
  normalizeLinks(editing.value)
}

const normalizeLinks = (value: Record<string, any>) => {
  if (!Array.isArray(value.links)) return
  value.links = value.links.map((link: PublicationLink, index: number) => ({
    ...link,
    label: link.label || defaultLinkLabel(link.type),
    routeName: link.type === 'video' ? link.routeName || '' : '',
    sortOrder: index + 1,
  }))
}

const linkMediaKind = (link: PublicationLink) => {
  if (link.type === 'video') return 'video'
  if (link.type === 'paper' || link.type === 'ppt') return 'document'
  return ''
}

const linkCanUseMedia = (link: PublicationLink) => link.type === 'paper' || link.type === 'ppt' || link.type === 'poster' || link.type === 'video'

const linkUploadAccept = (link: PublicationLink) => {
  if (link.type === 'paper' || link.type === 'poster') return '.pdf,application/pdf,image/*'
  if (link.type === 'ppt') return '.ppt,.pptx,application/vnd.ms-powerpoint,application/vnd.openxmlformats-officedocument.presentationml.presentation'
  if (link.type === 'video') return '.mp4,video/mp4'
  return ''
}

const linkURLPlaceholder = (link: PublicationLink) => {
  if (link.type === 'code') return 'GitHub、项目主页或其他外部链接'
  if (link.type === 'video') return '视频文件 URL 或外部链接'
  return '可粘贴外部链接，也可选择/上传文件'
}

const splitList = (value: string) => value.split('\n').map((item) => item.trim()).filter(Boolean)

const fieldOptions = computed(() => {
  const optionsByProp = new Map<string, Map<string | number, string>>()
  props.fields.forEach((field) => {
    if (!field.options) return
    optionsByProp.set(field.prop, new Map(field.options.map((option) => [option.value, option.label])))
  })
  return optionsByProp
})

const formatCell = (value: unknown, column?: FieldConfig) => {
  if (column) {
    const optionLabel = fieldOptions.value.get(column.prop)?.get(value as string | number)
    if (optionLabel) return optionLabel
  }
  if (Array.isArray(value)) return value.join('；')
  if (typeof value === 'boolean') return value ? '是' : '否'
  return value ?? ''
}

const formatStatus = (status: string) => status === 'published' ? '已发布' : '草稿'

const formatMediaKind = (kind: string) => {
  const labels: Record<string, string> = {
    image: '图片',
    document: '文档',
    archive: '压缩包',
    video: '视频',
  }
  return labels[kind] || kind
}

onMounted(() => {
  bindTableDragEvents()
  void load()
})

onUpdated(bindTableDragEvents)
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

.sort-hint {
  --el-alert-padding: 8px 12px;
}

.drag-handle {
  width: 28px;
  height: 28px;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  background: #ffffff;
  color: #6b7280;
  cursor: grab;
  font-weight: 700;
  line-height: 1;
}

.drag-handle:active {
  cursor: grabbing;
}

.admin-table :deep(.drag-over-row td) {
  background: #fff7ed;
}

.table-image {
  width: 72px;
  height: 48px;
  border-radius: 6px;
}

.empty-image {
  color: #9ca3af;
  font-size: 12px;
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

.links-toolbar {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.link-card {
  display: grid;
  grid-template-columns: 34px 1fr;
  gap: 10px;
  align-items: start;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  background: #ffffff;
  padding: 10px;
}

.link-card-over {
  border-color: #f59e0b;
  background: #fff7ed;
}

.link-drag-handle {
  margin-top: 2px;
}

.link-fields {
  display: flex;
  flex-direction: column;
  gap: 8px;
  min-width: 0;
}

.link-line {
  display: flex;
  align-items: center;
  gap: 8px;
  min-width: 0;
}

.link-type {
  width: 150px;
  flex-shrink: 0;
}

.link-label {
  width: 160px;
  flex-shrink: 0;
}

.media-picker {
  display: grid;
  grid-template-columns: 112px 1fr;
  gap: 12px;
  width: 100%;
}

.field-image-preview,
.field-image-empty {
  width: 112px;
  height: 76px;
  border-radius: 6px;
  border: 1px solid #e5e7eb;
  background: #f9fafb;
}

.field-image-empty {
  display: flex;
  align-items: center;
  justify-content: center;
  color: #9ca3af;
  font-size: 12px;
}

.media-picker-controls {
  display: flex;
  flex-direction: column;
  gap: 8px;
  min-width: 0;
}

.media-picker-buttons,
.media-picker-toolbar {
  display: flex;
  align-items: center;
  gap: 8px;
}

.media-picker-toolbar {
  margin-bottom: 14px;
}

.media-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
  gap: 12px;
  min-height: 180px;
  max-height: 58vh;
  overflow: auto;
}

.media-option {
  display: flex;
  flex-direction: column;
  gap: 8px;
  min-width: 0;
  border: 1px solid #e5e7eb;
  border-radius: 6px;
  background: #ffffff;
  padding: 8px;
  cursor: pointer;
  text-align: left;
}

.media-option:hover {
  border-color: #7d1231;
}

.media-option .el-image,
.media-kind {
  width: 100%;
  aspect-ratio: 4 / 3;
  border-radius: 4px;
  background: #f3f4f6;
}

.media-kind {
  display: flex;
  align-items: center;
  justify-content: center;
  color: #6b7280;
}

.media-name {
  overflow: hidden;
  color: #374151;
  font-size: 12px;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
