package pkg

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/ctrsploit/sploit-spec/pkg/env/vt"
)

func appendDev( info vt.Basic, Type vt.DeviceType, name string ) ( vt.Basic ) {
	for _, it := range info.DevList {
		if it.Type == Type && it.Name == name {
			return info
		}
	}
	info.DevList = append(info.DevList, vt.DeviceInfo { Type: Type, Name: name} )

	return info
}

func ParseLspci( info vt.Basic ) error {
	out, err := RunCmd("lspci", []string { "-n" })
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`[a-z0-9]{2}:[a-z0-9]{2}\.[a-z0-9]{1}`)
	devId := re.FindAllString(out, -1)

	for _, id := range devId {
		out, err = RunCmd("lspci", []string{"-vv", "-s", string(id)});
		if err != nil {
			continue
		}
		re2 := regexp.MustCompile(`subsystem: (.+?)\n`)
		result := re2.FindAllString(out, -1)
		if len(result) < 1 {
			continue
		}

		r := result[0]

		if !strings.Contains(r, "red hat") {
			continue
		}

		if strings.Contains(r, "tesla") {
			appendDev(info, vt.Display, "vfio-gpu")
		} else if strings.Contains(r, "xensource") {
			info.HyperType = "xen"
		} else {
			re3 := regexp.MustCompile(`kernel driver in use: (.+?)\n`)
			result = re3.FindAllString(out, -1)
			if len(result) > 1 {
				r2 := result[0]
				if strings.Contains(r2, "devdrv_device_driver") {
					appendDev(info, vt.Display, "vfio-ascend")
				} else {
					appendDev(info, vt.Other, "vfio-" + r2)
				}
			} else {
				appendDev(info, vt.Other, r)
				fmt.Printf("Unknown device: %s", r)
			}
		} 

	}
	return nil
}