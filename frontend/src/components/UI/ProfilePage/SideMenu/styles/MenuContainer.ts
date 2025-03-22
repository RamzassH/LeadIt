"use client"
import {Container, styled} from '@mui/material';

const MenuContainer = styled(Container)`
    /*background-color: ${({theme}) => theme.palette.dark?.main};*/
    background: linear-gradient(45.00deg, rgb(120, 189, 196),rgb(46, 99, 112) 33.515%,rgb(1, 44, 61) 100%);
    color: ${({theme}) => theme.palette.primary.main};
    width: 17rem;
    display: flex;
    flex-direction: column;
    box-shadow: none;
    justify-content: flex-start;
    margin-left: 0;
    gap: 0.375rem;
    
    &.close {
        width: calc(115rem / 16);
    }
`;

export default MenuContainer;