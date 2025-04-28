package api

import (
	"fmt"
	"net/http"
	vnets "vmwar/server/virtual_ops/vnets"
)

func Get_v_networks(w http.ResponseWriter, r *http.Request) {
	for _, vnet := range vnets.Load_vnets_from_vbox() {
		fmt.Fprintf(w, "%s\n", vnet.Get_v_network())
	}
}
