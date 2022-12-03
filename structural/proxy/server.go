package main

type server interface {
	// receives path(string), method(string)
	// returns http code(int), response body(string)
	handleRequest(string, string) (int, string)
}
