package Modules

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func Module8() {
	fmt.Println("module 8")
	regexPractice()
	practiceRegex()

	mymatch("abbba")
	mymatch("cbc")
	mymatch("cdddc")
	mymatch("adc")
	dataTime()
}

//Practice: Use Regexp
func regexPractice() {
	input := "parrot"

	//variations of gray
	m1, _ := regexp.MatchString("^(gr)[a-z](y)", input)
	fmt.Println(m1, "variations of gray/grey")
	//properly-formatted email address
	m2, _ := regexp.MatchString("^(.+)@(.+)$", input)
	fmt.Println(m2, "properly formatted email address")
	//Next Find any three-letter words that start with the same letter and end with the same letter, but which might have a different letter in between, such as cat or cot.
	m3, _ := regexp.MatchString(`a[a-z]a|b[a-z]b|c[a-z]c|d[a-z]d|e[a-z]e|f[a-z]f|g[a-z]g|h[a-z]h|i[a-z]i|j[a-z]j|k[a-z]k|l[a-z]l|m[a-z]m|n[a-z]n|o[a-z]o|p[a-z]p|q[a-z]q|r[a-z]r|s[a-z]s|t[a-z]t|u[a-z]u|v[a-z]v|w[a-z]w|x[a-z]x|y[a-z]y|z[a-z]z|A[A-Z]A|B[A-Z]B|C[A-Z]C|[A-Z]DD|E[A-Z]E|F[A-Z]F|G[A-Z]G|H[A-Z]H|I[A-Z]I|J[A-Z]J|K[A-Z]K|L[A-Z]L|M[A-Z]M|N[A-Z]N|O[A-Z]O|P[A-Z]P|Q[A-Z]Q|R[A-Z]R|S[A-Z]S|T[A-Z]T|U[A-Z]U|V[A-Z]V|W[A-Z]W|X[A-Z]X|Y[A-Z]Y|Z[A-Z]Z`, input)
	fmt.Println(m3, "three-letter words that start with the same letter and end with the same letter, but which might have a different letter in between, such as cat or cot")
	//words that contain a defined set of characters (a, e, i, o, u)
	m4, _ := regexp.MatchString("[a|e|i|o|u]", input)
	fmt.Println(m4, "contains a, e, i, o, or u")
	//words that contain double letters
	m5, _ := regexp.MatchString(`aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz|AA|BB|CC|DD|EE|FF|GG|HH|II|JJ|KK|LL|MM|NN|OO|PP|QQ|RR|SS|TT|UU|VV|WW|XX|YY|ZZ`, input)
	fmt.Println(m5, "contains double letters")
	//input = "+1 925-487-9335"
	m6, _ := regexp.MatchString(`^(\+[0-9]+)?\s?([0-9]{3}\-?[0-9]{3}\-?[0-9]{4})$`, input)
	fmt.Println(m6, "valid phone number")
	//IP address in range of 192.160.1.1    - 192.170.1.255
	m7, _ := regexp.MatchString(`^192\.1((6[0-9])|(70))\.1\.(([1|2]?[0-9]?[0-9])|[1-9])$`, input)
	fmt.Println(m7, "IP address in range")
}

func practiceRegex() {
	input := "parroting="
	//Write a program to check that a string contains only letters and numbers (e.g., a-z, A-Z, 0-9).
	m1, _ := regexp.MatchString(`^\w+$`, input)
	fmt.Println(m1, "contains only letters and numbers")
	//Write a program that finds all strings that include the letter i followed by zero or more instances of the letter n.
	m2, _ := regexp.MatchString("in*", input)
	fmt.Println(m2, "include the letter i followed by zero or more instances of the letter n.")
	//Write a program that finds all strings that include the letter i followed by one or more instances of the letter n.
	m3, _ := regexp.MatchString("in+", input)
	fmt.Println(m3, "include the letter i followed by one or more instances of the letter n.")
	//Write a program that finds all strings that include the letter i followed by one or two instances of the letter n.
	m4, _ := regexp.MatchString("in|inn", input)
	fmt.Println(m4, "include the letter i followed by one or two instances of the letter n.")
	//Write a program that finds all strings that include the letter i followed by three instances of the letter n.
	m5, _ := regexp.MatchString("innn", input)
	fmt.Println(m5, "include the letter i followed by three instances of the letter n.")
}

func dataTime() { //Module 8: Date-Time Calculator
	fmt.Println("Please enter all values as integers")
	op := 0
	for op != 1 && op != 2 && op != 3 {
		fmt.Println("Enter the operation you want")
		fmt.Println("1. Add/Subtract two different lengths of time")
		fmt.Println("2. Given a date and time, add or subtract an inputed time duration")
		fmt.Println("3. Calculate the time difference between two inputted dates")

		fmt.Scanln(&op)

		switch op {
		case 1:
			dataTime1()
		case 2:
			dataTime2()
		case 3:
			dataTime3()
		default:
			fmt.Println("Invalid operation")
		}
	}

}

