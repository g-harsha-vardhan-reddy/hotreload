package watcher

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func StartWatcher(root string, changes chan bool) error {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	err = watcher.Add(root)
	if err != nil {
		return err
	}

	go func() {

		for {

			select {

			case event := <-watcher.Events:

				log.Println("File changed:", event.Name)
				changes <- true

			case err := <-watcher.Errors:

				log.Println("Error:", err)
			}

		}

	}()

	return nil
}
