package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	pwd, _ := os.Getwd()
	var dir string
	var port int
	//var port int
	flag.StringVar(&dir, "serveDir", pwd, "directory the server serve")
	flag.IntVar(&port, "port", 18080, "port the server listen")
	flag.Parse()
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(dir))))
	log.Printf("server running on port %d", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
