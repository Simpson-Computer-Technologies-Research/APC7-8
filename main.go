package main

// Import Packages
import (
	"fmt"
	"net/http"
)

// The HomePage() function
func HomePage(w http.ResponseWriter, r *http.Request) {

	// Use the Fprintf(w, "Response Content") function to return
	// a string to an user
	fmt.Fprintf(w, "Welcome to the HomePage!")

	// Print to the console that an user has visited your api site
	fmt.Printf("\nEndpoint Hit (%s): /home\n", r.Method)
}

// The Main function
func main() {
	// Handle the "/home" Page
	//
	// This means that when an user goes to http://localhost:8080/home
	//
	// the content in the HomePage() function will display
	//
	http.HandleFunc("/home", HomePage)

	// Listen to incoming requests on the port 8080
	//
	// http://localhost:8080
	//
	http.ListenAndServe(":8080", nil)
}
