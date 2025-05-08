import axios, {AxiosResponse} from "axios";
import {source} from "@/api/data";

interface Data {
    id: number;
    role_id: number;
}

export async function updateEmployeeAPI(id: number, role_id: number, data: Data, token: string): Promise<AxiosResponse> {
    const article = {...data};
    const response = await axios.patch(`${source}/v1/employees/${id}/role/${role_id}`, article, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}