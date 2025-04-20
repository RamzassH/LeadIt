import React, {forwardRef, useEffect, useImperativeHandle, useState} from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import useOrganizationInfoStore, {
    ContactInfo,
    Organization
} from "@/components/UI/OrganizationPages/OrganizationPage/store";
import ModalWindow, {ModalWindowProps} from "@/components/UI/ModalWindowTemplate/ModalWindow";

interface Form extends Organization, ContactInfo {

}

const EditOrganizationDataWindow = ({open, handleClose}:ModalWindowProps) => {
    const organization = useOrganizationInfoStore();
    const [form, setForm] = useState<Form>({
        name: "",
        director: "",
        email: "",
        messenger: "",
        phone: ""

    });

    useEffect(() => {
        setForm({...organization.info.organization, ...organization.info.contacts});
    }, [organization.info]);

    const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setForm(prevForm => ({
            ...prevForm,
            [name]: value
        }));
        console.log(name)
    };

    const handleSubmit = () => {
        organization.setOrganization({name: form.name, director: form.director});
        organization.setContactInfo({email: form.email, messenger: form.messenger, phone: form.phone});
        handleClose();
    }
    const close = () => {
        handleClose();
        setForm({...organization.info.organization, ...organization.info.contacts});
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
        <ModalWindow open={open} handleClose={() => {}} title="Основная информация" actions={actions()}>
            <TextField
                autoFocus
                margin="dense"
                name="name"
                label="Наименование"
                type="text"
                fullWidth
                value={form.name}
                onChange={handleInputChange}
            />
            <TextField
                margin="dense"
                name="director"
                label="Директор"
                type="text"
                fullWidth
                value={form.director}
                onChange={handleInputChange}
            />
            <TextField
                margin="dense"
                name="email"
                label="Email"
                type="text"
                fullWidth
                value={form.email}
                onChange={handleInputChange}
            />
            <TextField
                margin="dense"
                name="messenger"
                label="Контакты"
                type="text"
                fullWidth
                value={form.messenger}
                onChange={handleInputChange}
            />
            <TextField
                margin="dense"
                name="phone"
                label="Телефон"
                type="text"
                fullWidth
                value={form.phone}
                onChange={handleInputChange}
            />
        </ModalWindow>
    );
};

export default EditOrganizationDataWindow;