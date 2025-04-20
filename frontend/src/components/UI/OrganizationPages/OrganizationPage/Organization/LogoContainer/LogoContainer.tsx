import {
    Background,
    BackgroundContainer, ContainerColumn,
    ContainerColumnCenter,
} from "@/components/UI/OrganizationPages/OrganizationPage/Organization/styled/Containers";
import {Logo} from "@/components/UI/OrganizationPages/OrganizationPage/Organization/styled/Logo";
import {
    LogoEditButton,
    LogoEditButtonText
} from "@/components/UI/OrganizationPages/OrganizationPage/Organization/OrganizationButton/styled/styled";
import React from "react";
import useUserInfoStore from "@/components/UI/OrganizationPages/OrganizationPage/store";

interface AvatarContainerProps {
    style?: React.CSSProperties;
}

interface ImageData {
    xPosition: number;
    yPosition: number;
    scale: number;
}

export default function AvatarProfileContainer({style}:AvatarContainerProps) {
    const logo = useUserInfoStore(state => state.info.logo);

    return (
        <BackgroundContainer style={style}>
            <Background style={{
                width: "calc(100% - 36rem/16 * 2)",
                height: "calc(100% - 36rem/16 * 2)",
                margin: "calc(36rem/16) calc(36rem/16)"
            }}
            >
                <ContainerColumn style={{alignItems: "center"}}>
                    <Logo style={{marginTop: "calc(36rem/16)"}}>
                        <img src={logo.src} style={{
                            position: "relative",
                            left: `calc(${logo.positionX}rem/16)`,
                            top: `calc(${logo.positionY}rem/16)`
                        }} alt="Пахнешь слабостью"/>
                    </Logo>
                    <div style={{
                        width: "100%",
                        height: "fit-content",
                        display: "flex",
                        justifyContent: "center",
                        margin: "calc(10rem/16) calc(0rem/16)"
                    }}
                    >
                        <LogoEditButton onClick={(event) => {}}>
                            <LogoEditButtonText>
                                Изменить аватар
                            </LogoEditButtonText>
                        </LogoEditButton>
                    </div>
                </ContainerColumn>
            </Background>
        </BackgroundContainer>
    )
}