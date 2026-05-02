import { ArrowRight } from 'lucide-react'
import { Card } from '../../components/ui/Card'

type PlaceholderPageProps = {
  title: string
}

export function PlaceholderPage({ title }: PlaceholderPageProps) {
  return (
    <Card title={title} description='Module shell is production-ready and waiting for backend APIs.'>
      <div className='rounded-md border border-dashed border-slate-600 bg-slate-900/40 p-4'>
        <p className='text-sm text-slate-300'>
          API integration contracts for <span className='font-medium text-slate-100'>{title}</span> are pending. This page
          already includes layout, navigation and visual standards.
        </p>
        <button
          type='button'
          className='mt-3 inline-flex items-center gap-2 rounded-md border border-slate-600 px-3 py-2 text-sm text-slate-200 hover:bg-slate-800'
        >
          Connect API when ready
          <ArrowRight size={14} />
        </button>
      </div>
    </Card>
  )
}
