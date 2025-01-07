package vbox

import (
	"fmt"
	"os/exec"
	"strings"
	"vmwar/cmdops"
	"vmwar/vars"
	"vmwar/virtual_ops/vm"
	"vmwar/virtual_ops/vm/vm_templates"
)

func Create_VM_in_VBox(vmtemplate *vm_templates.VM_template) {
	cmdops.ExecuteCommand("VBoxManage" + "createvm" + "--name" + vmtemplate.Name + "--ostype" + vmtemplate.OStype + "--register")
}

func checkVMExists() bool {
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

func removePreviousVM() {
	exec.Command(vars.Get_hypervisor_path(), "unregistervm", vm.Get_VM_Name(), "--delete")
}

func launchVM() {
	cmd := exec.Command(vars.Get_hypervisor_path(), "startvm", vm.Get_VM_Name())
	if err := cmd.Run(); err != nil {
		fmt.Println("could not run command: ", err)
	}
}
