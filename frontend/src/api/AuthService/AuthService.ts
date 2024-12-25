import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';

// Путь к вашему .proto файлу
const PROTO_PATH = './auth.proto';

// Загружаем описание протокола
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
});

// Получаем объект, который представляет гRPC сервисы
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition) as grpc.GrpcObject;

const authProto = protoDescriptor.auth as { AuthService: grpc.ServiceClientConstructor };

// Создаем клиент для взаимодействия с gRPC сервером
const client = new authProto.AuthService(
    'localhost:50051', // Адрес gRPC сервера
    grpc.credentials.createInsecure(), // Используем небезопасное соединение (для разработки)
);

// Функция для авторизации
function authenticate(username: string, password: string): void {
    const request = { username, password };

    // Вызываем метод Authenticate на сервере
    client.Authenticate(request, (error: grpc.ServiceError, response: any) => {
        if (error) {
            console.error('Error during authentication:', error);
        } else {
            if (response.success) {
                console.log('Authentication successful! Token:', response.token);
            } else {
                console.log('Authentication failed:', response.message);
            }
        }
    });
}

