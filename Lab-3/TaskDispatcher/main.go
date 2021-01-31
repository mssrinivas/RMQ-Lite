package main

import (
  "net/http"
  "log"
  "os"
)

// Fetch the environment variable if it exists or else return the defaultValue set in the code.
func GetValueFromEnvVariable(variableName, defaultValue string) string {
	environmentValue := os.Getenv(variableName)
	if environmentValue == "" {
		return defaultValue
	}
	return environmentValue
}

func HttpKeepAlive(port string) {
	errChan := make(chan error)
	go func() {
    log.Println("HTTP KeepAlive :transport", "HTTP", "started on port", port)
		errChan <- http.ListenAndServe(port, nil)
	}()

	log.Fatal("exit", <-errChan)
}

func main() {
  port := GetValueFromEnvVariable("ENV_PORT", ":8086")
  HttpKeepAlive(port)
}