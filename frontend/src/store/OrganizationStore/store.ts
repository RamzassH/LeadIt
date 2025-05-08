import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import {devtools} from "zustand/middleware";
import React from "react";

export interface Organization {
    id: number,
    name: string,
    organizer_id: number,
    description: string,
    image:  string,
}

interface State {
    organizations: Organization[],

    setOrganizations: (organizations: Organization[]) => void,
}

const useOrganizationStore = create<State>()(
    devtools(
        immer((set) => ({
            organizations: [],

            setOrganizations: (organizations: Organization[]) => set((state) => {
                state.organizations = organizations;
            })
        }))
    )
);

export default useOrganizationStore;