package broker

import (
	"contaapi/apiutil"
	"contaapi/models"
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

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
	config.MaxConns = 3
	config.MinConns = 1
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

func GetTitular() (*[]models.Titular, error) {
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	rows, err := cn.Query(context.Background(), sqlTitular())
	if err != nil {
		log.Println("Erro no SQL \"GetTitular\":", err.Error())
		return nil, err
	}

	if err = rows.Err(); err != nil {
		log.Println("Erro no SQL \"GetTitular\":", err.Error())
		return nil, err
	}

	res := rowsToArrayTitular(rows)

	return res, nil
}

func GetTitularById(id int) (*[]models.Titular, error) {
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	rows, err := cn.Query(context.Background(), sqlTitularById(), id)
	if err != nil {
		log.Println("Erro no SQL \"GetTitularById\":", err.Error())
		return nil, err
	}

	if err = rows.Err(); err != nil {
		log.Println("Erro no SQL \"GetTitularById\":", err.Error())
		return nil, err
	}

	res := rowsToArrayTitular(rows)

	return res, nil
}

func PostTitular(titular models.Titular) error {
	//salva um novo titular
	if titular.ID < 1 {
		log.Println("Salvando novo Titular.")
		err := newTitular(titular)
		return err
	} else {
		log.Println("Atulizando Titular.")
		err := updateTitular(titular)
		return err
	}
}

//(nome, cpf, email, nascimento, nomepai, nomemae, cidade, estado)
func newTitular(titular models.Titular) error {
	//verifica a exitencia de outro cadastro já feito
	res, err := getTitularByCPF(titular.CPF)
	if err != nil {
		log.Println("Erro verificando existencia de titular antes da criacao:", err.Error())
		return err
	}
	if res {
		return errors.New(fmt.Sprintf("titular já cadastrado para o CPF: %s", titular.CPF))
	}

	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	_, err = cn.Exec(context.Background(), sqlNewTitular(), titular.Nome, titular.CPF, titular.Email, titular.Nascimento, titular.NomePai, titular.NomeMae, titular.Cidade, titular.Estado)
	if err != nil {
		return nil
	}

	return nil
}

func updateTitular(titular models.Titular) error {
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	_, err := cn.Exec(ctx, sqlUpdateTitular(), titular.Nome, titular.CPF, titular.Email, titular.Nascimento, titular.NomePai, titular.NomeMae, titular.Cidade, titular.Estado, titular.ID)
	if err != nil {
		return nil
	}

	return nil
}

func getTitularByCPF(cpf string) (bool, error) {
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	rows, err := cn.Query(ctx, sqlEncontraTitularByCPDF(), cpf)
	if err != nil {
		return false, err
	}

	encontrado := false
	for rows.Next() {
		doc := ""
		rows.Scan(&doc)
		if doc == cpf {
			encontrado = true
		}
	}
	return encontrado, nil
}

func DeleteTitular(titular models.Titular) error {
	res, err := GetTitularById(titular.ID)
	if err != nil {
		log.Println("Erro verificando existencia de titular antes da exclusao:", err.Error())
		return err
	}
	if len(*res) < 1 {
		return errors.New("titular não encontrado para exclusao")
	}

	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	_, err = cn.Exec(ctx, sqlDeleteTitular(), titular.ID)
	if err != nil {
		return nil
	}

	return nil
}

func GetContaBancaria() (*[]models.ContaBancaria, error) {
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	rows, err := cn.Query(ctx, slqGetContaBancaria())
	if err != nil {
		log.Println("Erro no SQL \"GetContaBancaria\":", err.Error())
		return nil, err
	}

	if err = rows.Err(); err != nil {
		log.Println("Erro no SQL \"GetContaBancaria\":", err.Error())
		return nil, err
	}

	res := rowsToArrayContaBancaria(rows)

	return res, nil
}

func GetContaBancariaById(id int) (*[]models.ContaBancaria, error) {
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	rows, err := cn.Query(context.Background(), slqGetContaBancariaById(), id)
	if err != nil {
		log.Println("Erro no SQL \"GetContaBancariaById\":", err.Error())
		return nil, err
	}

	if err = rows.Err(); err != nil {
		log.Println("Erro no SQL \"GetContaBancariaById\":", err.Error())
		return nil, err
	}

	res := rowsToArrayContaBancaria(rows)

	return res, nil
}

func SaveContaBancaria(cb models.ContaBancaria) error {
	res, err := GetTitularById(cb.TitularId)
	if err != nil {
		log.Println("Erro verificando existencia de titular antes da criar nova conta bancaria:", err.Error())
		return err
	}
	if len(*res) < 1 {
		return errors.New("titular não encontrado para criar conta bancaria")
	}

	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	rows, err := cn.Query(context.Background(), sqlNewContaBancaria(), cb.CodigoBanco, cb.Agencia, cb.Conta, cb.Digito, cb.TitularId)
	if err != nil {
		return err
	}

	id, err := getReturningId(rows)
	if err != nil {
		return err
	}
	log.Println("Id Retornado: ", id)
	cb.ID = id
	go sendCriarNovoSaldo(cb)
	return nil
}

func getReturningId(rows pgx.Rows) (int, error) {
	var id int

	for rows.Next() {
		err := rows.Err()
		if err != nil {
			return 0, err
		}
		rows.Scan(&id)
	}
	return id, nil
}

func DeleteContaBancaria(cb models.ContaBancaria) error {
	//verificando existencia da conta
	res, err := GetContaBancariaById(cb.ID)
	if err != nil {
		log.Println("Erro verificando existencia de conta antes da exclusao:", err.Error())
		return err
	}
	if len(*res) < 1 {
		return errors.New("conta não encontrado para exclusao")
	}

	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	_, err = cn.Exec(context.Background(), sqlDeleteContaBancaria(), cb.ID)
	if err != nil {
		return nil
	}

	return nil
}

func rowsToArrayTitular(rows pgx.Rows) *[]models.Titular {
	var ret []models.Titular
	for rows.Next() {
		id := 0
		nome := ""
		cpf := ""
		email := ""
		nascimento := ""
		nomepai := ""
		nomemae := ""
		cidade := ""
		estado := ""
		err := rows.Scan(&id, &nome, &cpf, &email, &nascimento, &nomepai, &nomemae, &cidade, &estado)
		if err != nil {
			log.Println("Erro escaneando valores: ", err.Error())
		} else {
			titular := models.Titular{
				ID:         id,
				Nome:       nome,
				CPF:        cpf,
				Email:      email,
				Nascimento: nascimento,
				NomePai:    nomepai,
				NomeMae:    nomemae,
				Cidade:     cidade,
				Estado:     estado,
			}
			ret = append(ret, titular)
		}
	}
	return &ret
}

func rowsToArrayContaBancaria(rows pgx.Rows) *[]models.ContaBancaria {
	var ret []models.ContaBancaria
	for rows.Next() {
		id := 0
		codbanco := 0
		agencia := ""
		conta := ""
		digito := ""
		titular := 0
		err := rows.Scan(&id, &codbanco, &agencia, &conta, &digito, &titular)
		if err != nil {
			log.Println("Erro escaneando valores: ", err.Error())
		} else {
			contab := models.ContaBancaria{
				ID:          id,
				CodigoBanco: codbanco,
				Agencia:     agencia,
				Conta:       conta,
				Digito:      digito,
				TitularId:   titular,
			}
			ret = append(ret, contab)
		}
	}
	return &ret
}
