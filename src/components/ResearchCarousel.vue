
<template>
  <div class="research-carousel">
    <!-- 桌面端轮播 -->
    <div class="desktop-carousel" v-if="!isMobile">
      <div class="carousel-container">
        <div class="carousel-track" :style="trackStyle">
          <div
              v-for="(publication, index) in publications"
              :key="publication.id"
              class="carousel-slide"
              :class="{ active: currentIndex === index }"
          >
            <div class="slide-content">
              <div class="image-container">
                <img
                    :src="publication.image"
                    :alt="publication.title"
                    class="publication-image"
                    @error="handleImageError"
                />
                <div class="publication-badges">
                  <span class="badge type-badge">{{ publication.type }}</span>
                  <span class="badge year-badge">{{ publication.year }}</span>
                </div>
              </div>

              <div class="content-container">
                <h3 class="publication-title">{{ publication.title }}</h3>
                <p class="publication-authors">{{ publication.authors }}</p>
                <p class="publication-venue">{{ publication.venue }}</p>

                <div class="publication-links">
                  <el-button
                      v-if="publication.links.paper"
                      type="primary"
                      :icon="Document"
                      @click="openLink(publication.links.paper)"
                      class="link-button"
                  >
                    Paper
                  </el-button>

                  <el-button
                      v-if="publication.links.code"
                      type="success"
                      :icon="Promotion"
                      @click="openLink(publication.links.code)"
                      class="link-button"
                  >
                    Code
                  </el-button>

                  <el-button
                      v-if="publication.links.video"
                      type="info"
                      :icon="VideoPlay"
                      @click="handleVideoClick(publication.links.video)"
                      class="link-button"
                  >
                    Video
                  </el-button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 导航按钮 - 移除禁用状态 -->
        <button class="nav-button prev-button" @click="prevSlide">
          <el-icon><ArrowLeft /></el-icon>
        </button>
        <button class="nav-button next-button" @click="nextSlide">
          <el-icon><ArrowRight /></el-icon>
        </button>

        <!-- 指示器 -->
        <div class="carousel-indicators">
          <button
              v-for="(_, index) in publications"
              :key="index"
              class="indicator"
              :class="{ active: currentIndex === index }"
              @click="goToSlide(index)"
          ></button>
        </div>
      </div>
    </div>

    <!-- 移动端卡片列表 -->
    <div class="mobile-carousel" v-else>
      <div class="mobile-slides"
           @touchstart="handleTouchStart"
           @touchend="handleTouchEnd">
        <div
            v-for="(publication, index) in publications"
            :key="publication.id"
            class="mobile-slide"
            :class="{ active: currentIndex === index }"
        >
          <div class="mobile-card">
            <div class="card-image">
              <img
                  :src="publication.image"
                  :alt="publication.title"
                  @error="handleImageError"
              />
              <div class="mobile-badges">
                <span class="badge type-badge">{{ publication.type }}</span>
                <span class="badge year-badge">{{ publication.year }}</span>
              </div>
            </div>

            <div class="card-content">
              <h4 class="mobile-title">{{ publication.title }}</h4>
              <p class="mobile-authors">{{ publication.authors }}</p>
              <p class="mobile-venue">{{ publication.venue }}</p>

              <div class="mobile-links">
                <el-button
                    v-if="publication.links.paper"
                    type="primary"
                    size="small"
                    @click="openLink(publication.links.paper)"
                    class="mobile-link"
                >
                  Paper
                </el-button>
                <el-button
                    v-if="publication.links.code"
                    type="success"
                    size="small"
                    @click="openLink(publication.links.code)"
                    class="mobile-link"
                >
                  Code
                </el-button>
                <el-button
                    v-if="publication.links.video"
                    type="info"
                    size="small"
                    @click="handleVideoClick(publication.links.video)"
                    class="mobile-link"
                >
                  Video
                </el-button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="mobile-indicators">
        <button
            v-for="(_, index) in publications"
            :key="index"
            class="mobile-indicator"
            :class="{ active: currentIndex === index }"
            @click="goToSlide(index)"
        ></button>
      </div>
    </div>

    <!-- 查看全部按钮 -->
    <div class="view-all-section">
      <el-button
          type="primary"
          @click="goToPublications"
          class="view-all-button"
          size="large"
      >
        View All Publications
        <el-icon><ArrowRight /></el-icon>
      </el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { Document, Promotion, VideoPlay, ArrowRight, ArrowLeft } from '@element-plus/icons-vue'

