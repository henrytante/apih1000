package routes

import (
	"apiHALL/handlers/DB_Basico/cpfV1"
	"apiHALL/handlers/DB_Basico/nomeV1"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Configurar CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Permitir todas as origens (em produção, defina apenas as origens permitidas)
	config.AllowMethods = []string{"GET"}
	r.Use(cors.New(config))

	// Adicionar rotas
	r.GET("/api/cpf", cpfv1.CPF_V1)
	r.GET("/api/nome", nomev1.NomeV1)

	return r
}

func RunServer() {
	router := SetupRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro nas rotas", err)
	}
	fmt.Println("Servidor http://localhost:8080/")
}
