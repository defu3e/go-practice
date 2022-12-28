package storage

import (
	log "cityService/pkg/glog"
	"cityService/pkg/storage"
	f "fmt"
	"os"
	"os/signal"
)

var (
	glog = log.Init()
)

func CheckInterrupt(store *storage.Store) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		Interrupted(store)
	}()
}

func Interrupted(store *storage.Store) {
	f.Println("\nЗавершение работы программы...")
	glog.Println("Поступил сигнал прерывания...")
	store.Close()
	glog.FileHandle.Close()
	os.Exit(1)
}
