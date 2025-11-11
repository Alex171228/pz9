package config

import (
    "log"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

type Config struct {
    DB_DSN     string
    BcryptCost int
}

func Load() Config {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment variables")
    }

    cost := 12
    if v := os.Getenv("BCRYPT_COST"); v != "" {
        if parsed, err := strconv.Atoi(v); err == nil && parsed >= 8 && parsed <= 16 {
            cost = parsed
        }
    }

    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        log.Fatal("DB_DSN is not set — please configure it in .env or environment")
    }

    return Config{
        DB_DSN:     dsn,
        BcryptCost: cost,
    }
}
