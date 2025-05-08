// middleware.ts
import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

const protectedPaths = [
    '/organization',
    '/profile',
    // Добавьте другие защищённые пути
];

export function middleware(request: NextRequest) {
    const token = request.cookies.get('auth-token')?.value;

    if (!token && protectedPaths.some(path =>
        request.nextUrl.pathname.startsWith(path)
    )) {
        const loginUrl = new URL('/auth', request.url);
        loginUrl.searchParams.set('from', request.nextUrl.pathname);
        //return NextResponse.redirect(loginUrl);
    } else if (token && request.nextUrl.pathname.startsWith('/auth')) {
        const loginUrl = new URL('/', request.url);
        //return NextResponse.redirect(loginUrl);
    }

    return NextResponse.next();
}

export const config = {
    matcher: [
        '/((?!api|_next/static|favicon.ico).*)', // Исключает API и статику
    ],
};