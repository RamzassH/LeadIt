'use server';
import { cookies } from 'next/headers';
import {createOrganizationAPI} from "@/api/organization/create";

interface CreateOrganizationRequest {
    name: string;
    description: string;
    image: string;
}

export async function createOrganizationServerFunction(data: CreateOrganizationRequest) {
    try {
        console.log('Server create organization');
        // Асинхронное получение cookies
        const cookieStore = await cookies();
        const token = cookieStore.get("auth-token");
        if (!token) {
            throw new Error('Invalid token');
        }
        console.log('Create organization token', token);
        const response = await createOrganizationAPI(data, token?.value)
        console.log(response)

    } catch (error) {
        console.error('Login failed:', error);
        throw new Error('Login failed');
    }
}