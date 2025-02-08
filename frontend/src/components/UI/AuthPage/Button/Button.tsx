import styles from "./Button.module.css";
import React from "react";

interface ButtonProps {
    callback: (event: React.MouseEvent<HTMLButtonElement>) => void;
    classStyles?: string[];
    style?: React.CSSProperties;
    children?: React.ReactNode;
    type?: "button" | "submit" | "reset"; // Добавляем тип для пропа `type`
}

export default function Button({
                                   callback,
                                   classStyles,
                                   style,
                                   children,
                                   type = "button", // Значение по умолчанию для `type` — "button"
                               }: ButtonProps) {
    let componentStyle = styles.button + " " + (classStyles?.join(" ") || "");

    return (
        <button
            type={type} // Устанавливаем атрибут `type` на кнопку
            className={componentStyle}
            onClick={(event: React.MouseEvent<HTMLButtonElement>) => callback(event)}
            style={style}
        >
            {children}
        </button>
    );
}
