package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.POST("/signup",Signup)
	r.POST("/login",Login)

	r.Run(":2020")
}

func Signup(c *gin.Context) {

}
func Login(c *gin.Context)  {
	
}
