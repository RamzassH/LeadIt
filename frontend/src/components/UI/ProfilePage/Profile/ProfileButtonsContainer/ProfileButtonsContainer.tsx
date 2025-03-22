import {BackgroundContainer, ContainerRow} from "@/components/UI/ProfilePage/Profile/styled/Containers";
import MenuButton from "@/components/UI/ProfilePage/SideMenu/MenuButton/MenuButton";
import TextButton from "@/components/UI/ProfilePage/SideMenu/MenuButton/styles/Text";
import ProfileButton from "@/components/UI/ProfilePage/Profile/ProfileButton/ProfileButton";
import React from "react";
import SvgIcon from "@/components/UI/ProfilePage/SideMenu/MenuDropList/svg/SvgIcon";

interface ButtonProps {
    id?: string;
    icon?: React.ReactNode;
    text?: string;
    children?: React.ReactNode;
    callback: (event: React.MouseEvent<HTMLButtonElement>) => void;
}

export default function ProfileButtonsContainer() {
    const buttons: ButtonProps[] = [
        {text: "Задачи", callback: (event) => {}},
        {text: "Календарь", callback: (event) => {}},
        {text: "Мои документы", callback: (event) => {}},
    ]

    return (
        <div style={{width: "100%", height: "calc(96rem / 16)", marginBottom: "calc(6rem / 16)"}}>
            <BackgroundContainer>
                <ContainerRow style={{padding: "calc(12rem/16) calc(16rem/16)", alignItems: "center", gap: "calc(6rem / 16)"}}>
                    {buttons.map((item, index) => (
                        <ProfileButton key={index} callback={item.callback} icon={item.icon} text={item.text}>
                            {item.children}
                        </ProfileButton>)
                    )}
                    <div style={{marginRight: "auto"}}/>
                    <ProfileButton callback={()=>{}} text="Ещё">
                        <SvgIcon/>
                    </ProfileButton>
                </ContainerRow>
            </BackgroundContainer>
        </div>
    )
}