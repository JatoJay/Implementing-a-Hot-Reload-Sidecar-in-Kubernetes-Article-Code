package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = watcher.Add("/etc/config")
	if err != nil {
		log.Fatal(err)
	}

	for {

		select {

		case event := <-watcher.Events:
			log.Println("Config changed:", event)

		case err := <-watcher.Errors:
			log.Println(err)

		}

	}

} 
