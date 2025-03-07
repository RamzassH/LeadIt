// store.ts
import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import {devtools} from "zustand/middleware";

interface Avatar {
    src: string;
    positionX: number;
    positionY: number;
}

interface FullName {
    name: string;
    surname: string;
    patronymic: string;
}

interface ContactInfo {
    email: string;
    messenger: string;
}

interface ProjectInfo {
    organization: string;
    projects: string;
    position: string;
}

interface UserInfo {
    avatar: Avatar;
    fullName: FullName;
    date: string;
    contacts: ContactInfo;
    projectInfo: ProjectInfo;
    description: string;
}

interface UserState {
    info: UserInfo;
    isLoading: boolean;
    error?: string;

    setAvatar: (avatar: Avatar) => void;
    setFullName: (fullName: FullName) => void;
    setDate: (date: string) => void;
    setContactInfo: (contactInfo: ContactInfo) => void;
    setProjectInfo: (projectInfo: ProjectInfo) => void;
    setDescription: (description: string) => void;
}

const useUserInfoStore = create<UserState>()(
    devtools(
    immer((set) => ({
        info: {
            avatar: {
                src: '',
                positionX: 0,
                positionY: 0,
            },
            fullName: {
                name: '',
                surname: '',
                patronymic: '',
            },
            date: '',
            contacts: {
                email: '',
                messenger: '',
            },
            projectInfo: {
                organization: '',
                projects: '',
                position: '',
            },
            description: '',
        },
        isLoading: false,
        error: '',

        setAvatar: (avatar: Avatar) =>
            set((state) => {
                state.info.avatar = avatar; // Обновляем аватар
            }),

        setFullName: (fullName: FullName) =>
            set((state) => {
                state.info.fullName = fullName; // Обновляем полное имя
            }),

        setDate: (date: string) =>
            set((state) => {
                state.info.date = date; // Обновляем дату
            }),

        setContactInfo: (contactInfo: ContactInfo) =>
            set((state) => {
                state.info.contacts = contactInfo; // Обновляем контактную информацию
            }),

        setProjectInfo: (projectInfo: ProjectInfo) =>
            set((state) => {
                state.info.projectInfo = projectInfo; // Обновляем информацию о проекте
            }),

        setDescription: (description: string) =>
            set((state) => {
                state.info.description = description; // Обновляем описание
            }),
    }))
    )
);

export default useUserInfoStore;