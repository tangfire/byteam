<script setup lang="ts">
import { ArrowRightBold } from "@element-plus/icons-vue";
import { computed } from 'vue'
import { useSitePage } from '../composables/useSitePage'

defineProps<{ msg: string }>()

const { content, hidden } = useSitePage('research-direction', { directions: [] })
const directions = computed(() => Array.isArray(content.value.directions) ? content.value.directions : [])
</script>

<template>
  <div class="research-container">
    <el-empty v-if="hidden" description="页面暂未发布" />
    <template v-else>
    <el-space direction="vertical" :size="30" style="width: 100%">

      <div v-for="(direction, index) in directions" :key="direction.title" class="research-item">
        <div class="research-header">
          <div class="title-wrapper">
            <el-icon class="header-icon">
              <ArrowRightBold/>
            </el-icon>
            <h2 class="research-title">{{ direction.title }}</h2>
          </div>
        </div>

        <el-card class="research-card" shadow="hover">
          <div class="card-content" :class="{ 'reverse-layout': index % 2 === 1 }">
            <div class="image-wrapper">
              <img
                  :src="direction.image"
                  :alt="direction.alt || direction.title"
                  class="research-image"
              />
            </div>
            <div class="text-content">
              <div v-for="section in direction.sections || []" :key="section.title" class="content-section">
                <h3 class="section-title">{{ section.title }}</h3>
                <p class="content-text">{{ section.text }}</p>
              </div>
            </div>
          </div>
        </el-card>
      </div>

    </el-space>

    <!-- 底部间隔 -->
    <div style="height: 50px"></div>

    <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
    </template>
  </div>
</template>

<style scoped>
.research-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.research-item {
  margin-bottom: 40px;
}

.research-header {
  margin-bottom: 20px;
}

.title-wrapper {
  display: flex;
  align-items: center;
}

.header-icon {
  color: #7d1231;
  font-size: 28px;
  margin-right: 12px;
  flex-shrink: 0;
}

.research-title {
  color: #2c3e50;
  font-size: 1.8rem;
  font-weight: 600;
  margin: 0;
  background: linear-gradient(135deg, #7d1231, #3498db);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.research-card {
  max-width: 1000px;
  margin: 0 auto;
  border: none;
  border-radius: 16px;
  transition: all 0.3s ease;
  background: #fff;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  overflow: hidden;
}

.research-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.12);
}

.card-content {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  padding: 20px;
}

.image-wrapper {
  position: relative;
  flex: 0 0 400px;
  margin-right: 30px;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
}

.research-image {
  width: 100%;
  height: auto;
  max-height: 250px;
  object-fit: contain;
  display: block;
}

.text-content {
  flex: 1;
  padding: 10px;
}

.content-section {
  margin-bottom: 20px;
}

.section-title {
  color: #7d1231;
  font-size: 1.2rem;
  font-weight: 600;
  margin-bottom: 8px;
  position: relative;
  padding-left: 16px;
}

.section-title::before {
  content: '';
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 8px;
  height: 8px;
  background: #7d1231;
  border-radius: 50%;
}

.content-text {
  color: #5a6c7d;
  line-height: 1.6;
  font-size: 16px;
  margin: 0;
}

/* 交替布局 - 桌面端 */
@media (min-width: 769px) {
  .reverse-layout {
    flex-direction: row-reverse;
  }

  .reverse-layout .image-wrapper {
    margin-right: 0;
    margin-left: 30px;
  }
}

/* 移动端适配 */
@media (max-width: 768px) {
  .research-container {
    padding: 15px;
  }

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }

  .research-title {
    font-size: 1.3rem;
    line-height: 1.4;
  }

  .header-icon {
    font-size: 22px;
    margin-right: 10px;
  }

  .card-content {
    flex-direction: column;
    padding: 15px;
  }

  .image-wrapper {
    flex: none;
    width: 100%;
    max-width: 100%;
    margin-right: 0;
    margin-left: 0;
    margin-bottom: 20px;
  }

  .research-image {
    width: 100%;
    height: auto;
    max-height: 220px;
    object-fit: contain;
  }

  .text-content {
    width: 100%;
    padding: 0;
  }

  .section-title {
    font-size: 1.1rem;
  }

  .content-text {
    font-size: 14px;
    line-height: 1.5;
  }
}

/* 小屏手机适配 */
@media (max-width: 480px) {
  .research-container {
    padding: 10px;
  }

  .research-title {
    font-size: 1.1rem;
  }

  .research-card {
    border-radius: 12px;
  }

  .card-content {
    padding: 12px;
  }

  .research-image {
    max-height: 180px;
  }
}

/* 大屏幕优化 */
@media (min-width: 1400px) {
  .research-container {
    max-width: 1300px;
  }

  .research-card {
    max-width: 1100px;
  }
}

::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}
</style>
