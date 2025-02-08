import StyleMenuButton from "@/components/UI/ProfilePage/SideMenu/MenuButton/styles/MenuButton";
import React from "react";

interface ButtonProps {
    id?: string

    callback: (event: React.MouseEvent<HTMLButtonElement>) => void;
    children?: React.ReactNode;
    type?: "button" | "submit" | "reset"; // Добавляем тип для пропа `type`
}

export default function MenuButton({
                                       callback,
                                       children,
                                       type = "button", // Значение по умолчанию для `type` — "button"
}: ButtonProps) {
    return (
        <StyleMenuButton variant="contained" onClick={callback}>
            {children}
        </StyleMenuButton>
    )
}