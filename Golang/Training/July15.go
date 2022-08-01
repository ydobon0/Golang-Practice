package Training

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"
)

func July15() {
	activity4()
}

/*
1)Create a slice of 20 of type int and take 20 number create 4 goroutines to take average of every 5 elements ex[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20] ,so 1 goroutine should give average of first 5 elements and next for another 5 so on
*/
func activity1() {
	var nums []int
	num := 0
	ch := make(chan int)
	for ii := 0; ii < 20; ii++ {
		num = rand.Intn(100)
		nums = append(nums, num)
	}

	go avg(nums[0:5], ch)
	go avg(nums[5:10], ch)
	go avg(nums[10:15], ch)
	go avg(nums[15:20], ch)

	aa, bb, cc, dd := <-ch, <-ch, <-ch, <-ch

	fmt.Println(nums)
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)
	fmt.Println((aa + bb + cc + dd) / 4)
}

func avg(nums []int, c chan int) {
	avg := 0
	for ii := 0; ii < len(nums); ii++ {
		avg += nums[ii]
	}

	avg /= len(nums)
	c <- avg
}

/*
Activity 2
2)Take a Paragraphs of text Max word  up to 200 words in a Slice or String and then Pass each  Sentence(up to full stop) to another go routine/routines   1)print it in reverse order(1 goroutine) 2 )count its words(2 goroutine)
*/
var wg sync.WaitGroup

func activity2() {
	paragraph := "Take a Paragraphs of text. Max word up to 200 words in a Slice or String. And then Pass each Sentence (up to full stop) to another go routine/routines. Print it in reverse order (1 goroutine). Count its words (2 goroutine)"
	fmt.Println(paragraph)
	sentences := strings.Split(paragraph, ".")
	ch := make(chan string)
	wg.Add(len(sentences))
	for _, ii := range sentences {
		go reverseSentence(ii, ch)
		go wordCount(ii)
	}
	wg.Wait()
	for range sentences {
		str := <-ch
		fmt.Println(str)
	}
}

func reverseSentence(sentence string, ch chan string) {
	letters := strings.Split(sentence, "")
	result := ""
	for _, ii := range letters {
		result = ii + result
	}

	ch <- result
}

func wordCount(sentence string) {
	words := strings.Split(sentence, " ")
	count := len(words)

	fmt.Println(sentence, "contains ", count, "words.")
	wg.Done()
}

/*
3)Create a goroutine/channels to send and receive structure data type  free to design for any purpose
*/
func activity3() {
	chS := make(chan struct{}, 1)
	chR := make(chan struct{}, 1)
	// fmt.Println("How many structs do you want to send?")
	numSends := 3
	// fmt.Scanln(&numSends)

	for ii := 0; ii < numSends; ii++ {
		chS <- struct{}{}
		send(chS, chR)
		recv(chR, chS)
	}
}

func send(chS chan struct{}, chR chan struct{}) {
	item := <-chS
	fmt.Println("Sending:", item, "!")
	chR <- item
}

func recv(chR chan struct{}, chS chan struct{}) {
	item := <-chR
	fmt.Println("Recieved:", item, "			!")

}

/*
4)Create Two go routine where routine 1 generate random number and append them in  slice where another goroutine sort the slice at the same time . print the slice after every append and sorted array at the same time  side by side
*/

func activity4() {
	snd := make(chan []int, 0)
	rcv := make(chan []int, 0)
	nums := make([]int, 0)

	go addNum(snd, rcv)
	go sortNums(rcv, snd)
	snd <- nums
	select {}
}

func addNum(ch1 chan []int, ch2 chan []int) {
	for {
		nums := <-ch1
		num := rand.Intn(100)
		nums = append(nums, num)
		fmt.Println(nums)
		ch2 <- nums

	}
}

func sortNums(ch2 chan []int, ch1 chan []int) {
	for {
		nums := <-ch2
		// sort.Slice(nums, func(i, j int) bool {
		// 	return i < j
		// })
		sort.Ints(nums)
		fmt.Println("  ", nums)
		ch1 <- nums
	}
}

func July16() { //generate spinner while calculating fibonacci
	const n = 45
	go spinner(100 * time.Millisecond)
	fibN := fib(n) // naive and slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

}

func spinner(delay time.Duration) {
	state := 0
	for {
		state = state % 4
		switch state {
		case 0:
			fmt.Printf("\r-")
			time.Sleep(delay)
		case 1:
			fmt.Printf("\r/")
			time.Sleep(delay)
		case 2:
			fmt.Printf("\r|")
			time.Sleep(delay)
		case 3:
			fmt.Printf("\r\\")
			time.Sleep(delay)
		}
		state += 1
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
