package handlers

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	var buff bytes.Buffer

	file, err := os.Open("index.html")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			http.Error(w, os.ErrNotExist.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = io.Copy(&buff, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(buff.Bytes())
}

func PostForm(w http.ResponseWriter, r *http.Request) {
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
	
	var str string
	if service.IsMorse(data.String()) {
		str = morse.ToText(data.String())
	} else {
		str = morse.ToMorse(data.String())
	}

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

	w.Header().Set("Content-type", "text/plain")
	w.Write([]byte(str))

}