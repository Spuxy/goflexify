package database_test

import (
	"testing"

	"github.com/Spuxy/Goflexify/database"
	"github.com/Spuxy/Goflexify/model"
)

func TestInsert(t *testing.T) {
	user := model.User{
		ID:    2,
		Name:  "Spuxy",
		Email: "Spuxy@seznam.cz",
		Age:   "2",
	}

	conn, err := database.Connect()
	if err != nil {
		t.Error(err, "not pass")
	}
	db := database.NewHandler(conn)
	err = db.InsertUser(&user)

	if err != nil {
		t.Error("not pass bro")
	}

}
