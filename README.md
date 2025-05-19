# Pet-проект "Социальная сеть"

## Локальный запуск

### Требования

1. docker
2. docker-compose

### Подготовка окружения

```
# copy .env file
cp .env.example .env

# copy config file
cp config/example.config.yaml config/local.config.yaml
```

### Запуск
```
make run
```

## Миграции

Все миграции находятся в папке [database/migrations](./database/migrations/)

### Создание новой миграции
```
make migrate-new name="<name of migration>"
```

### Накатывание миграции
```
make migrate
```

### Откат миграции
```
make migrate-down
```