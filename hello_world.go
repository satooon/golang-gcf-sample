package hello_world

import (
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello World!")

	if _, err := w.Write([]byte("Hello World Response!")); err != nil {
		log.Fatalln(err)
	}
}
