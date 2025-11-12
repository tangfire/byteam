
<template>
  <div class="home-container">
    <!-- 背景图片层 -->
    <div class="background-image-layer">
      <div class="image-overlay"></div>
    </div>

    <!-- 科技感背景元素 -->
    <div class="tech-background">
      <div class="floating-particles">
        <div v-for="n in 15" :key="n" class="particle" :style="particleStyle(n)"></div>
      </div>
      <div class="gradient-mesh"></div>
    </div>

    <!-- 主要内容 -->
    <div class="main-content">
      <!-- 实验室介绍区域 -->
      <section class="hero-section">
        <div class="hero-content">
          <div class="lab-badge">
            <div class="badge-icon">
              <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M19 14C19 16.973 16.837 19.441 14 19.917V21H10V19.917C7.163 19.441 5 16.973 5 14V7C5 6.447 5.447 6 6 6H18C18.553 6 19 6.447 19 7V14Z" stroke="currentColor" stroke-width="2"/>
                <path d="M12 10V14" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
                <path d="M9 12H15" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              </svg>
            </div>
          </div>

          <h1 class="lab-title">
            <span class="title-main">Beyond Machine Learning Group</span>
          </h1>

          <p class="lab-mission">
            Pioneering the next generation of intelligent systems through interdisciplinary research
            in machine learning, computer vision, and biomedical computing.
          </p>

          <div class="mission-stats">
            <div class="stat-item">
              <span class="stat-number">30+</span>
              <span class="stat-label">Publications</span>
            </div>
            <div class="stat-item">
              <span class="stat-number">30+</span>
              <span class="stat-label">Research Projects</span>
            </div>
            <div class="stat-item">
              <span class="stat-number">20+</span>
              <span class="stat-label">Team Members</span>
            </div>
          </div>

          <div class="hero-actions">
            <el-button
                type="primary"
                size="large"
                @click="scrollToResearch"
                class="cta-button"
            >
              Explore Our Research
              <el-icon><ArrowRight /></el-icon>
            </el-button>
          </div>
        </div>

        <div class="hero-visual">
          <div class="floating-cards">
            <div class="card card-1">
              <div class="card-icon">🧠</div>
              <span>AI Research</span>
            </div>
            <div class="card card-2">
              <div class="card-icon">👁️</div>
              <span>Computer Vision</span>
            </div>
            <div class="card card-3">
              <div class="card-icon">🏥</div>
              <span>Medical AI</span>
            </div>
            <div class="card card-4">
              <div class="card-icon">🔄</div>
              <span>Federated Learning</span>
            </div>
          </div>
        </div>
      </section>

      <!-- 最新研究成果 -->
      <section class="research-section" id="research-highlights">
        <div class="section-header">
          <h2 class="section-title">Latest Research</h2>
          <p class="section-subtitle">Cutting-edge publications and breakthroughs from our team</p>
        </div>

        <research-carousel class="research-carousel-wrapper" />
      </section>

      <!-- 快速导航 -->
      <section class="quick-nav-section">
        <div class="section-header">
          <h2 class="section-title">Explore More</h2>
          <p class="section-subtitle">Discover our research areas and team</p>
        </div>

        <div class="nav-grid">
          <el-card
              class="nav-card"
              shadow="hover"
              @click="$router.push('/research-direction')"
          >
            <div class="nav-card-content">
              <div class="nav-icon">🎯</div>
              <h3>Research Directions</h3>
              <p>Explore our innovative research areas and methodologies</p>
              <span class="nav-arrow">→</span>
            </div>
          </el-card>

          <el-card
              class="nav-card"
              shadow="hover"
              @click="$router.push('/our-group')"
          >
            <div class="nav-card-content">
              <div class="nav-icon">👥</div>
              <h3>Our Team</h3>
              <p>Meet our talented researchers and collaborators</p>
              <span class="nav-arrow">→</span>
            </div>
          </el-card>

          <el-card
              class="nav-card"
              shadow="hover"
              @click="$router.push('/international-journals-conferences')"
          >
            <div class="nav-card-content">
              <div class="nav-icon">📚</div>
              <h3>Publications</h3>
              <p>Browse our latest papers and conference proceedings</p>
              <span class="nav-arrow">→</span>
            </div>
          </el-card>
        </div>
      </section>
    </div>

    <el-backtop class="mobile-backtop" :right="100" :bottom="100" />
  </div>
</template>

<script setup lang="ts">
import ResearchCarousel from '../components/ResearchCarousel.vue'
import { ArrowRight } from '@element-plus/icons-vue'
import { ref } from 'vue'

