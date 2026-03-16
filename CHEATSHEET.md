# ⚡ 快速参考手册

## 📦 常用导入

```typescript
// 类型定义
import type { NewsItem, TeamMember, Publication } from '@/types'

// 工具函数
import { getDay, getMonth, getYear, formatDate } from '@/utils'
import { scrollToElement, debounce, throttle } from '@/utils'

// 组件
import NavigationMenu from '@components/NavigationMenu.vue'
import Footer from '@components/Footer.vue'
```

---

## 🎨 CSS 变量速查

### 颜色
```css
var(--primary-color)      /* #7d1231 - 主色调 */
var(--primary-light)      /* #9a2c4d - 浅色 */
var(--primary-dark)       /* #5a0c22 - 深色 */
var(--secondary-color)    /* #13393e - 辅助色 */
var(--accent-color)       /* #3498db - 强调色 */
```

### 文字
```css
var(--text-primary)       /* #2c3e50 - 主文字 */
var(--text-secondary)     /* #5d6d7e - 次要文字 */
var(--text-light)         /* #7f8c8d - 浅色文字 */
```

### 背景
```css
var(--bg-light)           /* #f8f9fa - 浅灰背景 */
var(--bg-white)           /* #ffffff - 白色背景 */
var(--bg-card)            /* rgba(255,255,255,0.95) - 卡片背景 */
var(--bg-dark)            /* #2c2c2c - 深色背景 */
```

### 间距
```css
var(--spacing-xs)         /* 0.5rem - 超小间距 */
var(--spacing-sm)         /* 1rem - 小间距 */
var(--spacing-md)         /* 1.5rem - 中间距 */
var(--spacing-lg)         /* 2rem - 大间距 */
var(--spacing-xl)         /* 3rem - 超大间距 */
```

### 字体大小
```css
var(--font-size-xs)       /* 0.75rem */
var(--font-size-sm)       /* 0.875rem */
var(--font-size-base)     /* 1rem */
var(--font-size-lg)       /* 1.25rem */
var(--font-size-xl)       /* 1.5rem */
var(--font-size-2xl)      /* 2rem */
var(--font-size-3xl)      /* 2.8rem */
var(--font-size-4xl)      /* 3.5rem */
```

### 阴影
```css
var(--shadow-soft)        /* 柔和阴影 */
var(--shadow-medium)      /* 中等阴影 */
var(--shadow-large)       /* 大阴影 */
```

### 圆角
```css
var(--border-radius-sm)   /* 8px */
var(--border-radius)      /* 16px */
var(--border-radius-lg)   /* 20px */
```

---

## 🔧 工具函数速查

### 日期格式化
```typescript
getDay('2025-11-06')      // 6
getMonth('2025-11-06')    // 'Nov'
getYear('2025-11-06')     // 2025
formatDate('2025-11-06')  // 'November 6, 2025'
```

### DOM 操作
```typescript
scrollToElement('section-id')                    // 滚动到元素
scrollToElement('section-id', 'auto')           // 立即滚动
```

### 性能优化
```typescript
// 防抖 - 延迟执行
const debouncedFn = debounce((query: string) => {
  console.log(query)
}, 300)

// 节流 - 限制频率
const throttledFn = throttle(() => {
  console.log('scroll')
}, 1000)
```

### 设备检测
```typescript
isMobile()  // true | false
```

---

## 📝 Vue 组件模板

### 基础模板
```vue
<script setup lang="ts">
import { ref, computed } from 'vue'
import type { PropType } from 'vue'

// Props
interface Props {
  title?: string
  items?: string[]
}
const props = withDefaults(defineProps<Props>(), {
  title: '默认标题',
  items: () => []
})

// Emits
const emit = defineEmits<{
  update: [value: string]
  click: [event: MouseEvent]
}>()

// 响应式数据
const count = ref(0)

// 计算属性
const doubled = computed(() => count.value * 2)

// 方法
const handleClick = (e: MouseEvent) => {
  emit('click', e)
}
</script>

<template>
  <div class="component">
    <h2>{{ title }}</h2>
    <button @click="handleClick">{{ count }} - {{ doubled }}</button>
  </div>
</template>

<style scoped>
.component {
  padding: var(--spacing-lg);
  background: var(--bg-card);
  border-radius: var(--border-radius);
  box-shadow: var(--shadow-soft);
}
</style>
```

---

## 🎯 路由配置模板

```typescript
// src/router/index.ts
{
  path: '/your-page',
  name: 'your-page',
  component: () => import('../pages/YourPageView.vue')
}
```

---

## 🍱 菜单项配置模板

```typescript
// src/components/NavigationMenu.vue
{
  index: '/your-page',
  title: 'Your Page'
}

// 带子菜单
{
  index: 'parent',
  title: 'Parent Menu',
  children: [
    { index: '/child-1', title: 'Child 1' },
    { index: '/child-2', title: 'Child 2' }
  ]
}

// 外部链接
{
  index: 'https://github.com',
  title: 'GitHub',
  external: true
}
```

---

## 📊 TypeScript 类型模板

### 接口定义
```typescript
export interface YourType {
  id: number
  name: string
  description?: string  // 可选属性
  createdAt: string
  tags: string[]
}
```

### 使用示例
```typescript
const data = ref<YourType>({
  id: 1,
  name: 'Example',
  createdAt: '2025-11-06'
})

// 数组类型
const list = ref<YourType[]>([])
```

---

## 🐛 常见错误排查

### 1. 类型错误
```typescript
// ❌ 错误
const user = ref({ name: 'John' })

// ✅ 正确
interface User {
  name: string
  age: number
}
const user = ref<User>({ name: 'John', age: 25 })
```

### 2. 路径错误
```typescript
// ❌ 错误 - 相对路径过深
import utils from '../../../utils'

// ✅ 正确 - 使用别名
import utils from '@/utils'
```

### 3. 样式不生效
```vue
<!-- ❌ 忘记 scoped -->
<style>
.title { color: red; }
</style>

<!-- ✅ 正确 -->
<style scoped>
.title { color: var(--primary-color); }
</style>
```

---

## 💡 开发快捷键

### VS Code 推荐插件
- Volar (Vue 语言支持)
- ESLint (代码检查)
- Prettier (代码格式化)
- Path Intellisense (路径补全)

### 常用命令
```bash
npm run dev          # 启动开发服务器
npm run build        # 生产构建
npm run preview      # 预览构建结果
npm run lint         # 代码检查
```

---

## 📱 响应式断点

```css
/* 移动端 */
@media (max-width: 768px) {
  /* 移动端样式 */
}

/* 平板 */
@media (min-width: 769px) and (max-width: 1024px) {
  /* 平板样式 */
}

/* 桌面端 */
@media (min-width: 1025px) {
  /* 桌面端样式 */
}
```

---

## 🎨 动画过渡类

```css
/* 淡入淡出 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* 滑动 */
.slide-enter-active,
.slide-leave-active {
  transition: transform 0.3s ease;
}

.slide-enter-from {
  transform: translateX(-100%);
}

.slide-leave-to {
  transform: translateX(100%);
}
```

---

## ✅ 提交前检查清单

- [ ] TypeScript 无类型错误
- [ ] 删除了 console.log
- [ ] 删除了注释的无用代码
- [ ] 使用了 CSS 变量
- [ ] 添加了错误处理
- [ ] 测试了移动端
- [ ] 更新了路由（如有新页面）
- [ ] 代码已格式化

---

**打印此页作为桌面参考！** 📋
