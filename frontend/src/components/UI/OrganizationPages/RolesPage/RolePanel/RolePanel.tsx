import {Box, Tab, Tabs} from '@mui/material';
import React from 'react';

interface RolePanelProps {
    children?: React.ReactNode;
    index: number;
    value: number;
}


const RolePanel: React.FC<RolePanelProps> = (props) => {
    const { children, value, index, ...other } = props;

    return (
        <div
            role="tabpanel"
            hidden={value !== index}
            id={`simple-tabpanel-${index}`}
            aria-labelledby={`simple-tab-${index}`}
            {...other}
        >
            {value === index && <Box sx={{ p: 3 }}>{children}</Box>}
        </div>
    );
};

export default RolePanel;