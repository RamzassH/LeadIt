import React from "react";
import styles from "./SuccessMessage.module.css";

interface SuccessMessageProps {
    message: string;
}

export default function SuccessMessage({ message }:SuccessMessageProps) {
    return (
        <div className={styles.successMessage}>
            <span className={styles.icon}>✅</span>
            <span>{message}</span>
        </div>
    );
};