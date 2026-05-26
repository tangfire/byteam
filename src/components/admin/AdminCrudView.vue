<template>
  <div ref="crudRoot" class="admin-crud">
    <div class="admin-page-header">
      <div>
        <h1>{{ title }}</h1>
        <p>{{ description }}</p>
      </div>
      <el-button type="primary" @click="openCreate">新增{{ title.replace('管理', '') }}</el-button>
    </div>

    <div class="admin-panel">
      <div class="admin-toolbar">
        <el-input v-model="query.q" placeholder="搜索标题、姓名、编号等" clearable @clear="refresh" @keyup.enter="refresh" />
        <el-segmented v-model="query.status" :options="statusOptions" @change="refresh" />
        <div class="toolbar-spacer" />
        <span class="toolbar-count">共 {{ total }} 条</span>
        <el-button :loading="loading" @click="refresh">刷新</el-button>
      </div>

      <div v-if="sortable" class="sort-hint">
        <span class="drag-hint-icon">⋮⋮</span>
        <span>拖动每行左侧手柄即可调整展示顺序，同一年份或同一分类内排序会自动保存。</span>
      </div>
    </div>

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
            :disabled="dragSaving"
            @dragstart="handleDragStart(row, $event)"
            @dragend="handleDragEnd"
          >
            ⋮⋮
          </button>
        </template>
      </el-table-column>
      <el-table-column v-for="column in columns" :key="column.prop" :prop="column.prop" :label="column.label" :min-width="column.width || 120">
        <template #default="{ row }">
          <span v-if="column.type === 'avatar' && row[column.prop]" class="table-avatar">
            <img :src="row[column.prop]" :alt="row.name || '头像'" :style="avatarImageStyle(row)">
          </span>
          <el-image v-else-if="column.type === 'image' && row[column.prop]" :src="row[column.prop]" fit="cover" class="table-image" />
          <span v-else-if="column.type === 'image' || column.type === 'avatar'" class="empty-image">未设置</span>
          <el-tag v-else-if="column.prop === 'status'" :type="row.status === 'published' ? 'success' : 'info'">{{ formatStatus(row.status) }}</el-tag>
          <el-tag v-else-if="column.prop === 'featured' && row[column.prop]" type="warning">精选</el-tag>
          <span v-else-if="column.prop === 'featured'">-</span>
          <span v-else>{{ formatCell(row[column.prop], column) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" fixed="right" width="220">
        <template #default="{ row }">
          <el-button link type="primary" @click="openEdit(row)">编辑</el-button>
          <el-popconfirm
            :title="row.status === 'published' ? '确认隐藏这条内容？隐藏后前台将不再显示。' : '确认发布这条内容？发布后前台将显示。'"
            @confirm="toggleStatus(row)"
          >
            <template #reference>
              <el-button
                link
                :type="row.status === 'published' ? 'warning' : 'success'"
                :loading="togglingID === row.id"
              >
                {{ row.status === 'published' ? '隐藏' : '发布' }}
              </el-button>
            </template>
          </el-popconfirm>
          <el-popconfirm title="确认移入回收站？之后可在回收站恢复。" @confirm="remove(row)">
            <template #reference>
              <el-button link type="danger">移入回收站</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
      <template #empty>
        <el-empty :description="emptyDescription" :image-size="92">
          <el-button v-if="hasActiveFilters" @click="resetFilters">清空筛选</el-button>
          <el-button v-else type="primary" @click="openCreate">新增{{ title.replace('管理', '') }}</el-button>
        </el-empty>
      </template>
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

    <el-dialog
      v-model="dialogVisible"
      :title="editing?.id ? `编辑${title.replace('管理', '')}` : `新增${title.replace('管理', '')}`"
      width="900px"
      class="admin-edit-dialog"
      :before-close="beforeDialogClose"
      @closed="handleDialogClosed"
    >
      <el-form :model="editing" label-width="120px" class="admin-form" v-if="editing">
        <el-form-item v-for="field in fields" :key="field.prop" :label="field.label">
          <div v-if="field.type === 'color'" class="color-field">
            <el-color-picker v-model="editing[field.prop]" />
            <el-input v-model="editing[field.prop]" placeholder="#7d1231" />
          </div>
          <el-input v-else-if="!field.type || field.type === 'text'" v-model="editing[field.prop]" />
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
                <el-popconfirm
                  v-if="editing[field.prop]"
                  title="确认清空这个媒体地址？保存后前台将不再显示这个资源。"
                  @confirm="editing[field.prop] = ''"
                >
                  <template #reference>
                    <el-button>清空</el-button>
                  </template>
                </el-popconfirm>
              </div>
            </div>
          </div>
          <div v-else-if="field.type === 'avatarCrop'" class="avatar-crop-editor">
            <div
              class="avatar-crop-preview"
              @pointerdown.prevent="startAvatarDrag"
            >
              <img
                v-if="editing.avatarUrl"
                :src="editing.avatarUrl"
                :alt="editing.name || '头像预览'"
                :style="avatarImageStyle(editing)"
                draggable="false"
              >
              <div v-else class="avatar-crop-empty">请先选择头像</div>
            </div>
            <div class="avatar-crop-controls">
              <div class="avatar-crop-control">
                <span>水平位置</span>
                <el-slider v-model="editing.avatarObjectX" :min="0" :max="100" />
                <el-input-number v-model="editing.avatarObjectX" :min="0" :max="100" size="small" />
              </div>
              <div class="avatar-crop-control">
                <span>垂直位置</span>
                <el-slider v-model="editing.avatarObjectY" :min="0" :max="100" />
                <el-input-number v-model="editing.avatarObjectY" :min="0" :max="100" size="small" />
              </div>
              <div class="avatar-crop-control">
                <span>缩放</span>
                <el-slider v-model="editing.avatarScale" :min="100" :max="200" />
                <el-input-number v-model="editing.avatarScale" :min="100" :max="200" size="small" />
              </div>
              <el-button size="small" @click="resetAvatarCrop">居中显示</el-button>
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
                  <el-popconfirm title="确认删除这个附件/链接？保存后前台将不再显示。" @confirm="removePublicationLink(index)">
                    <template #reference>
                      <el-button link type="danger">删除</el-button>
                    </template>
                  </el-popconfirm>
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
                  <el-select
                    class="video-page-select"
                    :model-value="selectedVideoSlug(link)"
                    filterable
                    placeholder="选择站内视频页"
                    @focus="() => loadVideoPages()"
                    @change="selectVideoPage(link, $event)"
                  >
                    <el-option v-for="page in videoPageOptions" :key="page.slug" :label="page.title || page.slug" :value="page.slug">
                      <div class="video-page-option">
                        <span>{{ page.title || page.slug }}</span>
                        <small>{{ videoPagePath(page.slug) }}</small>
                      </div>
                    </el-option>
                  </el-select>
                  <el-button :loading="videoPagesLoading" @click="() => loadVideoPages()">刷新视频页</el-button>
                  <el-button link type="primary" @click="openAdminPages">去创建视频页</el-button>
                </div>
                <p v-if="link.type === 'video'" class="link-help">
                  站内视频页会在前台打开播放页；如果这里只填 MP4 或外部地址，前台会直接打开这个地址。
                  <span v-if="!videoPagesLoaded">点击选择框会加载后台已有视频页。</span>
                </p>
                <div v-if="link.type === 'video' && link.routeName" class="link-line legacy-route-line">
                  <el-tag type="warning">旧路由已兼容：{{ link.routeName }}</el-tag>
                </div>
              </div>
            </div>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="requestCloseDialog">取消</el-button>
        <el-button v-if="editing && editing.status !== 'draft'" :loading="saving" @click="saveWithStatus('draft')">存为草稿</el-button>
        <el-button v-if="editing && editing.status !== 'published'" :loading="saving" @click="saveWithStatus('published')">保存并发布</el-button>
        <el-button type="primary" :loading="saving" @click="save">保存</el-button>
      </template>
    </el-dialog>

    <AdminMediaPicker
      v-model="mediaPickerVisible"
      :kind="mediaKind"
      show-original-name
      @choose="chooseMedia"
    />
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, onUpdated, reactive, ref, toRaw } from 'vue'
import type { UploadRequestOptions } from 'element-plus'
import { ElMessage, ElMessageBox } from 'element-plus'
import { createAdmin, deleteAdmin, listAdmin, listSitePages, placeAdmin, updateAdmin, uploadMedia } from '../../api/admin'
import type { PublicationLink, SitePage } from '../../api/client'
import AdminMediaPicker from './AdminMediaPicker.vue'
import { formatStatus } from '../../utils/adminFormat'
import {
  defaultLinkLabel,
  defaultPublicationLinkLabels,
  linkCanUseMedia,
  linkMediaKind,
  linkTypeOptions,
  linkUploadAccept,
  linkURLPlaceholder,
  normalizePublicationLinks,
} from '../../utils/adminPublicationLinks'
import { openExternalLink } from '../../utils/links'
import { legacyVideoRouteSlugs, videoPagePath, videoSlugFromPath } from '../../utils/videoLinks'
import { avatarDisplayStyle } from '../../utils/avatarDisplay'

export interface FieldConfig {
  prop: string
  label: string
  width?: number
  type?: 'text' | 'textarea' | 'number' | 'select' | 'date' | 'boolean' | 'list' | 'links' | 'image' | 'avatar' | 'color' | 'avatarCrop'
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
const editingSnapshot = ref('')
const skipCloseGuard = ref(false)
const query = reactive({ page: 1, pageSize: 20, q: '', status: '' })
const draggingRow = ref<Record<string, any> | null>(null)
const dragOverID = ref<number | null>(null)
const dragSaving = ref(false)
const mediaPickerVisible = ref(false)
const mediaUploading = ref(false)
const mediaTargetProp = ref('')
const mediaTargetLinkIndex = ref<number | null>(null)
const mediaKind = ref('')
const videoPageOptions = ref<SitePage[]>([])
const videoPagesLoaded = ref(false)
const videoPagesLoading = ref(false)
const crudRoot = ref<HTMLElement | null>(null)
const dragEventsBound = ref(false)
const linkDragIndex = ref<number | null>(null)
const linkDragOverIndex = ref<number | null>(null)
const avatarDragState = ref<{ startX: number; startY: number; objectX: number; objectY: number; rect: DOMRect } | null>(null)

const statusOptions = [
  { label: '全部', value: '' },
  { label: '已发布', value: 'published' },
  { label: '草稿', value: 'draft' },
]

const sortGroupKeys = computed(() => {
  if (!props.sortGroupKey) return []
  return Array.isArray(props.sortGroupKey) ? props.sortGroupKey : [props.sortGroupKey]
})

const clonePlain = (value: Record<string, any>) => JSON.parse(JSON.stringify(toRaw(value)))

const serializeEditing = () => editing.value ? JSON.stringify(toRaw(editing.value)) : ''

const hasUnsavedChanges = computed(() => Boolean(editing.value) && serializeEditing() !== editingSnapshot.value)

const hasActiveFilters = computed(() => Boolean(query.q || query.status))

const emptyDescription = computed(() => {
  if (hasActiveFilters.value) return '没有符合当前筛选条件的内容'
  return `还没有${props.title.replace('管理', '')}内容`
})

const load = async () => {
  loading.value = true
  try {
    const result = await listAdmin<Record<string, any>>(props.resource, query)
    items.value = result.items
    total.value = result.total
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '列表加载失败')
  } finally {
    loading.value = false
  }
}

const refresh = async () => {
  query.page = 1
  await load()
}

const resetFilters = async () => {
  query.q = ''
  query.status = ''
  query.page = 1
  await load()
}

const openCreate = () => {
  editing.value = clonePlain(props.defaults)
  ensureAvatarDisplay()
  ensurePublicationLinks()
  editingSnapshot.value = serializeEditing()
  void loadVideoPagesIfNeeded()
  dialogVisible.value = true
}

const openEdit = (row: Record<string, any>) => {
  const next = clonePlain(row)
  if (props.resource === 'publications' && !next.links) {
    next.links = []
  }
  editing.value = next
  ensureAvatarDisplay()
  ensurePublicationLinks()
  editingSnapshot.value = serializeEditing()
  void loadVideoPagesIfNeeded()
  dialogVisible.value = true
}

const confirmDiscardChanges = async () => {
  if (!hasUnsavedChanges.value) return true
  try {
    await ElMessageBox.confirm('当前弹窗有未保存内容，确认关闭并放弃这些修改？', '未保存修改', {
      type: 'warning',
      confirmButtonText: '放弃修改',
      cancelButtonText: '继续编辑',
    })
    return true
  } catch {
    return false
  }
}

const beforeDialogClose = async (done: () => void) => {
  if (skipCloseGuard.value || await confirmDiscardChanges()) {
    done()
  }
}

const requestCloseDialog = async () => {
  if (await confirmDiscardChanges()) {
    skipCloseGuard.value = true
    dialogVisible.value = false
  }
}

const handleDialogClosed = () => {
  skipCloseGuard.value = false
  editingSnapshot.value = ''
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
    skipCloseGuard.value = true
    dialogVisible.value = false
    await load()
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '保存失败')
  } finally {
    saving.value = false
  }
}

