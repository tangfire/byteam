<template>
  <div class="home-container">
    <!-- 背景图片层 -->
    <div class="background-image-layer">
      <div class="image-overlay"></div>
    </div>

    <!-- 科技感背景元素 -->
    <div class="tech-background">
      <div class="floating-particles">
        <div v-for="n in 20" :key="n" class="particle" :style="particleStyle(n)"></div>
      </div>
      <div class="gradient-mesh"></div>
    </div>

    <!-- 主要内容 -->
    <div class="main-content">
      <!-- 实验室介绍区域 -->
      <section class="hero-section">
        <div class="hero-content">
          <div class="badge-icon">
            <img src="/logo/001.png" alt="BYML Logo" class="logo-image">
          </div>

          <h1 class="lab-title">
            <span class="title-main">
              <span class="word word-1">Beyond</span>
              <span class="word word-2">Machine</span>
              <span class="word word-3">Learning</span>
              <span class="word word-4">Group</span>
            </span>
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
              <span class="stat-number">15+</span>
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
                class="cta-button primary"
            >
              Explore Our Research
              <el-icon><ArrowRight /></el-icon>
            </el-button>
            <el-button
                size="large"
                @click="$router.push('/about')"
                class="cta-button secondary"
            >
              Learn More About Us
            </el-button>
          </div>
        </div>

        <div class="hero-visual">
          <div class="visual-container">
            <div class="floating-cards">
              <div class="card card-1">
                <div class="card-icon">🧠</div>
                <div class="card-content">
                  <span class="card-title">AI Research</span>
                  <span class="card-desc">Advanced ML Algorithms</span>
                </div>
              </div>
              <div class="card card-2">
                <div class="card-icon">👁️</div>
                <div class="card-content">
                  <span class="card-title">Computer Vision</span>
                  <span class="card-desc">Image & Video Analysis</span>
                </div>
              </div>
              <div class="card card-3">
                <div class="card-icon">🏥</div>
                <div class="card-content">
                  <span class="card-title">Medical AI</span>
                  <span class="card-desc">Healthcare Solutions</span>
                </div>
              </div>
              <div class="card card-4">
                <div class="card-icon">🔄</div>
                <div class="card-content">
                  <span class="card-title">Federated Learning</span>
                  <span class="card-desc">Privacy-Preserving AI</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- 研究亮点 -->
      <section class="highlights-section">
        <div class="container">
          <div class="highlight-item">
            <div class="highlight-icon">🚀</div>
            <div class="highlight-content">
              <h3>Innovation Driven</h3>
              <p>Pushing the boundaries of AI research with novel methodologies and approaches</p>
            </div>
          </div>
          <div class="highlight-item">
            <div class="highlight-icon">🤝</div>
            <div class="highlight-content">
              <h3>Collaborative Spirit</h3>
              <p>Working with leading institutions and industry partners worldwide</p>
            </div>
          </div>
          <div class="highlight-item">
            <div class="highlight-icon">🎯</div>
            <div class="highlight-content">
              <h3>Real-World Impact</h3>
              <p>Translating research into practical solutions for societal challenges</p>
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

      <!-- 最新新闻 -->
      <section class="news-section">
        <div class="section-header">
          <h2 class="section-title">Latest News</h2>
          <p class="section-subtitle">Stay updated with our recent activities and achievements</p>
        </div>

        <div class="news-timeline">
          <div
              v-for="(news, index) in latestNews"
              :key="index"
              class="news-item"
              @click="$router.push('/news')"
          >
            <div class="news-date">
              <span class="date-day">{{ getDay(news.date) }}</span>
              <span class="date-month">{{ getMonth(news.date) }}</span>
            </div>
            <div class="news-content">
              <div class="news-header">
                <span class="news-type" :class="getNewsTypeClass(news.type)">{{ news.typeLabel }}</span>
                <span class="news-year">{{ getYear(news.date) }}</span>
              </div>
              <h3 class="news-title">{{ news.title }}</h3>
              <p class="news-excerpt">{{ news.excerpt }}</p>
            </div>
            <div class="news-arrow">
              <el-icon><ArrowRight /></el-icon>
            </div>
          </div>
        </div>

        <div class="news-actions">
          <el-button
              type="primary"
              plain
              @click="$router.push('/news')"
              class="view-all-button"
          >
            View All News
            <el-icon><View /></el-icon>
          </el-button>
        </div>
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
import { ArrowRight, View } from '@element-plus/icons-vue'
import { ref } from 'vue'

