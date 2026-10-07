package sum

type Person struct {
	Name string
}

type Account struct {
	Name    string
	Balance float64
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func Sum(nums []int) int {
	sum := func(acc, x int) int { return acc + x }
	return Reduce(nums, sum, 0)
}

func SumAllTails(nums ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) != 0 {
			return append(acc, Sum(x[1:]))
		} else {
			return append(acc, 0)
		}
	}
	return Reduce(nums, sumTail, []int{})
}

func Reduce[A, B any](collection []B, fn func(A, B) A, initialValue A) A {
	var result A = initialValue
	for _, value := range collection {
		result = fn(result, value)
	}
	return result
}

func BalanceFor(transactions []Transaction, name string) float64 {
	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			return currentBalance - t.Sum
		}
		if t.To == name {
			return currentBalance + t.Sum
		}
		return currentBalance
	}

	return Reduce(transactions, adjustBalance, 0)
}

func NewBalanceFor(account *Account, transactions []Transaction) *Account {
	applyTransaction := func(account *Account, transaction Transaction) *Account {
		if (*account).Name == transaction.From {
			(*account).Balance -= transaction.Sum
		}
		if (*account).Name == transaction.To {
			(*account).Balance += transaction.Sum
		}
		return account
	}

	return Reduce(transactions, applyTransaction, account)
}

func Find[A any](collection []A, fn func(A) bool) (A, bool) {
	var empty A
	for _, item := range collection {
		if fn(item) {
			return item, true
		}
	}

	return empty, false
}
