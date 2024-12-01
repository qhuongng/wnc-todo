import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";

export interface User {
  username: string | null
}

const initialState: User = {
  username: null
};

export const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    setUsername: (state, action: PayloadAction<{ data: string }>) => {
      state.username = action.payload.data;
    },
  },
});

export const { setUsername } = userSlice.actions;
export const userReducer = userSlice.reducer;