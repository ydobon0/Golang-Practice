package main

import "fmt"

//"math"
//"strconv"
//"unicode"

//Alan Xu
//keep all practice in a go file will select from them later
type Node struct {
	Value int
}

type Stack struct {
	nodes []*Node
	count int
}

func NewStack() *Stack {
	return &Stack{}
}
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)

	s.count++ //pont to next position
}
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}
func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

func Sum(sli []int) int {
	sum := 0
	for ii := 0; ii < len(sli); ii++ {
		sum += sli[ii]
	}
	return sum
}

func half(in *int) bool {
	if *in%2 == 0 {
		*in /= 2
		return true
	} else {
		*in /= 2
		return false
	}
}

func findMax(in ...int) int {
	max := in[0]
	for ii, num := range in {
		if ii == 0 {
			max = num
		} else {
			if num > max {
				max = num
			}
		}
	}
	return max
}

func makeOddGenerator() func() int {
	odd := -1
	return func() int {
		odd += 2
		return odd
	}
}
func makeEvenGenerator() func() int {
	even := -2 //I'm counting 0 as an even number which is why I start at -2 instead of 0 for the even generator
	return func() int {
		even += 2
		return even
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func incArr(in *[]int) {
	arr := *in
	for ii := 0; ii < len(*in); ii++ {
		arr[ii] += 1
	}
}

func incArr2(in []int) {
	for ii := 0; ii < len(in); ii++ {
		in[ii] += 1
	}
}

func sum(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		s += i
	}
	return s
}

func multi(x, y, z int) (int, int, int) {
	x += 1
	y += 1
	z += 1
	return x, y, z
}

func fibonacii() func() int {
	first, second := 0, 1
	return func() int {
		temp := first
		first, second = second, first+second
		return temp
	}
}

func nextPrime() func() int {
	prev := 1
	return func() int {
		next := prev + 1
		found := false
		for found == false {
			found = true
			for ii := 2; ii < next; ii++ {
				if next%ii == 0 {
					found = false
				}
			}
			if found == false {
				next++
			}
		}
		prev = next
		return next
	}
}

func coffeeShop() {
	var coffeeSize, coffeeType, coffeeFlavor string
	price := 0.0
	fmt.Printf("Do you want small, medium, or large? ")
	fmt.Scanln(&coffeeSize)
	switch coffeeSize {
	case "small":
		price += 2
	case "medium":
		price += 3
	case "large":
		price += 4
	default:
		price += 2
		coffeeSize = "small"
	}

	fmt.Printf("\nDo you want brewed, espresso, or cold brew? ")
	fmt.Scanln(&coffeeType)
	fmt.Println(coffeeType)
	switch coffeeType {
	case "espresso":
		price += 0.5
	case "cold":
		price += 1
		coffeeType = "cold brew"
	case "cold brew":
		price += 1
	default:
		coffeeSize = "brewed"
	}

	fmt.Printf("\nDo you want a flavored syrup? (yes or no) ")
	fmt.Scanln(&coffeeFlavor)
	fmt.Println(coffeeFlavor)
	if coffeeFlavor == "yes" {
		fmt.Printf("\nDo you want hazelnut, vanilla, or caramel? ")
		fmt.Scanln(&coffeeFlavor)
		switch coffeeFlavor {
		case "hazelnut":
			price += 0.5
			coffeeFlavor = " with hazelnut syrup"
		case "vanilla":
			price += 0.5
			coffeeFlavor = " with vanilla syrup"
		case "caramel":
			price += 0.5
			coffeeFlavor = " with caramel syrup"
		default:
			coffeeFlavor = ""
		}
	} else {
		coffeeFlavor = ""
	}
	fmt.Printf("\nYou asked for a %s cup of %s coffee%s", coffeeSize, coffeeType, coffeeFlavor)
	fmt.Printf("\nYour cup of coffee costs %.2f", price)
	fmt.Printf("\nThe price with a tip is %.2f", price*1.15)
}

func fizzBuzz() {
	count := 0
	fmt.Printf("How many fizzing and buzzing units do you need in your life? ")
	fmt.Scanln(&count)
	ii := 0
	for count > 0 {
		if ii > 0 && ii%15 == 0 {
			fmt.Printf("\nfizz buzz")
			count--
		} else if ii > 0 && ii%5 == 0 {
			fmt.Printf("\nbuzz")
			count--
		} else if ii > 0 && ii%3 == 0 {
			fmt.Printf("\nfizz")
			count--
		} else {
			fmt.Printf("\n%d", ii)
		}
		ii++
	}
	fmt.Printf("\nTRADITION!!")
}

func revArray() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	var empty [len(arr)]int
	arrSize := len(arr)
	for ii, val := range arr {
		empty[arrSize-ii-1] = val
	}
	fmt.Println(empty)
}

