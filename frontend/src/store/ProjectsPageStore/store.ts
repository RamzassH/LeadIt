import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import {devtools} from "zustand/middleware";

export interface Project {
    id: number,
    name: string,
    description: string
    organization_id: number,
    image: string,
}

interface State {
    projects: Project[]

    addProject: (project: Project) => void,
    deleteProject: (projectId: number) => void,
    setProjects: (projects: Project[]) => void,
}

const useProjectStore = create<State>()(
    devtools(
        immer((set) => ({
            projects: [],

            addProject: (project: Project) => set((state) => {
                state.projects.push(project)
            }),
            deleteProject: (projectId: number) => set((state) => {
                state.projects = state.projects.filter(project => project.id !== projectId);
            }),
            setProjects: (projects: Project[]) => set((state) => {
                state.projects = projects;
            }),
        }))
    )
);

export default useProjectStore;