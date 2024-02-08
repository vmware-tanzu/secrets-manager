package logger

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"syscall"
)

var mutex sync.Mutex

func CreateLogger() {
	err := os.Remove(LogPipePath)
	if err != nil {
		log.Printf("Failed to remove %s : %s\n", LogPipePath, err)
	}
	err = syscall.Mkfifo(LogPipePath, 0666)
	if err != nil {
		log.Printf("Failed to create %s : %s\n", LogPipePath, err)
	} else {
		log.Printf("Successfully created %s\n", LogPipePath)
	}
}

func SendLog(a ...any) {
	mutex.Lock()
	defer mutex.Unlock()

	pipeFile, err := os.OpenFile(LogPipePath, os.O_WRONLY, os.ModeNamedPipe)
	if err != nil {
		log.Printf("LogPiper.SendLog error opening file: %v\n", err)
		return
	}
	defer pipeFile.Close()

	logMessage := LogTextBuilder(a...)

	_, err = pipeFile.WriteString(logMessage)
	if err != nil {
		log.Printf("LogPiper.SendLog error writing file: %v\n", err)
	}
}

func ReceiveLog() {
	pipeFile, err := os.OpenFile(LogPipePath, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		log.Printf("LogPiper.ReceiveLog error opening named pipe: %v\n", err)
		return
	}
	defer pipeFile.Close()

	reader := bufio.NewReader(pipeFile)

	for {
		line, err := reader.ReadBytes('\n')
		if err == nil {
			fmt.Print(string(line))
		}
	}
}
