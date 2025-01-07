package verbose

import (
	"fmt"
	"vmwar/vars"
)

func SpeakIf(text string) {
	if vars.GetVerboseMode() <= 1 {
		fmt.Println(text)
	}
	if vars.GetVerboseMode() <= 2 {
		EnumerateVariables(nil)
	}
}

func EnumerateVariables(variables map[string]string) {
	for varName, value := range variables {
		fmt.Printf("Variable: %s, Value: %s\n", varName, value)
	}
}
