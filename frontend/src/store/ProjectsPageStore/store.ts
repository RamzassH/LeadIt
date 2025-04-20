import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import {devtools} from "zustand/middleware";

export interface Project {
    id: string,
    name: string,
    description: string
}

interface State {
    projects: Project[]

    addProject: (project: Project) => void
}

const useProjectStore = create<State>()(
    devtools(
        immer((set) => ({
            projects: [],

            addProject: (project: Project) => set((state) => {
                state.projects.push(project)
            })
        }))
    )
);

export default useProjectStore;