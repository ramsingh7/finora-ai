import { Navigate, useLocation } from 'react-router-dom'
import { useAppSelector } from '../../store/hooks'

type ProtectedRouteProps = {
  children: React.ReactNode
}

export function ProtectedRoute({ children }: ProtectedRouteProps) {
  const isAuthenticated = useAppSelector((state) => state.auth.isAuthenticated)
  const location = useLocation()

  if (!isAuthenticated) {
    return <Navigate to='/login' replace state={{ from: location }} />
  }

  return children
}
