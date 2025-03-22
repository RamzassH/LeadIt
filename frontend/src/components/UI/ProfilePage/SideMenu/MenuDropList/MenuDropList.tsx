"use client";
import React, {useEffect, useRef, useState} from "react";
import { motion } from "framer-motion"; // Импортируем motion из Framer Motion
import MenuButton from "@/components/UI/ProfilePage/SideMenu/MenuButton/MenuButton";
import Text from "@/components/UI/ProfilePage/SideMenu/MenuButton/styles/Text";
import MenuDropListContainer from "@/components/UI/ProfilePage/SideMenu/MenuDropList/styles/MenuDropListContainer";
import {DropList, DropListBackground} from "@/components/UI/ProfilePage/SideMenu/MenuDropList/styles/DropList";
import {LastIcon} from "@/components/UI/ProfilePage/SideMenu/MenuDropList/styles/LastIcon";
import SvgIcon from "@/components/UI/ProfilePage/SideMenu/MenuDropList/svg/SvgIcon";
import Icon from "@/components/UI/ProfilePage/SideMenu/MenuButton/styles/StyleIcon";
import Button from "@/components/UI/AuthPage/Button/Button";
import CreateButton from "@/components/UI/ProfilePage/SideMenu/MenuDropList/CreateButton/CreateButton";


interface MenuDropListProps {
    firstIcon?: React.ReactNode;
    text?: string;
    visibility?: string;
    children?: React.ReactNode;
}

export default function MenuDropList({
    firstIcon, text, children, visibility = ""
                                     }: MenuDropListProps) {
    const [isOpen, setIsOpen] = useState(false);
    const dropListRef = useRef<HTMLDivElement | null>(null); // Реф для получения высоты
    // Состояние для высоты DropList
    const [dropListHeight, setDropListHeight] = useState(0);

    const handleClick = () => {
        if (visibility == "non-visible") {
            return
        }
        setIsOpen((prev) => !prev);
    };

    // Эффект для получения высоты компонента после рендера
    useEffect(() => {
        if (dropListRef.current) {
            setDropListHeight(dropListRef.current.scrollHeight + 6); // Получаем высоту компонента
        }
    }, []); // Обновляем высоту при каждом изменении состояния isOpen

    useEffect(() => {
        if (visibility == "non-visible" && isOpen) {
            setIsOpen(false)
        }
    }, [visibility]);

    return (
        <MenuDropListContainer>
            <MenuButton callback={handleClick}>
                <Icon className="first-icon">
                    {firstIcon}
                </Icon>
                <Text className={visibility}>{text}</Text>
                <motion.span
                    animate={{
                        rotate: isOpen ? -90 : 0, // Поворот на 90 градусов или обратно
                    }}
                    transition={{
                        duration: 0.3, // Длительность анимации
                    }}
                >
                    <LastIcon className={"last-icon " + visibility}>
                        <SvgIcon/>
                    </LastIcon>
                </motion.span>
            </MenuButton>

            {/* Используем motion.div для анимации появления/исчезновения */}
            <motion.div
                initial={{ opacity: 0, maxHeight: 0 }} // Начальное состояние (невидимый и с максимальной высотой 0)
                animate={{
                    opacity: isOpen ? 1 : 0, // Плавное изменение прозрачности
                    maxHeight: isOpen ? dropListHeight : 0, // Плавный сдвиг по высоте
                }}
                exit={{
                    opacity: 0, // Убираем прозрачность при скрытии
                    maxHeight: 0, // Сжимаем до 0 высоты при закрытии
                }}
                transition={{ duration: 0.3 }} // Длительность анимации
                style={{ overflow: "hidden" }} // Управляем overflow для скрытия содержимого при сжатии
            >
                <DropList ref={dropListRef}>
                    {children ?
                        children :
                        <CreateButton/>
                    }
                </DropList>
            </motion.div>
        </MenuDropListContainer>
    );
}
