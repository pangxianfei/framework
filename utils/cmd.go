package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

// ExecShell执行shell命令
func ExecShell(shell string) (stdout string, stderr string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", shell)
	} else {
		cmd = exec.Command("/bin/bash", "-c", shell)
	}

	var out bytes.Buffer
	var stderrBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderrBuf
	err := cmd.Run()
	if err != nil {
		fmt.Println("shell exec have an error", "err", err)
	}

	return out.String(), stderrBuf.String()
}
