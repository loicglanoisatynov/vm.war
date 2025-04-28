package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	server "vmwar/server"
	"vmwar/server/vars"
	"vmwar/server/virtual_ops"
	"vmwar/server/virtual_ops/vbox"
	vms "vmwar/server/virtual_ops/vm"
	vnets "vmwar/server/virtual_ops/vnets"
)

var vmName string = "VMwar-Client"

func main() {
	init_vmwar(os.Args)

	// vnets.Load_vnets_from_vbox()
	vnets.Wipe_vnets()
	vms.Wipe_vms()

	http.HandleFunc("/", server.Serve)
	http.ListenAndServe(":8080", nil)
}

func init_vmwar(args []string) {
	parse_args(args)
	fmt.Println("Initializing VMwar...")
	fmt.Println("Initialize local environment...")

	vbox.Get_vbox_path()
}

func parse_args(args []string) {
	for arg := range args {
		if args[arg] == "-d" || args[arg] == "--debug" {
			vars.Set_dbg_mode(1)
		}
		if args[arg] == "-v" || args[arg] == "--verbose" {
			vars.Set_verbose_mode(1)
		}
		if args[arg] == "-vv" || args[arg] == "--very-verbose" {
			vars.Set_verbose_mode(2)
		}
		if args[arg] == "-vvv" || args[arg] == "--debug-verbose" {
			vars.Set_verbose_mode(3)
		}
		if args[arg] == "-h" || args[arg] == "--help" {
			fmt.Println("TODO: print help")
		}
		if strings.Contains(args[arg], "--hypervisor=") {
			virtual_ops.Designate_hypervisor(strings.Split(args[arg], "=")[1])
			virtual_ops.Set_hypervisor_path(virtual_ops.Get_hypervisor())
		}
	}
}
