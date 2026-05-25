<template>
  <div class="pages-admin">
    <div class="admin-page-header">
      <div>
        <h1>页面内容</h1>
        <p>维护 About、Contact、Research Direction、VideoMind 等长页面内容</p>
      </div>
      <el-button :loading="loading" @click="loadPages">刷新</el-button>
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
            </div>
          </div>

          <el-form label-width="110px" class="page-form">
            <el-form-item label="页面标题">
              <el-input v-model="editing.title" />
            </el-form-item>
            <el-form-item label="后台说明">
              <el-input v-model="editing.description" />
            </el-form-item>

            <template v-if="editing.slug === 'about'">
              <el-divider content-position="left">顶部与使命</el-divider>
              <el-form-item label="副标题">
                <el-input v-model="content.subtitle" />
              </el-form-item>
              <el-form-item label="使命标题">
                <el-input v-model="content.mission.title" />
              </el-form-item>
              <el-form-item label="使命正文">
                <el-input v-model="content.mission.text" type="textarea" :rows="4" />
              </el-form-item>

              <el-divider content-position="left">Research Thrusts</el-divider>
              <EditableList
                :items="content.researchThrusts"
                add-label="添加研究领域"
                @add="content.researchThrusts.push({ icon: 'Search', title: '', items: [] })"
                @remove="removeAt(content.researchThrusts, $event)"
              >
                <template #default="{ item }">
                  <el-input v-model="item.title" placeholder="标题" />
                  <el-select v-model="item.icon" placeholder="图标">
                    <el-option v-for="icon in iconOptions" :key="icon" :label="icon" :value="icon" />
                  </el-select>
                  <TextListEditor v-model="item.items" placeholder="每行一条要点" />
                </template>
              </EditableList>

              <el-divider content-position="left">Core Expertise</el-divider>
              <EditableList
                :items="content.expertise"
                add-label="添加核心能力"
                @add="content.expertise.push({ icon: 'Aim', title: '', description: '' })"
                @remove="removeAt(content.expertise, $event)"
              >
                <template #default="{ item }">
                  <el-input v-model="item.title" placeholder="标题" />
                  <el-input v-model="item.description" placeholder="说明" />
                  <el-select v-model="item.icon" placeholder="图标">
                    <el-option v-for="icon in iconOptions" :key="icon" :label="icon" :value="icon" />
                  </el-select>
                </template>
              </EditableList>

              <el-divider content-position="left">Strategic Vision</el-divider>
              <el-form-item label="愿景标题">
                <el-input v-model="content.vision.title" />
              </el-form-item>
              <el-form-item label="愿景正文">
                <el-input v-model="content.vision.text" type="textarea" :rows="4" />
              </el-form-item>
              <el-form-item label="领域标签">
                <TextListEditor v-model="content.vision.domains" placeholder="每行一个标签" />
              </el-form-item>
              <el-form-item label="引用语">
                <el-input v-model="content.vision.quote" type="textarea" :rows="3" />
              </el-form-item>
            </template>

            <template v-else-if="editing.slug === 'contact'">
              <el-form-item label="标题">
                <el-input v-model="content.title" />
              </el-form-item>
              <el-form-item label="可联系时间">
                <el-input v-model="content.availability" />
              </el-form-item>
              <el-form-item label="邮箱">
                <el-input v-model="content.email" />
              </el-form-item>
              <ImageField v-model="content.backgroundImage" label="背景图" />
            </template>

            <template v-else-if="editing.slug === 'research-direction'">
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

            <template v-else-if="editing.slug === 'videomind'">
              <el-form-item label="页面大标题">
                <el-input v-model="content.title" type="textarea" :rows="2" />
              </el-form-item>
              <ProjectSectionEditor v-model="content.introduction" image-label="示例图" list-label="要点" />
              <ProjectSectionEditor v-model="content.statistics" image-label="分类图" />
              <EditableList
                :items="content.statistics.charts"
                add-label="添加统计图"
                compact
                @add="content.statistics.charts.push({ image: '', alt: '' })"
                @remove="removeAt(content.statistics.charts, $event)"
              >
                <template #default="{ item }">
                  <ImageField v-model="item.image" label="统计图" />
                  <el-input v-model="item.alt" placeholder="图片说明" />
                </template>
              </EditableList>
              <ProjectSectionEditor v-model="content.results" image-label="" />
              <EditableList
                :items="content.results.figures"
                add-label="添加结果图"
                compact
                @add="content.results.figures.push({ title: '', image: '' })"
                @remove="removeAt(content.results.figures, $event)"
              >
                <template #default="{ item }">
                  <el-input v-model="item.title" placeholder="图标题" />
                  <ImageField v-model="item.image" label="结果图" />
                </template>
              </EditableList>
              <EditableList
                :items="content.links"
                add-label="添加下载链接"
                @add="content.links.push({ type: 'code', label: 'Link', url: '' })"
                @remove="removeAt(content.links, $event)"
              >
                <template #default="{ item }">
                  <el-select v-model="item.type" placeholder="类型">
                    <el-option label="Code" value="code" />
                    <el-option label="Paper" value="paper" />
                    <el-option label="Data" value="data" />
                    <el-option label="Link" value="link" />
                  </el-select>
                  <el-input v-model="item.label" placeholder="按钮文字" />
                  <el-input v-model="item.url" placeholder="URL" />
                </template>
              </EditableList>
              <el-form-item label="引用">
                <el-input v-model="content.citation" type="textarea" :rows="10" />
              </el-form-item>
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

            <template v-else-if="editing.slug === 'vknow'">
              <el-form-item label="页面标题">
                <el-input v-model="content.title" />
              </el-form-item>
              <ProjectSectionEditor v-model="content.introduction" image-label="" list-label="模块" />
              <el-divider content-position="left">Dataset Card</el-divider>
              <el-form-item label="卡片标题">
                <el-input v-model="content.dataset.title" />
              </el-form-item>
              <el-form-item label="卡片说明">
                <el-input v-model="content.dataset.description" type="textarea" :rows="3" />
              </el-form-item>
              <ImageField v-model="content.dataset.image" label="卡片图片" />
              <el-form-item label="图片说明">
                <el-input v-model="content.dataset.imageAlt" />
              </el-form-item>
              <el-form-item label="按钮文字">
                <el-input v-model="content.dataset.linkText" />
              </el-form-item>
              <el-form-item label="跳转路径">
                <el-input v-model="content.dataset.linkPath" />
              </el-form-item>
            </template>

            <template v-else-if="editing.slug.startsWith('video-')">
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
  </div>
