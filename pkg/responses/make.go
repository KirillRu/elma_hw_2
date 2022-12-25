package responses

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

func Make(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		makeResponseError(w, err)

		return
	}

	if data == nil {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	responseContent, err := json.Marshal(data)
	if err != nil {
		makeResponseError(w, err)
		return
	}

	w.Write(responseContent)
}

func DrawPage(w http.ResponseWriter, fileName string, data interface{}, err error) {
	if err != nil {
		makeResponseError(w, err)

		return
	}

	_, err = os.Stat("d:\\web\\Projects\\Go\\src\\elma_hw_2\\templates\\" + fileName)
	if err != nil {
		makeResponseError(w, err)
		//TODO:: сделать лог
		return
	}
	buf := new(bytes.Buffer)

	files := []string{
		"d:\\web\\Projects\\Go\\src\\elma_hw_2\\templates\\" + fileName,
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		makeResponseError(w, err)
		//TODO:: сделать лог
		return
	}
	err = ts.ExecuteTemplate(buf, "htmlcontent", data) //
	if err != nil {
		makeResponseError(w, err)
		//TODO:: сделать лог
		return
	}
	w.Write([]byte(buf.String()))
}

func makeResponseError(w http.ResponseWriter, err error) {
	resp := ErrorWriter{err: err}
	resp.write(w)
}

type ErrorWriter struct {
	err error
}

func (e ErrorWriter) write(w http.ResponseWriter) {
	if e.err == nil {
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	response, err := json.Marshal(ErrorResponse{Message: e.err.Error()})
	if err != nil {
		return
	}

	w.Write(response)
}

type ErrorResponse struct {
	Message string `json:"message"`
}
