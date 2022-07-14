package main

import (
	"fmt"
	"strconv"
)

var (
	str        string
	indexEqual int
	temp       int
	num        [100]int
	opr        int32
	judge      [100]int
	flag2      int
	new1       string
	n          = 0
)

func main() {
	fmt.Scan(&str)
	expression := []int32(str)
	tmpExp := expression
	for i, v := range expression {
		if v == '=' {
			indexEqual = i
		}
	}
	for i, v := range tmpExp {
		if (i == 0 && (v > 47 && v < 58)) || (i == indexEqual+1 && (v > 47 && v < 58)) {
			opr = '+'
			temp1, _ := strconv.Atoi(string(v))
			temp += temp1
		} else {
			if v > 47 && v < 58 {
				temp1, _ := strconv.Atoi(string(v))
				temp = temp*10 + temp1
				if i == len(tmpExp)-1 {
					if opr == '+' {
						num[n] = temp
					} else {
						num[n] = -temp
					}
				}
			} else {
				if i != 0 {
					if v > 96 && v < 123 {
						if tmpExp[i-1] < 48 || tmpExp[i-1] > 57 {
							temp = 1
						}
						if opr == '=' {
							temp = 1
						}
					}
					if opr == '+' {
						num[n] = temp
					} else {
						if opr == '-' {
							num[n] = -temp
						}
						if tmpExp[i-1] == '=' {
							num[n] = temp
						}
					}
					if v > 96 && v < 123 {
						new1 = string(v)
						judge[n] = 1
					}
					if v == '=' {
						flag2 = n
					}
					temp = 0
					opr = 0
					n++
				}
				if v == '-' || v == '+' || v == '=' {
					opr = v
				}
			}
		}
	}
	sum1 := 0
	sum2 := 0
	for n, _ := range num {
		if n <= flag2 && judge[n] == 0 {
			sum1 += -num[n]
		}
		if n > flag2 && judge[n] == 0 {
			sum1 += num[n]
		}
		if n < flag2 && judge[n] == 1 {
			sum2 += num[n]
		}
		if n > flag2 && judge[n] == 1 {
			sum2 += -num[n]
		}
	}

	if float64(sum1)/float64(sum2) < 0.000000001 && float64(sum1)/float64(sum2) > -0.00000000001 {
		fmt.Printf("%s=%.3f", new1, 0.000)
	} else {
		fmt.Printf("%s=%.3f", new1, float64(sum1)/float64(sum2))
	}
}
