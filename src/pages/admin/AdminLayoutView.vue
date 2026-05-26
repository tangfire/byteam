<template>
  <el-container class="admin-shell">
    <el-aside width="248px" class="admin-aside">
      <div class="admin-brand">
        <span class="brand-mark">B</span>
        <div>
          <strong>BYML 后台</strong>
          <small>官网内容管理系统</small>
        </div>
      </div>
      <el-menu
        router
        :default-active="$route.path"
        background-color="#111827"
        text-color="#d1d5db"
        active-text-color="#ffffff"
        class="admin-menu"
      >
        <template v-for="section in navSections" :key="section.title">
          <div class="menu-section-title">{{ section.title }}</div>
          <el-menu-item v-for="item in section.items" :key="item.path" :index="item.path">
            <span class="menu-dot">{{ item.short }}</span>
            <span>{{ item.label }}</span>
          </el-menu-item>
        </template>
      </el-menu>
    </el-aside>
    <el-container class="admin-content">
      <el-header class="admin-header">
        <div class="header-title">
          <strong>{{ activePage?.label || '后台管理' }}</strong>
          <span>{{ activePage?.description || '维护官网内容、资源和恢复快照' }}</span>
        </div>
        <div class="admin-header-actions">
          <router-link class="header-link" to="/">查看官网</router-link>
          <router-link class="header-link subtle" to="/admin/guide">运维说明</router-link>
          <el-button plain @click="handleLogout">退出登录</el-button>
        </div>
      </el-header>
      <el-main class="admin-main">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { logout } from '../../api/admin'
import { useAdminElementPlus } from '../../composables/useAdminElementPlus'

useAdminElementPlus()

const router = useRouter()
const route = useRoute()

const navSections = [
  {
    title: '概览',
    items: [
      { path: '/admin/dashboard', label: '后台首页', short: '首', description: '查看内容数量和常用维护入口' },
    ],
  },
  {
    title: '内容管理',
    items: [
      { path: '/admin/news', label: '新闻动态', short: '新', description: '维护首页和新闻页展示的动态' },
      { path: '/admin/people', label: '成员管理', short: '人', description: '维护研究生、校友和成员资料' },
      { path: '/admin/undergraduates', label: '本科生培养', short: '本', description: '维护本科生培养成果' },
      { path: '/admin/publications', label: '论文管理', short: '论', description: '维护论文、附件和首页精选' },
      { path: '/admin/patents', label: '专利与标准', short: '专', description: '维护专利、标准和相关编号' },
      { path: '/admin/research-projects', label: '科研项目', short: '项', description: '维护科研项目列表' },
    ],
  },
  {
    title: '页面与资源',
    items: [
      { path: '/admin/pages', label: '页面内容', short: '页', description: '维护仍由后台管理的长页面和论文视频页' },
      { path: '/admin/media', label: '媒体文件', short: '媒', description: '上传、命名和管理官网资源文件' },
    ],
  },
  {
    title: '维护',
    items: [
      { path: '/admin/trash', label: '回收站', short: '回', description: '恢复误删内容和媒体索引' },
      { path: '/admin/guide', label: '运维说明', short: '维', description: '查看备份、同步和服务器恢复流程' },
    ],
  },
]

const navItems = computed(() => navSections.flatMap((section) => section.items))
const activePage = computed(() => navItems.value.find((item) => route.path === item.path))

const handleLogout = async () => {
  logout()
  await router.push('/admin/login')
}
</script>

<style scoped>
.admin-shell {
  min-height: 100vh;
  height: 100vh;
  background: #f6f7f9;
  display: flex;
  align-items: stretch;
}

.admin-aside {
  background: #111827;
  flex-shrink: 0;
  display: flex;
  flex-direction: column;
  box-shadow: 8px 0 24px rgba(17, 24, 39, 0.08);
}

.admin-content {
  min-width: 0;
  height: 100vh;
}

.admin-brand {
  min-height: 72px;
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 0 18px;
  color: white;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.brand-mark {
  width: 38px;
  height: 38px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  background: #7d1231;
  color: #ffffff;
  font-weight: 800;
}

.admin-brand strong,
.header-title strong {
  display: block;
  font-size: 17px;
  font-weight: 700;
}

.admin-brand small {
  display: block;
  margin-top: 3px;
  color: #9ca3af;
  font-size: 12px;
}

.admin-menu {
  flex: 1;
  min-height: 0;
  overflow: auto;
  border-right: 0;
  padding: 10px 10px 18px;
}

.admin-menu :deep(.el-menu-item) {
  height: 42px;
  margin: 2px 0;
  border-radius: 8px;
  padding: 0 12px !important;
}

.admin-menu :deep(.el-menu-item.is-active) {
  background: #7d1231;
}

.menu-section-title {
  margin: 16px 10px 6px;
  color: #9ca3af;
  font-size: 12px;
  font-weight: 700;
}

.menu-dot {
  width: 24px;
  height: 24px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  margin-right: 10px;
  border-radius: 6px;
  background: rgba(255, 255, 255, 0.08);
  color: #e5e7eb;
  font-size: 12px;
  font-weight: 700;
}

.admin-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  height: 64px;
  background: #ffffff;
  border-bottom: 1px solid #e5e7eb;
  padding: 0 24px;
}

.header-title {
  min-width: 0;
}

.header-title span {
  display: block;
  margin-top: 3px;
  color: #6b7280;
  font-size: 12px;
}

.admin-header-actions {
  display: flex;
  align-items: center;
  gap: 12px;
  flex-shrink: 0;
}

.admin-main {
  padding: 24px;
  overflow: auto;
}

.header-link {
  color: #7d1231;
  text-decoration: none;
  font-size: 14px;
  font-weight: 600;
}

.header-link.subtle {
  color: #4b5563;
}

@media (max-width: 920px) {
  .admin-aside {
    width: 212px !important;
  }

  .admin-header {
    align-items: flex-start;
    height: auto;
    padding: 14px 16px;
    flex-direction: column;
  }

  .admin-main {
    padding: 16px;
  }
}
</style>
