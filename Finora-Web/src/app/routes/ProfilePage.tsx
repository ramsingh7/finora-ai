import { Mail, Shield, User } from 'lucide-react'
import { Card } from '../../components/ui/Card'

export function ProfilePage() {
  return (
    <Card title='Admin Profile' description='Identity and access details for secure operations'>
      <div className='grid gap-3 sm:grid-cols-2'>
        <div className='rounded-md border border-slate-700 bg-slate-900/60 p-3'>
          <p className='mb-2 flex items-center gap-2 text-xs uppercase tracking-wider text-slate-400'>
            <User size={14} /> Name
          </p>
          <p className='text-sm text-slate-100'>Finora Administrator</p>
        </div>
        <div className='rounded-md border border-slate-700 bg-slate-900/60 p-3'>
          <p className='mb-2 flex items-center gap-2 text-xs uppercase tracking-wider text-slate-400'>
            <Mail size={14} /> Email
          </p>
          <p className='text-sm text-slate-100'>admin@finora.ai</p>
        </div>
        <div className='rounded-md border border-slate-700 bg-slate-900/60 p-3 sm:col-span-2'>
          <p className='mb-2 flex items-center gap-2 text-xs uppercase tracking-wider text-slate-400'>
            <Shield size={14} /> Role & Access
          </p>
          <p className='text-sm text-slate-100'>Super Admin - full access to AI agents, models, jobs and system controls.</p>
        </div>
      </div>
    </Card>
  )
}
