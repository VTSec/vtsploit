package env

import (
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
	"vtsploit/pkg"
)

func Hypervisor() (result printer.Interface) {

	hyper_type := "other"

	if pkg.FindFile("/dev","kvm") {
		hyper_type = "kvm"
	} else if pkg.FindFile("/dev","xen") {
		hyper_type = "xen"
	}

	result = item.Short{
		Name:        "hypervisor",
		Description: "second of current time",
		Result:      hyper_type,
	}
	return result
}
