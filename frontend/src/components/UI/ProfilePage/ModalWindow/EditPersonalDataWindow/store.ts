// store.ts
import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import {devtools} from "zustand/middleware";
import React from "react";

export interface PersonalData {
    name: string;
    surname: string;
    patronymic: string;
    date: string;
    email: string;
    messenger: string;
}

interface State {
    isOpen: boolean;

    form: PersonalData;

    setOpen: (isOpen: boolean) => void;
    setForm: (form: PersonalData) => void;

    handleOpen: (defaultValue: PersonalData) => void;
    handleClose: () => void;
    handleSubmit: (callback: (data: PersonalData) => void) => void;
    handleInputChange: (event:React.ChangeEvent<HTMLInputElement>) => void;
}

const usePersonalDataStore = create<State>()(
    devtools(
    immer((set) => ({
        isOpen: false,

        form: {
            name: `Personal Data`,
            surname: `Surname`,
            patronymic: `Patronymic Data`,
            date: `Date of Birth`,
            email: `Email Data`,
            messenger: `Messenger Data`,
        },

        setOpen: (isOpen: boolean) => set((state) => {
            state.isOpen = isOpen; // Обновляем аватар
        }),
        setForm: (form: PersonalData) => set((state) => {
            state.form = form;
        }),


        handleOpen: (defaultValue: PersonalData) => set((state) => {
            state.form = defaultValue;
            state.isOpen = true;
        }),

        handleClose: () => set((state) => {
            state.isOpen = false;
        }),

        handleSubmit: (callback: (data: PersonalData) => void) => set((state) => {
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

export default usePersonalDataStore;