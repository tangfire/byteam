<script lang="ts" setup>
import { computed } from 'vue'
import { getPublicNews } from '../api/public'
import { usePublicList } from '../composables/usePublicList'
import { fallbackNewsActivities, type NewsActivity } from '../data/fallbacks/publicContent'

const activities = usePublicList<NewsActivity>({
  fallback: fallbackNewsActivities,
  load: async () => {
    const result = await getPublicNews()
    return result.items.map((item) => ({
      content: item.content,
      timestamp: item.date,
      color: item.color || '#7d1231',
      type: item.type as NewsActivity['type'],
    }))
  },
  fallbackMessage: 'Using local news fallback data',
})

const activitiesByYear = computed(() => {
  const groups: Record<string, NewsActivity[]> = {}
  activities.value.forEach((activity) => {
    const year = String(activity.timestamp || '').slice(0, 4) || 'Other'
    groups[year] = groups[year] || []
    groups[year].push(activity)
  })
  return Object.keys(groups)
    .sort((a, b) => Number(b) - Number(a))
    .map((year) => ({
      year,
      items: groups[year].slice().sort((a, b) => String(b.timestamp || '').localeCompare(String(a.timestamp || ''))),
    }))
})

const getNewsIcon = (type: string = 'general') => {
  const icons = {
    publication: '📄',
    team: '👥',
    award: '🏆',
    event: '🎉',
    general: '📢'
  }
  return icons[type as keyof typeof icons] || icons.general
}

const monthNames = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']

const getDateParts = (timestamp?: string) => {
  const match = String(timestamp || '').match(/^(\d{4})-(\d{2})-(\d{2})/)
  if (!match) {
    return { month: 'Date', day: '--', full: timestamp || 'Undated' }
  }

  const monthIndex = Number(match[2]) - 1
  return {
    month: monthNames[monthIndex] || 'Date',
    day: match[3],
    full: `${match[1]}-${match[2]}-${match[3]}`,
  }
}
</script>

<template>
  <div class="news-container">
    <!-- 页面标题区域 -->
    <div class="page-header">
      <h1 class="page-title">News & Announcements</h1>
      <p class="page-subtitle">Latest updates and achievements from our research team</p>
    </div>

    <section v-for="group in activitiesByYear" :key="group.year" class="year-section">
      <div class="year-header">
        <h2 class="year-title">{{ group.year }}</h2>
        <span class="year-count">{{ group.items.length }} {{ group.items.length === 1 ? 'update' : 'updates' }}</span>
      </div>

      <div class="news-list">
        <article
            v-for="(activity, index) in group.items"
            :key="`${group.year}-${activity.timestamp || 'undated'}-${index}`"
            class="news-item"
        >
          <time class="news-date" :datetime="getDateParts(activity.timestamp).full">
            <span class="date-month">{{ getDateParts(activity.timestamp).month }}</span>
            <span class="date-day">{{ getDateParts(activity.timestamp).day }}</span>
          </time>
          <div class="news-icon">{{ getNewsIcon(activity.type) }}</div>
          <p class="news-content">{{ activity.content }}</p>
        </article>
      </div>
    </section>

    <el-backtop class="mobile-backtop" :right="100" :bottom="100"/>
  </div>
</template>

<style scoped>
.news-container {
  max-width: 1080px;
  margin: 0 auto;
  padding: 42px 24px 56px;
  min-height: calc(100vh - 220px);
}

/* 页面标题样式 */
.page-header {
  text-align: center;
  margin-bottom: 48px;
  padding: 0 20px;
}

.page-title {
  font-size: 3rem;
  color: #7d1231;
  margin-bottom: 15px;
  font-weight: 700;
  letter-spacing: 0;
}

.page-subtitle {
  font-size: 1.3rem;
  color: #7f8c8d;
  max-width: 500px;
  margin: 0 auto;
  line-height: 1.6;
  font-weight: 400;
}

