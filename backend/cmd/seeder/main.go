package main

import (
	"log"
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

	// 2. Masukin buku-buku dummy yak wkwk
	books := []domain.Book{
		// Gaya-gaya buku jadul (Classics)
		{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Publisher: "Scribner", Year: 1925, Category: "Fiction", ISBN: "9780743273565", Stock: 5},
		{Title: "1984", Author: "George Orwell", Publisher: "Secker & Warburg", Year: 1949, Category: "Dystopian", ISBN: "9780451524935", Stock: 3},
		
		// Buku anak IT jaman now (Modern Tech)
		{Title: "Clean Architecture", Author: "Robert C. Martin", Publisher: "Prentice Hall", Year: 2017, Category: "Programming", ISBN: "9780134494166", Stock: 10},
		{Title: "Introduction to Algorithms", Author: "Thomas H. Cormen", Publisher: "MIT Press", Year: 2009, Category: "Computer Science", ISBN: "9780262033848", Stock: 2},
		
		// Buku-buku lokal kita raya (National Literature)
		{Title: "Laskar Pelangi", Author: "Andrea Hirata", Publisher: "Bentang Pustaka", Year: 2005, Category: "Novel", ISBN: "9789793062791", Stock: 7},
		{Title: "Bumi Manusia", Author: "Pramoedya Ananta Toer", Publisher: "Hasta Mitra", Year: 1980, Category: "Historical Fiction", ISBN: "9789799731234", Stock: 4},
	}

	for _, b := range books {
		err := bookRepo.Create(&b)
		if err != nil {
			log.Printf("Skipping book %s (likely already exists): %v", b.Title, err)
			continue
		}
		log.Printf("Seeded book: %s", b.Title)
	}
}
