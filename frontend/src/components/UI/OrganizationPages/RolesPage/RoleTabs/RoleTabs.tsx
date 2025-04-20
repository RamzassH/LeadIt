import {Box, IconButton, Tab, Tabs} from '@mui/material';
import React, {useState} from 'react';
import RolePanel from "@/components/UI/OrganizationPages/RolesPage/RolePanel/RolePanel";
import Permission from "@/components/UI/OrganizationPages/RolesPage/Permission/Permission";
import {ArrowBack} from "@mui/icons-material";
import RoleVisualization from "@/components/UI/OrganizationPages/RolesPage/RoleVisualization/RoleVisualization";

interface RoleSettingsProps {
    callback: (roleId: string | null) => void;
}

function a11yProps(index: number) {
    return {
        id: `simple-tab-${index}`,
        'aria-controls': `simple-tabpanel-${index}`,
    };
}

const RoleTabs: React.FC<RoleSettingsProps> = ({callback}) => {
    const [value, setValue] = React.useState(0);
    const [permissions, setPermissions] = useState([
            {
                name: 'Создавать проекты',
                description: 'Дает право создать новый проект в организации',
                enabled: true,
            },
            {
                name: 'Назначать роли',
                description: 'Позволяет добавлять сотрудникам новые роли. Выбор ролей ограничен самой высокой ролью владельца.',
                enabled: false
            }
        ]
    );

    const handleTogglePermission = (permissionName: string) => {
        setPermissions(prev => prev.map(p =>
            p.name === permissionName ? { ...p, enabled: !p.enabled } : p
        ));
    };

    const handleResetPermissions = () => {
        setPermissions(prev => prev.map(p => ({ ...p, enabled: false })));
    };


    const handleChange = (event: React.SyntheticEvent, newValue: number) => {
        setValue(newValue);
    };

    return (
        <div style={{width:'fit-content', marginLeft:'auto', marginRight:'auto' }}>
            <div style={{width:'fit-content', marginLeft:'auto', marginRight:'auto', display:'flex', justifyContent:'center', flexWrap:'nowrap'}}>
                <IconButton onClick={() => {callback(null)}} size="small">
                    <ArrowBack/>
                </IconButton>
                <Box sx={{ borderBottom: 1, borderColor: 'divider' , width: 'fit-content' }}>
                    <Tabs value={value} onChange={handleChange} aria-label="basic tabs example">
                        <Tab label="Настройка отображения" {...a11yProps(0)} />
                        <Tab label="Настройка роли" {...a11yProps(1)} />
                        <Tab label="Управление пользователями" {...a11yProps(2)} />
                    </Tabs>
                </Box>
            </div>
            <RolePanel index={0} value={value}>
                <RoleVisualization/>
            </RolePanel>
            <RolePanel index={1} value={value}>
                <Permission permissions={permissions} onTogglePermission={handleTogglePermission} onResetPermissions={handleResetPermissions}/>
            </RolePanel>
            <RolePanel index={2} value={value}>
                Goida 2
            </RolePanel>
        </div>
    );
};

export default RoleTabs;