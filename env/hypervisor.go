package env

import (
	"vtsploit/pkg"

	"github.com/ctrsploit/sploit-spec/pkg/env/vt"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

func Hypervisor() (result printer.Interface) {

	var info vt.Basic
	pkg.ParseFile(&info)
	pkg.ParseDemsg(&info)
	pkg.ParseLscpu(&info)

	hyper_type := "Other"

	if info.HyperType != ""  {
		hyper_type = info.HyperType
	} else {
		if len(pkg.EnumFile("/dev/kvm")) > 0 {
			hyper_type = "KVM"
		} else if len(pkg.EnumFile("/dev/kvm")) > 0  {
			hyper_type = "XEN"
		}
	}

	result = item.Short{
		Name:        "hypervisor",
		Description: "second of current time",
		Result:      hyper_type,
	}
	return result
}
