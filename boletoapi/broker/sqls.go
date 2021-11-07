package broker

func sqlNewBoleto() string {
	return "INSERT INTO public.boleto (contaorigem_id, contadestino_id, codigobanco, agencia, carteira, datavencimento, valor, nossonumero, codigobeneficiario, linhadigitavel, codigobarras) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);"
}

func sqlGetAllBoleto() string {
	return "SELECT id, contaorigem_id, contadestino_id, codigobanco, agencia, carteira, datavencimento, valor, nossonumero, codigobeneficiario, linhadigitavel, codigobarras FROM public.boleto order by created desc;"
}

//Realiza a selecao filtrando pelo seu ID
func sqlGetBoletoById() string {
	return "SELECT id, contaorigem_id, contadestino_id, codigobanco, agencia, carteira, datavencimento, valor, nossonumero, codigobeneficiario, linhadigitavel, codigobarras FROM public.boleto where id=$1 order by created desc;"
}

//Realiza a seleção filtrando por codigo banco, agencia, carteira
func sqlGetBoletoByDados1() string {
	return "SELECT id, contaorigem_id, contadestino_id, codigobanco, agencia, carteira, datavencimento, valor, nossonumero, codigobeneficiario, linhadigitavel, codigobarras FROM public.boleto where codigobanco=$1 and agencia=$2 and carteira=$3 order by created desc;"
}

//Realiza a selecao filtrando pela conta de origem
func sqlGetBoletoByContaOrigem() string {
	return "SELECT id, contaorigem_id, contadestino_id, codigobanco, agencia, carteira, datavencimento, valor, nossonumero, codigobeneficiario, linhadigitavel, codigobarras FROM public.boleto where contaorigem_id=$1 order by created desc;"
}
