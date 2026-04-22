package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/config"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/domain"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/internal/repository"
	"github.com/Z-ZxZ/sistem-pinjam-perpus/backend/pkg/utils"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepositoryPostgres(db)
	bookRepo := repository.NewBookRepositoryPostgres(db)

	// 1. Bikin akun Admin dlu buat testing
	adminEmail := "admin@perpus.com"
	existingAdmin, _ := userRepo.GetByEmail(adminEmail)
	if existingAdmin == nil {
		hashedPassword, _ := utils.HashPassword("admin123")
		admin := &domain.User{
			Name:     "Librarian Admin",
			Email:    adminEmail,
			Password: hashedPassword,
			Role:     domain.RoleAdmin,
			Status:   "active",
		}
		userRepo.Create(admin)
		log.Println("Admin user seeded")
	}

	// 2. Baca dataset dari JSON cuy
	content, err := os.ReadFile("data/books.json")
	if err != nil {
		log.Fatalf("Gagal baca file JSON: %v", err)
	}

	var books []domain.Book
	err = json.Unmarshal(content, &books)
	if err != nil {
		log.Fatalf("Gagal unmarshal JSON: %v", err)
	}

	log.Printf("Mulai seeding %d buku...", len(books))
	for _, b := range books {
		err := bookRepo.Create(&b)
		if err != nil {
			log.Printf("Skipping book %s (likely already exists): %v", b.Title, err)
			continue
		}
		log.Printf("Seeded book: %s", b.Title)
	}
}
