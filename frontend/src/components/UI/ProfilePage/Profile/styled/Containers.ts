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
    
    background-color: rgba(120, 188, 196, 0.3);
    border-radius: calc(8rem/16);
`;