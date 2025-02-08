"use client"
import {Container, styled} from '@mui/material';

const MenuContainer = styled(Container)`
    background-color: ${({theme}) => theme.palette.dark?.main};
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