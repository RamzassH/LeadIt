/**
 * Сохраняет данные в localStorage.
 * @param key - Ключ, по которому будут храниться данные.
 * @param value - Данные для сохранения (объект, строка, число и т.д.).
 */
export const setLocalStorage = <T>(key: string, value: T): void => {
    if (typeof window === 'undefined') {
        console.warn('localStorage доступен только в браузере.');
        return;
    }
    try {
        const serializedValue = JSON.stringify(value);
        localStorage.setItem(key, serializedValue);
    } catch (error) {
        console.error('Ошибка записи в localStorage:', error);
    }
};

/**
 * Получает данные из localStorage.
 * @param key - Ключ, по которому хранятся данные.
 * @returns Данные или null, если их нет или произошла ошибка.
 */
export const getLocalStorage = <T>(key: string): T | null => {
    if (typeof window === 'undefined') {
        console.warn('localStorage доступен только в браузере.');
        return null;
    }
    try {
        const serializedValue = localStorage.getItem(key);
        return serializedValue ? JSON.parse(serializedValue) : null;
    } catch (error) {
        console.error('Ошибка чтения из localStorage:', error);
        return null;
    }
};

/**
 * Удаляет данные из localStorage по ключу.
 * @param key - Ключ, который нужно удалить.
 */
export const removeLocalStorage = (key: string): void => {
    if (typeof window === 'undefined') {
        console.warn('localStorage доступен только в браузере.');
        return;
    }
    try {
        localStorage.removeItem(key);
    } catch (error) {
        console.error('Ошибка удаления из localStorage:', error);
    }
};

/**
 * Проверяет, доступен ли localStorage в текущей среде.
 * @returns true, если localStorage доступен.
 */
export const isLocalStorageSupported = (): boolean => {
    try {
        const testKey = '__test__';
        localStorage.setItem(testKey, testKey);
        localStorage.removeItem(testKey);
        return true;
    } catch (error) {
        return false;
    }
};