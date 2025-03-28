import React, {ChangeEvent, useState} from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';

interface ModalWindowProps {
    open: boolean;
    handleClose: () => void;
    title?: string;
    children?: React.ReactNode;
    actions?: React.ReactNode;
}

function ModalWindow({open, handleClose, title, children, actions}: ModalWindowProps) {
    return (
        <Dialog open={open} onClose={handleClose}>
            <DialogTitle>{title}</DialogTitle>
            <DialogContent>
                {children}
            </DialogContent>
            <DialogActions>
                {actions}
            </DialogActions>
        </Dialog>
    );
}

export default ModalWindow;