# Pet-проект "Социальная сеть"

## Локальный запуск

### Подготовка окружения

Копируем файл `.env.example` в `.env`. При необходимости меняем значения переменных

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