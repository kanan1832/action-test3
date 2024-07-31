package main

import (
	"fmt"
	// "net/http"
	"os"
)

func main() {
	fmt.Println("Testing PR from fork")
	token := os.Getenv("GITHUB_TOKEN")
	fmt.Printf("TOKEN: %s", token)
}
