package Training

import "fmt"

type account struct {
	accNum       int //changed from string to int for the challenge
	owner        *entity
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
	ID       int //changed from string to int for the challenge
	accounts []*account
	owner    *entity
}

// ----- account methods -----

//withdraw
func (a *account) withdraw(ammount float64) {
	if ammount < 0 { //ensure we do not withdraw a negative ammount of money
		fmt.Println("You cannot withdraw a negative ammount of money")
		return
	}
	if a.balance >= ammount {
		a.balance -= ammount
		fmt.Printf("Withdrew $%.2f. The account now contains %.2f\n", ammount, a.balance)
	} else { //ensure the account has enough money to make the transfer
		fmt.Println("This account does not have enough money")
		return
	}
}

//deposit - deposit money into account
func (a *account) deposit(ammount float64) {
	if ammount < 0 { //ensure we do not deposit a negative ammount of money
		fmt.Println("You cannot deposit a negative ammount of money")
		return
	}
	a.balance += ammount
	fmt.Printf("Deposited $%.2f. The account now contains %.2f\n", ammount, a.balance)
}

//apply interest - apply interest according to the account type and owner type
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

//wire - transfer money from this account to another account
func (a *account) wire(target *account, ammount float64) { //source account is a, the account calling this method
	if a.balance < ammount { //ensure the account has enough money to make the transfer
		fmt.Println("The source account's balance is too low")
		return
	}
	if ammount < 0 { //ensure we do not transfer a negative ammount of money
		fmt.Println("Cannot wire a negative ammount of money")
		return
	}
	a.withdraw(ammount)
	target.deposit(ammount)
}

// ----- entity methods -----

//change address - change the address of the owner
func (e *entity) changeAddress(newAddress string) {
	e.address = newAddress
	fmt.Println(e.address)
}

// ----- wallet methods -----

//display accounts - display all accounts in the wallet. Display them in the order: checking, investment, then savings
func (w *wallet) displayAccounts() {
	for _, aa := range w.accounts {
		if aa.accType == "checking" {
			fmt.Println(aa.accNum, *aa.owner, aa.balance, aa.interestRate, aa.accType)
		}
	}
	for _, aa := range w.accounts {
		if aa.accType == "investment" {
			fmt.Println(aa.accNum, *aa.owner, aa.balance, aa.interestRate, aa.accType)
		}
	}
	for _, aa := range w.accounts {
		if aa.accType == "savings" {
			fmt.Println(aa.accNum, *aa.owner, aa.balance, aa.interestRate, aa.accType)
		}
	}
}

//balance - display the total balance of the wallet
func (w *wallet) balance() {
	total := 0.0
	for _, aa := range w.accounts {
		total += aa.balance
	}
	fmt.Println(total)
}

//wallet wire - transfer money from one account to another account. the source account must be in the wallet
func (w *wallet) wire(source *account, target *account, ammount float64) {
	if ammount < 0 { //make sure the ammount transfered is not negative
		fmt.Println("Cannot wire a negative ammount of money")
		return
	}

	//check if the source account is in the wallet
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

	//check to make sure the account has enough money to make the transfer
	if source.balance < ammount {
		fmt.Println("The source account does not have enough money to make this transfer")
		for ii := 0; ii < len(w.accounts); ii++ {
			if w.accounts[ii].balance >= ammount {
				fmt.Println("However, account ", w.accounts[ii].accNum, " does have enough to make the transfer")
				break
			}
		}
		return
	}

	source.withdraw(ammount)
	target.deposit(ammount)

}

