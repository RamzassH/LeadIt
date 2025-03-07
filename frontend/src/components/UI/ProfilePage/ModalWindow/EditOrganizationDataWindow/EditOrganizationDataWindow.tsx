import React, {forwardRef, useImperativeHandle, useState} from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import TextField from '@mui/material/TextField';
import useUserInfoStore from "@/components/UI/ProfilePage/store";
import useProjectInfoStore from "@/components/UI/ProfilePage/ModalWindow/EditOrganizationDataWindow/store";

interface EditOrganizationDataProps {

}

const EditOrganizationDataWindow = () => {
    const info = useUserInfoStore(state => state.info);
    const setProjectInfo = useUserInfoStore(state => state.setProjectInfo);
    const {isOpen, form, handleClose, handleInputChange, handleSubmit} = useProjectInfoStore();

    return (
        <Dialog open={isOpen} onClose={handleClose}>
            <DialogTitle>Проектная деятельность</DialogTitle>
            <DialogContent>
                <TextField
                    autoFocus
                    margin="dense"
                    name="organization"
                    label="Организация"
                    type="text"
                    fullWidth
                    value={form.organization}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="projects"
                    label="Проекты"
                    type="text"
                    fullWidth
                    value={form.projects}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="position"
                    label="Должность"
                    type="text"
                    fullWidth
                    value={form.position}
                    onChange={handleInputChange}
                />
            </DialogContent>
            <DialogActions>
                <Button onClick={handleClose}>Отмена</Button>
                <Button onClick={() => {handleSubmit(setProjectInfo)}} variant="contained" color="primary">
                    Сохранить
                </Button>
            </DialogActions>
        </Dialog>
    );
};

export default EditOrganizationDataWindow;