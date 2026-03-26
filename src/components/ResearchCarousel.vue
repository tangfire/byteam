<template>
  <div class="research-carousel">
    <!-- 桌面端 -->
    <div v-if="!isMobile" class="desktop-carousel">
      <div class="carousel-container">
        <div class="carousel-track" :style="{ transform: `translateX(-${currentIndex * 100}%)` }">
          <div v-for="pub in publications" :key="pub.id" class="carousel-slide">
            <div class="slide-card">
              <div class="slide-image">
                <img :src="pub.image" :alt="pub.title" @error="handleImageError" />
                <div class="badges">
                  <span class="badge type">{{ pub.type }}</span>
                  <span class="badge year">{{ pub.year }}</span>
                </div>
              </div>
              <div class="slide-info">
                <h3 class="title">{{ pub.title }}</h3>
                <p class="authors">{{ pub.authors }}</p>
                <p class="venue">{{ pub.venue }}</p>
                <div class="links">
                  <el-button v-if="pub.links.paper" type="primary" size="small" @click="openLink(pub.links.paper)">Paper</el-button>
                  <el-button v-if="pub.links.code" type="success" size="small" @click="openLink(pub.links.code)">Code</el-button>
                  <el-button v-if="pub.links.video" type="info" size="small" @click="handleVideoClick(pub.links.video)">Video</el-button>
                </div>
              </div>
            </div>
          </div>
        </div>
        <button class="nav prev" @click="prevSlide"><el-icon><ArrowLeft /></el-icon></button>
        <button class="nav next" @click="nextSlide"><el-icon><ArrowRight /></el-icon></button>
        <div class="indicators">
          <span v-for="(_, i) in publications" :key="i" class="dot" :class="{ active: i === currentIndex }" @click="goToSlide(i)"></span>
        </div>
      </div>
    </div>

    <!-- 移动端卡片列表 -->
    <div v-else class="mobile-carousel">
      <div class="mobile-track" ref="trackRef" @touchstart="handleTouchStart" @touchend="handleTouchEnd">
        <div v-for="pub in publications" :key="pub.id" class="mobile-card">
          <div class="card-image">
            <img :src="pub.image" :alt="pub.title" @error="handleImageError" />
            <div class="badges">
              <span class="badge type">{{ pub.type }}</span>
              <span class="badge year">{{ pub.year }}</span>
            </div>
          </div>
          <div class="card-info">
            <h4 class="title">{{ pub.title }}</h4>
            <p class="authors">{{ pub.authors }}</p>
            <p class="venue">{{ pub.venue }}</p>
            <div class="links">
              <el-button v-if="pub.links.paper" type="primary" size="small" @click="openLink(pub.links.paper)">Paper</el-button>
              <el-button v-if="pub.links.code" type="success" size="small" @click="openLink(pub.links.code)">Code</el-button>
              <el-button v-if="pub.links.video" type="info" size="small" @click="handleVideoClick(pub.links.video)">Video</el-button>
            </div>
          </div>
        </div>
      </div>
      <div class="mobile-indicators">
        <span v-for="(_, i) in publications" :key="i" class="dot" :class="{ active: i === currentIndex }" @click="goToSlide(i)"></span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { ArrowLeft, ArrowRight } from '@element-plus/icons-vue'

const router = useRouter()
const isMobile = ref(false)
const currentIndex = ref(0)
const touchStartX = ref(0)
const touchEndX = ref(0)

const publications = ref([
  {
    id: 1,
    image: '/publications/028.png',
    title: 'Multi-modal Brain Tumor Segmentation via Multi-category Interaction and Graph Co-reasoning',
    authors: 'Baoyao Yang*, Dongzhe Li, Chong Yin, Fei Lyu, Xiaochen He',
    venue: 'IEEE Transactions on Multimedia, 2025',
    type: 'Journal',
    year: '2025',
    links: { paper: 'https://ieeexplore.ieee.org/abstract/document/11205277' }
  },
  {
    id: 2,
    image: '/publications/026.png',
    title: 'CAM-interacted Vision GNN for Multi-label Medical Images',
    authors: 'Jiangchao Wang, Baoyao Yang*, Siqi Liu, Xiaoqi Zheng, Wenbin Yao* and Junxiang Chen',
    venue: 'IEEE Journal of Biomedical and Health Informatics, 2025',
    type: 'Journal',
    year: '2025',
    links: {
      code: 'https://github.com/BaoyaoGroup/JBHI_code',
      paper: 'https://ieeexplore.ieee.org/abstract/document/11205277'
    }
  },
  {
    id: 3,
    image: '/publications/027.png',
    title: 'FedCD: A Hybrid Federated Learning Framework for Adaptive Training under Data Heterogeneity',
    authors: 'Weide Zhan, Baoyao Yang*',
    venue: 'PRCV, 2025',
    type: 'Conference',
    year: '2025',
    links: {}
  },
  {
    id: 4,
    image: '/publications/004.png',
    title: 'Image-assisted Label Connective Completion for Vessel Segmentation with Insufficient Annotations',
    authors: 'Xiaoqi Zheng, Baoyao Yang*, Xiuwen Fang, Mang Ye',
    venue: 'ICASSP, 2025',
    type: 'Conference',
    year: '2025',
    links: {
      code: 'https://github.com/BaoyaoGroup/LabelCompletion',
      paper: 'https://ieeexplore.ieee.org/document/10888997',
      video: { name: 'video-player-XiaoqiZheng01' }
    }
  },
  {
    id: 5,
    image: '/publications/023.png',
    title: 'Simple but Effective: Sub-Volume Contrastive Learning for Class-Imbalanced Semi-Supervised 3D Medical Image Segmentation',
    authors: 'Xianrun Xu, Baoyao Yang*, Wanyun Li, Jingsong Lin, Yufei Xu',
    venue: 'ACM Multimedia, 2025',
    type: 'Conference',
    year: '2025',
    links: {
      paper: 'https://dl.acm.org/doi/abs/10.1145/3746027.3755652',
      video: { name: 'video-player-XianrunXu01' }
    }
  }
])

