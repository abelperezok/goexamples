package main

import (
	"fmt"
)

type customErr struct {
	Type  string
	Fatal bool
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Error type:%v, is fatal:%v", ce.Type, ce.Fatal)
}

func foo(err error) {
	fmt.Println(err)
}

func main() {

	ce := customErr{
		Type:  "Too Bad Error",
		Fatal: true,
	}
	foo(ce)
}
