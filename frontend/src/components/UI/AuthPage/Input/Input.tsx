import React from "react";
import Error from "@/components/UI/AuthPage/Error/Error";
import {InputComponent, InputContainer, InputWrapper, Message} from "@/components/UI/AuthPage/Input/styled/Input";

interface InputProps {
    type: string;
    name: string;
    placeholder: string;
    value: string;
    onChange: (event: React.ChangeEvent<HTMLInputElement>) => void; // Теперь ожидаем onChange вместо setValue
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
                                  isError,
                                  errorMessage,
                              }: InputProps) {
    // Если тип 'tel', то используем стиль для ввода телефона
    let errorStyle = `${isError ? "error" : ""}`;

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
            target: { name: event.target.name, value: inputValue },
        } as React.ChangeEvent<HTMLInputElement>);
    };

    return (
        <InputContainer>
            <InputWrapper className={errorStyle}>
                <InputComponent className={errorStyle}
                                name={name}
                                type={type}
                                value={value}
                                onChange={type === "tel" ? handlePhoneChange : onChange} // Используем onChange из props или handlePhoneChange
                                placeholder={placeholder}
                />
            </InputWrapper>
            {isError && errorMessage ? (
                <Message>
                    <Error>{errorMessage}</Error>
                </Message>
            ) : null}
        </InputContainer>
    );
}
