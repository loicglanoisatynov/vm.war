package cmdops

import (
	"bytes"
	"fmt"
	"os/exec"
	"vmwar/server/vars"
	"vmwar/server/vars/logs"
)

func ExecuteCommand(command string) string {
	logs.Throw("Executing command: ", command, nil)

	cmd := exec.Command(vars.Get_hypervisor_path(), "list", "vms")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		fmt.Println("could not run command: ", err)
	}

	return out.String()
}
