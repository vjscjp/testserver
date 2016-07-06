// Simple API server
// Test Server api to generate load and return cutom http status code and
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"
)

func main() {
	http.HandleFunc("/", page)
	http.HandleFunc("/hits", hitServer)
	port := ":9090"
	log.Printf("Server listening at %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

//Request input request parameters
type Request struct {
	URL        string
	StatusCode int
	Delay      int
}

func page(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func hitServer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Error, Invalid Request")
	}
	reg := new(Request)
	if err = json.Unmarshal(b, reg); err != nil {
		fmt.Fprintf(w, "Error, Failled to decode Request err: %s", err)
	}

	time.Sleep(time.Duration(reg.Delay) * time.Millisecond)

	if reg.StatusCode != 0 {
		w.WriteHeader(reg.StatusCode)
	}

	if strings.TrimSpace(reg.URL) != "" {
		hitExternalServer(reg, w)
	} else {
		fmt.Fprint(w, "Done ")
	}

}

func hitExternalServer(regParam *Request, w http.ResponseWriter) {

	resp, err := http.Get(regParam.URL)
	if err != nil {
		fmt.Printf("failed to get valid response. Err: %s \n", err.Error())
		w.WriteHeader(500)
		fmt.Fprintf(w, "Err: %s \n", err.Error())
	} else {
		//Dont Print the endpoint response as we dont need it
		//fmt.Fprintf(w, "%s :: Response %s", regParam.URL, string(body))
		fmt.Fprintf(w, "%s :: Done ", regParam.URL)
	}
	if regParam.StatusCode != 0 {
		w.WriteHeader(regParam.StatusCode)
	} else {
		if resp != nil {
			w.WriteHeader(resp.StatusCode)
		}

	}

}
