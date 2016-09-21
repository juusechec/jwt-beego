# jwt-beego
Una implementación simple de dgrijalva/jwt-go para beego. 

Pasos para implementar:

1) Crear una llave RSA con los comandos de ***generar_key.log***

2) Generar una ruta con action POST en el controlador especificado, en este caso la ruta elegida fue /user/getToken:

```go
// ./controllers/user.go
package controllers

import (
	...
	
	"github.com/juusechec/jwt-beego"
)

...

// @Title getToken
// @Description Get token from user session
// @Param	username		query 	string	true		"The username for get token"
// @Param	password		query 	string	true		"The password for get token"
// @Success 200 {string} Obtain Token
// @router /getToken [post]
func (u *UserController) GetToken() {
	username := u.Ctx.Input.Query("username")
	password := u.Ctx.Input.Query("password")

	tokenString := ""
	if username == "admin" && password == "mipassword" {
		et := jwtbeego.EasyToken{
			Username: username,
			Expires:  time.Now().Unix() + 3600, //Segundos
		}
		tokenString, _ = et.GetToken()
	}

	u.Data["json"] = "{'tokenString': '" + tokenString + "'}"
	u.ServeJSON()
	return
}

...
```

3) Agregar la validación del token en cada controlador que se necesite. Esto se hace a través de la función ***Prepare***.

```go
// ./controllers/tipo_cancelacion_semestre.go
package controllers

import (
	...
	
	"github.com/juusechec/jwt-beego"
)

func (c *TipoCancelacionSemestreController) Prepare() {
	tokenString := c.Ctx.Input.Query("tokenString")

	et := jwtbeego.EasyToken{}
	valido, _ := et.ValidateToken(tokenString)
	if !valido {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = "Permission Deny"
		c.ServeJSON()
	}
	return
}

...
```

Se basa en:
* https://github.com/someone1/gcp-jwt-go
* https://github.com/dgrijalva/jwt-go
