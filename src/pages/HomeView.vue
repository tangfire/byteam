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
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.7) 0%, rgba(248, 249, 250, 0.8) 100%);
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
  padding: 4rem 0;
}

.hero-content {
  padding-right: 2rem;
}

.lab-badge {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 2.5rem;
}

.badge-icon {
  /* 保持原有的样式 */
  width: 70px;
  height: 70px;
  background: white;
  border-radius: 18px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: transform 0.3s ease;
  /* 其他样式 */
}

.badge-icon .logo-image {
  width: 100px;  /* 根据logo的实际比例调整 */
  height: 100px;
  object-fit: contain;
}

.badge-icon:hover {
  transform: translateY(-5px);
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
  display: block;
  font-size: 3.5rem;
  font-weight: 800;
  color: var(--primary-color);
  line-height: 1.2;
  margin-bottom: 0.5rem;
  letter-spacing: -0.5px;
  white-space: nowrap;
}

.lab-mission {
  font-size: 1.3rem;
  line-height: 1.7;
  color: var(--text-secondary);
  margin-bottom: 3rem;
  font-weight: 400;
  max-width: 90%;
}

.mission-stats {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 2rem;
  margin-bottom: 3rem;
}

.stat-item {
  text-align: center;
  padding: 1.5rem;
  background: var(--bg-card);
  border-radius: var(--border-radius);
  box-shadow: var(--shadow-soft);
  transition: transform 0.3s ease;
}

.stat-item:hover {
  transform: translateY(-5px);
}

.stat-number {
  display: block;
  font-size: 2.5rem;
  font-weight: 800;
  color: var(--primary-color);
  line-height: 1;
  margin-bottom: 0.5rem;
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
  border-radius: 12px;
  padding: 1rem 2rem;
  font-weight: 600;
  font-size: 1.1rem;
  transition: all 0.3s ease;
  border: none;
}

.cta-button.primary {
  background: var(--gradient-primary);
  color: white;
  box-shadow: var(--shadow-soft);
}

.cta-button.primary:hover {
  transform: translateY(-3px);
  box-shadow: 0 12px 30px rgba(125, 18, 49, 0.4);
}

.cta-button.secondary {
  background: transparent;
  color: var(--primary-color);
  border: 2px solid var(--primary-color);
}

.cta-button.secondary:hover {
  background: var(--primary-color);
  color: white;
  transform: translateY(-3px);
}

/* 英雄区域视觉元素 */
.hero-visual {
  position: relative;
  height: 500px;
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
  background: var(--bg-card);
  padding: 1.2rem;
  border-radius: var(--border-radius);
  box-shadow: var(--shadow-soft);
  display: flex;
  align-items: center;
  gap: 1rem;
  font-weight: 500;
  color: var(--text-primary);
  animation: floatCard 8s ease-in-out infinite;
  border: 1px solid rgba(125, 18, 49, 0.1);
  backdrop-filter: blur(10px);
  max-width: 220px;
  transition: all 0.3s ease;
}

.card:hover {
  transform: translateY(-10px) scale(1.05);
  box-shadow: var(--shadow-large);
}

.card-1 {
  top: 10%;
  left: 5%;
  animation-delay: 0s;
}

.card-2 {
  top: 30%;
  right: 10%;
  animation-delay: -2s;
}

.card-3 {
  bottom: 25%;
  left: 15%;
  animation-delay: -4s;
}

.card-4 {
  bottom: 5%;
  right: 5%;
  animation-delay: -6s;
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
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-15px);
  }
}

/* 亮点区域 */
.highlights-section {
  padding: 4rem 0;
  background: linear-gradient(135deg, rgba(125, 18, 49, 0.03) 0%, rgba(19, 57, 62, 0.03) 100%);
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
  gap: 1.5rem;
  padding: 2rem;
  background: var(--bg-white);
  border-radius: var(--border-radius);
  box-shadow: var(--shadow-soft);
  transition: transform 0.3s ease;
}

