import axios, {AxiosResponse} from "axios";
import {source} from "@/api/data";

export async function getProjectAPI(id: number, token: string): Promise<AxiosResponse> {
    const response = await axios.get(`${source}/v1/projects/${id}`, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}

export async function getProjectsAPI(organization_id: number, token: string): Promise<AxiosResponse> {
    const response = await axios.get(`${source}/v1/organizations/${organization_id}/projects`, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}