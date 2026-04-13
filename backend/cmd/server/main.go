package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/config"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/handler"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/middleware"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/repository"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/service"
	"github.com/rs/cors"
)

func runMigrations(db *sql.DB) {
	files, err := ioutil.ReadDir("migrations")
	if err != nil {
		log.Fatalf("failed to read migrations folder: %v", err)
	}

	for _, file := range files {
		path := "migrations/" + file.Name()

		query, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("failed to read migration file %s: %v", path, err)
		}

		_, err = db.Exec(string(query))
		if err != nil {
			log.Printf("migration skipped for %s: %v", path, err)
		} else {
			log.Printf("migration executed: %s", path)
		}
	}
}

func main() {
	// Inisialisasi DB bang, moga aja connect wkwk
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()
	runMigrations(db)

	// Ini Repositori nya ya ges
	userRepo := repository.NewUserRepositoryPostgres(db)
	bookRepo := repository.NewBookRepositoryPostgres(db)
	borrowRepo := repository.NewBorrowRepositoryPostgres(db)
	fineRepo := repository.NewFineRepositoryPostgres(db)

	// Ini Service buat logic-logic mumet
	userService := service.NewUserService(userRepo)
	bookService := service.NewBookService(bookRepo)
	borrowService := service.NewBorrowService(borrowRepo, bookRepo, fineRepo)

	// Handlers buat nagkep request dari depan
	userHandler := handler.NewUserHandler(userService)
	bookHandler := handler.NewBookHandler(bookService)
	borrowHandler := handler.NewBorrowHandler(borrowService)

	// Pake Mux biar gampang routingnya
	mux := http.NewServeMux()

	// Route yg public nih, siapa aja boleh masuk
	mux.HandleFunc("/auth/register", userHandler.Register)
	mux.HandleFunc("/auth/login", userHandler.Login)
	mux.HandleFunc("/books", bookHandler.Handle) // Handle public list sama admin CRUD sekalian, males misahin wkwk
	mux.HandleFunc("/books/", bookHandler.Handle)

	// Route yg di protect, harus login dlu bos
	mux.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(userHandler.Profile)))
	mux.Handle("/borrow", middleware.AuthMiddleware(http.HandlerFunc(borrowHandler.Borrow)))
	mux.Handle("/return", middleware.AuthMiddleware(http.HandlerFunc(borrowHandler.Return)))
	mux.Handle("/history", middleware.AuthMiddleware(http.HandlerFunc(borrowHandler.History)))
	mux.Handle("/fines", middleware.AuthMiddleware(http.HandlerFunc(borrowHandler.Fines)))

	// Route khusus admin doang
	mux.Handle("/admin/users", middleware.AuthMiddleware(middleware.AdminOnly(http.HandlerFunc(userHandler.List))))
	mux.Handle("/admin/borrows", middleware.AuthMiddleware(middleware.AdminOnly(http.HandlerFunc(borrowHandler.AllBorrows))))

	// Middleware CORS, biar ga kena block pas fetch
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Ntar kalo udah kelar diganti yak jgn bintang gini
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	// Wrapper buat logging, biar gampang debug kalo error
	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, finalHandler); err != nil {
		log.Fatal(err)
	}
}
