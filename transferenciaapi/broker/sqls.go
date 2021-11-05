package broker

func sqlGetSaldo() string {
	return "select conta_id, titular_id, saldo from public.saldo where conta_id=$1 and titular_id=$2"
}

func sqlNewSaldo() string {
	return "INSERT INTO public.saldo (titular_id, conta_id, saldo, created, updated) VALUES($1, $2, 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);"
}

func sqlUpdateSaldo() string {
	return "UPDATE public.saldo SET saldo=$1, updated=CURRENT_TIMESTAMP where titular_id=$2 AND conta_id=$3;"
}

func sqlGetRegistrosSaldo() string {
	return "select titular_id, conta_id, sinal, valor from public.registrosaldo where titular_id=$1 and conta_id=$2"
}

func sqlRegistroSaldo() string {
	return "insert into public.registrosaldo (titular_id, conta_id, sinal, valor) values ($1,$2,$3,$4);"
}

func sqlNewTranferencia() string {
	return "INSERT INTO public.transferencia (contaorigem_id, contadestino_id, valor, created) VALUES($1, $2, $3, CURRENT_TIMESTAMP);"
}