const saveWithStatus = async (status: 'draft' | 'published') => {
  if (!editing.value) return
  editing.value.status = status
  await save()
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
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '状态更新失败')
  } finally {
    togglingID.value = null
  }
}

const remove = async (row: Record<string, any>) => {
  try {
    await deleteAdmin(props.resource, row.id)
    ElMessage.success('已移入回收站')
    await load()
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '删除失败')
  }
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
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '排序保存失败')
  } finally {
    dragSaving.value = false
  }
}

defineExpose({ load })

const openMediaPicker = (prop: string, kind = '') => {
  mediaTargetProp.value = prop
  mediaTargetLinkIndex.value = null
  mediaKind.value = kind
  mediaPickerVisible.value = true
}

const openLinkMediaPicker = (index: number) => {
  const link = editing.value?.links?.[index]
  if (!link) return
  mediaTargetProp.value = ''
  mediaTargetLinkIndex.value = index
  mediaKind.value = linkMediaKind(link)
  mediaPickerVisible.value = true
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
    ensureAvatarDisplay()
  }
  mediaPickerVisible.value = false
}

const uploadFieldMedia = async (options: UploadRequestOptions, prop: string) => {
  if (!editing.value) return
  mediaUploading.value = true
  try {
    const asset = await uploadMedia(options.file)
    editing.value[prop] = asset.url
    ensureAvatarDisplay()
    ElMessage.success('已上传并填入')
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '上传失败')
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
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '上传失败')
  } finally {
    mediaUploading.value = false
  }
}

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
  if (!link.label || defaultPublicationLinkLabels.includes(link.label)) {
    link.label = defaultLinkLabel(link.type)
  }
  if (link.type !== 'video') {
    link.routeName = ''
  }
}

