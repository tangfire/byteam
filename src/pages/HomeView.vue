<template>
  <div class="home-container">
    <section class="hero-band">
      <div class="hero-section" aria-labelledby="home-title">
        <div class="hero-copy">
          <div class="identity-row">
            <img src="/logo/001.png" alt="BYML Logo" class="logo-image">
            <span>Guangdong University of Technology</span>
          </div>

          <h1 id="home-title" class="lab-title">Beyond Machine Learning Group</h1>

          <p class="lab-mission">
            We study machine learning methods for multimodal data, with a focus on medical image
            analysis, video understanding, federated learning, and robust model training.
          </p>

          <div class="hero-actions">
            <el-button type="primary" size="large" @click="scrollToResearch" class="primary-action">
              Latest Research
              <el-icon><ArrowRight /></el-icon>
            </el-button>
            <el-button size="large" @click="$router.push('/about')" class="secondary-action">
              About the Group
            </el-button>
          </div>
        </div>

        <aside class="hero-summary" aria-label="BYML summary">
          <div class="stats-grid">
            <div class="stat-item">
              <span class="stat-number">{{ stats.publications }}+</span>
              <span class="stat-label">Publications</span>
            </div>
            <div class="stat-item">
              <span class="stat-number">{{ stats.projects }}+</span>
              <span class="stat-label">Research Projects</span>
            </div>
            <div class="stat-item">
              <span class="stat-number">{{ stats.teamMembers }}+</span>
              <span class="stat-label">Team Members</span>
            </div>
          </div>

          <div class="theme-list">
            <h2>Research Focus</h2>
            <ul>
              <li v-for="theme in researchThemes" :key="theme">{{ theme }}</li>
            </ul>
          </div>
        </aside>
      </div>
    </section>

    <main class="main-content">
      <section class="overview-section" aria-labelledby="overview-title">
        <div class="section-heading">
          <span class="section-kicker">Overview</span>
          <h2 id="overview-title" class="section-title">Research With Practical Data Conditions</h2>
        </div>

        <div class="overview-grid">
          <article v-for="item in overviewItems" :key="item.title" class="overview-item">
            <h3>{{ item.title }}</h3>
            <p>{{ item.description }}</p>
          </article>
        </div>
      </section>

      <section class="research-section" id="research-highlights" aria-labelledby="research-title">
        <div class="section-heading">
          <span class="section-kicker">Publications</span>
          <h2 id="research-title" class="section-title">Latest Research</h2>
          <p class="section-subtitle">Recent representative work from the group.</p>
        </div>

        <ResearchCarousel class="research-carousel-wrapper" />
      </section>

      <section class="news-section" aria-labelledby="news-title">
        <div class="section-heading">
          <span class="section-kicker">Updates</span>
          <h2 id="news-title" class="section-title">Latest News</h2>
          <p class="section-subtitle">Recent publications, group activities, and announcements.</p>
        </div>

        <div class="news-list">
          <article
              v-for="(news, index) in latestNews"
              :key="`${news.date}-${news.title}-${index}`"
              class="news-item"
              @click="$router.push('/news')"
          >
            <time class="news-date" :datetime="news.date">
              <span class="date-day">{{ getDay(news.date) }}</span>
              <span class="date-month">{{ getMonth(news.date) }} {{ getYear(news.date) }}</span>
            </time>
            <div class="news-content">
              <span class="news-type" :class="getNewsTypeClass(news.type)">{{ news.typeLabel }}</span>
              <h3 class="news-title">{{ news.title }}</h3>
              <p class="news-excerpt">{{ news.excerpt }}</p>
            </div>
            <el-icon class="news-arrow"><ArrowRight /></el-icon>
          </article>
        </div>

        <div class="section-actions">
          <el-button type="primary" plain @click="$router.push('/news')" class="link-button">
            View All News
            <el-icon><View /></el-icon>
          </el-button>
        </div>
      </section>

      <section class="quick-nav-section" aria-labelledby="explore-title">
        <div class="section-heading">
          <span class="section-kicker">Explore</span>
          <h2 id="explore-title" class="section-title">More From BYML</h2>
        </div>

        <div class="nav-list">
          <button
              v-for="link in homeQuickLinks"
              :key="link.route"
              class="nav-item"
              type="button"
              @click="$router.push(link.route)"
          >
            <span>
              <strong>{{ link.title }}</strong>
              <small>{{ link.description }}</small>
            </span>
            <el-icon><ArrowRight /></el-icon>
          </button>
        </div>
      </section>
    </main>

    <el-backtop class="mobile-backtop" :right="100" :bottom="100" />
  </div>
</template>

<script setup lang="ts">
import { ArrowRight, View } from '@element-plus/icons-vue'
import ResearchCarousel from '../components/ResearchCarousel.vue'
import { useHomeContent } from '../composables/useHomeContent'
import { homeQuickLinks } from '../data/homeSections'
import { getDay, getMonth, getNewsTypeClass, getYear, scrollToResearch } from '../utils/homeView'

