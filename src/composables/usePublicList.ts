import { onMounted, ref, type Ref } from 'vue'

type UsePublicListOptions<T> = {
  fallback: T[]
  load: () => Promise<T[]>
  fallbackMessage: string
}

export function usePublicList<T>({ fallback, load, fallbackMessage }: UsePublicListOptions<T>): Ref<T[]> {
  const items = ref<T[]>(fallback) as Ref<T[]>

  onMounted(async () => {
    try {
      items.value = await load()
    } catch (error) {
      console.warn(fallbackMessage, error)
    }
  })

  return items
}
