package main

import (
	"fmt"
	"log"
	"net/http"
	"transferenciabancariaapi/apiutil"
	"transferenciabancariaapi/models"
	"transferenciabancariaapi/router"
)

var Config models.Config

func main() {
	log.Println("Iniciando Transferencia Bancaria API")
	log.Println("Carregando configurações")
	Config = apiutil.GetConfig()
	log.Printf("Configurações: %+v\n", Config)
	r := router.StartRouter()
	url := fmt.Sprintf("%s:%s", Config.URL_HOST, Config.URL_PORT)
	log.Printf("Iniciando API em: %s\n", url)
	http.ListenAndServe(url, r.Router)
}
