package main

import (
    "fmt"
    "github.com/shopspring/decimal"
)

func main() {
    // 创建高精度的圆周率
    pi := decimal.NewFromFloat(3.141592653589793238462643383279)
    
    // 计算圆的面积（假设半径为 2）
    radius := decimal.NewFromInt(2)
    area := pi.Mul(radius.Mul(radius))
    
    fmt.Println("Pi =", pi)
    fmt.Println("Area of circle with radius 2 =", area)
    
    // 进行其他运算
    // 加法
    sum := pi.Add(decimal.NewFromInt(1))
    // 减法
    diff := pi.Sub(decimal.NewFromInt(1))
    // 乘法
    product := pi.Mul(decimal.NewFromInt(2))
    // 除法（设置精度为 10 位）
    quotient := pi.Div(decimal.NewFromInt(2))
    
    fmt.Println("Pi + 1 =", sum)
    fmt.Println("Pi - 1 =", diff)
    fmt.Println("Pi * 2 =", product)
    fmt.Println("Pi / 2 =", quotient)
}