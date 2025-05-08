import axios, {AxiosResponse} from "axios";
import {source} from "@/api/data";

export async function deleteRoleAPI(id: number, token: string): Promise<AxiosResponse> {
    const response = await axios.delete(`${source}/v1/roles/${id}`, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}
