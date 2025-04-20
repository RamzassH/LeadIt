import {
    List,
    ListItemButton,
    styled,
    Paper,
    Box
} from '@mui/material';

// Стилизованные компоненты
export const StyledPaper = styled(Paper)(({ theme }) => ({
    width: 'calc(1100rem/16)',
    margin: theme.spacing(2),
    padding: theme.spacing(2),
    borderRadius: theme.shape.borderRadius,
    marginLeft: 'auto',
    marginRight: 'auto'
}));

export const ProjectsHeader = styled(Box)(({ theme }) => ({
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: theme.spacing(2),
}));

export const StyledList = styled(List)({
    width: '100%',
});

export const ProjectListItem = styled(ListItemButton)(({ theme }) => ({
    borderRadius: theme.shape.borderRadius,
    marginBottom: theme.spacing(1),
    '&:hover': {
        backgroundColor: theme.palette.action.hover,
    },
}));