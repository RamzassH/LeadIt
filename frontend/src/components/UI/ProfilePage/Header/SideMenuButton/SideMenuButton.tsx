import {
    SideMenuButtonContainer,
    SideMenuButtonContent, SideMenuButtonStripe
} from "@/components/UI/ProfilePage/Header/SideMenuButton/styles/SideMenuButton";
import { motion } from "framer-motion";
import React, {useState} from "react";

interface SideMenuButtonProps {
    callback: () => void;
}

export default function SideMenuButton({callback}: SideMenuButtonProps) {
    const [isOpen, setIsOpen] = useState(false)

    const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        setIsOpen((prev) => !prev);
        callback()
    }

    return (
        <SideMenuButtonContainer onClick={handleClick}>
            <SideMenuButtonContent animate={{
                rotate: isOpen ? 0 : 0, // Поворот на 90 градусов или обратно
            }}
                transition={{
                duration: 0.3, // Длительность анимации
            }}>
                <SideMenuButtonStripe/>
                <SideMenuButtonStripe/>
                <SideMenuButtonStripe/>
            </SideMenuButtonContent>

        </SideMenuButtonContainer>
    )
}