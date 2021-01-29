package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main(){
	switch os.Args[1] {
	case "run":
			run()
	default:
			panic("arg not recognized")
	}
}

func run(){
	fmt.Printf("running %v\n", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd

	must(cmd.Run())
}

func must(err error){
	if err != nil {
		panic(err)
	}
}