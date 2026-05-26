import { onMounted, ref } from 'vue'
import { getHome } from '../api/public'
import type { HomePayload, NewsItem } from '../api/client'
import { fallbackHomeNews } from '../data/fallbacks/publicContent'

const fallbackStats: HomePayload['stats'] = {
  publications: 50,
  projects: 15,
  teamMembers: 20,
}

export const useHomeContent = () => {
  const latestNews = ref<NewsItem[]>(fallbackHomeNews)
  const stats = ref<HomePayload['stats']>(fallbackStats)

  onMounted(async () => {
    try {
      const data = await getHome()
      latestNews.value = data.latestNews
      stats.value = data.stats
    } catch (error) {
      console.warn('Using local home fallback data', error)
    }
  })

  return {
    latestNews,
    stats,
  }
}
