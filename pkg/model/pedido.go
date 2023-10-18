package model

type Pedido struct {
	Operacao                string `json:"Operacao"`
	CodigoUnidade           string `json:"CodigoUnidade"`
	CodigoAtendimento       string `json:"CodigoAtendimento"`
	DataAtendimento         string `json:"DataAtendimento"`
	CodigoMedicoSolicitante string `json:"codigoMedicoSolicitante"`
	NomeMedicoSolicitante   string `json:"NomeMedicoSolicitante"`
	CrmMedicoSolicitante    string `json:"CrmMedicoSolicitante"`
	CodigoConvenio          string `json:"CodigoConvenio"`
	NomeConvenio            string `json:"NomeConvenio"`
	CodigoSetorSolicitante  string `json:"CodigoSetorSolicitante"`
	NomeSetorSolicitante    string `json:"NomeSetorSolicitante"`
	CodigoSetorExecutante   string `json:"CodigoSetorExecutante"`
	NomeSetorExecutante     string `json:"NomeSetorExecutante"`
	CodigoPaciente          string `json:"CodigoPaciente"`
	NomePaciente            string `json:"NomePaciente"`
	DataNascimento          string `json:"DataNascimento"`
	Sexo                    string `json:"Sexo"`
	Endereco                string `json:"Endereco"`
	Bairro                  string `json:"Bairro"`
	Cidade                  string `json:"Cidade"`
	Uf                      string `json:"Uf"`
	Cep                     string `json:"Cep"`
	Telefone                string `json:"Telefone"`
	Cpf                     string `json:"Cpf"`
	Email                   string `json:"Email"`
	Rg                      string `json:"Rg"`
	Celular                 string `json:"Celular"`
	NomeSocial              string `json:"NomeSocial"`
	Cns                     string `json:"cns"`
	Mae                     string `json:"Mae"`
	CodigoPedido            string `json:"CodigoPedido"`
	DataPedido              string `json:"DataPedido"`
	Acnumber                string `json:"Acnumber"`
	CodigoExame             string `json:"CodigoExame"`
	NomeExame               string `json:"NomeExame"`
	Observacao              string `json:"Observacao"`
	Justificativa           string `json:"Justificativa"`
	AlturaPaciente          string `json:"AlturaPaciente"`
	PesoPaciente            string `json:"PesoPaciente"`
	PrescricaoUrgente       string `json:"PrescricaoUrgente"`
	Modalidade              string `json:"Modalidade"`
}
