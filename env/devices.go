package env

import (
	"fmt"
	"sort"
	"vtsploit/pkg"

	"github.com/ctrsploit/sploit-spec/pkg/env/vt"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

func Devices() (result printer.Interface) {

	var info vt.Basic

	pkg.ParseLspci(&info)
	pkg.ParseLsmod(&info)
	pkg.ParseLshw(&info)
	pkg.EnumPort(&info)
	pkg.ParseFile(&info)

	sort.Slice(info.DevList , func(i, j int ) bool  {
		return info.DevList[i].Type < info.DevList[j].Type
	})

	if pkg.ContainsDev(info, "vfio-nvme") {
		pkg.RemoveDev(&info, vt.Disk, "nvme")
	}

	list := "\n"
	for _ , it := range info.DevList {
		list += fmt.Sprintf("\tType: %d, Name: %s\n",rune(it.Type), it.Name )
	}

	result = item.Short{
		Name:        "devices",
		Description: "",
		Result:      list,
	}

	return
}