func shiftArray(pos int, direction string) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	shift := pos

	if direction == "right" {
		shift = len(arr) - shift
	}

	arr = append(arr[shift:], arr[0:shift]...)

	fmt.Println(arr)
}

func rowColSum() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rowSum, colSum := 0, 0
	for nn := 0; nn < len(matrix); nn++ {
		rowSum, colSum = 0, 0
		for xx := 0; xx < len(matrix); xx++ {
			rowSum += matrix[nn][xx]
			colSum += matrix[xx][nn]
		}
		fmt.Printf("The sum of row %d is %d\n", nn+1, rowSum)
		fmt.Printf("The sum of column %d is %d\n", nn+1, colSum)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	incArr2(nums)
	fmt.Println(nums)
	s := NewStack()
	s.Push(&Node{10})
	s.Push(&Node{12})
	s.Push(&Node{14})
	s.Push(&Node{16})

	fmt.Println(s.Pop(), s.Pop(), s.Pop(), s.Pop())
	//fizzBuzz()
	//coffeeShop()
	//rowColSum()
	//f := nextPrime()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	/*
				//activity 1
				fmt.Println((5 + 3) % 2 * 9)

				//activity 2
				fmt.Println("Enter Float")
				var ff float32
				fmt.Scanln(&ff)
				fmt.Println(int32(ff))

				//activity 3
				var V, P, r float32
				var n, t int32

				fmt.Println("Enter P (float)")
				fmt.Scanln(&P)
				fmt.Println("Enter r (float)")
				fmt.Scanln(&r)
				fmt.Println("Enter n (int)")
				fmt.Scanln(&n)
				fmt.Println("Enter t (int)")
				fmt.Scanln(&t)

				xx := 1 + r/float32(n)
				yy := n * t
				zz := int32(xx) ^ yy
				V = P * float32(zz)

				fmt.Printf("Value = %f \n", V)
				fmt.Printf("Initial Deposit = %f \n", P)
				fmt.Printf("Interest Rate = %f \n", r)
				fmt.Printf("Interest is calculated %d times per year \n", n)
				fmt.Printf("It has been %d years since the initial deposit \n", t)

				//activity 4
				var interest, principal, rate float32
				var days int32
				fmt.Println("Enter principal (float)")
				fmt.Scanln(&principal)
				fmt.Println("Enter rate (float)")
				fmt.Scanln(&rate)
				fmt.Println("Enter number of days (int)")
				fmt.Scanln(&days)

				interest = principal * rate * float32(days) / 365
				fmt.Println(interest)

				//activity 5
				fmt.Println("a = 0, b = 1")
				fmt.Println("a < b = True")
				fmt.Println("a <= b = True")
				fmt.Println("a != b = True")
				fmt.Println("a > b = False")
				fmt.Println("a = b = False")
				fmt.Println("a >= b = False")
				//activity 6
				var input int
				fmt.Println("Enter number (int)")
				fmt.Scanln(&input)

				fmt.Printf("You selected %d\n", input)
				var B bool
				if input != 0 {
					B = true
				}
				fmt.Printf("The boolean of your number is %t\n", B)
				fmt.Print("The binary equivalent of your number is ")
				fmt.Println(strconv.FormatInt(int64(input), 2))
				fmt.Printf("The square root of your number is %f\n", math.Sqrt(float64(input)))

				//activity 7
				var num1, num2, num3, num4, num5 int
				fmt.Printf("Enter num1\n")
				fmt.Scanln(&num1)
				fmt.Printf("Enter num2\n")
				fmt.Scanln(&num2)
				fmt.Printf("Enter num3\n")
				fmt.Scanln(&num3)
				fmt.Printf("Enter num4\n")
				fmt.Scanln(&num4)
				fmt.Printf("Enter num5\n")
				fmt.Scanln(&num5)

				fmt.Printf("You entered: %d, %d, %d, %d, and %d\n", num1, num2, num3, num4, num5)
				fmt.Printf("The product of the numbers is %d\n", num1*num2*num3*num4*num5)
				fmt.Printf("The average of the numbers is %d\n", (num1+num2+num3+num4+num5)/5)
				fmt.Printf("The sum of the numbers is %d\n", num1+num2+num3+num4+num5)

				//practice exercise

					var X, Y, Z int
					fmt.Printf("Enter x\n")
					fmt.Scanln(&X)
					fmt.Printf("Enter y\n")
					fmt.Scanln(&Y)
					fmt.Printf("Enter z\n")
					fmt.Scanln(&Z)

					if x > y && y > z {
						fmt.Println("X > Y > Z")
					}
					if x > z && z > y {
						fmt.Println("X > Z > Y")
					}
					if y > x && x > z {
						fmt.Println("Y > X > Z")
					}
					if y > z && z > x {
						fmt.Println("Y > Z > X")
					}
					if z > x && x > y {
						fmt.Println("Z > X > Y")
					}
					if z > y && y > x {
						fmt.Println("Z > Y > X")
					}
					fmt.Println(x + y + z)

				//Practice 2: Multiple Conditions
				var user, pass string
				fmt.Printf("Enter user\n")
				fmt.Scanln(&user)
				fmt.Printf("Enter pass\n")
				fmt.Scanln(&pass)
				if user == "hi" && pass == "hello" {
					fmt.Println("Access Granted")
				}
				if user != "hi" || pass != "hello" {
					fmt.Println("Access Denied")
				}

				//Practice 2: Else If


				//Write a golang code to convert centimeter into inches and feet ,take centimeter as input ex- input is 25 centimeter than it should first convert it into inch.
				var centimeter float32
				var inch, feet int
				fmt.Println("Enter Centimeters")
				fmt.Scanln(&centimeter)

				inch = int(centimeter / 2.54)
				feet = inch / 12
				inch = inch % 12

				fmt.Printf("%f centimeters is about %d feet and %d inches\n", centimeter, feet, inch)

				//Electricity bill from 0 to 100 unit rate $5 per unit from 101 to 200 unit $7 per unit for 201 to 350 unit $10 per unit calculate total bill
				var units, bill int
				fmt.Println("Enter Units")
				fmt.Scanln(&units)

				if units >= 0 && units <= 100 {
					bill = 5 * units
				} else if units <= 200 {
					bill = 500 + 7*(units-100)
				} else if units <= 350 {
					bill = 500 + 700 + 10*(units-200)
				}
				fmt.Printf("Bill is $%d.00\n", bill)

				//Go program to print twin prime number
				var aa, bb, next1 int = 0, 0, 0
				var prime1, prime2 bool = true, true
				fmt.Println("2 and 3 are twin prime numbers")
				for aa = 3; aa < 100; aa++ {
					prime1 = true
					prime2 = true
					next1 = aa + 2
					for bb = 1; bb < aa; bb++ {
						if aa%bb == 0 {
							prime1 = false
						}
						if next1%bb == 0 || next1%(bb+2) == 0 {
							prime2 = false
						}
					}
					if prime1 && prime2 {
						fmt.Printf("%d and %d are twin prime numbers\n", aa, next1)
					}
				}

				for ii := 7; ii >= 1; ii-- {
					for jj := 1; jj <= ii; jj++ {
						fmt.Print(jj)
					}
					fmt.Println("")
				}

				for ii := 0; ii < 4; ii++ {
					for jj := 0; jj <= ii; jj++ {
						fmt.Printf(("%c"), jj+65)
					}
					fmt.Println()
				}

				name := "Robert Pike"
				var letter string
				for _, s := range name {
					letter = string(s) + letter
				}
				fmt.Println(letter)

				var space int
				for pos, val := range name {
					if unicode.IsSpace(val) {
						space = pos
					}
				}
				rev1 := ""
				rev2 := ""
				for ii := space - 1; ii >= 0; ii-- {
					rev1 += string(name[ii])
				}
				for ii := len(name) - 1; ii >= space; ii-- {
					rev2 += string(name[ii])
				}
				fmt.Println(rev1 + " " + rev2)

				//Practice Activities: For Loops
				//8 9 10 11 12 13 14 15
				//activity 8: exponential series
				var ll, x int
				fmt.Println("enter x")
				fmt.Scanln(&x)
				fmt.Println("enter n")
				fmt.Scanln(&ll)
				factorial := 1
				exp := 1

				var output float64 = 1
				for ii := 1; ii < ll; ii++ {
					factorial = 1
					exp = 1
					for jj := 1; jj <= ii; jj++ {
						factorial *= jj
					}
					for kk := 0; kk < ii; kk++ {
						exp *= x
					}
					output += (float64(exp) / float64(factorial))
				}
				fmt.Printf("%f\n", output)

				//activity 9
				var integer int
				fmt.Println("enter an integer")
				fmt.Scanln(&integer)
				if integer == 0 {
					fmt.Printf("You entered: 0\nThis number has 1 digit\nThe first digit is 0\nThe last digit is 0\nThe sum of the digits is 0\nThe product of the digits is 0\nThis number is not prime\nThe factorial of this number is 1\n")
				}

				var numDigits, fDigit, lDigit, sumDigits, proDigits, fact int = 1, 0, 0, 0, 1, 1
				factors := 0
				ii := integer

				lDigit = ii % 10
				fDigit = lDigit
				sumDigits += fDigit
				proDigits *= fDigit
				ii /= 10
				for ii > 0 {
					fDigit = ii % 10
					numDigits++
					sumDigits += fDigit
					proDigits *= fDigit
					ii /= 10
				}

				for jj := 1; jj <= integer; jj++ {
					fact *= jj
					if integer%jj == 0 {
						factors += 1
					}
				}

				if factors == 2 {
					fmt.Printf("You entered: %d\nThis number has %d digit(s)\nThe first digit is %d\nThe last digit is %d\nThe sum of the digits is %d\nThe product of the digits is %d\nThis number is prime\nThe factorial of this number is %d\n", integer, numDigits, fDigit, lDigit, sumDigits, proDigits, fact)
				} else {
					fmt.Printf("You entered: %d\nThis number has %d digit(s)\nThe first digit is %d\nThe last digit is %d\nThe sum of the digits is %d\nThe product of the digits is %d\nThis number is not prime\nThe factorial of this number is %d\n", integer, numDigits, fDigit, lDigit, sumDigits, proDigits, fact)
				}

				//activity 10
				var rows int
				fmt.Println("enter number of rows")
				fmt.Scanln(&rows)
				for ii := 1; ii <= rows; ii++ {
					for jj := 0; jj < ii; jj++ {
						fmt.Print(ii)
					}
					fmt.Println("")
				}

				//activity 11: Greatest Common Divisor and 12: Lowest Common Divisor
				//activity 11
				var number1, number2, divisor, upBound int
				fmt.Println("enter number 1")
				fmt.Scanln(&number1)
				fmt.Println("enter number 2")
				fmt.Scanln(&number2)

				if number1 > number2 {
					divisor = number2
					upBound = number1
				} else {
					divisor = number1
					upBound = number2
				}

				for divisor >= 1 {
					if number1%divisor == 0 && number2%divisor == 0 {
						break
					}
					divisor--
				}
				fmt.Printf("The greatest common divisor is %d\n", divisor)

				//activity 12
				for ii := 1; ii <= upBound; ii++ {
					if number1%ii == 0 && number2%ii == 0 {
						fmt.Printf("The lowest common divisor is %d\n", ii)
						break
					}
				}

				//activity 13
				number := 0
				var factors1, factors2 int = 1, 1
				fmt.Println("enter number")
				fmt.Scanln(&number)
				for ii := 1; ii <= number; ii++ {
					factors1, factors2 = 1, 1
					fmt.Printf("ii = %d, number - ii = %d\n", ii, number-ii)
					for jj := 2; jj <= ii; jj++ {
						if ii%jj == 0 {
							factors1++
						}
					}
					for jj := 2; jj <= number-ii; jj++ {
						if (number-ii)%jj == 0 {
							factors2++
						}
					}

					if factors1 == 2 && factors2 == 2 {
						break
					}
				}
				if factors1 == 2 && factors2 == 2 {
					fmt.Printf("%d can be expressed as the sum of two primes", number)
				} else {
					fmt.Printf("%d cannot be expressed as the sum of two primes", number)
				}

				//activity 14

				//activity 15

				//activity 16

				//activity 17

				//activity 18

				//activity 19

				//activity 20

				//activity 21

				//activity 22

		arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
			fmt.Println(arr)

			fmt.Printf("1: sum function\n\n")
			arrSum := Sum(arr[0:4])
			fmt.Println(arr[0:4])
			fmt.Println(arrSum)

			fmt.Printf("\n2: halve an integer\n\n")
			fmt.Println(arrSum)
			fmt.Println(half(&arrSum))
			fmt.Println(arrSum)

			fmt.Printf("\n3: find max element in list of numbers\n\n")
			fmt.Println(findMax(1, 2, -3, 4, -5, 6, -7, 8, -9, 0))

			fmt.Printf("\n4: makeOddGenerator\n\n")
			f := makeOddGenerator()
			for ii := 0; ii < 10; ii++ {
				fmt.Println(f())
			}
			fmt.Printf("\nmakeEvenGenerator\n")
			F := makeEvenGenerator()
			for ii := 0; ii < 10; ii++ {
				fmt.Println(F())
			}

			fmt.Printf("\n5: recursive Fibonacci\n\n")
			for ii := 1; ii <= 10; ii++ {
				fmt.Println(fib(ii))
			}
	*/
}
