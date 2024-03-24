package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	//pid, err = os.StartProcess("/bin/ps", []string{"ps", "-e", "-opid,ppid,comm"}, procAttr)
	if err != nil {
		log.Fatalf("Error %v, starting process! ", err)
	}
	fmt.Printf("The pid is: %v", pid)

	cmd := exec.Command("ls", "-l")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error %v, exec command! ", err)
	}
	fmt.Printf("The command is %v", cmd)
}