// 最新新闻数据 - 增加到4条
const latestNews = ref([
  {
    type: 'publication',
    typeLabel: 'Publication',
    date: '2025-11-06',
    title: 'Multi-modal Brain Tumor Segmentation Paper Accepted',
    excerpt: 'Our paper on brain tumor segmentation has been accepted by IEEE Transactions on Multimedia.'
  },
  {
    type: 'publication',
    typeLabel: 'Publication',
    date: '2025-10-22',
    title: 'Federated Learning Framework Accepted by PRCV',
    excerpt: 'FedCD framework for adaptive training under data heterogeneity accepted by PRCV conference.'
  },
  {
    type: 'team',
    typeLabel: 'Team Update',
    date: '2025-05-06',
    title: 'Welcome New Group Members',
    excerpt: 'Warm welcome to 7 new group members joining our research team this semester.'
  },
  {
    type: 'event',
    typeLabel: 'Event',
    date: '2025-04-08',
    title: 'Official Website Launch',
    excerpt: 'Beyond Machine Learning team introduction official website officially launched!'
  }
])

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

const getNewsTypeClass = (type: string) => {
  const classMap: { [key: string]: string } = {
    publication: 'type-publication',
    team: 'type-team',
    award: 'type-award',
    event: 'type-event'
  }
  return classMap[type] || 'type-general'
}

// 日期处理函数
const getDay = (dateString: string) => {
  return new Date(dateString).getDate()
}

const getMonth = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleString('en-US', { month: 'short' })
}

const getYear = (dateString: string) => {
  return new Date(dateString).getFullYear()
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
  --primary-dark: #5a0c22;
  --secondary-color: #13393e;
  --accent-color: #3498db;
  --accent-light: #5dade2;
  --text-primary: #2c3e50;
  --text-secondary: #5d6d7e;
  --text-light: #7f8c8d;
  --bg-light: #f8f9fa;
  --bg-white: #ffffff;
  --bg-card: rgba(255, 255, 255, 0.95);
  --gradient-primary: linear-gradient(135deg, #7d1231 0%, #9a2c4d 100%);
  --gradient-secondary: linear-gradient(135deg, #13393e 0%, #3498db 100%);
  --gradient-hero: linear-gradient(135deg, #7d1231 0%, #13393e 100%);
  --shadow-soft: 0 8px 30px rgba(0, 0, 0, 0.08);
  --shadow-medium: 0 15px 40px rgba(0, 0, 0, 0.12);
  --shadow-large: 0 25px 50px rgba(0, 0, 0, 0.15);
  --border-radius: 16px;
  --border-radius-lg: 20px;

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
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.85) 0%, rgba(248, 249, 250, 0.9) 50%, rgba(255, 255, 255, 0.85) 100%);
}

/* 科技背景 */
.tech-background {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 1;
  overflow: hidden;
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
  opacity: 0.1;
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
      radial-gradient(circle at 20% 80%, rgba(125, 18, 49, 0.08) 0%, transparent 50%),
      radial-gradient(circle at 80% 20%, rgba(19, 57, 62, 0.08) 0%, transparent 50%),
      radial-gradient(circle at 40% 40%, rgba(52, 152, 219, 0.05) 0%, transparent 50%);
  animation: pulseMesh 15s ease-in-out infinite;
}

@keyframes pulseMesh {
  0%, 100% {
    opacity: 1;
    transform: scale(1);
  }
  50% {
    opacity: 0.7;
    transform: scale(1.1);
  }
}

/* 主内容 */
.main-content {
  position: relative;
  z-index: 2;
  max-width: 1400px;
  margin: 0 auto;
  padding: 0 2rem;
}

/* 英雄区域 */
.hero-section {
  display: grid;
  grid-template-columns: 1.2fr 1fr;
  gap: 5rem;
  align-items: center;
  min-height: 100vh;
  padding: 2rem 0;
  position: relative;
  z-index: 3;
}

.hero-content {
  padding-right: 2rem;
  position: relative;
  z-index: 10;
  animation: fadeInUp 1s ease-out;
  isolation: isolate;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.lab-badge {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 2.5rem;
}

.badge-icon {
  width: 90px;
  height: 90px;
  background: linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%);
  border-radius: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
  box-shadow: 0 12px 36px rgba(125, 18, 49, 0.18), inset 0 1px 0 rgba(255, 255, 255, 0.9);
  border: 2px solid rgba(125, 18, 49, 0.12);
  animation: floatLogo 6s ease-in-out infinite;
}

@keyframes floatLogo {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-10px);
  }
}

