import {Button, styled} from "@mui/material";

export const ProfileButtonStyle = styled(Button)`
    
`;

export const AvatarEditButton = styled(Button)`
    position: relative; /* Для позиционирования псевдоэлемента */
    width: calc(300rem/16);
    height: calc(48rem/16);
    padding: calc(10rem/16) calc(20rem/16);
    background: transparent; /* Прозрачный фон */
    border: none; /* Убираем стандартную рамку */
    border-radius: calc(10rem/16); /* Закругление углов */
    color: #ffffff; /* Цвет текста */
    font-size: calc(16rem/16);
    cursor: pointer;
    overflow: hidden; /* Чтобы псевдоэлемент не выходил за границы */

    /* Псевдоэлемент для градиентной рамки */
    &::before {
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;

        background: linear-gradient(225deg, rgb(248, 68, 79), rgb(120, 189, 196)); /* Градиент */
        border-radius: calc(10rem/16); /* Закругление углов */
        padding: calc(4rem/16); /* Толщина градиентной рамки */
        mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
        -webkit-mask-composite: xor; /* Обрезаем внутреннюю часть */
        mask-composite: exclude; /* Обрезаем внутреннюю часть */
    }

    &:hover {
        &::before {
            background: linear-gradient(225deg, rgb(255, 100, 110), rgb(140, 209, 216)); /* Измененный градиент при наведении */
        }
    }

    &.reverse-background::before {
        background: linear-gradient(45deg, rgb(248, 68, 79), rgb(120, 189, 196)); /* Градиент */
    }

    &.reverse-background:hover {
        &::before {
            background: linear-gradient(45deg, rgb(255, 100, 110), rgb(140, 209, 216)); /* Измененный градиент при наведении */
        }
    }
`;

export const AvatarEditButtonText = styled("div")`
    background: linear-gradient(225.00deg, rgb(248, 68, 79),rgb(184, 219, 220));
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
    text-fill-color: transparent;
    text-transform: none;
    font-family: Roboto, sans-serif;
    font-size: calc(24rem/16);
    font-weight: 500;
    line-height: calc(40rem/16);
    letter-spacing: 0;
    text-align: left;

    &.reverse-background {
        background: linear-gradient(45.00deg, rgb(248, 68, 79),rgb(184, 219, 220));
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
    }
`;
