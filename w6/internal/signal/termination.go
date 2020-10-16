package signal

import (
	"os"
	"os/signal"
	"syscall"
)

func Term() {
	var terminationSignals = []os.Signal{
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGHUP,
		syscall.SIGUSR2,
	}

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, terminationSignals...)

	<-stopSignal
}
