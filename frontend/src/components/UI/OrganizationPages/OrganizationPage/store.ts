// store.ts
import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import {devtools} from "zustand/middleware";

interface Logo {
    src: string;
    positionX: number;
    positionY: number;
}

export interface Organization {
    name: string;
    director: string;
}

export interface ContactInfo {
    email: string;
    messenger: string;
    phone: string;
}

interface OrganizationInfo {
    logo: Logo;
    organization: Organization;
    contacts: ContactInfo;
    description: string;
}

interface State {
    info: OrganizationInfo;
    isLoading: boolean;
    error?: string;

    setLogo: (avatar: Logo) => void;
    setOrganization: (organization: Organization) => void;
    setContactInfo: (contactInfo: ContactInfo) => void;
    setDescription: (description: string) => void;
}

const useOrganizationInfoStore = create<State>()(
    devtools(
    immer((set) => ({
        info: {
            logo: {
                src: '',
                positionX: 0,
                positionY: 0,
            },
            organization: {
                name: '',
                director: ''
            },
            contacts: {
                email: '',
                messenger: '',
                phone: '',
            },
            description: '',
        },
        isLoading: false,
        error: '',

        setLogo: (logo: Logo) =>
            set((state) => {
                state.info.logo = logo;
            }),

        setOrganization: (organization: Organization) =>
            set((state) => {
                state.info.organization = organization;
            }),

        setContactInfo: (contactInfo: ContactInfo) =>
            set((state) => {
                state.info.contacts = contactInfo;
            }),

        setDescription: (description: string) =>
            set((state) => {
                state.info.description = description;
            }),
    }))
    )
);

export default useOrganizationInfoStore;