import React, { useState } from "react";
import {CheckboxContainer, CheckboxInput, CheckboxLabel} from "@/components/UI/AuthPage/Checkbox/styled/Checkbox";

// Интерфейс для пропсов компонента
interface CheckboxProps {
    label: string; // Метка для чекбокса
    checked?: boolean; // Начальное состояние (по умолчанию false)
    onChange?: (checked: boolean) => void; // Функция, которая вызывается при изменении состояния
    style?: React.CSSProperties;
}

export default function Checkbox({ label, checked = false, onChange, style}: CheckboxProps)  {
    const [isChecked, setIsChecked] = useState(checked);

    // Обработчик изменения состояния
    const handleChange = () => {
        const newCheckedState = !isChecked;
        setIsChecked(newCheckedState);
        if (onChange) {
            onChange(newCheckedState); // Если передан onChange, вызываем его
        }
    };

    return (
        <CheckboxContainer style={style}>
            <CheckboxInput
                className="checkbox-input"
                type="checkbox"
                checked={isChecked}
                onChange={handleChange}
            />
            <CheckboxLabel className="checkbox-label">
                {label}
            </CheckboxLabel>
        </CheckboxContainer>
    );
};
