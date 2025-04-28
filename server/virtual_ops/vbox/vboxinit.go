package vbox

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"vmwar/server/vars"
)

func Get_vbox_path() {
	var cmd string
	extension := ""
	if runtime.GOOS == "windows" {
		cmd = "where"
		extension = ".exe"
	} else {
		cmd = "which"
	}
	if os.Getenv("WSL_DISTRO_NAME") != "" {
		extension = ".exe"
	}
	var output []byte
	var err error

	output, err = exec.Command(cmd, "VBoxManage"+extension).Output()

	if err != nil {
		fmt.Println("VirtualBox is not installed:", err)
		fmt.Println("Please install VirtualBox and try again.")
		os.Exit(1)
	}
	vars.Set_hypervisor_path(strings.TrimSpace(string(output)))

	fmt.Println("VBoxManage installation found at:", vars.Get_hypervisor_path())
}
