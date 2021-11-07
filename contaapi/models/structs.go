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
		erros += "\"cpf\" vazio."
	}
	if t.ID == 0 {
		erros += "\"id\" vazio."
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
		erros += "\"cpf\" vazio."
	}
	if t.Cidade == "" {
		erros += "\"cidade\" vazio."
	}
	if t.Email == "" {
		erros += "\"email\" vazio."
	}
	if t.Estado == "" {
		erros += "\"estado\" vazio."
	}
	if t.Nascimento == "" {
		erros += "\"nascimento\" vazio."
	}
	if t.Nome == "" {
		erros += "\"nome\" vazio"
	}
	if t.NomeMae == "" {
		erros += "\"nomemae\" vazio."
	}
	if t.NomePai == "" {
		erros += "\"nomepai\" vazio."
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
		erros += "\"agencia\" vazio."
	}
	if cb.CodigoBanco < 1 {
		erros += "\"codigobanco\" vazio."
	}
	if cb.Conta == "" {
		erros += "\"conta\" vazio."
	}
	if cb.Digito == "" {
		erros += "\"digito\" vazio."
	}
	if cb.TitularId < 1 {
		erros += "\"id\" vazio."
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
		erros += "\"id\" vazio."
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
	BOLETO_API       string `json:"boletoapi"`
}

type ReturnMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type RequestSaldo struct {
	TitularId int `json:"titularid"`
	ContaId   int `json:"contaid"`
}
