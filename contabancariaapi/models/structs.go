package models

import "errors"

type Titular struct {
	ID         int    `json:"id"`
	Nome       string `json:"nome"`
	CPF        string `json:"cpf"`
	Email      string `json:"email"`
	Nascimento string `json:"nascimento"`
	NomePai    string `json:"nomepai"`
	NomeMae    string `json:"nomemae"`
	Cidade     string `json:"cidade"`
	Estado     string `json:"estado"`
}

//Verifica na estrutura TITULAR erros para operação de delete
func (t *Titular) ErrosDelete() error {
	erros := ""
	if t.CPF == "" {
		erros += "CPF vazio."
	}
	if t.ID == 0 {
		erros += "ID vazio."
	}

	if erros != "" {
		return errors.New(erros)
	} else {
		return nil
	}
}

//Veriica na estrutura TITULAR erros para gravação ou atualização de dados
func (t *Titular) Erros() error {
	erros := ""
	if t.CPF == "" {
		erros += "CPF vazio."
	}
	if t.Cidade == "" {
		erros += "Cidade vazia."
	}
	if t.Email == "" {
		erros += "Email vazio."
	}
	if t.Estado == "" {
		erros += "Estado vazio."
	}
	if t.Nascimento == "" {
		erros += "Nascimento vazio."
	}
	if t.Nome == "" {
		erros += "Nome vazio"
	}
	if t.NomeMae == "" {
		erros += "Nome da Mae vazio."
	}
	if t.NomePai == "" {
		erros += "Nome do Pai vazio."
	}

	if erros != "" {
		return errors.New(erros)
	} else {
		return nil
	}
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

func (cb *ContaBancaria) ErrosDelete() error {
	erros := ""
	if cb.ID < 1 {
		erros += "ID vazio."
	}

	if erros != "" {
		return errors.New(erros)
	} else {
		return nil
	}
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
}

type ReturnMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type RequestSaldo struct {
	TitularId int `json:"titularid"`
	ContaId   int `json:"contaid"`
}