//*
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
	op := 0

	for op != 1 && op != 2 {
		fmt.Println("enter 1 to add times. enter 2 to subtract times")
		fmt.Scanln(&op)
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
			break
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
			break
		default:
			fmt.Println("invalid operation")
		}
	}

}

//!
func dataTime2() {
	//add or subtract time from a date
	sYear := 0
	sMonth := 0
	sDay := 0
	sHour := 0
	sMinute := 0
	sSecond := 0

	years := 0
	months := 0
	days := 0
	hours := 0
	minutes := 0
	seconds := 0
	fmt.Println("enter starting year")
	fmt.Scanln(&sYear)
	fmt.Println("enter starting month")
	fmt.Scanln(&sMonth)
	fmt.Println("enter starting day")
	fmt.Scanln(&sDay)
	fmt.Println("enter starting hour")
	fmt.Scanln(&sHour)
	fmt.Println("enter starting minutes")
	fmt.Scanln(&sMinute)
	fmt.Println("enter starting second")
	fmt.Scanln(&sSecond)

	fmt.Println("enter time:")
	fmt.Println("number of years")
	fmt.Scanln(&years)
	fmt.Println("number of months")
	fmt.Scanln(&months)
	fmt.Println("number of days")
	fmt.Scanln(&days)
	fmt.Println("number of hours")
	fmt.Scanln(&hours)
	fmt.Println("number of minutes")
	fmt.Scanln(&minutes)
	fmt.Println("number of seconds")
	fmt.Scanln(&seconds)

	start := time.Now()

	switch (sMonth % 12) + 1 {
	case 1:
		start = time.Date(sYear, 1, sDay, sHour, sMinute, sSecond, 0, time.Local)
	case 2:
		start = time.Date(sYear, 2, sDay, sHour, sMinute, sSecond, 0, time.Local)
	case 3:
		start = time.Date(sYear, 3, sDay, sHour, sMinute, sSecond, 0, time.Local)
	case 4:
		start = time.Date(sYear, 4, sDay, sHour, sMinute, sSecond, 0, time.Local)
	case 5:
		start = time.Date(sYear, 5, sDay, sHour, sMinute, sSecond, 0, time.Local)
	case 6:
		start = time.Date(sYear, 6, sDay, sHour, sMinute, sSecond, 0, time.Local)
	case 7:
		start = time.Date(sYear, 7, sDay, sHour, sMinute, sSecond, 0, time.Local)
	case 8:
		start = time.Date(sYear, 8, sDay, sHour, sMinute, sSecond, 0, time.Local)
	case 9:
		start = time.Date(sYear, 9, sDay, sHour, sMinute, sSecond, 0, time.Local)
	case 10:
		start = time.Date(sYear, 10, sDay, sHour, sMinute, sSecond, 0, time.Local)
	case 11:
		start = time.Date(sYear, 11, sDay, sHour, sMinute, sSecond, 0, time.Local)
	case 12:
		start = time.Date(sYear, 12, sDay, sHour, sMinute, sSecond, 0, time.Local)
	}

	totalHrs := (years * 8760) + (months * 730) + (days * 24) + hours
	totalSecs := (minutes * 60) + seconds
	durStr := strconv.Itoa(totalHrs) + "h" + strconv.Itoa(totalSecs) + "s"

	dur, _ := time.ParseDuration(durStr)

	op := 0
	fmt.Println("Press 1 to add. 2 to subtract")
	for op != 1 && op != 2 {
		fmt.Scanln(&op)
		fmt.Println(dur)
		switch op {
		case 1:
			fmt.Println(start.Add(dur))
			break
		case 2:
			fmt.Println(start.Add(-dur))
			break
		default:
			fmt.Println("Invalid operation!")
		}
	}

}

