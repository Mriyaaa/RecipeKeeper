package main

import (
	"net/http"

	"github.com/Mriyaaa/RecipeKeeper/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.ListenAndServe(":8000", nil)
}
