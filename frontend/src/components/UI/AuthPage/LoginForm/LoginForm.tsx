"use client";
import { useForm, Controller } from "react-hook-form";
import Input from "@/components/UI/AuthPage/Input/Input";
import Button from "@/components/UI/AuthPage/Button/Button";
import Checkbox from "@/components/UI/AuthPage/Checkbox/Checkbox";
import Loader from "@/components/UI/AuthPage/Loader/Loader";
import {
    BackgroundContainer, ButtonContainer,
    Container, Content,
    Form, NonBackgroundContainer,
    Title,
    TitleContainer
} from "@/components/UI/AuthPage/LoginForm/styled/LoginForm";
import {useFetching} from "@/hooks/useFetching";
import {data} from "framer-motion/m";
import {loginAPI} from "@/api/auth/login";
import useGlobalStore from "@/app/store";
import {useEffect} from "react";
import LoginModalWindow from "@/components/UI/AuthPage/ModalWindow/LoginModalWindow";
import {useRouter} from "next/navigation";


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
        setValue,
        clearErrors,
    } = useForm<LoginData>();
    const router = useRouter();
    const globalStore = useGlobalStore();
    const [login, isLoading, error] = useFetching(async (data) => {
        await new Promise(resolve => setTimeout(resolve, 500));
        const response = await loginAPI(data);
        globalStore.setRefreshToken(response.data.refreshToken);
        globalStore.setLogin(true)
    });

    // Обработчик отправки формы
    const onSubmit = async (data: LoginData) => {
        login({email: data.login, password: data.password});
    };

    useEffect(() => {
        if (globalStore.isLogin) {
            router.push("/profile");
        }
    }, [globalStore.isLogin]);

    useEffect(() => {
        if (error) {
            setValue("password", "")
            setError("login", {
                type: "manual",
                message: "",
            });
            setError("password", {
                type: "manual",
                message: "Ошибка авторизации. Попробуйте снова.",
            });
        }
    }, [error]);

    const handleCreateAccount = () => {
        callback();
    };

    return (
        <Container>
            <NonBackgroundContainer>
                <Content>
                    <TitleContainer>
                        <Title>
                            Нет аккаунта?
                        </Title>
                    </TitleContainer>
                    <div style={{width:"100%", height:"calc(45rem/16)"}}/>
                    <ButtonContainer>
                        <Button callback={handleCreateAccount} isReverseBackground={true}>
                            Зарегистрироваться
                        </Button>
                    </ButtonContainer>
                </Content>
            </NonBackgroundContainer>
            <BackgroundContainer>
                <Content>
                    <TitleContainer>
                        <Title>
                            Добро пожаловать!
                        </Title>
                    </TitleContainer>
                    <Form onSubmit={handleSubmit(onSubmit)}>
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
                                    isError={!!errors.password}
                                    errorMessage={errors.password?.message}
                                />
                            )}
                        />

                        <Checkbox label="Запомнить меня" style={{margin: "0 0 calc(16rem/16)  calc(20rem/16)"}} />
                    </Form>
                    <ButtonContainer>
                        <Button type="submit" callback={handleSubmit(onSubmit)}>
                            Войти
                        </Button>
                    </ButtonContainer>
                    <LoginModalWindow open={isLoading} handleClose={() => {}}/>
                </Content>
            </BackgroundContainer>
        </Container>
    );
}
