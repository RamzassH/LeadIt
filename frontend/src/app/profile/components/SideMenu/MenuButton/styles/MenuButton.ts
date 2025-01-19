"use client"
import {Button, styled} from '@mui/material';

const MenuButton = styled(Button)`
    background-color: ${({theme}) => theme.palette.dark?.main};
    height: fit-content;

    &:hover {
        background-color: ${({theme}) => theme.palette.primary.main};
    }
`;

export default MenuButton;