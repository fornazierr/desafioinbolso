package broker

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"transferenciaapi/apiutil"
	"transferenciaapi/models"
)

func getContaBancariaAPI(id int) (models.ContaBancaria, error) {
	config := apiutil.GetConfig()
	var conta models.ContaBancaria

	//realiza a busca pela conta
	url := fmt.Sprintf(config.CONTABANCARIA_API+"/contabancaria/%d", id)
	res, err := http.Get(url)
	if err != nil {
		log.Println("Erro ao identificar conta:", err.Error())
		return conta, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Erro lendo body: ", err.Error())
		return conta, err
	}

	if res.StatusCode != 200 {
		return conta, errors.New(string(body))
	} else {
		var rt models.ReturnMessage
		err := json.Unmarshal(body, &rt)
		if err != nil {
			log.Println("getContaBancariaAPI: erro durante parse da resposta", err.Error())
			return conta, err
		}

		//caso a minha mensagem possua um erro
		if rt.Status < 0 {
			log.Println("getContaBancariaAPI: resposta com erro ", err.Error())
			return conta, err
		}

		var aux []models.ContaBancaria
		err = json.Unmarshal([]byte(rt.Message), &aux)
		if err != nil {
			log.Println("getContaBancariaAPI: erro ao realizar parse para tipo ContaBancaria")
			return conta, err
		}

		if len(aux) < 1 {
			msgErr := fmt.Sprintf("getContaBancariaAPI: nenhuma conta encontrada para o ID fornecido [%d]\n", id)
			log.Printf(msgErr)
			return conta, errors.New(msgErr)
		}

		//adicionando conta à estrutura e verificando seus erros
		conta = aux[0]
		if err = conta.Erros(); err != nil {
			return conta, err
		}

		return conta, nil
	}
}
