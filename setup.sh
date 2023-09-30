#!/bin/bash

# Создаем основную директорию проекта
mkdir go-url-shortener
cd go-url-shortener

# Создаем внутренние директории для проекта
mkdir cmd pkg

# Создаем файлы main.go и shortener.go
touch cmd/main.go pkg/shortener.go

echo "Структура проекта успешно создана!"
