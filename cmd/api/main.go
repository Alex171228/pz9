package main

import (
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"

    "pz9-auth/internal/http/handlers"
    "pz9-auth/internal/platform/config"
    "pz9-auth/internal/repo"
)

func main() {
    cfg := config.Load()

    db, err := repo.Open(cfg.DB_DSN)
    if err != nil {
        log.Fatal("db connect:", err)
    }

    if err := db.Exec("SET timezone TO 'UTC'").Error; err != nil {
        log.Println("warning: failed to set timezone to UTC:", err)
    }

    users := repo.NewUserRepo(db)
    if err := users.AutoMigrate(); err != nil {
        log.Fatal("migrate:", err)
    }

    auth := &handlers.AuthHandler{
        Users:      users,
        BcryptCost: cfg.BcryptCost,
    }

    r := chi.NewRouter()
    r.Post("/auth/register", auth.Register)
    r.Post("/auth/login", auth.Login)

    const addr = ":8081"
    log.Println("listening on", addr)
    log.Fatal(http.ListenAndServe(addr, r))
}
