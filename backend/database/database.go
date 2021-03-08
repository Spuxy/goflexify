package database

import (
	"fmt"
	"log"

	"errors"
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

func (d *DbHandler) GetAllUsers(u *model.User) model.Users {

	return []model.User{{}, {}}
}

func (d *DbHandler) InsertUser(u model.User) error {
	var existsUser model.User

	d.Db.Raw("SELECT * FROM users WHERE id = ?", u.ID).Scan(&existsUser)
	if existsUser.ID > 0 {
		return errors.New("Database could not create the user")
	}

	rslt := d.Db.Create(&u)
	if rslt.Error != nil {
		return errors.New("Database could not create the user")
	}

	return nil
}

func (d *DbHandler) GetUserById(id int) model.User {
	var user model.User
	d.Db.First(&user, "id = ?", id)
	return user
}
func (d *DbHandler) GetUser(u model.User) model.User {
	var newUser model.User
	d.Db.First(newUser, u.ID)
	if newUser.ID < 1 {
		log.Fatal("wtf")
	}
	return newUser
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

	if err != nil {
		log.Fatal(err)
	}

	migrate(db)

	return db, err
}
func NewHandler(db *gorm.DB) *DbHandler {
	return &DbHandler{
		Db: db,
	}
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Something went wrong with migration")
	}
}

func (d *DbHandler) GetUserByEmailPassword(email, password string) model.User {
	var user model.User
	d.Db.Where("email = ? AND password = ?", email, password).Find(&user)
	return user
}
