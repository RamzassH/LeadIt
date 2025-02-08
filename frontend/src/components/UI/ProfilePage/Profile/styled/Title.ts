import {styled} from "@mui/material";

export const Title = styled("div")`
    color: ${({theme}) => theme.palette.dark?.main};
    font-family: Roboto, sans-serif;
    font-weight: 300; /* Light */
    font-style: italic;
    font-size: 24px;
    line-height: 28px;
    letter-spacing: 0;
    text-align: left;
`;

export const ChangeButton = styled("div")`
    color: ${({theme}) => theme.palette.accent?.main};
    font-family: Roboto,sans-serif;
    font-size: 20px;
    font-weight: 300;
    line-height: 23px;
    letter-spacing: 0;
    text-align: right;
    cursor: pointer;
    height: fit-content;
    transition: color 0.3s ease;
    &:hover {
        color: ${({theme}) => theme.palette.background?.default};
    }
`;
