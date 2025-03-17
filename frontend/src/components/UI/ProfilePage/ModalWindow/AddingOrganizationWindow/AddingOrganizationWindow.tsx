import React from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import TextField from '@mui/material/TextField';
import useAddingOrganizationStore from "@/components/UI/ProfilePage/ModalWindow/AddingOrganizationWindow/store";


const AddingOrganizationWindow = () => {
    const {isOpen, form, handleClose, handleInputChange, handleSubmit} = useAddingOrganizationStore();

    return (
        <Dialog open={isOpen} onClose={handleClose}>
            <DialogTitle>Добавить организацию</DialogTitle>
            <DialogContent style={{width: "calc(400rem/16)", height: "fit-content"}} >
                <TextField
                    margin="dense"
                    name="description"
                    label="Наименование организации"
                    type="text"
                    fullWidth
                    value={form.name}
                    onChange={handleInputChange}
                />
            </DialogContent>
            <DialogActions>
                <Button onClick={handleClose}>Отмена</Button>
                <Button onClick={() => {handleSubmit((data) => {console.log(data)})}} variant="contained" color="primary">
                    Сохранить
                </Button>
            </DialogActions>
        </Dialog>
    );
};

export default AddingOrganizationWindow;