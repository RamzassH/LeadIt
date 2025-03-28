import React from 'react';
import ModalWindow from "@/components/UI/ModalWindowTeamplate/ModalWindow";
import Loader from "@/components/UI/AuthPage/Loader/Loader";

interface ModalWindowProps {
    open: boolean;
    handleClose: () => void;
}

function LoginModalWindow({open, handleClose}: ModalWindowProps) {
    return (
        <ModalWindow title={"Пожалуйста подождите..."} open={open} handleClose={handleClose}>
            <div style={{display: "flex", justifyContent: "center"}}>
                <Loader/>
            </div>
        </ModalWindow>
    );
}

export default LoginModalWindow;