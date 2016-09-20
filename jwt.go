package jwtbeego

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
)

type EasyToken struct {
	Username string
	Expires  int64
}

// https://gist.github.com/cryptix/45c33ecf0ae54828e63b
// location of the files used for signing and verification
const (
	privKeyPath = "keys/rsakey.pem"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "keys/rsakey.pem.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey    *rsa.PublicKey
	mySigningKey *rsa.PrivateKey
)

func init() {
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
	}

	signBytes, err := ioutil.ReadFile(privKeyPath)

	if err != nil {
		log.Fatal(err)
	}

	mySigningKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err)
	}
}

func (e EasyToken) GetToken() (string, error) {

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: e.Expires, //time.Unix(c.ExpiresAt, 0)
		Issuer:    e.Username,
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, err
}

func (e EasyToken) ValidateToken(tokenString string) bool {
	// Token from another example.  This token is expired
	//var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if token.Valid {
		//fmt.Println("You look nice today")
		return true
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			//fmt.Println("That's not even a token")
			return false
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			//fmt.Println("Timing is everything")
			return false
		} else {
			//fmt.Println("Couldn't handle this token:", err)
			return false
		}
	} else {
		//fmt.Println("Couldn't handle this token:", err)
		return false
	}
}