const researchThemes = [
  'Multimodal medical image analysis',
  'Federated and privacy-aware learning',
  'Video understanding and cross-modal retrieval',
  'Robust learning with noisy or limited labels',
]

const overviewItems = [
  {
    title: 'Multimodal Learning',
    description: 'Connecting visual, textual, signal, and structured information for reliable model reasoning.',
  },
  {
    title: 'Healthcare Applications',
    description: 'Developing learning methods for medical image segmentation, diagnosis support, and clinical data analysis.',
  },
  {
    title: 'Distributed Intelligence',
    description: 'Studying federated optimization and collaborative learning when data remains across different sites.',
  },
]

const { latestNews, stats } = useHomeContent()
</script>

<style scoped>
.home-container {
  --primary-color: #7d1231;
  --primary-dark: #5a0c22;
  --text-primary: #273445;
  --text-secondary: #566273;
  --text-muted: #7a8492;
  --border-color: #e8edf1;
  --theme-border: #e4cfd7;
  --surface-muted: #faf7f8;

  background: #fff;
  color: var(--text-primary);
  min-height: 100vh;
}

.hero-band {
  background:
      linear-gradient(90deg, rgba(255, 255, 255, 0.94) 0%, rgba(255, 255, 255, 0.86) 48%, rgba(255, 255, 255, 0.72) 100%),
      url('/background/ContactBackground.jpg') center/cover;
  border-bottom: 1px solid var(--border-color);
}

.main-content {
  margin: 0 auto;
  max-width: 1180px;
  padding: 0 24px 76px;
}

.hero-section {
  align-items: start;
  display: grid;
  gap: 70px;
  grid-template-columns: minmax(0, 1.1fr) minmax(340px, 0.9fr);
  margin: 0 auto;
  max-width: 1180px;
  min-height: 620px;
  padding: 82px 24px 88px;
}

.hero-copy {
  max-width: 760px;
}

.identity-row {
  align-items: center;
  color: var(--primary-color);
  display: flex;
  gap: 14px;
  font-size: 0.98rem;
  font-weight: 700;
  margin-bottom: 26px;
}

.logo-image {
  display: block;
  height: 52px;
  object-fit: contain;
  width: 52px;
}

.lab-title {
  color: var(--primary-color);
  font-size: 4rem;
  font-weight: 750;
  letter-spacing: 0;
  line-height: 1.08;
  margin: 0;
}

.lab-mission {
  border-left: 4px solid var(--primary-color);
  color: var(--text-secondary);
  font-size: 1.22rem;
  line-height: 1.85;
  margin: 30px 0 0;
  max-width: 720px;
  padding-left: 22px;
}

.hero-actions {
  display: flex;
  flex-wrap: wrap;
  gap: 14px;
  margin-top: 34px;
}

.primary-action,
.secondary-action,
.link-button {
  border-radius: 6px;
  font-weight: 650;
}

.primary-action {
  background: var(--primary-color);
  border-color: var(--primary-color);
}

.primary-action:hover,
.primary-action:focus {
  background: var(--primary-dark);
  border-color: var(--primary-dark);
}

.secondary-action {
  border-color: var(--theme-border);
  color: var(--primary-color);
}

.hero-summary {
  border-top: 4px solid var(--primary-color);
  background: rgba(250, 247, 248, 0.93);
  padding: 28px;
}

.stats-grid {
  display: grid;
  gap: 0;
  grid-template-columns: repeat(3, 1fr);
  border: 1px solid var(--theme-border);
  background: #fff;
}

.stat-item {
  padding: 22px 14px;
  text-align: center;
}

.stat-item + .stat-item {
  border-left: 1px solid var(--theme-border);
}

.stat-number {
  color: var(--primary-color);
  display: block;
  font-size: 2rem;
  font-weight: 750;
  line-height: 1;
}

.stat-label {
  color: var(--text-muted);
  display: block;
  font-size: 0.9rem;
  line-height: 1.4;
  margin-top: 8px;
}

.theme-list {
  margin-top: 30px;
}

.theme-list h2 {
  color: var(--primary-color);
  font-size: 1.35rem;
  font-weight: 700;
  margin: 0 0 14px;
}

.theme-list ul {
  list-style: none;
  margin: 0;
  padding: 0;
}

.theme-list li {
  border-top: 1px solid var(--theme-border);
  color: var(--text-secondary);
  line-height: 1.6;
  padding: 13px 0;
}

.overview-section,
.research-section,
.news-section,
.quick-nav-section {
  border-top: 1px solid var(--border-color);
  padding: 58px 0;
}

.section-heading {
  margin-bottom: 28px;
  max-width: 760px;
}

.section-kicker {
  color: var(--primary-color);
  display: block;
  font-size: 0.92rem;
  font-weight: 700;
  margin-bottom: 8px;
}

.section-title {
  color: var(--primary-color);
  font-size: 2.1rem;
  font-weight: 700;
  letter-spacing: 0;
  line-height: 1.25;
  margin: 0;
}

