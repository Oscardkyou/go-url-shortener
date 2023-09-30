package main

import (
	"fmt"              // Импортируем пакет для форматирования ввода/вывода
	"urlshortener/pkg" // Импортируем ваш собственный пакет для сокращения URL
)

func main() {
	// Создаем новый экземпляр UrlShortener из вашего пакета
	us := pkg.New()

	// Сокращаем длинную ссылку, используя метод Encode
	shortUrl := us.Encode("http://hdrezka.me/films/adventures/57217-gran-turizmo-2023.html")

	// Выводим сокращенную ссылку на экран
	fmt.Println("Short URL:", shortUrl)

	// Пытаемся восстановить длинную ссылку из сокращенной, используя метод Decode
	longUrl, exists := us.Decode(shortUrl)

	// Проверяем, найдена ли короткая ссылка в хранилище
	if !exists {
		// Если ссылка не найдена, выводим сообщение об ошибке и завершаем выполнение функции
		fmt.Println("Ошибка при декодировании: ссылка не найдена")
		return
	}

	// Выводим исходную длинную ссылку на экран
	fmt.Println("Long URL:", longUrl)
}
