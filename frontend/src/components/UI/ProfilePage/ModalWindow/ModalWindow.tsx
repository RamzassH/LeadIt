import React, {ChangeEvent, useState} from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import TextField, {TextFieldProps} from '@mui/material/TextField';

interface ModalWindowProps {
    title?: string;
    children?: React.ReactNode;
}

function ModalWindow({title, children}: ModalWindowProps) {
    const [open, setOpen] = useState(false);
    const [formData, setFormData] = useState({
        name: '',
        email: '',
        message: '',
    });

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const handleInputChange = (event:React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setFormData({
            ...formData,
            [name]: value,
        });
    };

    // Клонируем children и добавляем свойство onChange
    const childrenWithProps = React.Children.map(children, (child) => {
        // Указываем тип TextFieldProps для child
        if (React.isValidElement<TextFieldProps>(child)) {
            return React.cloneElement<TextFieldProps>(child, {
                onChange: (event: ChangeEvent<HTMLInputElement>) => {
                    if (child.props.onChange) {
                        child.props.onChange(event); // Вызов существующего обработчика
                    }
                    handleInputChange(event); // Вызов нового обработчика
                },
            });
        }
        return child;
    });

    const handleSubmit = () => {
        console.log('Данные формы:', formData);
        handleClose(); // Закрываем модальное окно после отправки
    };

    return (
        <Dialog open={open} onClose={handleClose}>
            <DialogTitle>{title}</DialogTitle>
            <DialogContent>
                {childrenWithProps}
                {/*<TextField
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
                    name="email"
                    label="Email"
                    type="email"
                    fullWidth
                    value={formData.email}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="message"
                    label="Сообщение"
                    type="text"
                    fullWidth
                    multiline
                    rows={4}
                    value={formData.message}
                    onChange={handleInputChange}
                />
                */}
            </DialogContent>
            <DialogActions>
                <Button onClick={handleClose}>Отмена</Button>
                <Button onClick={handleSubmit} variant="contained" color="primary">
                    Сохранить
                </Button>
            </DialogActions>
        </Dialog>
    );
}

export default ModalWindow;