const selectedVideoSlug = (link: PublicationLink) => {
  return videoSlugFromPath(link.url) || legacyVideoRouteSlugs[link.routeName || ''] || ''
}

const selectVideoPage = (link: PublicationLink, slug: string) => {
  if (slug) {
    link.url = videoPagePath(slug)
    link.routeName = ''
    if (!link.label) link.label = defaultLinkLabel('video')
  } else if (videoSlugFromPath(link.url)) {
    link.url = ''
  }
}

const loadVideoPagesIfNeeded = async () => {
  if (props.resource !== 'publications') return
  if (videoPagesLoaded.value && videoPageOptions.value.length) return
  await loadVideoPages()
}

const loadVideoPages = async () => {
  if (videoPagesLoading.value) return
  videoPagesLoading.value = true
  try {
    const result = await listSitePages({ page: 1, pageSize: 100, q: 'video-' })
    videoPageOptions.value = result.items.filter((page) => page.slug.startsWith('video-'))
    videoPagesLoaded.value = true
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '视频页加载失败')
  } finally {
    videoPagesLoading.value = false
  }
}

const openAdminPages = () => {
  openExternalLink(`${window.location.origin}${window.location.pathname}#/admin/pages`)
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
  ensureAvatarDisplay()
  normalizeLinks(editing.value)
}

