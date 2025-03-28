import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import {devtools} from "zustand/middleware";
import React from "react";

interface State {
    isLogin: boolean;
    refreshToken: string;

    setLogin: (isLogin: boolean) => void;
    setRefreshToken: (token: string) => void;
}

const useGlobalStore = create<State>()(
    devtools(
        immer((set) => ({
            isLogin: false,
            refreshToken: "",

            setLogin: (isLogin: boolean) => set((state) => {
                state.isLogin = isLogin;
            }),
            setRefreshToken: (token: string) => set((state) => {
                state.refreshToken = token;
            }),
        }))
    )
);

export default useGlobalStore;