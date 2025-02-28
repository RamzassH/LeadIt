import {DrawerContainer, DrawerBackground} from "@/components/UI/ProfilePage/DrawerSideMenu/styled/DrawerContainer";
import MenuButton from "@/components/UI/ProfilePage/SideMenu/MenuButton/MenuButton";
import Icon from "@/components/UI/ProfilePage/SideMenu/MenuButton/styles/StyleIcon";
import VIPIcon from "@/images/icons-svg/VIPIcon";
import Text from "@/components/UI/ProfilePage/SideMenu/MenuButton/styles/Text";
import MenuDropList from "@/components/UI/ProfilePage/SideMenu/MenuDropList/MenuDropList";
import BookIcon from "@/images/icons-svg/BookIcon";
import React, {forwardRef, useImperativeHandle, useState} from "react";

interface ButtonContent {
    icon: React.ReactNode;
    text: string;
    callback: (event: React.MouseEvent<HTMLButtonElement>) => void;
}

interface DrawerSideMenuProps {
    // В родительском компоненте будет вызываться handleClick через ref
}

const DrawerSideMenu = forwardRef((props: DrawerSideMenuProps, ref) => {
    const [open, setOpen] = React.useState(false);

    const toggleDrawer = (newOpen: boolean) => {
        setOpen(newOpen);
    };

    // С помощью useImperativeHandle мы делаем handleClick доступным для родителя
    useImperativeHandle(ref, () => ({
        triggerHandleClick: () => {toggleDrawer(true)}
    }));

    const list1: ButtonContent[] = [
        { icon: <VIPIcon />, text: "Dada", callback: (event) => {} },
        { icon: <VIPIcon />, text: "Dada", callback: (event) => {} },
    ];
    const list2: ButtonContent[] = [
        { icon: <VIPIcon />, text: "Dada", callback: (event) => {} },
        { icon: <BookIcon />, text: "Dada", callback: (event) => {} },
    ];

    return (
        <DrawerContainer open={open} onClose={() => {toggleDrawer(false)}}>
            <DrawerBackground>
                <MenuButton callback={(event) => {}}>
                    <Icon className="first-icon">
                        <VIPIcon />
                    </Icon>
                    <Text>Гойда</Text>
                </MenuButton>
                <MenuDropList firstIcon={<VIPIcon />} text="Гойда №1">
                    {list1.map((item, index) => (
                        <MenuButton className="list-item" callback={item.callback} key={index}>
                            <Icon className="first-icon">{item.icon}</Icon>
                            <Text>{item.text}</Text>
                        </MenuButton>
                    ))}
                </MenuDropList>
                <MenuDropList firstIcon={<BookIcon />} text="Гойда №2">
                    {list2.map((item, index) => (
                        <MenuButton className="list-item" callback={item.callback} key={index}>
                            <Icon className="first-icon">{item.icon}</Icon>
                            <Text>{item.text}</Text>
                        </MenuButton>
                    ))}
                </MenuDropList>
                <MenuDropList firstIcon={<BookIcon />} text="Гойда №3">
                </MenuDropList>
            </DrawerBackground>
        </DrawerContainer>
    )
});

export default DrawerSideMenu;