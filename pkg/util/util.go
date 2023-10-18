package util

import (
	"application/pkg/model"
)

type UtilInterface interface {
	ValidaCamposInsere(pedido model.Pedido) error
	ValidaCampoDelete(pedido model.Pedido) error
}

type Util struct {
}

func NewUtil() *Util {
	return &Util{}
}

/*
func (u Util) ValidaCamposInsere(pedido model.Pedido) error {
	if pedido.Operacao == "" {
		return errors.New("pedido não pode ser nulo")
	}
	if pedido.CodigoMedicoSolicitante == "" {
		return errors.New("código do médico não pode ser nulo")
	}
	if pedido.CodigoAtendimento == "" {
		return errors.New("código atendimento não pode ser nulo")
	}
	if pedido.NomeMedicoSolicitante == "" {
		return errors.New("nome do médico não pode ser nulo")
	}
	if pedido.CrmMedicoSolicitante == "" {
		return errors.New("crm do médico não pode ser nulo")
	}
	if pedido.CodigoPaciente == "" {
		return errors.New("código do paciente não pode ser nulo")
	}
	if pedido.NomePaciente == "" {
		return errors.New("nome do paciente não pode ser nulo")
	}
	if pedido.Cpf == "" {
		return errors.New("cpf não pode ser nulo")
	}
	if pedido.CodigoExame == "" {
		return errors.New("código do exame não pode ser nulo")
	}
	if pedido.Acnumber == "" {
		return errors.New("ac number não pode ser nulo")
	}
	return nil
}

func (u Util) ValidaCampoDelete(pedido model.Pedido) (erro error) {

	if pedido.CodigoAtendimento == "" {
		return errors.New("código do atendimento não pode ser nulo")
	}

	if pedido.CodigoExame == "" {
		return errors.New("código do exame não pode ser nulo")
	}

	return nil
}
*/
