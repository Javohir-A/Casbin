package main

import (
	"file-management/storage"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func main() {
	db := storage.ConnectRedisDB()

	router := gin.Default()
	fmt.
		router.GET()
}
