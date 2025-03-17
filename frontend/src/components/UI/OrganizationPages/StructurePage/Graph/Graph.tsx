"use client"
import React, {useCallback, useEffect, useMemo, useState} from 'react';
import CustomNode from '@/components/UI/OrganizationPages/StructurePage/Graph/Node/Node';
import ReactFlow, {
    addEdge,
    Background,
    Connection,
    Controls,
    Edge,
    Node,
    Position,
    useEdgesState,
    useNodesState,
} from 'reactflow';
import 'reactflow/dist/style.css';
import CustomEdge from "@/components/UI/OrganizationPages/StructurePage/Graph/Edge/Edge";
import ModalNode from "@/components/UI/OrganizationPages/StructurePage/Graph/ModalNode/ModalNode";
import useDevelopersStore from "@/components/UI/OrganizationPages/StructurePage/Graph/store";

export default function Graph(props: any) {
    const {nodes, edges, setEdges} = useDevelopersStore()
    const [nodesReactFlow, setNodesReactFlow, onNodesChange] = useNodesState(nodes);
    const [edgesReactFlow, setEdgesReactFlow, onEdgesChange] = useEdgesState(edges);
    const [selectedNode, setSelectedNode] = useState<Node | null>(null);
    const [selectedEdge, setSelectedEdge] = useState<Edge | null>(null);
    const [menuPosition, setMenuPosition] = useState<{ x: number; y: number } | null>(null);

    const onNodeClick = useCallback((event: React.MouseEvent, node: Node) => {
        setSelectedNode(node);
    }, []);

    const nodeTypes = useMemo(() => ({
        custom: CustomNode,
    }), []);
    const edgeTypes = useMemo(() => ({
        custom: CustomEdge,
    }), []);

    const onConnect = useCallback(
        (params: Connection | Edge) => {
            const newEdge = {
                ...params,
                type: 'custom', // Указываем кастомный тип ребра
                data: { label: 'Goida' }, // Добавляем данные, если нужно
            };
            setEdgesReactFlow((eds) => addEdge(newEdge, eds));
        },
        [setEdgesReactFlow],
    );

    const onEdgeClick = useCallback(
        (event: React.MouseEvent, edge: Edge) => {
            // Открываем контекстное меню при клике правой кнопкой мыши
            if (event.button === 0) {
                setSelectedEdge(edge);
                setMenuPosition({ x: event.clientX, y: event.clientY });
            }
        },
        [],
    );

    const deleteEdge = useCallback(() => {
        if (selectedEdge) {
            setEdgesReactFlow((eds) => eds.filter((e) => e.id !== selectedEdge.id));
            setSelectedEdge(null);
            setMenuPosition(null);
        }
    }, [selectedEdge, setEdgesReactFlow]);

    useEffect(() => {
        setEdges(edgesReactFlow)
    }, [edgesReactFlow])

    return (
        <div style={{ width: '1000px', height: '1000px', /*border: '1px solid gray'*/ }}>
            <ReactFlow
                nodes={nodesReactFlow}
                edges={edgesReactFlow}
                onNodesChange={onNodesChange}
                onEdgesChange={onEdgesChange}
                onConnect={onConnect}
                nodeTypes={nodeTypes}
                edgeTypes={edgeTypes}
                onNodeClick={onNodeClick}
                onEdgeClick={onEdgeClick}
                fitView
            >
                <Controls/>
            </ReactFlow>
            <ModalNode node={selectedNode} onClose={() => setSelectedNode(null)} />
            {/* Контекстное меню */}
            {menuPosition && (
                <div
                    style={{
                        position: 'fixed',
                        top: menuPosition.y,
                        left: menuPosition.x,
                        backgroundColor: 'white',
                        boxShadow: '0 2px 4px rgba(0, 0, 0, 0.2)',
                        padding: '8px',
                        borderRadius: '4px',
                        zIndex: 1000,
                    }}
                >
                    <button onClick={deleteEdge}>Удалить ребро</button>
                </div>
            )}
        </div>
    );
}