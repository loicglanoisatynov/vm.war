package vm

var vm_name string = ""

func Get_VM_Name() string {
	return vm_name
}

func Set_VM_Name(new_vm_name string) {
	vm_name = new_vm_name
}

func Init_vm_name(new_vm_name string) {
	if new_vm_name == "" {
		vm_name = "VMwar-Client"
	}
}
