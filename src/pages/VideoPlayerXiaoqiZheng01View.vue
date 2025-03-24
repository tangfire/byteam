<template>
  <div class="video-container">
    <el-card style="max-width: 1200px; margin: 20px auto;">
      <!-- 返回按钮 -->
      <div class="navigation-control">
        <el-button
            type="primary"
            @click="goBack"
            style="background-color: #7d1231; border-color: #7d1231; margin-bottom: 20px"
        >
          <el-icon :size="24"><ArrowLeft /></el-icon>

        </el-button>
      </div>

      <h2 style="text-align: center;color: #7d1231">Image-assisted Label Connective Completion for Vessel Segmentation with Insufficient Annotations</h2>

      <video
          controls
          width="100%"
      >
        <source src="/videos/001.mp4" type="video/mp4">
        Your browser does not support the video tag.
      </video>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import {ArrowLeft} from '@element-plus/icons-vue'
import { onMounted } from 'vue'

const router = useRouter()


// 新增：组件挂载时滚动到顶部
onMounted(() => {
  window.scrollTo(0, 0)
  document.documentElement.scrollTop = 0
})


const goBack = () => {
  // 优先使用历史记录返回
  if (window.history.state?.back) {
    router.go(-1)
  } else {
    // 如果没有历史记录则返回首页
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

/* 按钮悬停效果 */
.el-button:hover {
  opacity: 0.9;
  transform: translateY(-1px);
  transition: all 0.3s ease;
}

/* 移动端适配 */
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