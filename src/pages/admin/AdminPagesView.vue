<template>
  <div class="pages-admin">
    <div class="admin-page-header">
      <div>
        <h1>页面内容</h1>
        <p>只维护仍需要后台编辑的长页面和论文视频页；About、Contact、VideoMind、vKnow 已恢复为源码硬编码。</p>
      </div>
      <div class="header-actions">
        <el-button type="primary" @click="openCreateVideoPage">新增视频页</el-button>
        <el-button :loading="loading" @click="loadPages">刷新</el-button>
      </div>
    </div>

    <div class="page-layout">
      <aside class="page-list">
        <button
          v-for="item in pages"
          :key="item.slug"
          type="button"
          class="page-item"
          :class="{ active: item.slug === activeSlug }"
          @click="selectPage(item.slug)"
        >
          <strong>{{ pageLabel(item.slug, item.title) }}</strong>
          <span>{{ item.slug }}</span>
          <el-tag size="small" :type="item.status === 'published' ? 'success' : 'info'">{{ formatStatus(item.status) }}</el-tag>
        </button>
      </aside>

      <section v-loading="loadingPage" class="editor-panel">
        <el-empty v-if="!editing" description="请选择一个页面" />
        <template v-else>
          <div class="editor-header">
            <div>
              <h2>{{ pageLabel(editing.slug, editing.title) }}</h2>
              <p>{{ editing.slug }}</p>
            </div>
            <div class="editor-actions">
              <el-tag :type="editing.status === 'published' ? 'success' : 'info'">{{ formatStatus(editing.status) }}</el-tag>
              <el-switch
                v-model="editing.status"
                active-text="发布"
                inactive-text="草稿"
                active-value="published"
                inactive-value="draft"
              />
              <el-button type="primary" :loading="saving" @click="savePage">保存</el-button>
              <el-popconfirm
                v-if="editing.slug.startsWith('video-')"
                title="确认移入回收站？之后可在回收站恢复。"
                @confirm="deleteCurrentPage"
              >
                <template #reference>
                  <el-button type="danger" plain :loading="deleting">移入回收站</el-button>
                </template>
              </el-popconfirm>
            </div>
          </div>

          <el-form label-width="110px" class="page-form">
            <el-form-item label="页面标题">
              <el-input v-model="editing.title" />
            </el-form-item>
            <el-form-item label="后台说明">
              <el-input v-model="editing.description" />
            </el-form-item>

            <template v-if="editing.slug === 'research-direction'">
              <EditableList
                :items="content.directions"
                add-label="添加研究方向"
                @add="content.directions.push({ title: '', image: '', alt: '', sections: [] })"
                @remove="removeAt(content.directions, $event)"
              >
                <template #default="{ item }">
                  <el-input v-model="item.title" placeholder="方向标题" />
                  <ImageField v-model="item.image" label="配图" />
                  <el-input v-model="item.alt" placeholder="图片说明" />
                  <EditableList
                    :items="item.sections"
                    add-label="添加文本段"
                    compact
                    @add="item.sections.push({ title: '', text: '' })"
                    @remove="removeAt(item.sections, $event)"
                  >
                    <template #default="{ item: section }">
                      <el-input v-model="section.title" placeholder="小标题，如 Challenges" />
                      <el-input v-model="section.text" type="textarea" :rows="2" placeholder="正文" />
                    </template>
                  </EditableList>
                </template>
              </EditableList>
            </template>

            <template v-else-if="editing.slug === 'dr-baoyao-yang'">
              <el-form-item label="姓名标题">
                <el-input v-model="content.name" />
              </el-form-item>
              <ImageField v-model="content.image" label="照片" />
              <el-form-item label="图片说明">
                <el-input v-model="content.alt" />
              </el-form-item>
              <el-form-item label="简介段落">
                <TextListEditor v-model="content.paragraphs" placeholder="每行一段；较长段落也可以直接粘贴" :rows="8" />
              </el-form-item>
            </template>

            <template v-else-if="editing.slug.startsWith('video-')">
              <el-alert
                type="info"
                :closable="false"
                show-icon
                :title="`前台访问路径：${videoPagePath(editing.slug)}`"
              />
              <el-form-item label="视频标题">
                <el-input v-model="content.title" type="textarea" :rows="2" />
              </el-form-item>
              <ImageField v-model="content.video" label="视频文件" kind="video" />
            </template>

            <template v-else>
              <el-alert type="warning" :closable="false" show-icon title="这个页面暂时没有专用编辑器，可以编辑下面的 JSON。" />
              <el-form-item label="JSON">
                <el-input v-model="rawJSON" type="textarea" :rows="18" @change="applyRawJSON" />
              </el-form-item>
            </template>
          </el-form>
        </template>
      </section>
    </div>

    <el-dialog v-model="mediaPickerVisible" title="选择媒体" width="860px">
      <div class="media-picker-toolbar">
        <el-input v-model="mediaQuery" placeholder="按显示名称、文件名或 URL 搜索" clearable @keyup.enter="loadMediaOptions" />
        <el-button @click="loadMediaOptions">搜索</el-button>
      </div>
      <div v-loading="mediaLoading" class="media-grid">
        <button v-for="asset in mediaOptions" :key="asset.id" class="media-option" type="button" @click="chooseMedia(asset.url)">
          <el-image v-if="asset.kind === 'image'" :src="asset.url" fit="cover" />
          <span v-else class="media-kind">{{ formatMediaKind(asset.kind) }}</span>
          <span class="media-name">{{ asset.displayName || asset.originalName }}</span>
        </button>
      </div>
    </el-dialog>

    <el-dialog v-model="createDialogVisible" title="新增论文视频页" width="640px">
      <el-form :model="newVideoPage" label-width="110px" class="create-video-form">
        <el-form-item label="视频页标题">
          <el-input v-model="newVideoPage.title" placeholder="通常填写论文标题" @input="syncSlugFromTitle" />
        </el-form-item>
        <el-form-item label="Slug">
          <el-input v-model="newVideoPage.slug" placeholder="video-paper-title" @input="slugTouched = true">
            <template #prepend>/video/</template>
          </el-input>
        </el-form-item>
        <el-form-item label="视频文件">
          <ImageField v-model="newVideoPage.video" label="视频文件" kind="video" />
        </el-form-item>
        <el-form-item label="发布状态">
          <el-switch
            v-model="newVideoPage.status"
            active-text="发布"
            inactive-text="草稿"
            active-value="published"
            inactive-value="draft"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="createDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="creating" @click="createVideoPage">创建</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { computed, defineComponent, h, nextTick, onMounted, ref, watch, type PropType } from 'vue'
