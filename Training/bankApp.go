package Training

import "fmt"

type account struct {
	accNum       string
	owner        entity
	balance      float64
	interestRate float64
	accType      string
}

type entity struct {
	ID         int
	address    string
	entityType string
}

type wallet struct {
	ID       string
	accounts []account
	owner    entity
}

// ----- account methods -----

//withdraw
func (a *account) withdraw(ammount float64) {
	if ammount < 0 {
		fmt.Println("You cannot withdraw a negative ammount of money")
		return
	}
	if a.balance >= ammount {
		a.balance -= ammount
		fmt.Printf("Withdrew $%.2f. The account now contains %.2f\n", ammount, a.balance)
	} else {
		fmt.Println("This account does not have enough money")
		return
	}
}

//deposit
func (a *account) deposit(ammount float64) {
	if ammount < 0 {
		fmt.Println("You cannot deposit a negative ammount of money")
		return
	}
	a.balance += ammount
	fmt.Printf("Deposited $%.2f. The account now contains %.2f\n", ammount, a.balance)
}

//apply interest
func (a *account) applyInterest() {
	if a.owner.entityType == "individual" {
		switch a.accType {
		case "checking":
			a.balance *= 1.01
		case "investment":
			a.balance *= 1.02
		case "savings":
			a.balance *= 1.05
		default:
			a.balance *= 1
		}
	} else {
		switch a.accType {
		case "checking":
			a.balance *= 1.005
		case "investment":
			a.balance *= 1.01
		case "savings":
			a.balance *= 1.02
		default:
			a.balance *= 1
		}
	}
}

//wire
func (a *account) wire(target *account, ammount float64) { //source account is a, the account calling this method
	if a.balance < ammount {
		fmt.Println("The source account's balance is too low")
		return
	}
	if ammount < 0 {
		fmt.Println("Cannot wire a negative ammount of money")
		return
	}
	a.withdraw(ammount)
	target.deposit(ammount)
}

// ----- entity methods -----

//change address
func (e *entity) changeAddress(newAddress string) {
	e.address = newAddress
	fmt.Println(e.address)
}

// ----- wallet methods -----

//display accounts
func (w *wallet) displayAccounts() {
	for _, aa := range w.accounts {
		if aa.accType == "checking" {
			fmt.Println(aa)
		}
	}
	for _, aa := range w.accounts {
		if aa.accType == "investment" {
			fmt.Println(aa)
		}
	}
	for _, aa := range w.accounts {
		if aa.accType == "savings" {
			fmt.Println(aa)
		}
	}
}

//balance
func (w *wallet) balance() {
	total := 0.0
	for _, aa := range w.accounts {
		total += aa.balance
	}
	fmt.Println(total)
}

//wire
func (w *wallet) wire(source *account, target *account, ammount float64) {
	if ammount < 0 {
		fmt.Println("Cannot wire a negative ammount of money")
		return
	}
	notFound := true
	for _, aa := range w.accounts {
		if aa.accNum == *&source.accNum {
			notFound = false
		}
	}
	if notFound {
		fmt.Println("Source account is not in this wallet")
		return
	}
	if source.balance < ammount {
		fmt.Printf("The source account does not have enough money to make this transfer")
		return
	}

	source.withdraw(ammount)
	target.deposit(ammount)

}

func BankApp() { // main function for testing code

	E1 := entity{}
	E1.ID = 0
	E1.address = "somewhere"
	E1.entityType = "individual"

	E2 := entity{}
	E2.ID = 1
	E2.address = "Somewhere"
	E2.entityType = "business"

	A1 := account{}
	A1.accNum = "0"
	A1.owner = E1
	A1.balance = 100.00
	A1.interestRate = 1.01
	A1.accType = "checking"

	A2 := account{}
	A2.accNum = "1"
	A2.owner = E1
	A2.balance = 100.00
	A2.interestRate = 1.02
	A2.accType = "investment"

	A3 := account{}
	A3.accNum = "2"
	A3.owner = E1
	A3.balance = 100.00
	A3.interestRate = 1.05
	A3.accType = "savings"

	A4 := account{}
	A4.accNum = "3"
	A4.owner = E2
	A4.balance = 100.00
	A4.interestRate = 1.01
	A4.accType = "checking"

	A5 := account{}
	A5.accNum = "4"
	A5.owner = E2
	A5.balance = 100.00
	A5.interestRate = 1.02
	A5.accType = "investment"

	A6 := account{}
	A6.accNum = "5"
	A6.owner = E2
	A6.balance = 100.00
	A6.interestRate = 1.05
	A6.accType = "savings"

	var accs1 []account
	accs1 = append(accs1, A1, A2, A3)
	W1 := wallet{}
	W1.ID = "0"
	W1.accounts = accs1
	W1.owner = E1

	var accs2 []account
	accs1 = append(accs2, A3, A4, A5)
	W2 := wallet{}
	W2.ID = "0"
	W2.accounts = accs2
	W2.owner = E2

	A1.withdraw(20)
	A1.withdraw(90)
	A1.deposit(20)
	A1.wire(&A2, 20)

	E1.changeAddress("Earth")

	W1.wire(&A1, &A2, 20)
	W1.balance()
	W1.displayAccounts()

}
