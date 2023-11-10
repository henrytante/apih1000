package nomev1

import (
    database "apiHALL/db"
    dbbasico "apiHALL/handlers/DB_Basico"
    "log"
    "strings"

    "github.com/gin-gonic/gin"
)

func NomeV1(c *gin.Context) {
    db, err := database.DB_Connect()
    if err != nil {
        log.Fatal("Erro na conexão com o banco (nome)", err)
    }
    defer db.Close()
    
    nome := c.Query("nome")
    if nome == "" {
        c.JSON(400, gin.H{"error": "Parametro 'nome' vazio!"})
        return
    }
    
    nome = strings.ToLower(nome)

    // Consulta para buscar todas as pessoas com o mesmo nome
    rows, err := db.Query("SELECT cpf, nome, sexo, nascimento FROM pessoas WHERE nome = ?", nome)
    if err != nil {
        c.JSON(500, gin.H{"error na consulta": err})
        return
    }
    defer rows.Close()

    var resultados []dbbasico.DadosCPF
    for rows.Next() {
        var dados dbbasico.DadosCPF
        err := rows.Scan(&dados.CPF, &dados.NOME, &dados.SEXO, &dados.NASCIMENTO)
        if err != nil {
            log.Println("Erro ao escanear linha:", err)
            continue
        }
        resultados = append(resultados, dados)
    }

    if len(resultados) == 0 {
        c.JSON(404, gin.H{"error": "Nenhum nome encontrado!"})
    } else {
        c.JSON(200, resultados)
    }
}
