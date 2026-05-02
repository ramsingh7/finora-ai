import { motion } from 'framer-motion'

type CardProps = {
  title: string
  description?: string
  children: React.ReactNode
}

export function Card({ title, description, children }: CardProps) {
  return (
    <motion.section
      initial={{ opacity: 0, y: 14 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.4 }}
      className='glass-panel rounded-xl p-5'
    >
      <header className='mb-4'>
        <h2 className='text-base font-semibold text-white'>{title}</h2>
        {description ? <p className='mt-1 text-sm text-slate-400'>{description}</p> : null}
      </header>
      {children}
    </motion.section>
  )
}
