const API_BASE = import.meta.env.VITE_API_BASE || ''
const TOKEN_KEY = 'byml_admin_token'

export interface ListResponse<T> {
  items: T[]
  total: number
  page: number
  pageSize: number
}

export interface NewsItem {
  id?: number
  title: string
  content: string
  excerpt: string
  type: string
  typeLabel: string
  date: string
  color: string
  status: string
  sortOrder: number
}

export interface Person {
  id?: number
  name: string
  avatarUrl: string
  category: string
  research: string
  graduationDate: string
  status: string
  sortOrder: number
}

export interface UndergraduateEducation {
  id?: number
  name: string
  grade: string
  major: string
  direction: string
  achievements: string[]
  status: string
  sortOrder: number
}

export interface PublicationLink {
  id?: number
  type: string
  label: string
  url: string
  routeName: string
  sortOrder: number
}

export interface Publication {
  id?: number
  image: string
  title: string
  authors: string
  venue: string
  year: number
  kind: string
  status: string
  featured: boolean
  sortOrder: number
  links: PublicationLink[]
}

export interface Patent {
  id?: number
  authors: string
  title: string
  date: string
  country: string
  number: string
  category: string
  status: string
  sortOrder: number
}

export interface ResearchProject {
  id?: number
  title: string
  fund: string
  number: string
  period: string
  amount: string
  projectStatus: string
  role: string
  status: string
  sortOrder: number
}

export interface MediaAsset {
  id: number
  fileName: string
  originalName: string
  url: string
  mimeType: string
  size: number
  kind: string
  createdAt: string
}

export interface TrashItem {
  resource: string
  label: string
  id: number
  title: string
  subtitle: string
  status: string
  deletedAt?: string
}

export interface HomePayload {
  latestNews: NewsItem[]
  featuredPublications: Publication[]
  stats: {
    publications: number
    projects: number
    teamMembers: number
  }
}

export const authStore = {
  get token() {
    return localStorage.getItem(TOKEN_KEY) || ''
  },
  setToken(token: string) {
    localStorage.setItem(TOKEN_KEY, token)
  },
  clear() {
    localStorage.removeItem(TOKEN_KEY)
  },
}

export async function apiRequest<T>(path: string, options: RequestInit = {}): Promise<T> {
  const headers = new Headers(options.headers)
  if (!(options.body instanceof FormData) && !headers.has('Content-Type')) {
    headers.set('Content-Type', 'application/json')
  }
  const token = authStore.token
  if (token) {
    headers.set('Authorization', `Bearer ${token}`)
  }

  const response = await fetch(`${API_BASE}${path}`, {
    ...options,
    headers,
  })

  if (!response.ok) {
    let message = `Request failed: ${response.status}`
    try {
      const body = await response.json()
      message = body.error || message
    } catch {
      // Keep the HTTP status message.
    }
    throw new Error(message)
  }

  if (response.status === 204) {
    return undefined as T
  }
  return response.json() as Promise<T>
}

export function listQuery(params: Record<string, string | number | undefined>) {
  const search = new URLSearchParams()
  Object.entries(params).forEach(([key, value]) => {
    if (value !== undefined && value !== '') {
      search.set(key, String(value))
    }
  })
  const query = search.toString()
  return query ? `?${query}` : ''
}
