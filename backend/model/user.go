package model

type Users []User

type User struct {
	ID    uint
	Name  string
	Email *string
	Age   uint8
}
