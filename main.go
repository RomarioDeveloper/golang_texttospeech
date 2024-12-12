package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func generateAudioFile(text string) (string, error) {
	// URL API сервиса text-to-speech
	apiURL := "https://api.voicerss.org/"
	apiKey := "45df21933ceb4bb6899d7504bededbe5" // Вставьте ваш API-ключ

	// Кодируем текст для предотвращения ошибок передачи данных
	encodedText := url.QueryEscape(text)

	// Формируем запрос к API
	query := fmt.Sprintf("key=%s&hl=ru-ru&src=%s&c=MP3&v=Peter", apiKey, encodedText)
	response, err := http.Get(apiURL + "?" + query)
	if err != nil {
		return "", fmt.Errorf("Ошибка при выполнении запроса: %v", err)
	}
	defer response.Body.Close()

	// Проверяем статус ответа
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API вернул ошибку: %s", response.Status)
	}

	// Читаем содержимое ответа
	audioContent, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Ошибка чтения ответа: %v", err)
	}

	// Сохраняем аудио в уникальный файл MP3
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("output_%s.mp3", timestamp)
	err = os.WriteFile(filename, audioContent, 0644)
	if err != nil {
		return "", fmt.Errorf("Ошибка сохранения MP3 файла: %v", err)
	}

	return filename, nil
}

func main() {
	text := "Привет, это пример текста для озвучки в формате MP3. Если текст слишком длинный, он должен быть полностью озвучен без обрезки."
	file, err := generateAudioFile(text)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
	log.Printf("Аудиофайл успешно сохранен как %s", file)
}
