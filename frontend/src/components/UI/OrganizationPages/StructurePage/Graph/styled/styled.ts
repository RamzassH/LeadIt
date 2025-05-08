import { styled } from '@mui/material/styles';
import { Box, Typography, Button, Switch, Divider } from '@mui/material';

export const GraphContainer = styled(Box)(({ theme }) => ({
    backgroundColor: theme.palette.background.paper,
    borderRadius: theme.shape.borderRadius,
    boxShadow: theme.shadows[1],
    padding: theme.spacing(3),
    width: 'calc(100% - 4rem * 2)',
    height: 'calc(100% - 4rem * 2)',
    margin: theme.spacing(8),
}));

export const Header = styled(Box)(({ theme }) => ({
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: theme.spacing(3),
}));
