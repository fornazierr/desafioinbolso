package broker

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"transferenciabancariaapi/apiutil"
	"transferenciabancariaapi/models"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var Pool *pgxpool.Pool
var mu sync.Mutex = sync.Mutex{}

func InitBroker() {
	log.Println("Iniciando database.")
	mu.Lock()
	defer mu.Unlock()

	var connStr string = apiutil.GetUrlConnection()
	log.Printf("Url de conexao: %s\n", connStr)
	config, err := pgxpool.ParseConfig(connStr)
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 2 * time.Minute
	config.MaxConnLifetime = 2 * time.Minute

	if err != nil {
		msg := fmt.Sprintf("Falha ao configurar pool. Url utilizada: %s\n", connStr)
		log.Fatalln(msg, err)
		return
	}

	ctx := context.Background()

	Pool, err = pgxpool.ConnectConfig(ctx, config)

	if err != nil {
		log.Fatal("Erro conectando ao banco de dados: ", err)
	}

	// inicia as conexões
	for i := 0; i < int(config.MinConns); i++ {
		cn, _ := Pool.Acquire(ctx)
		defer cn.Release()
	}
	log.Println("Banco de dados inicializado")
}

func GetSaldo(saldo models.Saldo) (*[]models.Saldo, error) {
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()
	rows, err := cn.Query(context.Background(), sqlGetSaldo())
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
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	go atualizaSaldo(rs)
	return nil
}

func atualizaSaldo(rs models.RegistroSaldo) {

}

func RealizaTransferencia(tr models.Transferencia) error {
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()
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
