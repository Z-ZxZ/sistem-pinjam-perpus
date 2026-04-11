package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/config"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/handler"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/middleware"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/repository"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/service"
)

func main() {
	// Initialize DB
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	// Repositories
	userRepo := repository.NewUserRepositoryPostgres(db)
	bookRepo := repository.NewBookRepositoryPostgres(db)
	borrowRepo := repository.NewBorrowRepositoryPostgres(db)
	fineRepo := repository.NewFineRepositoryPostgres(db)

	// Services
	userService := service.NewUserService(userRepo)
	bookService := service.NewBookService(bookRepo)
	borrowService := service.NewBorrowService(borrowRepo, bookRepo, fineRepo)

	// Handlers
	userHandler := handler.NewUserHandler(userService)
	bookHandler := handler.NewBookHandler(bookService)
	borrowHandler := handler.NewBorrowHandler(borrowService)

	// Mux
	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("/auth/register", userHandler.Register)
	mux.HandleFunc("/auth/login", userHandler.Login)
	mux.HandleFunc("/books", bookHandler.Handle) // Handle both public list and admin CRUD
	mux.HandleFunc("/books/", bookHandler.Handle)

	// Protected routes
	mux.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(userHandler.Profile)))
	mux.Handle("/borrow", middleware.AuthMiddleware(http.HandlerFunc(borrowHandler.Borrow)))
	mux.Handle("/return", middleware.AuthMiddleware(http.HandlerFunc(borrowHandler.Return)))
	mux.Handle("/history", middleware.AuthMiddleware(http.HandlerFunc(borrowHandler.History)))
	mux.Handle("/fines", middleware.AuthMiddleware(http.HandlerFunc(borrowHandler.Fines)))

	// Admin routes
	mux.Handle("/admin/users", middleware.AuthMiddleware(middleware.AdminOnly(http.HandlerFunc(userHandler.List))))
	mux.Handle("/admin/borrows", middleware.AuthMiddleware(middleware.AdminOnly(http.HandlerFunc(borrowHandler.AllBorrows))))

	// Middleware: CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Adjust for production
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	// Logging wrapper
	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})

	port := os.Getenv("PORT")
	if port == "" { port = "8080" }

	log.Printf("server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, finalHandler); err != nil {
		log.Fatal(err)
	}
}
