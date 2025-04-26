package gateway_handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) UploadBook(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != http.MethodPost {
		http.Error(w, "Неверный метод запроса", http.StatusMethodNotAllowed)
		return
	}

	// Получаем файл из формы
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Ошибка при получении файла: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Отправляем файл на gRPC сервер
	bookId, err := h.service
	if err != nil {
		http.Error(w, "Ошибка при отправке файла на gRPC сервер: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Ответ на успешную загрузку
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Файл успешно передан на gRPC сервер")
}

func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetBookById(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) DeleteBookById(w http.ResponseWriter, r *http.Request) {

}
