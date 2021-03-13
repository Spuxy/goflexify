package database

import (
	"fmt"
	"log"

	"github.com/Spuxy/Goflexify/model"
	"github.com/Spuxy/Goflexify/utils/reader"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(r reader.Reader) (*gorm.DB, error) {
	cfg, err := r.ReadGivenFileIntoMap()
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbDatabase, cfg.DbPort)), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	migrate(db)

	return db, err
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Something went wrong with migration")
	}
}
