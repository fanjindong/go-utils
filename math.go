package utils

import "math"

// math.Round 有问题
// Round 四舍五入，ROUND_HALF_UP 模式实现
// 返回将 val 根据指定精度 precision（十进制小数点后数字的数目）进行四舍五入的结果。precision 也可以是负数或零。
func Round(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}

func Max(values ...float64) float64 {
	var target float64 = values[0]
	for i := 1; i < len(values); i++ {
		if target < values[i] {
			target = values[i]
		}
	}
	return target
}

func Min(values ...float64) float64 {
	var target float64 = values[0]
	for i := 1; i < len(values); i++ {
		if target > values[i] {
			target = values[i]
		}
	}
	return target
}

func Mean(values ...float64) float64 {
	n := float64(len(values))
	if n == 0 {
		return 0
	}
	return Sum(values...) / n
}

func Sum(values ...float64) float64 {
	var target float64
	for _, num := range values {
		target += num
	}
	return target
}
