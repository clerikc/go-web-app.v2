package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/bye", byeHandler)
	http.HandleFunc("/refresh", refreshHandler)

	// Server config
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting server v2 on port %s...", port)
	log.Fatal(server.ListenAndServe())
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	podName := os.Getenv("HOSTNAME")
	message := r.URL.Query().Get("message")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>Go Web App v2</title>
		<link rel="stylesheet" href="/static/styles.css">
	</head>
	<body>
		%s
		<div class="buttons">
			<a href="/hello" class="button green">Привет</a>
			<a href="/bye" class="button red">Пока</a>
			<a href="/refresh" class="button blue">Опять</a>
		</div>
		<img src="/static/image.jpg" alt="Example Image" class="main-image">
		<div class="pod-info">
			<div>Pod: %s</div>
			<div>Version: 2.0</div>
		</div>
	</body>
	</html>
	`, getMessageHTML(message), podName)

	w.Write([]byte(html))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	showLink := r.URL.Query().Get("showlink") == "true"
	goBack := r.URL.Query().Get("goback") == "true"

	if goBack {
		http.Redirect(w, r, "/?message=Пока пока=)))", http.StatusFound)
		return
	}

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>Привет v2</title>
		<link rel="stylesheet" href="/static/styles.css">
	</head>
	<body>
		<img src="/static/image2.jpg" alt="Special Image" class="main-image">
		%s
		<div class="buttons">
			<a href="/hello?showlink=true" class="button green">Показать ссылку</a>
			<a href="/hello?goback=true" class="button blue">Вернуться</a>
		</div>
	</body>
	</html>
	`, getLinkHTML(showLink))

	w.Write([]byte(html))
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>Пока! v2</title>
		<link rel="stylesheet" href="/static/styles.css">
	</head>
	<body>
		<h1>Пока!</h1>
	</body>
	</html>
	`
	w.Write([]byte(html))
}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}

func getMessageHTML(message string) string {
	if message == "" {
		return ""
	}
	return fmt.Sprintf(`<div class="message">%s</div>`, message)
}

func getLinkHTML(show bool) string {
	if !show {
		return ""
	}
	return `
	<div class="link-message">
		Проект на GitHub: 
		<a href="https://github.com/clerikc/go-web-app.v2" class="github-link">
			github.com/clerikc/go-web-app.v2
		</a>
	</div>
	`
}
