import { onMounted, ref } from 'vue'
import { getHome } from '../api/public'
import type { Publication } from '../api/client'
import { fallbackFeaturedPublications } from '../data/fallbacks/featuredPublications'

export const useFeaturedPublications = () => {
  const publications = ref<Publication[]>(fallbackFeaturedPublications)

  onMounted(async () => {
    try {
      const data = await getHome()
      if (data.featuredPublications.length > 0) {
        publications.value = data.featuredPublications
      }
    } catch (error) {
      console.warn('Using local carousel fallback data', error)
    }
  })

  return publications
}
