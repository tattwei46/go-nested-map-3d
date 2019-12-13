package main

import "fmt"


func main(){

	
	yearList := []string{"2018", "2019"}
	monthList := []string{"01", "02", "03","04", "05", "06","07", "08", "09","10", "11", "12"}

	year := make(map[string]map[string]map[string]int)

	for _, y := range yearList {
		month, ok := year[y]
		if !ok {
			month = make(map[string]map[string]int)
			year[y] = month

			for _, m := range monthList {
				day, ok := month[m]
				if !ok {
					day = make(map[string]int)
					month[m] = day
				}
			}
		
		}
	}

	year["2018"]["01"]["1"] = 42
	year["2019"]["02"]["1"] = 52
	fmt.Println(year["2018"]["01"]["1"])
	fmt.Println(year["2019"]["02"]["1"])
}