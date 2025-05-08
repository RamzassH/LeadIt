import axios, {AxiosResponse} from "axios";
import {source} from "@/api/data";

interface Permission {
    key: string;
    value: boolean;
}

interface Data {

    id: number;
    name: string;
    organization_id: number;
    permissions:  Permission[];
    parent_role_id: number;
}

export async function createRoleAPI(data: Data, token: string): Promise<AxiosResponse> {
    const article = {...data};
    const response = await axios.post(`${source}/v1/roles`, article, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}