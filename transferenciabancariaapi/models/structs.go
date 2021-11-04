package models

import "errors"

type Saldo struct {
	TitularId int     `json:"titularid"`
	ContaId   int     `json:"contaid"`
	Saldo     float64 `json:"saldo"`
}

func (s *Saldo) ErroGetSaldo() error {
	erros := ""
	if s.ContaId < 1 {
		erros += "ID da Conta vazio"
	}
	if s.TitularId < 1 {
		erros += "ID do Titular vazio."
	}

	if erros != "" {
		return errors.New(erros)
	} else {
		return nil
	}
}

type RegistroSaldo struct {
	TitularId int     `json:"titularid"`
	ContaId   int     `json:"contaid"`
	Sinal     int     `json:"sinal"`
	Valor     float64 `json:"valor"`
}

func (t *RegistroSaldo) Erros() error {
	erros := ""
	if t.ContaId < 1 {
		erros += "ID da Conta vazio."
	}
	if t.TitularId < 1 {
		erros += "ID do Titular vazio."
	}
	if t.Sinal != 0 && t.Sinal != 1 {
		erros += "Sinal diferente de 0 e 1"
	}

	if erros != "" {
		return errors.New(erros)
	} else {
		return nil
	}
}

func (t *RegistroSaldo) ErrosGET() error {
	erros := ""
	if t.ContaId < 1 {
		erros += "ID da Conta vazio."
	}
	if t.TitularId < 1 {
		erros += "ID do Titular vazio."
	}

	if erros != "" {
		return errors.New(erros)
	} else {
		return nil
	}
}

type Transferencia struct {
	ContaOrigemId  int     `json:"contaorigemid"`
	ContaDestinoId int     `json:"contadestinoid"`
	Valor          float64 `json:"valor"`
}

func (t *Transferencia) Erros() error {
	erros := ""
	if t.ContaDestinoId < 1 {
		erros += "Conta de Destino não informada."
	}
	if t.ContaOrigemId < 1 {
		erros += "Conta de Origem não informada."
	}
	if t.Valor < 0.0000001 {
		erros += "Valores negativos não são aceitos."
	}

	if erros != "" {
		return errors.New(erros)
	} else {
		return nil
	}
}

type Config struct {
	DB_NAME           string `json:"bdname"`
	DB_PORT           string `json:"dbport"`
	DB_USER           string `json:"dbuser"`
	DB_PASS           string `json:"dbpass"`
	DB_HOST           string `json:"dbhost"`
	URL_PORT          string `json:"urlport"`
	URL_HOST          string `json:"urlhost"`
	CONTABANCARIA_API string `json:"contabancariaapi"`
}

type ReturnMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ContaBancaria struct {
	ID          int    `json:"id"`
	CodigoBanco int    `json:"codigobanco"`
	Agencia     string `json:"agencia"`
	Conta       string `json:"conta"`
	Digito      string `json:"digito"`
	TitularId   int    `json:"titularid"`
}

func (cb *ContaBancaria) Erros() error {
	erros := ""
	if cb.Agencia == "" {
		erros += "Agencia vazio."
	}
	if cb.CodigoBanco < 1 {
		erros += "Condigo baancario vazio."
	}
	if cb.Conta == "" {
		erros += "Conta vazio."
	}
	if cb.Digito == "" {
		erros += "Digito verificador vazio."
	}
	if cb.TitularId < 1 {
		erros += "ID do Titular vazio."
	}

	if erros != "" {
		return errors.New(erros)
	} else {
		return nil
	}
}