import { ElButton, ElInput, ElMessage, ElUpload } from 'element-plus'
import type { UploadRequestOptions } from 'element-plus'
import { createSitePage, deleteSitePage, getSitePage, listAdmin, listSitePages, updateSitePage, uploadMedia } from '../../api/admin'
import type { MediaAsset, SitePage } from '../../api/client'
import { videoPagePath } from '../../utils/videoLinks'

const pages = ref<SitePage[]>([])
const activeSlug = ref('')
const editing = ref<SitePage | null>(null)
const loading = ref(false)
const loadingPage = ref(false)
const saving = ref(false)
const deleting = ref(false)
const rawJSON = ref('')
const createDialogVisible = ref(false)
const creating = ref(false)
const slugTouched = ref(false)
const newVideoPage = ref({ title: '', slug: '', video: '', status: 'published' })
const mediaPickerVisible = ref(false)
const mediaLoading = ref(false)
const mediaUploading = ref(false)
const mediaQuery = ref('')
const mediaKind = ref('')
const mediaOptions = ref<MediaAsset[]>([])
const pendingMediaSetter = ref<((url: string) => void) | null>(null)

const sourceManagedSlugs = new Set(['about', 'contact', 'videomind', 'vknow'])

const content = computed<Record<string, any>>(() => editing.value?.content || {})

