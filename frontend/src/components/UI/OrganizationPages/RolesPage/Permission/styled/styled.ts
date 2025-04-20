import { styled } from '@mui/material/styles';
import { Box, Typography, Button, Switch, Divider } from '@mui/material';

export const PermissionCard = styled(Box)(({ theme }) => ({
    backgroundColor: theme.palette.background.paper,
    borderRadius: theme.shape.borderRadius,
    boxShadow: theme.shadows[1],
    padding: theme.spacing(3),
    maxWidth: 600,
    margin: 'auto',
}));

export const Header = styled(Box)(({ theme }) => ({
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: theme.spacing(3),
}));

export const PermissionItem = styled(Box)(({ theme }) => ({
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: theme.spacing(2, 0),
}));

export const PermissionDescription = styled(Typography)(({ theme }) => ({
    color: theme.palette.text.secondary,
    fontSize: '0.875rem',
    marginTop: theme.spacing(0.5),
}));
