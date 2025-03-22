import {
    Background,
    BackgroundContainer, ContainerColumn,
    ContainerColumnCenter,
} from "@/components/UI/ProfilePage/Profile/styled/Containers";
import {AvatarContainer} from "@/components/UI/ProfilePage/Profile/styled/Avatar";
import {
    AvatarEditButton,
    AvatarEditButtonText
} from "@/components/UI/ProfilePage/Profile/ProfileButton/styled/ProfileButton";
import React from "react";
import useUserInfoStore from "@/components/UI/ProfilePage/store";

interface AvatarContainerProps {
    style?: React.CSSProperties;
}

interface ImageData {
    xPosition: number;
    yPosition: number;
    scale: number;
}

export default function AvatarProfileContainer({style}:AvatarContainerProps) {
    const avatar = useUserInfoStore(state => state.info.avatar);

    return (
        <BackgroundContainer style={style}>
            <Background style={{
                width: "calc(100% - 36rem/16 * 2)",
                height: "calc(100% - 36rem/16 * 2)",
                margin: "calc(36rem/16) calc(36rem/16)"
                }}
            >
                <ContainerColumn style={{alignItems: "center"}}>
                    <AvatarContainer style={{marginTop: "calc(36rem/16)"}}>
                        <img src={avatar.src} style={{
                                position: "relative",
                                left: `calc(${avatar.positionX}rem/16)`,
                                top: `calc(${avatar.positionY}rem/16)`
                            }} alt="Пахнешь слабостью"/>
                    </AvatarContainer>
                    <div style={{
                        width: "100%",
                        height: "fit-content",
                        display: "flex",
                        justifyContent: "center",
                        margin: "calc(10rem/16) calc(0rem/16)"
                        }}
                    >
                        <AvatarEditButton onClick={(event) => {}}>
                            <AvatarEditButtonText>
                                Изменить аватар
                            </AvatarEditButtonText>
                        </AvatarEditButton>
                    </div>
                </ContainerColumn>
            </Background>
        </BackgroundContainer>
    )
}