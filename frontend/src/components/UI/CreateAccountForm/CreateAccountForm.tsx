"use client";
import styles from "./CreateAccountForm.module.css";
import Input from "@/components/UI/Input/Input";
import { useForm, Controller } from "react-hook-form";
import Button from "@/components/UI/Button/Button";
import Checkbox from "@/components/UI/Checkbox/Checkbox";
import Loader from "@/components/UI/Loader/Loader";
import SuccessMessage from "@/components/UI/SuccessMessage/SuccessMessage";
import { useState } from "react";

interface AccountData {
    surname: string;
    name: string;
    patronymic: string;
    email: string;
    password: string;
    repeatPassword: string;
}

interface CreateAccountFormProps {
    returnCallback: () => void;
}

export default function CreateAccountForm({ returnCallback }: CreateAccountFormProps) {
    const {
        control,
        handleSubmit,
        formState: { errors },
        getValues,
    } = useForm<AccountData>({
        defaultValues: {
            surname: "",
            name: "",
            patronymic: "",
            email: "",
            password: "",
            repeatPassword: "",
        },
    });


    const [isLoading, setLoading] = useState(false);
    const [isSuccess, setSuccess] = useState(false);

    // Обработчик отправки формы
    const onSubmit = async (data: AccountData) => {
        console.log('Account Data:', data);

        // Здесь будет логика для отправки данных на сервер
        setLoading(true);

        // Симуляция отправки данных
        setTimeout(() => {
            setLoading(false);
            setSuccess(true);
            setTimeout(() => {
                setSuccess(false);
                returnCallback(); // Возвращаемся к предыдущему шагу после успешной регистрации
            }, 5000);
        }, 5000); // Задержка для имитации асинхронной операции
    };

    if (isLoading) {
        return <Loader />;
    }

    if (isSuccess) {
        return <SuccessMessage message="Аккаунт успешно создан" />;
    }

    return (
        <div className={styles.container}>
            <div className={styles.title}>Регистрация</div>

            <form onSubmit={handleSubmit(onSubmit)} className={styles.form}>
                <Controller
                    name="surname"
                    control={control}
                    rules={{
                        required: "Обязательное поле",
                    }}
                    render={({ field }) => (
                        <Input
                            {...field}
                            type="text"
                            placeholder="Фамилия"
                            errorMessage={errors.surname?.message}
                            isError={!!errors.surname}
                        />
                    )}
                />

                <Controller
                    name="name"
                    control={control}
                    rules={{
                        required: "Обязательное поле",
                    }}
                    render={({ field }) => (
                        <Input
                            {...field}
                            type="text"
                            placeholder="Имя"
                            errorMessage={errors.name?.message}
                            isError={!!errors.name}
                        />
                    )}
                />

                <Controller
                    name="patronymic"
                    control={control}
                    render={({ field }) => <Input {...field} type="text" placeholder="Отчество (Необязательно)" />}
                />

                <Controller
                    name="email"
                    control={control}
                    rules={{
                        required: "Обязательное поле",
                        pattern: {
                            value: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
                            message: "Введите корректный email",
                        },
                    }}
                    render={({ field }) => (
                        <Input
                            {...field}
                            type="text"
                            placeholder="Email"
                            errorMessage={errors.email?.message}
                            isError={!!errors.email}
                        />
                    )}
                />

                <Controller
                    name="password"
                    control={control}
                    rules={{
                        required: "Обязательное поле",
                        minLength: {
                            value: 6,
                            message: "Пароль должен содержать минимум 6 символов",
                        },
                    }}
                    render={({ field }) => (
                        <Input
                            {...field}
                            placeholder="Пароль"
                            type="password"
                            errorMessage={errors.password?.message}
                            isError={!!errors.password}
                        />
                    )}
                />

                <Controller
                    name="repeatPassword"
                    control={control}
                    rules={{
                        required: "Обязательное поле",
                        validate: (value) =>
                            value === getValues("password") || "Пароли не совпадают",
                    }}
                    render={({ field }) => (
                        <Input
                            {...field}
                            placeholder="Повторить пароль"
                            type="password"
                            errorMessage={errors.repeatPassword?.message}
                            isError={!!errors.repeatPassword}
                        />
                    )}
                />

                <Checkbox label="Согласен на обработку персональных данных" classStyles={[styles.checkbox]} />

                <Button type="submit" callback={handleSubmit(onSubmit)}>Создать аккаунт</Button>
            </form>
        </div>
    );
}
