package main

import (
	"os/exec"
)

// CmdExec function to perform external xcommand execution
func CmdExec(args ...string) ([]byte, error) {

	baseCmd := args[0]
	cmdArgs := args[1:]

	cmd := exec.Command(baseCmd, cmdArgs...)
	out, err := cmd.Output()
	if err != nil {
		return []byte(""), err
	}

	return out, nil
}

func main() {
	cnp, err := CmdExec("kubectl", "get", "cnp", "--all-namespaces", "-o", "name")
	val := ""
	var val2 []string
	if err == nil {
		for i := 0; i < len(cnp); i++ {
			if string(cnp[i]) != "\n" {
				val += string(cnp[i])
			} else {
				val2 = append(val2, val)
				val = ""
			}

		}
		for i := 0; i < len(val2); i++ {
			println()
			println(i+1, " --> ", val2[i])
			desc, _ := CmdExec("kubectl", "describe", val2[i])
			println()
			println(string(desc))
		}

	}
}