.badge-icon .logo-image {
  width: 100px;
  height: 100px;
  object-fit: contain;
}

.badge-icon:hover {
  transform: translateY(-10px) scale(1.06) rotate(5deg);
  box-shadow: 0 20px 50px rgba(125, 18, 49, 0.3), inset 0 1px 0 rgba(255, 255, 255, 0.95);
  border-color: var(--primary-color);
  animation-play-state: paused;
}

.badge-icon svg {
  width: 35px;
  height: 35px;
}

.badge-text {
  font-size: 1.1rem;
  font-weight: 600;
  color: var(--primary-color);
  background: rgba(125, 18, 49, 0.1);
  padding: 0.5rem 1rem;
  border-radius: 20px;
}

.lab-title {
  margin-bottom: 2rem;
}

.title-main {
  display: inline-block;
  font-size: 4.5rem;
  font-weight: 900;
  color: var(--primary-color);
  letter-spacing: -1.5px;
  white-space: nowrap;
  text-shadow: 0 4px 12px rgba(125, 18, 49, 0.25);
}

.word {
  display: inline-block;
  margin-right: 1.2rem;
  animation: bounceWord 2.5s ease-in-out infinite;
}

.word:last-child {
  margin-right: 0;
}

.word-1 {
  animation-delay: 0s;
}

.word-2 {
  animation-delay: 0.1s;
}

.word-3 {
  animation-delay: 0.2s;
}

.word-4 {
  animation-delay: 0.3s;
}

@keyframes bounceWord {
  0%, 100% {
    transform: translateY(0);
  }
  20% {
    transform: translateY(-10px);
  }
  40% {
    transform: translateY(0);
  }
}

@keyframes slideInLeft {
  from {
    opacity: 0;
    transform: translateX(-50px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.lab-mission {
  font-size: 1.5rem;
  line-height: 1.9;
  color: var(--text-secondary);
  margin-bottom: 3.5rem;
  font-weight: 400;
  max-width: 95%;
  position: relative;
  padding-left: 24px;
  border-left: 5px solid var(--primary-color);
  animation: fadeInLeft 1s ease-out 0.4s both;
  background: linear-gradient(90deg, rgba(125, 18, 49, 0.03) 0%, transparent 100%);
  padding: 1.5rem 1.5rem 1.5rem 24px;
  border-radius: 0 12px 12px 0;
}

@keyframes fadeInLeft {
  from {
    opacity: 0;
    transform: translateX(-30px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.mission-stats {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 2rem;
  margin-bottom: 3rem;
  animation: fadeInUp 1s ease-out 0.6s both;
}

.stat-item {
  text-align: center;
  padding: 2.2rem 1.8rem;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.95) 0%, rgba(248, 249, 250, 0.9) 100%);
  backdrop-filter: blur(20px);
  border-radius: var(--border-radius-lg);
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.1), inset 0 1px 0 rgba(255, 255, 255, 0.9);
  border: 2px solid rgba(125, 18, 49, 0.12);
  transition: all 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
  position: relative;
  overflow: hidden;
}

.stat-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(125, 18, 49, 0.05) 0%, transparent 50%);
  opacity: 0;
  transition: opacity 0.4s ease;
}

.stat-item:hover::before {
  opacity: 1;
}

.stat-item:hover {
  transform: translateY(-12px) scale(1.08);
  box-shadow: 0 24px 70px rgba(125, 18, 49, 0.25);
  border-color: var(--primary-color);
}

.stat-number {
  display: block;
  font-size: 3.5rem;
  font-weight: 900;
  background: linear-gradient(135deg, var(--primary-color) 0%, #a51c41 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  line-height: 1;
  margin-bottom: 0.6rem;
  position: relative;
  z-index: 1;
}

.stat-label {
  font-size: 1rem;
  color: var(--text-secondary);
  font-weight: 500;
}

.hero-actions {
  display: flex;
  gap: 1.5rem;
  margin-top: 2rem;
}

.cta-button {
  border-radius: 18px;
  padding: 1.3rem 2.6rem;
  font-weight: 700;
  font-size: 1.15rem;
  transition: all 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
  border: none;
  position: relative;
  overflow: hidden;
  cursor: pointer;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
}

.cta-button::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.35), transparent);
  transition: left 0.7s ease;
}

