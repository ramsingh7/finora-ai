import { motion } from 'framer-motion'
import { Link } from 'react-router-dom'
import { Card } from '../../components/ui/Card'

const metrics = [
  { label: 'Signal Confidence', value: '94.2%', tone: 'text-emerald-300' },
  { label: 'Active AI Agents', value: '12', tone: 'text-cyan-300' },
  { label: 'Predictions / Hour', value: '1,840', tone: 'text-indigo-300' },
]

export function DashboardHomePage() {
  return (
    <div className='space-y-6'>
      <Card title='Finora AI Control Center' description='Real-time autonomous stock prediction operations'>
        <p className='text-sm text-slate-300'>
          Monitor model intelligence, agent execution, and market signal quality from one command center.
        </p>
      </Card>

      <div className='grid grid-cols-1 gap-4 md:grid-cols-3'>
        {metrics.map((metric, index) => (
          <motion.div
            key={metric.label}
            initial={{ opacity: 0, y: 12 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.35, delay: index * 0.08 }}
            className='glass-panel rounded-xl p-4'
          >
            <p className='text-xs uppercase tracking-wider text-slate-400'>{metric.label}</p>
            <p className={`mt-2 text-2xl font-semibold ${metric.tone}`}>{metric.value}</p>
          </motion.div>
        ))}
      </div>

      <Card title='Quick Actions'>
        <div className='flex flex-wrap gap-3'>
          <Link
            to='/dashboard/system-health'
            className='rounded-md bg-gradient-to-r from-indigo-500 via-cyan-500 to-emerald-500 px-4 py-2 text-sm font-medium text-white'
          >
            View system health
          </Link>
          <Link
            to='/dashboard/predictions'
            className='rounded-md border border-slate-600 bg-slate-900/70 px-4 py-2 text-sm font-medium text-slate-100 hover:bg-slate-800'
          >
            Open predictions module
          </Link>
        </div>
      </Card>
    </div>
  )
}
