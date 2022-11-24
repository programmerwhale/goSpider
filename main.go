package main

import (
	"fmt"
)

func main()  {
	res:=books.SelectSql()
	fmt.Println(res)
}
