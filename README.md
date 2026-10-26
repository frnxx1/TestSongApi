# TestSongApi

REST API на go для для получения, записи, удаления информации про музыку.

## О проекте
Учебный проект по разработке бэкенда. Позволяет работать с информацией, так же подключена база данных.

docker run -it --name user_test -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=user -e POSTGRES_DB=db  -e POSTGRES_HOST=localhost -p 5432:5432  postgres
go run main.go
