package pkg

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ctrsploit/sploit-spec/pkg/env/vt"
)

func EnumPort ( info * vt.Basic ) {
	file, err := os.Open("/proc/ioports")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// 创建一个 Scanner 对象
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容
	for scanner.Scan() {
		// 获取当前行的文本
		line := scanner.Text()

		if strings.Contains(line, "vesafb") {
			AppendDev(info, vt.Display, "qxl")
		}
	}
	// 检查扫描是否发生错误
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}

func ParseFile( info * vt.Basic ) {
    for _, it := range DeviceKeyFile {
        for k, v := range it.Pair {
			fs := EnumFile(k)
			if len(fs) > 0 {
                AppendDev(info, it.Type, v)
			}
        }
    }

	f1 := EnumFile("/dev/virtio-ports/*")
	f2 := EnumFile("/dev/vport*")
	if len(f2) != 0 && len(f2) != len(f1) {
		AppendDev(info, vt.Other, "virtconsole")
	}

	fs := EnumFile("/dev/disk/by-id/*")
	for _, f := range fs {
		f = strings.ToLower(f)
		if strings.Contains(f, "qemu") {
			continue
		}

        // vfio 和 DPDK 两种直通方案？
	 	if regexp.MustCompile(`/dev/disk/by-id/nvme-.+?_`).FindAllString(f, -1) != nil {
			AppendDev(info, vt.Disk, "vfio-nvme")
		} else if ( regexp.MustCompile(`/dev/disk/by-id/ata-.+?_`).FindAllString(f, -1) != nil ) {
			// PCI Passthrough 和 vfio 两种直通方案？
			AppendDev(info, vt.Disk, "vfio-sata")
		}
	}
}