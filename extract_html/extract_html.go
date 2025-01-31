package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetchHTML(url string) {

	// Send a GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	// Write the HTML content to a file
	file, err := os.Create("output.html")
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(body)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
	}

	// OR, Print the HTML content
	// fmt.Println(string(body))

	//Extract HTML

	// Define target URL or PATH
	link_url := "https://github.com/mohanmal553/SE"
	fetchHTML(link_url)
}
