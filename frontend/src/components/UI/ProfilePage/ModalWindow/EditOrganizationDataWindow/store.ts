// store.ts
import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import {devtools} from "zustand/middleware";
import React from "react";

interface ProjectInfo {
    organization: string;
    projects: string;
    position: string;
}

interface State {
    isOpen: boolean;

    form: ProjectInfo;

    setOpen: (isOpen: boolean) => void;
    setForm: (form: ProjectInfo) => void;

    handleOpen: (defaultValue: ProjectInfo) => void;
    handleClose: () => void;
    handleSubmit: (callback: (data: ProjectInfo) => void) => void;
    handleInputChange: (event:React.ChangeEvent<HTMLInputElement>) => void;
}

const useProjectInfoStore = create<State>()(
    devtools(
    immer((set) => ({
        isOpen: false,

        form: {organization: "", projects: "", position: ""},

        setOpen: (isOpen: boolean) => set((state) => {
            state.isOpen = isOpen;
        }),
        setForm: (form: ProjectInfo) => set((state) => {
            state.form = form;
        }),

        handleOpen: (defaultValue: ProjectInfo) => set((state) => {
            state.form = defaultValue;
            state.isOpen = true;
        }),

        handleClose: () => set((state) => {
            state.isOpen = false;
        }),

        handleSubmit: (callback: (data: ProjectInfo) => void) => set((state) => {
            callback({...state.form});
            state.isOpen = false;
        }),

        handleInputChange: (event: React.ChangeEvent<HTMLInputElement>) => set((state) => {
            const { name, value } = event.target;
            state.form = {
                ...state.form,
                [name]: value,
            };
        }),
    }))
    )
);

export default useProjectInfoStore;