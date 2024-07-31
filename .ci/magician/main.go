package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	fmt.Printf("TOKEN: %s", token)
}
