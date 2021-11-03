package broker

import (
	"contabancariaapi/models"
	"contabancariaapi/utilapi"
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func getConnection() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.Connect(context.Background(), utilapi.GetUrlConnection())
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}

func GetTitular() (*[]models.Titular, error) {
	pool, err := getConnection()
	if err != nil {
		log.Println("Erro na função \"GetTitular\": " + err.Error())
		return nil, err
	}
	defer pool.Close()

	rows, err := pool.Query(context.Background(), sqlTitular())
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
	pool, err := getConnection()
	if err != nil {
		log.Println("Erro na função \"GetTitularById\": " + err.Error())
		return nil, err
	}
	defer pool.Close()

	rows, err := pool.Query(context.Background(), sqlTitularById(), id)
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
	pool, err := getConnection()
	if err != nil {
		log.Println("Erro na função \"newTitular\": " + err.Error())
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(context.Background(), sqlNewTitular(), titular.Nome, titular.CPF, titular.Email, titular.Nascimento, titular.NomePai, titular.NomeMae, titular.Cidade, titular.Estado)
	if err != nil {
		return nil
	}

	return nil
}

func updateTitular(titular models.Titular) error {
	pool, err := getConnection()
	if err != nil {
		log.Println("Erro na função \"updateTitular\": " + err.Error())
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(context.Background(), sqlUpdateTitular(), titular.Nome, titular.CPF, titular.Email, titular.Nascimento, titular.NomePai, titular.NomeMae, titular.Cidade, titular.Estado, titular.ID)
	if err != nil {
		return nil
	}

	return nil
}

func DeleteTitular(titular models.Titular) error {
	pool, err := getConnection()
	if err != nil {
		log.Println("Erro na função \"DeleteTitular\": " + err.Error())
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(context.Background(), sqlDeleteTitular(), titular.ID)
	if err != nil {
		return nil
	}

	return nil
}

func GetContaBancaria() (*[]models.ContaBancaria, error) {
	pool, err := getConnection()
	if err != nil {
		log.Println("Erro na conexão \"GetContaBancaria\": " + err.Error())
		return nil, err
	}
	defer pool.Close()

	rows, err := pool.Query(context.Background(), slqGetContaBancaria())
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
	pool, err := getConnection()
	if err != nil {
		log.Println("Erro na conexão \"GetContaBancariaById\": " + err.Error())
		return nil, err
	}
	defer pool.Close()

	rows, err := pool.Query(context.Background(), slqGetContaBancariaById(), id)
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
	pool, err := getConnection()
	if err != nil {
		log.Println("Erro na função \"SaveContaBancaria\": " + err.Error())
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(context.Background(), sqlNewContaBancaria(), cb.CodigoBanco, cb.Agencia, cb.Conta, cb.Digito, cb.TitularId)
	if err != nil {
		return err
	}

	return nil
}

func DeleteContaBancaria(cb models.ContaBancaria) error {
	pool, err := getConnection()
	if err != nil {
		log.Println("Erro na função \"DeleteTitular\": " + err.Error())
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(context.Background(), sqlDeleteContaBancaria(), cb.ID)
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
