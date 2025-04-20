import React from 'react';

import AddIcon from '@mui/icons-material/Add';
import {Button, ListItem, ListItemText, Typography} from "@mui/material";
import {
    ProjectListItem,
    ProjectsHeader,
    StyledList,
    StyledPaper
} from "@/components/UI/OrganizationPages/ProjectsPage/ProjectList/styled/styled";

// Тип для проекта
export interface Project {
    id: string;
    name: string;
    description?: string;
}

// Пропсы компонента
interface ProjectListProps {
    projects: Project[];
    onProjectClick: (projectId: string) => void;
    onAddProject: () => void;
}



const ProjectList: React.FC<ProjectListProps> = ({
                                                     projects,
                                                     onProjectClick,
                                                     onAddProject
                                                 }) => {
    return (
        <StyledPaper elevation={3}>
            <ProjectsHeader>
                <Typography variant="h6" component="h2">
                    Проекты
                </Typography>
                <Button
                    variant="contained"
                    color="primary"
                    startIcon={<AddIcon />}
                    onClick={onAddProject}
                >
                    Добавить проект
                </Button>
            </ProjectsHeader>

            {projects.length === 0 ? (
                <Typography variant="body1" color="textSecondary">
                    Нет проектов. Нажмите "Добавить проект", чтобы создать первый.
                </Typography>
            ) : (
                <StyledList>
                    {projects.map((project) => (
                        <ListItem key={project.id} disablePadding>
                            <ProjectListItem onClick={() => onProjectClick(project.id)}>
                                <ListItemText
                                    primary={project.name}
                                    secondary={project.description}
                                    primaryTypographyProps={{ fontWeight: 'medium' }}
                                />
                            </ProjectListItem>
                        </ListItem>
                    ))}
                </StyledList>
            )}
        </StyledPaper>
    );
};

export default ProjectList;