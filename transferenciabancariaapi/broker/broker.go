package broker

import (
	"context"
	"log"
	"transferenciabancariaapi/apiutil"
	"transferenciabancariaapi/models"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func getConnection() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.Connect(context.Background(), apiutil.GetUrlConnection())
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}

func GetSaldo(saldo models.Saldo) (*[]models.Saldo, error) {
	pool, err := getConnection()
	if err != nil {
		log.Println("Erro na função \"GetSaldo\": " + err.Error())
		return nil, err
	}
	defer pool.Close()
	rows, err := pool.Query(context.Background(), sqlGetSaldo())
	if err != nil {
		log.Println("Erro no SQL \"GetSaldo\":", err.Error())
		return nil, err
	}

	if err = rows.Err(); err != nil {
		log.Println("Erro no SQL \"GetSaldo\":", err.Error())
		return nil, err
	}

	res := rowsToArraySaldo(rows)

	return res, nil
}

func RegistraSaldo(rs models.RegistroSaldo) error {
	pool, err := getConnection()
	if err != nil {
		log.Println("Erro na função \"RegistraSaldo\": " + err.Error())
		return err
	}
	defer pool.Close()

	go atualizaSaldo(rs)
	return nil
}

func atualizaSaldo(rs models.RegistroSaldo) {

}

func RealizaTransferencia(tr models.Transferencia) error {
	pool, err := getConnection()
	if err != nil {
		log.Println("Erro na função \"GetTitular\": " + err.Error())
		return err
	}
	defer pool.Close()
	return nil
}

func rowsToArraySaldo(rows pgx.Rows) *[]models.Saldo {
	var ret []models.Saldo
	for rows.Next() {
		conta := 0
		titular := 0
		var saldo float64
		err := rows.Scan(&conta, &titular, &saldo)
		if err != nil {
			log.Println("Erro escaneando valores: ", err.Error())
		} else {
			saldo := models.Saldo{
				TitularId: titular,
				ContaId:   conta,
				Saldo:     saldo,
			}
			ret = append(ret, saldo)
		}
	}
	return &ret
}

func rowsToArrayrRegistroSaldo(rows pgx.Rows) *[]models.RegistroSaldo {
	var ret []models.RegistroSaldo
	for rows.Next() {
		titular := 0
		conta := 0
		var sinal int
		var valor float64
		err := rows.Scan(&titular, &conta, &sinal, &valor)
		if err != nil {
			log.Println("Erro escaneando valores:", err.Error())
		} else {
			rs := models.RegistroSaldo{
				TitularId: titular,
				ContaId:   conta,
				Sinal:     sinal,
				Valor:     valor,
			}
			ret = append(ret, rs)
		}
	}
	return &ret
}
