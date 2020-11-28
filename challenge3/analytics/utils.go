package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	jsonString = "json"
)

var allowedFormats = map[string]string{
	jsonString: jsonString,
}

// Message ...
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond ...
func Respond(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}

// ErrorResponse ...
func ErrorResponse(errorCode int, w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(errorCode)
	json.NewEncoder(w).Encode(data)
}

// PaginationParams - Process supplied params
func PaginationParams(r *http.Request) (int64, int64, string) {
	params := r.URL.Query()
	fromParam := params.Get("from")
	toParam := params.Get("to")
	formatParam := params.Get("format")

	now := time.Now().UTC()
	from := now.Add(-time.Hour * 1).Unix()
	to := now.Unix()
	format := jsonString

	if len(fromParam) > 0 {
		fromParsedTime, err := parseTime(fromParam)
		if err == nil {
			from = fromParsedTime
		}
	}
	if len(toParam) > 0 {
		parsedTime, err := parseTime(toParam)
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
	unescape, err := url.QueryUnescape(dateParamString)
	formatString := strings.Replace(unescape, "T", " ", -1)
	parsedTime, err := time.Parse(timeFormat, formatString)

	// fmt.Println("formatString", formatString)
	if err == nil {
		return parsedTime.Unix(), nil
	}
	return 0, err
}
