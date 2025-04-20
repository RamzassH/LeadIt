import React, {forwardRef, useEffect, useImperativeHandle, useState} from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import ModalWindow, {ModalWindowProps} from "@/components/UI/ModalWindowTemplate/ModalWindow";
import useOrganizationInfoStore from "@/components/UI/OrganizationPages/OrganizationPage/store";


const EditDescriptionDataWindow = ({open, handleClose}: ModalWindowProps) => {
    const organization = useOrganizationInfoStore();
    const [description, setDescription] = useState<string>('');

    useEffect(() => {
        setDescription(organization.info.description);
    }, [organization.info.description]);

    const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setDescription(value);
    }
    const handleSubmit = () => {
        organization.setDescription(description);
        handleClose();
    }

    const close = () => {
        handleClose();
        setDescription(organization.info.description);
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
        <ModalWindow open={open} handleClose={() => {}} actions={actions()} title="Описание">
            <TextField
                margin="dense"
                name="description"
                label=""
                type="text"
                fullWidth
                multiline
                rows={8}
                value={description}
                onChange={handleInputChange}
            />
        </ModalWindow>
    );
};

export default EditDescriptionDataWindow;