const particleStyle = (index: number) => {
  const size = Math.random() * 4 + 2
  const delay = Math.random() * 5
  const duration = Math.random() * 8 + 8
  return {
    width: `${size}px`,
    height: `${size}px`,
    animationDelay: `${delay}s`,
    animationDuration: `${duration}s`,
    left: `${Math.random() * 100}%`,
    top: `${Math.random() * 100}%`
  }
}

const scrollToResearch = () => {
  const element = document.getElementById('research-highlights')
  if (element) {
    element.scrollIntoView({ behavior: 'smooth' })
  }
}
</script>

<style scoped>
.home-container {
  --primary-color: #7d1231;
  --primary-light: #9a2c4d;
  --secondary-color: #13393e;
  --accent-color: #3498db;
  --text-primary: #2c3e50;
  --text-secondary: #5d6d7e;
  --bg-light: #f8f9fa;
  --bg-white: #ffffff;
  --gradient-primary: linear-gradient(135deg, #7d1231 0%, #9a2c4d 100%);
  --gradient-secondary: linear-gradient(135deg, #13393e 0%, #3498db 100%);
  --shadow-soft: 0 8px 30px rgba(0, 0, 0, 0.08);
  --shadow-medium: 0 15px 40px rgba(0, 0, 0, 0.12);

  min-height: 100vh;
  position: relative;
  background: var(--bg-light);
  overflow-x: hidden;
}

/* 背景图片层 */
.background-image-layer {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-image: url('/background/ContactBackground.jpg');
  background-size: cover;
  background-position: center;
  background-attachment: fixed;
  z-index: 0;
}

.image-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(255, 255, 255, 0.55);
}

/* 科技背景 */
.tech-background {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 1;
}

.floating-particles {
  position: absolute;
  width: 100%;
  height: 100%;
}

.particle {
  position: absolute;
  background: var(--primary-color);
  border-radius: 50%;
  opacity: 0.15;
  animation: floatParticle linear infinite;
}

@keyframes floatParticle {
  0% {
    transform: translateY(0px) rotate(0deg);
    opacity: 0.1;
  }
  50% {
    opacity: 0.05;
  }
  100% {
    transform: translateY(-100vh) rotate(360deg);
    opacity: 0.1;
  }
}

.gradient-mesh {
  position: absolute;
  width: 100%;
  height: 100%;
  background:
      radial-gradient(circle at 20% 80%, rgba(125, 18, 49, 0.03) 0%, transparent 50%),
      radial-gradient(circle at 80% 20%, rgba(19, 57, 62, 0.03) 0%, transparent 50%),
      radial-gradient(circle at 40% 40%, rgba(52, 152, 219, 0.02) 0%, transparent 50%);
}

/* 主内容 */
.main-content {
  position: relative;
  z-index: 2;
  max-width: 1400px;
  margin: 0 auto;
  padding: 2rem;
}

/* 英雄区域 */
.hero-section {
  display: grid;
  grid-template-columns: 1.2fr 1fr;
  gap: 4rem;
  align-items: center;
  min-height: 100vh;
  padding: 2rem 0;
}

.hero-content {
  padding-right: 2rem;
}

.lab-badge {
  margin-bottom: 2rem;
}

.badge-icon {
  width: 80px;
  height: 80px;
  background: var(--gradient-primary);
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 2rem;
  box-shadow: var(--shadow-soft);
}

.badge-icon svg {
  width: 40px;
  height: 40px;
}

.lab-title {
  margin-bottom: 1.5rem;
}

.title-main {
  display: block;
  font-size: 3.2rem;
  font-weight: 700;
  background: var(--gradient-primary);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  line-height: 1.5;
  margin-bottom: 0.5rem;
  white-space: nowrap; /* 防止换行 */
}




.lab-mission {
  font-size: 1.2rem;
  line-height: 1.7;
  color: var(--text-secondary);
  margin-bottom: 3rem;
  font-weight: 400;
}

.mission-stats {
  display: flex;
  gap: 3rem;
  margin-bottom: 3rem;
}

.stat-item {
  text-align: center;
}

.stat-number {
  display: block;
  font-size: 2.2rem;
  font-weight: 700;
  color: var(--primary-color);
  line-height: 1;
}

.stat-label {
  font-size: 0.9rem;
  color: var(--text-secondary);
  font-weight: 500;
}

.hero-actions {
  margin-top: 2rem;
}

.cta-button {
  background: var(--gradient-primary);
  border: none;
  border-radius: 12px;
  padding: 1rem 2rem;
  font-weight: 600;
  font-size: 1.1rem;
  box-shadow: var(--shadow-soft);
  transition: all 0.3s ease;
}

.cta-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 25px rgba(125, 18, 49, 0.3);
}

/* 英雄区域视觉元素 */
.hero-visual {
  position: relative;
  height: 400px;
}

.floating-cards {
  position: relative;
  height: 100%;
}