.section-subtitle {
  color: var(--text-secondary);
  font-size: 1.06rem;
  line-height: 1.7;
  margin: 12px 0 0;
}

.overview-grid {
  display: grid;
  gap: 28px;
  grid-template-columns: repeat(3, 1fr);
}

.overview-item {
  border-top: 3px solid var(--primary-color);
  padding-top: 18px;
}

.overview-item h3 {
  color: var(--text-primary);
  font-size: 1.2rem;
  font-weight: 700;
  margin: 0;
}

.overview-item p {
  color: var(--text-secondary);
  font-size: 1rem;
  line-height: 1.75;
  margin: 10px 0 0;
}

.research-carousel-wrapper {
  max-width: 1000px;
  margin: 0 auto;
}

.news-list {
  border-top: 1px solid var(--theme-border);
}

.news-item {
  align-items: start;
  background: transparent;
  border: 0;
  border-bottom: 1px solid var(--border-color);
  cursor: pointer;
  display: grid;
  gap: 26px;
  grid-template-columns: 110px minmax(0, 1fr) 28px;
  padding: 24px 0;
  text-align: left;
  width: 100%;
}

.news-item:hover .news-title,
.news-item:focus-within .news-title,
.news-item:hover .news-arrow {
  color: var(--primary-color);
}

.news-date {
  color: var(--text-muted);
  display: flex;
  flex-direction: column;
  font-variant-numeric: tabular-nums;
}

.date-day {
  color: var(--primary-color);
  font-size: 1.65rem;
  font-weight: 750;
  line-height: 1;
}

.date-month {
  font-size: 0.9rem;
  line-height: 1.4;
  margin-top: 7px;
}

.news-content {
  min-width: 0;
}

.news-type {
  color: var(--primary-color);
  display: inline-block;
  font-size: 0.82rem;
  font-weight: 700;
  margin-bottom: 8px;
}

.news-title {
  color: var(--text-primary);
  font-size: 1.2rem;
  font-weight: 700;
  line-height: 1.45;
  margin: 0;
}

.news-excerpt {
  color: var(--text-secondary);
  font-size: 0.98rem;
  line-height: 1.7;
  margin: 8px 0 0;
}

.news-arrow {
  color: var(--text-muted);
  margin-top: 30px;
}

.section-actions {
  margin-top: 26px;
}

.link-button {
  border-color: var(--primary-color);
  color: var(--primary-color);
}

.nav-list {
  border-top: 1px solid var(--theme-border);
  display: grid;
  grid-template-columns: repeat(3, 1fr);
}

.nav-item {
  align-items: center;
  background: transparent;
  border: 0;
  border-bottom: 1px solid var(--border-color);
  border-right: 1px solid var(--border-color);
  color: inherit;
  cursor: pointer;
  display: flex;
  gap: 22px;
  justify-content: space-between;
  min-height: 150px;
  padding: 24px;
  text-align: left;
}

.nav-item:nth-child(3n) {
  border-right: 0;
}

.nav-item:hover strong,
.nav-item:hover .el-icon {
  color: var(--primary-color);
}

.nav-item strong {
  color: var(--text-primary);
  display: block;
  font-size: 1.16rem;
  line-height: 1.35;
}

.nav-item small {
  color: var(--text-secondary);
  display: block;
  font-size: 0.95rem;
  line-height: 1.55;
  margin-top: 8px;
}

.nav-item .el-icon {
  color: var(--text-muted);
  flex: 0 0 auto;
}

.mobile-backtop {
  right: 100px;
  bottom: 100px;
}

@media (max-width: 1024px) {
  .hero-section {
    grid-template-columns: 1fr;
    min-height: auto;
    padding-top: 30px;
  }

  .hero-summary {
    max-width: 680px;
  }

  .overview-grid {
    grid-template-columns: 1fr;
  }

  .nav-list {
    grid-template-columns: 1fr;
  }

  .nav-item,
  .nav-item:nth-child(3n) {
    border-right: 0;
  }
}

@media (max-width: 768px) {
  .main-content {
    padding: 28px 16px 58px;
  }

  .lab-title {
    font-size: 2.45rem;
  }

  .lab-mission {
    font-size: 1.08rem;
  }

  .hero-actions {
    flex-direction: column;
  }

  .stats-grid {
    grid-template-columns: 1fr;
  }

  .stat-item + .stat-item {
    border-left: 0;
    border-top: 1px solid var(--theme-border);
  }

  .section-title {
    font-size: 1.75rem;
  }

  .news-item {
    grid-template-columns: 74px minmax(0, 1fr);
  }

  .news-arrow {
    display: none;
  }

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }
}

@media (max-width: 480px) {
  .main-content {
    padding: 22px 12px 50px;
  }

  .identity-row {
    align-items: flex-start;
    flex-direction: column;
  }

  .lab-title {
    font-size: 2rem;
  }

  .hero-summary {
    padding: 20px;
  }

  .news-item {
    grid-template-columns: 1fr;
    gap: 12px;
  }
}
</style>
