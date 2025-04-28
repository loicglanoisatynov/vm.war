package logs

import (
	"fmt"
	"vmwar/server/vars"
)

func Throw(function_caller string, message string, varsToEnumerate map[string]string) {
	if vars.Get_verbose_mode() <= 1 {
		log_to(function_caller, message)
	}
	if vars.Get_verbose_mode() <= 2 {
		for varName, value := range varsToEnumerate {
			log_to(function_caller, "Variable: "+varName+", Value: "+value)
		}
	}
}

func log_to(caller string, message string) {
	if vars.Get_logfile() != "" {
		log_to_logfile(caller + message)
	}
	if vars.Get_verbose_mode() <= 1 {
		log_to_stdout(caller + message)
	}
}

func log_to_logfile(message string) { // TODO : function not implemented
	// log message to logfile
	fmt.Println("TODO : function not implemented (vmwar/vars/logs.log_to_logfile())")
}

func log_to_stdout(message string) { // log message to stdout
	fmt.Println(message)
}

func log_to_dbgout(message string) { // log message to debug output
	fmt.Println("TODO : function not implemented (vmwar/vars/logs.log_to_dbgout())")
}
