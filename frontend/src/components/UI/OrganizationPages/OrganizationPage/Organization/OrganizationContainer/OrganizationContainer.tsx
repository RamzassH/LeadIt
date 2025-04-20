import {
    ContainerColumn,
    ContainerRow
} from "@/components/UI/OrganizationPages/OrganizationPage/Organization/styled/Containers";
import LogoContainer from "@/components/UI/OrganizationPages/OrganizationPage/Organization/LogoContainer/LogoContainer";
import MainInfoContainer from "@/components/UI/OrganizationPages/OrganizationPage/Organization/MainInfoContainer/MainInfoContainer";
import AdditionalInfoContainer
    from "@/components/UI/OrganizationPages/OrganizationPage/Organization/AdditionalInfoContainer/AdditionalInfoContainer";
import React from "react";

interface ProfileContainerProps {
    style?: React.CSSProperties;
}

export default function OrganizationContainer({style}:ProfileContainerProps) {
    return (
        <ContainerRow style={style}>
            <div style={{width: "calc(452rem/16)", margin: "calc(10rem/16)"}}>
                <ContainerColumn>
                    <LogoContainer style={{height: "calc(570rem/16)", marginBottom: "calc(10rem/16)"}}/>
                    <AdditionalInfoContainer style={{height: "calc(100% - 570rem/16 - 20rem/16)", marginTop:"calc(10rem/16)"}}/>
                </ContainerColumn>
            </div>
            <div style={{width: "calc(100% - 452rem/16)", margin: "calc(10rem/16)"}}>
                <ContainerColumn>
                    <MainInfoContainer/>
                </ContainerColumn>
            </div>
        </ContainerRow>
    )
}