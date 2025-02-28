import React, {forwardRef, useImperativeHandle, useState} from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import TextField from '@mui/material/TextField';
import useUserInfoStore from "@/components/UI/ProfilePage/store";

interface EditDescriptionDataProps {

}

const EditDescriptionDataWindow = forwardRef((props: EditDescriptionDataProps, ref) => {
    const info = useUserInfoStore(state => state.info);
    const setDescription = useUserInfoStore(state => state.setDescription);


    const [open, setOpen] = useState(false);
    const [formData, setFormData] = useState({
        description : info.description,
    });

    const handleClickOpen = () => {
        setFormData({
            description : info.description,
        })
        setOpen(true);
    };

    // С помощью useImperativeHandle мы делаем handleClick доступным для родителя
    useImperativeHandle(ref, () => ({
        triggerHandleClick: handleClickOpen
    }));

    const handleClose = () => {
        setFormData({
            description : info.description,
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
        setDescription(formData.description);
        setOpen(false);
    };

    return (
        <Dialog open={open} onClose={handleClose}>
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
                    value={formData.description}
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

export default EditDescriptionDataWindow;