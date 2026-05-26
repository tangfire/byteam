<template>
  <div class="editable-list" :class="{ compact }">
    <div
      v-for="(item, index) in items"
      :key="index"
      class="editable-item"
      :class="{ 'compact-item': compact, 'drag-over': overIndex === index }"
      @dragover="handleDragOver(index, $event)"
      @dragleave="handleDragLeave(index)"
      @drop="handleDrop(index, $event)"
    >
      <div class="editable-item-tools">
        <button
          class="editable-drag-handle"
          type="button"
          draggable="true"
          title="拖动排序"
          @dragstart="handleDragStart(index, $event)"
          @dragend="resetDrag"
        >
          ⋮⋮
        </button>
        <el-button size="small" text :disabled="index === 0" @click="move(index, index - 1)">上移</el-button>
        <el-button size="small" text :disabled="index === items.length - 1" @click="move(index, index + 1)">下移</el-button>
        <el-popconfirm title="确认删除这一项？保存后前台将不再显示。" @confirm="emit('remove', index)">
          <template #reference>
            <el-button size="small" type="danger" link>删除</el-button>
          </template>
        </el-popconfirm>
      </div>
      <div class="editable-item-fields">
        <slot :item="item" :index="index" />
      </div>
    </div>

    <el-button class="editable-add" @click="emit('add')">{{ addLabel }}</el-button>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const props = withDefaults(defineProps<{
  items: any[]
  addLabel?: string
  compact?: boolean
}>(), {
  addLabel: '添加',
  compact: false,
})

const emit = defineEmits<{
  add: []
  remove: [index: number]
}>()

const dragIndex = ref<number | null>(null)
const overIndex = ref<number | null>(null)

const move = (from: number, to: number) => {
  if (to < 0 || to >= props.items.length) return
  const [item] = props.items.splice(from, 1)
  props.items.splice(to, 0, item)
}

const resetDrag = () => {
  dragIndex.value = null
  overIndex.value = null
}

const handleDragStart = (index: number, event: DragEvent) => {
  dragIndex.value = index
  overIndex.value = null
  event.dataTransfer?.setData('text/plain', String(index))
}

const handleDragOver = (index: number, event: DragEvent) => {
  if (dragIndex.value === null || dragIndex.value === index) return
  event.preventDefault()
  overIndex.value = index
}

const handleDragLeave = (index: number) => {
  if (overIndex.value === index) overIndex.value = null
}

const handleDrop = (index: number, event: DragEvent) => {
  event.preventDefault()
  if (dragIndex.value === null || dragIndex.value === index) {
    resetDrag()
    return
  }
  move(dragIndex.value, index)
  resetDrag()
}
</script>

<style scoped>
.editable-list,
.editable-item-fields {
  display: flex;
  flex-direction: column;
  gap: 16px;
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
</style>
