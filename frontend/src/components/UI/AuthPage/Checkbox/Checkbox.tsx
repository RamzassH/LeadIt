import React, { useState } from "react";
import styles from "./Checkbox.module.css"

// Интерфейс для пропсов компонента
interface CheckboxProps {
    label: string; // Метка для чекбокса
    checked?: boolean; // Начальное состояние (по умолчанию false)
    onChange?: (checked: boolean) => void; // Функция, которая вызывается при изменении состояния
    classStyles?: string[]; // Кастомные стили
}

export default function Checkbox({ label, checked = false, onChange, classStyles }: CheckboxProps)  {
    const [isChecked, setIsChecked] = useState(checked);

    let componentStyle = styles["checkbox-container"] + " " + (classStyles?.join(" ") || "");

    // Обработчик изменения состояния
    const handleChange = () => {
        const newCheckedState = !isChecked;
        setIsChecked(newCheckedState);
        if (onChange) {
            onChange(newCheckedState); // Если передан onChange, вызываем его
        }
    };

    return (
        <label className={componentStyle}>
            <input
                type="checkbox"
                checked={isChecked}
                onChange={handleChange}
                className={styles["checkbox-input"]}
            />
            <div className={styles["checkbox-label"]}>
                {label}
            </div>
        </label>
    );
};
