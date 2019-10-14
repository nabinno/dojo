package utilities

import (
	"fmt"
	"os/exec"

	"github.com/mattn/go-shellwords"
)

// Cmd ....
func Cmd(cmdstr string) {
	if out, err := OutCmd(cmdstr); err == nil {
		fmt.Println(string(out))
	} else {
		fmt.Println(err)
	}
}

// OutCmd ....
func OutCmd(cmdstr string) (out []byte, err error) {
	args, err := shellwords.Parse(cmdstr)
	if err != nil {
		return
	}

	switch len(args) {
	case 0:
		return
	case 1:
		out, err = exec.Command(args[0]).Output()
	default:
		out, err = exec.Command(args[0], args[1:]...).Output()
	}
	if err != nil {
		return
	}
	return
}
