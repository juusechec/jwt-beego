# jwt-beego
Una implementación simple de dgrijalva/jwt-go para beego. 

```go
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
```

Se basa en:
* https://github.com/someone1/gcp-jwt-go
* https://github.com/dgrijalva/jwt-go
