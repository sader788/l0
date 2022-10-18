package main

import (
	"github.com/nats-io/stan.go"
	"io"
	"log"
	"os"
)

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func main() {
	dir, err := os.Open("publisher/test")
	if err != nil {
		log.Fatal(err)
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	sc, err := stan.Connect("test-cluster", "client")

	for _, f := range files {
		file, err := os.Open("publisher/test/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		sending, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		sc.Publish("json-notification", sending)
		file.Close()
	}

	sc.Close()
}
