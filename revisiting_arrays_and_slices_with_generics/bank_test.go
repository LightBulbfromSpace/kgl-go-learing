package revisiting_arrays_and_slices

import "testing"

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		artem = Account{Name: "Artem", Balance: 0}
		denis = Account{Name: "Denis", Balance: 75}

		transactions = []Transaction{
			NewTransaction(riya, artem, 50),
			NewTransaction(artem, denis, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(riya), 50)
	AssertEqual(t, newBalanceFor(artem), 25)
	AssertEqual(t, newBalanceFor(denis), 100)
}
