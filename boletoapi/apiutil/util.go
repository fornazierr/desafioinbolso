package apiutil

import (
	"boletoapi/models"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const formatoData = "2006-01-02T15:04:05Z"
const dataBasePadrao = "1997-10-07T00:00:00Z"

func CalculaFatorVencimento(data string) (int16, error) {
	//sem vencimento ou zerado
	if data == "" {
		return int16(0), nil
	}

	dataVencimento, err := time.Parse(formatoData, formataData(data))
	if err != nil {
		log.Println("CalculaFatoVencimento: Erro no parse de data:", err.Error())
		return int16(0), err
	}

	dataBase, _ := time.Parse(formatoData, dataBasePadrao)

	if dataVencimento.Before(dataBase) {
		msg := "data de vencimento anterior à data base"
		log.Println("CalculaFatoVencimento: Erro,", msg)
		return int16(0), errors.New(msg)
	}

	dias := dataVencimento.Sub(dataBase).Hours() / 24

	return int16(dias), nil
}

func formataData(data string) string {
	return data + "T00:00:00Z"
}

func FormataValorBoleto(valor float64, tamanho int, zerar bool) string {
	if zerar {
		return strings.Repeat("0", tamanho)
	}

	valorFormatado := fmt.Sprintf("%.2f", valor)
	valorFormatado = strings.Replace(valorFormatado, ".", "", -1)
	valorFormatado = strings.Repeat("0", tamanho-len(valorFormatado)) + valorFormatado

	return valorFormatado
}

func FormataNossoNumero(numero string, tamanho int) string {
	if _, err := strconv.Atoi(numero); err != nil {
		return ""
	}

	if len(numero) > tamanho {
		return numero
	}

	return strings.Repeat("0", tamanho-len(numero)) + numero
}

func CalculaBaseCodigoBarras(codigo string, min int, max int) int {
	i := len(codigo)
	mult := min
	var digito int
	var total int

	for i >= 1 {
		digito, _ = strconv.Atoi(string(codigo[i-1]))
		total += digito + mult
		i--
		if mult > max {
			mult = min
		}
	}

	return total
}

func CalculaBaseLinhaDigitavel(codigo string, min int, max int, acima bool) int {
	i := len(codigo)
	mult := max
	var digito int
	var total int

	for i >= 1 {
		digito, _ = strconv.Atoi(string(codigo[i-1]))
		totalParcial := (digito + mult)

		if totalParcial > 9 && acima {
			valor := strconv.Itoa(totalParcial)
			digito1, _ := strconv.Atoi(string(valor[0]))
			digito2, _ := strconv.Atoi(string(valor[1]))

			total += (digito1 + digito2)
		} else {
			total += totalParcial
		}

		i--
		mult--
		if mult < min {
			mult = max
		}
	}

	return total
}

func CalculaModulo(valor, subtrair, modulo int) int {
	return (subtrair - (valor % modulo))
}

func GetConfig() models.Config {
	c := models.Config{
		DB_NAME:          "",
		DB_PORT:          "",
		DB_USER:          "",
		DB_HOST:          "",
		DB_PASS:          "",
		URL_PORT:         "",
		URL_HOST:         "",
		TRANFERENCIA_API: "",
		CONTA_API:        "",
	}

	c.DB_NAME = os.Getenv("DB_NAME")
	c.DB_PASS = os.Getenv("DB_PASS")
	c.DB_PORT = os.Getenv("DB_PORT")
	c.DB_HOST = os.Getenv("DB_HOST")
	c.DB_USER = os.Getenv("DB_USER")
	c.URL_PORT = os.Getenv("URL_PORT")
	c.URL_HOST = os.Getenv("URL_HOST")
	c.TRANFERENCIA_API = os.Getenv("TRANFERENCIA_API")
	c.CONTA_API = os.Getenv("CONTA_API")

	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		arg := strings.Split(args[i], "=")
		switch strings.ToUpper(arg[0]) {
		case "DB_NAME":
			c.DB_NAME = arg[1]
		case "DB_PASS":
			c.DB_PASS = arg[1]
		case "DB_PORT":
			c.DB_PORT = arg[1]
		case "DB_HOST":
			c.DB_HOST = arg[1]
		case "DB_USER":
			c.DB_USER = arg[1]
		case "URL_PORT":
			c.URL_PORT = arg[1]
		case "URL_HOST":
			c.URL_HOST = arg[1]
		case "TRANFERENCIA_API":
			c.TRANFERENCIA_API = arg[1]
		case "CONTA_API":
			c.CONTA_API = arg[1]
		}
	}

	if c.DB_NAME == "" {
		c.DB_NAME = "inbolso"
	}
	if c.DB_PASS == "" {
		c.DB_PASS = "12345678"
	}
	if c.DB_PORT == "" {
		c.DB_PORT = "15432"
	}
	if c.DB_HOST == "" {
		c.DB_HOST = "localhost"
	}
	if c.DB_USER == "" {
		c.DB_USER = "postgres"
	}
	if c.URL_PORT == "" {
		c.URL_PORT = "15003"
	}
	if c.URL_HOST == "" {
		c.URL_HOST = "127.0.0.1"
	}
	if c.TRANFERENCIA_API == "" {
		c.TRANFERENCIA_API = "http://127.0.0.1:15002"
	}
	if c.CONTA_API == "" {
		c.CONTA_API = "http://127.0.0.1:15001"
	}

	return c
}

