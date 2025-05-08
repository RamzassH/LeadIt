import axios, {AxiosResponse} from "axios";
import {source} from "@/api/data";

export async function getRoleAPI(id: number, token: string): Promise<AxiosResponse> {
    const response = await axios.get(`${source}/v1/roles/${id}`, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}

export async function getRolesAPI(organization_id: number, token: string): Promise<AxiosResponse> {
    const response = await axios.get(`${source}/v1/organizations/${organization_id}/roles`, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}