const normalizeLinks = (value: Record<string, any>) => {
  if (!Array.isArray(value.links)) return
  value.links = normalizePublicationLinks(value.links)
}

const splitList = (value: string) => value.split('\n').map((item) => item.trim()).filter(Boolean)

const clampNumber = (value: unknown, min: number, max: number, fallback: number) => {
  const numberValue = Number(value)
  if (!Number.isFinite(numberValue)) return fallback
  return Math.min(max, Math.max(min, Math.round(numberValue)))
}

const ensureAvatarDisplay = () => {
  if (!editing.value || !('avatarUrl' in editing.value)) return
  editing.value.avatarObjectX = clampNumber(editing.value.avatarObjectX, 0, 100, 50)
  editing.value.avatarObjectY = clampNumber(editing.value.avatarObjectY, 0, 100, 50)
  editing.value.avatarScale = clampNumber(editing.value.avatarScale, 100, 200, 100)
}

const avatarImageStyle = (record: Record<string, any>) => {
  return avatarDisplayStyle(record)
}

const resetAvatarCrop = () => {
  if (!editing.value) return
  editing.value.avatarObjectX = 50
  editing.value.avatarObjectY = 50
  editing.value.avatarScale = 100
}

const updateAvatarPositionFromPointer = (event: PointerEvent) => {
  if (!editing.value || !avatarDragState.value) return
  const state = avatarDragState.value
  const nextX = state.objectX - ((event.clientX - state.startX) / state.rect.width) * 100
  const nextY = state.objectY - ((event.clientY - state.startY) / state.rect.height) * 100
  editing.value.avatarObjectX = clampNumber(nextX, 0, 100, 50)
  editing.value.avatarObjectY = clampNumber(nextY, 0, 100, 50)
}

const stopAvatarDrag = () => {
  avatarDragState.value = null
  window.removeEventListener('pointermove', updateAvatarPositionFromPointer)
  window.removeEventListener('pointerup', stopAvatarDrag)
}