const fallbackContent: Record<string, () => Record<string, any>> = {
  'research-direction': () => ({ directions: [] }),
  'dr-baoyao-yang': () => ({ name: '', image: '', alt: '', paragraphs: [] }),
}

const pageLabel = (slug: string, title: string) => {
  const labels: Record<string, string> = {
    'research-direction': 'Research Direction',
    'dr-baoyao-yang': 'Baoyao Yang',
    'video-xiaoqi-zheng-01': '视频：Xiaoqi Zheng',
    'video-xianrun-xu-01': '视频：Xianrun Xu',
    'video-yali-ma-01': '视频：Yali Ma',
  }
  return labels[slug] || title
}

const normalizePageContent = (page: SitePage) => {
  const fallback = fallbackContent[page.slug]?.() || (page.slug.startsWith('video-') ? { title: '', video: '' } : {})
  page.content = deepMerge(fallback, page.content || {})
}

const deepMerge = (base: Record<string, any>, value: Record<string, any>) => {
  const out = structuredClone(base)
  Object.entries(value || {}).forEach(([key, nextValue]) => {
    if (isPlainObject(nextValue) && isPlainObject(out[key])) {
      out[key] = deepMerge(out[key], nextValue)
    } else {
      out[key] = nextValue
    }
  })
  return out
}

const isPlainObject = (value: unknown): value is Record<string, any> => Boolean(value) && typeof value === 'object' && !Array.isArray(value)

const loadPages = async () => {
  loading.value = true
  try {
    const result = await listSitePages({ page: 1, pageSize: 50 })
    pages.value = result.items.filter((item) => !sourceManagedSlugs.has(item.slug))
    if (activeSlug.value && !pages.value.some((item) => item.slug === activeSlug.value)) {
      activeSlug.value = ''
      editing.value = null
    }
    if (!activeSlug.value && pages.value.length) {
      await selectPage(pages.value[0].slug)
    }
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '页面列表加载失败')
  } finally {
    loading.value = false
  }
}

const selectPage = async (slug: string) => {
  activeSlug.value = slug
  loadingPage.value = true
  try {
    const page = await getSitePage(slug)
    normalizePageContent(page)
    editing.value = page
    rawJSON.value = JSON.stringify(page.content, null, 2)
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '页面内容加载失败')
  } finally {
    loadingPage.value = false
  }
}

const savePage = async () => {
  if (!editing.value) return
  saving.value = true
  try {
    const saved = await updateSitePage(editing.value.slug, editing.value)
    normalizePageContent(saved)
    editing.value = saved
    rawJSON.value = JSON.stringify(saved.content, null, 2)
    ElMessage.success('页面内容已保存')
    await loadPages()
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '保存失败')
  } finally {
    saving.value = false
  }
}

const openCreateVideoPage = () => {
  newVideoPage.value = { title: '', slug: '', video: '', status: 'published' }
  slugTouched.value = false
  createDialogVisible.value = true
}

const syncSlugFromTitle = () => {
  if (slugTouched.value) return
  newVideoPage.value.slug = slugifyVideoTitle(newVideoPage.value.title)
}

const slugifyVideoTitle = (value: string) => {
  const slug = value
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-+|-+$/g, '')
    .replace(/-{2,}/g, '-')
  return `video-${slug || 'publication'}`
}