//Monta uma url de conexão como a seguinte "postgres://username:password@localhost:5432/database_name"
func GetUrlConnection() string {
	conf := GetConfig()
	base := "postgres://%s:%s@%s:%s/%s"
	return fmt.Sprintf(base, conf.DB_USER, conf.DB_PASS, conf.DB_HOST, conf.DB_PORT, conf.DB_NAME)
}

func FormataLinhaDigitavel(codigoBarras string, min int, max int, somarAcima bool) models.LinhaDigitavel {

	var dv, valorBase int
	var linhaDigitavel, linhaDigitavelFormatada strings.Builder

	// 1° campo
	linhaDigitavel.WriteString(codigoBarras[0:4])
	linhaDigitavel.WriteString(codigoBarras[19:20])
	linhaDigitavel.WriteString(codigoBarras[20:24])
	valorBase = CalculaBaseLinhaDigitavel(linhaDigitavel.String(), 1, 2, true)
	dv = CalculaModulo(valorBase, 10, 10)
	linhaDigitavel.WriteString(strconv.Itoa(dv))

	// 2° campo
	linhaDigitavel.WriteString(codigoBarras[24:34])
	valorBase = CalculaBaseLinhaDigitavel(linhaDigitavel.String()[10:20], 1, 2, true)
	dv = CalculaModulo(valorBase, 10, 10)
	linhaDigitavel.WriteString(strconv.Itoa(dv))

	// 3° campo
	linhaDigitavel.WriteString(codigoBarras[34:44])
	valorBase = CalculaBaseLinhaDigitavel(linhaDigitavel.String()[21:31], 1, 2, true)
	dv = CalculaModulo(valorBase, 10, 10)
	linhaDigitavel.WriteString(strconv.Itoa(dv))

	// 4° campo
	linhaDigitavel.WriteString(codigoBarras[4:5])

	// 5° campo
	linhaDigitavel.WriteString(codigoBarras[5:9])
	linhaDigitavel.WriteString(codigoBarras[9:19])

	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[0:5] + ".")
	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[5:10] + " ")
	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[10:15] + ".")
	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[15:21] + " ")
	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[21:26] + ".")
	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[26:32] + " ")
	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[32:33] + " ")
	linhaDigitavelFormatada.WriteString(linhaDigitavel.String()[33:47])

	linha := models.LinhaDigitavel{
		LinhaDigitavel:          linhaDigitavel.String(),
		LinhaDigitavelFormatada: linhaDigitavelFormatada.String(),
	}

	return linha
}

func GeraCodigoBarras(mt models.MontaBoleto) (models.CodigoBarras, error) {
	codigo := models.CodigoBarras{
		CodigoBanco:     mt.CodigoBanco,
		Moeda:           "9",
		FatorVencimento: mt.FatorVencimento,
		Valor:           mt.Valor,
		CampoLivre:      GetCampoLivre(mt),
	}
	fmt.Println("Campo livre -> ", GetCampoLivre(mt))
	err := mt.VerificaCampos()
	if err != nil {
		return codigo, err
	}

	codigo.DV = getDV(codigo, mt)
	codigo.CodigoBarras = fmt.Sprintf("%s%s%s%s%s%s", codigo.CodigoBanco, codigo.Moeda, codigo.DV, codigo.FatorVencimento, codigo.Valor, codigo.CampoLivre)

	return codigo, nil
}

func GetLinhaDigitavel(codigoBarras string, mt models.MontaBoleto) (models.LinhaDigitavel, error) {
	var linha models.LinhaDigitavel
	if len(codigoBarras) != 44 {
		return linha, errors.New(fmt.Sprintf("Codigo de barras com tamanho incorreto %s", codigoBarras))
	}

	linha = FormataLinhaDigitavel(codigoBarras, 1, 2, true)

	return linha, nil
}

func getDV(cb models.CodigoBarras, mt models.MontaBoleto) string {
	codigo := fmt.Sprintf("%s%s%s%s%s%s", cb.CodigoBanco, cb.Moeda, cb.DV, cb.FatorVencimento, cb.Valor, cb.CampoLivre)
	total := CalculaBaseCodigoBarras(codigo, 2, 9)
	resultado := 11 - (total % 11)

	if resultado == 0 || resultado == 1 || resultado > 9 {
		return "1"
	}

	return strconv.Itoa(resultado)
}

func GetCampoLivre(mt models.MontaBoleto) string {
	return fmt.Sprintf("%s%s%s%s%d", mt.Agencia, mt.Carteira, mt.NossoNumero, mt.CodigoBeneficiario, 0)
}
