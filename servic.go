package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	argsLen := len(os.Args)
	if argsLen < 2 || argsLen > 3 {
		fmt.Printf("Usage: %s [dir] ([port])\n", os.Args[0])
		return
	}
	staticFolder := os.Args[1]
	port := 8080
	if argsLen > 2 {
		p, err := strconv.Atoi(os.Args[2])
		if err != nil || p < 1 || p > 65535 {
			fmt.Println("port isn't in range [1-65535]")
			return
		}
		port = p
	}
	exists, err := exists(staticFolder)
	if err != nil {
		fmt.Printf("Error whilst checking if directory '%s' exists\n  %s\n", staticFolder, err)
		return
	}
	if !exists {
		fmt.Printf("Directory '%s' doesn't exists.\n", staticFolder)
		return
	}
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticFolder)))
	http.Handle("/", r)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("Error listening on port %d\n  %s\n", port, err)
	}
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
