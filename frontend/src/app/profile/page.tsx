"use client"
import "./local.css"
import MenuButton from "@/components/UI/ProfilePage/SideMenu/MenuButton/MenuButton";
import theme from "../../../theme/theme";
import {margin, ThemeProvider} from "@mui/system";
import SideMenu from "@/components/UI/ProfilePage/SideMenu/SideMenu";
import React, {useRef} from "react";
import Header from "@/components/UI/ProfilePage/Header/Header";
import ProfileComponent from "@/components/UI/ProfilePage/Profile/ProfileComponent";

export default function Profile() {
    const sideMenuRef = useRef<{ triggerHandleClick: () => void }>(null);

    const handleButtonClick = () => {
        if (sideMenuRef.current) {
            sideMenuRef.current.triggerHandleClick();
        }
    };

    return (
        <ThemeProvider theme={theme}>
            <body>
                <Header menuOpenFunction={handleButtonClick}/>
                <main>
                    <SideMenu ref={sideMenuRef}/>
                    <ProfileComponent/>
                </main>
                <footer>
                </footer>
            </body>
        </ThemeProvider>
    );
}
