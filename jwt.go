package jwt_beego

import (
	"github.com/dgrijalva/jwt-go"
)

type EasyToken struct{}

func init() {

}

func (e *EasyToken) GetToken() string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	//fmt.Println(tokenString, err)
	return tokenString
}

func (e *EasyToken) ValidateToken(tokenString string) bool {
	// Token from another example.  This token is expired
	//var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
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
