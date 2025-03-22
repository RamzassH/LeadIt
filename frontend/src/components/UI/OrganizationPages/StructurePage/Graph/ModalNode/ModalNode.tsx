import React, { useState } from 'react';
import { Modal, Box, Typography, Button, List, ListItem, ListItemText, IconButton, Chip } from '@mui/material';
import CloseIcon from '@mui/icons-material/Close'; // Иконка крестика
import { Node } from 'reactflow';
import useDevelopersStore from "@/components/UI/OrganizationPages/StructurePage/Graph/store";

// Предположим, что у нас есть список всех возможных ролей и проектов
const allRoles = ['Администратор', 'Разработчик', 'Дизайнер', 'Тестировщик'];
const allProjects = ['Проект 1', 'Проект 2', 'Проект 3'];

const ModalNode: React.FC<{ node: Node | null; onClose: () => void }> = ({ node, onClose }) => {
    const developer = useDevelopersStore(state => state.developers.find(item => item.id === node?.id));
    const {addRole, addProject, deleteRole, deleteProject} = useDevelopersStore();
    const [openRolesModal, setOpenRolesModal] = useState(false);
    const [openProjectsModal, setOpenProjectsModal] = useState(false);

    const handleAddRole = (role: string) => {
        if (developer) {
            addRole(developer.id, role);
        }
        setOpenRolesModal(false);
    };

    const handleAddProject = (project: string) => {
        if (developer) {
            addProject(developer.id, project);
        }
        setOpenProjectsModal(false);
    };

    const handleRemoveRole = (role: string) => {
        if (developer) {
            deleteRole(developer.id, role);
        }
    };

    const handleRemoveProject = (project: string) => {
        if (developer) {
            deleteRole(developer.id, project);
        }
    };

    return (
        <>
            <Modal open={Boolean(node)} onClose={onClose}>
                <Box sx={{
                    position: 'absolute',
                    top: '50%',
                    left: '50%',
                    transform: 'translate(-50%, -50%)',
                    width: 400,
                    bgcolor: 'background.paper',
                    boxShadow: 24,
                    p: 4,
                }}>
                    <Typography variant="h6" component="h2">
                        {node?.data.label}
                    </Typography>
                    <Typography sx={{ mt: 2 }}>
                        Роли:
                        <Button onClick={() => setOpenRolesModal(true)}>Добавить роль</Button>
                        <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1 }}>
                            {developer?.roles.map((role, index) => (
                                <Chip
                                    key={index}
                                    label={role}
                                    onDelete={() => handleRemoveRole(role)}
                                    deleteIcon={<CloseIcon />}
                                    sx={{
                                        '&:hover': {
                                            cursor: 'pointer',
                                        },
                                    }}
                                />
                            ))}
                        </Box>
                    </Typography>
                    <Typography sx={{ mt: 2 }}>
                        Проекты:
                        <Button onClick={() => setOpenProjectsModal(true)}>Добавить проект</Button>
                        <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1 }}>
                            {developer?.projects.map((project, index) => (
                                <Chip
                                    key={index}
                                    label={project}
                                    onDelete={() => handleRemoveProject(project)}
                                    deleteIcon={<CloseIcon />}
                                    sx={{
                                        '&:hover': {
                                            cursor: 'pointer',
                                        },
                                    }}
                                />
                            ))}
                        </Box>
                    </Typography>
                </Box>
            </Modal>

            {/* Модальное окно для добавления ролей */}
            <Modal open={openRolesModal} onClose={() => setOpenRolesModal(false)}>
                <Box sx={{
                    position: 'absolute',
                    top: '50%',
                    left: '50%',
                    transform: 'translate(-50%, -50%)',
                    width: 400,
                    bgcolor: 'background.paper',
                    boxShadow: 24,
                    p: 4,
                }}>
                    <Typography variant="h6">Выберите роль</Typography>
                    <List>
                        {allRoles.filter(role => !developer?.roles.includes(role)).map((role, index) => (
                            <ListItem component="button" key={index} onClick={() => handleAddRole(role)}>
                                <ListItemText primary={role} />
                            </ListItem>
                        ))}
                    </List>
                </Box>
            </Modal>

            {/* Модальное окно для добавления проектов */}
            <Modal open={openProjectsModal} onClose={() => setOpenProjectsModal(false)}>
                <Box sx={{
                    position: 'absolute',
                    top: '50%',
                    left: '50%',
                    transform: 'translate(-50%, -50%)',
                    width: 400,
                    bgcolor: 'background.paper',
                    boxShadow: 24,
                    p: 4,
                }}>
                    <Typography variant="h6">Выберите проект</Typography>
                    <List>
                        {allProjects.filter(project => !developer?.projects.includes(project)).map((project, index) => (
                            <ListItem component="button" key={index} onClick={() => handleAddProject(project)}>
                                <ListItemText primary={project} />
                            </ListItem>
                        ))}
                    </List>
                </Box>
            </Modal>
        </>
    );
};

export default ModalNode;