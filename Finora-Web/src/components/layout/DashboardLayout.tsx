import { motion } from 'framer-motion'
import {
  Activity,
  BrainCircuit,
  Database,
  FileText,
  Home,
  LineChart,
  Radar,
  Settings,
  ShieldCheck,
  Users,
  Workflow,
} from 'lucide-react'
import { NavLink, Outlet, useNavigate } from 'react-router-dom'
import { logout } from '../../features/auth/authSlice'
import { useAppDispatch } from '../../store/hooks'
import { AnimatedBackdrop } from '../ui/AnimatedBackdrop'

const navItems = [
  { to: '/dashboard', label: 'Overview', icon: Home },
  { to: '/dashboard/system-health', label: 'System Health', icon: Activity },
  { to: '/dashboard/profile', label: 'Profile', icon: ShieldCheck },
  { to: '/dashboard/settings', label: 'Settings', icon: Settings },
  { to: '/dashboard/users', label: 'Users', icon: Users },
  { to: '/dashboard/roles', label: 'Roles', icon: Radar },
  { to: '/dashboard/market-data', label: 'Market Data', icon: LineChart },
  { to: '/dashboard/models', label: 'Models', icon: BrainCircuit },
  { to: '/dashboard/predictions', label: 'Predictions', icon: Database },
  { to: '/dashboard/jobs', label: 'Jobs', icon: Workflow },
  { to: '/dashboard/logs', label: 'Logs', icon: FileText },
]

export function DashboardLayout() {
  const dispatch = useAppDispatch()
  const navigate = useNavigate()

  const handleLogout = () => {
    dispatch(logout())
    navigate('/login', { replace: true })
  }

  return (
    <div className='relative flex min-h-screen overflow-hidden bg-slate-950 text-slate-100'>
      <AnimatedBackdrop />
      <aside className='glass-panel relative z-10 hidden w-72 border-r border-slate-700/60 lg:block'>
        <div className='border-b border-slate-700/60 px-6 py-5'>
          <p className='text-xs uppercase tracking-widest text-slate-400'>Finora</p>
          <h1 className='mt-1 text-xl font-semibold'>AI Trading Admin</h1>
          <p className='mt-1 text-xs text-slate-400'>Enterprise Operations</p>
        </div>
        <nav className='space-y-1 p-4'>
          {navItems.map((item) => {
            const Icon = item.icon
            return (
              <NavLink
                key={item.to}
                to={item.to}
                end={item.to === '/dashboard'}
                className={({ isActive }) =>
                  `group flex items-center gap-3 rounded-md px-3 py-2 text-sm transition ${
                    isActive
                      ? 'bg-gradient-to-r from-indigo-500/80 to-cyan-500/80 text-white shadow-lg shadow-cyan-500/10'
                      : 'text-slate-300 hover:bg-slate-800/80 hover:text-white'
                  }`
                }
              >
                <Icon size={16} className='opacity-90 transition group-hover:scale-105' />
                <span>{item.label}</span>
              </NavLink>
            )
          })}
        </nav>
      </aside>
      <div className='relative z-10 flex min-h-screen flex-1 flex-col'>
        <header className='glass-panel flex items-center justify-between border-b border-slate-700/60 px-6 py-4'>
          <div>
            <p className='text-xs uppercase tracking-wider text-cyan-300'>Autonomous market operations</p>
            <p className='text-sm text-slate-300'>AI agents, prediction models, and system telemetry</p>
          </div>
          <div className='flex items-center gap-3'>
            <span className='rounded-full border border-emerald-400/40 bg-emerald-500/15 px-3 py-1 text-xs text-emerald-300'>
              All systems active
            </span>
            <button
              type='button'
              onClick={handleLogout}
              className='rounded-md border border-slate-500 bg-slate-900/60 px-4 py-2 text-sm font-medium text-white hover:bg-slate-800'
            >
              Logout
            </button>
          </div>
        </header>
        <motion.main
          key='dashboard-content'
          initial={{ opacity: 0, y: 12 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.35 }}
          className='flex-1 p-6'
        >
          <Outlet />
        </motion.main>
      </div>
    </div>
  )
}
