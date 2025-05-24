import { Box, Typography, List, ListItem, ListItemText, Divider, Chip } from '@mui/material';
import { Event, Schedule, Person, CheckCircle } from '@mui/icons-material';
import { Fragment } from 'react';

interface HistoryItem {
    id: number;
    date: string;
    user: string;
    action: string;
    changes: string[];
}

const TaskHistory = () => {
    const historyData: HistoryItem[] = [
        {
            id: 1,
            date: '2023-05-15 14:30',
            user: 'Алексей Петров',
            action: 'Создал задачу',
            changes: ['Статус: To do']
        },
        {
            id: 2,
            date: '2023-05-16 09:15',
            user: 'Мария Сидорова',
            action: 'Взяла в работу',
            changes: ['Статус: In Progress', 'Исполнитель: Мария Сидорова']
        },
        {
            id: 3,
            date: '2023-05-18 16:45',
            user: 'Мария Сидорова',
            action: 'Обновила задачу',
            changes: ['Затраченное время: 4ч → 8ч']
        }
    ];

    const getActionIcon = (action: string) => {
        switch (action.toLowerCase()) {
            case 'создал задачу': return <Event color="info" />;
            case 'взяла в работу': return <Person color="primary" />;
            case 'обновила задачу': return <Schedule color="warning" />;
            default: return <CheckCircle color="success" />;
        }
    };

    return (
        <Box sx={{ maxHeight: 400, overflow: 'auto' }}>
            <List dense>
                {historyData.map((item) => (
                    <Fragment key={item.id}>
                        <ListItem alignItems="flex-start">
                            <Box sx={{ mr: 2 }}>{getActionIcon(item.action)}</Box>
                            <ListItemText
                                primary={
                                    <>
                                        <Typography component="span" fontWeight="bold">
                                            {item.user}
                                        </Typography>
                                        {' ' + item.action}
                                    </>
                                }
                                secondary={
                                    <>
                                        <Typography variant="caption" color="text.secondary">
                                            {item.date}
                                        </Typography>
                                        <Box sx={{ mt: 1 }}>
                                            {item.changes.map((change, i) => (
                                                <Chip
                                                    key={i}
                                                    label={change}
                                                    size="small"
                                                    sx={{ mr: 1, mb: 1 }}
                                                />
                                            ))}
                                        </Box>
                                    </>
                                }
                            />
                        </ListItem>
                        <Divider variant="inset" component="li" />
                    </Fragment>
                ))}
            </List>
        </Box>
    );
};

export default TaskHistory;