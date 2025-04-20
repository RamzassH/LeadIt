import { styled } from '@mui/material/styles';
import { Box, Typography, Divider, IconButton } from '@mui/material';

export const RolesContainer = styled(Box)(({ theme }) => ({
    backgroundColor: theme.palette.background.paper,
    borderRadius: theme.shape.borderRadius,
    padding: theme.spacing(3),
    boxShadow: theme.shadows[1],
    width: '100%',
    maxWidth: "calc(660rem/16)",
    marginLeft: "auto",
    marginRight: "auto",
    marginTop: "calc(36rem/16)"
}));

export const Header = styled(Box)(({ theme }) => ({
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: theme.spacing(2),
}));

export const RoleItem = styled(Box)(({ theme }) => ({
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: theme.spacing(1.5, 0),
    position: 'relative',
    border: 'calc(1rem/16) solid',
    borderRadius: 'calc(4rem/16)',
    borderColor: "rgb(0 0 0 / 12%)"
}));

export const RoleInfo = styled(Box)({
    display: 'flex',
    alignItems: 'center',
    gap: 8,
});

export const RoleName = styled(Typography)({
    marginLeft: 'calc(16rem/16)',
    fontWeight: 500,
});

export const MemberCount = styled(Typography)(({ theme }) => ({
    color: theme.palette.text.secondary,
    position: 'absolute',
    top: 'calc(50% - 11rem/16)',
    right: 'calc(50% - 20rem/16)',
}));

interface ColorIndicatorProps {
    color?: string;
}

export const ColorIndicator = styled(Box, {
    shouldForwardProp: (prop) => prop !== 'color',
})<ColorIndicatorProps>(({ theme, color }) => ({
    width: 12,
    height: 12,
    borderRadius: '50%',
    backgroundColor: color || theme.palette.warning.main, // Используем цвет из темы
}));