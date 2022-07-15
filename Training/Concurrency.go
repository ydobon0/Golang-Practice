package Training

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func Practice() {
	// printNow = make(chan bool)

	// wait := new(sync.WaitGroup)
	// wait.Add(2)
	// xx := gen1()
	// yy := gen2()

	// go xx(wait)
	// go yy(wait)

	// wait.Wait()

	// go checkParity()
	// go genNumber()
	// for {
	// }

	ConcurrAct5()
}

func gen1() func(wait *sync.WaitGroup) {
	return func(wait *sync.WaitGroup) {
		for ii := 0; ii <= 100; ii++ {
			if _, ok := <-printNow; ok {
				fmt.Println("Func 1: ", ii)
			}
		}
		defer wait.Done()
	}
}

func gen2() func(wait *sync.WaitGroup) {
	return func(wait *sync.WaitGroup) {
		for ii := 100; ii >= 0; ii-- {
			fmt.Println("Func 2:		", ii)

			printNow <- true

			time.Sleep(1 * time.Millisecond)
		}
		defer wait.Done()
	}
}

//* ---------- Using 2 goroutines, generate a number with one, then check if it's even or odd. The goroutines should go one after another alternating ----------
var (
	printNow chan bool
	i        int
)

func checkParity() {
	for {
		if _, ok := <-printNow; ok {
			//fmt.Println("Recieved !", i)
			if i%2 == 0 {
				fmt.Println(i, "is even")
			} else {
				fmt.Println(i, "is odd")
			}
		} //else{
		//return
		//   os.Exit(0)
		// }

	}
}
func genNumber() {
	for {
		for i = 0; i < 10; i++ {
			fmt.Println(i)

			printNow <- true

			time.Sleep(1 * time.Millisecond)
		}
		// close(printNow)
		//return
	}
}

/*
if channels are written to, they stop until they are read from
channels can be blocked
If there are multiple goroutines, you can have one goroutine that reads the other goroutines and displays them
channels are pipelines between goroutines

we need to close channels after we are done with them. The sender has to close the channel. program will panic if the channel is not closed
if you read a channel using <-
-> to write to a channel - puts something inside the channel
channels can be specified as only read or only write
*/
var wg sync.WaitGroup

func ConcurrAct1() {
	// Creating a channel without a buffer.
	// This will make the calling go routine to wait until a recevier is up and processing the data.
	// If we wanted a buffer we could specify it this way '
	cS := make(chan int, 1)
	cR := make(chan int, 1)
	// Closing the channel on the sender when the job is done.
	// NB : The sender shall allways close the channel, never the receiver.
	go sender(cS, cR)
	go reciever(cR, cS)
	cS <- 1
	select {} //allows both routines to complete? removing it makes it so you need to have wait groups
}

func ConcurrAct2() {
	// Creating a channel without a buffer.
	// This will make the calling go routine to wait until a recevier is up and processing the data.
	// If we wanted a buffer we could specify it this way '
	cS := make(chan struct{}, 1)
	cR := make(chan struct{}, 1)
	// Closing the channel on the sender when the job is done.
	// NB : The sender shall allways close the channel, never the receiver.
	cS <- struct{}{}
	go mix(cS, cR)

	select {} //allows both routines to complete? removing it makes it so you need to have wait groups
}

func ConcurrAct3() { //2 goroutines, 1 generates numbers, 1 checks if they are prime
	sender := make(chan int, 1)
	checker := make(chan int, 1)
	go sendNum(sender, checker)
	go checkPrime(checker, sender)
	sender <- 1
	select {}
}

func sendNum(snd <-chan int, rec chan<- int) {
	for {
		inform := <-snd
		rec <- inform
	}
}

func checkPrime(rec <-chan int, snd chan<- int) {
	for {
		inform := <-rec
		factors := 0
		for ii := 1; ii <= inform; ii++ {
			if inform%ii == 0 {
				factors++
			}
		}
		if factors == 2 {
			fmt.Println(inform, " is a prime number")
		}
		snd <- inform + 1
	}
}

func inc(s string) chan int {
	out := make(chan int)
	go func() {

		for i := 0; i <= 15; i++ {
			//fmt.Println(i, "A")
			out <- i

		}
		close(out)
	}()

	return out
}
func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var factors int
		for n := range c {
			factors = 0
			for ii := 1; ii <= n; ii++ {

				if n%ii == 0 {

					factors++
					fmt.Println(n, factors)
				}
				if factors == 2 {
					out <- n
					fmt.Println(n)
				}
			}

		}
		//fmt.Println("done looping")
		close(out)
	}()
	return out
}

func ConcurrAct4() {
	n1 := inc("Genarate1")
	p1 := puller(n1)
	fmt.Println("Sum ", <-p1)

}

var counter int64

func incrementor(s string) {
	for i := 0; i <= 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		// counter++
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Countert ", counter)
	}
	wg.Done()
}
func ConcurrAct5() {
	wg.Add(2)
	go incrementor("Routine 1")
	go incrementor("Routine 2")

	wg.Wait()
	fmt.Println("Countert ", counter)
}

func mix(s, r chan struct{}) {
	for {
		select { // like switch but it works with channels. If any of the case situations occur, the appropriate code executes
		case d1 := <-s:
			fmt.Println("recieved from send", d1)
			r <- struct{}{}
		case d2 := <-r:
			fmt.Println("recieved from recieve", d2)
			s <- struct{}{}
		}
	}
}

func sender(snd <-chan int, rec chan<- int) { // snd is the input, and rec is the output from the input
	for {
		inform := <-snd // inform is 1 on the first iteration //sender manipulates some data. the data we get from sender goes into inform
		fmt.Println("get info from snd ", inform)
		//time.Sleep(1 * time.Second)
		rec <- inform + 1 //write the data into rec channel
	}
}

//no need to sleep since writing to a channel blocks the execution until the channel is read from
func reciever(rec <-chan int, snd chan<- int) { //recieves data and resends it
	for {
		inform := <-rec //sender manipulates some data. the data we get from sender goes into inform
		fmt.Println("get info from rec          ", inform)
		//time.Sleep(1 * time.Second)
		snd <- inform + 1 //write the data into snd channel
	}
}

func get(c chan int) {
	for {
		// Using two variables with channels, the 2nd variable will have the state of the channel as type bool
		myVar, ok := <-c // reading    <-ch    writing  ch<- 12
		fmt.Printf("The value of 'ok' = %v, and the type is %T\n", ok, ok)
		if ok {
			fmt.Println("content of myVar = ", myVar, "and the content of c = ", c)
			time.Sleep(time.Millisecond * 500)
		} else {
			// The channel is closed since 'ok=false', and we do the logic to decrement the waitgroup, and leave the function
			fmt.Println("Channel was closed, returning for loop")
			// Decrement the waitgroup value by 1 by calling wg.Done()
			wg.Done()
			return
		}
	}
}
