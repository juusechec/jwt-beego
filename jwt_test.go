package jwtbeego_test

import (
	"testing"
	"time"

	"github.com/juusechec/jwt-beego"
)

func TestGetToken(t *testing.T) {
	et := jwtbeego.EasyToken{
		Username: "username",
		Expires:  time.Now().Unix() + 3600, //Segundos
	}
	tokenString, err := et.GetToken()
	if tokenString == "" {
		t.Errorf("Token String empty.")
	}
	if err != nil {
		t.Errorf("Error while verifying GetToken: %v", err)
	}
}
