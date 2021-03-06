package router

import (
	"boletoapi/broker"
	"boletoapi/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type BoletoRouter struct {
	Router *mux.Router
}

func StartRouter() *BoletoRouter {
	r := &BoletoRouter{
		Router: mux.NewRouter(),
	}

	r.initRoutes()

	return r
}

func (br *BoletoRouter) initRoutes() {
	br.Router.HandleFunc("/getallboleto", getAllBoleto).Methods("GET")
	br.Router.HandleFunc("/getboleto", getBoleto).Methods("POST")
	br.Router.HandleFunc("/saveboleto", postBoleto).Methods("POST")
	broker.InitBroker()
}

func getAllBoleto(w http.ResponseWriter, r *http.Request) {
	log.Println("Nova requisição em \"getAllBoleto\".")
	res, err := broker.GetAllBoleto()
	if err != nil {
		sendError(err, w)
	} else {
		sendOk(res, w)
	}
}

func getBoleto(w http.ResponseWriter, r *http.Request) {
	log.Println("Nova requisição em \"getBoleto\".")
	var boleto models.Boleto
	json.NewDecoder(r.Body).Decode(&boleto)
	log.Printf("Objeto da requisição: %+v\n", boleto)
	res, err := broker.GetBoleto(boleto)
	if err != nil {
		sendError(err, w)
		return
	}
	if len(res) > 0 {
		sendOk(res, w)
		return
	}
	if res == nil {
		sendOk("Nenhum resultado encontrado.", w)
	}

}

func postBoleto(w http.ResponseWriter, r *http.Request) {
	log.Println("Nova requisição em \"postBoleto\".")
	var boleto models.BoletoRequest
	json.NewDecoder(r.Body).Decode(&boleto)
	log.Printf("Objeto da requisição: %+v\n", boleto)
	err := boleto.Erros()
	if err != nil {
		sendError(err, w)
	} else {
		err = broker.NewBoleto(boleto)
		if err != nil {
			sendError(err, w)
		} else {
			sendOk("Gravado/Atualizado com sucesso", w)
		}
	}
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
