<template>
  <main class="admin-login">
    <el-card class="login-card">
      <h1>BYML Admin</h1>
      <p>登录后管理官网内容</p>
      <el-form :model="form" @submit.prevent="submit">
        <el-form-item>
          <el-input v-model="form.username" placeholder="用户名" autocomplete="username" />
        </el-form-item>
        <el-form-item>
          <el-input v-model="form.password" placeholder="密码" type="password" autocomplete="current-password" show-password />
        </el-form-item>
        <el-button type="primary" native-type="submit" :loading="loading" class="login-button">登录</el-button>
      </el-form>
    </el-card>
  </main>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { login } from '../../api/admin'

const route = useRoute()
const router = useRouter()
const loading = ref(false)
const form = reactive({ username: 'admin', password: 'admin123456' })

const submit = async () => {
  loading.value = true
  try {
    await login(form.username, form.password)
    ElMessage.success('登录成功')
    await router.push(String(route.query.redirect || '/admin/dashboard'))
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.admin-login {
  min-height: 100vh;
  display: grid;
  place-items: center;
  background: linear-gradient(135deg, #f7f8fb, #edf2f7);
}

.login-card {
  width: min(420px, calc(100vw - 32px));
  border-radius: 8px;
}

.login-card h1 {
  margin: 0;
  color: #7d1231;
}

.login-card p {
  margin: 8px 0 24px;
  color: #6b7280;
}

.login-button {
  width: 100%;
}
</style>
