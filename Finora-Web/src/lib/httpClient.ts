import axios from 'axios'

const timeoutMs = Number(import.meta.env.VITE_API_TIMEOUT_MS ?? 10000)

export const httpClient = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL ?? '',
  timeout: Number.isNaN(timeoutMs) ? 10000 : timeoutMs,
  withCredentials: true,
})
