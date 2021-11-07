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
		erros += "\"contaid\" vazio"
	}
	if s.TitularId < 1 {
		erros += "\"titularid\" vazio."
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
		erros += "\"contaid\" vazio."
	}
	if t.TitularId < 1 {
		erros += "\"titularid\" vazio."
	}
	if t.Sinal != 0 && t.Sinal != 1 {
		erros += "\"sinal\" diferente de 0 e 1"
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
		erros += "\"contaid\" vazio."
	}
	if t.TitularId < 1 {
		erros += "\"titularid\" vazio."
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
		erros += "\"contadestinoid\" vazio."
	}
	if t.ContaOrigemId < 1 {
		erros += "\"contaorigemid\" vazio."
	}
	if t.Valor < 0.0000001 {
		erros += "\"valor\" negativos ou zerados não são aceitos."
	}

	if erros != "" {
		return errors.New(erros)
	} else {
		return nil
	}
}

type Config struct {
	DB_NAME    string `json:"bdname"`
	DB_PORT    string `json:"dbport"`
	DB_USER    string `json:"dbuser"`
	DB_PASS    string `json:"dbpass"`
	DB_HOST    string `json:"dbhost"`
	URL_PORT   string `json:"urlport"`
	URL_HOST   string `json:"urlhost"`
	CONTA_API  string `json:"contaapi"`
	BOLETO_API string `json:"boletoapi"`
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
		erros += "\"titularid\" vazio."
	}

	if erros != "" {
		return errors.New(erros)
	} else {
		return nil
	}
}

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