const startAvatarDrag = (event: PointerEvent) => {
  if (!editing.value?.avatarUrl) return
  const target = event.currentTarget as HTMLElement | null
  const rect = target?.getBoundingClientRect()
  if (!rect) return
  ensureAvatarDisplay()
  avatarDragState.value = {
    startX: event.clientX,
    startY: event.clientY,
    objectX: editing.value.avatarObjectX,
    objectY: editing.value.avatarObjectY,
    rect,
  }
  window.addEventListener('pointermove', updateAvatarPositionFromPointer)
  window.addEventListener('pointerup', stopAvatarDrag)
}

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

onMounted(() => {
  bindTableDragEvents()
  void load()
})

onUnmounted(stopAvatarDrag)

onUpdated(bindTableDragEvents)
</script>

<style scoped>
.admin-crud {
  display: flex;
  flex-direction: column;
  gap: 16px;
  min-width: 0;
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
  line-height: 1.25;
}

.admin-page-header p {
  margin: 6px 0 0;
  color: #6b7280;
}

.admin-toolbar {
  justify-content: flex-start;
  flex-wrap: wrap;
}

.admin-toolbar .el-input {
  width: 320px;
}

.admin-panel {
  display: flex;
  flex-direction: column;
  gap: 10px;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  background: #ffffff;
  padding: 12px;
}

.toolbar-spacer {
  flex: 1;
  min-width: 16px;
}

.toolbar-count {
  color: #6b7280;
  font-size: 13px;
  white-space: nowrap;
}

.admin-table {
  width: 100%;
  border-radius: 8px;
  overflow: hidden;
}

.sort-hint {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #6b7280;
  font-size: 13px;
}

.drag-hint-icon {
  width: 22px;
  height: 22px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border-radius: 6px;
  background: #f3f4f6;
  color: #7d1231;
  font-weight: 700;
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

.drag-handle:disabled {
  cursor: wait;
  opacity: 0.55;
}

.admin-table :deep(.drag-over-row td) {
  background: #fff7ed;
}

.table-image {
  width: 72px;
  height: 48px;
  border-radius: 6px;
}

.table-avatar {
  display: block;
  width: 54px;
  height: 54px;
  overflow: hidden;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  background: #f9fafb;
}

.table-avatar img {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover;
  transform-origin: center;
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

.admin-form :deep(.el-input),
.admin-form :deep(.el-select),
.admin-form :deep(.el-date-editor) {
  width: 100%;
}

.color-field {
  display: grid;
  grid-template-columns: auto minmax(0, 1fr);
  gap: 10px;
  width: 260px;
  align-items: center;
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
  flex-wrap: wrap;
}

.link-type {
  width: 150px;
  flex-shrink: 0;
}

.link-label {
  width: 160px;
  flex-shrink: 0;
}

.video-page-select {
  min-width: min(100%, 360px);
  flex: 1;
}

.video-page-option {
  display: flex;
  flex-direction: column;
  line-height: 1.25;
}

.video-page-option small {
  color: #9ca3af;
  font-size: 11px;
}

.link-help {
  margin: -2px 0 0;
  color: #6b7280;
  font-size: 12px;
  line-height: 1.5;
}

.legacy-route-line {
  margin-top: -2px;
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

.media-picker-buttons {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
}

.avatar-crop-editor {
  display: grid;
  grid-template-columns: 180px minmax(0, 1fr);
  gap: 18px;
  width: 100%;
  align-items: start;
}

.avatar-crop-preview {
  width: 180px;
  height: 180px;
  overflow: hidden;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  background: #f9fafb;
  cursor: move;
  user-select: none;
  touch-action: none;
}

.avatar-crop-preview img {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover;
  transform-origin: center;
}

.avatar-crop-empty {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #9ca3af;
  font-size: 13px;
}

.avatar-crop-controls {
  display: flex;
  flex-direction: column;
  gap: 10px;
  min-width: 0;
}

.avatar-crop-control {
  display: grid;
  grid-template-columns: 72px minmax(0, 1fr) 108px;
  gap: 10px;
  align-items: center;
}

.avatar-crop-control span {
  color: #4b5563;
  font-size: 13px;
}

@media (max-width: 760px) {
  .admin-page-header {
    align-items: flex-start;
    flex-direction: column;
  }

  .admin-toolbar .el-input {
    width: 100%;
  }

  .toolbar-spacer,
  .toolbar-count {
    display: none;
  }

  .media-picker {
    grid-template-columns: 1fr;
  }

  .avatar-crop-editor {
    grid-template-columns: 1fr;
  }

  .avatar-crop-control {
    grid-template-columns: 1fr;
  }
}
</style>
