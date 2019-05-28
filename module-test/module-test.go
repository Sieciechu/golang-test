package main

import (
	"fmt"
	"time"

	"golang.org/x/text/currency"

	_ "gopkg.in/yaml.v2"
)

func main() {

	fmt.Println("hello2")

	t1799, _ := time.Parse("2006-01-02", "1799-01-01")
	for it := currency.Query(currency.Date(t1799)); it.Next(); {
		from := ""
		if t, ok := it.From(); ok {
			from = t.Format("2006-01-02")
		}
		fmt.Printf("%v is used in %v since: %v\n", it.Unit(), it.Region(), from)
	}
}
