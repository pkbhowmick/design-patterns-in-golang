package main

type Nginx struct {
	app               *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNginxServer() *Nginx {
	return &Nginx{
		app:               &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) handleRequest(path string, method string) (int, string) {
	if !n.checkRateLimiting(path) {
		return 403, "Not allowed"
	}

	return n.app.handleRequest(path, method)
}

func (n *Nginx) checkRateLimiting(path string) bool {
	if n.rateLimiter[path] == 0 {
		n.rateLimiter[path] = 1
	}

	if n.rateLimiter[path] > n.maxAllowedRequest {
		return false
	}

	n.rateLimiter[path] += 1
	return true
}
