package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"transferenciabancariaapi/broker"
	"transferenciabancariaapi/models"

	"github.com/gorilla/mux"
)

type TransferenciaRouter struct {
	Router *mux.Router
}

func StartRouter() *TransferenciaRouter {
	r := &TransferenciaRouter{
		Router: mux.NewRouter(),
	}

	r.initRoutes()

	return r
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

func (r *TransferenciaRouter) initRoutes() {
	r.Router.HandleFunc("/getsaldo", getSaldoConta).Methods("POST")
	r.Router.HandleFunc("/registrasaldo", registraSaldo).Methods("POST")
	r.Router.HandleFunc("/realizatransferencia", realizaTransferencia).Methods("POST")
}

func getSaldoConta(w http.ResponseWriter, r *http.Request) {
	log.Println("Nova requisição em \"getSaldoConta\".")
	var saldo models.Saldo
	json.NewDecoder(r.Body).Decode(&saldo)
	err := saldo.ErroGetSaldo()
	if err != nil {
		sendError(err, w)
	} else {
		log.Printf("Objeto da requisição: %+v\n", saldo)
		res, err := broker.GetSaldo(saldo)
		if err != nil {
			sendError(err, w)
		} else {
			sendOk(res, w)
		}
	}
}

func registraSaldo(w http.ResponseWriter, r *http.Request) {
	log.Println("Nova requisição em \"registraSaldo\".")
	var rs models.RegistroSaldo
	json.NewDecoder(r.Body).Decode(&rs)
	err := rs.Erros()
	if err != nil {
		sendError(err, w)
	} else {
		log.Printf("Objeto da requisição: %+v\n", rs)
		err := broker.RegistraSaldo(rs)
		if err != nil {
			sendError(err, w)
		} else {
			sendOk("Sucesso ao registrar saldo.", w)
		}
	}
}

func realizaTransferencia(w http.ResponseWriter, r *http.Request) {
	log.Println("Nova requisição em \"realizaTransferencia\".")
	var transf models.Transferencia
	json.NewDecoder(r.Body).Decode(&transf)
	err := transf.Erros()
	if err != nil {
		sendError(err, w)
	} else {
		log.Printf("Objeto da requisição: %+v\n", transf)
		err := broker.RegistraSaldo(transf)
		if err != nil {
			sendError(err, w)
		} else {
			sendOk("Sucesso ao registrar saldo.", w)
		}
	}
}
