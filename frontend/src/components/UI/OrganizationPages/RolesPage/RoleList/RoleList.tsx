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
import {Role} from "@/store/RolePageStore/store";

interface RoleListProps {
    roles: Role[];
    totalRoles: number;
    totalMembers: number;
    callback: (roleId: number | null) => void;
}

const RoleList: React.FC<RoleListProps> = ({ roles, totalRoles, totalMembers, callback }) => {
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
                            <RoleName variant="body1">{role.name}</RoleName>
                        </RoleInfo>
                        <MemberCount variant="body2">
                            {//role.memberCount
                            }
                            <PersonOutlineTwoTone/>
                        </MemberCount>
                        <IconButton onClick={() => {callback(role.id)}} size="small">
                            <MoreHorizIcon fontSize="small" />
                        </IconButton>
                    </RoleItem>
                </React.Fragment>
            ))}
        </RolesContainer>
    );
};

export default RoleList;