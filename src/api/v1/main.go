package main

import (
	"chat_real_time_go/src/api/v1/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	eng := gin.Default()
	router := routes.NewRouter(eng)
	router.MapRoutes()

	// Listen and Server in 0.0.0.0:8080
	if err := eng.Run(); err != nil {
		panic(err)
	}

}
