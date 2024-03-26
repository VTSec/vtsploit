package env

import (
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
	"vtsploit/pkg"
)

func Version() (result printer.Interface) {
	out, err := pkg.RunCmd("dmidecode",[]string {"-s","chassis-version"})
	if err != nil {
		out = err.Error()
	}
	result = item.Short{
		Name:        "version",
		Description: "qemu verion",
		Result:      out,
	}
	return result
}
