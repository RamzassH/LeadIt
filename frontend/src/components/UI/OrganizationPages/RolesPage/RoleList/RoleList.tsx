import React from 'react';
import {Typography, Divider, IconButton } from '@mui/material';
import MoreHorizIcon from '@mui/icons-material/MoreHoriz';
import {
    Header,
    RolesContainer,
    RoleItem,
    RoleInfo,
    RoleName,
    ColorIndicator, MemberCount
} from "@/components/UI/OrganizationPages/RolesPage/RoleList/styled/styled";
import {PersonOutlineTwoTone} from "@mui/icons-material";

interface Role {
    id: string;
    name: string;
    memberCount: number;
    callback: (value: string) => void;
    color?: string;
    icon?: string;
}

interface RoleListProps {
    roles: Role[];
    totalRoles: number;
    totalMembers: number;
}

const RoleList: React.FC<RoleListProps> = ({ roles, totalRoles, totalMembers }) => {
    return (
        <RolesContainer>
            <Header>
                <Typography variant="subtitle1">
                    РОЛИ — {totalRoles}
                </Typography>
                <Typography variant="subtitle1">
                    УЧАСТНИКИ — {totalMembers}
                </Typography>
            </Header>
            {roles.map((role, index) => (
                <React.Fragment key={role.id}>
                    <RoleItem>
                        <RoleInfo>
                            {role.icon && <span>{role.icon}</span>}
                            <RoleName variant="body1">{role.name}</RoleName>
                        </RoleInfo>
                        <MemberCount variant="body2">
                            {role.memberCount}
                            <PersonOutlineTwoTone/>
                        </MemberCount>
                        <IconButton onClick={() => {role.callback(role.id)}} size="small">
                            <MoreHorizIcon fontSize="small" />
                        </IconButton>
                    </RoleItem>
                </React.Fragment>
            ))}
        </RolesContainer>
    );
};

export default RoleList;