"use client"
import React, {useRef} from "react";
import ProfileComponent from "@/components/UI/ProfilePage/Profile/ProfileComponent";

export default function Profile() {
    const drawerSideMenuRef = useRef<{ triggerHandleClick: () => void }>(null);

    const handleButtonClick = () => {
        if (drawerSideMenuRef.current) {
            drawerSideMenuRef.current.triggerHandleClick();
        }
    };

    return (
        <ProfileComponent/>
    );
}
