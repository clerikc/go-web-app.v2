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
    </head>
    <body class="bye-page">
        <div class="bye-message">Пока!</div>
        <div class="rectangle-container" id="container"></div>
        <script>
            const container = document.getElementById('container');
            const rectCount = 20;
            const rects = [];
            const message = document.querySelector('.bye-message');
            
            // Показываем сообщение через 2 секунды
            setTimeout(function() {
                message.style.opacity = '1';
            }, 2000);
            
            function createRectangles() {
                for (let i = 0; i < rectCount; i++) {
                    // Левые прямоугольники
                    const leftRect = document.createElement('div');
                    leftRect.className = 'bye-rectangle';
                    leftRect.style.left = '0';
                    leftRect.style.top = (i * (window.innerHeight / rectCount)) + 'px';
                    container.appendChild(leftRect);
                    rects.push(leftRect);
                    
                    // Правые прямоугольники
                    const rightRect = document.createElement('div');
                    rightRect.className = 'bye-rectangle';
                    rightRect.style.right = '0';
                    rightRect.style.top = (i * (window.innerHeight / rectCount)) + 'px';
                    container.appendChild(rightRect);
                    rects.push(rightRect);
                }
            }
            
            function animateToCenter() {
                const centerX = window.innerWidth / 2;
                const duration = 2000;
                const startTime = performance.now();
                
                function update(time) {
                    const elapsed = time - startTime;
                    const progress = Math.min(elapsed / duration, 1);
                    
                    rects.forEach(function(rect, index) {
                        const rectX = parseInt(rect.style.left || 
                                        (window.innerWidth - parseInt(rect.style.right) - 30));
                        const targetX = centerX - 15;
                        const newX = rectX + (targetX - rectX) * progress;
                        
                        rect.style.left = newX + 'px';
                        rect.style.right = 'auto';
                        
                        // Волновой эффект
                        const delay = index * 50;
                        if (elapsed >= delay) {
                            rect.style.opacity = 0.8 - (0.7 * ((elapsed - delay) / (duration - delay)));
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
                const duration = 2000;
                const startTime = performance.now();
                
                function update(time) {
                    const elapsed = time - startTime;
                    const progress = Math.min(elapsed / duration, 1);
                    
                    rects.forEach(function(rect, index) {
                        const currentX = parseInt(rect.style.left);
                        const targetX = currentX < window.innerWidth / 2 ? 
                                        window.innerWidth : 
                                        -30;
                        const newX = currentX + (targetX - currentX) * progress;
                        
                        rect.style.left = newX + 'px';
                        
                        // Возвращаем прозрачность
                        const delay = index * 50;
                        if (elapsed >= delay) {
                            rect.style.opacity = 0.1 + (0.7 * ((elapsed - delay) / (duration - delay)));
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
                    if (index % 2 === 0) {
                        rect.style.left = '0';
                        rect.style.right = 'auto';
                    } else {
                        rect.style.right = '0';
                        rect.style.left = 'auto';
                    }
                    rect.style.top = (Math.floor(index/2) * (window.innerHeight / rectCount)) + 'px';
                    rect.style.opacity = '0.8';
                });
            }
            
            function startAnimation() {
                animateToCenter();
            }
            
            // Инициализация
            createRectangles();
            setTimeout(startAnimation, 500);
        </script>
    </body>
    </html>
    `
	w.Write([]byte(html))
}
