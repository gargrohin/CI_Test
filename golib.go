package main

import (
	"fmt"
)

func Validip(str string) int {

	var k int
	for i := 0; i < len(str); i++ {
		if str[i] == '.' || i == len(str) {

			str1 := str[k:i]
			fmt.Println(str1)
			k = i + 1
			//	fmt.Print(k)
			if len(str1) >= 4 {
				fmt.Println("Invalid IP")
			} else {
				num := (str1[0]-'0')*100 + (str1[1]-'0')*10 + (str[2] - '0')
				//fmt.Println(num)
				if num >= 255 {
					fmt.Println("Invalid IP")
				}
			}
		}

	}
	fmt.Print("valid IP")
	return 1
}
