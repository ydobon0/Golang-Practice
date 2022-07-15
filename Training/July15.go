package Training

import "fmt"

func July15() {
	var nums []int
	num := 0
	ch := make(chan int)
	fmt.Println("Enter 20 numbers")
	for ii := 0; ii < 20; ii++ {
		fmt.Scanln(&num)
		nums = append(nums, num)
		fmt.Println(ii+1, "numbers entered.")
	}

	go avg(nums[0:5], ch)
	go avg(nums[5:10], ch)
	go avg(nums[10:15], ch)
	go avg(nums[15:20], ch)

	aa, bb, cc, dd := <-ch, <-ch, <-ch, <-ch

	fmt.Println()
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
