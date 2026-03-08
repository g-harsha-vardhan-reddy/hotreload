package process

import (
	"os"
	"os/exec"
)

type Server struct {
	Cmd *exec.Cmd
}

func (s *Server) Start(command string) error {

	s.Cmd = exec.Command("cmd", "/C", command)

	s.Cmd.Stdout = os.Stdout
	s.Cmd.Stderr = os.Stderr

	return s.Cmd.Start()
}

func (s *Server) Stop() {

	if s.Cmd != nil && s.Cmd.Process != nil {
		s.Cmd.Process.Kill()
	}
}
