package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	register(r)
	err := r.Run()
	if err != nil {
		log.Println("run api service fail")
	}
}