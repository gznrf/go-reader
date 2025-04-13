package gateway_handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func (h *Handler) UploadBook(w http.ResponseWriter, r *http.Request) {
	log.Println("📥 Новый запрос на /upload")

	// Получаем имя файла из заголовка
	filename := r.Header.Get("X-Filename")
	if filename == "" {
		http.Error(w, "❌ Заголовок 'X-Filename' обязателен", http.StatusBadRequest)
		log.Println("❌ Нет заголовка X-Filename")
		return
	}

	// Получаем домашнюю директорию пользователя
	homeDir, err := os.UserHomeDir()
	if err != nil {
		http.Error(w, "❌ Не удалось определить домашнюю директорию", http.StatusInternalServerError)
		log.Printf("❌ os.UserHomeDir error: %v\n", err)
		return
	}

	// Абсолютный путь к папке для сохранения
	saveDir := filepath.Join(homeDir, "Desktop", "код", "goProjects", "go_reader", "files")
	err = os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		http.Error(w, "❌ Не удалось создать папку", http.StatusInternalServerError)
		log.Printf("❌ os.MkdirAll error: %v\n", err)
		return
	}

	// Полный путь до нового файла
	savePath := filepath.Join(saveDir, filepath.Base(filename))

	// Создаем файл
	dst, err := os.Create(savePath)
	if err != nil {
		http.Error(w, "❌ Не удалось создать файл", http.StatusInternalServerError)
		log.Printf("❌ os.Create error: %v\n", err)
		return
	}
	defer dst.Close()

	// Копируем тело запроса в файл
	written, err := io.Copy(dst, r.Body)
	if err != nil {
		http.Error(w, "❌ Ошибка при записи файла", http.StatusInternalServerError)
		log.Printf("❌ io.Copy error: %v\n", err)
		return
	}

	log.Printf("✅ Файл сохранён (%d байт) в %s\n", written, savePath)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "✅ Файл '%s' успешно сохранён!\n", filename)
}

func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetBookById(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) DeleteBookById(w http.ResponseWriter, r *http.Request) {

}
