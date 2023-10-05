package main

import (
	"fmt"
	"net/http"
	"urlshortener/pkg"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page, GO is power!")
}

func contact_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is good!!!")
}

func main() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts", contact_page)
	http.ListenAndServe(":8080", nil)
	us := pkg.New()
	shortUrl := us.Encode("http://hdrezka.me/films/adventures/57217-gran-turizmo-2023.html")
	fmt.Println("Short URL:", shortUrl)

	longUrl, _ := us.Decode(shortUrl)
	fmt.Println("Long URL:", longUrl)

}
