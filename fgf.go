package main

import (
	"os"
)

const defaultWorkspace = "./"

func main() {
	instance, err := newFGFInstance()
	if err != nil {
		panic(err)
	}

	if err = instance.sync(); err != nil {
		panic(err)
	}

	if len(os.Args) > 1 { // MODE: MANUAL
		action := os.Args[1]

		if err = executeManualMode(defaultWorkspace, instance, action, os.Args[2:]...); err != nil {
			panic(err)
		}
	}

	// MODE: AUTO
	if err = executeAutoMode(defaultWorkspace, instance); err != nil {
		panic(err)
	}
}
