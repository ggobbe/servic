package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	staticFolder, port, err := processArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Servic: Starting server for directory '%s' on :%d\n", staticFolder, port)
	if err := server(staticFolder, port); err != nil {
		fmt.Println(err)
	}
}

func server(staticFolder string, port int) error {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticFolder)))
	http.Handle("/", r)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		return fmt.Errorf("Error listening on port %d\n  %s", port, err)
	}
	return nil
}

func processArgs(args []string) (string, int, error) {
	argsLen := len(args)
	if argsLen < 2 || argsLen > 3 {
		return "", 0, fmt.Errorf("Usage: %s [dir] ([port])", os.Args[0])
	}
	staticFolder := os.Args[1]
	port := 8080
	if argsLen > 2 {
		p, err := strconv.Atoi(os.Args[2])
		if err != nil || p < 1 || p > 65535 {
			return "", 0, errors.New("port isn't in range [1-65535]")
		}
		port = p
	}
	exists, err := exists(staticFolder)
	if err != nil {
		return "", 0, fmt.Errorf("Error whilst checking if directory '%s' exists\n  %s", staticFolder, err)
	}
	if !exists {
		return "", 0, fmt.Errorf("Directory '%s' doesn't exists", staticFolder)
	}
	return staticFolder, port, nil
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
