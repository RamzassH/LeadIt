import axios, {AxiosResponse} from "axios";

const source = 'http://localhost:8080'

export async function loginAPI(data: {email: string, password: string}): Promise<AxiosResponse> {
    const article = {...data};
    const response = await axios.post(`${source}/v1/auth/login`, article)
    console.log(response)
    return response;
}