.card {
  position: absolute;
  background: var(--bg-white);
  padding: 1.2rem;
  border-radius: 16px;
  box-shadow: var(--shadow-soft);
  display: flex;
  align-items: center;
  gap: 1rem;
  font-weight: 500;
  color: var(--text-primary);
  animation: floatCard 8s ease-in-out infinite;
  border: 1px solid rgba(125, 18, 49, 0.1);
  backdrop-filter: blur(10px);
}

.card-1 {
  top: 10%;
  left: 5%;
  animation-delay: 0s;
}

.card-2 {
  top: 40%;
  right: 5%;
  animation-delay: -2s;
}

.card-3 {
  bottom: 20%;
  left: 20%;
  animation-delay: -4s;
}

.card-4 {
  bottom: 10%;
  right: 15%;
  animation-delay: -6s;
}

.card-icon {
  font-size: 1.8rem;
}

@keyframes floatCard {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
  }
  50% {
    transform: translateY(-15px) rotate(2deg);
  }
}

/* 通用区域样式 */
.research-section,
.quick-nav-section {
  padding: 6rem 0;
}

.section-header {
  text-align: center;
  margin-bottom: 4rem;
}

.section-title {
  font-size: 2.5rem;
  font-weight: 700;
  color: var(--text-primary);
  margin-bottom: 1rem;
}

.section-subtitle {
  font-size: 1.2rem;
  color: var(--text-secondary);
  max-width: 600px;
  margin: 0 auto;
}

/* 快速导航网格 */
.nav-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  max-width: 1000px;
  margin: 0 auto;
}

.nav-card {
  border-radius: 16px;
  cursor: pointer;
  transition: all 0.3s ease;
  border: 1px solid rgba(125, 18, 49, 0.1);
}

.nav-card:hover {
  transform: translateY(-8px);
  box-shadow: var(--shadow-medium);
  border-color: var(--primary-color);
}

.nav-card-content {
  padding: 2rem;
  text-align: center;
  position: relative;
}

.nav-icon {
  font-size: 3rem;
  margin-bottom: 1.5rem;
}

.nav-card h3 {
  font-size: 1.4rem;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 1rem;
}

.nav-card p {
  color: var(--text-secondary);
  line-height: 1.6;
  margin-bottom: 1.5rem;
}

.nav-arrow {
  color: var(--primary-color);
  font-size: 1.5rem;
  font-weight: bold;
  transition: transform 0.3s ease;
}

.nav-card:hover .nav-arrow {
  transform: translateX(5px);
}

/* 轮播图包装器 */
.research-carousel-wrapper {
  max-width: 1000px;
  margin: 0 auto;
}

/* 响应式设计 */
@media (max-width: 1024px) {
  .hero-section {
    grid-template-columns: 1fr;
    gap: 3rem;
    text-align: center;
  }

  .hero-content {
    padding-right: 0;
  }

  .mission-stats {
    justify-content: center;
  }

  .title-main {
    font-size: 2.6rem;
  }

  .title-sub {
    font-size: 1.4rem;
  }
}

@media (max-width: 768px) {
  .main-content {
    padding: 1rem;
  }

  .hero-section {
    padding: 1rem 0;
    min-height: auto;
  }

  .title-main {
    font-size: 2rem;
  }

  .title-sub {
    font-size: 1.2rem;
  }

  .lab-mission {
    font-size: 1rem;
  }

  .mission-stats {
    gap: 2rem;
    flex-wrap: wrap;
  }

  .stat-number {
    font-size: 1.8rem;
  }

  .research-section,
  .quick-nav-section {
    padding: 4rem 0;
  }

  .section-title {
    font-size: 2rem;
  }

  .nav-grid {
    grid-template-columns: 1fr;
    gap: 1.5rem;
  }

  .hero-visual {
    height: 300px;
  }

  .card {
    padding: 1rem;
    font-size: 0.9rem;
  }

  .background-image-layer {
    background-attachment: scroll;
  }

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }

}

@media (max-width: 480px) {
  .title-main {
    font-size: 1.6rem;
  }

  .mission-stats {
    flex-direction: column;
    gap: 1.5rem;
  }

  .section-title {
    font-size: 1.6rem;
  }

  .nav-card-content {
    padding: 1.5rem;
  }
}

/* 在移动端小屏幕上允许换行 */
@media (max-width: 768px) {
  .title-main {
    white-space: normal; /* 允许换行 */
    font-size: 2.2rem;
    line-height: 1.2;
  }
}

/* 在更小的屏幕上确保可读性 */
@media (max-width: 480px) {
  .title-main {
    font-size: 1.8rem;
    white-space: normal; /* 确保在小屏幕上换行 */
  }
}

/* 保证 el-icon 的图标颜色在初始时是正确的 */
::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}
</style>