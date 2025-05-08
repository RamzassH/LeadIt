'use client';
import GlobalStyle from "@/app/styled";
import {metadata} from "./metadata"
import useGlobalStore from "@/store/GlobalStore/store";
import {useEffect} from "react";
import {getLocalStorage, setLocalStorage} from "@/utils/cookie";

export default function RootLayout({
                                       children,
                                   }: Readonly<{
    children: React.ReactNode;
}>) {
    const globalStore = useGlobalStore();

    useEffect(() => {
        if (getLocalStorage("isAuth") === null) {
            setLocalStorage("isAuth", false);
        }
        if (getLocalStorage("refreshToken") === null) {
            setLocalStorage("refreshToken", "");
        }
        globalStore.setLogin(getLocalStorage("isAuth") as boolean);
        globalStore.setRefreshToken(getLocalStorage("refreshToken") as string);
    }, []);

    useEffect(() => {
        setLocalStorage("isAuth", globalStore.isLogin);
    }, [globalStore.isLogin]);
    useEffect(() => {
        setLocalStorage("refreshToken", globalStore.refreshToken);
    }, [globalStore.refreshToken]);

    return (
        <html lang="en">
        <GlobalStyle/>
        <body>
        {children}
        </body>
        </html>
    );
}
