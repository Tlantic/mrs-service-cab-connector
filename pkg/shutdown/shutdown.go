package shutdown

import (
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
)

// Shutdown ...
type Shutdown struct {
	InitiateChannel chan bool
	DoneChannel     chan bool
	timeout         time.Duration
	Logger          *log.Entry
}

// NewShutdown ...
func NewShutdown(timeout time.Duration, l *log.Entry) *Shutdown {
	return &Shutdown{
		InitiateChannel: make(chan bool, 1),
		DoneChannel:     make(chan bool, 1),
		timeout:         timeout,
		Logger:          l,
	}
}

// WaitForSignal ...
func (sd *Shutdown) WaitForSignal() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	<-signalChannel

	sd.Logger.Info("received interrupt signal")
	sd.InitiateChannel <- true
	select {
	case <-signalChannel:
		sd.Logger.Warning("forcing shutdown")
		os.Exit(1)
	case <-sd.DoneChannel:
		sd.Logger.Info("cleanup successful, exiting")
	case <-time.After(time.Second * sd.timeout):
		sd.Logger.Info("cleanup timed out, exiting")
	}
}
