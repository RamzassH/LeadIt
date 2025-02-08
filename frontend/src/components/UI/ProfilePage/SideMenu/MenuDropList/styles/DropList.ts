"use client"
import {styled} from '@mui/material';

const DropList = styled("div")`
    margin-top: 0.375rem;
    width: 100%;
    display: flex;
    flex-direction: column;
    padding: 0 0 0 2rem;
    overflow: hidden; /* Скрывает содержимое, когда DropList сжат */
    gap: 0.25rem;

    /* Устанавливаем начальную max-height для скрытого состояния */
    transition: max-height 0.3s ease-out; /* Анимация max-height */
`;

export default DropList;