const router = useRouter()
const isMobile = ref(false)
const currentIndex = ref(0)

const publications = ref([
  {
    id: 1,
    image: "/publications/028.png",
    title: "Multi-modal Brain Tumor Segmentation via Multi-category Interaction and Graph Co-reasoning",
    authors: "Baoyao Yang*, Dongzhe Li, Chong Yin, Fei Lyu, Xiaochen He",
    venue: "IEEE Transactions on Multimedia, 2025",
    type: "Journal",
    year: "2025",
    links: {
      paper: "https://ieeexplore.ieee.org/abstract/document/11205277"
    }
  },
  {
    id: 2,
    image: "/publications/026.png",
    title: "CAM-interacted Vision GNN for Multi-label Medical Images",
    authors: "Jiangchao Wang, Baoyao Yang*, Siqi Liu, Xiaoqi Zheng, Wenbin Yao* and Junxiang Chen",
    venue: "IEEE Journal of Biomedical and Health Informatics (JBHI), 2025",
    type: "Journal",
    year: "2025",
    links: {
      code: "https://github.com/BaoyaoGroup/JBHI_code",
      paper: "https://ieeexplore.ieee.org/abstract/document/11205277"
    }
  },
  {
    id: 3,
    image: "/publications/027.png",
    title: "FedCD: A Hybrid Federated Learning Framework for Adaptive Training under Data Heterogeneity",
    authors: "Weide Zhan, Baoyao Yang*",
    venue: "Chinese Conference on Pattern Recognition and Computer Vision (PRCV), 2025",
    type: "Conference",
    year: "2025",
    links: {}
  },
  {
    id: 4,
    image: "/publications/004.png",
    title: "Image-assisted Label Connective Completion for Vessel Segmentation with Insufficient Annotations",
    authors: "Xiaoqi Zheng, Baoyao Yang*, Xiuwen Fang, Mang Ye",
    venue: "IEEE International Conference on Acoustics, Speech, and Signal Processing (ICASSP), 2025",
    type: "Conference",
    year: "2025",
    links: {
      code: "https://github.com/BaoyaoGroup/LabelCompletion",
      paper: "https://ieeexplore.ieee.org/document/10888997",
      video: { name: 'video-player-XiaoqiZheng01' }
    }
  },
  {
    id: 5,
    image: "/publications/023.png",
    title: "Simple but Effective: Sub-Volume Contrastive Learning for Class-Imbalanced Semi-Supervised 3D Medical Image Segmentation",
    authors: "Xianrun Xu, Baoyao Yang*, Wanyun Li, Jingsong Lin, Yufei Xu",
    venue: "the 33rd ACM International Conference on Multimedia (ACM MM), 2025",
    type: "Conference",
    year: "2025",
    links: {
      paper: "https://dl.acm.org/doi/abs/10.1145/3746027.3755652",
      video: { name: 'video-player-XianrunXu01' }
    }
  }
])

const trackStyle = computed(() => ({
  transform: `translateX(-${currentIndex.value * 100}%)`
}))

// 轮播控制 - 修改为循环模式
const nextSlide = () => {
  currentIndex.value = (currentIndex.value + 1) % publications.value.length
}

const prevSlide = () => {
  currentIndex.value = currentIndex.value === 0 ? publications.value.length - 1 : currentIndex.value - 1
}

const goToSlide = (index: number) => {
  currentIndex.value = index
}

// 自动轮播
let autoPlayInterval: number

