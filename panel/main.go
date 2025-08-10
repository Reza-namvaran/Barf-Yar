package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Reza-namvaran/Barf-Yar/panel/internal/config"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/di"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	config.Load()

	container := di.NewContainer()

	db, err := sql.Open("postgres", config.DBURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	container.SetDB(db)

	handlers := container.GetHandlers()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Setup routes
	http.HandleFunc("/", handlers.LoginPage)
	http.HandleFunc("/api/login", handlers.Login)
	http.HandleFunc("/api/logout", handlers.Logout)
	http.HandleFunc("/dashboard", handlers.Dashboard)

	log.Printf("Server starting on port %s", config.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.ServerPort, nil))
}
