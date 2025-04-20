"use client"
import React, {useEffect, useRef, useState} from "react";
import ProjectList, {Project} from "@/components/UI/OrganizationPages/ProjectsPage/ProjectList/ProjectList";
import CreateProjectModalWindow
    from "@/components/UI/OrganizationPages/ProjectsPage/CreateProjectModalWindow/CreateProjectModalWindow";
import useProjectStore from "@/store/ProjectsPageStore/store";
import {useRouter} from "next/navigation";

export default function Page() {
    const projectStore = useProjectStore(state => state.projects)
    const [isOpen, setOpen] = useState(false)
    const [projects, setProjects] = React.useState<Project[]>([])
    const router = useRouter()

    useEffect(() => {
        setProjects(projectStore)
    }, [projectStore]);

    const handleProjectClick = (projectId: string) => {
        router.push(`/organization/projects/${projectId}`)
    };

    const handleAddProject = () => {
        setOpen(true)
    };

    return (
        <>
            <ProjectList
                projects={projects}
                onProjectClick={handleProjectClick}
                onAddProject={handleAddProject}
            />
            <CreateProjectModalWindow open={isOpen} handleClose={() => {setOpen(false)}}/>
        </>
    );
}