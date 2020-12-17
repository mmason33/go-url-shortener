package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getRedirects(file string) *map[string]string {
	// Open our jsonFile
	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		fmt.Println("No redirects loaded.")
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]string
	json.Unmarshal([]byte(byteValue), &result)

	return &result
}

func middleware(h http.Handler, redirects *map[string]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit Middleware")
		for key, value := range *redirects {
			fmt.Println(key, value)
			if r.URL.Path == key {
				http.Redirect(w, r, value, 301)
				break
			}
		}

		h.ServeHTTP(w, r)
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("Hello World"))
}

func main() {
	fileArg := flag.String("redirects", "./redirects.json", "The redirect JSON file you want to read in.")
	flag.Parse()
	mux := http.NewServeMux()
	redirects := getRedirects(*fileArg)

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", middleware(finalHandler, redirects))

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
