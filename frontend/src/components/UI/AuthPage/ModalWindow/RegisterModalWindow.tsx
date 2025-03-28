import React from 'react';
import ModalWindow from "@/components/UI/ModalWindowTemplate/ModalWindow";
import Loader from "@/components/UI/AuthPage/Loader/Loader";
import Button from "@mui/material/Button";

interface ModalWindowProps {
    open: boolean;
    callback: () => void;
    handleClose: () => void;
    children?: React.ReactNode;
}

function RegisterModalWindow({open, handleClose, children, callback}: ModalWindowProps) {
    const actions = () => {
        return (
            <Button onClick={callback} variant="contained" color="primary">
                Ок
            </Button>
        )
    }

    return (
        <ModalWindow title={"Регистрация"} open={open} handleClose={handleClose} actions={children? actions(): null}>
            <div style={{display: "flex", justifyContent: "center"}}>
                {children ? children : <Loader/>}
            </div>
        </ModalWindow>
    );
}

export default RegisterModalWindow;