package handler

import (
	"application/pkg/database"
	"application/pkg/model"
	"application/pkg/repository"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var status int

func RecebePedidoVivace(w http.ResponseWriter, r *http.Request) {

	bytes, erro := io.ReadAll(r.Body)

	if erro != nil {

		erro = fmt.Errorf("erro ao tentar ler o corpo da requisição: %v", erro)

		log.Default().Print(erro)

		status = http.StatusBadRequest

		return
	}

	var pedidoVicace model.Pedido

	if erro = json.Unmarshal(bytes, &pedidoVicace); erro != nil {

		erro = fmt.Errorf("erro ao tentar realizar o unmarshal dos bytes na estrutura do pedido: %v", erro)

		log.Default().Print(erro)

		status = http.StatusBadRequest

		return
	}

	connection, erro := database.GetConnection()

	if erro != nil {

		erro = fmt.Errorf("erro ao tentar abrir a conexão: %v", erro)

		log.Default().Print(erro)

		status = http.StatusInternalServerError

		return

	}

	defer func(connection *sql.DB) {

		if erro = connection.Close(); erro != nil {
			erro = fmt.Errorf("erro ao tentar fechar a conexão com o banco de dados: %v", erro)

			log.Default().Print(erro)

			status = http.StatusInternalServerError

			return
		}

	}(connection)

	rep := repository.NewRepository(connection)

	switch pedidoVicace.Operacao {

	case "I":

		if erro = rep.RecebePedidoVivace(pedidoVicace); erro != nil {

			erro = fmt.Errorf("erro ao tentar receber o pedido: %v", erro)

			log.Default().Print(erro)

			status = http.StatusBadRequest

			return

		} else {

			status = http.StatusCreated

		}

	case "E":

		/*util := util.NewUtil()

		if erro = util.ValidaCampoDelete(pedidoVicace); erro != nil {
			status = http.StatusBadRequest
			log.Print(erro)
			w.WriteHeader(status)
			return
		}
		*/

		if erro = rep.DeletaExame(pedidoVicace); erro != nil {

			erro = fmt.Errorf("erro ao tentar deletar o exame: %v", erro)

			status = http.StatusBadRequest

			return

		} else {

			status = http.StatusNoContent

		}

		if erro = rep.DeletaAtendimento(pedidoVicace); erro != nil {

			erro = fmt.Errorf("erro ao tentar deletar o atendimento: %v", erro)

			log.Default().Print(erro)

			status = http.StatusBadRequest

			return

		} else {

			status = http.StatusNoContent

		}

	default:

		status = http.StatusUnauthorized

	}

	w.WriteHeader(status)
}
