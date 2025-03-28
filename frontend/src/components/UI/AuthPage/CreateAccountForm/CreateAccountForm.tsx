"use client";
import Input from "@/components/UI/AuthPage/Input/Input";
import { useForm, Controller } from "react-hook-form";
import Button from "@/components/UI/AuthPage/Button/Button";
import Checkbox from "@/components/UI/AuthPage/Checkbox/Checkbox";
import Loader from "@/components/UI/AuthPage/Loader/Loader";
import SuccessMessage from "@/components/UI/AuthPage/SuccessMessage/SuccessMessage";
import { useState } from "react";
import {
    BackgroundContainer, ButtonContainer,
    Container, Content,
    Form,
    Title, TitleContainer
} from "@/components/UI/AuthPage/CreateAccountForm/styled/CreateAccountForm";

interface AccountData {
    surname: string;
    name: string;
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
        <Container>
            <BackgroundContainer>
                <Content>
                    <TitleContainer>
                        <Title>Создайте аккаунт</Title>
                    </TitleContainer>

                    <Form onSubmit={handleSubmit(onSubmit)}>
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

                        <Checkbox label="Согласен на обработку персональных данных" style={{margin: "0 0 calc(18rem/16)  calc(20rem/16)"}}/>
                    </Form>

                    <ButtonContainer>
                        <Button type="submit" callback={handleSubmit(onSubmit)}>Создать аккаунт</Button>
                    </ButtonContainer>
                </Content>
            </BackgroundContainer>
            <Content>
                <TitleContainer>
                    <Title>
                        Уже зарегистрированы?
                    </Title>
                </TitleContainer>
                <div style={{width:"100%", height:"calc(45rem/16)"}}/>
                <ButtonContainer>
                    <Button callback={returnCallback}>
                        Войти
                    </Button>
                </ButtonContainer>
            </Content>
        </Container>
    );
}
