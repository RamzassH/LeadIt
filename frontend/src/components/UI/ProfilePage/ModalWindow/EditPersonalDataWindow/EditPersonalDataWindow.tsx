import React, {forwardRef, useImperativeHandle, useState} from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import TextField from '@mui/material/TextField';
import useUserInfoStore from "@/components/UI/ProfilePage/store";

interface EditPersonalDataProps {

}

const EditPersonalDataWindow = forwardRef((props: EditPersonalDataProps, ref) => {
    const info = useUserInfoStore(state => state.info);
    const setFullName = useUserInfoStore(state => state.setFullName);
    const setDate = useUserInfoStore(state => state.setDate);
    const setContactInfo = useUserInfoStore(state => state.setContactInfo);

    const [open, setOpen] = useState(false);
    const [formData, setFormData] = useState({
        name: info.fullName.name,
        surname: info.fullName.surname,
        patronymic: info.fullName.patronymic,
        date: info.date,
        email: info.contacts.email,
        messenger: info.contacts.messenger,
    });

    const handleClickOpen = () => {
        setFormData({
            name: info.fullName.name,
            surname: info.fullName.surname,
            patronymic: info.fullName.patronymic,
            date: info.date,
            email: info.contacts.email,
            messenger: info.contacts.messenger,
        })
        setOpen(true);
    };

    // С помощью useImperativeHandle мы делаем handleClick доступным для родителя
    useImperativeHandle(ref, () => ({
        triggerHandleClick: handleClickOpen
    }));

    const handleClose = () => {
        setFormData({
            name: info.fullName.name,
            surname: info.fullName.surname,
            patronymic: info.fullName.patronymic,
            date: info.date,
            email: info.contacts.email,
            messenger: info.contacts.messenger,
        })
        setOpen(false);
    };

    const handleInputChange = (event:React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setFormData({
            ...formData,
            [name]: value,
        });
    };

    const handleSubmit = () => {
        console.log('Данные формы:', formData);
        setFullName({name: formData.name, surname: formData.surname, patronymic: formData.patronymic});
        setDate(formData.date);
        setContactInfo({email: formData.email, messenger: formData.messenger});
        setOpen(false);
    };

    return (
        <Dialog open={open} onClose={handleClose}>
            <DialogTitle>Контактная информация</DialogTitle>
            <DialogContent>
                <TextField
                    autoFocus
                    margin="dense"
                    name="name"
                    label="Имя"
                    type="text"
                    fullWidth
                    value={formData.name}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="surname"
                    label="Фамилия"
                    type="text"
                    fullWidth
                    value={formData.surname}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="patronymic"
                    label="Отчество"
                    type="text"
                    fullWidth
                    value={formData.patronymic}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="date"
                    label="Дата рождения"
                    type="text"
                    fullWidth
                    value={formData.date}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="email"
                    label="Email"
                    type="email"
                    fullWidth
                    value={formData.email}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="messenger"
                    label="Мессенджер"
                    type="text"
                    fullWidth
                    value={formData.messenger}
                    onChange={handleInputChange}
                />
            </DialogContent>
            <DialogActions>
                <Button onClick={handleClose}>Отмена</Button>
                <Button onClick={handleSubmit} variant="contained" color="primary">
                    Сохранить
                </Button>
            </DialogActions>
        </Dialog>
    );
});

export default EditPersonalDataWindow;