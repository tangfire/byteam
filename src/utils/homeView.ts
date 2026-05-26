export const createParticleStyle = () => {
  const size = Math.random() * 4 + 2
  return {
    width: `${size}px`,
    height: `${size}px`,
    animationDelay: `${Math.random() * 5}s`,
    animationDuration: `${Math.random() * 8 + 8}s`,
    left: `${Math.random() * 100}%`,
    top: `${Math.random() * 100}%`,
  }
}

export const getNewsTypeClass = (type: string) => {
  const classMap: Record<string, string> = {
    publication: 'type-publication',
    team: 'type-team',
    award: 'type-award',
    event: 'type-event',
  }
  return classMap[type] || 'type-general'
}

export const getDay = (dateString: string) => new Date(dateString).getDate()

export const getMonth = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleString('en-US', { month: 'short' })
}

export const getYear = (dateString: string) => new Date(dateString).getFullYear()

export const scrollToResearch = () => {
  document.getElementById('research-highlights')?.scrollIntoView({ behavior: 'smooth' })
}
