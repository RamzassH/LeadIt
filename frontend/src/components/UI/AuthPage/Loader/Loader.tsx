import React from "react";
import styles from "./Loader.module.css"; // Импорт стилей

export default function Loader() {
    return (
        <div className={styles.loader}>
            <div className={styles.spinner}></div>
        </div>
    );
};

