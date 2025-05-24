"use client"
import React, {useEffect, useState} from 'react';
import {
    Container,
    Paper,
    Typography,
    Chip,
    Divider,
    Button,
    Box,
    Avatar,
    LinearProgress,
    Tabs,
    Tab
} from '@mui/material';
import { Edit, ArrowBack, PlayArrow, Pause, Check } from '@mui/icons-material';
import {useRouter} from "next/navigation";
import TaskHistory from "@/components/UI/TasksPage/TaskHistory/TaskHistory";
import TaskAttachments from "@/components/UI/TasksPage/TaskAttachments/TaskAttachments";
import TaskComments from "@/components/UI/TasksPage/TaskComments/TaskComments";

// Типы данных
interface Task {
    id: number;
    name: string;
    description: string;
    project_id: number;
    setter_id: number;
    solver_id: number;
    progress_bar_id: number;
    current_status_id: number;
    allocated_time: number;
    wasted_time: number;
    active: boolean;
    last_start_time: number;
}

interface User {
    id: number;
    name: string;
    avatar?: string;
}

interface Project {
    id: number;
    name: string;
}

interface Status {
    id: number;
    name: string;
    color: string;
}

const TaskPageContent = () => {
    const router = useRouter();
    const [tabValue, setTabValue] = useState(0);
    const [task, setTask] = useState<Task | null>(null);
    const [setter, setSetter] = useState<User | null>(null);
    const [solver, setSolver] = useState<User | null>(null);
    const [project, setProject] = useState<Project | null>(null);
    const [status, setStatus] = useState<Status | null>(null);

    useEffect(() => {
            const mockTask: Task = {
                id: 1,
                name: 'Разработка интерфейса панели управления',
                description: 'Необходимо создать интерфейс административной панели с графиками и статистикой',
                project_id: 1,
                setter_id: 1,
                solver_id: 2,
                progress_bar_id: 1,
                current_status_id: 2,
                allocated_time: 24,
                wasted_time: 8.5,
                active: true,
                last_start_time: Date.now() - 3600000 // 1 час назад
            };

            const mockSetter: User = {
                id: 1,
                name: 'Алексей Петров',
                avatar: '/avatars/1.jpg'
            };

            const mockSolver: User = {
                id: 2,
                name: 'Мария Сидорова',
                avatar: '/avatars/2.jpg'
            };

            const mockProject: Project = {
                id: 1,
                name: 'Панель администратора'
            };

            const mockStatus: Status = {
                id: 2,
                name: 'In Progress',
                color: 'primary'
            };

            setTask(mockTask);
            setSetter(mockSetter);
            setSolver(mockSolver);
            setProject(mockProject);
            setStatus(mockStatus);

    }, []);

    if (!task) return <LinearProgress />;

    // Расчет прогресса
    const progress = Math.min(Math.round((task.wasted_time / task.allocated_time) * 100), 100);
    const timeLeft = task.allocated_time - task.wasted_time;

    // Форматирование времени
    const formatHours = (hours: number) => {
        const fullHours = Math.floor(hours);
        const minutes = Math.round((hours - fullHours) * 60);
        return `${fullHours}ч ${minutes}м`;
    };

    // Время с последнего старта
    const getActiveTime = () => {
        if (!task.active) return 'Не активно';
        const hours = (Date.now() - task.last_start_time) / (1000 * 60 * 60);
        return `Активна ${formatHours(hours)}`;
    };

    const handleClick = (event: React.MouseEvent<HTMLElement>) => {
        setTask({...task, active: !task.active });
    };

    return (
        <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
            <Button
                startIcon={<ArrowBack />}
                onClick={() => router.back()}
                sx={{ mb: 2 }}
            >
                Назад
            </Button>

            <Paper elevation={3} sx={{ p: 3 }}>
                {/* Заголовок и действия */}
                <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 3 }}>
                    <Typography variant="h4" component="h1">
                        {task.name}
                    </Typography>
                    <Box>
                        {/*
                            <Button startIcon={<Edit />} sx={{ mr: 1 }}>
                                Редактировать
                            </Button>
                        */}
                        {
                            <Button
                                variant="contained"
                                startIcon={task.active ? <Pause /> : <PlayArrow />}
                                color={task.active ? 'warning' : 'success'}
                                onClick={handleClick}
                            >
                                {task.active ? 'Приостановить' : 'Запустить'}
                            </Button>
                        }
                    </Box>
                </Box>

                {/* Основная информация */}
                <Box sx={{ display: 'flex', gap: 4, mb: 3 }}>
                    <Box sx={{ flex: 1 }}>
                        <Typography variant="body1" paragraph>
                            {task.description}
                        </Typography>

                        <Box sx={{ display: 'flex', gap: 2, mb: 2 }}>
                            <Chip
                                label={`Проект: ${project?.name || 'Не указан'}`}
                                variant="outlined"
                                clickable
                                onClick={() => project && router.push(`/projects/${project.id}`)}
                            />
                            <Chip
                                label={`Статус: ${status?.name || 'Не указан'}`}
                                color={status?.color as any || 'default'}
                            />
                        </Box>
                    </Box>

                    <Box sx={{ width: 300 }}>
                        <Paper variant="outlined" sx={{ p: 2 }}>
                            <Typography variant="subtitle1" gutterBottom>
                                Временные метки
                            </Typography>
                            <Box sx={{ mb: 2 }}>
                                <Typography variant="body2">
                                    Выделено: {formatHours(task.allocated_time)}
                                </Typography>
                                <Typography variant="body2">
                                    Затрачено: {formatHours(task.wasted_time)}
                                </Typography>
                                <Typography variant="body2">
                                    Осталось: {formatHours(timeLeft)}
                                </Typography>
                            </Box>
                            <LinearProgress
                                variant="determinate"
                                value={progress}
                                sx={{ height: 8, borderRadius: 4, mb: 2 }}
                            />
                            <Typography variant="caption" display="block">
                                Прогресс: {progress}%
                            </Typography>
                        </Paper>
                    </Box>
                </Box>

                {/* Ответственные */}
                <Box sx={{ display: 'flex', gap: 4, mb: 3 }}>
                    <Paper variant="outlined" sx={{ p: 2, flex: 1 }}>
                        <Typography variant="subtitle1" gutterBottom>
                            Постановщик
                        </Typography>
                        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                            <Avatar src={setter?.avatar}>{setter?.name.charAt(0)}</Avatar>
                            <Typography>{setter?.name || 'Не назначен'}</Typography>
                        </Box>
                    </Paper>

                    <Paper variant="outlined" sx={{ p: 2, flex: 1 }}>
                        <Typography variant="subtitle1" gutterBottom>
                            Исполнитель
                        </Typography>
                        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                            <Avatar src={solver?.avatar}>{solver?.name.charAt(0)}</Avatar>
                            <Typography>{solver?.name || 'Не назначен'}</Typography>
                        </Box>
                    </Paper>

                    <Paper variant="outlined" sx={{ p: 2, flex: 1 }}>
                        <Typography variant="subtitle1" gutterBottom>
                            Состояние
                        </Typography>
                        <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                            <Chip
                                label={task.active ? 'Активна' : 'Не активна'}
                                color={task.active ? 'success' : 'error'}
                                size="small"
                            />
                            <Typography variant="body2">
                                {getActiveTime()}
                            </Typography>
                        </Box>
                    </Paper>
                </Box>

                {/* Табы с дополнительной информацией */}
                <Tabs value={tabValue} onChange={(_, newValue) => setTabValue(newValue)}>
                    <Tab label="История изменений" />
                    <Tab label="Вложенные файлы" />
                    <Tab label="Комментарии" />
                </Tabs>
                <Divider sx={{ mb: 2 }} />

                {/* Контент табов */}
                <Box sx={{ pt: 2 }}>
                    {tabValue === 0 && (
                        <TaskHistory/>
                    )}
                    {tabValue === 1 && (
                        <TaskAttachments/>
                    )}
                    {tabValue === 2 && (
                        <TaskComments/>
                    )}
                </Box>
            </Paper>
        </Container>
    );
};

export default TaskPageContent;