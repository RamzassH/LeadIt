"use client"
import theme from "../../../theme/theme";
import {ThemeProvider} from "@mui/system";
import React, {useRef} from "react";
import Header from "@/components/UI/ProfilePage/Header/Header";
import ProfileComponent from "@/components/UI/ProfilePage/Profile/ProfileComponent";
import DrawerSideMenu from "@/components/UI/ProfilePage/DrawerSideMenu/DrawerSideMenu";
import {Body, Footer, Main} from "@/app/profile/styled";

export default function Profile() {
    const drawerSideMenuRef = useRef<{ triggerHandleClick: () => void }>(null);

    const handleButtonClick = () => {
        if (drawerSideMenuRef.current) {
            drawerSideMenuRef.current.triggerHandleClick();
        }
    };

    return (
        <ThemeProvider theme={theme}>
            <Body>
                <Header menuOpenFunction={handleButtonClick}/>
                <Main>
                    {/*<SideMenu ref={sideMenuRef}/>*/}
                    <DrawerSideMenu ref={drawerSideMenuRef}/>
                    <ProfileComponent/>
                </Main>
                <Footer>
                </Footer>
            </Body>
        </ThemeProvider>
    );
}
