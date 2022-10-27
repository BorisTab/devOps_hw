package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	backClient := http.Client{}

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")

		if key == "" || value == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err := backClient.Get("http://back_server:8080/set?key=" + key + "&value=" + value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")

		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp, err := backClient.Get("http://back_server:8080/get?key=" + key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"value\": \"" + string(body) + "\"}"))
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		_, err := backClient.Get("http://back_server:8080/delete_all")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":9090", nil)
}
