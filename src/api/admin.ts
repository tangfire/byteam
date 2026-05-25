import { apiRequest, authStore, listQuery, type ListResponse, type MediaAsset, type NewsItem, type Patent, type Person, type Publication, type ResearchProject, type TrashItem, type UndergraduateEducation } from './client'

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

export async function uploadMedia(file: File) {
  const body = new FormData()
  body.append('file', file)
  return apiRequest<MediaAsset>('/api/admin/media', {
    method: 'POST',
    body,
  })
}

export type AdminEntity = NewsItem | Person | UndergraduateEducation | Publication | Patent | ResearchProject | MediaAsset
