"use client"
import "./local.css"
import MenuButton from "@/components/UI/ProfilePage/SideMenu/MenuButton/MenuButton";
import theme from "../../../theme/theme";
import {margin, ThemeProvider} from "@mui/system";
import SideMenu from "@/components/UI/ProfilePage/SideMenu/SideMenu";
import React, {useRef} from "react";
import Header from "@/components/UI/ProfilePage/Header/Header";
import ProfileComponent from "@/components/UI/ProfilePage/Profile/ProfileComponent";
import DrawerSideMenu from "@/components/UI/ProfilePage/DrawerSideMenu/DrawerSideMenu";
import ModalWindow from "@/components/UI/ProfilePage/ModalWindow/ModalWindow";

export default function Profile() {
    const drawerSideMenuRef = useRef<{ triggerHandleClick: () => void }>(null);

    const handleButtonClick = () => {
        if (drawerSideMenuRef.current) {
            drawerSideMenuRef.current.triggerHandleClick();
        }
    };

    return (
        <ThemeProvider theme={theme}>
            <body>
                <Header menuOpenFunction={handleButtonClick}/>
                <main>
                    {/*<SideMenu ref={sideMenuRef}/>*/}
                    <DrawerSideMenu ref={drawerSideMenuRef}/>
                    <ProfileComponent/>
                </main>
                <footer>
                    <ModalWindow/>
                </footer>
            </body>
        </ThemeProvider>
    );
}
