package main

import (
	"flag"
	"fmt"
	"testing"
)

var (
	job  Job
	name string
)

func init() {
	flag.Var(&job, "job", "Help")
}

func TestFlagParse(t *testing.T) {
	//flag.Parse() // auto parse in testing case
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "Go Lang", "Help")
	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "n", "PHP Lang", "Help")

	args := flag.Args()
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "php":
		_ = phpCmd.Parse(args[1:])
	}

	fmt.Printf("name: %s\n", name)
	fmt.Printf("%s\n", job)
}
