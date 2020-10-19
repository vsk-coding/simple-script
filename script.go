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

// nameSpaceReturn function returns the name and namespace from the log
func nameSpaceReturn(args []byte, i int, j int) (string, string) {

	nameSpace := ""
	name := ""
	temp := ""
	for ; j < i; j++ {
		for string(args[j]) != " " {
			if string(args[j]) != "\n" {
				temp += string(args[j])

			}
			j++
		}
		if string(args[j]) == " " {
			for string(args[j]) == " " {
				j++
			}
			nameSpace = temp
			break
		}
	}
	temp = ""
	for ; j < i; j++ {
		for string(args[j]) != " " {
			temp += string(args[j])
			j++
		}
		if string(args[j]) == " " {
			name = temp
			break
		}
	}
	return name, nameSpace
}

func main() {
	get := "get"
	cnp, err := CmdExec("kubectl", get, "cnp", "--all-namespaces")

	nameSpace := ""
	name := ""
	count := 0
	if err == nil {
		ruleNo := 1
		for i := 0; i < len(cnp); i++ {

			if string(cnp[i]) == "\n" {
				name, nameSpace = nameSpaceReturn(cnp, i, count)

				count = i
				rules, erro := CmdExec("kubectl", "describe", "cnp", name, "--namespace", nameSpace)
				if erro == nil {
					println()
					println("Rule ", ruleNo, " ==> ", name)
					println()
					println(string(rules))
					ruleNo++
				}

			}
		}

	}
}
