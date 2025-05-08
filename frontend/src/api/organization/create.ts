import axios, {AxiosResponse} from "axios";
import {source} from "@/api/data";

interface Data {
    name: string;
    description: string;
    image: string;
}

export async function createOrganizationAPI(data: Data, token: string): Promise<AxiosResponse> {
    const article = {...data};
    const response = await axios.post(`${source}/v1/organizations`, article, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}