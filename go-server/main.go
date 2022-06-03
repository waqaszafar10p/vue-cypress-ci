package main

import "fmt"
import "net/http"
import "log"


func main() {
	// http.HandleFunc("/", homePageHandler)
	// fmt.Println("Server listening on port 3050")
	// log.Panic(
	// 	http.ListenAndServe(":3050", nil),
	// )
	// Serve static files from the frontend/dist directory.
	fs := http.FileServer(http.Dir("../hello-world/dist"))
	http.Handle("/", fs)

	// Start the server.
	fmt.Println("Server listening on port 3050")
	log.Panic(
		http.ListenAndServe(":3050", nil),
	)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}