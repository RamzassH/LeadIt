'use server';
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';
import {loginAPI} from "@/api/auth/login";

interface LoginRequest {
    email: string;
    password: string;
}

export async function loginServerFunction(data: LoginRequest) {
    try {
        console.log('Server login');
        const response = await loginAPI(data)
        // Асинхронное получение cookies
        const cookieStore = await cookies();

        // Установка куки
        cookieStore.set('auth-token', response.data.token, {
            httpOnly: true,
            secure: process.env.NODE_ENV === 'production',
            sameSite: 'lax',
            maxAge: 60 * 15,
            path: '/',
        });

    } catch (error) {
        console.error('Login failed:', error);
        throw new Error('Login failed');
    }
}