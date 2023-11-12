package cpfv1

import (
	database "apiHALL/db"
	dbstruct "apiHALL/handlers/DB_Basico"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func CPF_V1(c *gin.Context)  {
	db, err := database.DB_Connect()
	if err != nil{
		panic(err)
	}
	defer db.Close()
	cpf := c.Query("cpf")
	if cpf == ""{
		c.JSON(400, gin.H{"error":"Parametro cpf está vazio"})
		return
	}
	var dados dbstruct.DadosCPF
	err_Query := db.QueryRow("SELECT cpf,nome,sexo,nascimento FROM pessoas WHERE cpf = ?",cpf).Scan(&dados.CPF,&dados.NOME,&dados.SEXO,&dados.NASCIMENTO)
	if err_Query != nil{
		if err_Query == sql.ErrNoRows{
			c.JSON(404,gin.H{"error":"CPF não encontrado"})
			return
		}else{
			c.JSON(400,gin.H{"error":"Erro na consulta"})
			return
		}
	}
	c.JSON(200,dados)
}