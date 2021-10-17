package main

import (
	"log"
	"net/http"
)

var users = map[string]string{
	"hamdan": "hamdan21",
}

func isAuthorized(user, password string) bool {
	pass, ok := users[user]
	if !ok {
		return false
	}

	return password == pass
}

func greeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Add("WWW-Authenticate", `Basic realm="Please provide a username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`"Message": "No basic auth present"`))
		return
	}

	if !isAuthorized(username, password) {
		w.Header().Add("WWW-Authenticate", `Basic realm="Please provide a username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`"Message": "Incorrect username or password"`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`"Message": "Welcome, my name is Hamdan. Let me guide you"`))
}

func main() {
	http.HandleFunc("/", greeting)
	log.Println("Starting server at port :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
