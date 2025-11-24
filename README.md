# практическое задание 9
## Шишков А.Д. ЭФМО-02-21
## Тема
Реализация регистрации и входа пользователей. Хэширование паролей с bcrypt
## Цели
- Научиться безопасно хранить пароли (bcrypt), валидировать вход и обрабатывать ошибки.
- Реализовать эндпоинты POST /auth/register и POST /auth/login.
- Закрепить работу с БД (PostgreSQL + GORM или database/sql) и валидацией ввода.
- Подготовить основу для JWT-аутентификации в следующем ПЗ
### Структура проекта 

<img width="316" height="471" alt="изображение" src="https://github.com/user-attachments/assets/e65a0964-af51-4848-b90e-a0a9c23b55e6" /> 

### Запуск проекта
1. Склонировать репозиторий и перейти в папку проекта:
   ```bash
   git clone https://github.com/Alex171228/Pz8
   cd pz9
    ```
2. Создать файл .env:
      ```bash
   cp .env.example .env
    ```
   Откройте и укажите параметры подключения к PostgreSQL:
   ```bash
   BCRYPT_COST=12
   DB_DSN=postgres://<user>:<password>@localhost:5432/<dbname>?sslmode=disable
   ```
3. Установите зависимости
   ```bash
   go mod tidy
   ```
4. Запуск приложения
   ```bash
   go run ./cmd/api
   ```
### Примеры запросов и результат их выполнения
1. Регистрация пользователя
   
   <img width="768" height="605" alt="изображение" src="https://github.com/user-attachments/assets/1c4d614b-cd6f-45d3-953c-c0f2697e6166" /> 

3. Повторная регистрация с тем же e-mail 

   <img width="785" height="478" alt="изображение" src="https://github.com/user-attachments/assets/dc4ec2a1-a577-4be3-a7e9-971fc1b43415" /> 

4. Успешный вход

   <img width="772" height="588" alt="изображение" src="https://github.com/user-attachments/assets/1611d030-bfe6-4a50-8157-16a6a5a0cde3" /> 

5. Неверный пароль

   <img width="525" height="492" alt="изображение" src="https://github.com/user-attachments/assets/2d3b4814-a433-4bbe-b5dc-145e5cb5db3d" /> 

Для выполнения запросов приложен json для Postman
