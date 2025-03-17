import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import {devtools} from "zustand/middleware";
import {initialEdges, initialNodes} from "@/components/UI/OrganizationPages/StructurePage/Graph/data";
import {Node, Edge} from "reactflow"

interface DeveloperInfo {
    id: string;
    roles: string[];
    projects: string[];
}

interface State {
    nodes: Node[];
    developers: DeveloperInfo[];
    edges: Edge[];

    addNode: (node: Node, developer: DeveloperInfo) => void;
    deleteNode: (id: string) => void;
    setEdges: (edges: Edge[]) => void;

    addRole: (id: string, role: string) => void;
    deleteRole: (id: string, role: string) => void;

    addProject: (id: string, project: string) => void;
    deleteProject: (id: string, project: string) => void;
}

const useDevelopersStore = create<State>()(
    devtools(
        immer((set) => ({
            nodes: initialNodes,
            developers: [
                {
                    id: '1',
                    roles: [],
                    projects: [],
                },
                {
                    id: '2',
                    roles: [],
                    projects: [],
                },
                {
                    id: '3',
                    roles: [],
                    projects: [],
                },
                {
                    id: '4',
                    roles: [],
                    projects: [],
                },
                {
                    id: '5',
                    roles: [],
                    projects: [],
                },
                {
                    id: '6',
                    roles: [],
                    projects: [],
                },
                {
                    id: '7',
                    roles: [],
                    projects: [],
                },
            ],
            edges: initialEdges,

            addNode: (node: Node, developer: DeveloperInfo) => set((state) => {
                if (node.id !== developer.id) {
                    throw new Error('Developer ID is missing');
                }
                state.nodes.push(node);
                state.developers.push(developer);
            }),

            deleteNode: (id: string) => set((state) => {
                state.nodes = state.nodes.filter((node: Node) => node.id !== id);
                state.developers = state.developers.filter((developer: DeveloperInfo) => developer.id !== id);
            }),

            setEdges: (edges: Edge[]) => set((state) => {
                state.edges = edges;
            }),

            addRole: (id: string, role: string) => set((state) => {
                let item = state.developers.find((developer: DeveloperInfo) => developer.id === id);
                if (item !== undefined && !item.roles.includes(role)) {
                    item.roles.push(role);
                }
            }),

            deleteRole: (id: string, role: string) => set((state) => {
                let item = state.developers.find((developer: DeveloperInfo) => developer.id === id);
                if (item !== undefined && item.roles.includes(role)) {
                    item.roles = item.roles.filter(it => it !== role);
                }
            }),

            addProject: (id: string, project: string) => set((state) => {
                let item = state.developers.find((developer: DeveloperInfo) => developer.id === id);
                if (item !== undefined && !item.projects.includes(project)) {
                    item.projects.push(project);
                }
            }),

            deleteProject: (id: string, project: string) => set((state) => {
                let item = state.developers.find((developer: DeveloperInfo) => developer.id === id);
                if (item !== undefined && item.projects.includes(project)) {
                    item.projects = item.projects.filter(it => it !== project);
                }
            }),
        }))
    )
);

export default useDevelopersStore;