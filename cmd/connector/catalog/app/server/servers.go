package server

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Server interface {
	Run() error
	Stop() error
}

func RunServers(servers ...Server) {
	for _, server := range servers {
		go func(s Server) {
			if err := s.Run(); err != nil {
				logrus.WithError(err).Error("Run server failed")
				os.Exit(1)
			}
		}(server)
	}
}

// StopServers ...
func StopServers(servers ...Server) {
	for _, server := range servers {
		if err := server.Stop(); err != nil {
			logrus.WithError(err).Error("Error at stopping server")
		}
	}
}
