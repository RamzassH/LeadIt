import {Box, IconButton, Tab, Tabs} from '@mui/material';
import React, {useEffect, useState} from 'react';
import RolePanel from "@/components/UI/OrganizationPages/RolesPage/RolePanel/RolePanel";
import PermissionList from "@/components/UI/OrganizationPages/RolesPage/PermissionList/PermissionList";
import {ArrowBack} from "@mui/icons-material";
import RoleVisualization from "@/components/UI/OrganizationPages/RolesPage/RoleVisualization/RoleVisualization";
import EmployeeList from "@/components/UI/OrganizationPages/RolesPage/EmployeeList/EmployeeList";
import useRoleStore, {Permission} from "@/store/RolePageStore/store";
import {updateRoleAPI} from "@/api/role/update";

interface RoleSettingsProps {
    callback: (roleId: number | null) => void;
    currentRoleId: number
}

function a11yProps(index: number) {
    return {
        id: `simple-tab-${index}`,
        'aria-controls': `simple-tabpanel-${index}`,
    };
}

const RoleTabs: React.FC<RoleSettingsProps> = ({callback, currentRoleId}) => {
    const roles = useRoleStore(state => state.roles);
    const setPermissionsInRole= useRoleStore(state => state.setPermissionsInRole);
    const [value, setValue] = React.useState(0);
    const [permissions, setPermissions] = useState<Permission[]>([]);

    useEffect(() => {
        const permissionList = roles.find(item => item.id === currentRoleId)?.permissions;
        if (permissionList) {
            setPermissions(permissionList);
        }
    }, [currentRoleId])

    const handleTogglePermission = (permissionId: string) => {
        setPermissions(prev => prev.map(p =>
            p.id === permissionId ? { ...p, state: !p.state } : p
        ));
    };
    const handleResetPermissions = () => {
        setPermissions(prev => prev.map(p => ({ ...p, state: false })));
    };
    const handleChange = (event: React.SyntheticEvent, newValue: number) => {
        setValue(newValue);
    };
    const handleSaveChanges = async () => {
        setPermissionsInRole(currentRoleId, permissions);
        const role = roles.find(item => item.id == currentRoleId);
        let perm = []
        if (role) {
            for (const item of role.permissions) {
                perm.push({key: item.id,  value: item.state});
            }
            await updateRoleAPI(currentRoleId, {...role, permissions: perm}, '')
        }
    }

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
                <PermissionList permissions={permissions} onTogglePermission={handleTogglePermission} onResetPermissions={handleResetPermissions} onSaveChanges={handleSaveChanges}/>
            </RolePanel>
            <RolePanel index={2} value={value}>
                <EmployeeList/>
            </RolePanel>
        </div>
    );
};

export default RoleTabs;