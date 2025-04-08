import axios, {AxiosResponse} from "axios";

const source = 'http://localhost:8080'

export async function verifyAPI(data: {code: string}): Promise<AxiosResponse> {
    const article = {...data};
    const response = await axios.post(`${source}/v1/auth/verify`,article)
    console.log(response)
    return response;
}