import { useState } from 'react'
import { t, raceName } from '../i18n'

function getRaces() {
  return [
    { id: 1, name: 'Human', stats: t("character.race.human.stats"), desc: t("character.race.human.desc"), ability: t("character.race.human.ability") },
    { id: 2, name: 'Aelfen', stats: t("character.race.aelfen.stats"), desc: t("character.race.aelfen.desc"), ability: t("character.race.aelfen.ability") },
    { id: 3, name: 'Highlander', stats: t("character.race.highlander.stats"), desc: t("character.race.highlander.desc"), ability: t("character.race.highlander.ability") },
    { id: 4, name: 'Wolfling', stats: t("character.race.wolfling.stats"), desc: t("character.race.wolfling.desc"), ability: t("character.race.wolfling.ability") },
    { id: 5, name: 'Murg', stats: t("character.race.murg.stats"), desc: t("character.race.murg.desc"), ability: t("character.race.murg.ability") },
    { id: 6, name: 'Drakin', stats: t("character.race.drakin.stats"), desc: t("character.race.drakin.desc"), ability: t("character.race.drakin.ability") },
    { id: 7, name: 'Mechanoid', stats: t("character.race.mechanoid.stats"), desc: t("character.race.mechanoid.desc"), ability: t("character.race.mechanoid.ability") },
    { id: 8, name: 'Ephemeral', stats: t("character.race.ephemeral.stats"), desc: t("character.race.ephemeral.desc"), ability: t("character.race.ephemeral.ability") },
  ]
}

interface Props {
  onCreated: (char: { firstName: string; lastName: string; race: number; gender: number }) => void
  onOpenManual?: () => void
}

export default function CharacterCreate({ onCreated, onOpenManual }: Props) {
  const [firstName, setFirstName] = useState('')
  const [lastName, setLastName] = useState('')
  const [race, setRace] = useState(1)
  const [gender, setGender] = useState(0)
  const RACES = getRaces()
  const selectedRace = RACES.find(r => r.id === race) ?? RACES[0]

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (!firstName.trim() || !lastName.trim()) return
    onCreated({ firstName: firstName.trim(), lastName: lastName.trim(), race, gender })
  }

  return (
    <div className="flex items-start sm:items-center justify-center h-full pt-4 px-4 pb-4 sm:p-8 overflow-y-auto">
      <div className="max-w-3xl w-full bg-[#111] border border-[#333] rounded-lg p-4 sm:p-8">
        <h2 className="text-amber-400 text-2xl font-mono mb-1 text-center">
          {t("character.title")}
        </h2>
        <p className="text-gray-500 text-sm font-mono mb-6 text-center">
          {t("character.subtitle")}
        </p>

        <form onSubmit={handleSubmit} className="space-y-6">
          {/* Name */}
          <div className="bg-[#0a0a0a] border border-amber-900/50 rounded-lg p-3 mb-2">
            <p className="text-gray-400 text-xs font-mono leading-relaxed">
              {t("character.nameNotice")}{' '}
              <button type="button" onClick={onOpenManual} className="text-amber-500 hover:text-amber-400 underline cursor-pointer">
                {t("character.readMoreRoleplaying")}
              </button>
            </p>
          </div>
          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div>
              <label className="block text-gray-400 text-sm font-mono mb-1">{t("character.firstName")}</label>
              <input
                type="text"
                value={firstName}
                onChange={e => setFirstName(e.target.value)}
                maxLength={15}
                className="w-full bg-[#0a0a0a] border border-[#444] rounded px-3 py-2 text-gray-200 font-mono focus:border-amber-500 focus:outline-none"
                placeholder="Balthazar"
                autoFocus
              />
            </div>
            <div>
              <label className="block text-gray-400 text-sm font-mono mb-1">{t("character.lastName")}</label>
              <input
                type="text"
                value={lastName}
                onChange={e => setLastName(e.target.value)}
                maxLength={15}
                className="w-full bg-[#0a0a0a] border border-[#444] rounded px-3 py-2 text-gray-200 font-mono focus:border-amber-500 focus:outline-none"
                placeholder="Thinvar"
              />
              <p className="text-gray-600 text-[10px] font-mono mt-1">{t("character.lastNameHint")}</p>
            </div>
          </div>

          {/* Gender */}
          <div>
            <label className="block text-gray-400 text-sm font-mono mb-1">{t("character.gender")}</label>
            <div className="flex gap-4">
              {[{ v: 0, l: t("character.male") }, { v: 1, l: t("character.female") }].map(g => (
                <button
                  key={g.v}
                  type="button"
                  onClick={() => setGender(g.v)}
                  className={`px-6 py-2.5 min-h-[44px] rounded font-mono text-sm transition-colors ${gender === g.v ? 'bg-amber-700 text-white' : 'bg-[#1a1a1a] text-gray-400 border border-[#444] hover:border-amber-600'}`}
                >
                  {g.l}
                </button>
              ))}
            </div>
          </div>

          {/* Race selection */}
          <div>
            <label className="block text-gray-400 text-sm font-mono mb-2">{t("character.race")}</label>
            <div className="grid grid-cols-2 sm:grid-cols-4 gap-2 mb-3">
              {RACES.map(r => (
                <button
                  key={r.id}
                  type="button"
                  onClick={() => setRace(r.id)}
                  className={`px-2 py-3 min-h-[44px] rounded font-mono text-sm transition-colors ${race === r.id ? 'bg-amber-700 text-white border border-amber-600' : 'bg-[#1a1a1a] text-gray-400 border border-[#444] hover:border-amber-600'}`}
                >
                  {raceName(r.id)}
                </button>
              ))}
            </div>

            {/* Race detail */}
            <div className="bg-[#0a0a0a] border border-[#333] rounded-lg p-4 space-y-3">
              <div className="flex items-center gap-3">
                <h3 className="text-amber-400 font-mono text-lg font-bold">{raceName(selectedRace.id)}</h3>
              </div>
              <p className="text-gray-300 font-mono text-xs leading-relaxed">{selectedRace.desc}</p>
              <div className="bg-[#111] border border-[#2a2a2a] rounded p-2">
                <p className="text-green-400 font-mono text-xs">
                  <span className="text-gray-500">{t("character.ability")}</span> {selectedRace.ability}
                </p>
              </div>
              <div className="bg-[#111] border border-[#2a2a2a] rounded p-2">
                <p className="text-cyan-400 font-mono text-[10px] tracking-wider">{selectedRace.stats}</p>
              </div>
            </div>
          </div>

          <button
            type="submit"
            disabled={!firstName.trim() || !lastName.trim()}
            className="w-full py-3 bg-amber-700 hover:bg-amber-600 disabled:bg-gray-700 disabled:text-gray-500 text-white font-mono rounded text-lg transition-colors cursor-pointer"
          >
            {t("character.enterWorld")}
          </button>
        </form>
      </div>
    </div>
  )
}