const handleImageError = (e: Event) => {
  (e.target as HTMLImageElement).src = '/publications/online.png'
}

const openLink = (url: string) => window.open(url, '_blank')

const handleVideoClick = (video: any) => {
  if (video.name) router.push({ name: video.name })
}

const nextSlide = () => {
  currentIndex.value = (currentIndex.value + 1) % publications.value.length
}

const prevSlide = () => {
  currentIndex.value = currentIndex.value === 0 ? publications.value.length - 1 : currentIndex.value - 1
}

const goToSlide = (index: number) => {
  currentIndex.value = index
}

const handleTouchStart = (e: TouchEvent) => {
  touchStartX.value = e.changedTouches[0].screenX
}

const handleTouchEnd = (e: TouchEvent) => {
  touchEndX.value = e.changedTouches[0].screenX
  const diff = touchStartX.value - touchEndX.value
  if (Math.abs(diff) > 50) {
    if (diff > 0) nextSlide()
    else prevSlide()
  }
}

const checkMobile = () => {
  isMobile.value = window.innerWidth <= 768
}

let interval: ReturnType<typeof setInterval>

onMounted(() => {
  checkMobile()
  window.addEventListener('resize', checkMobile)
  interval = setInterval(nextSlide, 5000)
})

onUnmounted(() => {
  window.removeEventListener('resize', checkMobile)
  clearInterval(interval)
})
</script>

<style scoped>
.research-carousel {
  width: 100%;
}

.desktop-carousel {
  position: relative;
  overflow: hidden;
  border-radius: 28px;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 249, 250, 0.95) 100%);
  /* 移除 backdrop-filter 以提升性能 */
  /* backdrop-filter: blur(20px); */
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.15), inset 0 1px 0 rgba(255, 255, 255, 0.95), 0 4px 12px rgba(0, 0, 0, 0.08);
  border: 2px solid rgba(125, 18, 49, 0.15);
  /* 启用硬件加速 */
  will-change: transform;
  -webkit-transform: translateZ(0);
  transform: translateZ(0);
}

.carousel-container {
  position: relative;
}

.carousel-track {
  display: flex;
  transition: transform 0.5s cubic-bezier(0.4, 0, 0.2, 1);
}

.carousel-slide {
  flex: 0 0 100%;
}

.slide-card {
  display: grid;
  grid-template-columns: 1fr 1fr;
  min-height: 450px;
  gap: 0;
  position: relative;
  overflow: hidden;
}

.slide-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(125, 18, 49, 0.04) 0%, rgba(19, 57, 62, 0.03) 100%);
  pointer-events: none;
}

.slide-card::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 4px;
  height: 100%;
  background: linear-gradient(180deg, var(--primary-color) 0%, #a51c41 50%, var(--primary-color) 100%);
  opacity: 0;
  transition: opacity 0.5s ease;
}

.desktop-carousel:hover .slide-card::after {
  opacity: 1;
}

.slide-image {
  position: relative;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.8) 0%, rgba(248, 249, 250, 0.7) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2.5rem;
  border-right: 1px solid rgba(125, 18, 49, 0.1);
}

.slide-image::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(circle at 50% 30%, rgba(125, 18, 49, 0.05) 0%, transparent 60%);
  pointer-events: none;
}

.slide-image img {
  max-width: 100%;
  max-height: 300px;
  object-fit: contain;
  border-radius: 16px;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.15), 0 4px 12px rgba(0, 0, 0, 0.08);
  transition: transform 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
}

.slide-image:hover img {
  transform: scale(1.08) rotate(-2deg);
  box-shadow: 0 16px 50px rgba(0, 0, 0, 0.2);
}

.badges {
  position: absolute;
  top: 1.5rem;
  left: 1.5rem;
  display: flex;
  gap: 0.5rem;
  z-index: 2;
}

