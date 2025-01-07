package vm_templates

import "fmt"

type VM_template struct {
	Name     string `json:"--name="`                  // VMwar-Client
	Arch     string `json:"--platform-architecture="` // x86
	OStype   string `json:"--ostype="`
	ISO      string `json:"--iso="`
	VboxFile string `json:"--vboxfile="` // Corresponds to the path of the VBoxManage executable
}

type Options struct {
	Name   string
	Arch   string
	OStype string
	ISO    string
}

type VM_Interface interface {
	Create()
}

func Create(vmOptions *Options) VM_template {
	fmt.Println("Creating VM...")
	var vm VM_template
	vm.Name = vmOptions.Name
	vm.Arch = vmOptions.Arch
	vm.OStype = vmOptions.OStype
	vm.ISO = vmOptions.ISO
	return vm
}
