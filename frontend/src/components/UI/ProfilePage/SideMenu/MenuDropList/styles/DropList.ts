"use client"
import {styled} from '@mui/material';

export const DropList = styled("div")`
    /*margin-top: 0.375rem;*/
    width: 100%;
    display: flex;
    flex-direction: column;
    /*padding: 0 0 0 2rem;*/
    overflow: hidden; /* Скрывает содержимое, когда DropList сжат */
    /*gap: 0.25rem;*/
    align-items: center;
    /* Устанавливаем начальную max-height для скрытого состояния */
    transition: max-height 0.3s ease-out; /* Анимация max-height */
`;

export const DropListBackground = styled("div")`
    background-color: ${({theme}) => theme.palette.dark?.main};
    width: 80%;
    height: 100%;
    padding: calc(6rem/16) calc(8rem/16);
    border-radius: 0 0 calc(15rem/16) calc(15rem/16);
`;