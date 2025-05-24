import {Fragment, useState} from 'react';
import {
    Box,
    Typography,
    List,
    ListItem,
    ListItemAvatar,
    Avatar,
    ListItemText,
    TextField,
    Button,
    Divider,
    IconButton
} from '@mui/material';
import { Send, Delete, Edit } from '@mui/icons-material';

interface Comment {
    id: number;
    user: string;
    avatar?: string;
    text: string;
    date: string;
    isEditable?: boolean;
}

const TaskComments = () => {
    const [comments, setComments] = useState<Comment[]>([
        {
            id: 1,
            user: 'Мария Сидорова',
            avatar: '/avatars/2.jpg',
            text: 'Начала работу над интерфейсом. Возникли вопросы по пункту 3 в требованиях.',
            date: '2023-05-16 10:30',
            isEditable: true
        },
        {
            id: 2,
            user: 'Алексей Петров',
            avatar: '/avatars/1.jpg',
            text: 'Пункт 3 можно пока опустить, сосредоточьтесь на основных разделах.',
            date: '2023-05-16 11:45',
            isEditable: false
        }
    ]);

    const [newComment, setNewComment] = useState('');
    const [editingCommentId, setEditingCommentId] = useState<number | null>(null);

    const handleAddComment = () => {
        if (newComment.trim()) {
            if (editingCommentId) {
                setComments(comments.map(c =>
                    c.id === editingCommentId ? { ...c, text: newComment } : c
                ));
                setEditingCommentId(null);
            } else {
                const newCommentObj: Comment = {
                    id: Date.now(),
                    user: 'Вы',
                    text: newComment,
                    date: new Date().toLocaleString(),
                    isEditable: true
                };
                setComments([...comments, newCommentObj]);
            }
            setNewComment('');
        }
    };

    const handleEditComment = (id: number) => {
        const comment = comments.find(c => c.id === id);
        if (comment) {
            setNewComment(comment.text);
            setEditingCommentId(id);
        }
    };

    const handleDeleteComment = (id: number) => {
        setComments(comments.filter(c => c.id !== id));
    };

    return (
        <Box>
            <Typography variant="h6" gutterBottom>
                Комментарии ({comments.length})
            </Typography>

            <List sx={{ maxHeight: 400, overflow: 'auto', mb: 2 }}>
                {comments.map((comment) => (
                    <Fragment key={comment.id}>
                        <ListItem alignItems="flex-start">
                            <ListItemAvatar>
                                <Avatar src={comment.avatar}>
                                    {comment.user.charAt(0)}
                                </Avatar>
                            </ListItemAvatar>
                            <ListItemText
                                primary={
                                    <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
                                        <Typography fontWeight="bold">
                                            {comment.user}
                                        </Typography>
                                        <Typography variant="caption" color="text.secondary">
                                            {comment.date}
                                        </Typography>
                                    </Box>
                                }
                                secondary={
                                    <>
                                        <Typography>{comment.text}</Typography>
                                        {comment.isEditable && (
                                            <Box sx={{ mt: 1 }}>
                                                <IconButton size="small" onClick={() => handleEditComment(comment.id)}>
                                                    <Edit fontSize="small" />
                                                </IconButton>
                                                <IconButton size="small" onClick={() => handleDeleteComment(comment.id)}>
                                                    <Delete fontSize="small" />
                                                </IconButton>
                                            </Box>
                                        )}
                                    </>
                                }
                            />
                        </ListItem>
                        <Divider variant="inset" component="li" />
                    </Fragment>
                ))}
            </List>

            <Box sx={{ display: 'flex', gap: 1 }}>
                <TextField
                    fullWidth
                    multiline
                    rows={3}
                    variant="outlined"
                    placeholder="Напишите комментарий..."
                    value={newComment}
                    onChange={(e) => setNewComment(e.target.value)}
                />
                <Button
                    variant="contained"
                    startIcon={<Send />}
                    onClick={handleAddComment}
                    sx={{ height: 'fit-content' }}
                >
                    {editingCommentId ? 'Обновить' : 'Отправить'}
                </Button>
            </Box>
        </Box>
    );
};

export default TaskComments;