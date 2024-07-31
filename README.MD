# Todo List

## Описание
Проект "Todo List" представляет собой веб-приложение для управления списком задач. Он включает в себя сервис для управления задачами, реализованный на Go.

## Требования
- Docker
- Docker Compose
- Go 1.22.4

## Установка

1. Клонируйте репозиторий:
    ```sh
    git clone https://github.com/sterino/ToDo-List.git
    cd todo-list
    ```

2. Скопируйте файл `.env.example` в `.env` и настройте переменные окружения:
    ```sh
    cp .env.example .env
    ```

## Запуск

### С использованием Docker Compose

1. Соберите и запустите контейнеры:
    ```sh
    docker-compose up --build
    ```

2. Приложение будет доступно по адресу `http://localhost:8080`.

### Локальный запуск

1. Установите зависимости:
    ```sh
    go mod download
    ```

2. Соберите и запустите приложение:
    ```sh
    go build -o main .
    ./main
    ```

3. Приложение будет доступно по адресу `http://localhost:8080`.

## Использование

### API Документация

API документация доступна по адресу `http://localhost:8080/swagger/index.html`.

### Команды Makefile

- Остановить контейнеры:
    ```sh
    make down
    ```

- Просмотр логов:
    ```sh
    make logs
    ```

- Перезапуск контейнеров:
    ```sh
    make restart
    ```