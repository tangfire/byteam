import type { CSSProperties } from 'vue'

export interface AvatarDisplayRecord {
  avatarObjectX?: number | null
  avatarObjectY?: number | null
  avatarScale?: number | null
}

const clampNumber = (value: unknown, min: number, max: number, fallback: number) => {
  const numberValue = Number(value)
  if (!Number.isFinite(numberValue)) return fallback
  return Math.min(max, Math.max(min, Math.round(numberValue)))
}

export const avatarDisplayStyle = (record: AvatarDisplayRecord): CSSProperties => {
  const x = clampNumber(record.avatarObjectX, 0, 100, 50)
  const y = clampNumber(record.avatarObjectY, 0, 100, 50)
  const scale = clampNumber(record.avatarScale, 100, 200, 100) / 100
  return {
    objectPosition: `${x}% ${y}%`,
    transform: `scale(${scale})`,
  }
}