.badge {
  padding: 0.5rem 1.4rem;
  border-radius: 24px;
  font-size: 0.85rem;
  font-weight: 700;
  color: white;
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.2), inset 0 1px 0 rgba(255, 255, 255, 0.3);
  /* 移除 backdrop-filter 以提升性能 */
  /* backdrop-filter: blur(10px); */
  transition: all 0.4s cubic-bezier(0.34, 1.56, 0.64, 1);
  border: 2px solid rgba(255, 255, 255, 0.3);
  /* 启用硬件加速 */
  will-change: transform;
  -webkit-transform: translateZ(0);
  transform: translateZ(0);
}

.badge:hover {
  transform: translateY(-3px) scale(1.08);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.3), inset 0 1px 0 rgba(255, 255, 255, 0.4);
}

.type {
  background: linear-gradient(135deg, var(--primary-color) 0%, #a51c41 100%);
}

.year {
  background: linear-gradient(135deg, #3498db 0%, #2980b9 100%);
}

.slide-info {
  padding: 3rem 2.5rem;
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 1rem;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.9) 0%, rgba(248, 249, 250, 0.85) 100%);
}

.title {
  font-size: 1.5rem;
  font-weight: 800;
  background: linear-gradient(135deg, var(--primary-color) 0%, #a51c41 50%, #7d1231 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 0.8rem;
  line-height: 1.4;
  letter-spacing: -0.3px;
}

.authors {
  color: var(--primary-color);
  font-weight: 600;
  margin-bottom: 0.5rem;
  font-size: 1rem;
  line-height: 1.5;
}

.venue {
  color: var(--text-secondary);
  font-style: italic;
  margin-bottom: 1.5rem;
  font-weight: 500;
  opacity: 0.9;
}

.links {
  display: flex;
  gap: 1rem;
  flex-wrap: wrap;
}

.nav {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  width: 56px;
  height: 56px;
  border: none;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 249, 250, 0.95) 100%);
  box-shadow: 0 10px 32px rgba(0, 0, 0, 0.18), inset 0 1px 0 rgba(255, 255, 255, 0.95), 0 4px 12px rgba(0, 0, 0, 0.1);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--primary-color);
  transition: all 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
  z-index: 10;
  border: 2px solid rgba(125, 18, 49, 0.15);
}

.nav:hover {
  background: linear-gradient(135deg, var(--primary-color) 0%, #a51c41 100%);
  color: white;
  transform: translateY(-50%) scale(1.2) rotate(180deg);
  box-shadow: 0 16px 48px rgba(125, 18, 49, 0.4), inset 0 1px 0 rgba(255, 255, 255, 0.3);
}

.prev {
  left: 1rem;
}

.next {
  right: 1rem;
}

.indicators {
  position: absolute;
  bottom: 1rem;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  gap: 0.5rem;
}

.dot {
  width: 14px;
  height: 14px;
  border-radius: 50%;
  background: linear-gradient(135deg, #d0d7de 0%, #b8bcc2 100%);
  cursor: pointer;
  transition: all 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
  border: 2px solid rgba(125, 18, 49, 0.25);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1), inset 0 1px 0 rgba(255, 255, 255, 0.5);
}

.dot.active {
  background: linear-gradient(135deg, var(--primary-color) 0%, #a51c41 100%);
  transform: scale(1.6);
  box-shadow: 0 6px 16px rgba(125, 18, 49, 0.5), inset 0 1px 0 rgba(255, 255, 255, 0.4);
  border-color: transparent;
  animation: pulse 2s ease-in-out infinite;
}

@keyframes pulse {
  0%, 100% {
    box-shadow: 0 6px 16px rgba(125, 18, 49, 0.5), inset 0 1px 0 rgba(255, 255, 255, 0.4);
  }
  50% {
    box-shadow: 0 6px 24px rgba(125, 18, 49, 0.7), inset 0 1px 0 rgba(255, 255, 255, 0.4);
  }
}

/* 移动端 */
.mobile-carousel {
  padding: 0 0.5rem;
}

.mobile-track {
  display: flex;
  overflow-x: auto;
  scroll-snap-type: x mandatory;
  gap: 1rem;
  padding-bottom: 1rem;
  scrollbar-width: none;
}

.mobile-track::-webkit-scrollbar {
  display: none;
}

.mobile-card {
  flex: 0 0 90%;
  scroll-snap-align: start;
  background: white;
  border-radius: var(--border-radius);
  box-shadow: var(--shadow-soft);
  overflow: hidden;
}

.card-image {
  position: relative;
  background: var(--bg-light);
  padding: 1.5rem;
  text-align: center;
}

.card-image img {
  max-width: 100%;
  max-height: 180px;
  object-fit: contain;
}

.card-info {
  padding: 1.5rem;
}

.mobile-indicators {
  display: flex;
  justify-content: center;
  gap: 0.5rem;
  margin-top: 1rem;
}

@media (max-width: 768px) {
  .slide-card {
    grid-template-columns: 1fr;
    min-height: auto;
  }

  .slide-info {
    padding: 1.5rem;
  }

  .title {
    font-size: 1.2rem;
  }

  .nav {
    width: 40px;
    height: 40px;
  }
}
</style>