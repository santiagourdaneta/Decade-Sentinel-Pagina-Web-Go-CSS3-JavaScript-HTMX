package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"time"
)

// Base de datos de la cronología
var details = map[string]string{
	"2016": "Docker estandariza el envío de software. Se acaba el 'en mi máquina funciona'.",
	"2017": "Aparece el paper 'Attention is All You Need'. Los Transformers cambian el futuro.",
	"2018": "Serverless y Edge Computing. El código se ejecuta cerca del usuario.",
	"2019": "Microservicios masivos. El reto es entender cómo se comunican.",
	"2020": "La nube se vuelve el cimiento de la sociedad ante la crisis global.",
	"2021": "GitHub Copilot: La IA empieza a escribir funciones reales.",
	"2022": "ChatGPT y Stable Diffusion rompen la barrera técnica.",
	"2023": "HTMX revive el Server-Driven UI. Menos complejidad, más velocidad.",
	"2024": "Bases de datos vectoriales: La IA ahora tiene memoria a largo plazo.",
	"2025": "GDPR 2.0 y Edge AI local. Soberanía total del dato.",
	"2026": "Sistemas auto-curables y Zero-Trust real. El software es autónomo.",
}

func SecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		
		// CSP PERMISIVO PARA HTMX (Permite estilos inline y scripts de unpkg)
		w.Header().Set("Content-Security-Policy", 
			"default-src 'self' unpkg.com; " +
			"style-src 'self' 'unsafe-inline' unpkg.com; " + 
			"script-src 'self' 'unsafe-inline' unpkg.com; " +
			"connect-src 'self';")
		
		next.ServeHTTP(w, r)
	})
}

func handleYearDetails(w http.ResponseWriter, r *http.Request) {
	// ⏳ RETRASO ARTIFICIAL: 600ms para que luzca la telemetría
	time.Sleep(300 * time.Millisecond)

	year := r.URL.Query().Get("year")
	val, ok := details[year]

	if !ok {
		slog.Warn("🚨 INTENTO DE ACCESO A AÑO INEXISTENTE", "año", year)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "<div class='error-alert'>⚠️ ERROR: Año %s fuera de los registros.</div>", year)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<strong>> Detalles:  %s:</strong><br>%s", year, val)
}

//go:embed all:static
var staticFiles embed.FS

func main() {
	mux := http.NewServeMux()

	// Servir CSS, JS e Imágenes 
	staticContent, _ := fs.Sub(staticFiles, "static")
	mux.Handle("/css/", http.FileServer(http.FS(staticContent)))
	mux.Handle("/js/", http.FileServer(http.FS(staticContent)))
	mux.Handle("/img/", http.FileServer(http.FS(staticContent)))

	// Ruta principal
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		index, err := staticFiles.ReadFile("static/index.html")
		if err != nil {
			http.Error(w, "Error interno: index.html no encontrado en embed", 500)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(index)
	})

	mux.HandleFunc("/api/details", handleYearDetails)

	fmt.Println("🚀 Sentinel 2026 AUTÓNOMO: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", SecurityMiddleware(mux)))
}