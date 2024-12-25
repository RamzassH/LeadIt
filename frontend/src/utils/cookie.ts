// Устанавливаем JWT в cookie с флагами Secure и SameSite
export const setValueInCookie = (name: string, value: string, days: number = 7) => {
    const date = new Date();
    date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000)); // Устанавливаем срок действия cookie
    const expires = `expires=${date.toUTCString()}`;
    document.cookie = `${name}=${value}; ${expires}; path=/; Secure; SameSite=Strict`;
};

// Получаем JWT из cookie
export const getValueFromCookie = (name: string): string | null => {
    const cookie = document.cookie.split('; ').find(row => row.startsWith(`${name}=`));
    return cookie ? cookie.split('=')[1] : null;
};

// Удаляем JWT из cookie
export const removeValueFromCookie = (name: string) => {
    document.cookie = `${name}=; path=/; max-age=0; Secure; SameSite=Strict`;
};
