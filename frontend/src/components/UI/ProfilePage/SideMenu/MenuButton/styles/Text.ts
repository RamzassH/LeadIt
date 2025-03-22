"use client"
import {Typography, styled} from '@mui/material';

const Text = styled(Typography)`
    margin-right: auto;
    text-transform: none;
    transition: color 0.3s ease;
    
    &.non-visible {
        visibility: hidden;
    }
`;

export default Text;