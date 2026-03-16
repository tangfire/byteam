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
  border-radius: var(--border-radius-lg);
  background: white;
  box-shadow: var(--shadow-medium);
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
  min-height: 400px;
}

.slide-image {
  position: relative;
  background: var(--bg-light);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem;
}

.slide-image img {
  max-width: 100%;
  max-height: 280px;
  object-fit: contain;
  border-radius: var(--border-radius-sm);
}

.badges {
  position: absolute;
  top: 1.5rem;
  left: 1.5rem;
  display: flex;
  gap: 0.5rem;
}

.badge {
  padding: 0.3rem 1rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
  color: white;
}

.type {
  background: var(--gradient-primary);
}

.year {
  background: var(--gradient-accent);
}

.slide-info {
  padding: 2.5rem;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.title {
  font-size: 1.4rem;
  font-weight: 700;
  color: var(--text-primary);
  margin-bottom: 1rem;
  line-height: 1.4;
}

.authors {
  color: var(--primary-color);
  font-weight: 500;
  margin-bottom: 0.5rem;
}

.venue {
  color: var(--text-light);
  font-style: italic;
  margin-bottom: 1.5rem;
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
  width: 48px;
  height: 48px;
  border: none;
  border-radius: 50%;
  background: white;
  box-shadow: var(--shadow-soft);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--primary-color);
  transition: all var(--transition-fast);
  z-index: 10;
}

.nav:hover {
  background: var(--primary-color);
  color: white;
  transform: translateY(-50%) scale(1.1);
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
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: #d0d7de;
  cursor: pointer;
  transition: all var(--transition-fast);
}

.dot.active {
  background: var(--primary-color);
  transform: scale(1.2);
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