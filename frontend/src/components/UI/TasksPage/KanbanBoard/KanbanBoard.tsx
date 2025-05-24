import { useState } from 'react';
import { DragDropContext, Droppable } from 'react-beautiful-dnd';
import {
    Typography,
    IconButton
} from '@mui/material';
import AddIcon from '@mui/icons-material/Add';
import {
    BoardContainer,
    ColumnHeader,
    ColumnWrapper,
    TaskList
} from './styled/styled';
import useTaskStore, {Status, Task} from "@/store/TaskStore/store";
import AddTaskModalWindow from "@/components/UI/TasksPage/AddTaskModalWindow/AddTaskModalWindow";
import KanbanItem from "@/components/UI/TasksPage/KanbanItem/KanbanItem";

const KanbanBoard = () => {
    const [open, setOpen] = useState(false);
    const [currentStatus, setCurrentStatus] = useState<Status>({id: -1, name: "Ошибка", progress_bar_id: -1, order: 0, notify_roles_ids: []});
    const [developers, setDevelopers] = useState([
        {
            id: 1,
            name: "Иванов иван Иванович",
        },
        {
            id: 2,
            name: "Петров Петр Петрович",
        },
        {
            id: 3,
            name: "Степанов Степан Степанович",
        },
    ]);
    const {status, tasks, setTasks} = useTaskStore()

    // Обработка перетаскивания
    const onDragEnd = (result: any) => {

        const { source, destination, draggableId } = result;

        if (!destination) return;

        if (
            source.droppableId === destination.droppableId &&
            source.index === destination.index
        ) {
            return;
        }

        const startColumn = status[source.droppableId];
        const finishColumn = status[destination.droppableId];

        // Перемещение внутри одной колонки
        if (startColumn === finishColumn) {
            return;
        }

    };

    // Добавление новой задачи
    const addTask = (task: Task) => {
        setTasks([...tasks, task]);
    };

    return (
        <DragDropContext onDragEnd={onDragEnd}>
            <BoardContainer>
                {status.map(item => (
                    <ColumnWrapper key={item.id} style={{order: item.order}}>
                        <ColumnHeader>
                            <Typography variant="h6">{item.name}</Typography>
                            <IconButton
                                size="small"
                                onClick={() => {setOpen(true); setCurrentStatus(item)}}
                            >
                                <AddIcon fontSize="small" />
                            </IconButton>
                        </ColumnHeader>

                        <Droppable droppableId={`${item.id}`} key={item.id}>
                            {(provided) => (
                                <TaskList
                                    ref={provided.innerRef}
                                    {...provided.droppableProps}
                                >
                                    {tasks.map((task, index) => (
                                        task.current_status_id == item.id ?
                                        <KanbanItem task={task} index={index} key={task.id}/> : null
                                    ))}
                                    {provided.placeholder}
                                </TaskList>
                            )}
                        </Droppable>
                    </ColumnWrapper>
                ))}
            </BoardContainer>
            <AddTaskModalWindow
                open={open}
                handleClose={() => setOpen(false)}
                status={currentStatus}
                callback={addTask}
                developers={developers}
            />
        </DragDropContext>
    );
};

export default KanbanBoard;