</template>

<script setup lang="ts">
import { computed, defineComponent, h, nextTick, onMounted, ref, watch, type PropType } from 'vue'
import { ElButton, ElInput, ElMessage, ElUpload } from 'element-plus'
import type { UploadRequestOptions } from 'element-plus'
import { getSitePage, listAdmin, listSitePages, updateSitePage, uploadMedia } from '../../api/admin'
import type { MediaAsset, SitePage } from '../../api/client'

const pages = ref<SitePage[]>([])
const activeSlug = ref('')
const editing = ref<SitePage | null>(null)
const loading = ref(false)
const loadingPage = ref(false)
const saving = ref(false)
const rawJSON = ref('')
const mediaPickerVisible = ref(false)
const mediaLoading = ref(false)
const mediaUploading = ref(false)
const mediaQuery = ref('')
const mediaKind = ref('')
const mediaOptions = ref<MediaAsset[]>([])
const pendingMediaSetter = ref<((url: string) => void) | null>(null)

const iconOptions = ['Search', 'Shield', 'Setting', 'Connection', 'Aim', 'Lock', 'Refresh']

const content = computed<Record<string, any>>(() => editing.value?.content || {})

const fallbackContent: Record<string, () => Record<string, any>> = {
  about: () => ({ subtitle: '', mission: { title: '', text: '' }, researchThrusts: [], expertise: [], vision: { title: '', text: '', domains: [], quote: '' } }),
  contact: () => ({ title: '', availability: '', email: '', backgroundImage: '' }),
  'research-direction': () => ({ directions: [] }),
  videomind: () => ({ title: '', introduction: emptyProjectSection(), statistics: { ...emptyProjectSection(), charts: [] }, results: { ...emptyProjectSection(), figures: [] }, links: [], citation: '' }),
  'dr-baoyao-yang': () => ({ name: '', image: '', alt: '', paragraphs: [] }),
  vknow: () => ({ title: '', introduction: emptyProjectSection(), dataset: { title: '', description: '', image: '', imageAlt: '', linkText: '', linkPath: '' } }),
}

function emptyProjectSection() {
  return { title: '', text: '', items: [], image: '' }
}

const pageLabel = (slug: string, title: string) => {
  const labels: Record<string, string> = {
    about: 'About',
    contact: 'Contact',
    'research-direction': 'Research Direction',
    videomind: 'VideoMind',
    'dr-baoyao-yang': 'Baoyao Yang',
    vknow: 'vKnow',
    'video-xiaoqi-zheng-01': '视频：Xiaoqi Zheng',
    'video-xianrun-xu-01': '视频：Xianrun Xu',
    'video-yali-ma-01': '视频：Yali Ma',
  }
  return labels[slug] || title
}

const normalizePageContent = (page: SitePage) => {
  const fallback = fallbackContent[page.slug]?.() || (page.slug.startsWith('video-') ? { title: '', video: '' } : {})
  page.content = deepMerge(fallback, page.content || {})
  if (page.slug === 'videomind') {
    page.content.statistics.charts ||= []
    page.content.results.figures ||= []
    page.content.links ||= []
  }
  if (page.slug === 'vknow') {
    page.content.dataset ||= {}
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

const loadPages = async () => {
  loading.value = true
  try {
    const result = await listSitePages({ page: 1, pageSize: 50 })
    pages.value = result.items
    if (!activeSlug.value && pages.value.length) {
      await selectPage(pages.value[0].slug)
    }
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
  } finally {
    saving.value = false
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

const ProjectSectionEditor = defineComponent({
  props: {
    modelValue: { type: Object as () => Record<string, any>, required: true },
    imageLabel: { type: String, default: '图片' },
    listLabel: { type: String, default: '要点' },
  },
  setup(props) {
    return () => h('div', { class: 'project-section-editor' }, [
      h('div', { class: 'inline-form-row' }, [
        h('label', '小标题'),
        h(ElInput, { modelValue: props.modelValue.title, 'onUpdate:modelValue': (v: string) => { props.modelValue.title = v } }),
      ]),
      h('div', { class: 'inline-form-row' }, [
        h('label', '正文'),
        h(ElInput, { modelValue: props.modelValue.text, 'onUpdate:modelValue': (v: string) => { props.modelValue.text = v }, type: 'textarea', rows: 3 }),
      ]),
      props.listLabel ? h('div', { class: 'inline-form-row' }, [
        h('label', props.listLabel),
        h(TextListEditor, { modelValue: props.modelValue.items || [], 'onUpdate:modelValue': (v: string[]) => { props.modelValue.items = v } }),
      ]) : null,
      props.imageLabel ? h(ImageField, { modelValue: props.modelValue.image || '', label: props.imageLabel, 'onUpdate:modelValue': (v: string) => { props.modelValue.image = v } }) : null,
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
.project-section-editor {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.admin-page-header,
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
