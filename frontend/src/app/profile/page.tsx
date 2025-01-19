"use client"
import "./local.css"
import MenuButton from "@/app/profile/components/SideMenu/MenuButton/styles/MenuButton";
import theme from "../../../theme/theme";
import {ThemeProvider} from "@mui/system";

export default function Profile() {
    return (
        <ThemeProvider theme={theme}>
            <body>
            <header className="header">
                <div className="menu-button">
                    <div className="menu-button-component"/>
                    <div className="menu-button-component"/>
                    <div className="menu-button-component"/>
                </div>
                <div className="logo">
                    <svg className="logo-image" width="27.257812" height="34.171875" viewBox="0 0 27.2578 34.1719"
                         fill="none"
                         xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink">
                        <desc>
                            Created with Pixso.
                        </desc>
                        <defs/>
                        <path id="P"
                              d="M0 0L0 34.17L7.44 34.17L7.44 21.87L15.11 21.87C18.71 21.87 21.64 20.97 23.9 19.14C26.1 17.32 27.25 14.58 27.25 10.99C27.25 7.33 26.1 4.6 23.84 2.77C21.59 0.96 18.57 0 14.73 0L0 0ZM7.44 6.14L14.15 6.14C18 6.14 19.96 7.81 19.96 11.08C19.96 14.34 18 15.98 14.1 15.98L7.44 15.98L7.44 6.14Z"
                              fill="#012C3D" fillOpacity="1.000000" fillRule="nonzero"/>
                    </svg>
                </div>
                <div className="header-space"/>
                <div className="login-button">
                    <svg className="login-button-component-1" width="35.083008" height="1.500000"
                         viewBox="0 0 35.083 1.5" fill="none"
                         xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink">
                        <desc>
                            Created with Pixso.
                        </desc>
                        <defs/>
                        <path id="path"
                              d="M0.76 1.47L0.75 1.5C0.33 1.5 0 1.17 0 0.75C0 0.32 0.33 0 0.75 0L0.76 0.02L0.76 1.47ZM34.31 0.02L34.33 0C34.75 0 35.08 0.32 35.08 0.75C35.08 1.17 34.75 1.5 34.33 1.5L34.31 1.47L34.31 0.02Z"
                              fill="#000000" fillOpacity="0" fillRule="nonzero"/>
                        <path id="path" d="M0.75 0.75L34.33 0.75" stroke="#78BCC4" strokeOpacity="1.000000"
                              strokeWidth="1.500000" strokeLinejoin="round" strokeLinecap="round"/>
                    </svg>
                    <svg className="login-button-component-2" width="13.171875" height="19.880859"
                         viewBox="0 0 13.1719 19.8809" fill="none"
                         xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink">
                        <desc>
                            Created with Pixso.
                        </desc>
                        <defs/>
                        <path id="path"
                              d="M12.85 18.54L12.88 18.53C13.21 18.79 13.26 19.26 13 19.59C12.74 19.92 12.28 19.97 11.95 19.71L11.95 19.68L12.85 18.54ZM11.95 0.19L11.95 0.16C12.28 -0.1 12.74 -0.05 13 0.28C13.26 0.61 13.21 1.08 12.88 1.34L12.85 1.33L11.95 0.19Z"
                              fill="#000000" fillOpacity="0" fillRule="nonzero"/>
                        <path id="path" d="M12.41 19.12L0.79 9.94L12.41 0.75" stroke="#78BCC4" strokeOpacity="1.000000"
                              strokeWidth="1.500000" strokeLinejoin="round" strokeLinecap="round"/>
                    </svg>
                    <svg className="login-button-component-3" width="24.750000" height="38.250000"
                         viewBox="0 0 24.75 38.25" fill="none"
                         xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink">
                        <desc>
                            Created with Pixso.
                        </desc>
                        <defs/>
                        <path id="path"
                              d="M23.96 0.02L23.98 0C24.4 0 24.73 0.32 24.73 0.75C24.73 1.17 24.4 1.5 23.98 1.5L23.96 1.47L23.96 0.02ZM23.98 36.77L24 36.75C24.41 36.75 24.75 37.08 24.75 37.5C24.75 37.91 24.41 38.25 24 38.25L23.98 38.23L23.98 36.77Z"
                              fill="#000000" fillOpacity="0" fillRule="nonzero"/>
                        <path id="path" d="M23.98 0.75L0.75 0.75L0.75 37.5L24 37.5" stroke="#78BCC4"
                              strokeOpacity="1.000000" strokeWidth="1.500000" strokeLinejoin="round"
                              strokeLinecap="round"/>
                    </svg>

                </div>
            </header>
            <main>
                <div className="menu open">
                    <MenuButton variant="contained" onClick={() => {console.log('suka')}}>
                        Стилизованная кнопка
                    </MenuButton>
                    <div className="menu-drop-list">
                        <div className="menu-item">
                            <div className="menu-item-logo">
                                <svg className="menu-item-logo-image-1" width="29.166504" height="26.250000"
                                     viewBox="0 0 29.1665 26.25" fill="none"
                                     xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink">
                                    <desc>
                                        Created with Pixso.
                                    </desc>
                                    <defs/>
                                    <path id="矢量 117"
                                          d="M0 23.33L29.16 23.33L29.16 26.25L0 26.25L0 23.33ZM13.12 7.29L16.04 7.29L16.04 18.95L13.12 18.95L13.12 7.29ZM8.69 7.29L11.79 7.29L7.44 18.9L4.52 18.9L0.17 7.29L3.27 7.29L5.98 14.74L8.69 7.29ZM21.87 16.04L21.87 18.95L18.95 18.95L18.95 7.29L24.79 7.29C25.95 7.29 27.06 7.75 27.88 8.57C28.7 9.39 29.16 10.5 29.16 11.66C29.16 12.82 28.7 13.93 27.88 14.76C27.06 15.58 25.95 16.04 24.79 16.04L21.87 16.04ZM21.87 10.2L21.87 13.12L24.79 13.12C25.17 13.12 25.54 12.97 25.82 12.69C26.09 12.42 26.25 12.05 26.25 11.66C26.25 11.27 26.09 10.9 25.82 10.63C25.54 10.36 25.17 10.2 24.79 10.2L21.87 10.2ZM0 0L29.16 0L29.16 2.91L0 2.91L0 0Z"
                                          fill="#78BCC4" fillOpacity="1.000000" fillRule="evenodd"/>
                                </svg>
                            </div>
                            <div className="menu-item-text">
                                Главная
                            </div>
                            <div className="menu-item-drop-list-icon">
                                <svg className="menu-item-drop-list-icon-image" width="7.548340" height="13.506104"
                                     viewBox="0 0 7.54834 13.5061" fill="none"
                                     xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink">
                                    <desc>
                                        Created with Pixso.
                                    </desc>
                                    <defs/>
                                    <path id="path"
                                          d="M7.29 12.22L7.32 12.22C7.62 12.51 7.62 12.98 7.32 13.28C7.02 13.58 6.56 13.58 6.26 13.28L6.26 13.25L7.29 12.22ZM6.26 0.25L6.26 0.22C6.56 -0.08 7.02 -0.08 7.32 0.22C7.62 0.51 7.62 0.98 7.32 1.28L7.29 1.28L6.26 0.25Z"
                                          fill="#000000" fillOpacity="0" fillRule="nonzero"/>
                                    <path id="path" d="M6.79 12.75L0.79 6.75L6.79 0.75" stroke="#78BCC4"
                                          strokeOpacity="1.000000" strokeWidth="1.500000" strokeLinejoin="round"
                                          strokeLinecap="round"/>
                                </svg>

                            </div>
                        </div>

                    </div>
                    <div/>
                </div>
                <div className="goida">
                    <img className="goida-image" src="/images/goida.gif" alt="Описание гифки"/>
                </div>
            </main>
            </body>
            </ThemeProvider>
    );
}
