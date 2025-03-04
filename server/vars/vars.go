package vars

var hypervisor_path string
var verbose_mode int
var dbg_mode int
var log_file string

func Set_hypervisor_path(new_hypervisor_path string) {
	hypervisor_path = new_hypervisor_path
}

func Get_hypervisor_path() string {
	return hypervisor_path
}

func Set_verbose_mode(new_verbose_mode int) {
	verbose_mode = new_verbose_mode
}

func Get_verbose_mode() int {
	return verbose_mode
}

func Set_dbg_mode(new_dbg_mode int) {
	dbg_mode = new_dbg_mode
}

func Get_dbg_mode() int {
	return dbg_mode
}

func Set_logfile(new_log_file string) {
	log_file = new_log_file
}

func Get_logfile() string {
	return log_file
}
