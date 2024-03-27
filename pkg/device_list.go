package pkg

import "github.com/ctrsploit/sploit-spec/pkg/env/vt"

type DeviceKeys struct {
	Type vt.DeviceType 
	Pair map[string]string
}

var DeviceKeyList = []DeviceKeys{ 
	DeviceKeys{ Type: vt.Disk,  Pair: map[string]string { 
		"virtio block device":"virtio_blk", "virtio scsi":"virtio_scsi",  "qemu harddisk":"ide-hd/scsi-hd", "dvd reader":"ide-cd", "QEMU CD-ROM":"scsi-cd", "vbs fileio":"huawei_vbs", 
		"non-volatile memory":"nvme", "dimm memory":"nvdimm/pc-dimm", "nvdimm":"nvdimm", "scsi disk":"scsi-hd/scsi-block",
		}}, 
	DeviceKeys{ Type: vt.Display,  Pair: map[string]string { 
		"product: qxl":"qxl", "cirrus":"cirrus", "pcivnic":"vfio-ascend", 
		}}, 
	DeviceKeys{ Type: vt.Network,  Pair: map[string]string { 
		"product: 8254":"e1000", "product: 8257":"e1000e", "virtio_net":"virtio_net", "virtio network":"virtio_net",
		}}, 
	DeviceKeys{ Type: vt.USB,  Pair: map[string]string { 
		"uhci_hcd":"usb_uhci", "xhci_hcd":"usb_xhci", "ehci_hcd":"usb_ehci", "usb tablet":"usb-tablet",
		}}, 
	DeviceKeys{ Type: vt.Other,  Pair: map[string]string {  // fix Audio
		"snd_hda_intel":"intel-hda/hda-duplex",
		}}, 
	DeviceKeys{ Type: vt.Other,  Pair: map[string]string {  
		"vfio-based":"vfio", "9pnet":"virtio_9p",  "virtio rng":"virtio_rng",
		// "virtio console":"virtserialport/virtconsole", // 可通过/dev确定
		}}, 
}

var DeviceKeyRgx = []DeviceKeys{ 
	DeviceKeys{ Type: vt.Other,  Pair: map[string]string { 
		"virtio_console":"virtio_console", "virtio_pci":"virtio_pci",  "virtio_ring":"virtio_ring",
		}}, 
}

var DeviceKeyFile = []DeviceKeys{ 
	DeviceKeys{ Type: vt.Disk,  Pair: map[string]string { 
		"/dev/nvme*":"nvme", "/dev/dvd":"ide-cd", "/dev/disk/by-id/*HARDDISK*":"ide-hd", "/dev/vda*":"virtio_blk", "/dev/sda*":"virtio_scsi" ,
		}}, 
	DeviceKeys{ Type: vt.Disk,  Pair: map[string]string { 
		"/dev/virtio-ports/*spice*":"spice",
		}}, 
	DeviceKeys{ Type: vt.Other,  Pair: map[string]string { 
		"/dev/virtio-ports/*":"virtserialport",
		}}, 
}
	