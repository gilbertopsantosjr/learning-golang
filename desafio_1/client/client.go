package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Client running...")
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	select {
	case <-ctx.Done():
		// Read the response body
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("failed to read response body: %v\n", err)
			return
		}

		// Convert the response body to a string
		bodyString := string(bodyBytes)

		writeResultToFile(bodyString)
	}
}

func writeResultToFile(bid string) {
	// The name of the file to create.
	fileName := "bids.txt"

	// Create a file with the specified name.
	file, err := os.Create(fileName)
	if err != nil {
		// Handle the error if the file creation fails.
		fmt.Printf("failed to create file: %v\n", err)
	}

	// Ensure the file is closed when the function exits.
	defer file.Close()

	formatted := fmt.Sprintf("Dólar: %s", bid)

	// Write the content to the file.
	_, err = file.WriteString(formatted)
	if err != nil {
		// Handle the error if writing to the file fails.
		fmt.Printf("failed to write to file: %v\n", err)
	}

	fmt.Println("File written successfully")
}
