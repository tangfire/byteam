import { apiRequest, authStore, listQuery, type ListResponse, type MaintenanceCommandResult, type MaintenanceStatus, type MediaAsset, type NewsItem, type Patent, type Person, type Publication, type ResearchProject, type SitePage, type TrashItem, type UndergraduateEducation } from './client'

export async function login(username: string, password: string) {
  const result = await apiRequest<{ token: string; admin: { id: number; username: string } }>('/api/admin/auth/login', {
    method: 'POST',
    body: JSON.stringify({ username, password }),
  })
  authStore.setToken(result.token)
  return result
}

export function logout() {
  authStore.clear()
}

export async function getMe() {
  return apiRequest<{ admin: { id: number; username: string } }>('/api/admin/me')
}

export async function getSummary() {
  return apiRequest<{ counts: Record<string, number> }>('/api/admin/summary')
}

export function listAdmin<T>(resource: string, params: Record<string, string | number | undefined>) {
  return apiRequest<ListResponse<T>>(`/api/admin/${resource}${listQuery(params)}`)
}

export function listSitePages(params: Record<string, string | number | undefined> = {}) {
  return apiRequest<ListResponse<SitePage>>(`/api/admin/pages${listQuery(params)}`)
}

export function getSitePage(slug: string) {
  return apiRequest<SitePage>(`/api/admin/pages/${slug}`)
}

export function createSitePage(payload: SitePage) {
  return apiRequest<SitePage>('/api/admin/pages', {
    method: 'POST',
    body: JSON.stringify(payload),
  })
}

export function updateSitePage(slug: string, payload: SitePage) {
  return apiRequest<SitePage>(`/api/admin/pages/${slug}`, {
    method: 'PUT',
    body: JSON.stringify(payload),
  })
}

export function deleteSitePage(slug: string) {
  return apiRequest<{ deleted: boolean }>(`/api/admin/pages/${slug}`, {
    method: 'DELETE',
  })
}

export function createAdmin<T>(resource: string, payload: T) {
  return apiRequest<T>(`/api/admin/${resource}`, {
    method: 'POST',
    body: JSON.stringify(payload),
  })
}

export function updateAdmin<T extends { id?: number }>(resource: string, payload: T) {
  return apiRequest<T>(`/api/admin/${resource}/${payload.id}`, {
    method: 'PUT',
    body: JSON.stringify(payload),
  })
}

export function placeAdmin(resource: string, id: number, targetId: number, position: 'before' | 'after' = 'before') {
  return apiRequest<AdminEntity>(`/api/admin/${resource}/${id}/place`, {
    method: 'POST',
    body: JSON.stringify({ targetId, position }),
  })
}

export function deleteAdmin(resource: string, id: number) {
  return apiRequest<{ deleted: boolean }>(`/api/admin/${resource}/${id}`, {
    method: 'DELETE',
  })
}

export function listTrash(params: Record<string, string | number | undefined>) {
  return apiRequest<ListResponse<TrashItem>>(`/api/admin/trash${listQuery(params)}`)
}

export function restoreTrash(resource: string, id: number) {
  return apiRequest<{ restored: boolean }>(`/api/admin/trash/${resource}/${id}/restore`, {
    method: 'POST',
  })
}

export function importPublicMedia() {
  return apiRequest<{ scanned: number; created: number; updated: number; skipped: number }>('/api/admin/media/import-public', {
    method: 'POST',
  })
}

export function updateMediaName(id: number, displayName: string) {
  return apiRequest<MediaAsset>(`/api/admin/media/${id}`, {
    method: 'PUT',
    body: JSON.stringify({ displayName }),
  })
}

export function getMaintenanceStatus() {
  return apiRequest<MaintenanceStatus>('/api/admin/maintenance/status')
}

export function runBackup() {
  return apiRequest<MaintenanceCommandResult>('/api/admin/maintenance/backup', {
    method: 'POST',
  })
}

export function runGitSync() {
  return apiRequest<MaintenanceCommandResult>('/api/admin/maintenance/git-sync', {
    method: 'POST',
  })
}

export async function uploadMedia(file: File) {
  const body = new FormData()
  body.append('file', file)
  return apiRequest<MediaAsset>('/api/admin/media', {
    method: 'POST',
    body,
  })
}

export type AdminEntity = NewsItem | Person | UndergraduateEducation | Publication | Patent | ResearchProject | MediaAsset | SitePage
