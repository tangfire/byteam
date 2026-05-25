<script setup lang="ts">
import { computed } from 'vue'
import { useSitePage } from '../composables/useSitePage'

const { content, hidden } = useSitePage('dr-baoyao-yang', {
  name: 'Dr. Baoyao Yang(Associate Professor)',
  image: '/currentprojects/yangbaoyao.jpg',
  alt: 'Dr. Baoyao Yang',
  paragraphs: [],
})
const paragraphs = computed(() => Array.isArray(content.value.paragraphs) ? content.value.paragraphs : [])
</script>

<template>
  <el-empty v-if="hidden" description="页面暂未发布" />
  <el-card v-else class="profile-card">
    <template #header>
      <h2 class="teachName">{{ content.name }}</h2>
    </template>

    <div class="profile-content">
      <img
          :src="content.image"
          :alt="content.alt"
          class="profile-image"
      />
      <div class="profile-description">
        <p v-for="paragraph in paragraphs" :key="paragraph">{{ paragraph }}</p>
      </div>
    </div>
  </el-card>
</template>

<style scoped>


@media (max-width: 768px) {
  .profile-content {
    flex-direction: column;
    align-items: center;
  }

  .profile-image {
    width: 100%;
    max-width: 400px;
  }

  .teachName{
    font-size: 1.2rem !important;
  }


}

.profile-card {
  max-width: 1000px;
  margin: 2rem auto;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.profile-card h2 {
  color: #2c3e50;
  font-size: 1.8rem;
  font-weight: bold;
  /*border-bottom: 2px solid #7d1231;*/
  padding-bottom: 0.5em;
  margin-bottom: 0.5em;
}

.profile-content {
  display: flex;
  flex-wrap: wrap;
  gap: 2rem;
  align-items: flex-start;
}

.profile-image {
  width: 300px;
  height: auto;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.profile-description {
  flex: 1;
  font-family: 'Segoe UI', Roboto, sans-serif;
  line-height: 1.8;
  color: #2d3436;
  font-size: 1rem;
  text-align: justify;  /* 添加两端对齐 */
  hyphens: auto;        /* 允许自动断字 */
  word-break: break-word; /* 长单词强制换行 */
}

.profile-description p {
  margin-bottom: 1.2em;
}

.profile-description p:last-child {
  margin-bottom: 0;
}


</style>
