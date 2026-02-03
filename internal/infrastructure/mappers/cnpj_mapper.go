package mappers

import "strings"

type Cnpj struct {
	Uf                          string          `json:"uf"`
	ZipCode                     string          `json:"cep"`
	Shareholders                []Shareholder   `json:"qsa"`
	Cnpj                        string          `json:"cnpj"`
	Country                     *string         `json:"pais"`
	CountryCode                 *string         `json:"codigo_pais"`
	Email                       *string         `json:"email"`
	Size                        string          `json:"porte"`
	SizeCode                    int             `json:"codigo_porte"`
	Phone                       string          `json:"ddd_telefone_1"`
	Phone2                      string          `json:"ddd_telefone_2"`
	DddFax                      string          `json:"ddd_fax"`
	SocialReason                string          `json:"razao_social"`
	FantasyName                 string          `json:"nome_fantasia"`
	CapitalSocial               int             `json:"capital_social"`
	Neighborhood                string          `json:"bairro"`
	Number                      string          `json:"numero"`
	Complement                  string          `json:"complemento"`
	City                        string          `json:"municipio"`
	Street                      string          `json:"logradouro"`
	CnaeFiscalCode              int             `json:"cnae_fiscal"`
	CnaeFiscalDescription       string          `json:"cnae_fiscal_descricao"`
	MeiOption                   *bool           `json:"opcao_pelo_mei"`
	DescriptionSize             string          `json:"descricao_porte"`
	CityCode                    int             `json:"codigo_municipio"`
	CityCodeIbge                int             `json:"codigo_municipio_ibge"`
	SecondaryCnaes              []SecondaryCnae `json:"cnaes_secundarios"`
	LegalNature                 string          `json:"natureza_juridica"`
	LegalNatureCode             int             `json:"codigo_natureza_juridica"`
	TaxRegime                   []TaxRegime     `json:"regime_tributario"`
	SpecialSituation            string          `json:"situacao_especial"`
	SimpleOption                *bool           `json:"opcao_pelo_simples"`
	SituationCode               int             `json:"situacao_cadastral"`
	ReasonSituationCode         int             `json:"motivo_situacao_cadastral"`
	MatrixBranchIdentifier      int             `json:"identificador_matriz_filial"`
	MeiOptionDate               *string         `json:"data_opcao_pelo_mei"`
	MeiExclusionDate            *string         `json:"data_exclusao_do_mei"`
	ActivityStartDate           string          `json:"data_inicio_atividade"`
	SpecialSituationDate        *string         `json:"data_situacao_especial"`
	SimpleOptionDate            *string         `json:"data_opcao_pelo_simples"`
	SituationDate               string          `json:"data_situacao_cadastral"`
	CityNameInForeign           string          `json:"nome_cidade_no_exterior"`
	SimpleExclusionDate         *string         `json:"data_exclusao_do_simples"`
	ResponsibleCode             int             `json:"qualificacao_do_responsavel"`
	FederativeEntityResponsible string          `json:"ente_federativo_responsavel"`
	SituationDescription        string          `json:"descricao_situacao_cadastral"`
	StreetTypeDescription       string          `json:"descricao_tipo_de_logradouro"`
	ReasonSituationDescription  string          `json:"descricao_motivo_situacao_cadastral"`
	MatrixBranchDescription     string          `json:"descricao_identificador_matriz_filial"`
}

type Shareholder struct {
	Country                              *string `json:"pais"`
	Name                                 string  `json:"nome_socio"`
	Code                                 *string `json:"codigo_pais"`
	Age                                  string  `json:"faixa_etaria"`
	AgeCode                              int     `json:"codigo_faixa_etaria"`
	CnpjCpf                              string  `json:"cnpj_cpf_do_socio"`
	Qualification                        string  `json:"qualificacao_socio"`
	QualificationCode                    int     `json:"codigo_qualificacao_socio"`
	EntryDate                            string  `json:"data_entrada_sociedade"`
	ShareholderIdentifier                int     `json:"identificador_de_socio"`
	LegalRepresentativeCpf               string  `json:"cpf_representante_legal"`
	LegalRepresentativeName              string  `json:"nome_representante_legal"`
	LegalRepresentativeQualification     string  `json:"qualificacao_representante_legal"`
	LegalRepresentativeQualificationCode int     `json:"codigo_qualificacao_representante_legal"`
}

type TaxRegime struct {
	Year          int     `json:"ano"`
	CnpjScp       *string `json:"cnpj_da_scp"`
	TaxationForm  string  `json:"forma_de_tributacao"`
	NumberOfBooks int     `json:"quantidade_de_escrituracoes"`
}

type SecondaryCnae struct {
	Code        int    `json:"codigo"`
	Description string `json:"descricao"`
}

// CleanCnpj removes all non-numeric characters from the CNPJ.
func CleanCnpj(cnpj string) string {
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, " ", "")
	return cnpj
}
