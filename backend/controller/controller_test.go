package controller

import (
	"net/http/httptest"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	controller := NewController()

	user := []byte(`{"name":FilipT,"email":"test@gg","age":"22","password":"kinesis"}`)

	handler := httptest.NewRequest("POST", "/api/v1/user", nil)
}
