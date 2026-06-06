import en from './en.json'
import zh from './zh.json'

const locale = import.meta.env.VITE_LOCALE || 'en'
const messages: Record<string, string> = locale === 'zh' ? zh : en

export function t(key: string): string {
  return messages[key] || key
}

/** Translate a race ID (1-8) to its localized name. */
export function raceName(id: number): string {
  const keys: Record<number, string> = {
    1: 'race.human', 2: 'race.aelfen', 3: 'race.highlander', 4: 'race.wolfling',
    5: 'race.murg', 6: 'race.drakin', 7: 'race.mechanoid', 8: 'race.ephemeral',
  }
  return t(keys[id]) || String(id)
}
