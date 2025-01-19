# Используем официальный образ Envoy
FROM envoyproxy/envoy:v1.26.0

# Копируем конфигурационный файл в контейнер
COPY envoy.yaml /etc/envoy/envoy.yaml

# Указываем команду запуска Envoy
CMD ["envoy", "-c", "/etc/envoy/envoy.yaml", "--service-cluster", "envoy-service"]

