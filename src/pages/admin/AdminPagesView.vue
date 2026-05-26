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
              <div class="content-editor-section">
                <div class="section-heading">
                  <div>
                    <h3>研究方向列表</h3>
                    <p>每一项对应前台 Research Direction 页面中的一个研究方向卡片。</p>
                  </div>
                  <el-tag type="info">{{ content.directions.length }} 项</el-tag>
                </div>
                <EditableList
                  :items="content.directions"
                  add-label="添加研究方向"
                  @add="content.directions.push({ title: '', image: '', alt: '', sections: [] })"
                  @remove="removeAt(content.directions, $event)"
                >
                  <template #default="{ item, index }">
                    <div class="direction-editor">
                      <div class="item-title-row">
                        <span class="item-index">方向 {{ index + 1 }}</span>
                        <el-input v-model="item.title" placeholder="方向标题" />
                      </div>

                      <div class="direction-body-grid">
                        <ImageField v-model="item.image" class="direction-image-field" label="配图" />
                        <div class="field-stack">
                          <label>图片说明</label>
                          <el-input v-model="item.alt" placeholder="用于图片 alt 文本，建议简短描述图片内容" />
                        </div>
                      </div>

                      <div class="nested-editor-block">
                        <div class="nested-heading">
                          <strong>文本段落</strong>
                          <span>{{ item.sections?.length || 0 }} 段</span>
                        </div>
                        <EditableList
                          :items="item.sections"
                          add-label="添加文本段"
                          compact
                          @add="item.sections.push({ title: '', text: '' })"
                          @remove="removeAt(item.sections, $event)"
                        >
                          <template #default="{ item: section, index: sectionIndex }">
                            <div class="section-editor-row">
                              <span class="subitem-index">{{ sectionIndex + 1 }}</span>
                              <el-input v-model="section.title" placeholder="小标题，如 Challenges" />
                              <el-input v-model="section.text" type="textarea" :rows="3" placeholder="正文" />
                            </div>
                          </template>
                        </EditableList>
                      </div>
                    </div>
                  </template>
                </EditableList>
              </div>
            </template>

            <template v-else-if="editing.slug === 'dr-baoyao-yang'">
              <div class="content-editor-section">
                <div class="section-heading">
                  <div>
                    <h3>个人主页资料</h3>
                    <p>维护前台 Baoyao Yang 页面里的姓名、照片和简介段落。</p>
                  </div>
                </div>
                <div class="profile-editor-grid">
                  <div class="profile-photo-panel">
                    <ImageField v-model="content.image" class="profile-photo-field" label="照片" preview-size="portrait" />
                  </div>
                  <div class="profile-fields">
                    <div class="field-stack">
                      <label>姓名标题</label>
                      <el-input v-model="content.name" />
                    </div>
                    <div class="field-stack">
                      <label>图片说明</label>
                      <el-input v-model="content.alt" placeholder="用于图片 alt 文本" />
                    </div>
                  </div>
                </div>
              </div>

              <div class="content-editor-section">
                <div class="section-heading">
                  <div>
                    <h3>简介段落</h3>
                    <p>每张卡片是一段正文，拖动左侧手柄可以调整顺序。</p>
                  </div>
                  <el-tag type="info">{{ content.paragraphs.length }} 段</el-tag>
                </div>
                <EditableList
                  :items="content.paragraphs"
                  add-label="添加简介段落"
                  @add="content.paragraphs.push('')"
                  @remove="removeAt(content.paragraphs, $event)"
                >
                  <template #default="{ index }">
                    <div class="paragraph-editor">
                      <span class="item-index">段落 {{ index + 1 }}</span>
                      <el-input v-model="content.paragraphs[index]" type="textarea" :rows="5" placeholder="请输入简介段落" />
                    </div>
                  </template>
                </EditableList>
              </div>
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
import { ElButton, ElInput, ElMessage, ElMessageBox, ElPopconfirm, ElUpload } from 'element-plus'
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
const pageSnapshot = ref('')
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

const toPlainPage = (page: SitePage) => ({
  slug: page.slug,
  title: page.title,
  description: page.description,
  content: page.content,
  status: page.status,
  sortOrder: page.sortOrder,
})

const serializePage = () => editing.value ? JSON.stringify(toPlainPage(editing.value)) : ''

const hasUnsavedPageChanges = computed(() => Boolean(editing.value) && serializePage() !== pageSnapshot.value)

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
  if (page.slug === 'research-direction') {
    page.content.directions = normalizeDirections(page.content.directions)
  } else if (page.slug === 'dr-baoyao-yang') {
    page.content.name = normalizeText(page.content.name)
    page.content.image = normalizeText(page.content.image)
    page.content.alt = normalizeText(page.content.alt)
    page.content.paragraphs = normalizeTextItems(page.content.paragraphs)
  } else if (page.slug.startsWith('video-')) {
    page.content.title = normalizeText(page.content.title)
    page.content.video = normalizeText(page.content.video)
  }
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

