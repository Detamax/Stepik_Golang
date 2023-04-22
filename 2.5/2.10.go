package main

import (
	"fmt"
	"strings"
)

func main()  {
	var listWord string
	_, _ = fmt.Scan(&listWord)
	var delteRepe strings.Builder
	for _, r := range listWord{
		if strings.Count(listWord, string(r)) == 1{
			delteRepe.WriteRune(r)
		}
	}
	fmt.Print(delteRepe.String())
}
