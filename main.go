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
	http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

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

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Go Web App v2</title>
		<link rel="stylesheet" href="/static/styles.css">
	</head>
	<body>
		<div class="buttons">
			<a href="/hello" class="button green">Привет</a>
			<a href="/bye" class="button red">Пока</a>
			<a href="/refresh" class="button blue">Опять</a>
		</div>
		<img src="/static/image.jpg" alt="Example Image">
		<div class="pod-info">
			<div>Pod: %s</div>
			<div>Version: 2.0</div>
		</div>
	</body>
	</html>
	`, podName)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Привет v2</title>
		<style>
			body {
				background-color: #121212;
				color: #ffffff;
				display: flex;
				justify-content: center;
				align-items: center;
				height: 100vh;
				margin: 0;
				font-size: 24px;
				font-family: Arial, sans-serif;
			}
			h1 {
				text-shadow: 0 0 10px rgba(0,255,0,0.7);
			}
		</style>
	</head>
	<body>
		<h1>Привет от версии 2!</h1>
	</body>
	</html>
	`
	w.Write([]byte(html))
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Пока v2</title>
		<style>
			body {
				background-color: #250000;
				color: #ffaaaa;
				display: flex;
				justify-content: center;
				align-items: center;
				height: 100vh;
				margin: 0;
				font-size: 24px;
			}
		</style>
	</head>
	<body>
		<h1>Пока! (но кнопка теперь работает)</h1>
	</body>
	</html>
	`
	w.Write([]byte(html))
}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}
