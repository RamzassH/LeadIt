import { styled } from '@mui/material/styles';
import { Box, Typography, Checkbox, Divider, TextField } from '@mui/material';

export const RoleSettingsContainer = styled(Box)(({ theme }) => ({
    backgroundColor: theme.palette.background.paper,
    borderRadius: theme.shape.borderRadius,
    padding: theme.spacing(3),
    maxWidth: 800,
}));

export const Header = styled(Box)(({ theme }) => ({
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: theme.spacing(2),
}));

export const SectionTitle = styled(Typography)(({ theme }) => ({
    fontWeight: 500,
    marginBottom: theme.spacing(2),
}));

export const ColumnsContainer = styled(Box)({
    display: 'grid',
    gridTemplateColumns: '1fr 1fr 1fr',
    gap: 24,
});

export const RoleNameInput = styled(TextField)(({ theme }) => ({
    marginBottom: theme.spacing(3),
    '& .MuiInputLabel-asterisk': {
        color: theme.palette.error.main,
    },
}));

export const PermissionItem = styled(Box)({
    display: 'flex',
    alignItems: 'center',
    marginBottom: 8,
});
