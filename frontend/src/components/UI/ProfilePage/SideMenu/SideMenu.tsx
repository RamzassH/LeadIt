import React, { useState, useImperativeHandle, forwardRef } from "react";
import MenuButton from "@/components/UI/ProfilePage/SideMenu/MenuButton/MenuButton";
import Icon from "@/components/UI/ProfilePage/SideMenu/MenuButton/styles/StyleIcon";
import Text from "@/components/UI/ProfilePage/SideMenu/MenuButton/styles/Text";
import MenuContainer from "@/components/UI/ProfilePage/SideMenu/styles/MenuContainer";
import MenuDropList from "@/components/UI/ProfilePage/SideMenu/MenuDropList/MenuDropList";
import VIPIcon from "@/images/icons-svg/VIPIcon";
import BookIcon from "@/images/icons-svg/BookIcon";

interface ButtonContent {
    icon: React.ReactNode;
    text: string;
    callback: (event: React.MouseEvent<HTMLButtonElement>) => void;
}

interface SideMenuProps {
    // В родительском компоненте будет вызываться handleClick через ref
}

const SideMenu = forwardRef((props: SideMenuProps, ref) => {
    const [visibleButton, setVisible] = useState("non-visible");
    const [openMenu, setOpen] = useState("close");

    const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        if (openMenu === "") {
            setOpen("close");
            setVisible("non-visible");
        } else {
            setOpen("");
            setVisible("");
        }
    };

    // С помощью useImperativeHandle мы делаем handleClick доступным для родителя
    useImperativeHandle(ref, () => ({
        triggerHandleClick: handleClick
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
        <MenuContainer className={openMenu}>
            <MenuButton callback={(event) => {}}>
                <Icon className="first-icon">
                    <VIPIcon />
                </Icon>
                <Text className={visibleButton}>Гойда</Text>
            </MenuButton>
            <MenuDropList firstIcon={<VIPIcon />} text="Гойда №1" visibility={visibleButton}>
                {list1.map((item, index) => (
                    <MenuButton callback={item.callback} key={index}>
                        <Icon className="first-icon">{item.icon}</Icon>
                        <Text className={visibleButton}>{item.text}</Text>
                    </MenuButton>
                ))}
            </MenuDropList>
            <MenuDropList firstIcon={<BookIcon />} text="Гойда №2" visibility={visibleButton}>
                {list2.map((item, index) => (
                    <MenuButton callback={item.callback} key={index}>
                        <Icon className="first-icon">{item.icon}</Icon>
                        <Text className={visibleButton}>{item.text}</Text>
                    </MenuButton>
                ))}
            </MenuDropList>
            <div style={{ marginBottom: "auto" }}></div>
            <div />
        </MenuContainer>
    );
});

export default SideMenu;
