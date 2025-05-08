import axios, {AxiosResponse} from "axios";
import {source} from "@/api/data";

export async function getOrganizationAPI(id: number, token: string): Promise<AxiosResponse> {
    const response = await axios.get(`${source}/v1/organizations/${id}`, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}

export async function getOrganizationsAPI(organizer_id: number, token: string): Promise<AxiosResponse> {
    const response = await axios.get(`${source}/v1/users/${organizer_id}/organizations`, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}