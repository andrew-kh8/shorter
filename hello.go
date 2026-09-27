package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/check", checkHandler)
	http.ListenAndServe(":8080", nil)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	var userUrl string = r.URL.Query().Get("url")

	log.Println("User url: ", userUrl)

	if userUrl != "" {
		res, err := isUrlReal(userUrl)

		if err != nil {
			log.Fatal("FATAL")
		}

		if res {
			w.Write([]byte("Such a good boy"))
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("IT'S NOT A FUCKING URL"))
		}
	} else {
		http.Error(w, "No url provided", http.StatusBadRequest)
	}
}
