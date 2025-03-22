"use client"
import {Icon, styled} from "@mui/material";

export const LastIcon = styled(Icon)`
    display: flex;
    justify-content: center;
    align-items: center;
    
    path {
        transition: stroke 0.3s ease;
    }
    
    &.non-visible {
        visibility: hidden;
    }
`;

export const SvgStyle = styled("svg")`
    /* Left (DropList-1) */

    width: calc(16rem / 16);
    height: calc(16rem / 16);
    /* Inside Auto Layout */
    flex: none;
    flex-grow: 0;
    margin: 0 0;
`;
