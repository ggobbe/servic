package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	argsLen := len(os.Args)
	if argsLen < 2 || argsLen > 3 {
		fmt.Println("Usage: servic [folder] ([port])")
		return
	}
	staticFolder := os.Args[1]
	port := 8080
	exists, err := exists(staticFolder)
	if err != nil {
		fmt.Printf("Error whilst checking if directory '%s' exists: %s\n", staticFolder, err)
		return
	}
	if !exists {
		fmt.Printf("Directory '%s' doesn't exists.\n", staticFolder)
		return
	}
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticFolder)))
	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
