package vbox

import (
	"fmt"
	"os/exec"
	"strings"
	"vmwar/server/cmdops"
	"vmwar/server/vars"
	vm_pack "vmwar/server/virtual_ops/vm"
	"vmwar/server/virtual_ops/vm/vm_templates"
)

func Create_VM_in_VBox(vmtemplate *vm_templates.VM_template) {
	cmdops.ExecuteCommand("VBoxManage" + "createvm" + "--name" + vmtemplate.Name + "--ostype" + vmtemplate.OStype + "--register")
}

func checkVMExists(vm vm_pack.Vm) bool {
	vms_list := cmdops.ExecuteCommand("VBoxManage " + "list " + "vms")

	for _, line := range strings.Split(vms_list, "\n") {
		if strings.Contains(line, vm.Get_VM_Name()) {
			fmt.Println("VM already exists")
			return true
		}
	}
	fmt.Println("VM does not exist")
	return false
}

func removePreviousVM(vm vm_pack.Vm) {
	exec.Command(vars.Get_hypervisor_path(), "unregistervm", vm.Get_VM_Name(), "--delete")
}

func launchVM(vm vm_pack.Vm) {
	cmd := exec.Command(vars.Get_hypervisor_path(), "startvm", vm.Get_VM_Name())
	if err := cmd.Run(); err != nil {
		fmt.Println("could not run command: ", err)
	}
}
