"use client";
import LoginForm from "@/components/UI/AuthPage/LoginForm/LoginForm";
import CreateAccountForm from "@/components/UI/AuthPage/CreateAccountForm/CreateAccountForm";
import { useState } from "react";
import theme from "../../../theme/theme";
import { ThemeProvider } from "@mui/system";
import { Main } from "@/app/auth/styled/AuthStyled";
import { motion, AnimatePresence } from "framer-motion";

export default function Auth() {
    const [isLoginForm, setLogin] = useState(true);

    const showCreateAccountForm = () => {
        setLogin(false);
    };

    const showLoginForm = () => {
        setLogin(true);
    };

    return (
        <ThemeProvider theme={theme}>
            <Main>
                <AnimatePresence>
                    {isLoginForm ? (
                        <motion.div
                            key="login"
                            initial={{ translateX : "-100%"}}
                            animate={{ translateX : 0}}
                            exit={{ translateX : "-100%"}}
                            transition={{ duration: 0.5 }}
                            style={{ width: "100%", height: "100%", position: "absolute" }}
                        >
                            <LoginForm callback={showCreateAccountForm} />
                        </motion.div>
                    ) : (
                        <motion.div
                            key="createAccount"
                            initial={{ translateX : "100%"}}
                            animate={{ translateX : 0}}
                            exit={{ translateX : "100%"}}
                            transition={{ duration: 0.5 }}
                            style={{ width: "100%", height: "100%", position: "absolute" }}
                        >
                            <CreateAccountForm returnCallback={showLoginForm} />
                        </motion.div>
                    )}
                </AnimatePresence>
            </Main>
        </ThemeProvider>
    );
}