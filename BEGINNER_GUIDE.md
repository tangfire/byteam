# 🚀 新手入门指南

欢迎加入 Beyond Machine Learning Group！这份指南将帮助你快速上手项目开发。

## 📁 项目结构说明

### 核心目录
```
src/
├── main.ts              # 应用入口文件
├── App.vue             # 根组件
├── router/             # 路由配置
│   └── index.ts       # 路由定义（使用懒加载）
├── types/             # TypeScript 类型定义
│   └── index.ts      # 所有类型接口
├── utils/             # 工具函数
│   └── index.ts      # 公共函数库
├── components/        # 可复用组件
│   ├── NavigationMenu.vue  # 导航菜单
│   ├── Footer.vue          # 页脚
│   └── ResearchCarousel.vue # 研究轮播
├── pages/             # 页面组件
│   ├── HomeView.vue         # 首页
│   ├── AboutView.vue        # 关于页
│   └── ...                 # 其他页面
└── assets/            # 静态资源
    └── styles/        # 全局样式
        └── variables.css  # CSS 变量定义
```

## 🎯 核心概念

### 1. Vue 3 Composition API
本项目使用 Vue 3 的 `<script setup>` 语法，这是最新的 Vue 开发方式。

**示例：**
```vue
<script setup lang="ts">
import { ref, computed } from 'vue'

// 响应式数据
const count = ref(0)

// 计算属性
const doubleCount = computed(() => count.value * 2)

// 方法
function increment() {
  count.value++
}
</script>

<template>
  <div>{{ count }} - {{ doubleCount }}</div>
  <button @click="increment">+1</button>
</template>
```

### 2. TypeScript 类型系统
所有代码必须使用 TypeScript，并遵循严格的类型检查。

**正确使用类型：**
```typescript
// ✅ 好的做法
interface User {
  name: string
  age: number
}

const user = ref<User>({
  name: 'John',
  age: 25
})

// ❌ 避免使用 any
const data: any = {}  // 不要这样做！
```

### 3. 路径别名
使用 `@` 符号简化导入路径：

```typescript
// ✅ 使用别名
import { formatDate } from '@/utils'
import type { NewsItem } from '@/types'
import NavigationMenu from '@components/NavigationMenu.vue'

// ❌ 避免相对路径
import { formatDate } from '../../../utils/index'
```

## 🛠️ 开发流程

### 环境搭建

1. **安装 Node.js** (v18+)
   ```bash
   # 检查版本
   node -v
   
   # 如果版本过低，请升级到 v18 或更高
   ```

2. **安装依赖**
   ```bash
   npm install
   ```

3. **启动开发服务器**
   ```bash
   npm run dev
   ```
   
   会自动打开浏览器访问 http://localhost:5173

### 创建新页面

1. **在 `src/pages/` 下创建新文件**
   ```bash
   # 例如创建团队成员页面
   src/pages/TeamMemberView.vue
   ```

2. **编写组件代码**
   ```vue
   <script setup lang="ts">
   import { ref } from 'vue'
   import type { TeamMember } from '@/types'
   
   const members = ref<TeamMember[]>([
     {
       name: '张三',
       avatar: '/avatar/zhangsan.jpg',
       role: '研究员'
     }
   ])
   </script>
   
   <template>
     <div class="team-container">
       <h1>团队成员</h1>
       <div v-for="member in members" :key="member.name">
         <img :src="member.avatar" :alt="member.name" />
         <p>{{ member.name }} - {{ member.role }}</p>
       </div>
     </div>
   </template>
   
   <style scoped>
   .team-container {
     padding: var(--spacing-lg);
   }
   </style>
   ```

3. **添加路由**
   
   编辑 `src/router/index.ts`：
   ```typescript
   {
     path: '/team-member',
     name: 'team-member',
     component: () => import('../pages/TeamMemberView.vue')
   }
   ```

4. **添加到导航菜单**
   
   编辑 `src/components/NavigationMenu.vue`，在 `menuItems` 数组中添加：
   ```typescript
   { index: '/team-member', title: 'Team Member' }
   ```

### 使用全局样式

所有颜色、间距等都应该使用 CSS 变量：

```vue
<style scoped>
.card {
  background: var(--bg-card);
  padding: var(--spacing-lg);
  border-radius: var(--border-radius);
  box-shadow: var(--shadow-soft);
  color: var(--text-primary);
}

.button {
  color: var(--primary-color);
  background: var(--gradient-primary);
}
</style>
```

### 使用工具函数

常用工具函数已封装在 `@/utils` 中：

```typescript
import { 
  getDay, 
  getMonth, 
  getYear, 
  formatDate,
  scrollToElement,
  debounce 
} from '@/utils'

// 格式化日期
const day = getDay('2025-11-06')  // 6
const month = getMonth('2025-11-06')  // 'Nov'
const year = getYear('2025-11-06')  // 2025

// 平滑滚动
scrollToElement('section-id')

// 防抖函数
const handleSearch = debounce((query: string) => {
  console.log('搜索:', query)
}, 300)
```

## 📝 代码规范

### 命名规范

1. **文件名**: PascalCase (大驼峰)
   - ✅ `HomeView.vue`
   - ✅ `ResearchDirection.vue`
   - ❌ `homeView.vue`
   - ❌ `home-view.vue`

