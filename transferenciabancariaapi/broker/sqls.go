package broker

func sqlGetSaldo() string {
	return "select conta_id, titular_id, saldo from public.saldo where conta_id=$1 and titular_id=$2"
}
