<template>
  <el-empty v-if="hidden" description="页面暂未发布" />
  <div v-else class="video-container">
    <el-card style="max-width: 1200px; margin: 20px auto;">
      <div class="navigation-control">
        <el-button
          type="primary"
          @click="goBack"
          style="background-color: #7d1231; border-color: #7d1231; margin-bottom: 20px"
        >
          <el-icon :size="24"><ArrowLeft /></el-icon>
        </el-button>
      </div>

      <h2 style="text-align: center;color: #7d1231">{{ content.title }}</h2>

      <video controls width="100%">
        <source :src="content.video" type="video/mp4">
        Your browser does not support the video tag.
      </video>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ArrowLeft } from '@element-plus/icons-vue'
import { useSitePage } from '../composables/useSitePage'

const props = defineProps<{
  slug: string
  fallbackTitle: string
  fallbackVideo: string
}>()

const router = useRouter()
const { content, hidden } = useSitePage(props.slug, {
  title: props.fallbackTitle,
  video: props.fallbackVideo,
})

onMounted(() => {
  window.scrollTo(0, 0)
  document.documentElement.scrollTop = 0
})

const goBack = () => {
  if (window.history.state?.back) {
    router.go(-1)
  } else {
    router.push({ name: 'home' })
  }
}
</script>

<style scoped>
.video-container {
  padding: 20px;
  min-height: 100vh;
}

.navigation-control {
  display: flex;
  justify-content: flex-start;
}

.el-button:hover {
  opacity: 0.9;
  transform: translateY(-1px);
  transition: all 0.3s ease;
}

@media (max-width: 768px) {
  .el-card {
    margin: 10px !important;
  }

  h2 {
    font-size: 0.8rem !important;
    padding: 0 10px;
  }

  .el-button {
    padding: 8px 12px;
    font-size: 14px;
  }
}
</style>
