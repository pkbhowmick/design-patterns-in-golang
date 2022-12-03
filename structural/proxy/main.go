package main

import (
	"fmt"
	"net/http"
)

func printResult(statusCode int, body string) {
	fmt.Println("StatusCode: ", statusCode, " response: ", body)
}

func main() {
	nginxServer := newNginxServer()
	statusPath := "/app/status"
	createUserPath := "/create/user"

	printResult(nginxServer.handleRequest(statusPath, http.MethodGet))
	printResult(nginxServer.handleRequest(createUserPath, http.MethodPost))
	printResult(nginxServer.handleRequest(createUserPath, http.MethodPost))
	printResult(nginxServer.handleRequest(createUserPath, http.MethodPost)) // will be rate limited
}
