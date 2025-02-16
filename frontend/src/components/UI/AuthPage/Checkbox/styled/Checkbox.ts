import {styled} from "@mui/material";

export const CheckboxContainer = styled("label")`
    display: flex;
    width: 100%;
    align-items: center;
    font-family: Arial, sans-serif;
    cursor: pointer;
    user-select: none; /* Отключаем выделение текста */
    margin-bottom: calc(8rem/16);

    /* Стили при наведении на чекбокс */
    &:hover .checkbox-input {
        border-color: #1481cf;
    }

    &:hover .checkbox-label {
        color: #1481cf;
    }

`;

export const CheckboxInput = styled("input")`
    width: calc(20rem/16);
    height: calc(20rem/16);
    border: calc(2rem/16) solid #ccc;
    border-radius: calc(4rem/16);
    margin-right: calc(6rem/16);
    appearance: none;
    background-color: white;
    position: relative;
    transition: background-color 0.3s, border-color 0.3s;
    cursor: pointer;
    box-sizing: border-box; /* Включаем бордеры в размеры */
    flex-shrink: 0; /* Не позволяем сжимать чекбокс */
    
    &:checked {
        background: linear-gradient(45deg, #1481cf, #6bc0f3);
        /*border-color: #2196F3;*/
        border: 0;
    }

    &:checked::after {
        content: "";
        position: absolute;
        top: calc(6rem/16);
        left: calc(5rem/16);
        width: calc(10rem/16);
        height: calc(6rem/16);
        border: solid calc(2rem/16) white;
        border-top: none;
        border-right: none;
        transform: rotate(-45deg);
        animation: checkmark 0.3s ease forwards;
    }

    @keyframes checkmark {
        0% {
            width: 0;
            height: 0;
            transform: rotate(-45deg);
        }
        100% {
            width: calc(10rem/16);
            height: calc(6rem/16);
            transform: rotate(-45deg);
        }
    }
`;

export const CheckboxLabel = styled("div")`
    font-size: calc(16rem/16);
    color: #989898;
    transition: color 0.2s;
    box-sizing: border-box; /* Включаем бордеры в размеры */
`;