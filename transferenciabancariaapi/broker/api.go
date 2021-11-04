package broker

import (
	"fmt"
	"log"
	"net/http"
	"transferenciabancariaapi/apiutil"
	"transferenciabancariaapi/models"
)

func getContaBancaria(id int) (models.ContaBancaria, error) {
	config := apiutil.GetConfig()
	var conta models.ContaBancaria
	url := fmt.Sprintf(config.CONTABANCARIA_API+"/contabancaria/%d", id)
	res, err := http.Get(url)
	if err != nil {
		log.Println("Erro ao identificar conta:", err.Error())
		return conta, err
	}

}
