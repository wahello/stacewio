package main

import (
	"log"
	"net/http"
	"os"

	"github.com/stacew/.io/snsp"
)

//complete
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	s, err := snsp.NewServerAndBaseHandler()
	defer s.Close()
	if err != nil {
		log.Fatalln(err)
	}

	go s.Serve()
	http.Handle("/socket.io/", s)
	http.Handle("/", http.FileServer(http.Dir("./public")))

	//호출 후에 handler 등록 더 못함
	port := ":" + os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, nil))
}
