import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { loginRequest } from './authApi'
import type { LoginRequest } from './authTypes'

export type AuthState = {
  isAuthenticated: boolean
  loading: boolean
  error: string | null
  sessionExpiresAt: number | null
  token: string | null
}

const initialState: AuthState = {
  isAuthenticated: false,
  loading: false,
  error: null,
  sessionExpiresAt: null,
  token: null,
}

export const login = createAsyncThunk(
  'auth/login',
  async (payload: LoginRequest, { rejectWithValue }) => {
    try {
      const response = await loginRequest(payload)
      return response
    } catch {
      return rejectWithValue('Invalid credentials or service unavailable.')
    }
  },
)

const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    logout: (state) => {
      state.isAuthenticated = false
      state.sessionExpiresAt = null
      state.token = null
      state.error = null
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(login.pending, (state) => {
        state.loading = true
        state.error = null
      })
      .addCase(login.fulfilled, (state, action) => {
        state.loading = false
        state.isAuthenticated = true
        state.token = action.payload.access_token
        state.sessionExpiresAt = Date.now() + action.payload.expires_in_seconds * 1000
      })
      .addCase(login.rejected, (state, action) => {
        state.loading = false
        state.error = (action.payload as string) ?? 'Login failed.'
      })
  },
})

export const { logout } = authSlice.actions
export const authReducer = authSlice.reducer
