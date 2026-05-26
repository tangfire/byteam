<template>
  <el-dialog :model-value="modelValue" title="选择媒体" width="860px" @update:model-value="emit('update:modelValue', $event)">
    <div class="admin-media-picker-toolbar">
      <el-input v-model="query" :placeholder="searchPlaceholder" clearable @keyup.enter="loadMediaOptions" />
      <el-button @click="loadMediaOptions">搜索</el-button>
    </div>

    <div v-loading="loading" class="admin-media-picker-grid">
      <button
        v-for="asset in mediaOptions"
        :key="asset.id"
        class="admin-media-picker-option"
        type="button"
        @click="chooseMedia(asset.url)"
      >
        <el-image v-if="asset.kind === 'image'" :src="asset.url" fit="cover" />
        <span v-else class="admin-media-picker-kind">{{ formatMediaKind(asset.kind) }}</span>
        <span class="admin-media-picker-name">{{ asset.displayName || asset.originalName }}</span>
        <span v-if="showOriginalName" class="admin-media-picker-url">{{ asset.originalName }}</span>
      </button>
      <el-empty v-if="!loading && mediaOptions.length === 0" description="没有找到媒体文件" :image-size="72" />
    </div>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { listAdmin } from '../../api/admin'
import type { MediaAsset } from '../../api/client'
import { formatMediaKind } from '../../utils/adminFormat'

const props = withDefaults(defineProps<{
  modelValue: boolean
  kind?: string
  searchPlaceholder?: string
  queryExtra?: Record<string, string>
  showOriginalName?: boolean
}>(), {
  kind: '',
  searchPlaceholder: '按文件名或 URL 搜索',
  queryExtra: () => ({}),
  showOriginalName: false,
})

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
  choose: [url: string]
}>()

const loading = ref(false)
const query = ref('')
const mediaOptions = ref<MediaAsset[]>([])

const loadMediaOptions = async () => {
  loading.value = true
  try {
    const result = await listAdmin<MediaAsset>('media', {
      page: 1,
      pageSize: 48,
      q: query.value,
      kind: props.kind,
      ...props.queryExtra,
    })
    mediaOptions.value = result.items
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '媒体加载失败')
  } finally {
    loading.value = false
  }
}

const chooseMedia = (url: string) => {
  emit('choose', url)
  emit('update:modelValue', false)
}

watch(() => props.modelValue, (visible) => {
  if (visible) void loadMediaOptions()
})
</script>

<style scoped>
.admin-media-picker-toolbar {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  gap: 8px;
  flex-wrap: wrap;
  margin-bottom: 14px;
}

.admin-media-picker-toolbar .el-input {
  width: 360px;
}

.admin-media-picker-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
  gap: 12px;
  min-height: 180px;
  max-height: 58vh;
  overflow: auto;
}

.admin-media-picker-option {
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

.admin-media-picker-option:hover {
  border-color: #7d1231;
}

.admin-media-picker-option .el-image,
.admin-media-picker-kind {
  width: 100%;
  aspect-ratio: 4 / 3;
  border-radius: 4px;
  background: #f3f4f6;
}

.admin-media-picker-kind {
  display: flex;
  align-items: center;
  justify-content: center;
  color: #6b7280;
}

.admin-media-picker-name {
  overflow: hidden;
  color: #374151;
  font-size: 12px;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.admin-media-picker-url {
  overflow: hidden;
  color: #9ca3af;
  font-size: 11px;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
