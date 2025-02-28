import React, {forwardRef, useImperativeHandle, useState} from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import TextField from '@mui/material/TextField';
import useUserInfoStore from "@/components/UI/ProfilePage/store";

interface EditOrganizationDataProps {

}

const EditOrganizationDataWindow = forwardRef((props: EditOrganizationDataProps, ref) => {
    const info = useUserInfoStore(state => state.info);
    const setProjectInfo = useUserInfoStore(state => state.setProjectInfo);


    const [open, setOpen] = useState(false);
    const [formData, setFormData] = useState({
        organization: info.projectInfo.organization,
        projects: info.projectInfo.projects,
        position: info.projectInfo.position,
    });

    const handleClickOpen = () => {
        setFormData({
            organization: info.projectInfo.organization,
            projects: info.projectInfo.projects,
            position: info.projectInfo.position,
        })
        setOpen(true);
    };

    // С помощью useImperativeHandle мы делаем handleClick доступным для родителя
    useImperativeHandle(ref, () => ({
        triggerHandleClick: handleClickOpen
    }));

    const handleClose = () => {
        setFormData({
            organization: info.projectInfo.organization,
            projects: info.projectInfo.projects,
            position: info.projectInfo.position,
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
        setProjectInfo({projects: formData.projects, position: formData.position, organization: formData.organization});
        setOpen(false);
    };

    return (
        <Dialog open={open} onClose={handleClose}>
            <DialogTitle>Проектная деятельность</DialogTitle>
            <DialogContent>
                <TextField
                    autoFocus
                    margin="dense"
                    name="organization"
                    label="Организация"
                    type="text"
                    fullWidth
                    value={formData.organization}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="projects"
                    label="Проекты"
                    type="text"
                    fullWidth
                    value={formData.projects}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="position"
                    label="Должность"
                    type="text"
                    fullWidth
                    value={formData.position}
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

export default EditOrganizationDataWindow;