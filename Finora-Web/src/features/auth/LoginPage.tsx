import { zodResolver } from '@hookform/resolvers/zod'
import { motion } from 'framer-motion'
import { useEffect } from 'react'
import { useForm } from 'react-hook-form'
import { useLocation, useNavigate } from 'react-router-dom'
import { z } from 'zod'
import { AnimatedBackdrop } from '../../components/ui/AnimatedBackdrop'
import { useAppDispatch, useAppSelector } from '../../store/hooks'
import { login } from './authSlice'

const loginSchema = z.object({
  email: z.email('Enter a valid email address.'),
  password: z.string().min(6, 'Password must be at least 6 characters.'),
})

type LoginFormValues = z.infer<typeof loginSchema>

export function LoginPage() {
  const dispatch = useAppDispatch()
  const navigate = useNavigate()
  const location = useLocation()
  const { loading, error, isAuthenticated } = useAppSelector((state) => state.auth)

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormValues>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: '',
      password: '',
    },
  })

  useEffect(() => {
    if (isAuthenticated) {
      const from = (location.state as { from?: { pathname: string } })?.from?.pathname
      navigate(from ?? '/dashboard', { replace: true })
    }
  }, [isAuthenticated, location.state, navigate])

  const onSubmit = (values: LoginFormValues) => {
    void dispatch(login(values))
  }

  return (
    <div className='relative flex min-h-screen items-center justify-center overflow-hidden px-4'>
      <AnimatedBackdrop />

      <motion.div
        initial={{ opacity: 0, y: 30 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.65, ease: 'easeOut' }}
        className='relative z-10 grid w-full max-w-6xl gap-8 lg:grid-cols-2'
      >
        <div className='glass-panel hidden rounded-2xl p-8 lg:block'>
          <p className='text-xs uppercase tracking-[0.28em] text-cyan-300'>AI market intelligence</p>
          <h1 className='mt-3 text-4xl font-semibold leading-tight text-white'>
            Predict markets with AI agents and actionable signals
          </h1>
          <p className='mt-5 text-sm text-slate-300'>
            Finora combines autonomous agents, model orchestration, and live market telemetry for faster decisions.
          </p>
          <div className='mt-8 space-y-3'>
            {['Multi-agent strategy execution', 'Real-time trend detection', 'Risk-aware prediction pipelines'].map((item) => (
              <div key={item} className='ticker-line rounded-lg border border-cyan-400/30 bg-slate-900/60 px-4 py-3 text-sm'>
                {item}
              </div>
            ))}
          </div>
        </div>

        <div className='glass-panel relative rounded-2xl p-8'>
          <p className='text-xs uppercase tracking-[0.2em] text-slate-300'>Finora Admin</p>
          <h2 className='mt-2 text-3xl font-semibold text-white'>Secure Sign in</h2>
          <p className='mt-1 text-sm text-slate-400'>Use your admin credentials to continue.</p>

          <form className='mt-6 space-y-4' onSubmit={handleSubmit(onSubmit)}>
            <div>
              <label htmlFor='email' className='mb-1 block text-sm font-medium text-slate-200'>
                Email
              </label>
              <input
                id='email'
                type='email'
                autoComplete='email'
                className='w-full rounded-md border border-slate-600 bg-slate-900/80 px-3 py-2 text-white outline-none ring-cyan-400 focus:ring-2'
                {...register('email')}
              />
              {errors.email ? <p className='mt-1 text-xs text-red-400'>{errors.email.message}</p> : null}
            </div>

            <div>
              <label htmlFor='password' className='mb-1 block text-sm font-medium text-slate-200'>
                Password
              </label>
              <input
                id='password'
                type='password'
                autoComplete='current-password'
                className='w-full rounded-md border border-slate-600 bg-slate-900/80 px-3 py-2 text-white outline-none ring-cyan-400 focus:ring-2'
                {...register('password')}
              />
              {errors.password ? <p className='mt-1 text-xs text-red-400'>{errors.password.message}</p> : null}
            </div>

            {error ? <p className='rounded-md bg-red-500/15 p-2 text-sm text-red-300'>{error}</p> : null}

            <button
              type='submit'
              disabled={loading}
              className='w-full rounded-md bg-gradient-to-r from-indigo-500 via-cyan-500 to-emerald-500 px-4 py-2 font-medium text-white transition hover:scale-[1.01] disabled:cursor-not-allowed disabled:opacity-70'
            >
              {loading ? 'Signing in...' : 'Sign in'}
            </button>
          </form>
        </div>
      </motion.div>
    </div>
  )
}
