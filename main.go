package main

import (
	"sys"
	"os"
	"net/http"
)

func main() {
	done := make(chan int, 1)
	go func() {
		signal := make(chan os.Signal, 1)
		// listen for incoming signals SIGINT and SIGTERM
		os.Notify(signal, os.Interrupt, sys.SIGTERM)
		<-signal
		close(done)
	}()
	go httpServer()
	<-done
	os.Exit(0)
}

func httpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":8080", nil)
}