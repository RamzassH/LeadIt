import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import { devtools } from 'zustand/middleware';
import React from "react";

interface Organization {
    name: string;
}

interface State {
    isOpen: boolean;

    form: Organization;

    setOpen: (isOpen: boolean) => void;
    setForm: (form: Organization) => void;

    handleOpen: (defaultValue: Organization) => void;
    handleClose: () => void;
    handleSubmit: (callback: (data: Organization) => void) => void;
    handleInputChange: (event:React.ChangeEvent<HTMLInputElement>) => void;
}

const useAddingOrganizationStore = create<State>()(
    devtools(
        immer((set) => ({
            isOpen: false,
            form: { name: "" },

            setOpen: (isOpen: boolean) => set((state) => {
                state.isOpen = isOpen;
            }),
            setForm: (form: Organization) => set((state) => {
                state.form = form;
            }),

            handleOpen: (defaultValue: Organization) => set((state) => {
                state.form = defaultValue;
                state.isOpen = true;
            }),

            handleClose: () => set((state) => {
                state.isOpen = false;
            }),

            handleSubmit: (callback: (data: Organization) => void) => set((state) => {
                callback(state.form);
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

export default useAddingOrganizationStore;