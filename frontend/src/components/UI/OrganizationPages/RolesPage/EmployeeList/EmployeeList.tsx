import React, { useState } from 'react';
import {
    List,
    ListItem,
    ListItemText,
    ListItemSecondaryAction,
    IconButton,
    Button,
    TextField,
    Box,
    Paper,
    Typography,
    Avatar
} from '@mui/material';
import { Add as AddIcon, Delete as DeleteIcon } from '@mui/icons-material';

interface Employee {
    id: string;
    name: string;
    position: string;
    avatar?: string;
}

const EmployeeList: React.FC = () => {
    const [employees, setEmployees] = useState<Employee[]>([
        { id: '1', name: 'Иван Петров', position: 'Frontend Developer', avatar: '/avatars/1.jpg' },
        { id: '2', name: 'Мария Сидорова', position: 'UX Designer', avatar: '/avatars/2.jpg' },
        { id: '3', name: 'Алексей Иванов', position: 'Backend Developer' }
    ]);

    const [newEmployee, setNewEmployee] = useState<Omit<Employee, 'id'>>({
        name: '',
        position: ''
    });

    const [isAdding, setIsAdding] = useState(false);

    const handleAddEmployee = () => {
        if (newEmployee.name.trim() && newEmployee.position.trim()) {
            setEmployees([
                ...employees,
                {
                    id: Date.now().toString(),
                    ...newEmployee
                }
            ]);
            setNewEmployee({ name: '', position: '' });
            setIsAdding(false);
        }
    };

    const handleDeleteEmployee = (id: string) => {
        setEmployees(employees.filter(employee => employee.id !== id));
    };

    return (
        <Paper elevation={3} sx={{ p: 3, maxWidth: 600, margin: '0 auto' }}>
            <Box display="flex" justifyContent="space-between" alignItems="center" mb={2}>
                <Typography variant="h6" component="h2">
                    Список сотрудников
                </Typography>
                <Button
                    variant="contained"
                    color="primary"
                    startIcon={<AddIcon />}
                    onClick={() => setIsAdding(true)}
                    disabled={isAdding}
                >
                    Добавить
                </Button>
            </Box>

            {isAdding && (
                <Box mb={3} p={2} sx={{ border: '1px dashed', borderColor: 'primary.main', borderRadius: 1 }}>
                    <TextField
                        label="ФИО"
                        fullWidth
                        margin="normal"
                        value={newEmployee.name}
                        onChange={(e) => setNewEmployee({...newEmployee, name: e.target.value})}
                    />
                    <TextField
                        label="Должность"
                        fullWidth
                        margin="normal"
                        value={newEmployee.position}
                        onChange={(e) => setNewEmployee({...newEmployee, position: e.target.value})}
                    />
                    <Box mt={2} display="flex" justifyContent="flex-end" gap={1}>
                        <Button variant="outlined" onClick={() => setIsAdding(false)}>
                            Отмена
                        </Button>
                        <Button
                            variant="contained"
                            color="primary"
                            onClick={handleAddEmployee}
                            disabled={!newEmployee.name.trim() || !newEmployee.position.trim()}
                        >
                            Сохранить
                        </Button>
                    </Box>
                </Box>
            )}

            <List>
                {employees.length === 0 ? (
                    <Typography variant="body1" color="textSecondary" textAlign="center" py={2}>
                        Список сотрудников пуст
                    </Typography>
                ) : (
                    employees.map((employee) => (
                        <ListItem key={employee.id} divider>
                            <Avatar
                                src={employee.avatar}
                                sx={{ width: 56, height: 56, mr: 2 }}
                            >
                                {employee.name.charAt(0)}
                            </Avatar>
                            <ListItemText
                                primary={employee.name}
                                secondary={employee.position}
                                primaryTypographyProps={{ fontWeight: 'medium' }}
                            />
                            <ListItemSecondaryAction>
                                <IconButton
                                    edge="end"
                                    aria-label="delete"
                                    onClick={() => handleDeleteEmployee(employee.id)}
                                    color="error"
                                >
                                    <DeleteIcon />
                                </IconButton>
                            </ListItemSecondaryAction>
                        </ListItem>
                    ))
                )}
            </List>
        </Paper>
    );
};

export default EmployeeList;