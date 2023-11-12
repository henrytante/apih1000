package veiculosv1

import (
	database "apiHALL/db"
	dbstruct "apiHALL/handlers/DB_Basico"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Veiculo(c *gin.Context)  {
	db, err := database.DB_ConnectV2()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro na conexão com o banco de dados V2"})
		return
	}
	defer db.Close()
	placa := c.Query("placa")
	if placa == ""{
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Parametro 'placa' vazio"})
		return
	}
	placa = strings.ToUpper(placa)
	fmt.Println(placa)
	var dados dbstruct.Veiculo
	err_Query := db.QueryRow("SELECT COALESCE(id, ''), COALESCE(data_atualiacao, ''), COALESCE(chassi, ''), COALESCE(placa, ''), COALESCE(faturado, ''), COALESCE(ano_fabricacao, ''), COALESCE(municipio, ''), COALESCE(uf_placa, ''), COALESCE(marca_modelo, ''), COALESCE(combustivel, ''), COALESCE(potencia, ''), COALESCE(capacidade_carga, ''), COALESCE(nacionalidade, ''), COALESCE(linha, ''), COALESCE(carroceria, ''), COALESCE(caixa_cambio, ''), COALESCE(eixo_traseiro_dif, ''), COALESCE(terceiro_eixo, ''), COALESCE(motor, ''), COALESCE(tipo_doc_faturado, ''), COALESCE(uf_faturado, ''), COALESCE(tipo_doc_prop, ''), COALESCE(ano_modelo, ''), COALESCE(tipo_veiculo, ''), COALESCE(especie_veiculo, ''), COALESCE(tipo_carroceria, ''), COALESCE(cor_veiculo, ''), COALESCE(quantidade_passageiro, ''), COALESCE(situacao_chassi, ''), COALESCE(eixos, ''), COALESCE(tipo_montagem, ''), COALESCE(tipo_doc_importadora, ''), COALESCE(ident_importadora, ''), COALESCE(di, ''), COALESCE(registro_di, ''), COALESCE(unidade_local_srf, ''), COALESCE(ultima_atualizacao, ''), COALESCE(restricao_1, ''), COALESCE(restricao_2, ''), COALESCE(restricao_3, ''), COALESCE(restricao_4, ''), COALESCE(limite_restricao_trib, ''), COALESCE(cilindradas, ''), COALESCE(cap_maxima_tracao, ''), COALESCE(peso_bruto_total, ''), COALESCE(situacao_veiculo, ''), COALESCE(placa_modelo_antigo, ''), COALESCE(placa_modelo_novo, ''), COALESCE(placa_nova, '') FROM vehicles WHERE placa = ?",placa).Scan(&dados.ID, &dados.DataAtualizacao, &dados.Chassi, &dados.Placa, &dados.Faturado, &dados.AnoFabricacao, &dados.Municipio, &dados.UFPlaca, &dados.MarcaModelo, &dados.Combustivel, &dados.Potencia, &dados.CapacidadeCarga, &dados.Nacionalidade, &dados.Linha, &dados.Carroceria, &dados.CaixaCambio, &dados.EixoTraseiroDif, &dados.TerceiroEixo, &dados.Motor, &dados.TipoDocFaturado, &dados.UFFaturado, &dados.TipoDocProp, &dados.AnoModelo, &dados.TipoVeiculo, &dados.EspecieVeiculo, &dados.TipoCarroceria, &dados.CorVeiculo, &dados.QuantidadePassageiro, &dados.SituacaoChassi, &dados.Eixos, &dados.TipoMontagem, &dados.TipoDocImportadora, &dados.IdentImportadora, &dados.DI, &dados.RegistroDI, &dados.UnidadeLocalSRF, &dados.UltimaAtualizacao, &dados.Restricao1, &dados.Restricao2, &dados.Restricao3, &dados.Restricao4, &dados.LimiteRestricaoTrib, &dados.Cilindradas, &dados.CapMaximaTracao, &dados.PesoBrutoTotal, &dados.SituacaoVeiculo, &dados.PlacaModeloAntigo, &dados.PlacaModeloNovo, &dados.PlacaNova)

	if err_Query != nil{
		if err_Query == sql.ErrNoRows{
			c.JSON(http.StatusNotFound, gin.H{"erro": "Placa não encontrada"})
			return
		}else{
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err_Query.Error()})
			return
		}
	}
	c.JSON(200, dados)
}