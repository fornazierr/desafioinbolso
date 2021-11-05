package broker

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"transferenciaapi/apiutil"
	"transferenciaapi/models"

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

func GetSaldo(saldo models.Saldo) ([]models.Saldo, error) {
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()
	rows, err := cn.Query(context.Background(), sqlGetSaldo(), saldo.ContaId, saldo.TitularId)
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

func NewSaldo(saldo models.Saldo) error {
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	com, err := cn.Exec(ctx, sqlNewSaldo(), saldo.TitularId, saldo.ContaId)
	if err != nil {
		log.Println("Erro no SQL \"RegistraSaldo\":", err.Error())
		return err
	}

	if com.RowsAffected() > 0 {
		log.Println("Registro salvo com sucesso.")
	}
	return nil
}

func NewRegistroSaldo(rs models.RegistroSaldo) error {
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	com, err := cn.Exec(ctx, sqlRegistroSaldo(), rs.TitularId, rs.ContaId, rs.Sinal, rs.Valor)
	if err != nil {
		log.Println("Erro no SQL \"RegistraSaldo\":", err.Error())
		return err
	}

	if com.RowsAffected() > 0 {
		log.Println("Registro salvo com sucesso.")
	}

	go atualizaSaldo(rs)

	return nil
}

func atualizaSaldo(rs models.RegistroSaldo) {
	log.Printf("Iniciando rotina de atualização de saldo para: %d - %d\n", rs.TitularId, rs.ContaId)

	var saldo models.Saldo

	aux, err := GetSaldo(models.Saldo{
		TitularId: rs.TitularId,
		ContaId:   rs.ContaId,
	})
	if err != nil {
		log.Println("Erro ao resgatar Saldo:", err.Error())
	}

	saldo = aux[0]

	err = saldo.ErroGetSaldo()
	if err != nil {
		log.Println("Cliente não possui registro de saldo válido")
		return
	}

	res, err := GetAllRegistrosSaldo(rs)
	if err != nil {
		log.Printf("Falha ao atualizar saldo de %d - %d\n", rs.TitularId, rs.ContaId)
		return
	}
	var novoSaldo float64
	for i := 0; i < len(res); i++ {
		registro := res[i]

		if registro.Sinal == 0 { //credita o valor
			novoSaldo += registro.Valor
		} else if registro.Sinal == 1 { //debita o valor
			novoSaldo -= registro.Valor
		}
	}
	//atualizando novo saldo
	saldo.Saldo = novoSaldo
	err = UpdateSaldo(saldo)
	if err != nil {
		log.Println("Não foi possível atualizado o saldo:", err.Error())
	} else {
		log.Printf("Saldo de %d - %d atualizado\n", saldo.TitularId, saldo.ContaId)
	}
}

func GetAllRegistrosSaldo(rs models.RegistroSaldo) ([]models.RegistroSaldo, error) {
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	rows, err := cn.Query(ctx, sqlGetRegistrosSaldo(), rs.TitularId, rs.ContaId)
	if err != nil {
		log.Println("Erro no SQL \"getAllRegistrosSaldo\":", err.Error())
		return nil, err
	}

	res := rowsToArrayRegistroSaldo(rows)

	return res, nil
}

func UpdateSaldo(saldo models.Saldo) error {
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()
	com, err := cn.Exec(ctx, sqlUpdateSaldo(), saldo.Saldo, saldo.TitularId, saldo.ContaId)
	if err != nil {
		log.Println("Erro no SQL \"RegistraSaldo\":", err.Error())
		return err
	}

	if com.RowsAffected() > 0 {
		log.Println("Registro salvo com sucesso.")
	}
	return nil
}

func RealizaTransferencia(tr models.Transferencia) error {
	log.Println("RealizaTransferencia: Iniciando operação")
	contaOrigem, err := getContaBancariaAPI(tr.ContaOrigemId)
	if err != nil {
		log.Println("RealizaTransferencia: erro ao buscar conta de origem ::", err.Error())
		return err
	}

	contaDestino, err := getContaBancariaAPI(tr.ContaOrigemId)
	if err != nil {
		log.Println("RealizaTransferencia: erro ao buscar conta de destino ::", err.Error())
		return err
	}

	log.Printf(fmt.Sprintf("Preparando transfercia de <%+v> para <%+v>\n", contaOrigem, contaDestino))

	return nil
}

func rowsToArraySaldo(rows pgx.Rows) []models.Saldo {
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
	return ret
}

func rowsToArrayRegistroSaldo(rows pgx.Rows) []models.RegistroSaldo {
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
	return ret
}
