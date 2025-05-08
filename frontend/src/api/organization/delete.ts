import axios, {AxiosResponse} from "axios";
import {source} from "@/api/data";


export async function deleteOrganizationAPI(id: number, token: string): Promise<AxiosResponse> {
    const response = await axios.delete(`${source}/v1/organizations/${id}`, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}
