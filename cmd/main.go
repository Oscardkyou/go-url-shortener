package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"urlshortener/pkg"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
}

func (u User) GetAllinfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) SetNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}
	tmpl, _ := template.ParseFiles("templates/home_page.html") // Используйте правильное имя файла
	tmpl.Execute(w, bob)

	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		longURL := string(body)
		us := pkg.New()
		shortUrl := us.Encode(longURL)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(shortUrl))
		return
	}

	//bob.SetNewName("Alexander")
	//w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//fmt.Fprintf(w, bob.GetAllinfo())
}

func contact_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is good!!!")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts", contact_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
	us := pkg.New()
	shortUrl := us.Encode("http://hdrezka.me/films/adventures/57217-gran-turizmo-2023.html")
	fmt.Println("Short URL:", shortUrl)

	longUrl, _ := us.Decode(shortUrl)
	fmt.Println("Long URL:", longUrl)
}
