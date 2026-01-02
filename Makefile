# Variables
BINARY_NAME=sentinel_pro.exe
PORT=8080

.PHONY: all build run clean dev

all: build run

# Regla de Desarrollo: Limpia, Compila y Corre
dev: clean build run

# Compilar el motor (Go)
build:
	@echo "🔨 Compilando binario autónomo..."
	go build -o $(BINARY_NAME) ./api/main.go

# Encender el motor
run:
	@echo "🚀 Sentinel 2026 activo en http://localhost:$(PORT)"
	./$(BINARY_NAME)

# Limpiar restos (Matar proceso y borrar exe)
clean:
	@echo "🧹 Limpiando procesos y binarios..."
	-taskkill //F //IM $(BINARY_NAME) //T 2>/dev/null || true
	-rm -f $(BINARY_NAME)