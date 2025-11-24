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
   Переменная APP_ADDR не применяется, так как сервер работает строго на порту 8081.
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
   
Для выполнения запросов приложен json для Postman https://github.com/Alex171228/pz9/blob/main/pz9-auth-postman-collection.json

### Фрагменты кода
Вызов bcrypt.GenerateFromPassword
   ```go
   // Хэширование пароля при регистрации
   hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), h.BcryptCost)
   if err != nil {
       writeErr(w, http.StatusInternalServerError, "hash_failed")
       return
   }
   ```
Вызов bcrypt.CompareHashAndPassword
   ```go
   // Сравнение хэша и введённого пароля при логине
   if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(in.Password)) != nil {
       writeErr(w, http.StatusUnauthorized, "invalid_credentials")
       return
   }
   ```
### SQL/миграции 
В проекте используется автоматическая миграция схемы базы данных с помощью GORM.
При запуске приложения вызывается метод AutoMigrate(), который автоматически создаёт таблицу users, если она отсутствует.

В cmd/api/main.go
   ```go
    users := repo.NewUserRepo(db)
    if err := users.AutoMigrate(); err != nil {
        log.Fatal("migrate:", err)
    }
   ```
В internal/repo/user_repo.go
   ```go
   func (r *UserRepo) AutoMigrate() error {
       return r.db.AutoMigrate(&core.User{})
   }
   ```
Таблица users 

<img width="425" height="183" alt="изображение" src="https://github.com/user-attachments/assets/1e3a9ca0-44b5-4255-b1fd-64d1ba5d440c" /> 

<img width="1234" height="334" alt="изображение" src="https://github.com/user-attachments/assets/7760a6aa-0838-4e52-afe5-e755892d1f6a" /> 

### Почему нельзя хранить пароли в открытом виде

Хранение паролей в текстовом виде является критической ошибкой, поскольку приводит к немедленной компрометации всех учётных записей при утечке базы данных. Злоумышленник получает полный доступ ко всем аккаунтам, особенно с учётом того, что пользователи часто используют один и тот же пароль в разных сервисах. 
Кроме того, это нарушает требования информационной безопасности и законодательства о защите персональных данных. 
