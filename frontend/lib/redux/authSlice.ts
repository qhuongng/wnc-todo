import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";

export interface authState {
    accessToken: string
}

const initialState: authState = {
    accessToken: ""
};

export const authSlice = createSlice({
    name: "auth",
    initialState,
    reducers: {
        setAuthUser: (state, action: PayloadAction<{ accessToken: string }>) => {
            state.accessToken = action.payload.accessToken;
          },
    },
});

export const { setAuthUser } = authSlice.actions;
export const authReducer = authSlice.reducer;
