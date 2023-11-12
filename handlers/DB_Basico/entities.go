package dbstruct


type DadosCPF struct {
	CPF        string
	NOME       string
	SEXO       string
	NASCIMENTO string
}
type DadosCPF_V2 struct {
	CPF            string
	PAI            string
	MAE            string
	MUNICIPIONASC  string
	MUNICIPIO      string
	LOGRADOURO     string
	NUMERO         string
	BAIRRO         string
	CEP            string
	RGNUMERO       string
	RGORGAOEMISSOR string
	RGUF           string
	RGDATAEMISSAO  string
	CNS            string
	TELEFONE       string
	TELEFONESEG    string
}

type Veiculo struct {
	ID                   string
	DataAtualizacao      string
	Chassi               string
	Placa                string
	Faturado             string
	AnoFabricacao        string
	Municipio            string
	UFPlaca              string
	MarcaModelo          string
	Combustivel          string
	Potencia             string
	CapacidadeCarga      string
	Nacionalidade        string
	Linha                string
	Carroceria           string
	CaixaCambio          string
	EixoTraseiroDif      string
	TerceiroEixo         string
	Motor                string
	TipoDocFaturado      string
	UFFaturado           string
	TipoDocProp          string
	AnoModelo            string
	TipoVeiculo          string
	EspecieVeiculo       string
	TipoCarroceria       string
	CorVeiculo           string
	QuantidadePassageiro string
	SituacaoChassi       string
	Eixos                string
	TipoMontagem         string
	TipoDocImportadora   string
	IdentImportadora     string
	DI                   string
	RegistroDI           string
	UnidadeLocalSRF      string
	UltimaAtualizacao    string
	Restricao1           string
	Restricao2           string
	Restricao3           string
	Restricao4           string
	LimiteRestricaoTrib  string
	Cilindradas          string
	CapMaximaTracao      string
	PesoBrutoTotal       string
	SituacaoVeiculo      string
	PlacaModeloAntigo    string
	PlacaModeloNovo      string
	PlacaNova            string
}