/* 年份区域 */
.year-section {
  position: relative;
  margin-bottom: 54px;
}

.year-section:last-of-type {
  margin-bottom: 0;
}

.year-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 20px;
  margin-bottom: 14px;
  padding-bottom: 12px;
  border-bottom: 2px solid #f1e7eb;
}

.year-title {
  color: #7d1231;
  font-size: 2rem;
  font-weight: 700;
  line-height: 1.2;
  margin: 0;
}

.year-count {
  color: #7a8492;
  font-size: 0.95rem;
  font-weight: 500;
  white-space: nowrap;
}

.news-list {
  display: flex;
  flex-direction: column;
}

/* 新闻项样式 */
.news-item {
  display: grid;
  grid-template-columns: 76px 42px minmax(0, 1fr);
  gap: 18px;
  align-items: flex-start;
  border-bottom: 1px solid #edf0f2;
  padding: 22px 0;
  transition: border-color 0.2s ease, background-color 0.2s ease;
}

.news-item:hover {
  background-color: #fbf8f9;
  border-bottom-color: #e5d4db;
}

.news-date {
  color: #5f6875;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  font-variant-numeric: tabular-nums;
  line-height: 1;
  padding-top: 2px;
}

.date-month {
  font-size: 0.82rem;
  font-weight: 700;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

.date-day {
  color: #273445;
  font-size: 1.7rem;
  font-weight: 700;
  margin-top: 6px;
}

.news-icon {
  align-items: center;
  background: #f7eef2;
  border: 1px solid #ead0d9;
  border-radius: 10px;
  color: #7d1231;
  display: flex;
  font-size: 1.25rem;
  height: 42px;
  justify-content: center;
  transition: transform 0.2s ease, background-color 0.2s ease;
  width: 42px;
}

.news-item:hover .news-icon {
  background: #f2e3e9;
  transform: translateY(-1px);
}

.news-content {
  color: #2c3e50;
  font-size: 1.08rem;
  line-height: 1.7;
  margin: 0;
  word-break: break-word;
  white-space: normal;
}

/* 返回顶部按钮 */
.mobile-backtop {
  right: 100px;
  bottom: 100px;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .news-container {
    padding: 25px 15px;
  }

  .page-title {
    font-size: 2.3rem;
  }

  .page-subtitle {
    font-size: 1.1rem;
    padding: 0 10px;
  }

  .year-header {
    align-items: flex-start;
    flex-direction: column;
    gap: 6px;
  }

  .year-title {
    font-size: 1.75rem;
  }

  .news-item {
    grid-template-columns: 58px 36px minmax(0, 1fr);
    gap: 12px;
    padding: 18px 0;
  }

  .news-icon {
    border-radius: 9px;
    font-size: 1.1rem;
    height: 36px;
    width: 36px;
  }

  .news-content {
    font-size: 15px;
    line-height: 1.6;
  }

  .date-day {
    font-size: 1.45rem;
  }

  .mobile-backtop {
    right: 20px !important;
    bottom: 80px !important;
  }
}

/* 小屏手机优化 */
@media (max-width: 480px) {
  .news-container {
    padding: 20px 10px;
  }

  .page-title {
    font-size: 2rem;
  }

  .page-header {
    margin-bottom: 40px;
  }

  .news-item {
    grid-template-columns: 48px 32px minmax(0, 1fr);
    gap: 10px;
    padding: 16px 0;
  }

  .date-month {
    font-size: 0.72rem;
  }

  .date-day {
    font-size: 1.25rem;
    margin-top: 5px;
  }

  .news-icon {
    border-radius: 8px;
    font-size: 1rem;
    height: 32px;
    width: 32px;
  }

  .news-content {
    font-size: 14px;
  }
}

/* 大屏幕优化 */
@media (min-width: 1400px) {
  .news-container {
    max-width: 1100px;
  }
}
</style>
