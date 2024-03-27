package pkg

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/ctrsploit/sploit-spec/pkg/env/vt"
)

// 字符包含
func MatchKey( info * vt.Basic , data string ) {
	data = strings.ToLower(data)
    for _, it := range DeviceKeyList {
        for k, v := range it.Pair {
            if strings.Contains(data, strings.ToLower(k)) {
                AppendDev(info, it.Type, v)
            }
        }
    }
}

// 正则匹配
func MatchRgx( info * vt.Basic , data string  ) {
	data = strings.ToLower(data)
	re := regexp.MustCompile(`virtio_\w+`)
	result := re.FindAllString(data, -1)
	for _, r := range result {
		for _, it := range DeviceKeyRgx {
			for k, v := range it.Pair {
				if k == r {
					AppendDev(info, vt.Other, v)
				}
			}
		}
	}
}

func ParseLspci( info * vt.Basic ) error {
	out, err := RunCmd("lspci", []string { "-n" })
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`[a-z0-9]{2}:[a-z0-9]{2}\.[a-z0-9]{1}`)
	devId := re.FindAllString(out, -1)

	for _, id := range devId {
		out, err = RunCmd("lspci", []string{"-vv", "-s", string(id)})
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
			AppendDev(info, vt.Display, "vfio-gpu")
		} else if strings.Contains(r, "xensource") {
			info.HyperType = "xen"
		} else {
			re3 := regexp.MustCompile(`kernel driver in use: (.+?)\n`)
			result = re3.FindAllString(out, -1)
			if len(result) > 1 {
				r2 := result[0]
				if strings.Contains(r2, "devdrv_device_driver") {
					AppendDev(info, vt.Display, "vfio-ascend")
				} else {
					AppendDev(info, vt.Other, "vfio-" + r2)
				}
			} else {
				AppendDev(info, vt.Other, r)
				fmt.Printf("Unknown device: %s", r)
			}
		} 

	}
	return nil
}

func ParseLsmod( info * vt.Basic ) error {
	out, err := RunCmd("lsmod", []string{})
	if err != nil {
		return err
	}

	MatchKey(info, out)
	MatchRgx(info, out)

	return nil
}

func ParseLshw( info * vt.Basic ) error {
	out, err := RunCmd("lshw", []string{})
	if err != nil {
		return err
	}

	MatchKey(info, out)
	MatchRgx(info, out)
	return nil
}

func ParseDemsg( info * vt.Basic ) error {
	out, err := RunCmd("dmesg", []string{})
	if err != nil {
		return err
	}

	s := regexp.MustCompile(`Hypervisor detected: (.+?)\n`).FindAllStringSubmatch(out, 1)
	if len(s) > 0 {
		info.HyperType = Trim(s[0][1])
	}
	
	return nil
}

func ParseLscpu( info * vt.Basic ) error {
	out, err := RunCmd("lscpu", []string{})
	if err != nil {
		return err
	}

	s := regexp.MustCompile(`Hypervisor vendor: (.+?)\n`).FindAllStringSubmatch(out, 1)
	if len(s) > 0 {
		info.HyperType = Trim(s[0][1])
	}

	return nil
}