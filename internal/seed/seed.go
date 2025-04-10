package seed

import (
	"log"

	"github.com/bayskie/test-book-api/internal/model"
	"gorm.io/gorm"
)

func SeedBooks(db *gorm.DB) {
	var count int64
	db.Model(&model.Book{}).Count(&count)
	if count > 0 {
		log.Println("[seeder] books table already seeded")
		return
	}

	books := []model.Book{
		{Title: "Laut Bercerita", Author: "Leila S. Chudori", Year: 2017},
		{Title: "Pulang", Author: "Leila S. Chudori", Year: 2012},
		{Title: "Bumi Manusia", Author: "Pramoedya Ananta Toer", Year: 1980},
		{Title: "Anak Semua Bangsa", Author: "Pramoedya Ananta Toer", Year: 1981},
		{Title: "Jejak Langkah", Author: "Pramoedya Ananta Toer", Year: 1985},
		{Title: "Rumah Kaca", Author: "Pramoedya Ananta Toer", Year: 1988},
		{Title: "Animal Farm", Author: "George Orwell", Year: 1945},
		{Title: "1984", Author: "George Orwell", Year: 1949},
		{Title: "Norwegian Wood", Author: "Haruki Murakami", Year: 1987},
		{Title: "Kafka on the Shore", Author: "Haruki Murakami", Year: 2002},
		{Title: "The Catcher in the Rye", Author: "J.D. Salinger", Year: 1951},
		{Title: "To Kill a Mockingbird", Author: "Harper Lee", Year: 1960},
		{Title: "Pride and Prejudice", Author: "Jane Austen", Year: 1813},
		{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: 1925},
		{Title: "The Hobbit", Author: "J.R.R. Tolkien", Year: 1937},
		{Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Year: 1954},
		{Title: "Harry Potter and the Philosopher's Stone", Author: "J.K. Rowling", Year: 1997},
		{Title: "The Da Vinci Code", Author: "Dan Brown", Year: 2003},
		{Title: "The Alchemist", Author: "Paulo Coelho", Year: 1988},
		{Title: "The Book Thief", Author: "Markus Zusak", Year: 2005},
	}

	if err := db.Create(&books).Error; err != nil {
		log.Printf("[seeder] failed to seed books: %v", err)
		return
	}

	log.Println("[seeder] books table seeded successfully")
}
