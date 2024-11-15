package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"log/slog"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	port := GetPortToServe()
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/{code}", Httpbin)
	slog.Info("Starting server on port", "PORT", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		slog.Error("There was an error deploying the server")
	}

}

func GetPortToServe() string {
	portToServe := os.Getenv("SERVER_PORT")
	if portToServe == "" {
		return "8080"
	}
	return portToServe

}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	serverName := os.Getenv("SERVER_NAME")

	response := Response{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("Hello from [ %v ] =D", serverName),
	}

	w.WriteHeader(http.StatusOK)
	encoder_err := json.NewEncoder(w).Encode(response)

	if encoder_err != nil {
		slog.Error("There was an error encoding the response on HelloServer", "error", encoder_err)
	}
}

func Httpbin(w http.ResponseWriter, r *http.Request) {
	httpCodeString := r.PathValue("code")
	httpCode, err := strconv.Atoi(httpCodeString)

	response := Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}

	if err != nil {
		slog.Error("Could not convert the HttpCode to Int", "HttpCode", httpCodeString)
		response.Code = 400
		response.Message = fmt.Sprintf("Could not convert the HttpCode %v to Int", httpCodeString)
		w.WriteHeader(response.Code)
		encoder_err := json.NewEncoder(w).Encode(response)
		if encoder_err != nil {
			slog.Error("There was an error encoding the response on HelloServer", "error", encoder_err)
		}
		return
	}

	httpStatusText := http.StatusText(httpCode)
	if httpStatusText == "" {
		slog.Error("Could not convert the HttpCode", "HttpCode", httpCodeString)
		response.Code = 400
		response.Message = fmt.Sprintf("Could not convert the HttpCode %v", httpCodeString)
		w.WriteHeader(response.Code)
		encoder_err := json.NewEncoder(w).Encode(response)
		if encoder_err != nil {
			slog.Error("There was an error encoding the response on HelloServer", "error", encoder_err)
		}
		return
	}

	response.Code = httpCode
	response.Message = httpStatusText

	slog.Info("Response to return", "Response", response)

	w.WriteHeader(response.Code)
	encoder_err := json.NewEncoder(w).Encode(response)

	if encoder_err != nil {
		slog.Error("There was an error encoding the response on HelloServer", "error", encoder_err)
	}

}
