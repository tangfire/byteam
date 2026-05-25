import { computed, onMounted, ref } from 'vue'
import { getPublicPage } from '../api/public'
import { ApiError, type SitePage } from '../api/client'

export function useSitePage(slug: string, fallbackContent: Record<string, any> = {}) {
  const page = ref<SitePage | null>(null)
  const loading = ref(true)
  const error = ref('')
  const hidden = ref(false)

  const content = computed<Record<string, any>>(() => page.value?.content || fallbackContent)

  onMounted(async () => {
    loading.value = true
    error.value = ''
    try {
      page.value = (await getPublicPage(slug)).page
    } catch (err) {
      error.value = err instanceof Error ? err.message : '页面内容加载失败'
      hidden.value = err instanceof ApiError && err.status === 404
      if (hidden.value) {
        page.value = null
        return
      }
      page.value = {
        slug,
        title: fallbackContent.title || '',
        description: '',
        content: fallbackContent,
        status: 'draft',
        sortOrder: 0,
      }
    } finally {
      loading.value = false
    }
  })

  return { page, content, loading, error, hidden }
}
