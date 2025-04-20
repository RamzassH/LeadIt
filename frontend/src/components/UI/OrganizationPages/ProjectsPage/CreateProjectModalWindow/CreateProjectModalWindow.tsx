import React, {useEffect, useState} from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import ModalWindow, {ModalWindowProps} from "@/components/UI/ModalWindowTemplate/ModalWindow";
import useProjectStore, {Project} from "@/store/ProjectsPageStore/store";


const CreateProjectModalWindow = ({open, handleClose}: ModalWindowProps) => {
    const [data, setData] = useState<Project>({id: "", name: "", description: ""});
    const {addProject} = useProjectStore()

    const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setData({
            ...data,
            [name]: value
        })
    }
    const handleSubmit = () => {
        addProject({...data, id: Date.now().toString()})
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
        <ModalWindow open={open} handleClose={() => {}} actions={actions()} title="Описание">
            <TextField
                autoFocus
                margin="dense"
                name="name"
                label="Название проекта"
                type="text"
                fullWidth
                value={data.name}
                onChange={handleInputChange}
            />
            <TextField
                margin="dense"
                name="description"
                label="Описание"
                type="text"
                fullWidth
                multiline
                rows={8}
                value={data.description}
                onChange={handleInputChange}
            />
        </ModalWindow>
    );
};

export default CreateProjectModalWindow;