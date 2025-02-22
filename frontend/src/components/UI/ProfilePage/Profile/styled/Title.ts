import {styled} from "@mui/material";

export const Title = styled("div")`
    color: ${({theme}) => theme.palette.primary?.main};
    font-family: Roboto, sans-serif;
    font-weight: 700; 
    /*font-style: italic;*/
    font-size: calc(26rem/16);
    line-height: calc(30rem/16);
    letter-spacing: 0;
    text-align: left;
`;

export const ChangeButton = styled("div")`
    background: linear-gradient(225.00deg, rgb(248, 68, 79),rgb(120, 189, 196));
    -webkit-background-clip:
            text;
    -webkit-text-fill-color:
            transparent;
    background-clip:
            text;
    text-fill-color:
            transparent;
    font-family: Roboto,sans-serif;
    font-size: calc(20rem/16);
    font-weight: 400;
    line-height: calc(20rem/16);
    letter-spacing: 0;
    text-align: right;
    cursor: pointer;
    height: fit-content;
    transition: color 0.3s ease;
    &:hover {
        color: ${({theme}) => theme.palette.background?.default};
    }
`;
