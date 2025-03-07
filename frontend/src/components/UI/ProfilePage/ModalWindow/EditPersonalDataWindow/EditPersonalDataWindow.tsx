import React, {forwardRef, useImperativeHandle, useState} from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import TextField from '@mui/material/TextField';
import useUserInfoStore from "@/components/UI/ProfilePage/store";
import usePersonalDataStore, {PersonalData} from "@/components/UI/ProfilePage/ModalWindow/EditPersonalDataWindow/store";

interface EditPersonalDataProps {

}

const EditPersonalDataWindow = () => {
    const info = useUserInfoStore(state => state.info);
    const setFullName = useUserInfoStore(state => state.setFullName);
    const setDate = useUserInfoStore(state => state.setDate);
    const setContactInfo = useUserInfoStore(state => state.setContactInfo);
    const {isOpen, form, handleClose, handleInputChange, handleSubmit} = usePersonalDataStore();

    const setPersonalInfo = (data: PersonalData)=> {
        setFullName({name: data.name, surname: data.surname, patronymic: data.patronymic});
        setDate(data.date);
        setContactInfo({messenger: data.messenger, email: data.email});
    }

    return (
        <Dialog open={isOpen} onClose={handleClose}>
            <DialogTitle>Контактная информация</DialogTitle>
            <DialogContent>
                <TextField
                    autoFocus
                    margin="dense"
                    name="name"
                    label="Имя"
                    type="text"
                    fullWidth
                    value={form.name}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="surname"
                    label="Фамилия"
                    type="text"
                    fullWidth
                    value={form.surname}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="patronymic"
                    label="Отчество"
                    type="text"
                    fullWidth
                    value={form.patronymic}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="date"
                    label="Дата рождения"
                    type="text"
                    fullWidth
                    value={form.date}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="email"
                    label="Email"
                    type="email"
                    fullWidth
                    value={form.email}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="messenger"
                    label="Мессенджер"
                    type="text"
                    fullWidth
                    value={form.messenger}
                    onChange={handleInputChange}
                />
            </DialogContent>
            <DialogActions>
                <Button onClick={handleClose}>Отмена</Button>
                <Button onClick={() => {handleSubmit(setPersonalInfo)}} variant="contained" color="primary">
                    Сохранить
                </Button>
            </DialogActions>
        </Dialog>
    );
};

export default EditPersonalDataWindow;