<script setup lang="ts">
import { computed } from 'vue'
import type { MenuItem } from '../types'

// 定义菜单数据结构
const menuItems: MenuItem[] = [
  { index: '/', title: 'Home' },
  { index: '/about', title: 'About' },
  { index: '/news', title: 'News' },
  { 
    index: 'research', 
    title: 'Research',
    children: [
      { index: '/research-direction', title: 'Research Direction' },
      { index: '/research-projects', title: 'Research Projects' },
      { index: 'https://github.com/BaoyaoGroup', title: 'Github-Repositories', external: true }
    ]
  },
  {
    index: 'our-team',
    title: 'Our Team',
    children: [
      { index: '/dr-Baoyao-Yang', title: 'Baoyao Yang' },
      { index: '/our-group', title: 'Our Group' },
      { index: '/Alumni', title: 'Alumni' }
    ]
  },
  {
    index: 'publications',
    title: 'Publications',
    children: [
      { index: '/international-journals-conferences', title: 'International Journals/Conferences' },
      { index: '/patents', title: 'Patents' }
    ]
  },
  {
    index: 'Project',
    title: 'Project',
    children: [
      { index: '/vKnow', title: 'vKnow' },
      { index: '/VideoMind', title: 'Dataset' }
    ]
  },
  { index: '/contact', title: 'Contact' }
]

// 定义 emits
const emit = defineEmits<{
  select: []
}>()

// 处理选择事件
const handleSelect = () => {
  emit('select')
}

// 检查是否为外部链接
const isExternalLink = (item: MenuItem) => item.external === true
</script>

<template>
  <el-menu
    :default-active="$route.path"
    class="navigation-menu"
    mode="horizontal"
    :ellipsis="false"
    active-text-color="#7d1231"
    @select="handleSelect"
    :router="true"
    :popper-offset="16"
    text-color="#2f3542"
  >
    <!-- Logo -->
    <div class="logo">
      <img
        class="logo-img"
        src="/logo/001.png"
        alt="BYML Logo"
      />
    </div>

    <!-- 菜单项 -->
    <template v-for="item in menuItems" :key="item.index">
      <!-- 无子菜单 -->
      <el-menu-item 
        v-if="!item.children" 
        :index="item.index"
      >
        <p class="menu-title">{{ item.title }}</p>
      </el-menu-item>

      <!-- 有子菜单 -->
      <el-sub-menu v-else :index="item.index">
        <template #title>
          <p class="menu-title">{{ item.title }}</p>
        </template>
        
        <el-menu-item
          v-for="child in item.children"
          :key="child.index"
          :index="child.index"
          style="height: 50px"
        >
          <!-- 外部链接 -->
          <a 
            v-if="isExternalLink(child)"
            :href="child.index"
            target="_blank"
            class="external-link"
          >
            {{ child.title }}
          </a>
          <!-- 内部路由 -->
          <span v-else>{{ child.title }}</span>
        </el-menu-item>
      </el-sub-menu>
    </template>
  </el-menu>
</template>

<style scoped>
.logo {
  margin-right: auto;
  order: -1;
  margin-left: 30px;
}

.logo-img {
  width: 180px;
  height: auto;
  margin: 0 20px;
}

.menu-title {
  font-size: var(--font-size-lg);
  margin: 0;
}

.external-link {
  color: inherit;
  text-decoration: none;
}

.external-link:hover {
  color: white;
}
</style>
