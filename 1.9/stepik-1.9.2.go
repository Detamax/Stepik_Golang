/*

По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".

Sample Input 1:

237
Sample Output 1:

YES
Sample Input 2:

117
Sample Output 2:

NO

*/
package main

import (
	"fmt"
)

var (
	number int = 121
	first_digit int
	current_number int
	next_digit int
)

func main() {
	fmt.Scan(&number)
	first_digit := number / 100
	next_digit := number / 10 % 10
	current_number := number % 10

	if (first_digit != current_number) && (current_number != next_digit) && (next_digit != first_digit ) {
		fmt.Println("YES")
	} else { fmt.Println("NO")

	}

}
