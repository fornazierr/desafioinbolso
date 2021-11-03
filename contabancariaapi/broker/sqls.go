package broker

func sqlTitular() string {
	return "SELECT id, nome, cpf, email, nascimento::text, nomepai, nomemae, cidade, estado FROM public.titular;"
}

func sqlTitularById() string {
	return "SELECT id, nome, cpf, email, nascimento::text, nomepai, nomemae, cidade, estado FROM public.titular where id = $1;"
}

func sqlNewTitular() string {
	return "INSERT INTO public.titular (nome, cpf, email, nascimento, nomepai, nomemae, cidade, estado) VALUES($1, $2, $3, $4, $5, $6, $7, $8);"
}

func sqlUpdateTitular() string {
	return "UPDATE public.titular SET nome=$1, cpf=$2, email=$3, nascimento=$4, nomepai=$5, nomemae=$6, cidade=$7, estado=$8 WHERE id=$9;"
}

func sqlDeleteTitular() string {
	return "DELETE FROM public.titular WHERE id=$1;"
}

func slqGetContaBancaria() string {
	return "SELECT id, codigobanco, agencia, conta, digito, titular_id FROM public.contabancaria;"
}

func slqGetContaBancariaById() string {
	return "SELECT id, codigobanco, agencia, conta, digito, titular_id FROM public.contabancaria where id=$1;"
}

func sqlNewContaBancaria() string {
	return "INSERT INTO public.contabancaria (codigobanco, agencia, conta, digito, titular_id) VALUES($1, $2, $3, $4, $5);"
}

func sqlDeleteContaBancaria() string {
	return "DELETE FROM public.contabancaria WHERE id=$1;"
}
