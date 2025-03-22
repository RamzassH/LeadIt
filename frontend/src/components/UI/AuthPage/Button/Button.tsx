import React from "react";
import {
    ButtonComponent,
    ButtonText,
    ButtonWrapper,
    GradientButton
} from "@/components/UI/AuthPage/Button/styled/Button";

interface ButtonProps {
    callback: (event: React.MouseEvent<HTMLButtonElement>) => void;
    style?: React.CSSProperties;
    children?: React.ReactNode;
    type?: "button" | "submit" | "reset"; // Добавляем тип для пропа `type`;
    isReverseBackground?: boolean;
}

export default function Button({
                                   callback,
                                   style,
                                   children,
                                   isReverseBackground = false,
                                   type = "button", // Значение по умолчанию для `type` — "button"
                               }: ButtonProps) {
    return (
        <GradientButton
            type={type} // Устанавливаем атрибут `type` на кнопку
            onClick={(event: React.MouseEvent<HTMLButtonElement>) => callback(event)}
            style={style}
            className={isReverseBackground ? "reverse-background": ""}
        >
            <ButtonText className={isReverseBackground ? "reverse-background": ""}>
                {children}
            </ButtonText>
        </GradientButton >
    );
}
