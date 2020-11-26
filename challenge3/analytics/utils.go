package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/render"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	jsonString = "json"
)

 var allowedFormats = map[string]string{
	jsonString: jsonString
 } 

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	render.JSON(w, r, data)
}

func ErrorResponse(errorCode int, w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(errorCode)
	render.JSON(w, r, data)
}

// PaginationParams - Process supplied params
func PaginationParams(r *http.Request) (int64, int64, string) {
	params := r.URL.Query()
	fromParam := params.Get("from")
	toParam := params.Get("to")
	formatParam := params.Get("format")

	now := time.Now().UtC()
	from := now.Sub(1 * time.Hour).Format(timeFormat).Unix()
	to := now.Format(timeFormat).Unix()
	format := jsonString

	if len(fromParam) > 0 {
		parsedTime, err := parseTime(fromParam)
		if err == nil {
			from = parsedTime
		}
	}
	if len(toParam) > 0 {
		parsedTime, err := parseTime(fromParam)
		if err == nil {
			to = parsedTime
		}
	}

	if _, ok := allowedFormats[formatParam]; ok {
		format = formatParam
	}

	return from, to, format
}

func parseTime(dateParamString string) (int64, error) {
	unescape, err := url.QueryEscape(dateParamString)
	if err == nil {
		parsedTime, err := time.Parse(timeFormat, unescape)
		if err == nil {
			parsedTime = parsedTime.UTC().Unix()

			return parsedTime, nil
		}
	}
	return nil, err
}
