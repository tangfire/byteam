import { apiRequest, type HomePayload, type NewsItem, type Person, type Publication, type Patent, type ResearchProject, type UndergraduateEducation } from './client'

export async function getHome() {
  return apiRequest<HomePayload>('/api/public/home')
}

export async function getPublicNews() {
  return apiRequest<{ items: NewsItem[] }>('/api/public/news')
}

export async function getPublicPeople(category?: string) {
  return apiRequest<{ items: Person[] }>(`/api/public/people${category ? `?category=${category}` : ''}`)
}

export async function getPublicUndergraduates() {
  return apiRequest<{ items: UndergraduateEducation[] }>('/api/public/undergraduates')
}

export async function getPublicPublications() {
  return apiRequest<{ items: Publication[]; groups: Record<string, Record<string, Publication[]>> }>('/api/public/publications')
}

export async function getPublicPatents() {
  return apiRequest<{ items: Patent[]; groups: Record<string, Patent[]> }>('/api/public/patents')
}

export async function getPublicResearchProjects() {
  return apiRequest<{ items: ResearchProject[] }>('/api/public/research-projects')
}