func BankApp() { // main function for testing code
	//* --------------------------------------------- setting up for the terminal ---------------------------------------------
	E1 := entity{}
	E1.ID = 0
	E1.address = "somewhere"
	E1.entityType = "individual"

	E2 := entity{}
	E2.ID = 1
	E2.address = "Somewhere"
	E2.entityType = "business"

	A1 := account{}
	A1.accNum = 0
	A1.owner = &E1
	A1.balance = 100.00
	A1.interestRate = 1.01
	A1.accType = "checking"

	A2 := account{}
	A2.accNum = 1
	A2.owner = &E1
	A2.balance = 100.00
	A2.interestRate = 1.02
	A2.accType = "investment"

	A3 := account{}
	A3.accNum = 2
	A3.owner = &E1
	A3.balance = 100.00
	A3.interestRate = 1.05
	A3.accType = "savings"

	A4 := account{}
	A4.accNum = 3
	A4.owner = &E2
	A4.balance = 100.00
	A4.interestRate = 1.01
	A4.accType = "checking"

	A5 := account{}
	A5.accNum = 4
	A5.owner = &E2
	A5.balance = 100.00
	A5.interestRate = 1.02
	A5.accType = "investment"

	A6 := account{}
	A6.accNum = 5
	A6.owner = &E2
	A6.balance = 100.00
	A6.interestRate = 1.05
	A6.accType = "savings"

	var accs1 []*account
	accs1 = append(accs1, &A1, &A2, &A3)
	W1 := wallet{}
	W1.ID = 0
	W1.accounts = accs1
	W1.owner = &E1

	var accs2 []*account
	accs2 = append(accs2, &A4, &A5, &A6)
	W2 := wallet{}
	W2.ID = 1
	W2.accounts = accs2
	W2.owner = &E2

	//* These variables below will be used in the terminal
	allAccs := append(accs1, accs2...)
	var allWall []wallet = []wallet{W1, W2}
	var allEnts []*entity = []*entity{&E1, &E2}

	done := false
	fmt.Println("Welcome to the banking terminal!")
	input := 0 //stores user inputs
	input2 := 0
	input3 := 0
	xx := 0
	ammount := 0.0 //stores user inputs for deposits, withdraws, and wiring

	//* --------------------------------------------- start of bank terminal ---------------------------------------------
	for true {
		fmt.Println("What would you like to do? Enter one of the numbers below")
		fmt.Println("1. View an account")
		fmt.Println("2. View a wallet")
		fmt.Println("3. Change user address")
		fmt.Println("4. Quit")
		fmt.Scanln(&input)
		fmt.Println()
		switch input {
		case 1:
			input = 0
			fmt.Println("Enter an account number")
			fmt.Scanln(&input)
			fmt.Println()
			if input >= len(allAccs) {
				fmt.Println("That is not a valid account number. Please try again")
				input = 0
			} else {
				xx = input
				for true {
					fmt.Println(allAccs[xx].accNum)
					fmt.Println(*allAccs[xx].owner)
					fmt.Println(allAccs[xx].balance)
					fmt.Println(allAccs[xx].accType)
					fmt.Println()

					input = 0
					fmt.Println("What do you want to do with this account?")
					fmt.Println("1. Deposit money")
					fmt.Println("2. Withdraw money")
					fmt.Println("3. Wire money")
					fmt.Println("4. Apply interest")
					fmt.Println("5. Back")
					fmt.Scanln(&input)
					fmt.Println()

					switch input {
					case 1:
						fmt.Println("How much money?")
						fmt.Scanln(&ammount)
						allAccs[xx].deposit(ammount)
					case 2:
						fmt.Println("How much money?")
						fmt.Scanln(&ammount)
						allAccs[xx].withdraw(ammount)
					case 3:
						fmt.Println("Enter target account ID")
						fmt.Scanln(&input2)
						if input2 >= len(allAccs) {
							fmt.Println("That is not a valid account ID. Please try again")
						} else {
							fmt.Println("How much money?")
							fmt.Scanln(&ammount)
							allAccs[xx].wire(allAccs[input2], ammount)
						}
					case 4:
						allAccs[xx].applyInterest()
					case 5:
						done = true
					default:
						fmt.Println("Invalid input! Try again.")
					}

					if done {
						done = false
						break
					}

					fmt.Println()
				}
			}
		case 2:
			input = 0
			fmt.Println("Enter a wallet ID number")
			fmt.Scanln(&input)
			fmt.Println()
			if input >= len(allWall) {
				fmt.Println("That is not a valid wallet ID. Please try again")
				input = 0
			} else {
				xx = input
				for true {
					fmt.Println(allWall[xx].ID)
					fmt.Println(*allWall[xx].owner)
					fmt.Println()

					input = 0
					fmt.Println("What do you want to do with this wallet?")
					fmt.Println("1. View accounts")
					fmt.Println("2. View balance")
					fmt.Println("3. Wire money")
					fmt.Println("4. Back")
					fmt.Scanln(&input)
					fmt.Println()

					switch input {
					case 1:
						allWall[xx].displayAccounts()
					case 2:
						fmt.Print("Total balance across all accounts: ")
						allWall[xx].balance()
					case 3:
						fmt.Println("Enter source account ID")
						fmt.Scanln(&input2)
						fmt.Println("Enter target account ID")
						fmt.Scanln(&input3)
						if input2 >= len(allAccs) || input3 >= len(allAccs) {
							fmt.Println("At least one of the values you entered was not a valid account ID. Please try again")
						} else {
							fmt.Println("How much money?")
							fmt.Scanln(&ammount)
							allWall[xx].wire(allAccs[input2], allAccs[input3], ammount)
						}
					case 4:
						done = true
					default:
						fmt.Println("Invalid input! Try again.")
					}

					if done {
						done = false
						break
					}

					fmt.Println()
				}
			}
		case 3:
			fmt.Println("Enter user ID number")
			fmt.Scanln(&input)
			fmt.Println()
			if input >= len(allEnts) {
				fmt.Println("That is not a valid user ID. Try again")
			} else {
				xx = input
				fmt.Println(allEnts[xx].ID)
				fmt.Println(allEnts[xx].address)
				fmt.Println(allEnts[xx].entityType)
				fmt.Println()

				fmt.Println("Do you want to change your address?")
				fmt.Println("1. Yes")
				fmt.Println("Any other number: No")
				fmt.Scanln(&input)

				if input == 1 {
					fmt.Println("Enter new address (Do not include spaces)")
					newAddr := ""
					fmt.Scanln(&newAddr)
					allEnts[xx].changeAddress(newAddr)
				}
			}
		case 4:
			done = true
		default:
			fmt.Println("Invalid input! Try again.")
		}
		if done {
			break
		}
		fmt.Println()
	}

	// A1.withdraw(20)
	// A1.withdraw(90)
	// A1.deposit(20)
	// A1.wire(&A2, 20)

	// E1.changeAddress("Earth")

	// W1.wire(&A1, &A2, 20)
	// W1.balance()
	// W1.displayAccounts()

	// W2.displayAccounts()
}
