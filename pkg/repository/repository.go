package repository

import (
	"application/pkg/model"
	"database/sql"
	"fmt"
	"log"
)

type InterfaceVivace interface {
	RecebePedidoVivace(pedido model.Pedido) error
	DeletaPedidoVivace(pedido model.Pedido) error
}

type Repository struct {
	DB *sql.DB
}

func NewRepository(connection *sql.DB) *Repository {
	return &Repository{DB: connection}
}

func (rep Repository) RecebePedidoVivace(pedido model.Pedido) (erro error) {
	begin, erro := rep.DB.Begin()

	if erro != nil {

		erro = fmt.Errorf("erro ao tentar inicializar a transação com banco de dados: %v", erro)

		log.Default().Print(erro)

		return

	}

	defer func() {
		if erro != nil {
			if erro = begin.Rollback(); erro != nil {

				erro = fmt.Errorf("erro ao tentar realizar o rollback: %v", erro)

				log.Default().Print(erro)

				return

			} else {
				if erro = begin.Commit(); erro != nil {

					erro = fmt.Errorf("erro ao tentar realizar o commit: %v", erro)

					log.Default().Print(erro)

					return

				}
			}

		}
	}()

	var cdPedido int
	var operacaoPedido string

	if erro = begin.QueryRow(`INSERT INTO mv_integra_clinux (
							                                 ds_operacao,
															 cd_paciente_vivace,
															 ds_paciente,
															 ds_cpf,
															 dt_nascimento,
															 ds_logradouro,
															 ds_cidade,
															 ds_estado,
															 nr_altura,
															 nr_peso,
							                                 cd_medico_vivace, 
							                                 ds_medico,
															 nr_crm,
							                                 cd_pedido, 
							                                 dt_pedido,
															 cd_exame_vivace
															) VALUES ($1, $2, 
																	  $3, $4, 
																	  $5, $6, 
																	  $7, $8
																	  $9, $10
																	  $11, $12
																	  $13, $14
																	  $15, $16
																	  ) RETURNING cd_pedido`,
	).Scan(&cdPedido); erro != nil {
		erro = fmt.Errorf("erro ao tentar escanear o código do pedido inserido: %v", erro)
		log.Default().Print(erro)
		return
	}

	if cdPedido >= 1 {
		log.Default().Printf("pedido inserido com sucesso: código = %v", cdPedido)
	}

	if erro = begin.QueryRow(`
				    SELECT ds_operacao 
				    FROM mv_integra_clinux 
					WHERE cd_pedido = $1`,
		cdPedido).Scan(&operacaoPedido); erro != nil {
		erro = fmt.Errorf("erro ao tentar escanear o ")
	}

	return

}

func (rep Repository) DeletaExame(pedido model.Pedido) (erro error) {

	begin, erro := rep.DB.Begin()

	if erro != nil {

		erro = fmt.Errorf("erro ao tentar começar uma transação com o banco de dados: %v", erro)

		log.Default().Print(erro)

		return
	}

	defer func() {

		if erro != nil {
			if erro = begin.Rollback(); erro != nil {

				erro = fmt.Errorf("erro ao tentar realizar o rollback: %v", erro)

				log.Default().Print(erro)

				return
			}
		} else {
			if erro = begin.Commit(); erro != nil {

				erro = fmt.Errorf("erro ao tentar realizar o commit: %v", erro)

				log.Default().Print(erro)

				return
			}
		}

	}()

	result, erro := begin.Exec(`DELETE FROM exames WHERE cd_exame_vivace = $1`, pedido.CodigoExame)

	if erro != nil {
		log.Printf("erro ao tentar deletar o exame. Erro: %v", erro)
	}

	affect, erro := result.RowsAffected()

	if erro != nil {
		log.Printf("erro ao tentar pegar as linhas afetadas. Erro: %v", erro)
	}

	if affect >= 1 {
		log.Printf("exame código = %v deletado", pedido.CodigoExame)
	} else {
		log.Printf("exame não encontrado")
	}

	return

}

func (rep Repository) DeletaAtendimento(pedido model.Pedido) (erro error) {
	begin, erro := rep.DB.Begin()

	if erro != nil {
		log.Printf("erro ao tentar inicializar a transação com banco de dados. Erro: %v", erro)
		return
	}

	defer func() {
		if erro != nil {
			if erro = begin.Rollback(); erro != nil {
				log.Printf("Erro ao tentar realizar o rollback. Erro: %v", erro)
			}
		} else {
			if erro = begin.Commit(); erro != nil {
				log.Printf("Erro ao tentar realizar o commit. Erro: %v", erro)
			}
		}
	}()

	result, erro := begin.Exec(`DELETE FROM atendimentos WHERE cd_atendimento_vivace = $1`, pedido.CodigoAtendimento)

	if erro != nil {
		log.Printf("erro ao tentar deletar o atendimento. Erro: %v", erro)
	}

	affect, erro := result.RowsAffected()

	if erro != nil {
		log.Printf("erro ao tentar pegar as linhas afetadas. Erro: %v", erro)
	}

	if affect >= 1 {
		log.Printf("atendimento código = %v deletado", pedido.CodigoAtendimento)
	} else {
		log.Printf("atendimento não encontrado")
	}

	return
}
