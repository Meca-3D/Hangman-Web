package main

import (
	"fmt"
	"net/http"
)

func web() {
	fs := http.FileServer(http.Dir("template"))
	http.Handle("/template/", http.StripPrefix("/template/", fs))

	fs = http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fs = http.FileServer(http.Dir("font"))
	http.Handle("/font/", http.StripPrefix("/font/", fs))

	http.HandleFunc("/", Home)
	http.HandleFunc("/solo", Solo)

	fmt.Println("(http://localhost:8080) - Server started on port ", PORT)
	http.ListenAndServe(PORT, nil)
}
