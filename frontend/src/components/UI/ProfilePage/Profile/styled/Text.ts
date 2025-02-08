import {styled} from "@mui/material";

export const Text = styled("div")`
    color: ${({theme}) => theme.palette.dark?.main};
    font-family: Roboto, sans-serif;
    font-weight: 100;
    font-size: 20px;
    height: fit-content;
    letter-spacing: 0;
    text-align: left;
`;

