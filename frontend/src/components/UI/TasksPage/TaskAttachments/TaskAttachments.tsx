import { Box, Typography, List, ListItem, ListItemIcon, ListItemText, Button } from '@mui/material';
import { InsertDriveFile, Image, PictureAsPdf, GetApp } from '@mui/icons-material';

interface Attachment {
    id: number;
    name: string;
    size: string;
    type: 'image' | 'document' | 'pdf' | 'other';
    date: string;
}

const TaskAttachments = () => {
    const attachments: Attachment[] = [
        { id: 1, name: 'Макет интерфейса.png', size: '2.4 MB', type: 'image', date: '2023-05-15' },
        { id: 2, name: 'Требования.pdf', size: '1.8 MB', type: 'pdf', date: '2023-05-16' },
        { id: 3, name: 'Описание задачи.docx', size: '0.5 MB', type: 'document', date: '2023-05-15' }
    ];

    const getFileIcon = (type: string) => {
        switch (type) {
            case 'image': return <Image color="primary" />;
            case 'pdf': return <PictureAsPdf color="error" />;
            case 'document': return <InsertDriveFile color="action" />;
            default: return <InsertDriveFile />;
        }
    };

    return (
        <Box>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 2 }}>
                <Typography variant="h6">Вложенные файлы ({attachments.length})</Typography>
                <Button variant="contained" size="small">
                    Добавить файл
                </Button>
            </Box>

            <List sx={{ maxHeight: 400, overflow: 'auto' }}>
                {attachments.map((file) => (
                    <ListItem
                        key={file.id}
                        secondaryAction={
                            <Button startIcon={<GetApp />} size="small">
                                Скачать
                            </Button>
                        }
                    >
                        <ListItemIcon>
                            {getFileIcon(file.type)}
                        </ListItemIcon>
                        <ListItemText
                            primary={file.name}
                            secondary={`${file.size} · ${file.date}`}
                        />
                    </ListItem>
                ))}
            </List>
        </Box>
    );
};

export default TaskAttachments;