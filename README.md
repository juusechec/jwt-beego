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

*) Adicionalmente se puede establecer que para todos los controladores se haga la validación excepto para el de login.

1) Configurar un nuevo paquete.

```go
//./myBeego/controller.go

//Se crea un espacio de nombres llamado myBeego
package myBeego

//Se agrega la biblioteca de beego
import (
	...
	"github.com/astaxie/beego"
)

//Se genera un tipo Controller que hereda de beego.Controller
type Controller struct {
	DisableJWT false
	beego.Controller
}

//Se redefine lo que hace la función Prepare
//* es un apuntador al igual que en C
//& hace referencia a la dirección de memoria
//La iniciación de una variable o funcion con * se traduce en que almacena
//u := 10 //var z *int  //z = &u //fmt.Println(z)//0x1040e0f8
//var s *string //var r **string = &s //fmt.Println(r)//0x1040a120
func (c *Controller) Prepare() {
	//Lo que quieras hacer en todos los controladores
	c.DisableJWT == false {
		tokenString := c.Ctx.Input.Query("tokenString")
	
		et := jwtbeego.EasyToken{}
		valido, _ := et.ValidateToken(tokenString)
		if !valido {
			c.Ctx.Output.SetStatus(401)
			c.Data["json"] = "Permission Deny"
			c.ServeJSON()
		}
	}
	return
}

```

2) Configurar llamado del nuevo Controller en todos los controladores:

```go
//./controllers/miObjeto.go

package controllers

import (
	...
	"github.com/juusechec/oas_be_cancelacion_semestre/myBeego"
	"github.com/juusechec/jwt-beego"
)

type MiObjetoController struct {
	//beego.Controller
	myBeego.Controller
	//myBeego.Controller.DisableJWT = DisableJWT //Si desea deshabilitar para este control
}

...
```

Se basa en:
* https://github.com/someone1/gcp-jwt-go
* https://github.com/dgrijalva/jwt-go
