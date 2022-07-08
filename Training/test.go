package Training

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

//*---------------------- JULY 5 -----------------------------------------------------------------------------------------------------------------------------------------

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

//*---------------------- JULY 6 -----------------------------------------------------------------------------------------------------------------------------------------

/*
todo 1) 3 X 3 Matrix of slice type ,create a function which return a slice from function
Rest complete the Activities from LMS
*/

func july6() {
	//todo create 3 X 3 Matrix , take input from user
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

//todo 2)Take a 3 X 3 slice and print its left triangle
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

//todo LMS slice activity
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
todo	3)input
todo	1 2 3
todo	4 5 6
todo	7 8 9
todo	output(interchange the first and last coloumn)
todo	3 2 1
todo	6 5 4
todo	9 8 7
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

/*
todo	input
todo	1 2 3
todo	4 5 6
todo	7 8 9
todo	output(interchange the first and last row )
todo	7 8 9
todo	4 5 6
todo	1 2 3
*/

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

//todo take in a slice of integers and return a slice that contains only the even numbers
func getEvens(in []int) []int {
	var out []int
	for _, jj := range in {
		if jj%2 == 0 {
			out = append(out, jj)
		}
	}
	return out
}

//*---------------------- JULY 7 -----------------------------------------------------------------------------------------------------------------------------------------

func july7() {
	// var greetings = make(map[string]string)
	// greetings["Abe"] = "Hi"
	// greetings["Bob"] = "Hey"
	// greetings["Clyde"] = "Hello"
	// greetings["Dan"] = "Sup"
	// greetings["Frank"] = "Good Day"
	// fmt.Println(greetings)

	// keys := make([]string, 0, len(greetings))
	// for i := range greetings {
	// 	keys = append(keys, i)
	// }
	// sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	// fmt.Println(keys)

	//todo 1)Create a map where values must be a structure type
	var accs = make(map[int]Account)
	numAccs := 0
	fmt.Println("Please enter number of accounts")
	fmt.Scanln(&numAccs)

	f := generateAccounts() //uses account structure from 7/5/2022

	for ii := 0; ii < numAccs; ii++ {
		a := f()
		accs[a.id] = a
	}

	fmt.Println(accs)
}

/*
golang is very good for APIs
maps and structures are very important

methods are faster than functions
methods are functions that use structs

interface - a collection of function but with only the function signature
only prototype of function is there.
if you declare a function from an interface, it becomes a method
*/

//todo do basic map activities on LMS

//todo Maps Activity: Keyword Search
func keywordSearch() {
	var mm = make(map[string]string)
	key := ""

	mm["apple"] = "fruit"
	mm["bananna"] = "fruit"
	mm["strawberry"] = "fruit"
	mm["watermelon"] = "fruit"
	mm["orange"] = "fruit"
	mm["onion"] = "vegetable"
	mm["celery"] = "vegetable"
	mm["peanut"] = "nut"
	mm["pistachio"] = "nut"
	mm["almond"] = "nut"

	done := ""
	found := false
	for true {
		found = false
		fmt.Println("Please enter a search term:")
		fmt.Scanln(&key)

		for i := range mm {
			if i == key || mm[i] == key {
				fmt.Println(i, " : ", mm[i])
				found = true
			}
		}

		if !found {
			fmt.Println("The term you entered does not appear in the map. Sorry!")
		}

		fmt.Println("Do you wish to stop? Type y to stop. Type anything else to continue")
		fmt.Scanln(&done)

		if done == "y" {
			break
		}
	}
	fmt.Println("All done!")
}

//todo	3) Create a map ex map[string]int{"orange": 5, "apple": 7,	"mango": 3, "strawberry": 9} ,sort the map based on key length in asecending order
func sortMap() {
	var mm = make(map[string]int)

	mm["orange"] = 5
	mm["apple"] = 7
	mm["mango"] = 3
	mm["strawberry"] = 9

	var keys []string
	for ii := range mm {

		keys = append(keys, ii)
	}
	fmt.Println(keys)

	sort.SliceStable(keys, func(ii, jj int) bool {
		return len(keys[ii]) < len(keys[jj])
	})

	for _, ii := range keys {
		fmt.Println(ii, "	:	", mm[ii])
	}

}

//todo	2) Take 20 (any count) from console between 1 -100 ) then print the summary like no between 1- 10 Count-5     11 -20 count -7    21-30 count -10  etc
func rangeSummary() {
	num := 0
	fmt.Println("How many numbers do you want to enter?")
	fmt.Scanln(&num)

	var mm = make(map[int]int)

	mm[0] = 0
	mm[1] = 0
	mm[2] = 0
	mm[3] = 0
	mm[4] = 0
	mm[5] = 0
	mm[6] = 0
	mm[7] = 0
	mm[8] = 0
	mm[9] = 0

	input := 0

	for ii := 0; ii < num; ii++ {
		fmt.Println("Enter an integer:")
		fmt.Scanln(&input)
		if input <= 10 {
			mm[0] += 1
		} else if input <= 20 {
			mm[1] += 1
		} else if input <= 30 {
			mm[2] += 1
		} else if input <= 40 {
			mm[3] += 1
		} else if input <= 50 {
			mm[4] += 1
		} else if input <= 60 {
			mm[5] += 1
		} else if input <= 70 {
			mm[6] += 1
		} else if input <= 80 {
			mm[7] += 1
		} else if input <= 90 {
			mm[8] += 1
		} else {
			mm[9] += 1
		}
	}

	for ii := 0; ii < 10; ii++ {
		switch ii {
		case 0:
			fmt.Println(mm[ii], "numbers less than or equal to 10")
		case 9:
			fmt.Println(mm[ii], "numbers greater than 90")
		default:
			fmt.Println(mm[ii], "numbers between", (ii*10)+1, "and", (ii*10)+10)
		}
	}
}

func Test() { //!--------------------------- TESTING FUNCTION ---------------------------------------
	// arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	// fmt.Println(getEvens(arr))
	//july7()
	//keywordSearch()
	// sortMap()
	rangeSummary()
}
