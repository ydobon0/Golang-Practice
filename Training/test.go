package Training

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Account struct {
	balance         float64
	minBal          float64
	id              int
	fName           string
	lName           string
	transactionDate []string
	deposit         func(*Account, float64)
	withdraw        func(*Account, float64)
}

func deposit(acc *Account, ammount float64) {
	if ammount < 0 {
		fmt.Println("ERROR! deposit ammount must be positive")
		return
	}
	bal := acc.balance
	bal += ammount
	acc.balance = bal
	fmt.Printf("Account ID = %d, %s %s. Account balance = %.2f\n", acc.id, acc.fName, acc.lName, acc.balance)

	date := time.Now()
	acc.transactionDate = append(acc.transactionDate, date.String())
}

func withdraw(acc *Account, ammount float64) {
	if ammount < 0 {
		fmt.Println("ERROR! withdraw ammount must be positive")
		return
	}
	bal := acc.balance
	bal -= ammount
	if bal < acc.minBal {
		fmt.Printf("ERROR! balance cannot go below minimum balance of %.2f!\n", acc.minBal)
		return
	}
	acc.balance = bal
	fmt.Printf("Account ID = %d, %s %s. Account balance = %.2f\n", acc.id, acc.fName, acc.lName, acc.balance)

	date := time.Now()
	acc.transactionDate = append(acc.transactionDate, date.String())
}

func sortFirst(accs []Account) {
	sort.Slice(accs[:], func(i, j int) bool {
		return accs[i].fName < accs[j].fName
	})
}

func sortAmm(accs []Account) {
	sort.Slice(accs[:], func(i, j int) bool {
		return accs[i].balance < accs[j].balance
	})
}

func generateAccounts() func() Account {
	ID := 0
	return func() Account {
		var iBal, mBal float64
		var fN, lN string
		fmt.Println("Please enter initial balance (float)")
		fmt.Scanln(&iBal)
		fmt.Println("Please enter minimum balance (float)")
		fmt.Scanln(&mBal)
		fmt.Println("Please enter account holder's first name")
		fmt.Scanln(&fN)
		fmt.Println("Please enter account holder's last name")
		fmt.Scanln(&lN)
		var transactions []string
		var a = Account{
			iBal,
			mBal,
			ID,
			fN,
			lN,
			transactions,

			func(acc *Account, ammount float64) {
				if ammount < 0 {
					fmt.Println("ERROR! deposit ammount must be positive")
					return
				}
				bal := acc.balance
				bal += ammount
				acc.balance = bal
				fmt.Printf("Account ID = %d, %s %s. Account balance = %.2f\n", acc.id, acc.fName, acc.lName, acc.balance)

				date := time.Now()
				acc.transactionDate = append(acc.transactionDate, date.String())
			},

			func(acc *Account, ammount float64) {
				if ammount < 0 {
					fmt.Println("ERROR! withdraw ammount must be positive")
					return
				}
				bal := acc.balance
				bal -= ammount
				if bal < acc.minBal {
					fmt.Printf("ERROR! balance cannot go below minimum balance of %.2f!\n", acc.minBal)
					return
				}
				acc.balance = bal
				fmt.Printf("Account ID = %d, %s %s. Account balance = %.2f\n", acc.id, acc.fName, acc.lName, acc.balance)

				date := time.Now()
				acc.transactionDate = append(acc.transactionDate, date.String())
			},
		}
		ID += 1

		return a
	}
}

func Test() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(getEvens(arr))
	//july6()
}

func july5() {
	var accs []Account

	// ID := 0
	numAccs := 0
	fmt.Println("Please enter number of accounts")
	fmt.Scanln(&numAccs)

	f := generateAccounts()

	for ii := 0; ii < numAccs; ii++ {
		a := f()
		accs = append(accs, a)
	}
	fmt.Println(accs)
	//withdraw(&accs[0], 60)
	accs[0].withdraw(&accs[0], 60)
	//deposit(&accs[1], 200)
	accs[1].deposit(&accs[1], 200)
	//withdraw(&accs[2], 50)
	accs[2].withdraw(&accs[2], 60)

	accs[1].deposit(&accs[1], -20)

	accs[2].withdraw(&accs[2], -60)

	fmt.Println()
	fmt.Println(accs)

	sortAmm(accs[:])
	fmt.Println()
	fmt.Println(accs)

	sortFirst(accs[:])
	fmt.Println()
	fmt.Println(accs)
}

