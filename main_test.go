package main

import (
	"log"
	"net/http"
	"testing"

	snsp "github.com/stacew/io/sNSP"
)

func Test(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("main()...")

	s := snsp.NewSocketServer()
	defer s.Close()
	go s.Serve()

	snsp.RegisterHandler(s)

	http.Handle("/socket.io/", s)                           //Socket
	http.Handle("/", http.FileServer(http.Dir("./public"))) //file Server

	log.Println("ListenAndServe()...")
	log.Fatal(http.ListenAndServe(":"+GetPort(), nil)) //http start
}
