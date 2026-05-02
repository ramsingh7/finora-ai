import { Bell, Gauge, Lock } from 'lucide-react'
import { Card } from '../../components/ui/Card'

const settings = [
  { label: 'Prediction alert notifications', description: 'Get alerts when confidence shifts rapidly.', icon: Bell },
  { label: 'Risk profile controls', description: 'Set guardrails for aggressive and conservative modes.', icon: Gauge },
  { label: 'Session security', description: 'Manage auth timeout and hard re-authentication rules.', icon: Lock },
]

export function SettingsPage() {
  return (
    <Card title='System Settings' description='Operational controls for AI-driven trading workflows'>
      <div className='space-y-3'>
        {settings.map((item) => {
          const Icon = item.icon
          return (
            <div
              key={item.label}
              className='flex items-start justify-between gap-4 rounded-md border border-slate-700 bg-slate-900/60 p-3'
            >
              <div>
                <p className='text-sm font-medium text-slate-100'>{item.label}</p>
                <p className='mt-1 text-xs text-slate-400'>{item.description}</p>
              </div>
              <Icon size={16} className='mt-0.5 text-cyan-300' />
            </div>
          )
        })}
      </div>
    </Card>
  )
}
