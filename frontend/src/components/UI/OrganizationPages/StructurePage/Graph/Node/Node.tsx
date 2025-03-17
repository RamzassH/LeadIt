import React from 'react';
import { Handle, Position } from 'reactflow';
import { Card, CardContent, Typography, Avatar } from '@mui/material';

interface CustomNodeProps {
    data: {
        label: string;
    };
}


const CustomNode: React.FC<CustomNodeProps> = ({ data }) => {
    return (
        <div>
            <Handle
                type="source"
                position={Position.Bottom}
                style={{ bottom: 'calc(-6rem/16)', backgroundColor: '#ff5722', width: 'calc(12rem/16)', height: 'calc(12rem/16)', zIndex: 1 }}
            />

            {/* Handle для входящих соединений */}
            <Handle
                type="target"
                position={Position.Top}
                style={{ top: 'calc(-6rem/16)', backgroundColor: '#4caf50', width: 'calc(12rem/16)', height: 'calc(12rem/16)', zIndex: 1 }}
            />
            <Card
                sx={{
                    minWidth: 120,
                    textAlign: 'center',
                    backgroundColor: '#1976d2',
                    color: 'white',
                    borderRadius: '8px',
                    boxShadow: '0 2px 4px rgba(0, 0, 0, 0.2)',
                    position: 'relative', // Для позиционирования Handle
                }}
            >
                {/* Handle для исходящих соединений */}


                <CardContent>
                    {/* Аватар */}
                    <Avatar
                        sx={{
                            bgcolor: 'white',
                            color: '#1976d2',
                            margin: 'auto',
                            marginBottom: '8px',
                        }}
                    >
                        {data.label[0]} {/* Первая буква имени */}
                    </Avatar>

                    {/* Название узла */}
                    <Typography variant="body1" sx={{ fontWeight: 'bold' }}>
                        {data.label}
                    </Typography>
                </CardContent>
            </Card>
        </div>
    );
};

export default CustomNode;