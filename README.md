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
