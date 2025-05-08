import React from 'react';
import {PermissionCard, PermissionDescription, PermissionItem, Header} from "@/components/UI/OrganizationPages/RolesPage/PermissionList/styled/styled";
import {Box, Button, Divider, Switch, Typography} from "@mui/material";
import {Permission} from "@/store/RolePageStore/store";

interface PermissionProps {
    permissions: Permission[];
    onTogglePermission: (permissionName: string) => void;
    onResetPermissions: () => void;
    onSaveChanges: () => void;
}

const PermissionList: React.FC<PermissionProps> = ({
                                                                 permissions,
                                                                 onTogglePermission,
                                                                 onResetPermissions,
                                                                 onSaveChanges
                                                             }) => {
    return (
        <PermissionCard>
            <Header>
                <Typography variant="h6">Настройки роли</Typography>
                <Button
                    variant="outlined"
                    color="secondary"
                    onClick={onResetPermissions}
                >
                    Сбросить права
                </Button>
            </Header>

            <Divider />

            {permissions.map((permission, index) => (
                <React.Fragment key={permission.id}>
                    <PermissionItem>
                        <Box>
                            <Typography variant="subtitle1">{permission.name}</Typography>
                            <PermissionDescription>
                                {permission.description}
                            </PermissionDescription>
                        </Box>
                        <Switch
                            checked={permission.state}
                            onChange={() => onTogglePermission(permission.id)}
                            color="primary"
                        />
                    </PermissionItem>
                    {index < permissions.length && <Divider />}
                </React.Fragment>
            ))}
            <Button
                style={{marginTop: "calc(16rem/16)"}}
                variant="contained"
                color="primary"
                onClick={onSaveChanges}
            >
                Сохранить
            </Button>
        </PermissionCard>
    );
};

export default PermissionList;