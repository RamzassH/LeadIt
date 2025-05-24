import { Draggable } from "react-beautiful-dnd";
import { TaskCard, TaskPlaceholder } from "./styled/styled";
import {Box, CardContent, Chip, IconButton, Typography} from "@mui/material";
import DeleteIcon from '@mui/icons-material/Delete';
import useTaskStore, { Task } from "@/store/TaskStore/store";

interface Props {
    task: Task,
    index: number,
}

export default function KanbanItem({task, index}: Props) {
    const {setTasks, tasks} = useTaskStore()
    // Удаление задачи
    const deleteTask = (taskId: number) => {
        setTasks(tasks.filter(task => task.id !== taskId));
    };

    return (
        <Draggable key={task.id} draggableId={`${task.id}`} index={index}>
            {(provided, snapshot) => (
            <>
                <TaskCard
                    ref={provided.innerRef}
                    {...provided.draggableProps}
                    {...provided.dragHandleProps}
                    style={{
                        ...provided.draggableProps.style,
                        opacity: snapshot.isDragging ? 0.8 : 1,
                    }}
                >
                    <CardContent>
                        {/* Заголовок-ссылка */}
                        <Box
                            component="a"
                            href={`/tasks/${task.id}`}
                            sx={{
                                textDecoration: 'none',
                                color: 'inherit',
                                '&:hover': {
                                    textDecoration: 'underline',
                                    color: 'primary.main'
                                }
                            }}
                            {...provided.dragHandleProps} // Добавляем возможность перетаскивания за заголовок
                        >
                            <Typography variant="h6" component="div" sx={{ mb: 1 }}>
                                {task.name}
                            </Typography>
                        </Box>

                        {/* Описание задачи */}
                        {task.description && (
                            <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
                                {task.description}
                            </Typography>
                        )}

                        {/* Временные метки */}
                        <Box sx={{
                            display: 'flex',
                            justifyContent: 'space-between',
                            mb: 1,
                            gap: 1
                        }}>
                            <Chip
                                size="small"
                                label={`План: ${task.allocated_time}ч`}
                                color="info"
                                variant="outlined"
                            />
                            <Chip
                                size="small"
                                label={`Факт: ${task.wasted_time || 0}ч`}
                                color={task.wasted_time ? "success" : "default"}
                                variant="outlined"
                            />
                        </Box>

                        {/* Статус и действия */}
                        <Box sx={{
                            display: 'flex',
                            justifyContent: 'space-between',
                            alignItems: 'center'
                        }}>
                            <Chip
                                label={task.active ? 'Активно' : 'Не активно'}
                                size="small"
                                color={task.active ? "success" : "error"}
                                variant={task.active ? "filled" : "outlined"}
                            />

                            <IconButton
                                size="small"
                                onClick={(e) => {
                                    e.stopPropagation();
                                    deleteTask(task.id);
                                }}
                                sx={{ ml: 'auto' }}
                            >
                                <DeleteIcon fontSize="small" />
                            </IconButton>
                        </Box>
                    </CardContent>
                </TaskCard>

                {snapshot.isDragging && (
                    <TaskPlaceholder style={{ height: 'calc(120rem/16)' }} />
                )}
            </>)}
        </Draggable>
    )
}