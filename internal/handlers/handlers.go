package handlers

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "index.html")
		w.Header().Set("Content-type", "text/html")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Недопустимый метод", http.StatusMethodNotAllowed)
	}
}

func PostForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {	
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		file, _, err := r.FormFile("myFile")
		if err != nil {
			http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
			return
		}
		defer file.Close()

		var data bytes.Buffer
		_, err = io.Copy(&data, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		str := service.Answer(data.String())

		localFile, err := os.OpenFile(time.Now().UTC().String(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = localFile.Write([]byte(str))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "text/plain; charset=UTF-8")
		_, err = w.Write([]byte(str))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Недопустимый метод", http.StatusMethodNotAllowed)
	}
}