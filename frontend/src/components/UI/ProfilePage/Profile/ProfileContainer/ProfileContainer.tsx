import {
    ContainerColumn,
    ContainerRow
} from "@/components/UI/ProfilePage/Profile/styled/Containers";
import AvatarProfileContainer from "@/components/UI/ProfilePage/Profile/AvatarProfileContainer/AvatarProfileContainer";
import MainInfoContainer from "@/components/UI/ProfilePage/Profile/MainInfoContainer/MainInfoContainer";
import AdditionalInfoContainer
    from "@/components/UI/ProfilePage/Profile/AdditionalInfoContainer/AdditionalInfoContainer";
import React from "react";

interface ProfileContainerProps {
    style?: React.CSSProperties;
}

export default function ProfileContainer({style}:ProfileContainerProps) {
    return (
        <ContainerRow style={style}>
            <div style={{width: "calc(452rem/16)", margin: "calc(10rem/16)"}}>
                <ContainerColumn>
                    <AvatarProfileContainer style={{height: "calc(570rem/16)", marginBottom: "calc(10rem/16)"}}/>
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