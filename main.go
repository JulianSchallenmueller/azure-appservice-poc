package main

import (
    "fmt"
    "net/http"
)

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// Custom handler function for the /hello endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    fmt.Fprintf(w, "Hello from Go!\n")
}

func main() {
    // Create a new mux
    mux := http.NewServeMux()

    // Add the /hello route to the mux
    mux.HandleFunc("/api/hello", helloHandler)

    // Start the server
    fmt.Println("Server starting on port 80")
    http.ListenAndServe(":80", mux)
}

