package main

import (
	"fmt"
	"io"
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

var green string = "\033[32m"
var blue string = "\033[34m"
var white string = "\033[0m"
var vmName string = "VMwar-Client"

func main() {
	// Nettoie le terminal
	fmt.Print("\033[H\033[2J")
	init_vmwar(os.Args)

	// vnets.Load_vnets_from_vbox()
	vnets.Wipe_vnets()
	vms.Wipe_vms()

	fmt.Println(green + "VMWar booted and ready to go !" + white)
	http.HandleFunc("/", server.Serve)
	http.ListenAndServe(":8080", nil)
}

func init_vmwar(args []string) {
	parse_args(args)
	fmt.Println("Initializing VMwar...")
	fmt.Println("Initialize local environment...")

	vbox.Get_vbox_path()
	// Check if Metasploitable3-ub1404.ova exists in ./server/virtual_ops/vm/
	if _, err := os.Stat("./server/virtual_ops/vm/Metasploitable3-ub1404-Origin.ova"); os.IsNotExist(err) {
		download_sourceforge_file()
	}
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

func printProgress(progress, total int64) {
	percent := float64(progress) / float64(total) * 100
	barlen := 40
	done := int(percent / 100 * float64(barlen))
	fmt.Printf("\r[%-*s] %6.2f%% %d/%d MB", barlen, strings.Repeat("=", done), percent, progress/1024/1024, total/1024/1024)
}

func download_sourceforge_file() {
	fmt.Println(blue + "socengai/main.go:func download_sourceforge_file() starting..." + white)
	url := "https://sourceforge.net/projects/metasploitable3-ub1404upgraded/files/Metasploitable3-ub1404.ova/download"

	// Crée une requête HTTP
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Déterminer le nom du fichier à partir des headers (Content-Disposition)
	filename := "./server/virtual_ops/vm/Metasploitable3-ub1404-Origin.ova"
	fmt.Println(blue + "socengai/main.go:func download_sourceforge_file() 1 filename : " + filename + white)
	// cd := resp.Header.Get("Content-Disposition")
	// if cd != "" {
	// 	_, params, err := mime.ParseMediaType(cd)
	// 	if err == nil && params["filename"] != "" {
	// 		filename = params["filename"]
	// 	}
	// } else {
	// 	parts := strings.Split(url, "/")
	// 	if len(parts) > 0 && parts[len(parts)-1] != "" {
	// 		filename = parts[len(parts)-1]
	// 	}
	// }
	fmt.Println(blue + "socengai/main.go:func download_sourceforge_file() 2 filename : " + filename + white)

	// Crée le fichier local
	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Récupère la taille totale pour la barre de progression
	size := resp.ContentLength
	if size <= 0 {
		fmt.Println("Taille inconnue, téléchargement sans barre de progression...")
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			panic(err)
		}
	} else {
		var written int64
		buffer := make([]byte, 32*1024)
		for {
			nr, er := resp.Body.Read(buffer)
			if nr > 0 {
				nw, ew := out.Write(buffer[0:nr])
				if nw > 0 {
					written += int64(nw)
					printProgress(written, size)
				}
				if ew != nil {
					panic(ew)
				}
				if nr != nw {
					panic("short write")
				}
			}
			if er == io.EOF {
				break
			}
			if er != nil {
				panic(er)
			}
		}
		fmt.Println() // ligne suivante après la barre
	}

	fmt.Println("Fichier téléchargé sous le nom :", filename)
	fmt.Println(blue + "socengai/main.go:func download_sourceforge_file() ending..." + white)
}
