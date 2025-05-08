"use client"
import React, {useEffect, useState} from "react";
import RoleList from "@/components/UI/OrganizationPages/RolesPage/RoleList/RoleList";
import RoleTabs from "@/components/UI/OrganizationPages/RolesPage/RoleTabs/RoleTabs";
import useRoleStore, {Role} from "@/store/RolePageStore/store";
import {useFetching} from "@/hooks/useFetching";
import {getRolesAPI} from "@/api/role/get";
import getPermissionList from "@/utils/getPermissionList";

export default function Page() {
    const {setRoles, roles} = useRoleStore()
    const [idRole, setIdRole] = useState<number | null>(null);

    const [getRoles, isLoading, error] = useFetching(async () => {
        const response = await getRolesAPI(1, "")
        let result: Role[] = []
        for (const role of response.data) {
            let rolePermissions = getPermissionList()
            for (const permission of role.permissions) {
                rolePermissions.forEach(item => {
                    if (item.id === permission.name) {
                        item.state = permission.value;
                    }
                });
            }
            result.push({...role, permissions: rolePermissions})
        }
        setRoles(result);
    })

    useEffect(() => {
        getRoles()
    }, [])

    useEffect(() => {
        // Временная заглушка
        if (error) {
            setRoles([
                {
                    id: 1,
                    organization_id: 3,
                    parent_role_id: 3,
                    name: "Директор",
                    permissions: getPermissionList()
                },
                {
                    id: 2,
                    organization_id: 3,
                    parent_role_id: 3,
                    name: "Разработчик",
                    permissions: getPermissionList()
                },
                {
                    id: 3,
                    organization_id: 3,
                    parent_role_id: 3,
                    name: "Долбаеб",
                    permissions: getPermissionList()
                }
            ])
        }
    }, [error]);

    return (
        <div style={{width:'100%', justifyContent:'center', alignItems: 'center', marginTop: "calc(0rem/16)"}}>
            {idRole ?
                <RoleTabs callback={setIdRole} currentRoleId={idRole}/>:
                <RoleList roles={roles} totalRoles={3} totalMembers={3} callback={setIdRole}/>
            }
        </div>
    )
}