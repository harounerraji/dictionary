
package main

import (
	"fmt"
	"exemple3/dictionary"
	"net/http"
)

func main() {
	
	myDictionary := make(dictionary.Dictionary)

	http.HandleFunc("/add", myDictionary.AddHandler)
	http.HandleFunc("/get", myDictionary.GetHandler)
	http.HandleFunc("/remove", myDictionary.RemoveHandler)
	http.HandleFunc("/list", myDictionary.ListHandler)

	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}

	
}
