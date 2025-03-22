import {styled} from "@mui/material";

export const InputContainer = styled("div")`
    /* InputWithLabel */
    position: static;
    width: 100%;
    height: calc(101rem/16);
    /* Автолейаут */
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    padding: 0 calc(15rem/16);
    /* Inside Auto Layout */
    flex: none;
`;

export const InputWrapper = styled("div")`
    position: relative;
    width: 100%;
    border-radius: calc(10rem/16);
    background: linear-gradient(225deg, rgb(248, 68, 79), rgb(120, 189, 196));
    padding: calc(4rem/16); /* Толщина градиентной рамки */
    
    /*transition: background 0.3s ease;*/
    
    &:hover {
        background: #4D90FE;
        outline: none;
        box-shadow: 0 0 calc(5rem/16) rgba(77, 144, 254, 0.5); /* Синяя тень */
    }

    &.error {
        background: red;
        box-shadow: 0 0 calc(5rem/16) rgba(255, 0, 0, 0.5); /* Красная тень при ошибке */
    }
`;

export const InputComponent = styled("input")`
    width: 100%;
    height: calc(52rem/16);

    /* Автолейаут */
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;

    /* Inside Auto Layout */
    margin: 0;
    padding: calc(4rem/16) calc(7rem/16);

    background: rgb(1, 34, 47); /* Основной фон */
    border: none; /* Убираем рамку */
    border-radius: calc(6rem/16); /* Закругление углов (меньше, чем у обертки) */

    color: rgb(247, 248, 243);
    font-family: Roboto, sans-serif;
    font-size: calc(24rem/16);
    font-weight: 500;
    line-height: calc(44rem/16);
    letter-spacing: 0;
    text-align: left;
    
    /*
    &:hover {
        border: calc(2rem/16) solid #4D90FE;
        outline: none;
        box-shadow: 0 0 calc(5rem/16) rgba(77, 144, 254, 0.5); /* Синяя тень 
    }

    &.error {
        border-color: red;
        box-shadow: 0 0 calc(5rem/16) rgba(255, 0, 0, 0.5); /* Красная тень при ошибке 
    }
    */
`;

export const Message = styled("div")`
    position: relative;
    top: calc(-6rem/16);
    margin: 0 calc(8rem/16);
`;