import React from "react";
import {ProfileButtonStyle} from "@/components/UI/ProfilePage/Profile/ProfileButton/styled/ProfileButton";

interface ProfileButtonProps {
    id?: string
    icon?: React.ReactNode;
    text?: string;
    callback: (event: React.MouseEvent<HTMLButtonElement>) => void;
    children?: React.ReactNode;
    type?: "button" | "submit" | "reset"; // Добавляем тип для пропа `type`
}

export default function ProfileButton({id, icon, text, callback, children}:ProfileButtonProps) {
    return (
        <ProfileButtonStyle id={id} onClick={callback}>
            {icon}
            {text}
            {children}
        </ProfileButtonStyle>
    )
}