//?
func dataTime3() {
	//age calculator
	/*
		Given a start date and an end date, calculate the amount of time passed between the dates, displayed in years, months, days, hours, minutes and seconds.
		Before submit the make sure it will be well documented                and in  separate folder

	*/
	sYears := 0
	sMonths := 0
	sDays := 0
	sHours := 0
	sMinutes := 0
	sSeconds := 0

	eYears := 0
	eMonths := 0
	eDays := 0
	eHours := 0
	eMinutes := 0
	eSeconds := 0

	start := time.Now()
	end := time.Now()
	fmt.Println("enter start date:")
	fmt.Println("enter year")
	fmt.Scanln(&sYears)
	fmt.Println("enter month")
	fmt.Scanln(&sMonths)
	fmt.Println("enter day")
	fmt.Scanln(&sDays)
	fmt.Println("enter hour")
	fmt.Scanln(&sHours)
	fmt.Println("enter minute")
	fmt.Scanln(&sMinutes)
	fmt.Println("enter second")
	fmt.Scanln(&sSeconds)
	switch (sMonths % 12) + 1 {
	case 1:
		start = time.Date(sYears, 1, sDays, sHours, sMinutes, sSeconds, 0, time.Local)
	case 2:
		start = time.Date(sYears, 2, sDays, sHours, sMinutes, sSeconds, 0, time.Local)
	case 3:
		start = time.Date(sYears, 3, sDays, sHours, sMinutes, sSeconds, 0, time.Local)
	case 4:
		start = time.Date(sYears, 4, sDays, sHours, sMinutes, sSeconds, 0, time.Local)
	case 5:
		start = time.Date(sYears, 5, sDays, sHours, sMinutes, sSeconds, 0, time.Local)
	case 6:
		start = time.Date(sYears, 6, sDays, sHours, sMinutes, sSeconds, 0, time.Local)
	case 7:
		start = time.Date(sYears, 7, sDays, sHours, sMinutes, sSeconds, 0, time.Local)
	case 8:
		start = time.Date(sYears, 8, sDays, sHours, sMinutes, sSeconds, 0, time.Local)
	case 9:
		start = time.Date(sYears, 9, sDays, sHours, sMinutes, sSeconds, 0, time.Local)
	case 10:
		start = time.Date(sYears, 10, sDays, sHours, sMinutes, sSeconds, 0, time.Local)
	case 11:
		start = time.Date(sYears, 11, sDays, sHours, sMinutes, sSeconds, 0, time.Local)
	case 12:
		start = time.Date(sYears, 12, sDays, sHours, sMinutes, sSeconds, 0, time.Local)
	}

	fmt.Println("enter end date:")
	fmt.Println("enter year")
	fmt.Scanln(&eYears)
	fmt.Println("enter month")
	fmt.Scanln(&eMonths)
	fmt.Println("enter day")
	fmt.Scanln(&eDays)
	fmt.Println("enter hour")
	fmt.Scanln(&eHours)
	fmt.Println("enter minute")
	fmt.Scanln(&eMinutes)
	fmt.Println("enter second")
	fmt.Scanln(&eSeconds)

	switch (eMonths % 12) + 1 {
	case 1:
		end = time.Date(eYears, 1, eDays, eHours, eMinutes, eSeconds, 0, time.Local)
	case 2:
		end = time.Date(eYears, 2, eDays, eHours, eMinutes, eSeconds, 0, time.Local)
	case 3:
		end = time.Date(eYears, 3, eDays, eHours, eMinutes, eSeconds, 0, time.Local)
	case 4:
		end = time.Date(eYears, 4, eDays, eHours, eMinutes, eSeconds, 0, time.Local)
	case 5:
		end = time.Date(eYears, 5, eDays, eHours, eMinutes, eSeconds, 0, time.Local)
	case 6:
		end = time.Date(eYears, 6, eDays, eHours, eMinutes, eSeconds, 0, time.Local)
	case 7:
		end = time.Date(eYears, 7, eDays, eHours, eMinutes, eSeconds, 0, time.Local)
	case 8:
		end = time.Date(eYears, 8, eDays, eHours, eMinutes, eSeconds, 0, time.Local)
	case 9:
		end = time.Date(eYears, 9, eDays, eHours, eMinutes, eSeconds, 0, time.Local)
	case 10:
		end = time.Date(eYears, 10, eDays, eHours, eMinutes, eSeconds, 0, time.Local)
	case 11:
		end = time.Date(eYears, 11, eDays, eHours, eMinutes, eSeconds, 0, time.Local)
	case 12:
		end = time.Date(eYears, 12, eDays, eHours, eMinutes, eSeconds, 0, time.Local)
	}

	age := end.Sub(start)

	fmt.Println("Duration is: ", age)
	tSecs := int(age.Seconds())
	tMins := tSecs / 60
	tHours := tMins / 60
	tDays := tHours / 24
	tMonths := tDays / 30
	tYears := tMonths / 12

	tSecs %= 60
	tMins %= 60
	tHours %= 24
	tDays %= 30
	tMonths %= 12

	fmt.Println(tYears, " years, ", tMonths, " months, ", tDays, " days ")
	fmt.Println(((tYears * 12) + tMonths), " months, ", tDays, " days ")
	fmt.Println(((((tYears * 12) + tMonths) * 30) + tDays), " days ")
	fmt.Println(((((((tYears * 12) + tMonths) * 30) + tDays) * 24) + tHours), " hours ")
	fmt.Println(((((((((tYears * 12) + tMonths) * 30) + tDays) * 24) + tHours) * 60) + tMins), " minutes ")
	fmt.Println(((((((((((tYears * 12) + tMonths) * 30) + tDays) * 24) + tHours) * 60) + tMins) * 60) + tSecs), " seconds ")
}

func mymatch(name string) {
	x := []byte(name)

	regEx1, _ := regexp.Compile("^" + string(x[0]) + ".*" + string(x[0]) + "$")

	match1 := regEx1.Match([]byte(name))
	fmt.Println(match1)
}

//Many of you has not updated your GitHub with Activities from all Modules Name I will post on slack tomorrow It has to be completed by Monday ,till Assessment will withhold
