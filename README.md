# 🛰️ Decade Sentinel 2026

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![HTMX](https://img.shields.io/badge/htmx-3366CC?style=for-the-badge&logo=htmx&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

**Decade Sentinel** es una aplicación web autónoma y ultra-ligera que explora la evolución tecnológica entre 2016 y 2026. Construida con **Go** para el motor de backend y **HTMX** para una experiencia de usuario dinámica sin la sobrecarga de pesados frameworks de JavaScript.

## 🚀 Características Clave
- **Binario Autónomo:** Utiliza `go:embed` para incluir HTML, CSS y JS dentro del ejecutable.
- **UX Fluida:** Tooltips dinámicos con telemetría de carga simulada mediante HTMX.
- **Arquitectura Limpia:** Middleware de seguridad integrado (CSP, XSS Protection).
- **Estética Cyberpunk:** Interfaz de terminal de comando optimizada para legibilidad y performance.

## 🛠️ Tecnologías
- **Backend:** Go (Golang) 1.2x
- **Frontend:** HTMX, CSS3 (Variables y Animaciones), HTML5.
- **Despliegue:** Makefile optimizado para entornos Windows/Linux.

## 📦 Instalación y Uso

1. Clona el repositorio:
   ```bash
   git clone [https://github.com/santiagourdaneta/decade-sentinel-2026.git](https://github.com/santiagourdaneta/decade-sentinel-2026.git)
   cd decade-sentinel-2026

2. Ejecuta el entorno de desarrollo (requiere Make):

make dev

3. Abre tu navegador en: http://localhost:8080

🛡️ Seguridad

El servidor incluye un SecurityMiddleware que implementa:

Content Security Policy (CSP) restrictivo.

Protección contra Sniffing de tipos (X-Content-Type-Options).

Prevención de Clickjacking (X-Frame-Options).

Desarrollado con ❤️ por Santiago Urdaneta

#Golang #HTMX #CSS3 #JavaScript #WebDevelopment #SoftwareArchitecture #Sentinel2026 #Coding #CyberpunkUI