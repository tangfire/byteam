import type { PublicationLink } from '../api/client'
import { resolveVideoPagePath } from './videoLinks'

export const linkTypeOptions = [
  { label: '论文 / PDF', shortLabel: '论文', value: 'paper', defaultLabel: 'Paper' },
  { label: '代码', shortLabel: '代码', value: 'code', defaultLabel: 'Code' },
  { label: '视频', shortLabel: '视频', value: 'video', defaultLabel: 'Video' },
  { label: 'PPT', shortLabel: 'PPT', value: 'ppt', defaultLabel: 'PPT' },
  { label: 'Poster', shortLabel: 'Poster', value: 'poster', defaultLabel: 'Poster' },
] as const

export const defaultLinkLabel = (type: string) => {
  return linkTypeOptions.find((option) => option.value === type)?.defaultLabel || 'Link'
}

export const defaultPublicationLinkLabels = linkTypeOptions.map((option) => option.defaultLabel)

const toRecord = (value: unknown): Record<string, any> => {
  return Boolean(value) && typeof value === 'object' && !Array.isArray(value) ? value as Record<string, any> : {}
}

const normalizeLinkText = (value: unknown) => typeof value === 'string' ? value : ''

export const normalizePublicationLinks = (value: unknown): PublicationLink[] => {
  if (!Array.isArray(value)) return []

  return value.map((link, index) => {
    const record = toRecord(link)
    const type = normalizeLinkText(record.type) || 'paper'
    const url = normalizeLinkText(record.url)
    const routeName = normalizeLinkText(record.routeName)
    const videoPageURL = type === 'video' && !url && routeName ? resolveVideoPagePath({ url, routeName }) : ''

    return {
      ...record,
      type,
      label: normalizeLinkText(record.label) || defaultLinkLabel(type),
      url: videoPageURL || url,
      routeName: '',
      sortOrder: index + 1,
    } as PublicationLink
  })
}

export const linkMediaKind = (link: PublicationLink) => {
  if (link.type === 'video') return 'video'
  if (link.type === 'paper' || link.type === 'ppt') return 'document'
  return ''
}

export const linkCanUseMedia = (link: PublicationLink) => {
  return link.type === 'paper' || link.type === 'ppt' || link.type === 'poster' || link.type === 'video'
}

export const linkUploadAccept = (link: PublicationLink) => {
  if (link.type === 'paper' || link.type === 'poster') return '.pdf,application/pdf,image/*'
  if (link.type === 'ppt') return '.ppt,.pptx,application/vnd.ms-powerpoint,application/vnd.openxmlformats-officedocument.presentationml.presentation'
  if (link.type === 'video') return '.mp4,video/mp4'
  return ''
}

export const linkURLPlaceholder = (link: PublicationLink) => {
  if (link.type === 'code') return 'GitHub、项目主页或其他外部链接'
  if (link.type === 'video') return '视频文件 URL 或外部链接'
  return '可粘贴外部链接，也可选择/上传文件'
}
