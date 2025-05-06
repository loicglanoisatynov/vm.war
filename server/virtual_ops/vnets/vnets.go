package vnets

import (
	"bytes"
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"vmwar/server/vars"
	// "vmwar/server/virtual_ops/vm"
)

type Vnet struct {
	Id   int
	Name string
	Ip   string
	Key  string
	Vm1  string
	Vm2  string
}

var Vnets []Vnet

func Wipe_vnets() {
	Vnets = Load_vnets_from_vbox()
	// fmt.Println("Deleting Vnets...")
	for _, vnet := range Vnets {
		Delete_v_network(vnet.Name)
	}
}

func Load_vnets_from_vbox() []Vnet {
	// fmt.Println("Loading Vnets from VirtualBox...")
	cmd, err := exec.Command(vars.Get_hypervisor_path(), "natnetwork", "list").Output()

	if err != nil {
		fmt.Println("could not run command: ", err)
		return nil
	}

	output := strings.Split(string(cmd), "\n")

	vnets := []Vnet{}

	for _, line := range output {
		if strings.Contains(string(line), "vnet_") {
			vnet := Vnet{
				Id:   Get_next_id(),
				Name: strings.ReplaceAll(strings.ReplaceAll(string(line), " ", ""), "Name:", ""),
				Ip:   "10.0.1.0/24",
			}
			vnets = AddVnet(vnets, vnet)
		}
	}
	return vnets
}

// Ajoute un Vnet à la liste des Vnets
func AddVnet(vnets []Vnet, vnet Vnet) []Vnet {
	vnets = append(vnets, vnet)
	return vnets
}

func Get_next_id() int {
	return len(Vnets) + 1
}

func Create_v_network(username string) Vnet {
	vnet := Vnet{
		Id:   Get_next_id(),
		Name: "vnet_" + strconv.Itoa(Get_next_id()) + "_" + username,
		Ip:   "10.0.1.0/24",
		Key:  generate_key(),
	}
	// AddVnet(vnet)

	cmd := exec.Command(vars.Get_hypervisor_path(), "natnetwork", "add", "--enable", "--netname="+vnet.Name, "--network=10.0.2.0/24")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		fmt.Println("could not run command: ", err)
	}

	return vnet
}

func (vnet Vnet) Get_v_network() string {
	return vnet.Name
}

func Get_vnets() string {
	returned := ""
	for _, vnet := range Load_vnets_from_vbox() {
		returned += fmt.Sprintf("%s\n", vnet.Name)
	}
	return returned
}

func Delete_v_network(vnet_name string) error {

	full_vnet_name := vnet_name
	// fmt.Println("Deleting Vnet: ", full_vnet_name)

	_, err := exec.Command(vars.Get_hypervisor_path(), "natnetwork", "remove", "--netname", full_vnet_name).Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	RemoveVnet(full_vnet_name)
	return nil
}

func RemoveVnet(vnet_name string) {
	for i, vnet := range Vnets {
		if vnet.Name == vnet_name {
			Vnets = append(Vnets[:i], Vnets[i+1:]...)
			break
		}
	}
}

func Get_vnet_id(vnet_name string) string {
	vnets := strings.Split(Get_vnets(), "\n")
	for _, vnet := range vnets {
		if strings.Contains(vnet, vnet_name) {
			return strings.Split(vnet, "_")[1]
		}
	}
	return ""
}

func generate_key() string {
	// Génère une clé aléatoire de 6 caractères

	return randStringBytes(6)
}

func randStringBytes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
