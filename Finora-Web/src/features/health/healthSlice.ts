import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { getHealthStatus } from './healthApi'

export type HealthState = {
  status: string
  version: string
  loading: boolean
  error: string | null
}

const initialState: HealthState = {
  status: 'Unknown',
  version: '-',
  loading: false,
  error: null,
}

export const fetchHealth = createAsyncThunk('health/fetchHealth', async (_, { rejectWithValue }) => {
  try {
    return await getHealthStatus()
  } catch {
    return rejectWithValue('Unable to reach health endpoint.')
  }
})

const healthSlice = createSlice({
  name: 'health',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchHealth.pending, (state) => {
        state.loading = true
        state.error = null
      })
      .addCase(fetchHealth.fulfilled, (state, action) => {
        state.loading = false
        state.status = action.payload.status
        state.version = action.payload.version
      })
      .addCase(fetchHealth.rejected, (state, action) => {
        state.loading = false
        state.error = (action.payload as string) ?? 'Health check failed.'
      })
  },
})

export const healthReducer = healthSlice.reducer
