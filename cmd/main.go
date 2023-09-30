package main

import (
	"fmt"
	"urlshortener/pkg"
)

func main() {
	us := pkg.New()
	shortUrl := us.Encode("http://hdrezka.me/films/adventures/57217-gran-turizmo-2023.html")
	fmt.Println("Short URL:", shortUrl)

	longUrl, _ := us.Decode(shortUrl)
	fmt.Println("Long URL:", longUrl)
}
