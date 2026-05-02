import { motion } from 'framer-motion'
import { ArrowUpRight, Bot, Sparkles, TrendingUp } from 'lucide-react'
import { Link } from 'react-router-dom'
import { Card } from '../../components/ui/Card'

const metrics = [
  { label: 'Signal Confidence', value: '94.2%', delta: '+3.1%', tone: 'text-emerald-300', icon: Sparkles },
  { label: 'Active AI Agents', value: '12', delta: '+2', tone: 'text-cyan-300', icon: Bot },
  { label: 'Predictions / Hour', value: '1,840', delta: '+11%', tone: 'text-indigo-300', icon: TrendingUp },
]

const marketSignals = [
  { symbol: 'AAPL', trend: 'Bullish', confidence: '93%', model: 'Temporal Fusion v2' },
  { symbol: 'MSFT', trend: 'Bullish', confidence: '88%', model: 'LSTM Ensemble v4' },
  { symbol: 'NVDA', trend: 'Volatile', confidence: '76%', model: 'Regime Detector' },
  { symbol: 'TSLA', trend: 'Bearish', confidence: '81%', model: 'Event Impact Agent' },
]

const activities = [
  'Agent #07 retrained on latest earnings data',
  'Risk guardrail auto-adjusted stop-loss thresholds',
  'Portfolio simulation completed for US Tech basket',
  'Anomaly detector flagged sudden volatility spike',
]

export function DashboardHomePage() {
  return (
    <div className='space-y-6'>
      <Card title='Finora AI Control Center' description='Real-time autonomous stock prediction operations'>
        <div className='flex flex-wrap items-center justify-between gap-4'>
          <p className='max-w-2xl text-sm text-slate-300'>
            Monitor model intelligence, agent execution, and market signal quality from one command center.
          </p>
          <Link
            to='/dashboard/predictions'
            className='inline-flex items-center gap-2 rounded-md border border-cyan-400/40 bg-cyan-400/10 px-3 py-2 text-sm text-cyan-200 hover:bg-cyan-400/20'
          >
            Open live predictions
            <ArrowUpRight size={14} />
          </Link>
        </div>
      </Card>

      <div className='grid grid-cols-1 gap-4 md:grid-cols-3'>
        {metrics.map((metric, index) => {
          const Icon = metric.icon
          return (
            <motion.div
              key={metric.label}
              initial={{ opacity: 0, y: 12 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.35, delay: index * 0.08 }}
              className='glass-panel rounded-xl p-4'
            >
              <div className='flex items-center justify-between'>
                <p className='text-xs uppercase tracking-wider text-slate-400'>{metric.label}</p>
                <Icon size={16} className='text-slate-300' />
              </div>
              <p className={`mt-2 text-2xl font-semibold ${metric.tone}`}>{metric.value}</p>
              <p className='mt-1 text-xs text-emerald-300'>{metric.delta} vs last 24h</p>
            </motion.div>
          )
        })}
      </div>

      <div className='grid grid-cols-1 gap-6 xl:grid-cols-3'>
        <Card title='Top AI Signals' description='Highest-confidence predictions currently active'>
          <div className='overflow-hidden rounded-lg border border-slate-700/80'>
            <table className='w-full text-left text-sm'>
              <thead className='bg-slate-900/80 text-xs uppercase tracking-wider text-slate-400'>
                <tr>
                  <th className='px-3 py-2'>Symbol</th>
                  <th className='px-3 py-2'>Trend</th>
                  <th className='px-3 py-2'>Confidence</th>
                  <th className='px-3 py-2'>Model</th>
                </tr>
              </thead>
              <tbody>
                {marketSignals.map((row) => (
                  <tr key={row.symbol} className='border-t border-slate-800/80 hover:bg-slate-800/40'>
                    <td className='px-3 py-2 font-medium text-slate-100'>{row.symbol}</td>
                    <td className='px-3 py-2 text-slate-300'>{row.trend}</td>
                    <td className='px-3 py-2 text-cyan-300'>{row.confidence}</td>
                    <td className='px-3 py-2 text-slate-400'>{row.model}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </Card>

        <Card title='Recent Activity' description='Automated events and system actions'>
          <div className='space-y-3'>
            {activities.map((item) => (
              <div key={item} className='rounded-md border border-slate-700 bg-slate-900/50 px-3 py-2 text-sm text-slate-300'>
                {item}
              </div>
            ))}
          </div>
        </Card>

        <Card title='Quick Actions'>
          <div className='flex flex-col gap-3'>
            <Link
              to='/dashboard/system-health'
              className='rounded-md bg-gradient-to-r from-indigo-500 via-cyan-500 to-emerald-500 px-4 py-2 text-center text-sm font-medium text-white'
            >
              View system health
            </Link>
            <Link
              to='/dashboard/models'
              className='rounded-md border border-slate-600 bg-slate-900/70 px-4 py-2 text-center text-sm font-medium text-slate-100 hover:bg-slate-800'
            >
              Manage model registry
            </Link>
            <Link
              to='/dashboard/jobs'
              className='rounded-md border border-slate-600 bg-slate-900/70 px-4 py-2 text-center text-sm font-medium text-slate-100 hover:bg-slate-800'
            >
              Monitor job pipeline
            </Link>
          </div>
        </Card>
      </div>
    </div>
  )
}
