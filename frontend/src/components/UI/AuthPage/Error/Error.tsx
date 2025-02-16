import React from "react";
import {ErrorText} from "@/components/UI/AuthPage/Error/styled/ErrorText";

interface ErrorProps {
    children: React.ReactNode;
}

export default function Error({ children }:ErrorProps)  {
    return <ErrorText>{children}</ErrorText>;
};
