package Training

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

//todo Module 6 Activites: Data Collections -----------------------------------------------------------------------------------------------------------------------------------------

func Test() { //!--------------------------- TESTING FUNCTION ---------------------------------------
	// arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	// fmt.Println(getEvens(arr))
	//july7()
	//keywordSearch()
	// sortMap()
	july8()
}

func july8() {
	//act2()
	//act3()
	act4()
	// act5()
	// act6()
}

//todo Activity 1
func generateArray(n int) []int { //should not be able to generate 100 and -100
	var out []int
	for ii := 0; ii < n; ii++ {
		out = append(out, rand.Intn(198)-99)
	}
	return out
}

//todo Activity 2 -----------------------------------------------------------------------------------------------------------------------------------------
func act2() {
	s := generateArray(100)
	//fmt.Println(s)
	var a [100]int
	copy(a[:], s)
	fmt.Println(a)
	//fmt.Printf("%T", a)
	// fmt.Println(maxArr(a))
	// fmt.Println(maxInd(a))
	// fmt.Println(minArr(a))
	// fmt.Println(minInd(a))
	// fmt.Println(sortD(a))
	// fmt.Println(sortA(a))
	// fmt.Println(meanArr(a))
	// fmt.Println(medianArr(a))
	// fmt.Println(findPos(a))
	// fmt.Println(findNeg(a))
	fmt.Println(longestSorted(a))
	// fmt.Println((removeDupes(a)))
}

func maxArr(in [100]int) int {
	out := in[0]

	for _, ii := range in {
		if ii > out {
			out = ii
		}
	}
	return out
}

func maxInd(in [100]int) int {
	max := in[0]
	out := 0

	for ii, jj := range in {
		if jj > max {
			max = jj
			out = ii
		}
	}
	return out
}

func minArr(in [100]int) int {
	out := in[0]

	for _, ii := range in {
		if ii < out {
			out = ii
		}
	}
	return out
}

func minInd(in [100]int) int {
	min := in[0]
	out := 0

	for ii, jj := range in {
		if jj < min {
			min = jj
			out = ii
		}
	}
	return out
}

func sortD(in [100]int) [100]int {
	var out [100]int
	copy(out[:], splitArr(in[:], false))
	return out
}

func sortA(in [100]int) [100]int {
	var out [100]int
	copy(out[:], splitArr(in[:], true))
	return out
}

func splitArr(in []int, asc bool) []int {
	if len(in) < 2 {
		return in
	}
	mid := len(in) / 2
	var LL, RR []int
	LL = splitArr(in[:mid], asc)
	RR = splitArr(in[mid:], asc)
	out := mergeSort(LL, RR, asc)
	return out
}

func mergeSort(in1 []int, in2 []int, asc bool) []int {
	var out []int
	xx := 0
	yy := 0
	if asc {
		for xx < len(in1) || yy < len(in2) {
			if xx >= len(in1) {
				out = append(out, in2[yy])
				yy++
			} else if yy >= len(in2) {
				out = append(out, in1[xx])
				xx++
			} else {
				if in1[xx] < in2[yy] {
					out = append(out, in1[xx])
					xx++
				} else {
					out = append(out, in2[yy])
					yy++
				}
			}
		}
	} else {
		for xx < len(in1) || yy < len(in2) {
			if xx >= len(in1) {
				out = append(out, in2[yy])
				yy++
			} else if yy >= len(in2) {
				out = append(out, in1[xx])
				xx++
			} else {
				if in1[xx] > in2[yy] {
					out = append(out, in1[xx])
					xx++
				} else {
					out = append(out, in2[yy])
					yy++
				}
			}
		}
	}
	return out
}

func meanArr(in [100]int) int {
	mean := 0
	for _, ii := range in {
		mean += ii
	}
	mean /= len(in)
	return mean
}

func medianArr(in [100]int) int {
	sorted := sortA(in)
	mid := len(in) / 2
	return sorted[mid]
}

func findPos(in [100]int) []int {
	var out []int
	for _, ii := range in {
		if ii > 0 {
			out = append(out, ii)
		}
	}
	return out
}

func findNeg(in [100]int) []int {
	var out []int
	for _, ii := range in {
		if ii < 0 {
			out = append(out, ii)
		}
	}
	return out
}

func longestSorted(in [100]int) []int {
	asc := 0 //-1 = decreasing; 1 = increasing; 0 = same
	prev := in[0]
	start := 0
	length := 0
	bestLen := 0

	for ii, jj := range in {
		if ii != 0 {
			if asc == -1 {
				length += 1
				if jj > prev {
					if length > bestLen {
						bestLen = length
						start = ii - 1
					}
					length = 1
					asc = 1
				}
				prev = jj
			} else if asc == 1 {
				length += 1
				if jj < prev {
					if length > bestLen {
						bestLen = length
						start = ii - 1
					}
					length = 1
					asc = -1
				}
				prev = jj
			} else {
				length += 1
				if jj > prev {
					asc = 1
				} else if jj < prev {
					asc = -1
				}
				prev = jj
			}
		}
	}
	out := in[start : start+bestLen]
	return out
}

func removeDupes(in [100]int) []int {
	var mm = make(map[int]int)
	var out []int
	new := false
	for _, ii := range in {
		_, new = mm[ii]

		if !new {
			out = append(out, ii)
			mm[ii] = 1
		}
	}
	return out
}

