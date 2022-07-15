package Training

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

func July15() {
	activity2()
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