const createVideoPage = async () => {
  const slug = newVideoPage.value.slug.trim().toLowerCase()
  if (!newVideoPage.value.title.trim()) {
    ElMessage.warning('请先填写视频页标题')
    return
  }
  if (!/^video-[a-z0-9]+(?:-[a-z0-9]+)*$/.test(slug)) {
    ElMessage.warning('Slug 需要以 video- 开头，只能包含小写字母、数字和横线')
    return
  }

  creating.value = true
  try {
    const saved = await createSitePage({
      slug,
      title: newVideoPage.value.title.trim(),
      description: 'Publication video',
      content: { title: newVideoPage.value.title.trim(), video: newVideoPage.value.video.trim() },
      status: newVideoPage.value.status,
      sortOrder: 0,
    })
    createDialogVisible.value = false
    ElMessage.success('视频页已创建')
    await loadPages()
    await selectPage(saved.slug)
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '创建失败')
  } finally {
    creating.value = false
  }
}

const deleteCurrentPage = async () => {
  if (!editing.value) return
  deleting.value = true
  try {
    const deletedSlug = editing.value.slug
    await deleteSitePage(deletedSlug)
    ElMessage.success('已移入回收站')
    activeSlug.value = ''
    editing.value = null
    await loadPages()
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '删除失败')
  } finally {
    deleting.value = false
  }
}

const applyRawJSON = () => {
  if (!editing.value) return
  try {
    editing.value.content = JSON.parse(rawJSON.value || '{}')
    normalizePageContent(editing.value)
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : 'JSON 格式不正确')
  }
}

const removeAt = (items: unknown[], index: number) => items.splice(index, 1)

const formatStatus = (status: string) => status === 'published' ? '已发布' : '草稿'

const openMediaPicker = async (setter: (url: string) => void, kind = '') => {
  pendingMediaSetter.value = setter
  mediaKind.value = kind
  mediaPickerVisible.value = true
  await loadMediaOptions()
}

const loadMediaOptions = async () => {
  mediaLoading.value = true
  try {
    const result = await listAdmin<MediaAsset>('media', { page: 1, pageSize: 48, q: mediaQuery.value, kind: mediaKind.value, usage: '' })
    mediaOptions.value = result.items
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '媒体加载失败')
  } finally {
    mediaLoading.value = false
  }
}

const chooseMedia = (url: string) => {
  pendingMediaSetter.value?.(url)
  mediaPickerVisible.value = false
}

const uploadAndSet = async (options: UploadRequestOptions, setter: (url: string) => void) => {
  mediaUploading.value = true
  try {
    const asset = await uploadMedia(options.file)
    setter(asset.url)
    ElMessage.success('已上传并填入')
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '上传失败')
  } finally {
    mediaUploading.value = false
  }
}

const formatMediaKind = (kind: string) => {
  const labels: Record<string, string> = { image: '图片', document: '文档', archive: '压缩包', video: '视频' }
  return labels[kind] || kind
}

const TextListEditor = defineComponent({
  props: {
    modelValue: { type: Array as PropType<string[]>, default: () => [] },
    placeholder: { type: String, default: '' },
    rows: { type: Number, default: 5 },
  },
  emits: ['update:modelValue'],
  setup(props, { emit }) {
    const value = ref((props.modelValue || []).join('\n'))
    watch(() => props.modelValue, (next) => {
      value.value = (next || []).join('\n')
    })
    const update = (next: string) => {
      value.value = next
      emit('update:modelValue', next.split('\n').map((item) => item.trim()).filter(Boolean))
    }
    return () => h(ElInput, {
      modelValue: value.value,
      'onUpdate:modelValue': update,
      type: 'textarea',
      rows: props.rows,
      placeholder: props.placeholder,
    })
  },
})

