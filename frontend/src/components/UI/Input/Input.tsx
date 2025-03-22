import styles from "./Input.module.css";
import React from "react";
import Error from "@/components/UI/Error/Error";

interface InputProps {
    type: string;
    name: string;
    placeholder: string;
    value: string;
    onChange: (event: React.ChangeEvent<HTMLInputElement>) => void; // Теперь ожидаем onChange вместо setValue
    classStyles?: string[];
    style?: React.CSSProperties;
    isError?: boolean;
    errorMessage?: string;
}

export default function Input({
                                  type,
                                  name,
                                  placeholder,
                                  value,
                                  onChange,
                                  classStyles,
                                  isError,
                                  errorMessage,
                              }: InputProps) {
    // Если тип 'tel', то используем стиль для ввода телефона
    let componentStyle =
        styles.input +
        ` ${isError ? styles.error + ` ` : ""}` +
        (classStyles?.join(" ") || "");

    // Функция для маскирования номера телефона (пример для формата +1 (234) 567-8901)
    const handlePhoneChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        let inputValue = event.target.value;

        // Удаляем все нецифровые символы
        inputValue = inputValue.replace(/\D/g, "");

        // Форматируем номер телефона
        if (inputValue.length <= 1) {
            inputValue = `+${inputValue}`;
        } else if (inputValue.length <= 4) {
            inputValue = `+${inputValue.slice(0, 1)} (${inputValue.slice(1)}`;
        } else if (inputValue.length <= 7) {
            inputValue = `+${inputValue.slice(0, 1)} (${inputValue.slice(1, 4)}) ${inputValue.slice(4)}`;
        } else if (inputValue.length <= 9) {
            inputValue = `+${inputValue.slice(0, 1)} (${inputValue.slice(1, 4)}) ${inputValue.slice(4, 7)}-${inputValue.slice(7)}`;
        } else {
            inputValue = `+${inputValue.slice(0, 1)} (${inputValue.slice(1, 4)}) ${inputValue.slice(4, 7)}-${inputValue.slice(7, 9)}-${inputValue.slice(9, 11)}`;
        }

        // Передаем измененное значение в onChange
        onChange({
            target: {name: event.target.name, value: inputValue},
        } as React.ChangeEvent<HTMLInputElement>);
    };

    return (
        <div>
            <input
                className={componentStyle}
                name={name}
                type={type}
                value={value}
                onChange={type === "tel" ? handlePhoneChange : onChange} // Используем onChange из props или handlePhoneChange
                placeholder={placeholder}
            />
            {isError && errorMessage ? (
                <div className={styles.message}>
                    <Error>{errorMessage}</Error>
                </div>
            ) : null}
        </div>
    );
}
