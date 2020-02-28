package mathhelper

import (
	"fmt"
	"strconv"
)

/*
round : 四捨五入
*/
func Round(v float64, decimals int) float64 {
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	return float64(int((v*pow)+0.5)) / pow
}

/*
Decimal : 小數第一位,四捨五入
*/
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", value), 64)
	return value
}

/*
Decimal2 : 小數第二位,四捨五入
*/
func Decimal2(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
