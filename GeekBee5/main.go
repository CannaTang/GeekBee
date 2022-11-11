package main

import (
	"GeekBee5/api"
	_ "github.com/appleboy/gin-jwt/v2"
)

func main() {
	api.InitRouter()
}
