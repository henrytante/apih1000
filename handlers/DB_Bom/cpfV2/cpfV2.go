package cpfv2

import (
	database "apiHALL/db"
	dbstruct "apiHALL/handlers/DB_Basico"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CPF_V2(c *gin.Context)  {
	db, err := database.DB_ConnectV2()
	if err != nil{
		log.Fatal("Erro na conexão do v2")
	}
	defer db.Close()
	cpfV2 := c.Query("cpf")
	if cpfV2 == ""{
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Parametro 'cpf' vazio!"})
		return
	}
	var dados dbstruct.DadosCPF_V2
	err_Query := db.QueryRow("SELECT COALESCE(cpf, ''),COALESCE(pai, ''),COALESCE(mae, ''),COALESCE(municipioNascimento, ''),COALESCE(municipio, ''),COALESCE(logradouro, ''),COALESCE(numero, ''),COALESCE(bairro, ''),COALESCE(cep, ''),COALESCE(rgNumero, ''),COALESCE(rgOrgaoEmisor, ''),COALESCE(rgUf,''),COALESCE(rgDataEmissao, ''),COALESCE(cns, ''),COALESCE(telefone, ''),COALESCE(telefoneSecundario, '') FROM datasus WHERE cpf = ?",cpfV2).Scan(&dados.CPF,&dados.PAI,&dados.MAE,&dados.MUNICIPIONASC,&dados.MUNICIPIO,&dados.LOGRADOURO,&dados.NUMERO,&dados.BAIRRO,&dados.CEP,&dados.RGNUMERO,&dados.RGORGAOEMISSOR,&dados.RGUF,&dados.RGDATAEMISSAO,&dados.CNS,&dados.TELEFONE,&dados.TELEFONESEG)
	if err_Query != nil{
		if err_Query == sql.ErrNoRows{
			c.JSON(http.StatusNotFound, gin.H{"erro":"CPF não encontrado"})
			return
		}else{
			c.JSON(http.StatusInternalServerError, gin.H{"erro":err_Query.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, dados)
}