.highlight-item:hover {
  transform: translateY(-5px);
}

.highlight-icon {
  font-size: 3rem;
  flex-shrink: 0;
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
  font-size: 2.8rem;
  font-weight: 700;
  color: var(--text-primary);
  margin-bottom: 1.2rem;
  position: relative;
  display: inline-block;
}

.section-title::after {
  content: '';
  position: absolute;
  bottom: -10px;
  left: 50%;
  transform: translateX(-50%);
  width: 80px;
  height: 4px;
  background: var(--gradient-primary);
  border-radius: 2px;
}

.section-subtitle {
  font-size: 1.3rem;
  color: var(--text-secondary);
  max-width: 600px;
  margin: 0 auto;
  line-height: 1.6;
}

/* 新闻模块样式 - 时间线布局 */
.news-section {
  background: linear-gradient(135deg, rgba(125, 18, 49, 0.02) 0%, rgba(19, 57, 62, 0.02) 100%);
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
  padding: 2rem;
  background: var(--bg-white);
  border-radius: var(--border-radius);
  box-shadow: var(--shadow-soft);
  margin-bottom: 1.5rem;
  cursor: pointer;
  transition: all 0.3s ease;
  border-left: 4px solid transparent;
  position: relative;
  overflow: hidden;
}

.news-item::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  height: 100%;
  width: 4px;
  background: var(--gradient-primary);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.news-item:hover {
  transform: translateX(10px);
  box-shadow: var(--shadow-medium);
  border-left-color: var(--primary-color);
}

.news-item:hover::before {
  opacity: 1;
}

.news-date {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-width: 80px;
  padding: 1rem;
  background: rgba(125, 18, 49, 0.05);
  border-radius: var(--border-radius);
  border: 1px solid rgba(125, 18, 49, 0.1);
}

.date-day {
  font-size: 2rem;
  font-weight: 700;
  color: var(--primary-color);
  line-height: 1;
}

.date-month {
  font-size: 0.9rem;
  font-weight: 600;
  color: var(--text-secondary);
  text-transform: uppercase;
  margin-top: 0.2rem;
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
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.type-publication {
  background: rgba(125, 18, 49, 0.1);
  color: var(--primary-color);
}

.type-team {
  background: rgba(52, 152, 219, 0.1);
  color: var(--accent-color);
}

.type-award {
  background: rgba(241, 196, 15, 0.1);
  color: #f39c12;
}

.type-event {
  background: rgba(46, 204, 113, 0.1);
  color: #27ae60;
}

.news-year {
  font-size: 0.9rem;
  color: var(--text-light);
  font-weight: 500;
}

.news-title {
  font-size: 1.4rem;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 0.8rem;
  line-height: 1.4;
}

.news-excerpt {
  color: var(--text-secondary);
  line-height: 1.6;
  margin: 0;
}

.news-arrow {
  color: var(--text-light);
  transition: all 0.3s ease;
}

.news-item:hover .news-arrow {
  color: var(--primary-color);
  transform: translateX(5px);
}

.news-actions {
  text-align: center;
}

.view-all-button {
  border-radius: 12px;
  padding: 0.8rem 2rem;
  font-weight: 600;
  border: 2px solid var(--primary-color);
  color: var(--primary-color);
  background: transparent;
  transition: all 0.3s ease;
}

.view-all-button:hover {
  background: var(--primary-color);
  color: white;
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(125, 18, 49, 0.3);
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
  transition: all 0.3s ease;
  border: 1px solid rgba(125, 18, 49, 0.1);
  background: var(--bg-white);
  overflow: hidden;
}

.nav-card:hover {
  transform: translateY(-8px);
  box-shadow: var(--shadow-large);
  border-color: var(--primary-color);
}

.nav-card-content {
  padding: 2.5rem 2rem;
  text-align: center;
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.nav-icon {
  font-size: 3.5rem;
  margin-bottom: 1.5rem;
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