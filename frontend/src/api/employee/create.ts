import axios, {AxiosResponse} from "axios";
import {source} from "@/api/data";

interface Data {
    user_id: number,
    organization_id: number
}

export async function createEmployeeAPI(data: Data, token: string): Promise<AxiosResponse> {
    const article = {...data};
    const response = await axios.post(`${source}/v1/employees`, article, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}