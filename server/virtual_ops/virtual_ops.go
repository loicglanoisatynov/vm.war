package virtual_ops

import (
	"vmwar/server/vars/logs"
	"vmwar/server/virtual_ops/vbox"
)

var hypervisor string
var default_hypervisor string = "vbox"
var available_hypervisors = []string{"vbox"}

func Designate_hypervisor(designated_hypervisor string) {
	for _, h := range available_hypervisors {
		if h == designated_hypervisor {
			hypervisor = designated_hypervisor
			return
		} else {
			logs.Throw("virtual_ops.designate_hypervisor()", hypervisor+" could not be init. Switching to default hypervisor (VirtualBox)", nil)
			hypervisor = default_hypervisor
		}
	}
}

func Get_hypervisor() string {
	return hypervisor
}

func Set_hypervisor_path(hypervisor string) {
	switch hypervisor {
	case "vbox":
		vbox.Get_vbox_path()
	}
}