const toRecord = (value: unknown): Record<string, any> => isPlainObject(value) ? value : {}

const normalizeText = (value: unknown) => typeof value === 'string' ? value : ''

const normalizeTextItems = (value: unknown) => {
  if (Array.isArray(value)) {
    return value.map((item) => typeof item === 'string' ? item : String(item ?? ''))
  }
  if (typeof value === 'string') {
    return value.split('\n').map((item) => item.trim()).filter(Boolean)
  }
  return []
}

const normalizeSections = (value: unknown) => {
  if (!Array.isArray(value)) return []
  return value.map((section) => {
    if (typeof section === 'string') return { title: '', text: section }
    const record = toRecord(section)
    return {
      ...record,
      title: normalizeText(record.title),
      text: normalizeText(record.text),
    }
  })
}

const normalizeDirections = (value: unknown) => {
  if (!Array.isArray(value)) return []
  return value.map((direction) => {
    const record = toRecord(direction)
    return {
      ...record,
      title: normalizeText(record.title),
      image: normalizeText(record.image),
      alt: normalizeText(record.alt),
      sections: normalizeSections(record.sections),
    }
  })
}

const loadPages = async () => {
  loading.value = true
  try {
    const result = await listSitePages({ page: 1, pageSize: 50 })
    pages.value = result.items.filter((item) => !sourceManagedSlugs.has(item.slug))
    if (activeSlug.value && !pages.value.some((item) => item.slug === activeSlug.value)) {
      activeSlug.value = ''
      editing.value = null
      pageSnapshot.value = ''
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

const confirmDiscardPageChanges = async () => {
  if (!hasUnsavedPageChanges.value) return true
  try {
    await ElMessageBox.confirm('当前页面有未保存内容，确认切换并放弃这些修改？', '未保存修改', {
      type: 'warning',
      confirmButtonText: '放弃修改',
      cancelButtonText: '继续编辑',
    })
    return true
  } catch {
    return false
  }
}

const selectPage = async (slug: string, force = false) => {
  if (!force && slug !== activeSlug.value && !(await confirmDiscardPageChanges())) return
  activeSlug.value = slug
  loadingPage.value = true
  try {
    const page = await getSitePage(slug)
    normalizePageContent(page)
    editing.value = page
    rawJSON.value = JSON.stringify(page.content, null, 2)
    pageSnapshot.value = serializePage()
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
    pageSnapshot.value = serializePage()
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
    await selectPage(saved.slug, true)
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
    pageSnapshot.value = ''
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

const previewFileName = (url: string) => {
  const clean = String(url || '').split('?')[0].split('#')[0]
  return clean.split('/').filter(Boolean).pop() || clean || '未选择'
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
        class: ['editable-item', props.compact ? 'compact-item' : '', overIndex.value === index ? 'drag-over' : ''],
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
          h(ElPopconfirm, {
            title: '确认删除这一项？保存后前台将不再显示。',
            onConfirm: () => emit('remove', index),
          }, {
            reference: () => h(ElButton, { size: 'small', type: 'danger', link: true }, () => '删除'),
          }),
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
    previewSize: { type: String, default: 'default' },
  },
  emits: ['update:modelValue'],
  setup(props, { emit }) {
    const set = (url: string) => emit('update:modelValue', url)
    const isImage = () => props.modelValue && props.kind === 'image'
    return () => h('div', { class: ['media-field', props.previewSize === 'portrait' ? 'portrait-preview' : ''] }, [
      h('label', props.label),
      isImage()
        ? h('img', { src: props.modelValue, alt: '', class: 'media-field-preview' })
        : h('div', { class: 'media-field-empty' }, previewFileName(props.modelValue)),
      h(ElInput, { modelValue: props.modelValue, 'onUpdate:modelValue': set, placeholder: '可粘贴 URL，也可选择/上传媒体' }),
      h('div', { class: 'media-field-actions' }, [
        h(ElButton, { onClick: () => openMediaPicker(set, props.kind) }, () => '选择媒体'),
        h(ElUpload, { accept: props.kind === 'video' ? '.mp4,video/mp4' : 'image/*', showFileList: false, httpRequest: (options: UploadRequestOptions) => uploadAndSet(options, set) }, () => h(ElButton, { loading: mediaUploading.value }, () => '上传并使用')),
        props.modelValue
          ? h(ElPopconfirm, {
            title: '确认清空这个媒体地址？保存后前台将不再显示这个资源。',
            onConfirm: () => set(''),
          }, {
            reference: () => h(ElButton, null, () => '清空'),
          })
          : null,
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
.editable-item-fields,
.content-editor-section,
.direction-editor,
.field-stack,
.nested-editor-block,
.paragraph-editor {
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
  max-width: 1040px;
}

.create-video-form {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.content-editor-section {
  gap: 14px;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  background: #ffffff;
  padding: 16px;
}

.section-heading,
.nested-heading,
.item-title-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.section-heading {
  padding-bottom: 12px;
  border-bottom: 1px solid #eef0f3;
}

.section-heading h3 {
  margin: 0;
  color: #25313b;
  font-size: 16px;
  line-height: 1.35;
}

.section-heading p {
  margin: 4px 0 0;
  color: #6b7280;
  font-size: 13px;
  line-height: 1.5;
}

.editable-list.compact {
  gap: 10px;
}

.editable-item {
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 14px;
  background: #fbfbfc;
  transition: border-color 0.18s ease, background 0.18s ease;
}

.editable-item.compact-item {
  padding: 10px;
  background: #ffffff;
}

.editable-item.drag-over {
  border-color: #f59e0b;
  background: #fff7ed;
}

.editable-item-tools {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  gap: 8px;
  margin-bottom: 10px;
  padding-bottom: 10px;
  border-bottom: 1px solid #eceff3;
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

.direction-editor {
  gap: 14px;
}

.item-title-row {
  align-items: flex-start;
}

.item-index,
.subitem-index {
  display: inline-flex;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  min-width: 74px;
  height: 32px;
  border-radius: 6px;
  background: #f3f4f6;
  color: #374151;
  font-size: 12px;
  font-weight: 600;
}

.subitem-index {
  min-width: 32px;
  height: 30px;
}

.direction-body-grid,
.profile-editor-grid {
  display: grid;
  grid-template-columns: minmax(0, 1fr) minmax(260px, 340px);
  gap: 14px;
  align-items: start;
}

.direction-image-field {
  min-width: 0;
}

.profile-fields {
  min-width: 0;
}

.field-stack {
  gap: 8px;
  min-width: 0;
}

.field-stack label,
.nested-heading strong {
  color: #374151;
  font-size: 13px;
  font-weight: 600;
}

.nested-editor-block {
  gap: 10px;
  border-radius: 8px;
  background: #f7f8fa;
  padding: 12px;
}

.nested-heading {
  color: #6b7280;
  font-size: 12px;
}

.section-editor-row {
  display: grid;
  grid-template-columns: 32px minmax(180px, 0.45fr) minmax(260px, 1fr);
  gap: 10px;
  align-items: start;
}

.profile-editor-grid {
  grid-template-columns: minmax(260px, 380px) minmax(0, 1fr);
}

.profile-photo-panel {
  min-width: 0;
}

.paragraph-editor {
  gap: 10px;
}

.media-field,
:deep(.media-field) {
  display: grid;
  grid-template-columns: 128px minmax(0, 1fr);
  gap: 10px 12px;
  align-items: start;
  min-width: 0;
  max-width: 100%;
  overflow: hidden;
}

.media-field label,
:deep(.media-field label) {
  grid-column: 1 / -1;
  color: #606266;
  font-size: 14px;
  font-weight: 600;
}

.media-field-preview,
:deep(.media-field-preview),
.media-field-empty,
:deep(.media-field-empty) {
  width: 128px;
  height: 82px;
  max-width: 100%;
  border: 1px solid #e5e7eb;
  border-radius: 6px;
  background: #f9fafb;
  object-fit: cover;
  display: block;
}

.media-field-empty,
:deep(.media-field-empty) {
  display: flex;
  align-items: center;
  justify-content: center;
  color: #6b7280;
  font-size: 12px;
  overflow: hidden;
  padding: 6px;
  word-break: break-all;
}

.media-field-actions,
:deep(.media-field-actions) {
  grid-column: 2;
  justify-content: flex-start;
  flex-wrap: wrap;
}

.media-field :deep(.el-input),
:deep(.media-field .el-input) {
  min-width: 0;
}

.media-field.portrait-preview,
:deep(.media-field.portrait-preview) {
  grid-template-columns: 160px minmax(0, 1fr);
}

.media-field.portrait-preview .media-field-preview,
.media-field.portrait-preview .media-field-empty,
:deep(.media-field.portrait-preview .media-field-preview),
:deep(.media-field.portrait-preview .media-field-empty) {
  width: 160px;
  height: 190px;
  object-fit: cover;
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

  .direction-body-grid,
  .profile-editor-grid,
  .section-editor-row {
    grid-template-columns: 1fr;
  }

  .item-title-row,
  .section-heading,
  .nested-heading {
    align-items: flex-start;
    flex-direction: column;
  }

  .item-index,
  .subitem-index {
    align-self: flex-start;
  }
}
</style>
