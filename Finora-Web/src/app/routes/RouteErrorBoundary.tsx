import { isRouteErrorResponse, useRouteError } from 'react-router-dom'

export function RouteErrorBoundary() {
  const error = useRouteError()

  if (isRouteErrorResponse(error)) {
    return (
      <div className='flex min-h-screen items-center justify-center p-6'>
        <div className='max-w-lg rounded-lg border border-red-200 bg-white p-6'>
          <h1 className='text-xl font-semibold text-red-700'>Something went wrong</h1>
          <p className='mt-2 text-sm text-slate-600'>
            {error.status} {error.statusText}
          </p>
        </div>
      </div>
    )
  }

  return (
    <div className='flex min-h-screen items-center justify-center p-6'>
      <div className='max-w-lg rounded-lg border border-red-200 bg-white p-6'>
        <h1 className='text-xl font-semibold text-red-700'>Unexpected application error</h1>
        <p className='mt-2 text-sm text-slate-600'>Please refresh and try again.</p>
      </div>
    </div>
  )
}
