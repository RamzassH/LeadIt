"use client"
import React, {useRef} from "react";
import OrganizationComponent
    from "@/components/UI/OrganizationPages/OrganizationPage/Organization/OrganizationComponent";
import {useFetching} from "@/hooks/useFetching";
import {createOrganizationServerFunction} from "@/app/(authorized)/organization/actions";
import {Button} from "@mui/material";

export default function Page() {

    return (
        <OrganizationComponent/>

    )
}