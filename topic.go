package main

import (
	"fmt"
	"math"
)

var charactor = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func CalDivideTimes(n int) (times int) {
	mathcount := float64(n % 2)
	if n == 1 {
		return 0
	}
	if mathcount == 1 {
		return -1
	} else {
		return int(n / 2)
	}
}
func DecimalToAZ52(dec int) string {
	strvals := string("")
	if dec == 0 {
		strvals = "a"
	} else {
		for dec > 0 {
			i := int(math.Floor(float64(dec % 52))) // 取餘數
			strvals = string(charactor[i]) + strvals
			dec = dec / 52 // 除以基數
		}
	}

	return strvals
}
func AZ52ToDecimal(az52 string) int {
	// 定義字元表
	charactor := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// 初始化結果
	totalNum := 0
	base := 1 // 對應 52^0

	// 從右往左處理字串
	for i := len(az52) - 1; i >= 0; i-- {
		// 找到字元在字元表的位置
		char := az52[i]
		index := -1
		for j := 0; j < len(charactor); j++ {
			if charactor[j] == char {
				index = j
				break
			}
		}

		// 如果找不到字元，直接報錯
		if index == -1 {
			fmt.Printf("Invalid character: %c\n", char)
			return -1
		}

		// 累加權重值
		totalNum += index * base

		// 更新基數（次方）
		base *= 52
	}

	return totalNum
}

func main() {
	// Get a greeting message and print it.
	fmt.Printf("題目一 \n")
	number := int64(0)
	number = int64(CalDivideTimes(1))
	fmt.Printf("\n CalDivideTimes(%d) %d", 1, number)
	number = int64(CalDivideTimes(7))
	fmt.Printf("\n CalDivideTimes(%d) %d", 7, number)
	number = int64(CalDivideTimes(4))
	fmt.Printf("\n CalDivideTimes(%d) %d", 4, number)
	fmt.Printf("\n 題目二 \n")
	DecimalNumber := 4073126405
	AZ52Number := "kLdXLX"
	// 預期執行結果
	fmt.Printf("\n DecimalToAZ52(%d) %s", DecimalNumber, DecimalToAZ52(DecimalNumber)) // kLdXLX
	fmt.Printf("\n DecimalToAZ52(%d) %s", 0, DecimalToAZ52(0))                         // a
	fmt.Printf("\n DecimalToAZ52(%d) %s", 26, DecimalToAZ52(26))                       // A
	fmt.Printf("\n AZ52ToDecimal(%s) %d", AZ52Number, AZ52ToDecimal(AZ52Number))       // 4073126405

}
