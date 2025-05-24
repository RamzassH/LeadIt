// Стилизованные компоненты
import {styled} from "@mui/material/styles";
import {Card, Paper} from "@mui/material";

export const BoardContainer = styled('div')({
    display: 'flex',
    padding: '16px',
    gap: '16px',
    overflowX: 'auto',
    height: 'calc(100vh - 100px)',
    alignItems: 'flex-start' // Важно: выравниваем колонки по верху
});

export const ColumnWrapper = styled(Paper)(({ theme }) => ({
    minWidth: '450px',
    padding: '16px',
    backgroundColor: theme.palette.grey[100],
    display: 'flex',
    flexDirection: 'column',
    gap: '12px'
}));

export const ColumnHeader = styled('div')({
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center'
});

export const TaskList = styled('div')({
    display: 'flex',
    flexDirection: 'column', // Гарантируем вертикальное направление
    gap: '8px',
    flexGrow: 1,
    minHeight: '100px', // Минимальная высота для пустой колонки
    padding: '4px',
    overflowY: 'auto',
    width: '100%' // Занимаем всю ширину колонки
});
