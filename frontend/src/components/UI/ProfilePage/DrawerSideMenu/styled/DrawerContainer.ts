"use client"
import {Container, Drawer, styled} from '@mui/material';

export const DrawerContainer = styled(Drawer)`
    width: 17rem;
    display: flex;
    box-shadow: none;
`;

export const DrawerBackground = styled("div")`
    background: linear-gradient(45.00deg, rgb(120, 189, 196),rgb(46, 99, 112) 33.515%,rgb(1, 44, 61) 100%);
    width: 17rem;
    height: 100%;
    display: flex;
    flex-direction: column;
    box-shadow: none;
    justify-content: flex-start;
    margin-left: 0;
    gap: 0.375rem;
    padding: calc(20rem/16) calc(32rem/16);
`;
