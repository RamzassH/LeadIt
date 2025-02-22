"use client"
import {Button, styled} from '@mui/material';

const StyleMenuButton = styled(Button)`
    background-color: ${({theme}) => theme.palette.dark?.main};
    color: ${({theme}) => theme.palette.primary.main};
    height: fit-content;
    width: 100%;
    display: flex;
    box-shadow: none;
    justify-content: flex-start;
    
    border-radius: calc(15rem/16);
    
    &.list-item {
        border-radius: 0;
        width: 88%;
    }
    
    &.list-item:last-child {
        border-radius: 0 0 calc(15rem/16) calc(15rem/16);
    }
    
    &:hover {
        background-color: ${({theme}) => theme.palette.primary.main};
        color: ${({theme}) => theme.palette.dark?.main};
        .first-icon path {
            fill: ${({theme}) => theme.palette.dark?.main};
        }
        .last-icon path {
            stroke: ${({theme}) => theme.palette.dark?.main};
        }
    }
`;

export default StyleMenuButton;