/**
 * 格式化日期 - 获取日
 */
export const getDay = (dateString: string): number => {
  return new Date(dateString).getDate()
}

/**
 * 格式化日期 - 获取月份缩写
 */
export const getMonth = (dateString: string): string => {
  const date = new Date(dateString)
  return date.toLocaleString('en-US', { month: 'short' })
}

/**
 * 格式化日期 - 获取年份
 */
export const getYear = (dateString: string): number => {
  return new Date(dateString).getFullYear()
}

/**
 * 格式化完整日期
 */
export const formatDate = (dateString: string, locale: string = 'en-US'): string => {
  const date = new Date(dateString)
  return date.toLocaleDateString(locale, {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

/**
 * 平滑滚动到指定元素
 */
export const scrollToElement = (elementId: string, behavior: ScrollBehavior = 'smooth'): void => {
  const element = document.getElementById(elementId)
  if (element) {
    element.scrollIntoView({ behavior })
  }
}

/**
 * 检查是否为移动端
 */
export const isMobile = (): boolean => {
  const viewportWidth = Math.max(document.documentElement.clientWidth || 0, window.innerWidth || 0)
  const isTouchDevice = 'ontouchstart' in window || navigator.maxTouchPoints > 0
  return viewportWidth <= 768 && isTouchDevice
}

/**
 * 防抖函数
 */
export const debounce = <T extends (...args: any[]) => any>(
  func: T,
  wait: number
): ((...args: Parameters<T>) => void) => {
  let timeout: ReturnType<typeof setTimeout> | null = null
  
  return function executedFunction(...args: Parameters<T>) {
    const later = () => {
      timeout = null
      func(...args)
    }
    
    if (timeout) {
      clearTimeout(timeout)
    }
    timeout = setTimeout(later, wait)
  }
}

/**
 * 节流函数
 */
export const throttle = <T extends (...args: any[]) => any>(
  func: T,
  limit: number
): ((...args: Parameters<T>) => void) => {
  let inThrottle: boolean = false
  
  return function executedFunction(...args: Parameters<T>) {
    if (!inThrottle) {
      func(...args)
      inThrottle = true
      setTimeout(() => (inThrottle = false), limit)
    }
  }
}
