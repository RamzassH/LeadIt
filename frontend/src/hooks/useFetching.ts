import { useState } from "react";

// Тип для callback функции, которая может быть асинхронной и принимать любые аргументы
type CallbackFunction = (...args: any[]) => Promise<void>;

export const useFetching = (callback: CallbackFunction) => {
    const [isLoading, setIsLoading] = useState<boolean>(false);
    const [error, setError] = useState<string>('');

    const fetching = async (...args: any[]): Promise<void> => {
        try {
            setIsLoading(true);
            setError("")
            await callback(...args);
        } catch (e) {
            setError(e instanceof Error ? e.message : String(e));  // Обрабатываем ошибку, чтобы получить строку
        } finally {
            setIsLoading(false);
        }
    };

    return [fetching, isLoading, error] as const; // Используем 'as const', чтобы сохранить типы возвращаемых значений
};