.cta-button:hover::before {
  left: 100%;
}

.cta-button.primary {
  background: linear-gradient(135deg, var(--primary-color) 0%, #a51c41 100%);
  color: white;
  box-shadow: 0 12px 32px rgba(125, 18, 49, 0.35), inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 4px 12px rgba(0, 0, 0, 0.1);
}

.cta-button.primary:hover {
  transform: translateY(-8px);
  box-shadow: 0 24px 60px rgba(125, 18, 49, 0.5), inset 0 1px 0 rgba(255, 255, 255, 0.2);
  filter: brightness(1.1);
}

.cta-button.secondary {
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 249, 250, 0.95) 100%);
  color: var(--primary-color);
  border: 2px solid var(--primary-color);
  box-shadow: 0 6px 20px rgba(125, 18, 49, 0.15), inset 0 1px 0 rgba(255, 255, 255, 0.9), 0 4px 12px rgba(0, 0, 0, 0.08);
}

.cta-button.secondary:hover {
  background: linear-gradient(135deg, var(--primary-color) 0%, #a51c41 100%);
  color: white;
  transform: translateY(-8px);
  box-shadow: 0 20px 50px rgba(125, 18, 49, 0.4), inset 0 1px 0 rgba(255, 255, 255, 0.2);
  border-color: transparent;
}

/* 英雄区域视觉元素 */
.hero-visual {
  position: relative;
  height: 500px;
  z-index: 1;
}

.visual-container {
  position: relative;
  width: 100%;
  height: 100%;
}

.floating-cards {
  position: relative;
  width: 100%;
  height: 100%;
}

.card {
  position: absolute;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 249, 250, 0.95) 100%);
  padding: 1.6rem;
  border-radius: 24px;
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.12), inset 0 1px 0 rgba(255, 255, 255, 0.95), 0 4px 12px rgba(0, 0, 0, 0.08);
  display: flex;
  align-items: center;
  gap: 1.4rem;
  font-weight: 600;
  color: var(--text-primary);
  animation: floatCard 8s ease-in-out infinite;
  border: 2px solid rgba(125, 18, 49, 0.18);
  backdrop-filter: blur(20px);
  max-width: 260px;
  transition: all 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
  position: relative;
  overflow: hidden;
}

.card::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(125, 18, 49, 0.12) 0%, transparent 70%);
  opacity: 0;
  transition: opacity 0.4s ease;
}

.card:hover::before {
  opacity: 1;
}

.card:hover {
  transform: translateY(-20px) scale(1.12) rotate(2deg);
  box-shadow: 0 28px 80px rgba(125, 18, 49, 0.35), inset 0 1px 0 rgba(255, 255, 255, 0.98);
  border-color: var(--primary-color);
}

.card-1 {
  top: 5%;
  left: 0%;
  animation-delay: 0s;
  z-index: 4;
}

.card-2 {
  top: 0%;
  right: 5%;
  animation-delay: -2s;
  z-index: 3;
}

.card-3 {
  bottom: 5%;
  left: 5%;
  animation-delay: -4s;
  z-index: 2;
}

.card-4 {
  bottom: 0%;
  right: 0%;
  animation-delay: -6s;
  z-index: 1;
}

.card-icon {
  font-size: 2rem;
  flex-shrink: 0;
}

.card-content {
  display: flex;
  flex-direction: column;
}

.card-title {
  font-weight: 600;
  font-size: 1rem;
  margin-bottom: 0.2rem;
}

.card-desc {
  font-size: 0.8rem;
  color: var(--text-light);
}

@keyframes floatCard {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
  }
  33% {
    transform: translateY(-20px) rotate(2deg);
  }
  66% {
    transform: translateY(-10px) rotate(-2deg);
  }
}

