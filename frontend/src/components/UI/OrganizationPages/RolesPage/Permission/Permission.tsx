import React from 'react';
import {PermissionCard, PermissionDescription, PermissionItem, Header} from "@/components/UI/OrganizationPages/RolesPage/Permission/styled/styled";
import {Box, Button, Divider, Switch, Typography} from "@mui/material";

interface Permission {
    name: string;
    description: string;
    enabled: boolean;
}

interface PermissionProps {
    permissions: Permission[];
    onTogglePermission: (permissionName: string) => void;
    onResetPermissions: () => void;
}

const Permission: React.FC<PermissionProps> = ({
                                                                 permissions,
                                                                 onTogglePermission,
                                                                 onResetPermissions,
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
                <React.Fragment key={permission.name}>
                    <PermissionItem>
                        <Box>
                            <Typography variant="subtitle1">{permission.name}</Typography>
                            <PermissionDescription>
                                {permission.description}
                            </PermissionDescription>
                        </Box>
                        <Switch
                            checked={permission.enabled}
                            onChange={() => onTogglePermission(permission.name)}
                            color="primary"
                        />
                    </PermissionItem>
                    {index < permissions.length - 1 && <Divider />}
                </React.Fragment>
            ))}
        </PermissionCard>
    );
};

export default Permission;