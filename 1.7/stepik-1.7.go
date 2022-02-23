/*

Исправьте ошибку в программе ниже:

Sample Input:


Sample Output:

18

*/

package main

import "fmt"

func main(){
    var a int = 8
    //var c int
    const b int = 10
    a = a + b
    //b = b + a 
    fmt.Println(a)
}
