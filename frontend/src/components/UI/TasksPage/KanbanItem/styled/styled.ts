import {styled} from "@mui/material/styles";
import {Card} from "@mui/material";

export const TaskCard = styled(Card)(({ theme }) => ({
    cursor: 'grab',
    '&:active': {
        cursor: 'grabbing',
        boxShadow: theme.shadows[8] // Более заметный эффект при перетаскивании
    },
    width: '100%', // Занимает всю ширину
    marginBottom: '8px', // Отступ между задачами
    transition: 'transform 0.2s',
}));

export const TaskContent = styled('div')({
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center'
});

export const AddTaskContainer = styled('div')({
    display: 'flex',
    gap: '8px',
    alignItems: 'center'
});

// 4. Добавляем стиль для placeholder (важно!)
export const TaskPlaceholder = styled('div')({
    backgroundColor: 'rgba(0, 0, 0, 0.05)',
    borderRadius: '4px',
    marginBottom: '8px'
});