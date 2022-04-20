package math

import (
	"math"
)

//FloatDecimal 浮点数保留小数点后prec位，去尾法
func FloatDecimal(f float64, prec int) float64 {
	precValue := math.Pow(10, float64(prec))   //Pow returns x**y, the base-x exponential of y.Pow返回x**y，以x为底的y的指数。
	return math.Trunc(f*precValue) / precValue //Trunc returns the integer value of x. 返回x 的整数值
}

//FloatDecimalRound 浮点数保留小数点后prec位，  四舍五入
func FloatDecimalRound(f float64, prec int) float64 {
	return FloatDecimal(f+5/math.Pow(10, float64(prec+1)), prec)
}

//FloatEqual判断浮点数是否相等,prec控制精度
func FloatEqual(x, y, prec float64) bool {
	return math.Dim(x, y) < prec //Dim returns the maximum of x-y or 0.
}

// Float32Equal 判断float32是否相等，prec控制精度
func Float32Equal(x, y, prec float32) bool {
	return FloatEqual(float64(x), float64(y), float64(prec))
}
