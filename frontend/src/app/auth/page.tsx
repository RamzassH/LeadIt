"use client"
import LoginForm from "@/components/UI/AuthPage/LoginForm/LoginForm";
import CreateAccountForm from "@/components/UI/AuthPage/CreateAccountForm/CreateAccountForm";
import {useState} from "react";

export default function Auth() {
    const [isLoginForm, setLogin] = useState(true)

    const showCreateAccountForm = () => {
        setLogin(false)
    }

    const showLoginForm = () => {
        setLogin(true)
    }

    return (
        <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
            <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start">
                {isLoginForm ?
                    <LoginForm callback={showCreateAccountForm}/>:
                    <CreateAccountForm returnCallback={showLoginForm}/>
                }
            </main>
        </div>
);
}
