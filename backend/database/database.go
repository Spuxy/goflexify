package database

import (
	"fmt"
	"log"

	"github.com/Spuxy/Goflexify/model"

	"github.com/Spuxy/Goflexify/utils/reader"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDb interface {
	Connect(cfg reader.Config) (*gorm.DB, error)
}

type DbHandler struct {
	Db *gorm.DB
}

func (d *DbHandler) SelectUser(u *model.User) model.Users {
	return []model.User{{}, {}}
}
func (d *DbHandler) InsertUser(u *model.User) error {
	var existsUser model.User
	d.Db.First(&existsUser, u.ID)
	if existsUser.ID > 0 {
		log.Fatalln("gg")
	}
	r := d.Db.Create(u)
	return r.Error
}
func (d *DbHandler) UpdateUser(u *model.User) error {
	return nil
}
func (d *DbHandler) DeleteUser(u *model.User) error {
	return nil
}

func Connect() (*gorm.DB, error) {
	cfg, err := reader.ReadGivenFileIntoMap()
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbDatabase, cfg.DbPort)), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
func NewHandler(db *gorm.DB) *DbHandler {
	return &DbHandler{
		Db: db,
	}
}
