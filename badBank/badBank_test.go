package badbank

import "testing"
import th "github.com/mrenrich84/learn-go-with-tests/testHelpers"

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	th.AssertEqual(t, newBalanceFor(riya), 200)
	th.AssertEqual(t, newBalanceFor(chris), 0)
	th.AssertEqual(t, newBalanceFor(adil), 175)
}