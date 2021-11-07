package broker

import (
	"boletoapi/apiutil"
	"boletoapi/models"
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

func NewBoleto(bol models.BoletoRequest) error {
	log.Println("NewBoleto: realizando chamada")
	var modelo models.MontaBoleto

	if bol.ContaOrigemId > 0 {
		existe, err := verificaContaBancaria(bol.ContaOrigemId)
		if err != nil {
			log.Println("NewBoleto:erro verificando exitencia da conta de origem:", err.Error())
			return err
		}
		if !existe {
			return errors.New("conta de origem não existe")
		}
	}

	if bol.ContaDestinoId > 0 {
		existe, err := verificaContaBancaria(bol.ContaDestinoId)
		if err != nil {
			log.Println("NewBoleto:erro verificando exitencia da conta de destino:", err.Error())
			return err
		}
		if !existe {
			return errors.New("conta de destino não existe")
		}
	}

	log.Println("Iniciando montagem de boleto")
	modelo.CodigoBanco = bol.CodigoBanco
	modelo.Agencia = fmt.Sprintf("%04d", bol.Agencia)
	modelo.Carteira = fmt.Sprintf("%02d", bol.Carteira)
	fatorVencimento, err := apiutil.CalculaFatorVencimento(bol.DataVencimento)
	if err != nil {
		return err
	}
	modelo.FatorVencimento = fmt.Sprintf("%04d", fatorVencimento)
	modelo.Valor = apiutil.FormataValorBoleto(bol.Valor, 10, false)
	modelo.NossoNumero = apiutil.FormataNossoNumero(bol.NossoNumero, 11)
	modelo.CodigoBeneficiario = fmt.Sprintf("%07d", bol.CodigoBeneficiario)
	fmt.Printf("ModeloGerado gerado: %+v\n", modelo)
	log.Println("Gerando código de barras")
	codigoBarras, err := apiutil.GeraCodigoBarras(modelo)
	if err != nil {
		return err
	}
	log.Println("Código de barras gerado:", codigoBarras.CodigoBarras)
	log.Println("Gerando linha digitável")
	linhaDigitavel, err := apiutil.GetLinhaDigitavel(codigoBarras.CodigoBarras, modelo)
	if err != nil {
		return err
	}

	log.Println("Linha digitavel gerada:", linhaDigitavel.LinhaDigitavelFormatada)
	salvarBoleto := models.Boleto{
		ID:                 0,
		ContaOrigemId:      bol.ContaOrigemId,
		ContaDestinoId:     bol.ContaDestinoId,
		CodigoBanco:        modelo.CodigoBanco,
		Agencia:            modelo.Agencia,
		Carteira:           modelo.Carteira,
		DataVencimento:     bol.DataVencimento,
		Valor:              bol.Valor,
		NossoNumero:        modelo.NossoNumero,
		CodigoBeneficiario: modelo.CodigoBeneficiario,
		CodigoBarras:       codigoBarras.CodigoBarras,
		LinhaDigitavel:     linhaDigitavel.LinhaDigitavel,
	}

	log.Printf("Boleto gerado: $%+v\n", salvarBoleto)

	err = saveBoleto(salvarBoleto)
	if err != nil {
		return err
	}

	return nil
}

//Função que busca
func verificaContaBancaria(contaId int) (bool, error) {
	_, err := apiGetContaBancaria(contaId)
	if err != nil {
		log.Println("verificaContaBancaria: erro resgatando conta: ", err.Error())
		return false, err
	}

	return true, nil
}

//(contaorigem_id, contadestino_id, codigobanco, agencia, carteira, datavencimento, valor, nossonumero, codigobeneficiario, linhadigitavel, codigobarras)
func saveBoleto(b models.Boleto) error {
	log.Println("GetAllBoleto: realizando chamada")
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	_, err := cn.Exec(context.Background(), sqlNewBoleto(), b.ContaOrigemId, b.ContaDestinoId, b.CodigoBanco, b.Agencia, b.Carteira, b.DataVencimento, b.Valor, b.NossoNumero, b.CodigoBeneficiario, b.LinhaDigitavel, b.CodigoBarras)
	if err != nil {
		log.Println("saveBoleto: Erro durante sql-> ", err.Error())
		return err
	}

	return nil
}

func GetAllBoleto() ([]models.Boleto, error) {
	log.Println("GetAllBoleto: realizando chamada")
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	rows, err := cn.Query(context.Background(), sqlGetAllBoleto())
	if err != nil {
		log.Println("saveBoleto: Erro durante sql-> ", err.Error())
		return nil, err
	}

	res, err := rowsToArray(rows)
	if err != nil {
		log.Println("getBoletoById: Erro no parse do array", err.Error())
		return nil, err
	}

	return res, nil
}

func GetBoleto(bol models.Boleto) ([]models.Boleto, error) {
	log.Println("GetBoleto: realizando chamada")
	//busca pelo id do boleto
	if bol.ID > 0 {
		res, err := getBoletoById(bol.ID)
		if err != nil {
			return nil, err
		}
		if len(res) > 0 {
			return res, nil
		} else {
			log.Println("GetBoleto: Nenhum boleto encontrado por ID.")
		}
	}
	//caso não ache resultados procura pela combinacao de banco & agencia & carteira
	if bol.Agencia != "" && bol.CodigoBanco != "" && bol.Carteira != "" { //busca por dados do booleto
		res, err := getBoletoByDados(bol)
		if err != nil {
			return nil, err
		}

		if len(res) > 0 {
			return res, nil
		} else {
			log.Println(fmt.Sprintf("GetBoleto: Nenhum boleto encontrado para a o banco %s da agencia %s da carteira %s", bol.CodigoBanco, bol.Agencia, bol.Carteira))
		}
	}
	//caso nao ache resultados, procura pelo conta de origem
	if bol.ContaOrigemId > 0 {
		res, err := getBoletosByContaOrigem(bol.ContaOrigemId)
		if err != nil {
			return nil, err
		}

		if len(res) > 0 {
			log.Println(fmt.Sprintf("GetBoleto: Nenhum boleto encontrado para a conta de origem %d", bol.ContaOrigemId))
			return res, nil
		}
	}
	return nil, nil
}

func getBoletoById(id int) ([]models.Boleto, error) {
	log.Println("getBoletoById: realizando chamada")
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	rows, err := cn.Query(context.Background(), sqlGetBoletoById(), id)
	if err != nil {
		log.Println("saveBoleto: Erro durante sql-> ", err.Error())
		return nil, err
	}

	res, err := rowsToArray(rows)
	if err != nil {
		log.Println("getBoletoById: Erro no parse do array", err.Error())
		return nil, err
	}

	return res, nil
}

func getBoletoByDados(bol models.Boleto) ([]models.Boleto, error) {
	log.Println("getBoletoByDados: realizando chamada")
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	rows, err := cn.Query(context.Background(), sqlGetBoletoByDados1(), bol.CodigoBanco, bol.Agencia, bol.Carteira)
	if err != nil {
		log.Println("saveBoleto: Erro durante sql-> ", err.Error())
		return nil, err
	}

	res, err := rowsToArray(rows)
	if err != nil {
		log.Println("getBoletoByDados: Erro no parse do array", err.Error())
		return nil, err
	}

	return res, nil
}

func getBoletosByContaOrigem(id int) ([]models.Boleto, error) {
	log.Println("getBoletosByContaOrigem: realizando chamada")
	ctx := context.Background()
	cn, _ := Pool.Acquire(ctx)
	defer cn.Release()

	rows, err := cn.Query(context.Background(), sqlGetBoletoByContaOrigem(), id)
	if err != nil {
		log.Println("getBoletosByContaOrigem: Erro durante sql-> ", err.Error())
		return nil, err
	}

	res, err := rowsToArray(rows)
	if err != nil {
		log.Println("getBoletosByContaOrigem: Erro no parse do array", err.Error())
		return nil, err
	}

	return res, nil
}

//id, contaorigem_id, contadestino_id, codigobanco, agencia, carteira, datavencimento, valor, nossonumero, codigobeneficiario, linhadigitavel, codigobarras
func rowsToArray(rows pgx.Rows) ([]models.Boleto, error) {
	var res []models.Boleto
	for rows.Next() {
		id := 0
		cori := 0
		cdest := 0
		codbanco := ""
		ag := ""
		car := ""
		dvenc := ""
		var valor float64
		nosso := ""
		bene := ""
		linha := ""
		barra := ""
		rows.Scan(&id, &cori, &cdest, &codbanco, &ag, &car, &dvenc, &valor, &nosso, &bene, &linha, &barra)
		if err := rows.Err(); err != nil {
			log.Println("rowsToArray Erro: ", err.Error())
			return nil, err
		}
		bol := models.Boleto{
			ID:                 id,
			ContaOrigemId:      cori,
			ContaDestinoId:     cdest,
			CodigoBanco:        codbanco,
			Agencia:            ag,
			Carteira:           car,
			DataVencimento:     dvenc,
			Valor:              valor,
			NossoNumero:        nosso,
			CodigoBeneficiario: bene,
			CodigoBarras:       barra,
			LinhaDigitavel:     linha,
		}
		res = append(res, bol)
	}

	return res, nil
}
