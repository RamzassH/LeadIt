"use client"
import {ThemeProvider} from "@mui/system";
import theme from "../../../../theme/theme";
import Header from "@/components/UI/ProfilePage/Header/Header";
import DrawerSideMenu from "@/components/UI/ProfilePage/DrawerSideMenu/DrawerSideMenu";
import ProfileComponent from "@/components/UI/ProfilePage/Profile/ProfileComponent";
import ModalWindow from "@/components/UI/ProfilePage/ModalWindow/ModalWindow";
import React, {useRef} from "react";
import Graph from "@/components/UI/OrganizationPages/StructurePage/Graph/Graph";

export default function Page() {
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
            <main style={{alignItems: "center", justifyContent: "center", display: "flex", flexDirection: "column", height: "100vh"}}>
                <DrawerSideMenu ref={drawerSideMenuRef}/>
                <Graph/>
            </main>
            <footer>
            </footer>
            </body>
        </ThemeProvider>)
}