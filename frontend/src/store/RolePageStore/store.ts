import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import {devtools} from "zustand/middleware";
import React from "react";

interface User {
    avatar: string,
    name: string
}

export interface Permission {
    id: string,
    name: string,
    description: string,
    state: boolean
}

export interface Role {
    id: number,
    name: string,
    organization_id: number
    permissions: Permission[],
    parent_role_id: number,
}

interface State {
    roles: Role[]

    setRoles: (roles: Role[]) => void
    addRole: (role: Role) => void,
    deleteRole: (id: number) => void,
    setPermissionsInRole: (idRole: number, value: Permission[]) => void
}

const useRoleStore = create<State>()(
    devtools(
        immer((set) => ({
            roles: [],


            setRoles: (roles) => set((state) => {
               state.roles = roles;
            }),
            addRole: (role: Role) => set((state) => {
                state.roles.push(role);
            }),
            deleteRole: (id: number) => set((state) => {
                state.roles = state.roles.filter(role => role.id !== id);
            }),

            setPermissionsInRole: (idRole: number, value: Permission[]) => set((state) => {
                const role = state.roles.find(role => role.id === idRole);
                if (role) {
                    role.permissions = value;
                }
            }),
        }))
    )
);

export default useRoleStore;