2. **变量名**: camelCase (小驼峰)
   ```typescript
   const userName = 'John'
   const researchData = ref([])
   ```

3. **组件名**: PascalCase
   ```vue
   <!-- ✅ -->
   <NavigationMenu />
   <ResearchCard />
   
   <!-- ❌ -->
   <navigation-menu />
   <researchCard />
   ```

4. **常量**: UPPER_SNAKE_CASE
   ```typescript
   const MAX_RETRY_COUNT = 3
   const API_BASE_URL = 'https://api.example.com'
   ```

### 代码组织

**Vue 文件标准顺序：**

```vue
<!-- 1. Script 标签 -->
<script setup lang="ts">
// 导入（按类型分组）
import { ref, computed } from 'vue'
import type { NewsItem } from '@/types'
import { formatDate } from '@/utils'
import NavigationMenu from '@components/NavigationMenu.vue'

// Props 和 Emits
interface Props {
  title?: string
}
const props = withDefaults(defineProps<Props>(), {
  title: '默认标题'
})

const emit = defineEmits<{
  update: [value: string]
}>()

// 响应式数据
const dataList = ref<NewsItem[]>([])

// 计算属性
const filteredList = computed(() => {
  return dataList.value.filter(item => item.type === 'publication')
})

// 方法
const handleClick = () => {
  console.log('点击了')
}
</script>

<!-- 2. Template 标签 -->
<template>
  <div class="container">
    <!-- 模板内容 -->
  </div>
</template>

<!-- 3. Style 标签 -->
<style scoped>
.container {
  /* 样式 */
}
</style>
```

## 🔍 调试技巧

### 1. Vue DevTools
安装 [Vue DevTools](https://devtools.vuejs.org/) 浏览器扩展，可以：
- 查看组件树
- 检查响应式数据
- 追踪事件

### 2. 控制台日志
```typescript
// 开发环境下才输出日志
if (import.meta.env.DEV) {
  console.log('调试信息:', data)
}
```

### 3. 错误处理
```typescript
try {
  await fetchData()
} catch (error) {
  console.error('获取数据失败:', error)
  // 显示错误提示
  ElMessage.error('加载失败')
}
```

## 🎨 UI 设计原则

### 1. 使用 Design Tokens
所有样式值都应该来自 CSS 变量：

```css
/* ✅ 好的做法 */
color: var(--primary-color);
padding: var(--spacing-md);
font-size: var(--font-size-lg);

/* ❌ 避免硬编码 */
color: #7d1231;
padding: 20px;
font-size: 18px;
```

### 2. 响应式设计
使用媒体查询适配移动端：

```css
.container {
  padding: var(--spacing-lg);
}

@media (max-width: 768px) {
  .container {
    padding: var(--spacing-md);
  }
  
  .title {
    font-size: var(--font-size-xl);
  }
}
```

### 3. 动画过渡
添加平滑的过渡效果：

```css
.card {
  transition: all var(--transition-base);
}

.card:hover {
  transform: translateY(-5px);
  box-shadow: var(--shadow-medium);
}
```

## 📦 构建部署

### 开发构建
```bash
npm run dev
```

### 生产构建
```bash
npm run build
```

构建产物在 `dist/` 目录

### 预览生产构建
```bash
npm run preview
```

## 🐛 常见问题

### Q1: 遇到 "Cannot find module" 错误
**解决方案：**
1. 检查路径别名是否正确
2. 确保安装了所有依赖：`npm install`
3. 重启 VSCode TypeScript 服务

### Q2: 样式不生效
**可能原因：**
1. 忘记加 `scoped` 属性
2. CSS 变量未定义
3. 选择器优先级不够

**解决方案：**
```vue
<!-- ✅ 正确 -->
<style scoped>
.title {
  color: var(--primary-color);
}
</style>
```

### Q3: 路由跳转后页面不更新
**解决方案：**
确保使用了正确的路由参数监听：
```typescript
watch(
  () => route.params.id,
  (newId) => {
    fetchDetail(newId)
  }
)
```

## 📚 学习资源

### 官方文档
- [Vue 3 文档](https://vuejs.org/)
- [TypeScript 文档](https://www.typescriptlang.org/)
- [Element Plus 文档](https://element-plus.org/)
- [Vite 文档](https://vitejs.dev/)

### 视频教程
- Vue 3 官方教程
- TypeScript 从入门到精通

### 实践建议
1. **多动手写代码** - 看懂不如写会
2. **阅读优秀源码** - 学习最佳实践
3. **参与开源项目** - 积累实战经验
4. **写技术博客** - 巩固知识点

## 💡 快速检查清单

在提交代码前，请确认：

- [ ] TypeScript 没有类型错误
- [ ] 使用了 ESLint 检查代码
- [ ] 删除了所有 `console.log` 调试代码
- [ ] 删除了注释掉的无用代码
- [ ] 使用了 CSS 变量而非硬编码值
- [ ] 添加了必要的错误处理
- [ ] 测试了移动端响应式布局
- [ ] 更新了路由配置（如有新页面）

---

## 🤝 获取帮助

遇到问题时：
1. 先查阅本文档
2. 搜索相关官方文档
3. 在团队群组提问
4. 向导师或有经验的同事请教

**记住：没有愚蠢的问题，只有没问出来的问题！**

祝你开发愉快！🎉
