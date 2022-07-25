package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"test/Modules"
	"time"
)

var port = flag.String("port", "8000", "the port number")
var NewYork = flag.String("NewYork", "8010", "the port number")
var tz = flag.String("tz", "8000", "time zone")

// call like ./clock1 -port 8010 -tz=Asia/Tokyo
func main() {
	//July21()
	Modules.Module8()
	// // os.Args holds the command-line arguments, starting with the program name.
	// for _, place := range os.Args[1:] {
	// 	//fmt.Println(place)
	// 	loc := strings.Split(place, "=")
	// 	fmt.Println(loc)
	// 	go net.Listen("tcp", loc[1])
	// 	go timeAt(loc[1])

	// 	// place is gonna look like NewYork=localhost:8010
	// 	// get the localhost:port part
	// 	// make a connection to that localhost:port
	// }
	// fmt.Println("!!")
	// for {

	// }
}

func timeAt(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	go shouldCopy(os.Stdout, conn)
	handleConn(conn)
}

func shouldCopy(dst io.Writer, src io.Reader) {
	fmt.Println("!")
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
	fmt.Println("?")
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		fmt.Println(c.RemoteAddr(), "	", time.Now().Format("15:04:05\n"))
		time.Sleep(1 * time.Second)
	}
}

// import (
// 	"flag"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// 	"time"
// )

// var port = flag.String("port", "8000", "the port number")
// var tz = flag.String("tz", "8000", "time zone")

// func main2() {
// 	flag.Parse()
// 	fmt.Println(*port)
// 	listener, err := net.Listen("tcp", "localhost:"+*port)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			log.Print(err) // e.g., connection aborted
// 			continue
// 		}
// 		go handleConn(conn) // handle one connection at a time
// 	}
// }

// func handleConn(c net.Conn) {
// 	defer c.Close()
// 	for {
// 		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
// 		if err != nil {
// 			return // e.g., client disconnected
// 		}
// 		fmt.Println(c.RemoteAddr())
// 		time.Sleep(1 * time.Second)
// 	}
// }

/*
where goroutines/concurrency are used

net package - http is built on it - net package is used for making requests
	for client/server programs communicating over sockets
	tcp - for making connections
	listener, err := net.Listen("tcp", "localhost:8000")

defer - do this action as the function is closing. will happen no matter how the function ends
	when opening file you need to close it before you exit, so you use defer to do that\
	use defer instead of putting the action everywhere because its more simple that way

	for windows
*/

func July16() {
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

var wg1 sync.WaitGroup
var wg2 sync.WaitGroup

func July21() {
	// go one()
	// wg2.Add(1)
	// wg2.Wait()
	// two()
	ch := make(chan bool)

	go one1(ch)
	ch <- false
	two1(ch)
	//time.Sleep(4 * time.Millisecond)
	fmt.Println("main Goroutine")
}

func one() {

	for i := 0; i < 100; i++ {
		wg1.Add(1)
		fmt.Println(i + 1)
		time.Sleep(1 * time.Millisecond)
		wg2.Done()
		wg1.Wait()
	}
	wg2.Done()
}

func two() {

	for i := 0; i < 100; i++ {
		wg2.Add(1)
		fmt.Println("	", 100-i)
		time.Sleep(1 * time.Millisecond)
		wg1.Done()
		wg2.Wait()
	}
}

func one1(ch chan bool) {

	for i := 0; i < 100; i++ {
		ok := <-ch
		if !ok {
			fmt.Println(i + 1)
			time.Sleep(1 * time.Millisecond)
		}
		ch <- true
	}
	close(ch)
}

func two1(ch chan bool) {

	for i := 0; i < 100; i++ {
		ok := <-ch
		if ok {
			fmt.Println("	", 100-i)
			time.Sleep(1 * time.Millisecond)
		}
		ch <- false

	}
}
