package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		_, _ = fmt.Fprint(w, "UP")
	})
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		resp, err := http.Get("https://golang.org") // requires CA certs
		if errorRespOnError(err, w) != nil {
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if errorRespOnError(err, w) != nil {
			return
		}
		_, _ = fmt.Fprint(w, len(body))
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func errorRespOnError(err error, w http.ResponseWriter) error {
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, err.Error())
	}
	return err
}

func logRequest(r *http.Request) {
	log.Println(r.Method, r.URL)
}
