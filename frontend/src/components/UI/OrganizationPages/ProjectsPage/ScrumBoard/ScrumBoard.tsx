import React from 'react';
import { styled } from '@mui/material/styles';
import {
    Paper,
    Typography,
    Button,
    Box,
    Avatar,
    Chip,
    IconButton,
    TextField
} from '@mui/material';
import {
    Add as AddIcon,
    DragIndicator as DragIcon,
    MoreVert as MoreIcon,
    Check as CheckIcon
} from '@mui/icons-material';
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd';
import {
    BoardContainer,
    BoardHeader,
    Column,
    ColumnHeader, TaskCard
} from "@/components/UI/OrganizationPages/ProjectsPage/ScrumBoard/styled/styled";

// Типы для Scrum-доски
export interface Task {
    id: string;
    content: string;
    assignee?: {
        name: string;
        avatar?: string;
    };
    points?: number;
    status: 'backlog' | 'todo' | 'in-progress' | 'review' | 'done';
}

interface Sprint {
    id: string;
    name: string;
    startDate: string;
    endDate: string;
    goal: string;
}

interface ScrumBoardProps {
    sprint: Sprint;
    tasks: Task[];
    onTaskUpdate: (updatedTasks: Task[]) => void;
    onAddTask: (status: Task['status'], content: string) => void;
}

const ScrumBoard: React.FC<ScrumBoardProps> = ({
                                                   sprint,
                                                   tasks,
                                                   onTaskUpdate,
                                                   onAddTask,
                                               }) => {
    const [newTaskContent, setNewTaskContent] = React.useState('');
    const [addingTaskTo, setAddingTaskTo] = React.useState<Task['status'] | null>(null);

    const onDragEnd = (result: any) => {
        if (!result.destination) return;

        const updatedTasks = Array.from(tasks);
        const [removed] = updatedTasks.splice(result.source.index, 1);
        removed.status = result.destination.droppableId as Task['status'];
        updatedTasks.splice(result.destination.index, 0, removed);

        onTaskUpdate(updatedTasks);
    };

    const handleAddTask = (status: Task['status']) => {
        if (newTaskContent.trim()) {
            onAddTask(status, newTaskContent);
            setNewTaskContent('');
            setAddingTaskTo(null);
        }
    };

    const columnConfig = [
        { id: 'backlog', title: 'Backlog', color: 'secondary' },
        { id: 'todo', title: 'To Do', color: 'info' },
        { id: 'in-progress', title: 'In Progress', color: 'warning' },
        { id: 'review', title: 'Review', color: 'primary' },
        { id: 'done', title: 'Done', color: 'success' },
    ];

    return (
        <BoardContainer elevation={3}>
            <BoardHeader>
                <Box>
                    <Typography variant="h5" component="h2">
                        {sprint.name}
                    </Typography>
                    <Typography variant="subtitle1" color="textSecondary">
                        {sprint.startDate} - {sprint.endDate}
                    </Typography>
                    <Typography variant="body1">{sprint.goal}</Typography>
                </Box>
                <Box>
                    <Button variant="contained" color="primary">
                        Start Sprint
                    </Button>
                </Box>
            </BoardHeader>

            <DragDropContext onDragEnd={onDragEnd}>
                <Box display="flex" overflow="auto" py={2}>
                    {columnConfig.map((column) => (
                        <Droppable key={column.id} droppableId={column.id}>
                            {(provided) => (
                                <Column
                                    ref={provided.innerRef}
                                    {...provided.droppableProps}
                                    elevation={0}
                                >
                                    <ColumnHeader>
                                        <Chip
                                            label={column.title}
                                            color={column.color as any}
                                            size="small"
                                        />
                                        <IconButton
                                            size="small"
                                            onClick={() => setAddingTaskTo(column.id as Task['status'])}
                                        >
                                            <AddIcon fontSize="small" />
                                        </IconButton>
                                    </ColumnHeader>

                                    {addingTaskTo === column.id && (
                                        <Box mb={2} display="flex" alignItems="center">
                                            <TextField
                                                fullWidth
                                                size="small"
                                                placeholder="Task description"
                                                value={newTaskContent}
                                                onChange={(e) => setNewTaskContent(e.target.value)}
                                            />
                                            <IconButton onClick={() => handleAddTask(column.id as Task['status'])}>
                                                <CheckIcon color="primary" />
                                            </IconButton>
                                        </Box>
                                    )}

                                    {tasks
                                        .filter((task) => task.status === column.id)
                                        .map((task, index) => (
                                            <Draggable key={task.id} draggableId={task.id} index={index}>
                                                {(provided) => (
                                                    <TaskCard
                                                        ref={provided.innerRef}
                                                        {...provided.draggableProps}
                                                        elevation={1}
                                                    >
                                                        <Box {...provided.dragHandleProps}>
                                                            <DragIcon color="action" />
                                                        </Box>
                                                        <Typography variant="body2">{task.content}</Typography>
                                                        {task.assignee && (
                                                            <Box mt={1} display="flex" alignItems="center">
                                                                <Avatar
                                                                    src={task.assignee.avatar}
                                                                    sx={{ width: 24, height: 24, mr: 1 }}
                                                                >
                                                                    {task.assignee.name.charAt(0)}
                                                                </Avatar>
                                                                <Typography variant="caption">
                                                                    {task.assignee.name}
                                                                </Typography>
                                                            </Box>
                                                        )}
                                                        {task.points && (
                                                            <Chip
                                                                label={`${task.points} pts`}
                                                                size="small"
                                                                sx={{ mt: 1 }}
                                                            />
                                                        )}
                                                    </TaskCard>
                                                )}
                                            </Draggable>
                                        ))}
                                    {provided.placeholder}
                                </Column>
                            )}
                        </Droppable>
                    ))}
                </Box>
            </DragDropContext>
        </BoardContainer>
    );
};

export default ScrumBoard;