import React from "react";
import styles from "./Error.module.css";

interface ErrorProps {
    children: React.ReactNode;
}

export default function Error({ children }:ErrorProps)  {
    return <div className={styles.errorText}>{children}</div>;
};
