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

export const DrawerLogo = styled("div")`
    width: auto;
    height: calc(48rem/16);
    background: linear-gradient(45.00deg, rgb(248, 68, 79),rgb(120, 189, 196));
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
    font-weight: 400;
    line-height: calc(48rem/16);
    letter-spacing: 0;
    text-align: center;
`;

