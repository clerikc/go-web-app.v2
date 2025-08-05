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
	return `<a href="https://example.com" class="special-link">Секретная ссылка!</a>`
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
				border-radius: 5px;
				transition: all 0.3s ease;
				box-shadow: none;
				border: none;
				outline: none;
			}
			.neon-link:hover {
				text-shadow: 0 0 5px, 0 0 10px, 0 0 15px;
			}
			.green-neon {
				color: #00ff00;
				text-shadow: 0 0 5px #00ff00;
			}
			.green-neon:hover {
				color: #00ff00;
				text-shadow: 0 0 10px #00ff00, 0 0 20px #00ff00;
			}
			.red-neon {
				color: #ff0000;
				text-shadow: 0 0 5px #ff0000;
			}
			.red-neon:hover {
				color: #ff0000;
				text-shadow: 0 0 10px #ff0000, 0 0 20px #ff0000;
			}
			.blue-neon {
				color: #00bfff;
				text-shadow: 0 0 5px #00bfff;
			}
			.blue-neon:hover {
				color: #00bfff;
				text-shadow: 0 0 10px #00bfff, 0 0 20px #00bfff;
			}
			.buttons {
				margin: 30px 0;
				text-align: center;
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
                background-color: black;
                margin: 0;
                height: 100vh;
                overflow: hidden;
            }
            .bye-message {
                position: absolute;
                top: 50%;
                left: 50%;
                transform: translate(-50%, -50%);
                color: white;
                font-size: 3em;
                opacity: 0;
                transition: opacity 1s;
                z-index: 100;
                cursor: pointer;
                text-decoration: none;
            }
            .bye-message:hover {
                color: #00BFFF;
            }
            .rectangle-container {
                position: relative;
                width: 100%;
                height: 100%;
            }
            .bye-rectangle {
                position: absolute;
                width: 30px;
                height: 60px;
                background-color: #00BFFF;
                opacity: 0.8;
            }
        </style>
    </head>
    <body class="bye-page">
        <a href="/" class="bye-message">Пока!</a>
        <div class="rectangle-container" id="container"></div>
        <script>
            document.addEventListener('DOMContentLoaded', function() {
                const container = document.getElementById('container');
                const rectCount = 20;
                const rects = [];
                const message = document.querySelector('.bye-message');
                
                // Показываем сообщение через 2 секунды
                setTimeout(function() {
                    message.style.opacity = '1';
                }, 2000);
                
                function createRectangles() {
                    for (var i = 0; i < rectCount; i++) {
                        // Левые прямоугольники
                        var leftRect = document.createElement('div');
                        leftRect.className = 'bye-rectangle';
                        leftRect.style.left = '0';
                        leftRect.style.top = (i * (window.innerHeight / rectCount)) + 'px';
                        container.appendChild(leftRect);
                        rects.push({element: leftRect, isRight: false});
                        
                        // Правые прямоугольники
                        var rightRect = document.createElement('div');
                        rightRect.className = 'bye-rectangle';
                        rightRect.style.right = '0';
                        rightRect.style.left = 'auto';
                        rightRect.style.top = (i * (window.innerHeight / rectCount)) + 'px';
                        container.appendChild(rightRect);
                        rects.push({element: rightRect, isRight: true});
                    }
                }
                
                function animateToCenter() {
                    var centerX = window.innerWidth / 2;
                    var duration = 2000;
                    var startTime = performance.now();
                    
                    function update(time) {
                        var elapsed = time - startTime;
                        var progress = Math.min(elapsed / duration, 1);
                        
                        rects.forEach(function(rect, index) {
                            var rectEl = rect.element;
                            if (rect.isRight) {
                                // Для правых прямоугольников
                                var currentRight = parseInt(rectEl.style.right || 0);
                                var currentX = window.innerWidth - currentRight - 30;
                                var targetX = centerX - 15;
                                var newX = currentX + (targetX - currentX) * progress;
                                rectEl.style.right = (window.innerWidth - newX - 30) + 'px';
                                rectEl.style.left = 'auto';
                            } else {
                                // Для левых прямоугольников
                                var currentX = parseInt(rectEl.style.left || 0);
                                var targetX = centerX - 15;
                                var newX = currentX + (targetX - currentX) * progress;
                                rectEl.style.left = newX + 'px';
                                rectEl.style.right = 'auto';
                            }
                            
                            // Волновой эффект
                            var delay = index * 50;
                            if (elapsed >= delay) {
                                rectEl.style.opacity = 0.8 - (0.7 * ((elapsed - delay) / (duration - delay)));
                            }
                        });
                        
                        if (progress < 1) {
                            requestAnimationFrame(update);
                        } else {
                            setTimeout(animateToOppositeSide, 500);
                        }
                    }
                    
                    requestAnimationFrame(update);
                }
                
                function animateToOppositeSide() {
                    var duration = 2000;
                    var startTime = performance.now();
                    
                    function update(time) {
                        var elapsed = time - startTime;
                        var progress = Math.min(elapsed / duration, 1);
                        
                        rects.forEach(function(rect, index) {
                            var rectEl = rect.element;
                            if (rect.isRight) {
                                // Для правых прямоугольников
                                var currentRight = parseInt(rectEl.style.right || 0);
                                var currentX = window.innerWidth - currentRight - 30;
                                var targetX = -30;
                                var newX = currentX + (targetX - currentX) * progress;
                                rectEl.style.right = (window.innerWidth - newX - 30) + 'px';
                                rectEl.style.left = 'auto';
                            } else {
                                // Для левых прямоугольников
                                var currentX = parseInt(rectEl.style.left || 0);
                                var targetX = window.innerWidth;
                                var newX = currentX + (targetX - currentX) * progress;
                                rectEl.style.left = newX + 'px';
                                rectEl.style.right = 'auto';
                            }
                            
                            // Возвращаем прозрачность
                            var delay = index * 50;
                            if (elapsed >= delay) {
                                rectEl.style.opacity = 0.1 + (0.7 * ((elapsed - delay) / (duration - delay)));
                            }
                        });
                        
                        if (progress < 1) {
                            requestAnimationFrame(update);
                        } else {
                            resetRectangles();
                            setTimeout(startAnimation, 1000);
                        }
                    }
                    
                    requestAnimationFrame(update);
                }
                
                function resetRectangles() {
                    rects.forEach(function(rect, index) {
                        var rectEl = rect.element;
                        if (rect.isRight) {
                            rectEl.style.right = '0';
                            rectEl.style.left = 'auto';
                        } else {
                            rectEl.style.left = '0';
                            rectEl.style.right = 'auto';
                        }
                        rectEl.style.top = (Math.floor(index/2) * (window.innerHeight / rectCount)) + 'px';
                        rectEl.style.opacity = '0.8';
                    });
                }
                
                function startAnimation() {
                    animateToCenter();
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
