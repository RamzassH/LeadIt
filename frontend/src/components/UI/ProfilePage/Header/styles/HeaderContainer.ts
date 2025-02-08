"use client"
import {Container, styled} from '@mui/material';

const HeaderContainer = styled("header")`
    background-color: ${({theme}) => theme.palette.dark?.main};
    color: ${({theme}) => theme.palette.primary.main};
    position: static;
    width: 100%;
    height: calc(65rem / 16);
    /* Автолейаут */
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    padding: calc(8rem / 16) calc(12rem / 16);


    /* Inside Auto Layout */
    flex: none;
    margin: 0 0;
    
    border-bottom: 1px solid rgb(1, 44, 61);
    border-radius: 0 0 calc(4rem / 16) 0;
    
`;

export default HeaderContainer;