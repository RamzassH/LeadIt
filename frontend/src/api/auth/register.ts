import axios, {AxiosResponse} from "axios";

const source = 'http://localhost:8080'

export async function registerAPI(data: {name: string, surname: string, email: string, password: string}): Promise<AxiosResponse> {
    const article = {...data};
    const response = await axios.post(`${source}/v1/auth/register`,article)
    console.log(response)
    return response;
}