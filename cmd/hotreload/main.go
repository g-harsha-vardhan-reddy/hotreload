package main

import (
	"flag"
	"fmt"

	"hotreload/internal/builder"
	"hotreload/internal/process"
	"hotreload/internal/watcher"
)

func main() {

	root := flag.String("root", ".", "project root")
	buildCmd := flag.String("build", "", "build command")
	execCmd := flag.String("exec", "", "run command")

	flag.Parse()

	changes := make(chan bool)

	err := watcher.StartWatcher(*root, changes)
	if err != nil {
		fmt.Println("Watcher error:", err)
		return
	}

	server := &process.Server{}

	fmt.Println("Watching folder:", *root)

	// First build when tool starts
	err = builder.Build(*buildCmd)
	if err != nil {
		fmt.Println("Initial build failed:", err)
	} else {
		server.Start(*execCmd)
	}

	for {

		<-changes

		fmt.Println("File change detected")

		server.Stop()

		err := builder.Build(*buildCmd)
		if err != nil {
			fmt.Println("Build failed:", err)
			continue
		}

		fmt.Println("Build successful")

		server.Start(*execCmd)

	}

}
