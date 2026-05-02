export type LoginRequest = {
  email: string
  password: string
}

export type LoginResponse = {
  ok: boolean
  expiresInSeconds: number
  token?: string
}
