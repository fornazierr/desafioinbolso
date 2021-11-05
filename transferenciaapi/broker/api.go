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
		var mess models.ReturnMessage
		err := json.Unmarshal(body, &mess)
		if err != nil {
			return conta, err
		}
		if mess.Status < 0 {
			return conta, errors.New(mess.Message)
		}
		err = json.Unmarshal([]byte(mess.Message), &conta)
		if err != nil {
			return conta, err
		}
		return conta, nil
	}
}
