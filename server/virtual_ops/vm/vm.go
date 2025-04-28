package vm

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"vmwar/server/vars"
	"vmwar/server/virtual_ops/vnets"
)

type Vm struct {
	id   int
	name string
	ip   string
}

func (vm Vm) Get_VM_Name() string {
	return vm.name
}

func (vm Vm) Set_VM_Name(new_vm_name string) {
	vm.name = new_vm_name
}

func Create_vms(vnet_name string) error {
	fmt.Println("func Create_vms")
	// Create VMs
	vm1 := Vm{
		// id:
		name: vnet_name + "_vm1",
		ip:   "10.0.2.1",
	}
	vm2 := Vm{
		name: vnet_name + "_vm2",
		ip:   "10.0.2.2",
	}
	fmt.Println("VM1 name: ", vm1.name)
	fmt.Println("VM2 name: ", vm2.name)

	fmt.Println("Preparing to launch commands...")

	bytes_copied, err := copy("./server/virtual_ops/vm/Metasploitable3-ub1404-Origin.ova", "./VMs/"+vm1.name+".ova")
	if err != nil {
		fmt.Println("Error copying file:", err)
	} else {
		fmt.Printf("Copied %d bytes from %s to %s\n", bytes_copied, "./server/virtual_ops/vm/Metasploitable3-ub1404-Origin.ova", "./VMs/"+vm1.name+".ova")
	}
	bytes_copied, err = copy("./server/virtual_ops/vm/Metasploitable3-ub1404-Origin.ova", "./VMs/"+vm2.name+".ova")
	if err != nil {
		fmt.Println("Error copying file:", err)
	} else {
		fmt.Printf("Copied %d bytes from %s to %s\n", bytes_copied, "./server/virtual_ops/vm/Metasploitable3-ub1404-Origin.ova", "./VMs/"+vm1.name+".ova")
	}

	fmt.Println("VMs copied")
	fmt.Println("Importing VMs...")
	// var output []byte

	// Créer le répertoire de destination s'il n'existe pas
	if _, err := os.Stat("./VMs/" + vm1.name); os.IsNotExist(err) {
		err = os.MkdirAll("./VMs/"+vm1.name, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
		}
	}
	if _, err := os.Stat("./VMs/" + vm2.name); os.IsNotExist(err) {
		err = os.MkdirAll("./VMs/"+vm2.name, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
		}
	}

	_, err = exec.Command("VBoxManage", "import", "./VMs/"+vm1.name+".ova", "--vsys", "0", "--vmname", vm1.name, "--basefolder", "./VMs/"+vm1.name+"/").Output()
	// _, err = exec.Command("VBoxManage", "import", "./VMs/"+vm1.name+".ova").Output()
	if err != nil {
		fmt.Println("Error importing VM:", err)
	}

	_, err = exec.Command("VBoxManage", "import", "./VMs/"+vm2.name+".ova", "--vsys", "0", "--vmname", vm2.name).Output()
	// _, err = exec.Command("VBoxManage", "import", "./VMs/"+vm2.name+".ova").Output()
	if err != nil {
		fmt.Println("Error importing VM:", err)
	}

	fmt.Println("VMs imported")
	return nil
}

func copy(src, dst string) (int64, error) {
	// Si le répertoire de destination n'existe pas, le créer
	repository := "./VMs/"
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		err = os.MkdirAll(repository, 0755)
		if err != nil {
			return 0, err
		}
	}

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func Load_vms_from_vbox() []Vm {
	var vms []Vm
	cmd, err := exec.Command(vars.Get_hypervisor_path(), "list", "vms").Output()
	if err != nil {
		fmt.Println("Error loading VMs:", err)
		return nil
	}

	vms_list := strings.Split(string(cmd), "\n")
	for _, vm := range vms_list {
		if strings.Contains(vm, "_vm") {
			vm_name := strings.Split(vm, "\"")[1]
			vm_id := strings.Split(vm, "{")[1]
			vm_id = strings.Split(vm_id, "}")[0]
			vm_id = strings.TrimSpace(vm_id)
			vm_name = strings.TrimSpace(vm_name)
			vms = append(vms, Vm{
				// id:   0,
				name: vm_name,
				// ip:   vm_id,
			})
			fmt.Println("VM name: ", vm_name)
		}
	}
	return vms
}

func Wipe_vms() {
	// TODO effacer les VMs avec VBoxManage
	vnets := vnets.Load_vnets_from_vbox()
	vms := Load_vms_from_vbox()
	fmt.Println("vms: ", vms)
	fmt.Println("vnets: ", vnets)
	fmt.Println("Deleting VMs...")
	for _, vm := range vms {
		fmt.Println("Deleting VMs for vnet: ", vm.name)
		fmt.Println("Deleting VMs: ", vm.name, " and ", vm.name)
		_, err := exec.Command("VBoxManage", "controlvm", vm.name, "poweroff").Output()
		if err != nil {
			fmt.Println("Error deleting VM:", err)
		}
		_, err = exec.Command("VBoxManage", "controlvm", vm.name, "poweroff").Output()
		if err != nil {
			fmt.Println("Error deleting VM:", err)
		}
		_, err = exec.Command("VBoxManage", "unregistervm", vm.name, "--delete").Output()
		if err != nil {
			fmt.Println("Error deleting VM:", err)
		}
		_, err = exec.Command("VBoxManage", "unregistervm", vm.name, "--delete").Output()
		if err != nil {
			fmt.Println("Error deleting VM:", err)
		}
		fmt.Println("VMs deleted")
	}
	// TODO effacer les fichiers de VMs

	files, err := os.ReadDir("./VMs/")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		err := os.Remove("./VMs/" + file.Name())
		if err != nil {
			fmt.Println("Error deleting file:", err)
		}
	}
	// Supprimer les sous-répertoires du répertoire VMs
	for _, file := range files {
		if file.IsDir() {
			err := os.RemoveAll("./VMs/" + file.Name())
			if err != nil {
				fmt.Println("Error deleting directory:", err)
			}
		}
	}

	fmt.Println("All VMs deleted")
	fmt.Println("All VMs files deleted")
}
