package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "syscall"

    "database/sql"
    _ "github.com/lib/pq"
    "golang.org/x/term"
    "github.com/joho/godotenv"
    "github.com/Reza-namvaran/Barf-Yar/panel/internal/repository"
    "github.com/Reza-namvaran/Barf-Yar/panel/internal/service"
)

func main() {
    _ = godotenv.Load()

    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        log.Fatal("DATABASE_URL not set")
    }

    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter admin username: ")
    username, _ := reader.ReadString('\n')
    username = strings.TrimSpace(username)

    fmt.Print("Enter admin password: ")
    passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
    if err != nil {
        log.Fatal("Failed to read password:", err)
    }
    password := string(passwordBytes)
    fmt.Println()

    fmt.Print("Confirm password: ")
    confirmBytes, _ := term.ReadPassword(int(syscall.Stdin))
    confirm := string(confirmBytes)
    fmt.Println()

    if password != confirm {
        log.Fatal("Passwords do not match")
    }

    if username == "" || password == "" {
        log.Fatal("Username and password cannot be empty")
    }

    adminRepo := repository.NewAdminRepository(db)
    adminService := service.NewAdminService(adminRepo)

    if err := adminService.CreateAdmin(username, password); err != nil {
        log.Fatalf("Failed to create admin: %v", err)
    }

    log.Println("✅ Admin account created successfully")
}