const EditableList = defineComponent({
  props: {
    items: { type: Array as PropType<any[]>, required: true },
    addLabel: { type: String, default: '添加' },
    compact: { type: Boolean, default: false },
  },
  emits: ['add', 'remove'],
  setup(props, { emit, slots }) {
    const dragIndex = ref<number | null>(null)
    const overIndex = ref<number | null>(null)
    const move = (from: number, to: number) => {
      if (to < 0 || to >= props.items.length) return
      const [item] = props.items.splice(from, 1)
      props.items.splice(to, 0, item)
    }
    const handleDrop = (index: number) => {
      if (dragIndex.value === null || dragIndex.value === index) {
        dragIndex.value = null
        overIndex.value = null
        return
      }
      move(dragIndex.value, index)
      dragIndex.value = null
      overIndex.value = null
    }
    return () => h('div', { class: ['editable-list', props.compact ? 'compact' : ''] }, [
      ...props.items.map((item, index) => h('div', {
        class: ['editable-item', overIndex.value === index ? 'drag-over' : ''],
        onDragover: (event: DragEvent) => {
          if (dragIndex.value === null || dragIndex.value === index) return
          event.preventDefault()
          overIndex.value = index
        },
        onDragleave: () => {
          if (overIndex.value === index) overIndex.value = null
        },
        onDrop: (event: DragEvent) => {
          event.preventDefault()
          handleDrop(index)
        },
      }, [
        h('div', { class: 'editable-item-tools' }, [
          h('button', {
            class: 'editable-drag-handle',
            type: 'button',
            draggable: true,
            title: '拖动排序',
            onDragstart: (event: DragEvent) => {
              dragIndex.value = index
              overIndex.value = null
              event.dataTransfer?.setData('text/plain', String(index))
            },
            onDragend: () => {
              dragIndex.value = null
              overIndex.value = null
            },
          }, '⋮⋮'),
          h(ElButton, { size: 'small', text: true, disabled: index === 0, onClick: () => move(index, index - 1) }, () => '上移'),
          h(ElButton, { size: 'small', text: true, disabled: index === props.items.length - 1, onClick: () => move(index, index + 1) }, () => '下移'),
          h(ElButton, { size: 'small', type: 'danger', link: true, onClick: () => emit('remove', index) }, () => '删除'),
        ]),
        h('div', { class: 'editable-item-fields' }, slots.default?.({ item, index })),
      ])),
      h(ElButton, { class: 'editable-add', onClick: () => emit('add') }, () => props.addLabel),
    ])
  },
})

const ImageField = defineComponent({
  props: {
    modelValue: { type: String, default: '' },
    label: { type: String, default: '图片' },
    kind: { type: String, default: 'image' },
  },
  emits: ['update:modelValue'],
  setup(props, { emit }) {
    const set = (url: string) => emit('update:modelValue', url)
    return () => h('div', { class: 'media-field' }, [
      h('label', props.label),
      props.modelValue && props.kind === 'image'
        ? h('img', { src: props.modelValue, alt: '', class: 'media-field-preview' })
        : h('div', { class: 'media-field-empty' }, props.modelValue || '未选择'),
      h(ElInput, { modelValue: props.modelValue, 'onUpdate:modelValue': set, placeholder: '可粘贴 URL，也可选择/上传媒体' }),
      h('div', { class: 'media-field-actions' }, [
        h(ElButton, { onClick: () => openMediaPicker(set, props.kind) }, () => '选择媒体'),
        h(ElUpload, { accept: props.kind === 'video' ? '.mp4,video/mp4' : 'image/*', showFileList: false, httpRequest: (options: UploadRequestOptions) => uploadAndSet(options, set) }, () => h(ElButton, { loading: mediaUploading.value }, () => '上传并使用')),
        props.modelValue ? h(ElButton, { onClick: () => set('') }, () => '清空') : null,
      ]),
    ])
  },
})

onMounted(loadPages)

watch(editing, async () => {
  await nextTick()
  if (editing.value) rawJSON.value = JSON.stringify(editing.value.content, null, 2)
})
</script>

