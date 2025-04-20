import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import {devtools} from "zustand/middleware";
import React from "react";

interface User {
    avatar: string,
    name: string
}

interface Permission {
    id: string,
    name: string,
    description: string,
    state: boolean
}

interface Role {
    id: string,
    name: string,
    permission: Permission[],
    users: User[]
}

interface State {
    roles: Role[]

    addRole: (role: Role) => void,
    deleteRole: (id: string) => void,
    addUserInRole: (idRole: string, user: User) => void,
    deleteUserInRole: (idRole: string, idUser: string) => void,
    setStateInPermission: (idRole: string, idPermission: string, value: boolean) => void
}

const useRoleStore = create<State>()(
    devtools(
        immer((set) => ({
            roles: [],

            addRole: (role: Role) => set((state) => {
                state.roles.push(role);
            }),
            deleteRole: (id: string) => set((state) => {
                state.roles = state.roles.filter(role => role.id !== id);
            }),

            addUserInRole: (idRole: string, user:User) => set((state) => {
                const role = state.roles.find(role => role.id === idRole);
                if (role) {
                    role.users.push(user);
                }
            }),
            deleteUserInRole: (idRole: string, idUser:string) => set((state) => {
                const role = state.roles.find(role => role.id === idRole);
                if (role) {
                    role.users = role.users.filter(user => user.name !== idUser);
                }
            }),

            setStateInPermission: (idRole: string, idPermission:string, value:boolean) => set((state) => {
                const role = state.roles.find(role => role.id === idRole);
                if (role) {
                    const permission = role.permission.find(p => p.id === idPermission);
                    if (permission) {
                        permission.state = value;
                    }
                }
            }),
        }))
    )
);

export default useRoleStore;