import { motion } from 'framer-motion'
import { useEffect } from 'react'
import { Card } from '../../components/ui/Card'
import { useAppDispatch, useAppSelector } from '../../store/hooks'
import { fetchHealth } from './healthSlice'

export function SystemHealthPage() {
  const dispatch = useAppDispatch()
  const { status, version, loading, error } = useAppSelector((state) => state.health)

  useEffect(() => {
    void dispatch(fetchHealth())
  }, [dispatch])

  return (
    <Card title='System Health' description='Live signal from Finora backend health endpoint.'>
      {loading ? <p className='text-sm text-slate-300'>Checking backend status...</p> : null}
      {error ? <p className='rounded-md bg-red-500/15 p-2 text-sm text-red-300'>{error}</p> : null}
      {!loading && !error ? (
        <dl className='grid grid-cols-1 gap-3 text-sm sm:grid-cols-2'>
          <motion.div
            initial={{ scale: 0.96, opacity: 0.7 }}
            animate={{ scale: 1, opacity: 1 }}
            transition={{ duration: 0.35 }}
            className='rounded-md border border-slate-700 bg-slate-900/70 p-3'
          >
            <dt className='text-slate-400'>Status</dt>
            <dd className='mt-1 font-semibold text-emerald-300'>{status}</dd>
          </motion.div>
          <motion.div
            initial={{ scale: 0.96, opacity: 0.7 }}
            animate={{ scale: 1, opacity: 1 }}
            transition={{ duration: 0.35, delay: 0.1 }}
            className='rounded-md border border-slate-700 bg-slate-900/70 p-3'
          >
            <dt className='text-slate-400'>Version</dt>
            <dd className='mt-1 font-semibold text-cyan-300'>{version}</dd>
          </motion.div>
        </dl>
      ) : null}
    </Card>
  )
}
