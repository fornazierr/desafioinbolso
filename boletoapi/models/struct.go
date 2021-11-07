package models

import (
	"errors"
	"fmt"
	"strings"
)

type Boleto struct {
	ID                 int     `json:"id"`
	ContaOrigemId      int     `json:"contaorigem"`
	ContaDestinoId     int     `json:"contadestino"`
	CodigoBanco        string  `json:"codbanco"`
	Agencia            string  `json:"agencia"`
	Carteira           string  `json:"carteira"`
	DataVencimento     string  `json:"datavencimento"`
	Valor              float64 `json:"valor"`
	NossoNumero        string  `json:"nossonumero"`
	CodigoBeneficiario string  `json:"codbeneficiario"`
	CodigoBarras       string  `json:"codigobarras"`
	LinhaDigitavel     string  `json:"linhadigitavel"`
}

type BoletoRequest struct {
	ID                 int     `json:"id"`
	ContaOrigemId      int     `json:"contaorigem"`
	ContaDestinoId     int     `json:"contadestino"`
	CodigoBanco        string  `json:"codbanco"`
	Agencia            int32   `json:"agencia"`
	Carteira           int16   `json:"carteira"`
	DataVencimento     string  `json:"datavencimento"`
	Valor              float64 `json:"valor"`
	NossoNumero        string  `json:"nossonumero"`
	CodigoBeneficiario int32   `json:"codbeneficiario"`
}

func (bol *BoletoRequest) Erros() error {
	e := ""

	if bol.Agencia < 1 {
		e += "\"agencia\" vazio."
	}
	if bol.Carteira < 1 {
		e += "\"carteira\" vazio."
	}
	if bol.CodigoBeneficiario < 1 {
		e += "\"codbeneficiario\" vazio."
	}
	if bol.DataVencimento == "" {
		e += "\"datavencimento\" vazio."
	}
	if bol.NossoNumero == "" {
		e += "\"nossonumero\" vazio."
	}
	if bol.CodigoBanco == "" {
		e += "\"codbanco\" vazio."
	}
	if len(strings.TrimSpace(bol.CodigoBanco)) > 3 {
		e += "\"codbanco\" pode ter no máximo 3 digitos."
	}
	if bol.Valor < 0.000001 {
		e += "\"valor\" vazio ou menor que 0.0."
	}

	if e != "" {
		return errors.New(e)
	}
	return nil
}

type MontaBoleto struct {
	CodigoBanco        string
	Moeda              string
	DigitoVerificador  string
	FatorVencimento    string
	Valor              string
	Agencia            string
	Carteira           string
	NossoNumero        string
	CodigoBeneficiario string
	Zero               string
}

func (mt *MontaBoleto) getCodBarrasAux() {

}

func (mt *MontaBoleto) VerificaCampos() error {
	e := ""

	if len(mt.Agencia) != 4 {
		e += fmt.Sprintf("Campo agencia deveria possuir 4 digitos, mas foi gerado %d digitos com valor %s ", len(mt.Agencia), mt.Agencia)
	}
	if len(mt.Carteira) != 2 {
		e += fmt.Sprintf("Campo carteira deveria possuir 2 digitos, mas foi gerado com %d digito e com valor %s", len(mt.Carteira), mt.Carteira)
	}
	if mt.FatorVencimento == "0" {
		e += "O vencimento informado é inválido. A data de vencimento deve ser no formato yyyy-MM-dd e precisa ser superior a 06/10/1997."
	}
	if len(mt.CodigoBeneficiario) != 7 {
		e += fmt.Sprintf("Campo código do beneficiario deveria possuir 7 digitos, mas foi gerado com %d digito e com valor %s", len(mt.CodigoBeneficiario), mt.CodigoBeneficiario)
	}
	if len(mt.NossoNumero) != 11 {
		e += fmt.Sprintf("Campo nosso numero deveria possuir 11 digitos, mas foi gerado com %d digito e com valor %s", len(mt.NossoNumero), mt.NossoNumero)
	}

	if e != "" {
		return errors.New(e)
	}
	return nil
}

type CodigoBarras struct {
	CodigoBanco     string `json:"codigobanco"`
	Moeda           string `json:"moeda"`
	DV              string `json:"dv"`
	FatorVencimento string `json:"fatorvencimento"`
	Valor           string `json:"valor"`
	CampoLivre      string `json:"campolivre"`
	CodigoBarras    string `json:"codigobarras"`
}

type LinhaDigitavel struct {
	LinhaDigitavel          string `json:"linhadigitavel"`
	LinhaDigitavelFormatada string `json:"linhaformatada"`
}

type Config struct {
	DB_NAME          string `json:"bdname"`
	DB_PORT          string `json:"dbport"`
	DB_USER          string `json:"dbuser"`
	DB_PASS          string `json:"dbpass"`
	DB_HOST          string `json:"dbhost"`
	URL_PORT         string `json:"urlport"`
	URL_HOST         string `json:"urlhost"`
	TRANFERENCIA_API string `json:"transferenciaapi"`
	CONTA_API        string `json:"boletoapi"`
}

type ReturnMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