const startAutoPlay = () => {
  autoPlayInterval = setInterval(nextSlide, 5000)
}

const stopAutoPlay = () => {
  clearInterval(autoPlayInterval)
}

// 方法
const handleImageError = (event: Event) => {
  const target = event.target as HTMLImageElement
  target.src = '/publications/online.png'
}

const openLink = (url: string) => {
  window.open(url, '_blank')
}

const handleVideoClick = (video: any) => {
  if (video.name) {
    router.push({ name: video.name })
  }
}

const goToPublications = () => {
  router.push('/international-journals-conferences')
}

// 响应式检测
const checkScreenSize = () => {
  isMobile.value = window.innerWidth <= 768
}

// 触摸滑动支持
const touchStartX = ref(0)
const touchEndX = ref(0)

const handleTouchStart = (e: TouchEvent) => {
  touchStartX.value = e.changedTouches[0].screenX
}

const handleTouchEnd = (e: TouchEvent) => {
  touchEndX.value = e.changedTouches[0].screenX
  handleSwipe()
}

const handleSwipe = () => {
  const swipeThreshold = 50
  const diff = touchStartX.value - touchEndX.value

  if (Math.abs(diff) > swipeThreshold) {
    if (diff > 0) {
      nextSlide()
    } else {
      prevSlide()
    }
  }
}

// 生命周期
onMounted(() => {
  checkScreenSize()
  window.addEventListener('resize', checkScreenSize)
  startAutoPlay()
})

onUnmounted(() => {
  window.removeEventListener('resize', checkScreenSize)
  stopAutoPlay()
})
</script>

<style scoped>
.research-carousel {
  width: 100%;
  position: relative;
}

/* 桌面端轮播样式 */
.carousel-container {
  position: relative;
  overflow: hidden;
  border-radius: 20px;
  background: white;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(125, 18, 49, 0.1);
}

.carousel-track {
  display: flex;
  transition: transform 0.5s cubic-bezier(0.25, 0.46, 0.45, 0.94);
}

.carousel-slide {
  flex: 0 0 100%;
  min-width: 0;
}

.slide-content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  min-height: 400px;
  gap: 0;
}

.image-container {
  position: relative;
  background: #f8f9fa;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem;
}

.publication-image {
  max-width: 100%;
  max-height: 300px;
  object-fit: contain;
  border-radius: 8px;
}

.publication-badges {
  position: absolute;
  top: 1.5rem;
  left: 1.5rem;
  display: flex;
  gap: 0.5rem;
}

.badge {
  padding: 0.5rem 1rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
  color: white;
}

