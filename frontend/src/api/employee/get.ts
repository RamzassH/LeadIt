import axios, {AxiosResponse} from "axios";
import {source} from "@/api/data";

export async function getEmployeeAPI(id: number, token: string): Promise<AxiosResponse> {
    const response = await axios.get(`${source}/v1/employees/${id}`, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}

export async function getEmployeesAPI(organization_id: number, token: string): Promise<AxiosResponse> {
    const response = await axios.get(`${source}/v1/organizations/${organization_id}/employees`, {
        headers: {
            "Authorization": token
        }
    })
    return response;
}