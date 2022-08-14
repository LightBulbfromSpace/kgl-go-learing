package revisiting_arrays_and_slices

type Account struct {
	Name    string
	Balance float64
}

type Transaction struct {
	From   string
	To     string
	Amount float64
}

func NewTransaction(from, to Account, amount float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Amount: amount}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

func applyTransaction(a Account, t Transaction) Account {
	if t.From == a.Name {
		a.Balance -= t.Amount
	}
	if t.To == a.Name {
		a.Balance += t.Amount
	}
	return a
}

func BalanceFor(transactions []Transaction, name string) float64 {
	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			return currentBalance - t.Amount
		}
		if t.To == name {
			return currentBalance + t.Amount
		}
		return currentBalance
	}
	return Reduce(transactions, adjustBalance, 0.0)
}
