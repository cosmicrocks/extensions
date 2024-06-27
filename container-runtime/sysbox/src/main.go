package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	// print the PATH
	fmt.Println(os.Getenv("PATH"))
	// create symlink in /bin/sh to /usr/local/bin/sh
	if err := os.Symlink("/usr/local/bin/sh", "/bin/sh"); err != nil {
		log.Fatalf("failed to create symlink: %v", err)
	}
	// run the executable /usr/local/bin/sysbox-mgr without args and print to stdout
	cmd := "/usr/local/bin/sysbox-mgr"
	args := []string{}
	if err := syscall.Exec(cmd, args, os.Environ()); err != nil {
		log.Fatalf("failed to exec %s: %v", cmd, err)
	}
}
