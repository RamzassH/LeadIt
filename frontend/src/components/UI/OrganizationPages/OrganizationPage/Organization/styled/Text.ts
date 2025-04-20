import {styled} from "@mui/material";

export const Text = styled("div")`
    color: ${({theme}) => theme.palette.secondary?.main};
    font-family: Roboto, sans-serif;
    font-style: italic;
    font-weight: 400;
    font-size: calc(16rem/16);
    line-height: calc(19rem/16);
    letter-spacing: 0;
    text-align: left;
    
    &.text {
        color: ${({theme}) => theme.palette.background?.default};
        font-weight: 700;  
        font-style: normal;
    }
`;

