import { Card } from '../../components/ui/Card'

type PlaceholderPageProps = {
  title: string
}

export function PlaceholderPage({ title }: PlaceholderPageProps) {
  return (
    <Card title={title} description='Backend APIs are pending for this module.'>
      <p className='text-sm text-slate-600'>This screen is scaffolded and ready for API integration.</p>
    </Card>
  )
}