//todo Activity 3 -----------------------------------------------------------------------------------------------------------------------------------------
func act3() {
	var nums []int
	num := 0
	done := ""
	for ii := 0; ii < 10; ii++ {
		fmt.Println("Enter an integer:")
		fmt.Scanln(&num)
		nums = append(nums, num)

		if ii >= 4 && ii < 9 {
			fmt.Println("Are you done entering numbers? type y to stop")
			fmt.Scanln(&done)
			if done == "y" {
				break
			}
		}
	}

	sum := 0
	prod := 1

	fmt.Print("You entered these numbers:")
	for _, ii := range nums {
		sum += ii
		prod *= ii
		fmt.Print(" ", ii)
	}
	fmt.Printf("\n")
	fmt.Println("The sum of these numbers is", sum)
	fmt.Println("The product of these numbers is", prod)
}

//todo Activity 4 -----------------------------------------------------------------------------------------------------------------------------------------
func act4() {
	paragraph := "Hi, I like really like cheese and crackers... They like taste good"
	fmt.Println(paragraph)
	paragraph = strings.ToLower(paragraph)
	paragraph = strings.ReplaceAll(paragraph, ".", "")
	paragraph = strings.ReplaceAll(paragraph, ",", "")
	paragraph = strings.ReplaceAll(paragraph, "?", "")
	paragraph = strings.ReplaceAll(paragraph, "!", "")
	paragraph = strings.ReplaceAll(paragraph, ":", "")
	paragraph = strings.ReplaceAll(paragraph, ";", "")
	words := strings.Split(paragraph, " ")
	fmt.Println(words)

	var mm = make(map[string]int)
	numDistinct := 0

	for _, ii := range words {
		_, found := mm[ii]
		if found {
			mm[ii] += 1
		} else {
			mm[ii] = 1
			numDistinct += 1
		}
	}

	fmt.Println("There were ", numDistinct, "distinct words")
	fmt.Println("Here are the number of times each word appeared")
	for ii := range mm {
		fmt.Println(ii, "	", mm[ii])
	}
}

//todo Activity 5 -----------------------------------------------------------------------------------------------------------------------------------------
func act5() {
	length := 0
	fmt.Println("Enter slice 1 length")
	fmt.Scanln(&length)

	ss1 := generateSlice(length)

	fmt.Println("Enter slice 2 length")
	fmt.Scanln(&length)

	ss2 := generateSlice(length)

	fmt.Println("Slice 1: ", ss1)
	fmt.Println("Slice 2: ", ss2)

	ss1 = sortSlice(ss1, true)
	ss2 = sortSlice(ss2, false)

	fmt.Println(combineSlices(ss1, ss2, true))

}

func generateSlice(length int) []int {
	var ss []int
	for ii := 0; ii < length; ii++ {
		ss = append(ss, rand.Intn(98)+1) //range does not include 0 and 100
	}
	return ss
}

func sortSlice(in []int, asc bool) []int {
	if len(in) < 2 {
		return in
	}
	mid := len(in) / 2
	var LL, RR []int
	LL = splitArr(in[:mid], asc)
	RR = splitArr(in[mid:], asc)
	out := mergeSort(LL, RR, asc)
	fmt.Println("The sorted slice is:", out)
	return out
}

func combineSlices(in1 []int, in2 []int, asc bool) []int {
	ss := append(in1, in2...)
	out := sortSlice(ss, asc)
	return out
}

//todo Activity 6 -----------------------------------------------------------------------------------------------------------------------------------------
func act6() {
	cube := Cube{}
	cube.length = 10

	box := Box{}
	box.length = 10
	box.width = 20
	box.height = 30

	sphere := Sphere{}
	sphere.radius = 2

	fmt.Println(cube.volume())
	fmt.Println(box.volume())
	fmt.Println(sphere.volume())
}

type Cube struct {
	length float64
}

type Box struct {
	length float64
	width  float64
	height float64
}

type Sphere struct {
	radius float64
}

type shape interface {
	volume() float64
}

func (a *Cube) volume() float64 {
	return a.length * a.length * a.length
}

func (a *Box) volume() float64 {
	return a.length * a.width * a.height
}

func (a *Sphere) volume() float64 {
	return 3.14 * a.radius * a.radius
}

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

//*---------------------- JULY 8 -----------------------------------------------------------------------------------------------------------------------------------------
type AccountOperations interface {
	// Methods
	computeInterest() float64
}

type SavingsAccount struct {
	number   string
	balance  float64
	interest float64
}

type CheckingAccount struct {
	number   string
	balance  float64
	interest float64
}

func (a *SavingsAccount) computeInterest() float64 {
	return 0.005
}

func (a *CheckingAccount) computeInterest() float64 {
	return 0.001
}

func CheckType(i interface{}) {
	switch i.(type) {
	case *SavingsAccount:
		fmt.Println("This is a saving account")
	case *CheckingAccount:
		fmt.Println("This is a checking account")
	default:
		fmt.Println("Unknown account")
	}
}

func testInterface() {
	a := SavingsAccount{}
	a.number = "S21345345345355"
	a.balance = 159

	var ao1 AccountOperations
	ao1 = &a
	fmt.Println("Result for ao1")
	CheckType(ao1)

	b := CheckingAccount{}
	b.number = "C218678678345345355"
	b.balance = 2000

	var ao2 AccountOperations
	ao2 = &b
	fmt.Println("Result for ao2")
	CheckType(ao2)
}

//todo Bank Application

//todo Black Jack

//todo Goblin Tower
