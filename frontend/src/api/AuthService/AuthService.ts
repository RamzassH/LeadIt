import * as grpcWeb from 'grpc-web';
import { AuthClient } from './generate/AuthServiceClientPb';
import { LoginRequest, LoginResponse } from './generate/auth_pb';

// Создаем экземпляр клиента
const client = new AuthClient('localhost:8080', null, null);

// Функция для отправки запроса на сервер
export function loginUser(email: string, password: string): void {
    const request = new LoginRequest();
    request.setEmail(email);
    request.setPassword(password);

    // Внимание: Ошибка может быть типизирована как обычная ошибка, а не grpcWeb.Error
    client.login(request, {}, (error: grpcWeb.RpcError, response: LoginResponse) => {
        if (error) {
            console.error('Ошибка при авторизации:', error.message);
        } else {
            console.log('Авторизация успешна:', response.getToken());
        }
    });
}
