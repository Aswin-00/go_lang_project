// simple webserver using go

// package main imported
package main

// imported needed packages
import (
	"fmt"
	"log"
	"net/http"
)

// helloHandeler function to handle requests to the /hello endpoint
func helloHandeler(w http.ResponseWriter, r *http.Request) {

	// Logs that a request has been received for /hello
	log.Printf("received a request to hello ")

	// Check if the URL path is exactly /hello
	if r.URL.Path != "/hello" {
		// If not, send a 404 Not Found error
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	// Check if the request method is GET
	if r.Method != "GET" {
		// If not, send a 404 Not Found error
		http.Error(w, "method not found", http.StatusNotFound)
		return
	}

	// If the request is valid, respond with "came to hello"
	fmt.Fprint(w, "came to hello")
}

// formHandeler function to handle POST requests to /form
func formHandeler(w http.ResponseWriter, r *http.Request) {

	// Logs that a request has been received for /form
	log.Printf("received a request to form")

	// Check if the URL path is exactly /form
	if r.URL.Path != "/form" {
		// If not, send a 400 Bad Request error
		http.Error(w, "404 error", http.StatusBadRequest)
		return
	}

	// Check if the request method is POST
	if r.Method != "POST" {
		// If not, send a 400 Bad Request error
		http.Error(w, "wrong method", http.StatusBadRequest)
		return
	}

	// Extract form values for username and password
	name := r.FormValue("username")
	password := r.FormValue("password")

	// Send the username and password back in the response
	fmt.Fprintf(w, "username = %s\n ", name)
	fmt.Fprintf(w, "password = %s\n ", password)

	// Print the username and password to the server console
	fmt.Printf("name = %s\n password =%s \n ", name, password)
}

func main() {
	// File server to serve static files from the ./static directory
	fileserver := http.FileServer(http.Dir("./static"))

	// Print confirmation of connection to directory
	fmt.Println("connected to dir")

	// Log that the server has started
	log.Println("server started")

	// Set up handlers for the root path (serving static files), /form, and /hello
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandeler)
	http.HandleFunc("/hello", helloHandeler)

	// Start the server on port 8080 and log if there is an error
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// youtube vide :-https://youtu.be/jFfo23yIWac?si=anakht9Jf2P6N9nd&t=283
