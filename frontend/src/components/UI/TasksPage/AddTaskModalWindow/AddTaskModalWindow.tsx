import {Status, Task} from "@/store/TaskStore/store";
import ModalWindow from "@/components/UI/ModalWindowTemplate/ModalWindow";
import React, {useState} from "react";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import {Box, FormControl, InputAdornment, InputLabel, MenuItem, Select, SelectChangeEvent} from "@mui/material";

interface AddTaskModalWindowProps {
    open: boolean;
    status: Status;
    callback: Function;
    handleClose: () => void;
    developers: {id: number; name: string}[];
}

export default function AddTaskModalWindow({status, open, handleClose, callback, developers}: AddTaskModalWindowProps) {
    const [currentTaskId, setId] = useState<number>(1);
    const [form, setForm] = useState<{name: string, description: string, solver_id: number, allocated_time: number}>({
        name: "",
        description: "",
        solver_id:  -1,
        allocated_time: -1,
    });

    const handleDeveloperChange = (e: SelectChangeEvent) => {
        setForm(prev => ({ ...prev, solver_id: Number(e.target.value) }));
    };

    const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        setForm({
            ...form,
            [name]: value
        });
    }
    const handleSubmit = () => {
        callback({
            id: currentTaskId,
            name: form.name,
            description: form.description,
            project_id: 1,
            setter_id:  1,
            solver_id: Number(form.solver_id),
            progress_bar_id: 1,
            current_status_id: status.id,
            allocated_time: Number(form.allocated_time),
            wasted_time: 0,
            active: false,
            last_start_time: 0,
        })
        setId(currentTaskId + 1)
        handleClose();
    }

    const close = () => {
        handleClose();
    }

    const actions = () => {
        return (
            <div>
                <Button onClick={close}>Отмена</Button>
                <Button onClick={() => {handleSubmit()}} variant="contained" color="primary">
                    Добавить
                </Button>
            </div>
        )
    }
    return (
        <ModalWindow open={open} handleClose={() => {}} actions={actions()} title={`Добавить задачу в ${status.name}`}>
            <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
                <TextField
                    autoFocus
                    margin="dense"
                    name="name"
                    label="Имя задачи"
                    type="text"
                    fullWidth
                    value={form.name}
                    onChange={handleInputChange}
                />

                <FormControl fullWidth margin="dense">
                    <InputLabel id="developer-select-label">Исполнитель</InputLabel>
                    <Select
                        labelId="developer-select-label"
                        id="solver_id"
                        value={`${form.solver_id}`}
                        label="Исполнитель"
                        onChange={handleDeveloperChange}
                    >
                        {developers.map(dev => (
                            <MenuItem key={dev.id} value={dev.id}>
                                {dev.name}
                            </MenuItem>
                        ))}
                    </Select>
                </FormControl>

                <TextField
                    margin="dense"
                    name="allocated_time"
                    label="Оценка времени (часы)"
                    type="number"
                    fullWidth
                    InputProps={{
                        endAdornment: <InputAdornment position="end">ч</InputAdornment>,
                    }}
                    inputProps={{ min: 1 }}
                    value={form.allocated_time}
                    onChange={handleInputChange}
                />
                <TextField
                    margin="dense"
                    name="description"
                    label="Описание задачи"
                    type="text"
                    fullWidth
                    multiline
                    rows={4}
                    value={form.description}
                    onChange={handleInputChange}
                />
            </Box>
        </ModalWindow>
    )
}