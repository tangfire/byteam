<script setup lang="ts">
import { computed } from 'vue'
import { useSitePage } from '../composables/useSitePage'

const { content, hidden } = useSitePage('vknow', {
  title: 'vKnow: Decode Omni-modal Intent',
  introduction: { title: 'Introduction', text: '', items: [] },
  dataset: { title: '', description: '', image: '', imageAlt: '', linkText: 'Learn more →', linkPath: '/VideoMind' },
})
const introItems = computed(() => Array.isArray(content.value.introduction?.items) ? content.value.introduction.items : [])
</script>

<template>
  <el-empty v-if="hidden" description="页面暂未发布" />
  <template v-else>
  <el-space direction="vertical" :size="30" style="width: 100%">
    <div style="display: flex; flex-direction: column; align-items: center;">
      <h1 class="section-title">{{ content.title }}</h1>

      <!-- 项目介绍卡片 -->
      <el-card class="intro-card">
        <div class="introduction">
          <h3 class="subsection-title first-subtitle">{{ content.introduction?.title }}</h3>
          <p class="vision-text">
            {{ content.introduction?.text }}
          </p>
          <ul class="styled-list">
            <li v-for="item in introItems" :key="item">{{ item }}</li>
          </ul>
        </div>
      </el-card>

      <!-- 数据集介绍 -->
      <el-card class="dataset-card">
        <div class="dataset-card-content">
          <!-- 图片左边 -->
          <img :src="content.dataset?.image" :alt="content.dataset?.imageAlt" class="dataset-card-image" />
          <!-- 文字右边 -->
          <div class="dataset-card-info">
            <h3 class="dataset-title">{{ content.dataset?.title }}</h3>
            <p class="dataset-desc">
              {{ content.dataset?.description }}
            </p>
            <div class="dataset-link-row">
              <router-link :to="content.dataset?.linkPath || '/videomind'" class="dataset-link">{{ content.dataset?.linkText }}</router-link>
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
</template>

<style scoped>

@media (max-width: 768px) {
  .section-title {
    font-size: 1.35rem !important;
  }
  .subsection-title{
    font-size: 1.2rem !important;
    margin-top: -1rem !important;
  }
  /* 数据集内容 */
  .dataset-card-content {
    flex-direction: column;
    padding: 1rem 0.5rem;
  }
  /* 数据集图片 */
  .dataset-card-image {
    width: 90vw;
    max-width: 95vw;
    height: auto;
    margin-right: 0px;
  }
  /* 数据集信息区 */
  .dataset-card-info {
    width: 100%;
    align-items: center;
    text-align: center;
  }
  /* Learn more 按钮 */
  .dataset-link-row {
    justify-content: center;
  }
  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }
  .intro-card:hover,
  .dataset-card:hover {
    box-shadow: 0 4px 16px rgba(125,18,49,0.10);
    transform: none;
  }
  .dataset-link:hover {
    transform: none;
    box-shadow: 0 2px 8px rgba(125,18,49,0.08);
  }
}

/* 统一字体设置 */
.project-content {
  font-family: 'Segoe UI', system-ui, sans-serif;
  line-height: 1.7;
  color: #2c3e50;
}

/* 主标题样式 */
.section-title {
  color: #7d1231;
  font-size: 2.2rem;
  margin: 1rem 0 2rem;
  font-weight: 600;
  letter-spacing: -0.5px;
  text-align: center;
}

/* 小标题 */
.subsection-title {
  color: #7d1231;
  font-size: 1.5rem;
  margin: 0.5rem 0 1.5rem;
  padding-bottom: 0.5rem;
  border-bottom: 2px solid #eee;
}

/* 高亮文字 */
.highlight {
  color: #7d1231;
  font-weight: 600;
}

/* 项目介绍 */
.intro-card {
  max-width: 1000px;
  margin: 0 auto 1rem auto;
  box-shadow: 0 4px 16px rgba(44,62,80,0.10);
  border-radius: 14px;
  border: none;
}
.intro-card .introduction {
  padding: 1rem 1.5rem;
}

/* 列表样式 */
.styled-list {
  list-style: none;
  padding-left: 1.2rem;
  margin: 0.8rem 0;
}
.styled-list li {
  position: relative;
  padding-left: 1.2rem;
  margin-bottom: 0.8rem;
}
.styled-list li::before {
  content: "•";
  color: #7d1231;
  position: absolute;
  left: 0;
  font-weight: bold;
}

/* 数据集 */
.dataset-card {
  max-width: 1000px;
  margin: 2.5rem auto 1.5rem auto;
  box-shadow: 0 4px 16px rgba(44,62,80,0.10);
  border-radius: 14px;
  padding: 0;
  border: none;
}
/* 数据集内容 */
.dataset-card-content {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  padding: 1.2rem 1.5rem;
  gap: 2.2rem;
}
/* 数据集图片 */
.dataset-card-image {
  width: 320px;
  height: 200px;
  object-fit: cover;
  border-radius: 10px;
  box-shadow: 0 2px 12px rgba(44,62,80,0.10);
  background: #fff;
}
/* 数据集右侧信息区 */
.dataset-card-info {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-start;
  flex: 1;
}
.dataset-title {
  color: #7d1231;
  font-size: 1.2rem;
  font-weight: 700;
  margin-bottom: 0.5rem;
}
.dataset-desc {
  font-size: 1.05rem;
  color: #2c3e50;
  margin-bottom: 0.7rem;
}
/* Learn more 按钮右对齐行 */
.dataset-link-row {
  width: 100%;
  display: flex;
  justify-content: flex-end;
  margin-top: 0.5rem;
}

/* Learn more 按钮 */
.dataset-link {
  color: #fff;
  background: #7d1231;
  padding: 0.45em 1.1em;
  border-radius: 6px;
  font-size: 1rem;
  font-weight: 500;
  text-decoration: none;
  display: inline-block;
  position: relative;
  overflow: hidden;
  transition: background 0.18s, color 0.18s, transform 0.18s, box-shadow 0.18s;
  box-shadow: 0 2px 8px rgba(125,18,49,0.08);
}
.dataset-link:hover {
  background: #a41e44;
  color: #fff;
  transform: scale(1.06) translateY(-2px);
  box-shadow: 0 4px 18px rgba(125,18,49,0.16);
}
.dataset-link::after {
  content: '';
  position: absolute;
  left: -60%;
  top: 0;
  width: 60%;
  height: 100%;
  background: linear-gradient(120deg, rgba(255,255,255,0.18) 0%, rgba(255,255,255,0.01) 100%);
  transform: skewX(-20deg);
  transition: left 0.4s cubic-bezier(0.23, 1, 0.32, 1);
  pointer-events: none;
}
.dataset-link:hover::after {
  left: 110%;
}

/* 动画关键帧：卡片淡入上移 */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(40px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 卡片加载动画和悬浮动态效果 */
.intro-card,
.dataset-card {
  animation: fadeInUp 0.8s cubic-bezier(0.23, 1, 0.32, 1);
  transition: box-shadow 0.22s, transform 0.22s;
}
.intro-card:hover,
.dataset-card:hover {
  box-shadow: 0 8px 32px rgba(125,18,49,0.13);
  transform: translateY(-4px) scale(1.025);
}

</style>
