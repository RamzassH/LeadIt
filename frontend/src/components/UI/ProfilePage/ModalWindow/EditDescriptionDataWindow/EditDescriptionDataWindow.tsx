import React, {forwardRef, useImperativeHandle, useState} from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import TextField from '@mui/material/TextField';
import useUserInfoStore from "@/components/UI/ProfilePage/store";
import useDescriptionStore from "@/components/UI/ProfilePage/ModalWindow/EditDescriptionDataWindow/store";

const EditDescriptionDataWindow = () => {
    const info = useUserInfoStore(state => state.info);
    const setDescription = useUserInfoStore(state => state.setDescription);
    const {isOpen, form, handleClose, handleInputChange, handleSubmit} = useDescriptionStore();

    return (
        <Dialog open={isOpen} onClose={handleClose}>
            <DialogTitle>О себе</DialogTitle>
            <DialogContent style={{width: "calc(400rem/16)", height: "fit-content"}} >
                <TextField
                    margin="dense"
                    name="description"
                    label="дада"
                    type="text"
                    fullWidth
                    multiline
                    rows={8}
                    value={form.description}
                    onChange={handleInputChange}
                />
            </DialogContent>
            <DialogActions>
                <Button onClick={handleClose}>Отмена</Button>
                <Button onClick={() => {handleSubmit(setDescription)}} variant="contained" color="primary">
                    Сохранить
                </Button>
            </DialogActions>
        </Dialog>
    );
};

export default EditDescriptionDataWindow;