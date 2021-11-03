package broker

import "transferenciabancariaapi/models"

func getConnection() {

}

func GetSaldo(saldo models.Saldo) (*[]models.Saldo, error) {
	return nil, nil
}

func RegistraSaldo(rs models.RegistroSaldo) error {
	go atualizaSaldo(rs)
	return nil
}

func atualizaSaldo(rs models.RegistroSaldo) {

}
