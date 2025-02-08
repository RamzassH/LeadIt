"use client"
import {Button, styled} from '@mui/material';
import {motion} from "framer-motion";
import { Box } from '@mui/material';

export const SideMenuButtonContainer = styled(Button)`
    background-color: ${({theme}) => theme.palette.dark?.main};
    color: ${({theme}) => theme.palette.primary.main};
    min-width: 0;
    width: calc(53rem / 16);
    height: calc(52rem/ 16);
    box-shadow: none;
`;

export const SideMenuButtonContent = styled(motion(Box))`
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    
    justify-content: flex-start;
`;

export const SideMenuButtonStripe = styled("div")`
    /* Прямоугольник 4 */
    position: static;
    width: 100%;
    height: calc(4rem / 16);
    /* Inside Auto Layout */
    flex: none;
    order: 0;
    flex-grow: 0;
    margin: calc(5rem / 16) 0;

    background: currentColor;
`;