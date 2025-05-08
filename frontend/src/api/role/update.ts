import axios, {AxiosResponse} from "axios";
import {source} from "@/api/data";

interface Permission {
    key: string;
    value: boolean;
}

interface Data {
    id: number;
    name: string;
    permissions:  Permission[];
    parent_role_id: number;
}

export async function updateRoleAPI(id: number, data: Data, token: string): Promise<AxiosResponse> {
    const article = {...data};
    const response = await axios.patch(`${source}/v1/roles/${id}`, article, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}