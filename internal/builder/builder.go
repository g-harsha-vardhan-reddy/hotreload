package builder

import (
	"log"
	"os/exec"
)

func Build(command string) error {

	cmd := exec.Command("cmd", "/C", command)

	cmd.Stdout = nil
	cmd.Stderr = nil

	log.Println("Running build command:", command)

	return cmd.Run()
}
