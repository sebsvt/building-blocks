package transaction

type Transaction struct {
	reference  string
	merchantID string
	customer   Person
	payment    Payment
}

func CreateNewTransaction(ref, merchantID string, customer Person) (Transaction, error) {
	return Transaction{
		reference:  ref,
		merchantID: merchantID,
		customer:   customer,
		payment:    Payment{},
	}, nil
}
