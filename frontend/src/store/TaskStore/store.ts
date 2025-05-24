import {create} from "zustand";
import {devtools} from "zustand/middleware";
import {immer} from "zustand/middleware/immer";
import {progressBar, status, tasks} from "@/store/TaskStore/data";

export interface Task {
    id: number,
    name: string,
    description: string,
    project_id:  number,
    setter_id:  number,
    solver_id:  number,
    progress_bar_id: number,
    current_status_id: number,
    allocated_time: number,
    wasted_time:  number,
    active: boolean,
    last_start_time: number,
}

export interface ProgressBar {
    id: number,
    name: string,
    project_id: number,
}

export interface Status {
    id:  number,
    progress_bar_id: number,
    name: string,
    order:  number,
    notify_roles_ids: number[]
}

interface State {
    tasks: Task[],
    progressBar: ProgressBar,
    status: Status[],

    setStatus: (status: Status[]) => void,
    setProgressBar: (progressBar: ProgressBar) => void,
    setTasks: (tasks: Task[]) => void,
}


const useTaskStore = create<State>()(
    devtools(
        immer((set) => ({
            tasks: tasks,
            progressBar: progressBar,
            status: status,

            setStatus: (status) => set((state) => {
                state.status = status;
            }),
            setProgressBar: (progressBar: ProgressBar) => set((state) => {
                state.progressBar = progressBar
            }),
            setTasks: (tasks: Task[]) => set((state) => {
                state.tasks = tasks;
            }),
        }))
    )
);

export default useTaskStore;