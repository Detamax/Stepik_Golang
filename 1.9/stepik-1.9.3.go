/*

Дано неотрицательное целое число. Найдите и выведите первую цифру числа. 

Формат входных данных
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.


*/


package main

import (
	"fmt"
)

var number int


func main() {
	fmt.Scan(&number)
	switch   {
	case number > 9 && number < 100 : //Два знака
		number = number / 10
		fmt.Println(number)
	case number > 99 && number < 1000: //Три знака
		number = number / 100
		fmt.Println(number)
	case number > 999 && number < 10000: //Четыре знака
		number = number / 1000
		fmt.Println(number)
	case number < 9: //Число однозначное
		fmt.Println(number)
	case number >= 10000: //Много
		number = number / 10000
		fmt.Println(number)
	}
}
