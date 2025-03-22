import {styled} from "@mui/material";

export const ContainerColumnCenter = styled("div")`
    width: 100%;
    height: 100%;
    
    display: flex;
    justify-content: center;
    align-items: center;
    
    flex-direction: column;
`;

export const ContainerColumn = styled("div")`
    width: 100%;
    height: 100%;
    
    display: flex;
    flex-direction: column;
`;

export const ContainerRow = styled("div")`
    width: 100%;
    height: 100%;

    display: flex;
    flex-direction: row;
`;


export const BackgroundContainer = styled("div")`
    width: 100%;
    height: 100%;
    display: flex;
    background-color: #254A58; /*rgba(120, 188, 196, 0.3);*/
    border-radius: calc(20rem/16);
`;

export const Background = styled("div")`
    width: 100%;
    height: 100%;

    background-color: #012C3D; /*rgba(120, 188, 196, 0.3);*/
    border-radius: calc(20rem/16);
`;