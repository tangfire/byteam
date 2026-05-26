import type { PublicationLink } from '../api/client'

export const legacyVideoRouteSlugs: Record<string, string> = {
  'video-player-XiaoqiZheng01': 'video-xiaoqi-zheng-01',
  'video-player-XianrunXu01': 'video-xianrun-xu-01',
  'video-player-YaliMa01': 'video-yali-ma-01',
}

export type VideoLinkTarget = Pick<PublicationLink, 'url' | 'routeName'>

export const videoPagePath = (slug: string) => `/video/${slug}`

export const videoSlugFromPath = (url = '') => {
  const match = url.trim().match(/^\/video\/([^/?#]+)/)
  return match?.[1] || ''
}

export const isVideoPagePath = (url = '') => Boolean(videoSlugFromPath(url))

export const legacyVideoPagePath = (routeName = '') => {
  const slug = legacyVideoRouteSlugs[routeName.trim()]
  return slug ? videoPagePath(slug) : ''
}

export const resolveVideoPagePath = (link: VideoLinkTarget) => {
  const url = (link.url || '').trim()
  if (isVideoPagePath(url)) return url
  return legacyVideoPagePath(link.routeName || '')
}

export const resolveExternalVideoURL = (link: VideoLinkTarget) => {
  const url = (link.url || '').trim()
  return url && !isVideoPagePath(url) ? url : ''
}
