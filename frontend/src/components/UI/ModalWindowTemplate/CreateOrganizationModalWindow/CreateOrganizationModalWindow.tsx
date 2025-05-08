import ModalWindow, {ModalWindowProps} from "@/components/UI/ModalWindowTemplate/ModalWindow";
import React, {useEffect, useState} from "react";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";

interface CreateOrganizationModalWindowProps extends ModalWindowProps {
    callback: (data: CreateOrganizationForm) => void;
}

export interface CreateOrganizationForm {
    name: string;
    image: string;
    description: string;
}

const CreateOrganizationModalWindow = ({open, handleClose, callback}: CreateOrganizationModalWindowProps) => {
    const [form, setForm] = useState<CreateOrganizationForm>({
        name: '',
        image: '',
        description: '',
    });

    const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setForm({
            ...form,
            [name]: value
        });
    }
    const handleSubmit = () => {
        callback(form)
        handleClose();
    }

    const close = () => {
        handleClose();
    }

    const actions = () => {
        return (
            <div>
                <Button onClick={close}>Отмена</Button>
                <Button onClick={() => {handleSubmit()}} variant="contained" color="primary">
                    Сохранить
                </Button>
            </div>
        )
    }
    return (
        <ModalWindow open={open} handleClose={() => {}} actions={actions()} title="Создание организации">
            <TextField
                autoFocus
                margin="dense"
                name="name"
                label="Название организации"
                type="text"
                fullWidth
                value={form.name}
                onChange={handleInputChange}
            />
            <TextField
                margin="dense"
                name="image"
                label="Логотип"
                type="text"
                fullWidth
                value={form.image}
                onChange={handleInputChange}
            />
            <TextField
                margin="dense"
                name="description"
                label="Описание организации"
                type="text"
                fullWidth
                multiline
                rows={8}
                value={form.description}
                onChange={handleInputChange}
            />
        </ModalWindow>
    );
};

export default CreateOrganizationModalWindow;