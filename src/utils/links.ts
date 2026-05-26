export const openExternalLink = (url?: string) => {
  const target = url?.trim()
  if (!target) return
  window.open(target, '_blank', 'noopener,noreferrer')
}
