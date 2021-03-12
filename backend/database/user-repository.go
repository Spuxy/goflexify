package database

import (
	"fmt"
	"log"

	"github.com/Spuxy/Goflexify/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositer {
	return &UserRepository{db}
}

func (d *UserRepository) CreateUser(u *model.User) error {
	var existsUser model.User

	err := d.DB.Raw("SELECT * FROM users WHERE id = ?", u.ID).Scan(&existsUser).Error
	if err != nil {
		return fmt.Errorf("we could not find the user with id %d. Reason is %s", u.ID, err)
	}

	err = d.DB.Create(&u).Error
	if err != nil {
		return fmt.Errorf("we could not create the user with id %d. Reason is %s", u.ID, err)
	}

	return nil
}

func (u *UserRepository) GetUserById(id int) (model.User, error) {
	var user model.User

	err := u.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		return model.User{}, fmt.Errorf("we could not find the user with id %d. Reason is %s", id, err)
	}

	return user, nil
}

func (d *UserRepository) GetUserByEmailPassword(email, password string) (model.User, error) {
	var user model.User

	log.Println(email, password)

	d.DB.Where("email = ? AND password = ?", email, password).Find(&user)
	if user.ID <= 1 {

		return model.User{}, fmt.Errorf("we could not find the user with email => %s", email)
	}

	return user, nil
}

func (d *UserRepository) UpdateUser(u *model.User) error {
	err := d.DB.Save(u).Error

	if err != nil {
		return fmt.Errorf("we could not update the user with email => %s. Reason is %s", u.Email, err)
	}

	return nil
}

func (d *UserRepository) DeleteUserById(id int) error {
	err := d.DB.Delete(&model.User{}, id).Error

	if err != nil {
		return fmt.Errorf("we could not delete the user with id => %d. Reason is %s", id, err)
	}

	return nil
}
