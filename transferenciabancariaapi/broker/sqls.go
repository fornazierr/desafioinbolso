package broker

func sqlGetSaldo() string {
	return "select conta_id, titular_id, saldo from public.saldo where conta_id=$1 and titular_id=$2"
}

func sqlGetRegistrosSaldo() string {
	return "select titular_id, conta_id, sinal, valor from public.registrosaldo where titular_id=$1 and conta_id=$1"
}

func sqlRegistraSaldo() string {
	return "insert into public.registrosaldo (titular_id, conta_id, sinal, valor) values ($1,$2,$3,$4);"
}

func sqlUpdateSaldo() string {
	// UPDATE public.tipos SET descricao='', metadado='', "status"=false, datacriacao=current_timestamp(), dataalteracao=current_timestamp() WHERE codigo='';
	return "update public.saldo set saldo=$1 where titular_id=$2 and conta_id=$3"
}
