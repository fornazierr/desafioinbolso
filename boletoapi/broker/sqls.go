package broker

func sqlNewBoleto() string {
	return "INSERT INTO public.boleto (contaorigem_id, contadestino_id, codigobanco, agencia, carteira, datavencimento, valor, nossonumero, codigobeneficiario, linhadigitavel, codigobarras) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);"
}

func sqlGetBoletoById() string {
	return "SELECT id, contaorigem_id, contadestino_id, codigobanco, agencia, carteira, datavencimento, valor, nossonumero, codigobeneficiario, linhadigitavel, codigobarras FROM public.boleto where id=$1;"
}
