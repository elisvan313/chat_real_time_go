package main

import (
	"chat_real_time_go/src/api/v1/routes"
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

func main() {
	eng := gin.Default()

	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	url, err := pq.ParseURL(os.Getenv("DB_URL"))
	db, err := sql.Open("postgres", url)

	if err != nil {
		panic(err)
	}
	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	// Listen and Server in 0.0.0.0:8080
	if err := eng.Run(); err != nil {
		panic(err)
	}

}
