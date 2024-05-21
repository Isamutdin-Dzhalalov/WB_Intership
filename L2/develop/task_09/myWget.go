package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
	"log"
    "path/filepath"
)

func downloadFile(url string, path string) error {
    // GET запрос.
    resp, err := http.Get(url)
    if err!= nil {
		log.Fatal(err)
    }
    defer resp.Body.Close()

    // Проверяем статус ответа.
    if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
    }

    // Получаем имя файла из URL.
    filename := filepath.Base(url)

    // Открываем файл для записи.
    file, err := os.Create(path + "/" + filename)
    if err!= nil {
		log.Fatal(err)
    }
    defer file.Close()

    // Копируем данные из тела ответа в файл.
    _, err = io.Copy(file, resp.Body)
    if err!= nil {
		log.Fatal(err)
    }

    return nil
}

func main() {
    url := "http://example.com" // URL сайта.
    path := "./downloaded_files" // Путь, где будут сохраняться файлы.

    // Создаем директорию, если она не существует.
    if err := os.MkdirAll(path, 0755); err!= nil {
        fmt.Printf("Error creating directory: %v\n", err)
        return
    }

    err := downloadFile(url, path)
    if err!= nil {
		log.Fatal(err)
    }

    fmt.Println("Download success!")
}

