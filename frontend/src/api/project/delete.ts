import axios, {AxiosResponse} from "axios";
import {source} from "@/api/data";

export async function deleteProjectAPI(id: number, token: string): Promise<AxiosResponse> {
    const response = await axios.delete(`${source}/v1/projects/${id}`, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}
