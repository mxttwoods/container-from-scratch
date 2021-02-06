package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// main cli controller
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("arg not recognized")
	}
}

// main container runtime
func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2], os.Args[3:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	// run
	must(cmd.Run())
}

func child() {
	fmt.Printf("running %v as PID %d\n", os.Args[2:], os.Getpid())
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	// create root
	must(syscall.Chroot("home/matthew"))
	// cd into root
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	// run
	must(cmd.Run())
}

// catch errors
func must(err error) {
	if err != nil {
		panic(err)
	}
}