/*
1) 3 X 3 Matrix of slice type ,create a function which return a slice from function
2)Take a 3 X 3 slice and print its left triangle
Rest complete the Activities from LMS
3)input
1 2 3
4 5 6
7 8 9
output(interchange the first and last coloumn)
3 2 1
6 5 4
9 8 7
-------------------------------
input
1 2 3
4 5 6
7 8 9
output(interchange the first and last row )
7 8 9
4 5 6
1 2 3
*/

func july6() {
	//create 3 X 3 Matrix , take input from user
	//can generate any M X N Matrix including 3 X 3
	var inputs [][]int
	numR := 0
	numC := 0
	fmt.Println("Enter number of rows:")
	fmt.Scanln(&numR)
	fmt.Println("Enter number of cols:")
	fmt.Scanln(&numC)

	for ii := 0; ii < numR; ii++ {
		fmt.Printf("Row %d\n", ii+1)
		s := generateRow(numC)
		inputs = append(inputs, s)
	}
	fmt.Println(inputs)
	lowerTriangle(inputs)
	fmt.Println(exchangeR(inputs))
	fmt.Println(exchangeC(inputs))

	interfaceSlice()

}

//this function generates each row of the matrix
func generateRow(length int) []int {
	var s []int
	n := 0
	for jj := 0; jj < length; jj++ {
		fmt.Println("Enter integer")
		fmt.Scanln(&n)
		s = append(s, n)
	}
	return s
}

func lowerTriangle(in [][]int) {
	l1 := len(in)
	l2 := len(in[0])
	diag := 0
	if l1 < l2 {
		diag = l1
	} else {
		diag = l2
	}
	for ii := 0; ii < diag; ii++ {
		for jj := 0; jj <= ii; jj++ {
			fmt.Printf("%d ", in[ii][jj])
		}
		fmt.Printf("\n")
	}
}

func analyzeWords() {
	var arr [10]string
	arrLen := len(arr)
	wordLen := 0
	var word string
	for ii := 0; ii < arrLen; ii++ {
		fmt.Println("Enter word:")
		fmt.Scanln(&word)
		arr[ii] = strings.ToLower(word)
		wordLen += len(word)
	}

	fmt.Println("You entered these words")
	fmt.Println(arr)

	var avgLen int
	avgLen = wordLen / arrLen

	var s1, s2 []string
	for ii := 0; ii < arrLen; ii++ {
		if len(arr[ii]) > avgLen {
			s1 = append(s1, arr[ii])
		} else if len(arr[ii]) < avgLen {
			s2 = append(s2, arr[ii])
		}
	}

	fmt.Printf("The average word length is %d\n", avgLen)
	fmt.Println("These words are longer than average")
	fmt.Println(s1)

	fmt.Println("These words are shorter than average")
	fmt.Println(s2)
}

/*
3)input
1 2 3
4 5 6
7 8 9
output(interchange the first and last coloumn)
3 2 1
6 5 4
9 8 7
-------------------------------
input
1 2 3
4 5 6
7 8 9
output(interchange the first and last row )
7 8 9
4 5 6
1 2 3
*/

func exchangeC(in [][]int) [][]int {
	var out [][]int
	temp := 0

	for ii := 0; ii < len(in); ii++ {
		out = append(out, in[ii])
		temp = out[ii][0]
		out[ii][0] = out[ii][len(out[ii])-1]
		out[ii][len(out[ii])-1] = temp
	}

	return out
}

func exchangeR(in [][]int) [][]int {
	var out [][]int
	out = append(out, in[len(in)-1])
	for ii := 1; ii < len(in)-1; ii++ {
		out = append(out, in[ii])
	}
	out = append(out, in[0])
	return out
}

func interfaceSlice() {
	i := []interface{}{1, 2, 3}
	j := []interface{}{"A", "B", "C"}
	k := append(i, j...)
	fmt.Println(k)
}

func getEvens(in []int) []int {
	var out []int
	for _, jj := range in {
		if jj%2 == 0 {
			out = append(out, jj)
		}
	}
	return out
}

func july7() {

}
