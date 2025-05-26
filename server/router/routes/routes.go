package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	vms "vmwar/server/virtual_ops/vm"
	vnets "vmwar/server/virtual_ops/vnets"
)

type Get_vnet_request struct {
	Username string `json:"username"`
}

type Post_vnet_request struct {
	Username string `json:"username"`
}

func Test_Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test Handler")
}

// Créée un réseau virtuel
func Post_v_network(w http.ResponseWriter, r *http.Request) {
	var request Post_vnet_request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Fprintln(w, "Missing name field in JSON ?\ncurl example : curl -X POST -H \"Content-Type: application/json\" localhost:8080/v-network -d '{\"username\":\"test\"}'")
		return
	}

	if strings.Contains(vnets.Get_vnets(), request.Username) {
		fmt.Fprintln(w, "V-network already exists")
		return
	}

	Vnet := vnets.Create_v_network(request.Username)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, "V-network created")

	fmt.Println("V-network created : ", request.Username)
	fmt.Println("Creating VMs...")

	err = vms.Create_vms(request.Username, Vnet.Name)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "V-network created",
		"key":     Vnet.Key,
	})
}

func Delete_v_network(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete v-network")
	var request Get_vnet_request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Fprintln(w, "Missing name field in JSON ?\ncurl example : curl -X DELETE -H \"Content-Type: application/json\" localhost:8080/v-network -d '{\"username\":\"test\"}'")
		return
	}

	vnet_id := request.Username
	fmt.Println("Delete v-network : ", vnet_id)

	err = vnets.Delete_v_network(vnet_id)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, "V-network deleted")
	fmt.Println("V-network deleted : ", vnet_id)
}

func Get_v_network(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get v-network")
	var request Get_vnet_request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Fprintln(w, "Missing name field in JSON ?\ncurl example : curl -X GET -H \"Content-Type: application/json\" localhost:8080/v-network -d '{\"username\":\"test\"}'")
		return
	}

	vnet_id := request.Username
	fmt.Println("Get v-network : ", vnet_id)

}

func Get_v_networks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get v-networks")
	json.NewEncoder(w).Encode(vnets.Get_vnets())
}
