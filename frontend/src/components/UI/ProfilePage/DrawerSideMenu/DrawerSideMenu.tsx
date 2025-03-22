"use client";
import {
    DrawerContainer,
    DrawerBackground,
    DrawerLogo
} from "@/components/UI/ProfilePage/DrawerSideMenu/styled/DrawerContainer";
import MenuButton from "@/components/UI/ProfilePage/SideMenu/MenuButton/MenuButton";
import Icon from "@/components/UI/ProfilePage/SideMenu/MenuButton/styles/StyleIcon";
import VIPIcon from "@/images/icons-svg/VIPIcon";
import Text from "@/components/UI/ProfilePage/SideMenu/MenuButton/styles/Text";
import MenuDropList from "@/components/UI/ProfilePage/SideMenu/MenuDropList/MenuDropList";
import BookIcon from "@/images/icons-svg/BookIcon";
import React, { forwardRef, useImperativeHandle, useState, useEffect } from "react";
import { useRouter } from "next/navigation";

interface ButtonContent {
    icon: React.ReactNode;
    text: string;
    callback: (event: React.MouseEvent<HTMLButtonElement>) => void;
}

interface DrawerSideMenuProps {
    // В родительском компоненте будет вызываться handleClick через ref
}

interface DrawerSideMenuRef {
    triggerHandleClick: () => void;
}

const DrawerSideMenu = forwardRef<DrawerSideMenuRef, DrawerSideMenuProps>((props, ref) => {
    const router = useRouter();
    const [open, setOpen] = useState(false);

    const toggleDrawer = (newOpen: boolean) => {
        setOpen(newOpen);
    };

    useImperativeHandle(ref, () => ({
        triggerHandleClick: () => { toggleDrawer(true); }
    }));

    const list1: ButtonContent[] = [
        { icon: <VIPIcon />, text: "Главная", callback: (event) => {} },
        { icon: <VIPIcon />, text: "Структура", callback: (event) => {router?.push("/organization/structure")} },
        { icon: <VIPIcon />, text: "Роли", callback: (event) => {} },
        { icon: <VIPIcon />, text: "Проекты", callback: (event) => {} }
    ];

    return (
        <DrawerContainer open={open} onClose={() => { toggleDrawer(false); }}>
            <DrawerBackground>
                <DrawerLogo>
                    LeadIt
                </DrawerLogo>
                <MenuButton callback={(event) => { router?.push("/"); }}>
                    <Icon className="first-icon">
                        <VIPIcon />
                    </Icon>
                    <Text>Главная</Text>
                </MenuButton>
                <MenuButton callback={(event) => { router?.push("/profile"); }}>
                    <Icon className="first-icon">
                        <VIPIcon/>
                    </Icon>
                    <Text>Профиль</Text>
                </MenuButton>
                <MenuDropList firstIcon={<VIPIcon/>} text="Организация">
                    {list1.map((item, index) => (
                        <MenuButton className="list-item" callback={item.callback} key={index}>
                            <Icon className="first-icon">{item.icon}</Icon>
                            <Text>{item.text}</Text>
                        </MenuButton>
                    ))}
                </MenuDropList>
                <MenuDropList firstIcon={<BookIcon />} text="Гойда №3">
                    {/* Добавьте детей, если необходимо */}
                </MenuDropList>
            </DrawerBackground>
        </DrawerContainer>
    );
});

export default DrawerSideMenu;