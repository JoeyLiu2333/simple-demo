package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)
	// undefined initRouter: configuration中选择package,填入github路径
	// 或 go run main.go router.go

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
