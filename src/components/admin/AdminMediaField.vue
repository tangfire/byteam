<template>
  <div class="media-field" :class="{ 'portrait-preview': previewSize === 'portrait' }">
    <label>{{ label }}</label>
    <img v-if="isImage" :src="modelValue" alt="" class="media-field-preview">
    <div v-else class="media-field-empty">{{ previewFileName(modelValue) }}</div>
    <el-input :model-value="modelValue" placeholder="可粘贴 URL，也可选择/上传媒体" @update:model-value="setValue" />
    <div class="media-field-actions">
      <el-button @click="emit('pick', setValue, kind)">选择媒体</el-button>
      <el-upload :accept="uploadAccept" :show-file-list="false" :http-request="uploadAndSet">
        <el-button :loading="uploading">上传并使用</el-button>
      </el-upload>
      <el-popconfirm
        v-if="modelValue"
        title="确认清空这个媒体地址？保存后前台将不再显示这个资源。"
        @confirm="setValue('')"
      >
        <template #reference>
          <el-button>清空</el-button>
        </template>
      </el-popconfirm>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { UploadRequestOptions } from 'element-plus'

const props = withDefaults(defineProps<{
  modelValue?: string
  label?: string
  kind?: string
  previewSize?: string
  uploading?: boolean
}>(), {
  modelValue: '',
  label: '图片',
  kind: 'image',
  previewSize: 'default',
  uploading: false,
})

const emit = defineEmits<{
  'update:modelValue': [value: string]
  pick: [setter: (url: string) => void, kind: string]
  upload: [options: UploadRequestOptions, setter: (url: string) => void]
}>()

const isImage = computed(() => Boolean(props.modelValue) && props.kind === 'image')

const uploadAccept = computed(() => props.kind === 'video' ? '.mp4,video/mp4' : 'image/*')

const setValue = (value: string | number) => {
  emit('update:modelValue', String(value))
}

const uploadAndSet = (options: UploadRequestOptions) => {
  emit('upload', options, setValue)
}

const previewFileName = (url: string) => {
  const clean = String(url || '').split('?')[0].split('#')[0]
  return clean.split('/').filter(Boolean).pop() || clean || '未选择'
}
</script>

<style scoped>
.media-field {
  display: grid;
  grid-template-columns: 128px minmax(0, 1fr);
  gap: 10px 12px;
  align-items: start;
  min-width: 0;
  max-width: 100%;
  overflow: hidden;
}

.media-field label {
  grid-column: 1 / -1;
  color: #606266;
  font-size: 14px;
  font-weight: 600;
}

.media-field-preview,
.media-field-empty {
  width: 128px;
  height: 82px;
  max-width: 100%;
  border: 1px solid #e5e7eb;
  border-radius: 6px;
  background: #f9fafb;
  object-fit: cover;
  display: block;
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
  grid-column: 2;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  gap: 12px;
  flex-wrap: wrap;
}

.media-field :deep(.el-input) {
  min-width: 0;
}

.media-field.portrait-preview {
  grid-template-columns: 160px minmax(0, 1fr);
}

.media-field.portrait-preview .media-field-preview,
.media-field.portrait-preview .media-field-empty {
  width: 160px;
  height: 190px;
  object-fit: cover;
}

@media (max-width: 980px) {
  .media-field {
    grid-template-columns: 1fr;
  }

  .media-field-actions {
    grid-column: auto;
  }
}
</style>
