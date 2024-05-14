package main

import (
	"fmt"
	"log"

	"go-store-api/model/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Koneksi ke database
	dsn := "postgres://postgres:password@localhost:5462/store-api"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Migrasi model data ke database untuk setiap tabel
	err = db.AutoMigrate(&domain.Product{}, &domain.Categories{}, &domain.ShoppingCarts{},
		&domain.Users{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migrasi selesai.")
}
