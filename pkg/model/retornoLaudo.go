package model

type RetornoLaudo struct {
	Operacao               string `json:"operacao"`
	CodigoUnidade          string `json:"codigo_unidade"`
	Acnumber               string `json:"ac_number"`
	StudyId                string `json:"study_id"`
	CodigoMedicoExecutante string `json:"codigo_medico_executante"`
	NomeMedicoExecutante   string `json:"nome_medico_executante"`
	CrmMedicoExecutante    string `json:"crm_medico_executante"`
	CodigoMedicoRevisor    string `json:"codigo_medico_revisor"`
	DataLaudo              string `json:"data_laudo"`
	DataRevisado           string `json:"data_revisado"`
	CodigoPaciente         string `json:"codigo_paciente"`
	CodigoAtendimento      string `json:"codigo_atendimento"`
	CodigoPedido           string `json:"codigo_pedido"`
	UrlViewer              string `json:"url_viewer"`
	ResultadoPdf           string `json:"resultado_pdf"`
	ResultadoRft           string `json:"resultado_rft"`
}