<style scoped>
.pages-admin,
.page-form,
.editable-list,
.editable-item-fields {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.admin-page-header,
.header-actions,
.editor-header,
.editor-actions,
.media-picker-toolbar,
.media-field-actions {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.admin-page-header h1,
.editor-header h2 {
  margin: 0;
  color: #25313b;
}

.admin-page-header p,
.editor-header p {
  margin: 6px 0 0;
  color: #6b7280;
}

.header-actions {
  justify-content: flex-end;
  flex-wrap: wrap;
}

.page-layout {
  display: grid;
  grid-template-columns: 270px minmax(0, 1fr);
  gap: 18px;
  align-items: start;
}

.page-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
  position: sticky;
  top: 0;
  max-height: calc(100vh - 126px);
  overflow: auto;
}

.page-item {
  display: grid;
  grid-template-columns: 1fr auto;
  gap: 6px 10px;
  text-align: left;
  padding: 12px;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  background: #fff;
  cursor: pointer;
  transition: border-color 0.18s ease, box-shadow 0.18s ease, background 0.18s ease;
}

.page-item:hover {
  border-color: #c7ccd3;
  background: #fbfbfc;
}

.page-item.active {
  border-color: #7d1231;
  box-shadow: 0 0 0 2px rgba(125, 18, 49, 0.08);
}

.page-item span {
  color: #6b7280;
  font-size: 12px;
}

.editor-panel {
  min-height: 520px;
  background: #fff;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 0 18px 18px;
  overflow: hidden;
}

.editor-header {
  position: sticky;
  top: 0;
  z-index: 3;
  margin: 0 -18px 18px;
  padding: 16px 18px;
  background: rgba(255, 255, 255, 0.96);
  border-bottom: 1px solid #e5e7eb;
  backdrop-filter: blur(8px);
}

.page-form {
  max-width: 980px;
}

.create-video-form {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.editable-list.compact {
  gap: 10px;
}

.editable-item {
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 12px;
  background: #fafafa;
  transition: border-color 0.18s ease, background 0.18s ease;
}

.editable-item.drag-over {
  border-color: #f59e0b;
  background: #fff7ed;
}

.editable-item-tools {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 8px;
  margin-bottom: 10px;
}

.editable-drag-handle {
  width: 30px;
  height: 28px;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  background: #ffffff;
  color: #6b7280;
  cursor: grab;
  font-weight: 700;
  line-height: 1;
}

.editable-drag-handle:active {
  cursor: grabbing;
}

.editable-add {
  align-self: flex-start;
}

.media-field {
  display: grid;
  grid-template-columns: 110px 128px minmax(0, 1fr);
  gap: 10px;
  align-items: center;
}

.media-field label,
.inline-form-row label {
  color: #606266;
  font-size: 14px;
}

.media-field-preview,
.media-field-empty {
  width: 128px;
  height: 82px;
  border: 1px solid #e5e7eb;
  border-radius: 6px;
  background: #f9fafb;
  object-fit: cover;
}

.media-field-empty {
  display: flex;
  align-items: center;
  justify-content: center;
  color: #6b7280;
  font-size: 12px;
  overflow: hidden;
  padding: 6px;
  word-break: break-all;
}

.media-field-actions {
  grid-column: 3;
  justify-content: flex-start;
}

.inline-form-row {
  display: grid;
  grid-template-columns: 110px 1fr;
  gap: 10px;
  align-items: start;
}

.media-picker-toolbar {
  justify-content: flex-start;
  margin-bottom: 14px;
}

.media-picker-toolbar .el-input {
  width: 360px;
}

.media-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
  gap: 12px;
  min-height: 240px;
}

.media-option {
  display: flex;
  flex-direction: column;
  gap: 8px;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 8px;
  background: white;
  cursor: pointer;
  text-align: left;
}

.media-option .el-image,
.media-kind {
  width: 100%;
  height: 86px;
  border-radius: 6px;
  background: #f3f4f6;
}

.media-kind {
  display: flex;
  align-items: center;
  justify-content: center;
  color: #6b7280;
}

.media-name {
  color: #374151;
  font-size: 12px;
  word-break: break-all;
}

@media (max-width: 980px) {
  .page-layout {
    grid-template-columns: 1fr;
  }

  .page-list {
    position: static;
    max-height: none;
  }

  .media-field {
    grid-template-columns: 1fr;
  }

  .media-field-actions {
    grid-column: auto;
  }
}
</style>
