package Modules

import (
	"fmt"
	"regexp"
)

func Module8() {
	fmt.Println("module 8")
	regexPractice()
	practiceRegex()
}

//Practice: Use Regexp
func regexPractice() {
	input := "parrot"

	//variations of gray
	m1, _ := regexp.MatchString("^(gr)(.)(y)", input)
	fmt.Println(m1, "unfinished")
	//properly-formatted email address
	m2, _ := regexp.MatchString("^(.+)@(.+)$", input)
	fmt.Println(m2, "properly formatted email address")
	//Next Find any three-letter words that start with the same letter and end with the same letter, but which might have a different letter in between, such as cat or cot.
	m3, _ := regexp.MatchString(`([a-z]|[A-Z])([a-z]|[A-Z])\2`, input)
	fmt.Println(m3, "unfinished")
	//words that contain a defined set of characters (a, e, i, o, u)
	m4, _ := regexp.MatchString("[a|e|i|o|u]", input)
	fmt.Println(m4, "contains a, e, i, o, or u")
	//words that contain double letters
	m5, _ := regexp.MatchString(`.\1`, input)
	fmt.Println(m5, "contains double letters")
}

func practiceRegex() {
	input := "+19254879335"
	//input = "+1 925-487-9335"
	m1, _ := regexp.MatchString(`^(\+[0-9]+)?\s?([0-9]{3}\-?[0-9]{3}\-?[0-9]{4})$`, input)
	fmt.Println(m1, "valid phone number")

	//IP address in range of 192.160.1.1    - 192.170.1.255
	input = "192.161.1.0"
	m1, _ = regexp.MatchString(`^192\.1((6[0-9])|(70))\.1\.(([1|2]?[0-9]?[0-9])|[1-9])$`, input)
	fmt.Println(m1, "IP address in range")
}

func dataTime() {
	dataTime1()
}

func dataTime1() {
	//Time duration
	days1 := 0
	hours1 := 0
	minutes1 := 0
	seconds1 := 0

	days2 := 0
	hours2 := 0
	minutes2 := 0
	seconds2 := 0

	fmt.Println("enter time 1:")
	fmt.Println("number of days")
	fmt.Scanln(&days1)
	fmt.Println("number of hours")
	fmt.Scanln(&hours1)
	fmt.Println("number of minutes")
	fmt.Scanln(&minutes1)
	fmt.Println("number of seconds")
	fmt.Scanln(&seconds1)

	fmt.Println("enter time 2:")
	fmt.Println("number of days")
	fmt.Scanln(&days2)
	fmt.Println("number of hours")
	fmt.Scanln(&hours2)
	fmt.Println("number of hours")
	fmt.Scanln(&minutes2)
	fmt.Println("number of seconds")
	fmt.Scanln(&seconds2)

	totalSeconds1 := (days1 * 86400) + (hours1 * 3600) + (minutes1 * 60) + (seconds1)
	//fmt.Println(totalSeconds1)
	totalSeconds2 := (days2 * 86400) + (hours2 * 3600) + (minutes2 * 60) + (seconds2)
	//fmt.Println(totalSeconds2)

	var daysT, hoursT, minutesT, secondsT float64
	op := 1

	switch op {
	case 1: //add times
		total := totalSeconds1 + totalSeconds2
		secondsT = float64(total)
		daysT = float64(total / 86400)
		hoursT = float64(total / 3600)
		minutesT = float64(total / 60)

		days := total / 86400
		total = total % 86400
		hours := total / 3600
		total = total % 3600
		minutes := total / 60
		total = total % 60

		fmt.Println(days, " days ", hours, " hours ", minutes, " minutes ", total, " seconds ")
		fmt.Println(daysT, " days")
		fmt.Println(hoursT, " hours")
		fmt.Println(minutesT, " minutes")
		fmt.Println(secondsT, " seconds")
	case 2: //subtract times
		total := totalSeconds1 - totalSeconds2
		secondsT = float64(total)
		daysT = float64(total / 86400)
		hoursT = float64(total / 3600)
		minutesT = float64(total / 60)

		days := total / 86400
		total = total % 86400
		hours := total / 3600
		total = total % 3600
		minutes := total / 60
		total = total % 60

		fmt.Println(days, " days ", hours, " hours ", minutes, " minutes ", total, " seconds ")
		fmt.Println(daysT, " days")
		fmt.Println(hoursT, " hours")
		fmt.Println(minutesT, " minutes")
		fmt.Println(secondsT, " seconds")
	default:
	}

}

func dataTime2() {
	//add or subtract time from a date
}

func dataTime3() {
	//age calculator
}
