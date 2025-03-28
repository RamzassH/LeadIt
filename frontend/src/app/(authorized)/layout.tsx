import {ReactNode} from 'react';
// Типы для пропсов
interface DashboardLayoutProps {
    children: ReactNode;
}

export default function AuthorizedLayout({ children }: DashboardLayoutProps) {
    return (
        <body>
            {children}
        </body>
    );
}