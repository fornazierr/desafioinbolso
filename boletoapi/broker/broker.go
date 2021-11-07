package broker

import (
	"boletoapi/apiutil"
	"boletoapi/models"
	"context"
	"fmt"
	"log"
	"sync"
	"time"

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
	var modelo models.MontaBoleto

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
	codigoBarras, err := apiutil.GeraCodigoBarras(modelo)
	if err != nil {
		return err
	}
	log.Println("Código de barras gerado:", codigoBarras.CodigoBarras)

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

//(contaorigem_id, contadestino_id, codigobanco, agencia, carteira, datavencimento, valor, nossonumero, codigobeneficiario, linhadigitavel, codigobarras)
func saveBoleto(b models.Boleto) error {
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

func GetBoleto(bol models.Boleto) ([]models.Boleto, error) {
	if bol.ID > 0 { //busca pelo id do boleto
		res, err := getBoletoById(bol.ID)
		if err != nil {
			return nil, err
		}
		return res, nil
	} else { //busca por dados do booleto
		res, err := getBoletoByDados(bol)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

func getBoletoById(id int) ([]models.Boleto, error) {
	return nil, nil
}

func getBoletoByDados(bol models.Boleto) ([]models.Boleto, error) {
	return nil, nil
}