/* 亮点区域 */
.highlights-section {
  padding: 4rem 0;
  border-radius: var(--border-radius-lg);
  margin: 2rem 0;
}

.highlights-section .container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  max-width: 1200px;
  margin: 0 auto;
}

.highlight-item {
  display: flex;
  align-items: flex-start;
  gap: 1.8rem;
  padding: 2.5rem;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 249, 250, 0.95) 100%);
  backdrop-filter: blur(20px);
  border-radius: 24px;
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.1), inset 0 1px 0 rgba(255, 255, 255, 0.9), 0 4px 12px rgba(0, 0, 0, 0.06);
  border: 2px solid rgba(125, 18, 49, 0.12);
  transition: all 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
  position: relative;
  overflow: hidden;
}

.highlight-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(125, 18, 49, 0.08) 0%, rgba(19, 57, 62, 0.05) 100%);
  opacity: 0;
  transition: opacity 0.4s ease;
}

.highlight-item:hover::before {
  opacity: 1;
}

.highlight-item:hover {
  transform: translateY(-12px) scale(1.03);
  box-shadow: 0 24px 70px rgba(125, 18, 49, 0.25), inset 0 1px 0 rgba(255, 255, 255, 0.98);
  border-color: var(--primary-color);
}

.highlight-icon {
  font-size: 3.5rem;
  flex-shrink: 0;
  filter: drop-shadow(0 6px 12px rgba(125, 18, 49, 0.25));
  transition: all 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
}

.highlight-item:hover .highlight-icon {
  transform: scale(1.2) rotate(8deg);
  filter: drop-shadow(0 10px 20px rgba(125, 18, 49, 0.35));
}

.highlight-content h3 {
  font-size: 1.4rem;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 0.8rem;
}

.highlight-content p {
  color: var(--text-secondary);
  line-height: 1.6;
}

/* 通用区域样式 */
.research-section,
.quick-nav-section,
.news-section {
  padding: 6rem 0;
}

.section-header {
  text-align: center;
  margin-bottom: 4rem;
}

.section-title {
  font-size: 3rem;
  font-weight: 800;
  background: linear-gradient(135deg, var(--primary-color) 0%, #a51c41 50%, #7d1231 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 1.5rem;
  position: relative;
  display: inline-block;
  letter-spacing: -0.5px;
}

.section-title::after {
  content: '';
  position: absolute;
  bottom: -12px;
  left: 50%;
  transform: translateX(-50%);
  width: 100px;
  height: 5px;
  background: linear-gradient(90deg, transparent, var(--primary-color), transparent);
  border-radius: 3px;
  box-shadow: 0 2px 8px rgba(125, 18, 49, 0.3);
}

.section-subtitle {
  font-size: 1.4rem;
  color: var(--text-secondary);
  max-width: 650px;
  margin: 0 auto;
  line-height: 1.8;
  font-weight: 400;
  opacity: 0.9;
}

/* 新闻模块样式 - 时间线布局 */
.news-section {
  margin: 2rem 0;
  border-radius: var(--border-radius-lg);
}

.news-timeline {
  max-width: 900px;
  margin: 0 auto 3rem;
}

.news-item {
  display: flex;
  align-items: center;
  gap: 2rem;
  padding: 2.5rem;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 249, 250, 0.95) 100%);
  backdrop-filter: blur(20px);
  border-radius: 24px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.08), inset 0 1px 0 rgba(255, 255, 255, 0.9), 0 4px 12px rgba(0, 0, 0, 0.05);
  margin-bottom: 1.5rem;
  cursor: pointer;
  transition: all 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
  border: 2px solid rgba(125, 18, 49, 0.1);
  position: relative;
  overflow: hidden;
}

.news-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(125, 18, 49, 0.05) 0%, rgba(19, 57, 62, 0.03) 100%);
  opacity: 0;
  transition: opacity 0.4s ease;
}

.news-item:hover::before {
  opacity: 1;
}

.news-item:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 60px rgba(125, 18, 49, 0.25), inset 0 1px 0 rgba(255, 255, 255, 0.98);
  border-color: var(--primary-color);
}

