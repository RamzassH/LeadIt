import axios, {AxiosResponse} from "axios";
import {source} from "@/api/data";

interface Data {
    id: number;
    name: string;
    description: string;
    image: string;
}

export async function updateOrganizationAPI(id: number, data: Data, token: string): Promise<AxiosResponse> {
    const article = {...data};
    const response = await axios.patch(`${source}/v1/organizations/${id}`, article, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}