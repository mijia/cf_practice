package main

import "fmt"

func main() {
	weekDay := 1
	total := 0
	for year := 1901; year <= 2000; year += 1 {
		for month := 1; month <= 12; month += 1 {
			days := 31
			if month == 4 || month == 6 || month == 9 || month == 11 {
				days = 30
			}
			if month == 2 {
				isLeapYear := false
				if year%400 == 0 {
					isLeapYear = true
				} else {
					if year%4 == 0 && year%100 != 0 {
						isLeapYear = true
					}
				}
				if isLeapYear {
					days = 29
				} else {
					days = 28
				}
			}
			for day := 1; day <= days; day += 1 {
				if weekDay == 1 && day == 1 {
					total += 1
				}
				weekDay += 1
				weekDay = weekDay%7 + 1
			}
		}
	}
	fmt.Println(total)
}
