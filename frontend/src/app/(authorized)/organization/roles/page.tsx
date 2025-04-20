"use client"
import React, {useRef, useState} from "react";
import RoleList from "@/components/UI/OrganizationPages/RolesPage/RoleList/RoleList";
import RoleTabs from "@/components/UI/OrganizationPages/RolesPage/RoleTabs/RoleTabs";

export default function Page() {
    const [idRole, setIdRole] = useState<string | null>(null);

    const roles = [
        {
            id: '1',
            name: 'Директор',
            memberCount: 3,
            callback: setIdRole,
            icon: '',
        },
        {
            id: '2',
            name: 'Разработчик',
            memberCount: 10,
            callback: setIdRole,
            icon: '',
        },
        {
            id: '3',
            name: 'Долбаеб',
            memberCount: 6,
            callback: setIdRole,
            icon: '',
        },
    ];

    return (
        <div style={{width:'100%', justifyContent:'center', alignItems: 'center', marginTop: "calc(0rem/16)"}}>
            {idRole ?
                <RoleTabs callback={setIdRole}/>:
                <RoleList roles={roles} totalRoles={3} totalMembers={3}/>
            }
        </div>
    )
}