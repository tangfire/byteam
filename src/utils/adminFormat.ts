export const formatStatus = (status = '') => status === 'published' ? '已发布' : '草稿'

export const formatMediaKind = (kind = '') => {
  const labels: Record<string, string> = {
    image: '图片',
    document: '文档',
    archive: '压缩包',
    video: '视频',
  }
  return labels[kind] || kind
}
