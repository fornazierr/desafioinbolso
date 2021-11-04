package broker

import (
	"bytes"
	"contaapi/apiutil"
	"contaapi/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var Config models.Config = apiutil.GetConfig()

func sendCriarNovoSaldo(cb models.ContaBancaria) {
	log.Printf("Criando novo Saldo para: %+v\n", cb)
	url := fmt.Sprintf("%s/newsaldo", Config.TRANFERENCIA_API)

	req := models.RequestSaldo{
		TitularId: cb.TitularId,
		ContaId:   cb.ID,
	}

	reqByte, err := json.Marshal(&req)
	if err != nil {
		log.Printf("Não foi possivel criar saldo para %d - %d: %s", cb.TitularId, cb.ID, err.Error())
		return
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqByte))
	if err != nil {
		log.Printf("Não foi possivel criar saldo para %d - %d: %s", cb.TitularId, cb.ID, err.Error())
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		print(err)
	}
	if res.StatusCode == 200 {
		log.Printf("Retorno: %s", string(body))
	} else {
		log.Printf("Falha: %s", string(body))
	}
}
