package env

import (
	"fmt"
	"sort"
	"time"
	"vtsploit/pkg"

	"github.com/ctrsploit/sploit-spec/pkg/env/vt"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

func Devices() (result printer.Interface) {

	info := vt.Basic {}

	pkg.ParseLspci(info)

	sort.Slice(info , func(i, j int ) bool  {
		return info.DevList[i].Type < info.DevList[j].Type
	})

	list := ""
	for _ , it := range info.DevList {
		list += fmt.Sprintf("Type: %d, Name: %s\n",rune(it.Type), it.Name )
	}

	result = item.Short{
		Name:        "devices",
		Description: "show device list",
		Result:      list,
	}


	return
}
