"use client"
import {Container, styled} from '@mui/material';

const LogoContainer = styled("div")`
    position: static;
    width: calc(45rem / 16);
    height: calc(45rem / 16);
    
    margin: 0 calc(16rem / 16);
    border-radius: calc(8rem/ 16);
    background: ${({theme}) => theme.palette.primary.main};
    color: ${({theme}) => theme.palette.dark?.main};
    
    display: flex;
    justify-content: center;
    align-items: center;
    
    svg {
        position: relative;
        width: calc(27.26rem / 16);
        height: calc(34.17rem / 16);
        path {
            fill: currentColor;
        }
    }
`;

export default LogoContainer;