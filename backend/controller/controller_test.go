package controller

import (
	"errors"
	"testing"

	"github.com/Spuxy/Goflexify/model"
	"github.com/gofiber/fiber/v2"
)

type MockORMDatabase struct{}

func (m MockORMDatabase) CreateUser(u model.User) error {
	return errors.New("w")
}
func (m *MockORMDatabase) GetUserByEmailPassword(email, password string) *model.User {
	user := new(model.User)
	return user
}

func TestRegisteringAnUser(t *testing.T) {
	m := MockORMDatabase{}
	c := NewController(m)
	c.Register(fiber.App)
}
