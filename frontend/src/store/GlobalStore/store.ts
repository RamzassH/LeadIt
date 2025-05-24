import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import {devtools} from "zustand/middleware";
import React from "react";
import {Organization} from "@/store/OrganizationStore/store";

interface State {
    isLogin: boolean;
    refreshToken: string;

    currentOrganization: Organization | null;

    setLogin: (isLogin: boolean) => void;
    setRefreshToken: (token: string) => void;
    setCurrentOrganization: (organization: Organization | null) => void;
}

const useGlobalStore = create<State>()(
    devtools(
        immer((set) => ({
            isLogin: false,
            refreshToken: "",
            currentOrganization: null,

            setLogin: (isLogin: boolean) => set((state) => {
                state.isLogin = isLogin;
            }),
            setRefreshToken: (token: string) => set((state) => {
                state.refreshToken = token;
            }),
            setCurrentOrganization: (organization: Organization | null) => set((state) => {
                state.currentOrganization = organization
            })
        }))
    )
);

export default useGlobalStore;