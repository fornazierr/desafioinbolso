package router

import (
	"contaapi/broker"
	"contaapi/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CredencialRouter struct {
	Router *mux.Router
}

func StartRouter() *CredencialRouter {
	rt := &CredencialRouter{
		Router: mux.NewRouter(),
	}

	rt.initRoutes()

	return rt
}

func (cr *CredencialRouter) initRoutes() {
	//titular
	cr.Router.HandleFunc("/titular", getTitular).Methods("GET")
	cr.Router.HandleFunc("/titular/{id}", getTitularById).Methods("GET")
	cr.Router.HandleFunc("/titular", postTitular).Methods("POST")
	cr.Router.HandleFunc("/titular", deleteTitular).Methods("DELETE")
	cr.Router.HandleFunc("/titular/{id}", deleteTitularById).Methods("DELETE")

	//conta bancaria
	cr.Router.HandleFunc("/contabancaria", getConta).Methods("GET")
	cr.Router.HandleFunc("/contabancaria/{id}", getContaById).Methods("GET")
	cr.Router.HandleFunc("/contabancaria", postConta).Methods("POST")
	cr.Router.HandleFunc("/contabancaria", deleteConta).Methods("DELETE")
	cr.Router.HandleFunc("/contabancaria/{id}", deleteContaById).Methods("DELETE")

	broker.InitBroker()
}

func sendError(err error, w http.ResponseWriter) {
	res := models.ReturnMessage{
		Status:  -1,
		Message: fmt.Sprintf("Erro : %s", err.Error()),
	}
	resByte, _ := json.Marshal(res)
	fmt.Fprintln(w, string(resByte))
}

func sendOk(val interface{}, w http.ResponseWriter) {
	valJson, _ := json.Marshal(val)
	res := models.ReturnMessage{
		Status:  0,
		Message: string(valJson),
	}
	resByte, _ := json.Marshal(res)
	fmt.Fprintln(w, string(resByte))
}

func getTitular(w http.ResponseWriter, req *http.Request) {
	log.Println("Nova requisição em \"getTitular\".")
	arr, err := broker.GetTitular()
	if err != nil {
		sendError(err, w)
	} else {
		sendOk(arr, w)
	}
}

func getTitularById(w http.ResponseWriter, req *http.Request) {
	log.Println("Nova requisição em \"getTitularById\".")
	vars := mux.Vars(req)
	id := vars["id"]
	if id == "" {
		sendError(errors.New("ID do titular não informado"), w)
	} else {
		val, _ := strconv.Atoi(id)
		log.Println("Valor do id:", val)
		res, err := broker.GetTitularById(val)
		if err != nil {
			sendError(err, w)
		} else {
			sendOk(res, w)
		}
	}
}

func postTitular(w http.ResponseWriter, req *http.Request) {
	log.Println("Nova requisição em \"postTitular\".")
	var titular models.Titular
	json.NewDecoder(req.Body).Decode(&titular)
	log.Printf("Objeto da requisição: %+v\n", titular)
	err := titular.Erros()
	if err != nil {
		sendError(err, w)
	} else {
		err = broker.PostTitular(titular)
		if err != nil {
			sendError(err, w)
		} else {
			sendOk("Gravado/Atualizado com sucesso", w)
		}
	}
}

func deleteTitular(w http.ResponseWriter, req *http.Request) {
	log.Println("Nova requisição em \"deleteTitular\".")
	var titular models.Titular
	json.NewDecoder(req.Body).Decode(&titular)
	err := titular.ErrosDelete()
	if err != nil {
		sendError(err, w)
	} else {
		err := broker.DeleteTitular(titular)
		if err != nil {
			sendError(err, w)
		} else {
			sendOk("Sucesso ao excluir titular", w)
		}
	}
}

func deleteTitularById(w http.ResponseWriter, req *http.Request) {
	log.Println("Nova requisição em \"deleteTitularById\".")
	vars := mux.Vars(req)
	id := vars["id"]
	if id == "" {
		sendError(errors.New("ID do titular não informado"), w)
	} else {
		val, _ := strconv.Atoi(id)
		log.Println("Valor do id:", val)
		titular := models.Titular{
			ID: val,
		}
		err := broker.DeleteTitular(titular)
		if err != nil {
			sendError(err, w)
		} else {
			sendOk("Sucesso ao excluir titular", w)
		}
	}
}

func getConta(w http.ResponseWriter, req *http.Request) {
	log.Println("Nova requisição em \"getConta\".")
	arr, err := broker.GetContaBancaria()
	if err != nil {
		sendError(err, w)
	} else {
		sendOk(arr, w)
	}
}

func getContaById(w http.ResponseWriter, req *http.Request) {
	log.Println("Nova requisição em \"getContaById\".")

	vars := mux.Vars(req)
	id := vars["id"]
	if id == "" {
		sendError(errors.New("ID da conta não informado"), w)
	} else {
		val, _ := strconv.Atoi(id)
		arr, err := broker.GetContaBancariaById(val)
		if err != nil {
			sendError(err, w)
		} else {
			sendOk(arr, w)
		}
	}
}

func postConta(w http.ResponseWriter, req *http.Request) {
	log.Println("Nova requisição em \"postConta\".")
	var cb models.ContaBancaria
	json.NewDecoder(req.Body).Decode(&cb)
	err := cb.Erros()
	if err != nil {
		sendError(err, w)
	} else {
		log.Printf("Objeto da requisição: %+v\n", cb)
		err := broker.SaveContaBancaria(cb)
		if err != nil {
			sendError(err, w)
		} else {
			sendOk("Sucesso ao salvar conta bancaria", w)
		}
	}
}

func deleteConta(w http.ResponseWriter, req *http.Request) {
	log.Println("Nova requisição em \"deleteTitular\".")
	var cb models.ContaBancaria
	json.NewDecoder(req.Body).Decode(&cb)
	err := cb.ErrosDelete()
	if err != nil {
		sendError(err, w)
	} else {
		err := broker.DeleteContaBancaria(cb)
		if err != nil {
			sendError(err, w)
		} else {
			sendOk("Sucesso ao excluir conta bancaria", w)
		}
	}
}

func deleteContaById(w http.ResponseWriter, req *http.Request) {
	log.Println("Nova requisição em \"deleteContaById\".")
	vars := mux.Vars(req)
	id := vars["id"]
	if id == "" {
		sendError(errors.New("ID do titular não informado"), w)
	} else {
		val, _ := strconv.Atoi(id)
		log.Println("Valor do id:", val)
		conta := models.ContaBancaria{
			ID: val,
		}
		err := broker.DeleteContaBancaria(conta)
		if err != nil {
			sendError(err, w)
		} else {
			sendOk("Sucesso ao excluir titular", w)
		}
	}
}