.news-date {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-width: 90px;
  padding: 1.2rem;
  background: linear-gradient(135deg, rgba(125, 18, 49, 0.08) 0%, rgba(19, 57, 62, 0.05) 100%);
  border-radius: 16px;
  border: 2px solid rgba(125, 18, 49, 0.15);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.08), inset 0 1px 0 rgba(255, 255, 255, 0.5);
  position: relative;
  overflow: hidden;
}

.news-date::after {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(125, 18, 49, 0.05) 0%, transparent 70%);
  opacity: 0;
  transition: opacity 0.4s ease;
}

.news-item:hover .news-date::after {
  opacity: 1;
}

.date-day {
  font-size: 2.2rem;
  font-weight: 800;
  background: linear-gradient(135deg, var(--primary-color) 0%, #a51c41 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  line-height: 1;
  position: relative;
  z-index: 1;
}

.date-month {
  font-size: 0.95rem;
  font-weight: 700;
  color: var(--text-secondary);
  text-transform: uppercase;
  margin-top: 0.3rem;
  letter-spacing: 0.5px;
  position: relative;
  z-index: 1;
}

.news-content {
  flex: 1;
}

.news-header {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 0.8rem;
}

.news-type {
  padding: 6px 16px;
  border-radius: 20px;
  font-size: 0.85rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.8px;
  border: 2px solid transparent;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.type-publication {
  background: linear-gradient(135deg, rgba(125, 18, 49, 0.1) 0%, rgba(165, 28, 65, 0.08) 100%);
  color: var(--primary-color);
  border-color: rgba(125, 18, 49, 0.2);
}

.type-team {
  background: linear-gradient(135deg, rgba(52, 152, 219, 0.1) 0%, rgba(41, 128, 185, 0.08) 100%);
  color: var(--accent-color);
  border-color: rgba(52, 152, 219, 0.2);
}

.type-award {
  background: linear-gradient(135deg, rgba(241, 196, 15, 0.1) 0%, rgba(243, 156, 18, 0.08) 100%);
  color: #f39c12;
  border-color: rgba(241, 196, 15, 0.2);
}

.type-event {
  background: linear-gradient(135deg, rgba(46, 204, 113, 0.1) 0%, rgba(39, 174, 96, 0.08) 100%);
  color: #27ae60;
  border-color: rgba(46, 204, 113, 0.2);
}

.news-item:hover .news-type {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.news-year {
  font-size: 0.9rem;
  color: var(--text-light);
  font-weight: 500;
}

.news-title {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--text-primary);
  margin-bottom: 0.8rem;
  line-height: 1.4;
  letter-spacing: -0.3px;
  transition: color 0.3s ease;
}

.news-item:hover .news-title {
  color: var(--primary-color);
}

.news-excerpt {
  color: var(--text-secondary);
  line-height: 1.7;
  margin: 0;
  font-size: 1rem;
  opacity: 0.95;
}

.news-arrow {
  color: var(--text-light);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  opacity: 0.6;
}

.news-item:hover .news-arrow {
  color: var(--primary-color);
  transform: translateX(8px);
  opacity: 1;
}

.news-actions {
  text-align: center;
}

.view-all-button {
  border-radius: 16px;
  padding: 1rem 2.4rem;
  font-weight: 600;
  font-size: 1.05rem;
  border: 2px solid var(--primary-color);
  color: var(--primary-color);
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 249, 250, 0.95) 100%);
  transition: all 0.4s cubic-bezier(0.34, 1.56, 0.64, 1);
  box-shadow: 0 4px 12px rgba(125, 18, 49, 0.15), inset 0 1px 0 rgba(255, 255, 255, 0.9);
}

.view-all-button:hover {
  background: linear-gradient(135deg, var(--primary-color) 0%, #a51c41 100%);
  color: white;
  transform: translateY(-4px);
  box-shadow: 0 12px 32px rgba(125, 18, 49, 0.35), inset 0 1px 0 rgba(255, 255, 255, 0.2);
  border-color: transparent;
}

/* 快速导航网格 */
.nav-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 2rem;
  max-width: 1000px;
  margin: 0 auto;
}

.nav-card {
  border-radius: var(--border-radius);
  cursor: pointer;
  transition: all 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
  border: 2px solid rgba(125, 18, 49, 0.12);
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.98) 0%, rgba(248, 249, 250, 0.95) 100%);
  overflow: hidden;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08), inset 0 1px 0 rgba(255, 255, 255, 0.9);
}

