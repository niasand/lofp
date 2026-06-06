import { useEffect, useState } from 'react'
import { t } from '../i18n'

interface Props {
  onBack: () => void
}

export default function VerifyEmail({ onBack }: Props) {
  const [status, setStatus] = useState<'verifying' | 'success' | 'error'>('verifying')
  const [error, setError] = useState('')

  useEffect(() => {
    const params = new URLSearchParams(window.location.search)
    const token = params.get('token')
    if (!token) {
      setStatus('error')
      setError('Missing verification token.')
      return
    }
    fetch('/api/auth/verify-email', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ token }),
    }).then(async r => {
      if (r.ok) {
        setStatus('success')
      } else {
        const data = await r.json().catch(() => null)
        setStatus('error')
        setError(data?.error || 'Verification failed.')
      }
    }).catch(() => {
      setStatus('error')
      setError('Network error.')
    })
  }, [])

  return (
    <div className="flex items-center justify-center h-full p-8">
      <div className="max-w-md w-full text-center">
        {status === 'verifying' && (
          <p className="text-gray-400 font-mono">{t("verify.verifying")}</p>
        )}
        {status === 'success' && (
          <div>
            <p className="text-green-400 font-mono text-lg mb-4">{t("verify.success")}</p>
            <p className="text-gray-400 font-mono text-sm mb-6">{t("verify.successDesc")}</p>
            <button onClick={onBack} className="px-6 py-2 bg-amber-700 hover:bg-amber-600 text-white font-mono rounded">
              Continue to Login
            </button>
          </div>
        )}
        {status === 'error' && (
          <div>
            <p className="text-red-400 font-mono text-lg mb-4">{t("verify.failed")}</p>
            <p className="text-gray-400 font-mono text-sm mb-6">{error}</p>
            <button onClick={onBack} className="px-6 py-2 bg-[#333] hover:bg-[#444] text-gray-300 font-mono rounded">
              {t("main.back")}
            </button>
          </div>
        )}
      </div>
    </div>
  )
}
