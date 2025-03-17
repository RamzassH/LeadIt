import {styled} from "@mui/material";

export const Container = styled("div")`
    display: flex;
    flex-direction: row;
    width: 100%;
    height: 100%;

    align-items: center;
    background: linear-gradient(90.00deg, rgb(120, 189, 196),rgb(247, 248, 243) 68.189%);;
`;

export const BackgroundContainer = styled("div")`
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    
    width: 50%;
    background: linear-gradient(132.54deg, rgb(9, 54, 70) 12.848%,rgb(35, 85, 99) 57.839%,rgb(56, 111, 123) 75.084%,rgb(85, 147, 156) 89.324%,rgb(120, 189, 196) 100%);;
    border-radius: 0 calc(180rem/16) calc(180rem/16) 0;

    flex: none;
    order: 0;
    align-self: stretch;
    flex-grow: 0;
    margin: 0 calc(220rem/16) 0 0;
`;

export const Content = styled("div")`
    /* Регистрация */
    position: static;
    width: calc(480rem/16);
    height: calc(800rem/16);
    /* Автолейаут */
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

`;

export const TitleContainer = styled("div")`
    /* Шапка регистрации */
    position: static;
    width: 100%;
    height: calc(120rem/16);
    /* Автолейаут */
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 0;
    
`;

export const Title = styled("div")`
    width: fit-content;

    background: linear-gradient(225.00deg, rgb(248, 68, 79),rgb(184, 219, 220));
    -webkit-background-clip:
            text;
    -webkit-text-fill-color:
            transparent;
    background-clip:
            text;
    text-fill-color:
            transparent;
    font-family: Roboto, sans-serif;
    font-size: calc(36rem/16);
    font-weight: 500;
    line-height: calc(60rem/16);
    letter-spacing: 0;
    text-align: left;
`;

export const Form = styled("form")`
    display: flex;
    flex-direction: column; /* Располагаем элементы вертикально */
    width: 100%; /* Форма занимает всю доступную ширину */
`;

export const ButtonContainer = styled("div")`
    position: static;
    width: calc(369rem/16);
    height: calc(68rem/16);
    /* Автолейаут */
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
`;