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

// Добавляем недостающую функцию refreshHandler
func refreshHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Добавляем недостающую функцию getMessageHTML
func getMessageHTML(message string) string {
	if message == "" {
		return ""
	}
	return fmt.Sprintf(`<div class="message">%s</div>`, message)
}

// Добавляем недостающую функцию getLinkHTML
func getLinkHTML(show bool) string {
	if !show {
		return ""
	}
	return `<a href="https://github.com/clerikc/go-web-app.v2" class="special-link">Секретная ссылка!</a>`
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
		<style>
			.neon-link {
				display: inline-block;
				margin: 0 15px;
				padding: 10px 20px;
				text-decoration: none;
				font-size: 1.5em;
				font-weight: bold;
				transition: all 0.3s ease;
				background: transparent !important;
				border: none !important;
			}
			.neon-link:hover {
				text-shadow: 0 0 8px;
			}
			.green-neon {
				color: #0f0;
				text-shadow: 0 0 5px rgba(0, 255, 0, 0.3);
			}
			.green-neon:hover {
				color: #0f0;
				text-shadow: 0 0 10px rgba(0, 255, 0, 0.7), 0 0 20px rgba(0, 255, 0, 0.5);
			}
			.red-neon {
				color: #f33;
				text-shadow: 0 0 5px rgba(255, 0, 0, 0.3);
			}
			.red-neon:hover {
				color: #f33;
				text-shadow: 0 0 10px rgba(255, 0, 0, 0.7), 0 0 20px rgba(255, 0, 0, 0.5);
			}
			.blue-neon {
				color: #0af;
				text-shadow: 0 0 5px rgba(0, 170, 255, 0.3);
			}
			.blue-neon:hover {
				color: #0af;
				text-shadow: 0 0 10px rgba(0, 170, 255, 0.7), 0 0 20px rgba(0, 170, 255, 0.5);
			}
			.buttons {
				margin: 30px 0;
				text-align: center;
				background: transparent !important;
			}
			body {
				background-color: #111;
				color: #fff;
			}
		</style>
	</head>
	<body>
		%s
		<div class="buttons">
			<a href="/hello" class="neon-link green-neon">Привет</a>
			<a href="/bye" class="neon-link red-neon">Пока</a>
			<a href="/refresh" class="neon-link blue-neon">Опять</a>
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
		http.Redirect(w, r, "/?message=Пока+пока", http.StatusSeeOther)
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
			<a href="/hello?goback=true" class="button blue">Вернуться!</a>
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
        <style>
            .bye-page {
                background-color: #111;
                margin: 0;
                height: 100vh;
                overflow: hidden;
                font-family: Arial, sans-serif;
            }
            .bye-message {
                position: absolute;
                top: 50%;
                left: 50%;
                transform: translate(-50%, -50%);
                color: #00BFFF;
                font-size: 3em;
                opacity: 0;
                transition: all 0.5s ease;
                z-index: 100;
                cursor: pointer;
                text-decoration: none;
                text-shadow: 0 0 10px rgba(0, 191, 255, 0.7);
                font-weight: bold;
            }
            .bye-message:hover {
                text-shadow: 0 0 15px #00BFFF, 0 0 30px rgba(0, 191, 255, 0.5);
                transform: translate(-50%, -50%) scale(1.1);
            }
            .rectangle-container {
                position: relative;
                width: 100%;
                height: 100%;
            }
            .bye-rectangle {
                position: absolute;
                width: 20px;
                height: 40px;
                background-color: rgba(0, 191, 255, 0.5);
                border-radius: 4px;
                opacity: 0.8;
                transition: all 0.3s ease;
                transform-origin: center;
            }
            .bye-rectangle:hover {
                background-color: rgba(0, 191, 255, 0.8);
            }
        </style>
    </head>
    <body class="bye-page">
        <a href="/" class="bye-message">Пока!</a>
        <div class="rectangle-container" id="container"></div>
        <script>
            document.addEventListener('DOMContentLoaded', function() {
                const container = document.getElementById('container');
                const rectCount = 40;
                const rects = [];
                const message = document.querySelector('.bye-message');
                const centerX = window.innerWidth / 2;
                const centerY = window.innerHeight / 2;
                const radius = Math.min(window.innerWidth, window.innerHeight) * 0.4;
                let angle = 0;
                
                // Показываем сообщение через 2 секунды
                setTimeout(function() {
                    message.style.opacity = '1';
                }, 2000);
                
                function createRectangles() {
                    for (let i = 0; i < rectCount; i++) {
                        // Первая спираль (синяя)
                        const rect1 = document.createElement('div');
                        rect1.className = 'bye-rectangle';
                        rect1.style.backgroundColor = 'rgba(0, 191, 255, 0.6)';
                        container.appendChild(rect1);
                        rects.push({element: rect1, spiral: 1, index: i});
                        
                        // Вторая спираль (голубая)
                        const rect2 = document.createElement('div');
                        rect2.className = 'bye-rectangle';
                        rect2.style.backgroundColor = 'rgba(100, 210, 255, 0.6)';
                        container.appendChild(rect2);
                        rects.push({element: rect2, spiral: 2, index: i});
                    }
                }
                
                function updatePositions() {
                    const time = Date.now() * 0.001;
                    const speed = 0.5;
                    
                    rects.forEach(rect => {
                        const spiralAngle = angle + (rect.index / rectCount) * Math.PI * 2;
                        const spiralOffset = rect.spiral === 1 ? 0 : Math.PI;
                        const currentAngle = spiralAngle + spiralOffset;
                        
                        // Позиция по спирали
                        const spiralProgress = (rect.index / rectCount) * 2 * Math.PI;
                        const x = centerX + Math.cos(currentAngle + time * speed) * radius * (0.5 + 0.5 * Math.sin(spiralProgress));
                        const y = centerY + Math.sin(currentAngle + time * speed) * radius * (0.5 + 0.5 * Math.sin(spiralProgress));
                        
                        // Размер и поворот
                        const scale = 0.5 + 0.5 * Math.sin(time * 0.5 + rect.index * 0.1);
                        const rotation = currentAngle * (180 / Math.PI);
                        
                        rect.element.style.left = x + 'px';
                        rect.element.style.top = y + 'px';
                        rect.element.style.transform = 'rotate(' + rotation + 'deg) scale(' + scale + ')';
                        rect.element.style.opacity = 0.5 + 0.4 * Math.sin(time + rect.index * 0.2);
                        rect.element.style.zIndex = Math.floor(scale * 100);
                    });
                    
                    angle += 0.005;
                    requestAnimationFrame(updatePositions);
                }
                
                function startAnimation() {
                    updatePositions();
                }
                
                // Инициализация
                createRectangles();
                setTimeout(startAnimation, 500);
            });
        </script>
    </body>
    </html>
    `
	w.Write([]byte(html))
}
