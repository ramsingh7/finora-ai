import { httpClient } from '../../lib/httpClient'
import type { HealthResponse } from './healthTypes'

export async function getHealthStatus(): Promise<HealthResponse> {
  const { data } = await httpClient.get<HealthResponse>('/api/health')
  return data
}
