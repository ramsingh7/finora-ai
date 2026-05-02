import { Navigate, createBrowserRouter } from 'react-router-dom'
import { LoginPage } from '../../features/auth/LoginPage'
import { DashboardLayout } from '../../components/layout/DashboardLayout'
import { DashboardHomePage } from '../routes/DashboardHomePage'
import { SystemHealthPage } from '../../features/health/SystemHealthPage'
import { ProfilePage } from '../routes/ProfilePage'
import { SettingsPage } from '../routes/SettingsPage'
import { PlaceholderPage } from '../routes/PlaceholderPage'
import { ProtectedRoute } from './ProtectedRoute'
import { RouteErrorBoundary } from './RouteErrorBoundary'

export const router = createBrowserRouter([
  {
    path: '/login',
    element: <LoginPage />,
    errorElement: <RouteErrorBoundary />,
  },
  {
    path: '/',
    element: <Navigate to='/dashboard' replace />,
  },
  {
    path: '/dashboard',
    element: (
      <ProtectedRoute>
        <DashboardLayout />
      </ProtectedRoute>
    ),
    errorElement: <RouteErrorBoundary />,
    children: [
      { index: true, element: <DashboardHomePage /> },
      { path: 'system-health', element: <SystemHealthPage /> },
      { path: 'profile', element: <ProfilePage /> },
      { path: 'settings', element: <SettingsPage /> },
      { path: 'users', element: <PlaceholderPage title='Users' /> },
      { path: 'roles', element: <PlaceholderPage title='Roles' /> },
      { path: 'market-data', element: <PlaceholderPage title='Market Data' /> },
      { path: 'models', element: <PlaceholderPage title='Models' /> },
      { path: 'predictions', element: <PlaceholderPage title='Predictions' /> },
      { path: 'jobs', element: <PlaceholderPage title='Jobs' /> },
      { path: 'logs', element: <PlaceholderPage title='Logs' /> },
    ],
  },
  {
    path: '*',
    element: <Navigate to='/dashboard' replace />,
  },
])
