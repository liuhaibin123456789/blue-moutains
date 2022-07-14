package main

import (
	"fmt"
	"log"
)

func main() {
	str := ""
	n := -1
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatalf("输入错误: %v", err)
		return
	}
	_, err = fmt.Scan(&str)
	if err != nil {
		log.Fatalf("输入错误: %v", err)
		return
	}

	bytes := make([]uint8, 0, len(str))
	for i := 0; i < len(str); i++ {
		bytes = append(bytes, str[i])
	}
	password := getPassword(uint8(n), bytes)
	fmt.Println(string(password))
}

func getPassword(n uint8, bytes []uint8) (newBytes []uint8) {
	var chA uint8 = 'a'
	var chZ uint8 = 'z'
	for i := 0; i < len(bytes); i++ {
		if bytes[i]-chA+n+1 > 26 {
			bytes[i] = (n - 1) - (chZ - bytes[i]) + chA
		} else {
			bytes[i] = n + bytes[i]
		}
	}
	newBytes = bytes
	return
}
