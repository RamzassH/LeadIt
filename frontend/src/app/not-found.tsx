// app/not-found.tsx
import type { Metadata } from 'next';

export const metadata: Metadata = {
    title: '404 - Страница не найдена',
};

// app/not-found.tsx
export default function NotFound() {
    return (
        <div className="text-center py-20">
            <h1 className="text-4xl font-bold">404</h1>
            <p className="mt-4">Страница не найдена</p>
        </div>
    );
}