.nav-card:hover {
  transform: translateY(-12px) scale(1.03);
  box-shadow: 0 20px 50px rgba(125, 18, 49, 0.2);
  border-color: var(--primary-color);
  border-width: 2px;
}

.nav-card-content {
  padding: 2.5rem 2rem;
  text-align: center;
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
}

.nav-icon {
  font-size: 3.5rem;
  margin-bottom: 1.5rem;
  transition: all 0.4s cubic-bezier(0.34, 1.56, 0.64, 1);
  filter: drop-shadow(0 4px 8px rgba(125, 18, 49, 0.15));
}

.nav-card:hover .nav-icon {
  transform: scale(1.2) rotate(10deg);
  filter: drop-shadow(0 8px 16px rgba(125, 18, 49, 0.3));
}

.nav-card h3 {
  font-size: 1.5rem;
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
@media (max-width: 1200px) {
  .hero-section {
    gap: 3rem;
  }

  .title-main {
    font-size: 3rem;
  }
}

@media (max-width: 1024px) {
  .hero-section {
    grid-template-columns: 1fr;
    gap: 3rem;
    text-align: center;
    padding: 2rem 0;
  }

  .hero-content {
    padding-right: 0;
  }

  .lab-badge {
    justify-content: center;
  }

  .mission-stats {
    grid-template-columns: repeat(3, 1fr);
    max-width: 600px;
    margin-left: auto;
    margin-right: auto;
  }

  .title-main {
    font-size: 2.8rem;
  }

  .lab-mission {
    max-width: 100%;
    margin-left: auto;
    margin-right: auto;
  }

  .nav-grid {
    grid-template-columns: repeat(2, 1fr);
  }

  /* 隐藏 logo */
  .badge-icon {
    display: none !important;
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
    font-size: 2.2rem;
    white-space: normal;
    line-height: 1.2;
  }

  .lab-mission {
    font-size: 1.1rem;
  }

  .mission-stats {
    grid-template-columns: 1fr;
    gap: 1.5rem;
  }

  .stat-number {
    font-size: 2rem;
  }

  .hero-actions {
    flex-direction: column;
    gap: 1rem;
  }

  .research-section,
  .quick-nav-section,
  .news-section {
    padding: 4rem 0;
  }

  .section-title {
    font-size: 2.2rem;
  }

  .section-subtitle {
    font-size: 1.1rem;
  }

  .hero-visual {
    height: 400px;
  }

  .card {
    padding: 1rem;
    max-width: 180px;
  }

  .background-image-layer {
    background-attachment: scroll;
  }

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }

  .news-item {
    flex-direction: column;
    text-align: center;
    gap: 1.5rem;
    padding: 1.5rem;
  }

  .news-date {
    flex-direction: row;
    gap: 1rem;
    min-width: auto;
  }

  .date-day {
    font-size: 1.5rem;
  }

  .news-header {
    flex-direction: column;
    gap: 0.5rem;
  }

  .news-title {
    font-size: 1.2rem;
  }

  .highlights-section .container {
    grid-template-columns: 1fr;
  }

  .highlight-item {
    flex-direction: column;
    text-align: center;
    gap: 1rem;
  }

  .nav-grid {
    grid-template-columns: 1fr;
  }

  /* 隐藏 logo */
  .badge-icon {
    display: none !important;
  }
}

@media (max-width: 480px) {
  .title-main {
    font-size: 1.8rem;
  }

  .mission-stats {
    grid-template-columns: 1fr;
  }

  .section-title {
    font-size: 1.8rem;
  }

  .nav-card-content {
    padding: 1.5rem;
  }

  .lab-badge {
    flex-direction: column;
    gap: 0.8rem;
  }

  .news-item {
    padding: 1.2rem;
  }

  /* 隐藏 logo */
  .badge-icon {
    display: none !important;
  }
}

/* 大屏幕优化 */
@media (min-width: 1400px) {
  .main-content {
    max-width: 1400px;
  }
}

/* 保证 el-icon 的图标颜色在初始时是正确的 */
::v-deep(.el-icon svg) {
  color: #7d1231 !important;
}
</style>