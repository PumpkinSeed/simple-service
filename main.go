package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
)

const (
	port = 8080
)

var (
	serviceid string
)

func init() {
	serviceid = pseudoUUID()
}

func main() {
	http.HandleFunc("/", serviceidHandler)
	log.Printf("Server listening on %s", getPort())
	http.ListenAndServe(getPort(), nil)
}

func serviceidHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, serviceid)
}

func getPort() string {
	return fmt.Sprintf(":%d", port)
}

func pseudoUUID() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return
}