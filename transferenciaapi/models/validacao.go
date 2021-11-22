package models

import (
	"transferenciaapi/validacao"
)

//Realiza a validação da struct RegistroSaldo
func (rs *RegistroSaldo) Validacao() {

}

//Realiza a validação da struct Saldo
func (s *Saldo) Validacao() {

}

//Realiza a validacao da struct Transferencia
func (t *Transferencia) Validacao() {

}

//Realiza a validacao da struct ContaBancaria
func (cb *ContaBancaria) Validacao() {

}

//Realiza a validacao da struct Titular
func (t *Titular) Validacao() error {
	err := validacao.GetInstance().Struct(t)
	return err
}
