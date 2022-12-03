package main

import (
	"net/http"
)

type Application struct {
}

func (a *Application) handleRequest(path string, method string) (int, string) {
	if path == "/server/status" && method == http.MethodGet {
		return 200, "OK"
	}

	if path == "/create/user" && method == http.MethodPost {
		return 201, "User created"
	}

	return 404, "Not OK"
}
