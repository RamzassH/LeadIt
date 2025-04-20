// Стилизованные компоненты
import {Box, Paper, styled} from "@mui/material";

export const BoardContainer = styled(Paper)(({ theme }) => ({
    padding: theme.spacing(3),
    borderRadius: theme.shape.borderRadius * 2,
    boxShadow: theme.shadows[3],
}));

export const BoardHeader = styled(Box)(({ theme }) => ({
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: theme.spacing(4),
}));

export const Column = styled(Paper)(({ theme }) => ({
    padding: theme.spacing(2),
    minWidth: 300,
    minHeight: 500,
    backgroundColor: theme.palette.grey[50],
    marginRight: theme.spacing(2),
    borderRadius: theme.shape.borderRadius * 2,
}));

export const ColumnHeader = styled(Box)(({ theme }) => ({
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: theme.spacing(2),
    paddingBottom: theme.spacing(1),
    borderBottom: `1px solid ${theme.palette.divider}`,
}));

export const TaskCard = styled(Paper)(({ theme }) => ({
    padding: theme.spacing(2),
    marginBottom: theme.spacing(2),
    borderRadius: theme.shape.borderRadius,
    cursor: 'pointer',
    '&:hover': {
        boxShadow: theme.shadows[2],
    },
}));
