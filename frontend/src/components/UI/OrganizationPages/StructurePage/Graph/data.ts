// Начальные данные для дерева
import {Edge, Node} from "reactflow";

export const initialNodes: Node[] = [
    {
        id: '1',
        type: 'custom',
        data: { label: 'Руководитель' },
        position: { x: 250, y: 0 },
    },
    {
        id: '2',
        type: 'custom',
        data: { label: 'Начальник отдела 1' },
        position: { x: 100, y: 200 },
    },
    {
        id: '3',
        type: 'custom',
        data: { label: 'Начальник отдела 2' },
        position: { x: 400, y: 200 },
    },
    {
        id: '4',
        type: 'custom',
        data: { label: 'Сотрудник 1' },
        position: { x: 0, y: 400 },
    },
    {
        id: '5',
        type: 'custom',
        data: { label: 'Сотрудник 2' },
        position: { x: 200, y: 400 },
    },
    {
        id: '6',
        type: 'custom',
        data: { label: 'Сотрудник 3' },
        position: { x: 350, y: 400 },
    },
    {
        id: '7',
        type: 'custom',
        data: { label: 'Сотрудник 4' },
        position: { x: 500, y: 400 },
    },
];

export const initialEdges: Edge[] = [
    { id: 'e1-2', source: '1', target: '2', type: 'custom', data: {label: "Goida"}},
    { id: 'e1-3', source: '1', target: '3', type: 'custom', data: {label: "Goida"}},
    { id: 'e2-4', source: '2', target: '4', type: 'custom', data: {label: "Goida"}},
    { id: 'e2-5', source: '2', target: '5', type: 'custom', data: {label: "Goida"}},
    { id: 'e3-6', source: '3', target: '6', type: 'custom', data: {label: "Goida"}},
    { id: 'e3-7', source: '3', target: '7', type: 'custom', data: {label: "Goida"}},
];