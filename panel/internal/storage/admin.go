package storage

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/lib/pq"
    "github.com/Reza-namvaran/Barf-Yar/panel/internal/auth"
)

func CreateAdmin(username, password string) error {
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        log.Fatal("DATABASE_URL environment variable is not set")
    }

    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        return err
    }
    defer db.Close()

    // verify connection
    if err = db.Ping(); err != nil {
        return err
    }

    hashedPassword, err := auth.HashPassword(password)
    if err != nil {
        return err
    }

    _, err = db.Exec(`
        INSERT INTO admins (username, password_hash)
        VALUES ($1, $2)
        ON CONFLICT (username) DO NOTHING`,
        username, hashedPassword)
    
    return err
}