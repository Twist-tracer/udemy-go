package main

import (
	"os"
)

type logWriter struct{}

func main() {
	var logger logWriter

	_, err := logger.Write([]byte("Обана"))

	if err != nil {
		os.Exit(1)
	}
}

func (logWriter) Write(bs []byte) (int, error) {
	return os.Stdout.Write(bs)
}
