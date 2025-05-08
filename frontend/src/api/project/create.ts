import axios, {AxiosResponse} from "axios";
import {source} from "@/api/data";

interface Data {
    name: string;
    description: string;
    organization_id: number;
    image: string;
}

export async function createProjectAPI(data: Data, token: string): Promise<AxiosResponse> {
    const article = {...data};
    const response = await axios.post(`${source}/v1/projects`, article, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}