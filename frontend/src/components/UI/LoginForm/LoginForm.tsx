"use client";
import { useForm, Controller } from "react-hook-form";
import Input from "@/components/UI/Input/Input";
import styles from "./LoginForm.module.css";
import Button from "@/components/UI/Button/Button";
import Checkbox from "@/components/UI/Checkbox/Checkbox";
import Loader from "@/components/UI/Loader/Loader";
import {useFetching} from "@/hooks/useFetching";
import {loginUser} from  "@/api/AuthService/AuthService"

interface LoginFormProps {
    callback: () => void;
}

interface LoginData {
    login: string;
    password: string;
}

export default function LoginForm({ callback }: LoginFormProps) {
    const {
        control,
        handleSubmit,
        formState: { errors },
        setError,
        clearErrors,
    } = useForm<LoginData>();

    // Обработчик отправки формы
    const onSubmit = async (data: LoginData) => {
        try {
            console.log(data.login, data.password)
            loginUser(data.login, data.password)
        } catch (error) {
            setError("login", {
                type: "manual",
                message: "",
            });
            setError("password", {
                type: "manual",
                message: "Ошибка авторизации. Попробуйте снова.",
            });
        }
    };

    const handleCreateAccount = () => {
        callback();
    };

    //if (isLoading) {
    //    return <Loader />;
    //}

    return (
        <div className={styles.container}>
            <div className={styles.title}>LeadIt</div>

            <form onSubmit={handleSubmit(onSubmit)} className={styles.form}>
                <Controller
                    name="login"
                    control={control}
                    defaultValue=""
                    rules={{
                        required: "Логин обязателен", // Встроенная валидация
                        pattern: {
                            value: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
                            message: "Введите корректный email",
                        },
                    }}
                    render={({ field }) => (
                        <Input
                            {...field}
                            type="text"
                            placeholder="Логин (email)"
                            classStyles={[]}
                            isError={!!errors.login}
                            errorMessage={errors.login?.message}
                        />
                    )}
                />

                <Controller
                    name="password"
                    control={control}
                    defaultValue=""
                    rules={{
                        required: "Пароль обязателен", // Встроенная валидация
                    }}
                    render={({ field }) => (
                        <Input
                            {...field}
                            type="password"
                            placeholder="Пароль"
                            classStyles={[]}
                            isError={!!errors.password}
                            errorMessage={errors.password?.message}
                        />
                    )}
                />

                <Checkbox label="Запомнить меня" classStyles={[styles.checkbox]} />

                <Button type="submit" classStyles={[]} callback={handleSubmit(onSubmit)}>
                    Войти
                </Button>

                <Button
                    type="button"
                    classStyles={[]}
                    style={{ backgroundColor: "#989898" }}
                    callback={handleCreateAccount}
                >
                    Создать аккаунт
                </Button>
            </form>
        </div>
    );
}
