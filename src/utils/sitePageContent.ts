import type { SitePage } from '../api/client'

export const sourceManagedSlugs = new Set(['about', 'contact', 'videomind', 'vknow'])

export const pageLabel = (slug: string, title: string) => {
  const labels: Record<string, string> = {
    'research-direction': 'Research Direction',
    'dr-baoyao-yang': 'Baoyao Yang',
    'video-xiaoqi-zheng-01': '视频：Xiaoqi Zheng',
    'video-xianrun-xu-01': '视频：Xianrun Xu',
    'video-yali-ma-01': '视频：Yali Ma',
  }
  return labels[slug] || title
}

export const slugifyVideoTitle = (value: string) => {
  const slug = value
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-+|-+$/g, '')
    .replace(/-{2,}/g, '-')
  return `video-${slug || 'publication'}`
}

export const toPlainPage = (page: SitePage) => ({
  slug: page.slug,
  title: page.title,
  description: page.description,
  content: page.content,
  status: page.status,
  sortOrder: page.sortOrder,
})

const fallbackContent: Record<string, () => Record<string, any>> = {
  'research-direction': () => ({ directions: [] }),
  'dr-baoyao-yang': () => ({ name: '', image: '', alt: '', paragraphs: [] }),
}

const clonePlain = <T>(value: T): T => JSON.parse(JSON.stringify(value))

const isPlainObject = (value: unknown): value is Record<string, any> => {
  return Boolean(value) && typeof value === 'object' && !Array.isArray(value)
}

const toRecord = (value: unknown): Record<string, any> => isPlainObject(value) ? value : {}

const normalizeText = (value: unknown) => typeof value === 'string' ? value : ''

const normalizeTextItems = (value: unknown) => {
  if (Array.isArray(value)) {
    return value.map((item) => typeof item === 'string' ? item : String(item ?? ''))
  }
  if (typeof value === 'string') {
    return value.split('\n').map((item) => item.trim()).filter(Boolean)
  }
  return []
}

const normalizeSections = (value: unknown) => {
  if (!Array.isArray(value)) return []
  return value.map((section) => {
    if (typeof section === 'string') return { title: '', text: section }
    const record = toRecord(section)
    return {
      ...record,
      title: normalizeText(record.title),
      text: normalizeText(record.text),
    }
  })
}

const normalizeDirections = (value: unknown) => {
  if (!Array.isArray(value)) return []
  return value.map((direction) => {
    const record = toRecord(direction)
    return {
      ...record,
      title: normalizeText(record.title),
      image: normalizeText(record.image),
      alt: normalizeText(record.alt),
      sections: normalizeSections(record.sections),
    }
  })
}

const deepMerge = (base: Record<string, any>, value: Record<string, any>) => {
  const out = clonePlain(base)
  Object.entries(value || {}).forEach(([key, nextValue]) => {
    if (isPlainObject(nextValue) && isPlainObject(out[key])) {
      out[key] = deepMerge(out[key], nextValue)
    } else {
      out[key] = nextValue
    }
  })
  return out
}

export const normalizePageContent = (page: SitePage) => {
  const fallback = fallbackContent[page.slug]?.() || (page.slug.startsWith('video-') ? { title: '', video: '' } : {})
  page.content = deepMerge(fallback, page.content || {})

  if (page.slug === 'research-direction') {
    page.content.directions = normalizeDirections(page.content.directions)
  } else if (page.slug === 'dr-baoyao-yang') {
    page.content.name = normalizeText(page.content.name)
    page.content.image = normalizeText(page.content.image)
    page.content.alt = normalizeText(page.content.alt)
    page.content.paragraphs = normalizeTextItems(page.content.paragraphs)
  } else if (page.slug.startsWith('video-')) {
    page.content.title = normalizeText(page.content.title)
    page.content.video = normalizeText(page.content.video)
  }
}
