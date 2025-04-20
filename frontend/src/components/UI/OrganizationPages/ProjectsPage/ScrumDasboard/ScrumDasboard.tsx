import {useState} from "react";
import {Box} from "@mui/material";
import ScrumBoard, {Task} from "../ScrumBoard/ScrumBoard";
import RetrospectiveBoard from "../SprintRetrospective/SprintRetrospective";
import BurndownChart from "../BurndownChart/BurndownChart";


const currentSprint = {
    id: 'sprint-1',
    name: 'Sprint 15: Marketplace Features',
    startDate: '2023-05-01',
    endDate: '2023-05-14',
    goal: 'Implement checkout process and payment integration',
};

const initialTasks: Task[] = [
    {
        id: 'task-1',
        content: 'Design payment UI',
        status: 'done',
        assignee: { name: 'Alex' },
        points: 3,
    },
    {
        id: 'task-2',
        content: 'Implement Stripe API',
        status: 'in-progress',
        assignee: { name: 'Maria' },
        points: 5,
    },
];

const ScrumDashboard = () => {
    const [tasks, setTasks] = useState<Task[]>(initialTasks);

    const handleTaskUpdate = (updatedTasks: Task[]) => {
        setTasks(updatedTasks);
    };

    const handleAddTask = (status: Task['status'], content: string) => {
        const newTask = {
            id: `task-${Date.now()}`,
            content,
            status,
            points: Math.floor(Math.random() * 5) + 1,
        };
        setTasks([...tasks, newTask]);
    };

    return (
        <Box sx={{ p: 3 }}>
            <ScrumBoard
                sprint={currentSprint}
                tasks={tasks}
                onTaskUpdate={handleTaskUpdate}
                onAddTask={handleAddTask}
            />

            <BurndownChart
                sprintDays={14}
                idealData={[50, 45, 40, 35, 30, 25, 20, 15, 10, 5, 0]}
                actualData={[50, 48, 42, 40, 38, 35, 32, 28, 25, 20, 15]}
            />

            <RetrospectiveBoard />
        </Box>
    );
};

export default ScrumDashboard;