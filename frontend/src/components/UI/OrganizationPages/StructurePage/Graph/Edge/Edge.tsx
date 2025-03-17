import React from 'react';
import { BaseEdge, EdgeLabelRenderer, EdgeProps, getBezierPath } from 'reactflow';

const CustomEdge: React.FC<EdgeProps> = ({
                                             id,
                                             sourceX,
                                             sourceY,
                                             targetX,
                                             targetY,
                                             sourcePosition,
                                             targetPosition,
                                             data,
                                             selected, // Проп selected, который передается React Flow
                                         }) => {
    const [edgePath, labelX, labelY] = getBezierPath({
        sourceX,
        sourceY,
        sourcePosition,
        targetX,
        targetY,
        targetPosition,
    });

    // Стиль для выделенного ребра
    const edgeStyle = selected
        ? { stroke: '#ff0000', strokeWidth: 3 } // Красный цвет и толстая линия
        : { stroke: '#ff5722', strokeWidth: 2 }; // Оранжевый цвет и тонкая линия

    return (
        <>
            <BaseEdge id={id} path={edgePath} style={edgeStyle} />
            {/*<EdgeLabelRenderer>
                <div
                    style={{
                        position: 'absolute',
                        transform: `translate(-50%, -50%) translate(${labelX}px,${labelY}px)`,
                        background: '#fff',
                        padding: '4px 8px',
                        borderRadius: '4px',
                        fontSize: '12px',
                        boxShadow: '0 2px 4px rgba(0, 0, 0, 0.2)',
                    }}
                >
                    {data?.label} {/* Текст на ребре *}
                </div>
            </EdgeLabelRenderer>*/}
        </>
    );
};

export default CustomEdge;