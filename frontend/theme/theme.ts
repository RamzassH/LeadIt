import {createTheme} from '@mui/material/styles';

const theme = createTheme({
    palette: {
        background: {
            default: '#F7F8F3',
        },
        primary: {
            main: '#78BCC4',
        },
        secondary: {
            main: '#B8DBDC',
        },
        dark: {
            main: '#012C3D',
        },
        accent: {
            main: '#F8444F',
        },
    },
    typography: {
        fontFamily: '"Roboto", "Helvetica", "Arial", sans-serif',
        fontSize: 14,
    },
    components: {
        MuiButton: {
            styleOverrides: {
                root: {
                    borderRadius: '8px',
                },
            },
        },
    },
});

export default theme;
