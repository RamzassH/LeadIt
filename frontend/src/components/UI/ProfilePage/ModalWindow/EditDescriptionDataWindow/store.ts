// store.ts
import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import { devtools } from 'zustand/middleware';
import React from "react";

interface DescriptionData {
    description: string;
}

interface State {
    isOpen: boolean;

    form: DescriptionData;

    setOpen: (isOpen: boolean) => void;
    setForm: (form: DescriptionData) => void;

    handleOpen: (defaultValue: DescriptionData) => void;
    handleClose: () => void;
    handleSubmit: (callback: (data: string) => void) => void;
    handleInputChange: (event:React.ChangeEvent<HTMLInputElement>) => void;
}

const useDescriptionStore = create<State>()(
    devtools(
    immer((set) => ({
        isOpen: false,
        form: { description: "" },

        setOpen: (isOpen: boolean) => set((state) => {
            state.isOpen = isOpen;
        }),
        setForm: (form: DescriptionData) => set((state) => {
            state.form = form;
        }),

        handleOpen: (defaultValue: DescriptionData) => set((state) => {
            state.form = defaultValue;
            state.isOpen = true;
        }),

        handleClose: () => set((state) => {
            state.isOpen = false;
        }),

        handleSubmit: (callback: (data: string) => void) => set((state) => {
            callback(state.form.description);
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

export default useDescriptionStore;