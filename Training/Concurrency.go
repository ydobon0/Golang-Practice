package Training

import (
	"fmt"
	"sync"
)

func gen1() func(wait *sync.WaitGroup) {
	return func(wait *sync.WaitGroup) {
		for ii := 0; ii <= 100; ii++ {
			fmt.Println("Func 1: ", ii)
		}
		defer wait.Done()
	}
}

func gen2() func(wait *sync.WaitGroup) {
	return func(wait *sync.WaitGroup) {
		for ii := 100; ii >= 0; ii-- {
			fmt.Println("Func 2:		", ii)
		}
		defer wait.Done()
	}
}

func Practice() {
	wait := new(sync.WaitGroup)
	wait.Add(2)
	xx := gen1()
	yy := gen2()

	go xx(wait)
	go yy(wait)

	wait.Wait()
}
