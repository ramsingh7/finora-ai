import { httpClient } from '../../lib/httpClient'
import type { LoginRequest, LoginResponse } from './authTypes'

export async function loginRequest(payload: LoginRequest): Promise<LoginResponse> {
  const { data } = await httpClient.post<LoginResponse>('/api/auth/login', payload)
  return data
}