.type-badge {
  background: linear-gradient(135deg, #7d1231, #9a2c4d);
}

.year-badge {
  background: linear-gradient(135deg, #3498db, #2980b9);
}

.content-container {
  padding: 3rem;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.publication-title {
  font-size: 1.5rem;
  font-weight: 700;
  color: #2c3e50;
  line-height: 1.4;
  margin-bottom: 1rem;
}

.publication-authors {
  color: #7d1231;
  font-weight: 500;
  margin-bottom: 0.8rem;
  font-size: 1rem;
}

.publication-venue {
  color: #5d6d7e;
  font-size: 0.95rem;
  margin-bottom: 2rem;
  font-style: italic;
}

.publication-links {
  display: flex;
  gap: 1rem;
  flex-wrap: wrap;
}

.link-button {
  border-radius: 20px;
  font-weight: 500;
  padding: 0.75rem 1.5rem;
}

/* 导航按钮 - 移除禁用状态相关样式 */
.nav-button {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  width: 50px;
  height: 50px;
  border: none;
  border-radius: 50%;
  background: white;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.15);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
  z-index: 10;
  color: #7d1231;
}

.nav-button:hover {
  background: #7d1231;
  color: white;
  transform: translateY(-50%) scale(1.1);
}

.prev-button {
  left: 1rem;
}

.next-button {
  right: 1rem;
}

/* 指示器 */
.carousel-indicators {
  position: absolute;
  bottom: 1.5rem;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  gap: 0.5rem;
  z-index: 10;
}

.indicator {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  border: none;
  background: #bdc3c7;
  cursor: pointer;
  transition: all 0.3s ease;
}

.indicator.active {
  background: #7d1231;
  transform: scale(1.2);
}

/* 移动端样式 */
.mobile-carousel {
  position: relative;
}

.mobile-slides {
  display: flex;
  overflow-x: auto;
  scroll-snap-type: x mandatory;
  scroll-behavior: smooth;
  -webkit-overflow-scrolling: touch;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.mobile-slides::-webkit-scrollbar {
  display: none;
}

.mobile-slide {
  flex: 0 0 100%;
  scroll-snap-align: start;
  padding: 0 0.5rem;
}

.mobile-card {
  background: white;
  border-radius: 16px;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.08);
  overflow: hidden;
  margin: 1rem 0;
  border: 1px solid rgba(125, 18, 49, 0.1);
}

.card-image {
  position: relative;
  background: #f8f9fa;
  padding: 1.5rem;
  text-align: center;
}

.card-image img {
  max-width: 100%;
  max-height: 200px;
  object-fit: contain;
}

.mobile-badges {
  position: absolute;
  top: 1rem;
  left: 1rem;
  display: flex;
  gap: 0.5rem;
}

.mobile-badges .badge {
  padding: 0.4rem 0.8rem;
  font-size: 0.7rem;
}

.card-content {
  padding: 1.5rem;
}

.mobile-title {
  font-size: 1.1rem;
  font-weight: 700;
  color: #2c3e50;
  line-height: 1.4;
  margin-bottom: 0.8rem;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.mobile-authors {
  color: #7d1231;
  font-size: 0.9rem;
  font-weight: 500;
  margin-bottom: 0.5rem;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.mobile-venue {
  color: #5d6d7e;
  font-size: 0.8rem;
  margin-bottom: 1.2rem;
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.mobile-links {
  display: flex;
  gap: 0.8rem;
  flex-wrap: wrap;
}

.mobile-link {
  border-radius: 16px;
  font-size: 0.8rem;
  padding: 0.5rem 1rem;
}

.mobile-indicators {
  display: flex;
  justify-content: center;
  gap: 0.5rem;
  margin-top: 1rem;
}

.mobile-indicator {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  border: none;
  background: #bdc3c7;
  cursor: pointer;
  transition: all 0.3s ease;
}

.mobile-indicator.active {
  background: #7d1231;
  transform: scale(1.2);
}

/* 查看全部按钮 */
.view-all-section {
  text-align: center;
  margin-top: 3rem;
}

.view-all-button {
  background: linear-gradient(135deg, #7d1231, #9a2c4d);
  border: none;
  border-radius: 12px;
  padding: 1rem 2.5rem;
  font-weight: 600;
  font-size: 1.1rem;
  box-shadow: 0 8px 25px rgba(125, 18, 49, 0.3);
  transition: all 0.3s ease;
}

.view-all-button:hover {
  transform: translateY(-3px);
  box-shadow: 0 12px 35px rgba(125, 18, 49, 0.4);
}

/* 响应式调整 */
@media (max-width: 768px) {
  .slide-content {
    grid-template-columns: 1fr;
    min-height: auto;
  }

  .image-container {
    padding: 1.5rem;
    min-height: 200px;
  }

  .content-container {
    padding: 2rem 1.5rem;
  }

  .publication-title {
    font-size: 1.3rem;
  }

  .nav-button {
    width: 40px;
    height: 40px;
  }
}

@media (max-width: 480px) {
  .content-container {
    padding: 1.5rem 1rem;
  }

  .publication-title {
    font-size: 1.1rem;
  }

  .publication-links {
    justify-content: center;
  }

  .link-button {
    padding: 0.6rem 1.2rem;
    font-size: 0.8rem;
  }

  .view-all-button {
    width: 100%;
    padding: 0.9rem 2rem;
    font-size: 1rem;
  }
}
</style>