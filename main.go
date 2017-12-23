package main

import (
	"Anvil/parser"
	"Anvil/system"
	"fmt"
)

func anvil_init(opt system.Options) {
	fmt.Printf("Starting Anvil...\n")
}

func anvil_cleanup() {
	fmt.Printf("Shutting down Anvil...\n")
}

func process_scene_desc(filenames []string) {
	if len(filenames) == 0 {
		// read scene desc from standard input
		parser.ParseFile("-")
	} else {
		// parse files
		for _, file := range filenames {
			if !parser.ParseFile(file) {
				system.Error(file + " could not be parsed!")
			}
		}
	}
}

func main() {
	opt := system.Options{Desc: "Anvil Options"}
	filenames := make([]string, 0)

	//process command line args
	anvil_init(opt)

	process_scene_desc(filenames)

	anvil_cleanup()
}
