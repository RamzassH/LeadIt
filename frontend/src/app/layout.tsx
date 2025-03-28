import type { Metadata } from "next";
import localFont from "next/font/local";
import GlobalStyle from "@/app/styled";

const geistSans = localFont({
    src: "./fonts/GeistVF.woff",
    variable: "--font-geist-sans",
    weight: "100 900",
});
const geistMono = localFont({
    src: "./fonts/GeistMonoVF.woff",
    variable: "--font-geist-mono",
    weight: "100 900",
});

export const metadata: Metadata = {
    title: "Гойда по ссылке в описании",
    description: "Гойда и этим всё сказано",
};

export default function RootLayout({
                                       children,
                                   }: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="en">
        <GlobalStyle/>
        <body
            className={`${geistSans.variable} ${geistMono.variable}`}
        >
        {children}
        </body>
        </html>
    );
}
