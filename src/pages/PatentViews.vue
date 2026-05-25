<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { getPublicPatents } from '../api/public'
import type { Patent } from '../api/client'

const patents = ref<Patent[]>([
  { authors: '杨宝瑶，黄彦浩，陈涤新', title: '一种基于时空信息聚合的视频特征提取模型训练方法、系统及特征提取方法', date: '2025-09-09', country: '中国', number: 'ZL202510359255.5', category: 'granted', status: 'published', sortOrder: 1 },
  { authors: '杨宝瑶，麻亚利，詹伟德，唐彦超，卢泽坚', title: '一种联邦学习场景下的检测噪声标注的方法及系统', date: '2025-09-05', country: '中国', number: 'ZL202510375049.3', category: 'granted', status: 'published', sortOrder: 2 },
  { authors: '杨宝瑶，陈俊祥，黄彦浩，姚文彬', title: '一种基于多模态大模型的视频 - 文本检索方法', date: '2025-06-13', country: '中国', number: 'ZL202411271756.X', category: 'granted', status: 'published', sortOrder: 3 },
  { authors: '杨宝瑶，郑晓琦', title: '一种噪声标注的血管图像分割方法及系统', date: '2025-02-11', country: '中国', number: 'ZL 202311792631.7', category: 'granted', status: 'published', sortOrder: 4 },
  { authors: '杨宝瑶，詹伟德', title: '一种模型异构性联邦学习方法和系统', date: '', country: '中国', number: 'ZL202210989290.1', category: 'review', status: 'published', sortOrder: 14 },
  { authors: '李春林，吴恒，曾安，惠恩明，富锐，杨思维，曹洪江，骆有隆，杨宝瑶，刘俊，张勇，江焜', title: '产业聚集区域内业务资源服务平台规范', date: '2025-05-01', country: '广东省工业软件学会团体标准', number: 'T/GISF 002-2024', category: 'standard', status: 'published', sortOrder: 30 },
])

const sections = computed(() => [
  { key: 'granted', title: '授权专利 (Granted Patents)', items: patents.value.filter((item) => item.category === 'granted') },
  { key: 'review', title: '实审专利 (Patents Under Review)', items: patents.value.filter((item) => item.category === 'review') },
  { key: 'standard', title: '团体标准 (Group Standards)', items: patents.value.filter((item) => item.category === 'standard') },
])

const statusLabel = (category: string) => {
  if (category === 'granted') return '授权'
  if (category === 'review') return '实审'
  return '标准'
}

onMounted(async () => {
  try {
    patents.value = (await getPublicPatents()).items
  } catch (error) {
    console.warn('Using local patent fallback data', error)
  }
})
</script>

<template>
  <el-space direction="vertical" :size="30" style="width: 100%">
    <div class="patent-container">
      <h1 class="patent-title">Patents</h1>

      <el-card class="patent-card">
        <div class="patent-content">
          <div v-for="section in sections" :key="section.key" class="patent-section">
            <h2 class="section-subtitle">{{ section.title }}</h2>
            <ul class="patent-list">
              <li v-for="patent in section.items" :key="patent.title">
                <div class="patent-item">
                  <div class="patent-authors">{{ patent.authors }}</div>
                  <div class="patent-title-desc">{{ patent.title }}</div>
                  <div class="patent-details">
                    <span v-if="patent.date" class="patent-date">{{ patent.date }}</span>
                    <span v-if="patent.country" class="patent-country">{{ patent.country }}</span>
                    <span v-if="patent.number" class="patent-number">{{ patent.number }}</span>
                    <span class="patent-status" :class="patent.category">{{ statusLabel(patent.category) }}</span>
                  </div>
                </div>
              </li>
            </ul>
          </div>
        </div>
      </el-card>
    </div>
  </el-space>

  <div class="bottom-spacing"></div>
  <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
</template>

<style scoped>
.patent-container {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px 0;
}

.patent-title {
  color: #7d1231;
  font-size: 32px;
  font-weight: 600;
  margin-bottom: 20px;
  text-align: center;
  position: relative;
  padding-bottom: 10px;
}

.patent-title::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 80px;
  height: 3px;
  background: linear-gradient(90deg, transparent, #7d1231, transparent);
}

.patent-card {
  max-width: 1000px;
  width: 90%;
  margin: 0 auto;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  border: none;
}

.patent-content {
  padding: 10px 5px;
}

.patent-section {
  margin-bottom: 40px;
}

.section-subtitle {
  color: #13393e;
  font-size: 22px;
  font-weight: 600;
  margin-bottom: 20px;
  padding-left: 15px;
  border-left: 4px solid #7d1231;
}

.patent-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.patent-item {
  padding: 18px 20px;
  margin-bottom: 15px;
  border-radius: 8px;
  background: #f8f9fa;
  transition: all 0.3s ease;
  border-left: 3px solid transparent;
}

.patent-item:hover {
  transform: translateX(5px);
  border-left-color: #7d1231;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.patent-authors {
  color: #7d1231;
  font-weight: 600;
  margin-bottom: 8px;
}

.patent-title-desc {
  color: #2c3e50;
  font-size: 16px;
  line-height: 1.5;
  margin-bottom: 12px;
}

.patent-details {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  align-items: center;
}

.patent-details span {
  padding: 4px 10px;
  border-radius: 6px;
  font-size: 13px;
  background: white;
  color: #5f6670;
}

.patent-status.granted {
  background: rgba(39, 174, 96, 0.1);
  color: #27ae60;
}

.patent-status.review {
  background: rgba(243, 156, 18, 0.1);
  color: #d68910;
}

.patent-status.standard {
  background: rgba(52, 152, 219, 0.1);
  color: #2980b9;
}

.bottom-spacing {
  height: 50px;
}

@media (max-width: 768px) {
  .patent-card {
    width: 95%;
  }

  .patent-details {
    align-items: flex-start;
    flex-direction: column;
  }

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }
}
</style>
