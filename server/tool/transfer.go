/**
* @Author: yexichen
* @Date:2023/7/11-19:23
* @Desc
**/

package tool

import (
	"math"
	"strconv"
)

func StringTOInt(number string) int {
	len := len(number)
	var trannum int = 0
	for i := 0; i < len; i++ {
		trannum = trannum*10 + int(number[i]-'0')
	}
	return trannum
}

func StringToFloat(number string) float32 {
	var atrannum float32 = 0
	var btrannum float32 = 0
	var i = 0
	for ; number[i] != '.'; i++ {
	}

	for j := 0; j < i; j++ {
		atrannum = atrannum*10 + float32(number[j]-'0')
	}
	for j := len(number) - 1; j > i; j-- {
		btrannum = btrannum*0.1 + float32(number[j]-'0')
	}
	return atrannum + btrannum*0.1
}

func Float64ToFloat2(f float64) float64 {
	f1 := math.Trunc(f*1e2+0.5) * 1e-2
	f1Str := strconv.FormatFloat(f1, 'f', 2, 64)
	value, _ := strconv.ParseFloat(f1Str, 64)
	return value
}

func GetInterfaceToInt(t1 interface{}) int {
	var t2 int
	switch t1.(type) {
	default:
		t2 = t1.(int)
		break
	}
	return t2
}
