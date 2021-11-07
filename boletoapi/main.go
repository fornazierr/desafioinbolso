package main

import (
	"boletoapi/apiutil"
	"boletoapi/models"
	"boletoapi/router"
	"fmt"
	"log"
	"net/http"
)

var Config models.Config

func main() {
	log.Println("Iniciando Modulo de Gerenciamento de Contas Bancarias")
	log.Println("Carregando configurações")
	Config = apiutil.GetConfig()
	log.Printf("Configurações: %+v\n", Config)
	log.Println("Carregando router")
	r := router.StartRouter()
	log.Println("Iniciando servidor em:", fmt.Sprintf("%s:%s", Config.URL_HOST, Config.URL_PORT))
	http.ListenAndServe(fmt.Sprintf("%s:%s", Config.URL_HOST, Config.URL_PORT), r.Router)
}
