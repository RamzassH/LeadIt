"use client"
import React, {useEffect, useRef, useState} from "react";
import ProjectList from "@/components/UI/OrganizationPages/ProjectsPage/ProjectList/ProjectList";
import CreateProjectModalWindow
    from "@/components/UI/OrganizationPages/ProjectsPage/CreateProjectModalWindow/CreateProjectModalWindow";
import useProjectStore, {Project} from "@/store/ProjectsPageStore/store";
import {useRouter} from "next/navigation";
import {useFetching} from "@/hooks/useFetching";
import {getProjectsAPI} from "@/api/project/get";

export default function Page() {
    const projectStore = useProjectStore(state => state.projects)
    const serProjectStore = useProjectStore(state => state.setProjects)
    const [isOpen, setOpen] = useState(false)
    const [projects, setProjects] = React.useState<Project[]>([])
    const router = useRouter()
    const [getProjects, isLoadingRequest, errorProjectRequest] = useFetching(async () => {
        const response = await getProjectsAPI(5, "")
        setProjects(response.data)
    })

    useEffect(() => {
        getProjects()
    }, []);

    useEffect(() => {
        setProjects(projectStore)
    }, [projectStore]);

    const handleProjectClick = (projectId: number) => {
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