<template>
  <el-container class="admin-shell">
    <el-aside width="230px" class="admin-aside">
      <div class="admin-brand">BYML Admin</div>
      <el-menu router :default-active="$route.path" background-color="#111827" text-color="#d1d5db" active-text-color="#ffffff">
        <el-menu-item index="/admin/dashboard">Dashboard</el-menu-item>
        <el-menu-item index="/admin/news">News</el-menu-item>
        <el-menu-item index="/admin/people">People</el-menu-item>
        <el-menu-item index="/admin/undergraduates">Undergraduates</el-menu-item>
        <el-menu-item index="/admin/publications">Publications</el-menu-item>
        <el-menu-item index="/admin/patents">Patents</el-menu-item>
        <el-menu-item index="/admin/research-projects">Research Projects</el-menu-item>
        <el-menu-item index="/admin/media">Media</el-menu-item>
        <el-menu-item index="/admin/trash">Trash</el-menu-item>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header class="admin-header">
        <router-link to="/">查看官网</router-link>
        <el-button @click="handleLogout">退出登录</el-button>
      </el-header>
      <el-main class="admin-main">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { logout } from '../../api/admin'

const router = useRouter()

const handleLogout = async () => {
  logout()
  await router.push('/admin/login')
}
</script>

<style scoped>
.admin-shell {
  min-height: 100vh;
  background: #f3f4f6;
}

.admin-aside {
  background: #111827;
}

.admin-brand {
  height: 64px;
  display: flex;
  align-items: center;
  padding: 0 20px;
  color: white;
  font-size: 18px;
  font-weight: 700;
}

.admin-header {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 18px;
  background: white;
  border-bottom: 1px solid #e5e7eb;
}

.admin-main {
  padding: 24px;
}

.admin-header a {
  color: #7d1231;
  text-decoration: none;
}
</style>
