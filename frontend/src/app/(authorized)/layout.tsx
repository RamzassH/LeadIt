"use client"
import React, {ReactNode, useRef} from 'react';
import {ThemeProvider} from "@mui/system";
import theme from "../../../theme/theme";
import Header from "@/components/UI/ProfilePage/Header/Header";
import DrawerSideMenu from "@/components/UI/DrawerSideMenu/DrawerSideMenu";
import {Body, Main} from "@/app/(authorized)/styled";
// Типы для пропсов
interface DashboardLayoutProps {
    children: ReactNode;
}

export default function AuthorizedLayout({ children }: DashboardLayoutProps) {
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
                <DrawerSideMenu ref={drawerSideMenuRef}/>
                <Main>
                    {children}
                </Main>
                <footer>
                </footer>
            </Body>
        </ThemeProvider>
    );
}