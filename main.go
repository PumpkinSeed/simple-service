package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	serviceid string
	next string
	port string
)

func init() {
	serviceid = pseudoUUID()
	next = os.Getenv("SS_NEXT")
	port = os.Getenv("SS_PORT")
	if port == "" {
		port = "8080"
	}
}

func main() {
	http.HandleFunc("/", serviceidHandler)
	log.Printf("Server listening on %s", getPort())
	if next != "" {
		log.Printf("Next service: %s", next)
	}
	http.ListenAndServe(getPort(), nil)
}

func serviceidHandler(w http.ResponseWriter, r *http.Request) {
	var serviceStack []string
	if next != "" {
		resp, err := http.Get(next)
		if err != nil {
			panic(err)
		}
		respbytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		serviceStack = append(serviceStack, serviceid)
		result := strings.Split(string(respbytes), ",")
		serviceStack = append(serviceStack, result...)
		fmt.Fprintf(w, strings.Join(serviceStack, ","))
		return
	}
	fmt.Fprintf(w, serviceid)
}

func getPort() string {
	return fmt.Sprintf(":%s", port)
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