import {Button, styled} from "@mui/material";

export const ProfileButtonStyle = styled(Button)`
    
`;

export const AvatarEditButton = styled(Button)`
    
    background-color: ${({theme}) => theme.palette.accent?.main};;
    color: ${({theme}) => theme.palette.background?.default};;
`;
