import { styled } from '@mui/material/styles';
import { Box, Typography} from '@mui/material';
import TextField from "@mui/material/TextField";

export const VisualizationCard = styled(Box)(({ theme }) => ({
    backgroundColor: theme.palette.background.paper,
    borderRadius: theme.shape.borderRadius,
    boxShadow: theme.shadows[1],
    padding: theme.spacing(3),
    maxWidth: 660,
    margin: 'auto',
}));

export const Header = styled(Box)(({ theme }) => ({
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: theme.spacing(3),
}));

export const VisualizationItem = styled(Box)(({ theme }) => ({
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: theme.spacing(2, 0),
}));

export const VisualizationDescription = styled(Typography)(({ theme }) => ({
    color: theme.palette.text.secondary,
    fontSize: '0.875rem',
    marginTop: theme.spacing(0.5),
}));

export const RoleNameInput = styled(TextField)(({ theme }) => ({
    marginBottom: theme.spacing(3),
    '& .MuiInputLabel-asterisk': {
        color: theme.palette.error.main,
    },
}));