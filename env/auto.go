package env

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
)

type Result map[string]printer.Interface

func Auto() {
	result := Result{
	}
	fmt.Println(printer.Printer.Print(result))
}
